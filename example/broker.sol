pragma solidity >=0.5.6;
pragma experimental ABIEncoderV2;

contract Broker {
    // Only the contract in the whitelist can invoke the Broker for interchain operations.
    mapping(address => int64) whiteList;
    address[] contracts;
    address[] admins;
    string bxhID;
    string appchainID;

    event throwEvent(uint64 index, string dstFullID, string srcFullID, string funcs, string args, string argscb, string argsrb);

    string[] outServicePairs;
    string[] inServicePairs;
    string[] callbackServicePairs;

    mapping(string => uint64) outCounter; // mapping from contract address to out event last index
    mapping(string => mapping(uint64 => uint)) outMessages;
    mapping(string => uint64) inCounter;
    mapping(string => mapping(uint64 => uint)) inMessages;
    mapping(string => uint64) callbackCounter;
    mapping(string => mapping(uint64 => string)) invokeError;
    mapping(string => mapping(uint64 => string)) callbackError;
    mapping(string => uint64) dstRollbackCounter;

    // Authority control. Contracts need to be registered.
    modifier onlyWhiteList {
        require(whiteList[msg.sender] == 1, "Invoker are not in white list");
        _;
    }

    // Authority control. Only the administrator can audit the contract
    modifier onlyAdmin {
        bool flag = false;
        for (uint i = 0; i < admins.length; i++) {
            if (msg.sender == admins[i]) {flag = true;}
        }
        if (flag) {revert();}
        _;
    }

    function initialize() public {
        for (uint i = 0; i < inServicePairs.length; i++) {
            inCounter[inServicePairs[i]] = 0;
        }
        for (uint j = 0; j < outServicePairs.length; j++) {
            outCounter[outServicePairs[j]] = 0;
        }
        for (uint k = 0; k < callbackServicePairs.length; k++) {
            callbackCounter[callbackServicePairs[k]] = 0;
        }
        for (uint m = 0; m < inServicePairs.length; m++) {
            dstRollbackCounter[inServicePairs[m]] = 0;
        }
        for (uint n = 0; n < contracts.length; n++) {
            whiteList[contracts[n]] = 0;
        }
        delete outServicePairs;
        delete inServicePairs;
        delete callbackServicePairs;
    }

    // 0: auditting  1: approved  -1: refused
    function register(address addr) public {
        whiteList[addr] = 0;
    }

    function audit(address addr, int64 status) public returns (bool) {
        if (status != - 1 && status != 0 && status != 1) {return false;}
        whiteList[addr] = status;
        // Only approved contracts can be recorded
        if (status == 1) {
            contracts.push(addr);
        }
        return true;
    }

    constructor(string memory _bxhID, string memory _appchainID) public {
        bxhID = _bxhID;
        appchainID = _appchainID;
    }

    function invokeInterchain(string calldata srcChainServiceID, uint64 index, address destAddr, uint64 reqType, bytes calldata bizCallData) payable external {
        require(whiteList[destAddr] == 1);
        invokeIndexUpdate(srcChainServiceID, index, destAddr, reqType, "");

        assembly {
            let ptr := mload(0x40)

        // 获取bizCallData在calldata中的偏移
            calldatacopy(ptr, 132, 32)
            let off := mload(ptr)

        // 获取bizCallData的大小
            calldatacopy(ptr, add(4, off), 32)
            let datasize := mload(ptr)

        // 将bizCallData的内容copy到ptr所指的内存
            calldatacopy(ptr, add(36, off), datasize)

        // 调用业务合约
            let result := call(gas(), destAddr, callvalue(), ptr, datasize, 0, 0)
            let size := returndatasize()
            returndatacopy(ptr, 0, size)

            switch result
            case 0 {revert(ptr, size)}
            default {
                log0(ptr, size)
                return (ptr, size)
            }
        }
    }

    function invokeIndexUpdate(string memory srcChainServiceID, uint64 index, address destAddr, uint64 reqType, string memory err) private {
        string memory curServiceID = genFullServiceID(addressToString(destAddr));

        if (reqType == 0) {
            string memory inServicePair = genServicePair(srcChainServiceID, curServiceID);
            require(inCounter[inServicePair] + 1 == index);
            markInCounter(inServicePair);
            if (keccak256(abi.encodePacked(err)) != keccak256(abi.encodePacked(""))) {
                invokeError[inServicePair][index] = err;
            }
        } else if (reqType == 1) {
            string memory outServicePair = genServicePair(curServiceID, srcChainServiceID);
            // invoke src callback or rollback
            require(callbackCounter[outServicePair] + 1 == index);
            markCallbackCounter(outServicePair, index);
            if (keccak256(abi.encodePacked(err)) != keccak256(abi.encodePacked(""))) {
                callbackError[outServicePair][index] = err;
            }
        } else if (reqType == 2) {
            string memory inServicePair = genServicePair(srcChainServiceID, curServiceID);
            // invoke dst rollback
            require(dstRollbackCounter[inServicePair] + 1 <= index);
            markDstRollbackCounter(inServicePair, index);
        }
    }

    function invokeIndexUpdateWithError(string memory srcChainServiceID, uint64 index, address destAddr, uint64 reqType, string memory err) public {
        invokeIndexUpdate(srcChainServiceID, index, destAddr, reqType, err);
    }

    function emitInterchainEvent(
        string memory destChainServiceID,
        string memory funcs,
        string memory args,
        string memory argscb,
        string memory argsrb)
    public onlyWhiteList {
        string memory curFullID = genFullServiceID(addressToString(msg.sender));
        string memory outServicePair = genServicePair(curFullID, destChainServiceID);

        // Record the order of interchain contract which has been started.
        outCounter[outServicePair]++;
        if (outCounter[outServicePair] == 1) {
            outServicePairs.push(outServicePair);
        }
        outMessages[outServicePair][outCounter[outServicePair]] = block.number;

        // Throw interchain event for listening of plugin.
        emit throwEvent(outCounter[outServicePair], destChainServiceID, curFullID, funcs, args, argscb, argsrb);
    }
    
    

    // The helper functions that help document Meta information.
    function markCallbackCounter(string memory servicePair, uint64 index) private {
        if (callbackCounter[servicePair] == 0) {
            callbackServicePairs.push(servicePair);
        }
        callbackCounter[servicePair] = index;
    }

    function markDstRollbackCounter(string memory servicePair, uint64 index) private {
        dstRollbackCounter[servicePair] = index;
    }

    function markInCounter(string memory servicePair) private {
        inCounter[servicePair]++;
        if (inCounter[servicePair] == 1) {
            inServicePairs.push(servicePair);
        }

        inMessages[servicePair][inCounter[servicePair]] = block.number;
    }

    // The helper functions that help plugin query.
    function getOuterMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](outServicePairs.length);
        for (uint64 i = 0; i < outServicePairs.length; i++) {
            indices[i] = outCounter[outServicePairs[i]];
        }

        return (outServicePairs, indices);
    }

    function getOutMessage(string memory outServicePair, uint64 idx) public view returns (uint) {
        return outMessages[outServicePair][idx];
    }

    function getInMessage(string memory inServicePair, uint64 idx) public view returns (uint)  {
        return inMessages[inServicePair][idx];
    }

    function getInnerMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](inServicePairs.length);
        for (uint i = 0; i < inServicePairs.length; i++) {
            indices[i] = inCounter[inServicePairs[i]];
        }

        return (inServicePairs, indices);
    }

    function getCallbackMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](callbackServicePairs.length);
        for (uint64 i = 0; i < callbackServicePairs.length; i++) {
            indices[i] = callbackCounter[callbackServicePairs[i]];
        }

        return (callbackServicePairs, indices);
    }

    function getDstRollbackMeta() public view returns(string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](inServicePairs.length);
        for (uint i = 0; i < inServicePairs.length; i++) {
            indices[i] = dstRollbackCounter[inServicePairs[i]];
        }

        return (inServicePairs, indices);
    }

    function genFullServiceID(string memory serviceID) public view returns (string memory) {
        string memory fullID = concat(toSlice(bxhID), toSlice(":"));
        fullID = concat(toSlice(fullID), toSlice(appchainID));
        fullID = concat(toSlice(fullID), toSlice(":"));
        fullID = concat(toSlice(fullID), toSlice(serviceID));
        return fullID;
    }

    function genServicePair(string memory from, string memory to) internal pure returns (string memory) {
        string memory servicePair = concat(toSlice(from), toSlice("-"));
        servicePair = concat(toSlice(servicePair), toSlice(to));

        return servicePair;
    }

    function getChainID() public view returns (string memory, string memory) {
        return (bxhID, appchainID);
    }

    function _indexOf(string memory _base, string memory _value, uint _offset) internal pure returns (int) {
        bytes memory _baseBytes = bytes(_base);
        bytes memory _valueBytes = bytes(_value);

        assert(_valueBytes.length == 1);

        for (uint i = _offset; i < _baseBytes.length; i++) {
            if (_baseBytes[i] == _valueBytes[0]) {
                return int(i);
            }
        }

        return - 1;
    }

    struct slice {
        uint _len;
        uint _ptr;
    }

    function toSlice(string memory self) internal pure returns (slice memory) {
        uint ptr;
        assembly {
            ptr := add(self, 0x20)
        }
        return slice(bytes(self).length, ptr);
    }

    function concat(slice memory self, slice memory other) internal pure returns (string memory) {
        string memory ret = new string(self._len + other._len);
        uint retptr;
        assembly { retptr := add(ret, 32) }
        memcpy(retptr, self._ptr, self._len);
        memcpy(retptr + self._len, other._ptr, other._len);
        return ret;
    }

    function memcpy(uint dest, uint src, uint len) private pure {
        // Copy word-length chunks while possible
        for(; len >= 32; len -= 32) {
            assembly {
                mstore(dest, mload(src))
            }
            dest += 32;
            src += 32;
        }

        // Copy remaining bytes
        uint mask = 256 ** (32 - len) - 1;
        assembly {
            let srcpart := and(mload(src), not(mask))
            let destpart := and(mload(dest), mask)
            mstore(dest, or(destpart, srcpart))
        }
    }

    // function addressToString(address x) internal pure returns (string memory) {
    //     bytes memory s = new bytes(40);
    //     for (uint i = 0; i < 20; i++) {
    //         bytes1 b = bytes1(uint8(uint(uint160(x)) / (2**(8*(19 - i)))));
    //         bytes1 hi = bytes1(uint8(b) / 16);
    //         bytes1 lo = bytes1(uint8(b) - 16 * uint8(hi));
    //         s[2*i] = char(hi);
    //         s[2*i+1] = char(lo);
    //     }
    //     return string(s);
    // }
    
    
  function _getAsciiOffset(
    uint8 nibble, bool caps
  ) internal pure returns (uint8 offset) {
    // to convert to ascii characters, add 48 to 0-9, 55 to A-F, & 87 to a-f.
    if (nibble < 10) {
      offset = 48;
    } else if (caps) {
      offset = 55;
    } else {
      offset = 87;
    }
  }
    
    function addressToString(
    address account
  ) public pure returns (string memory asciiString) {
    // convert the account argument from address to bytes.
    bytes20 data = bytes20(account);

    // create an in-memory fixed-size bytes array.
    bytes memory asciiBytes = new bytes(40);

    // declare variable types.
    uint8 b;
    uint8 leftNibble;
    uint8 rightNibble;
    bool leftCaps;
    bool rightCaps;
    uint8 asciiOffset;

    // get the capitalized characters in the actual checksum.
    bool[40] memory caps = _toChecksumCapsFlags(account);

    // iterate over bytes, processing left and right nibble in each iteration.
    for (uint256 i = 0; i < data.length; i++) {
      // locate the byte and extract each nibble.
      b = uint8(uint160(data) / (2**(8*(19 - i))));
      leftNibble = b / 16;
      rightNibble = b - 16 * leftNibble;

      // locate and extract each capitalization status.
      leftCaps = caps[2*i];
      rightCaps = caps[2*i + 1];

      // get the offset from nibble value to ascii character for left nibble.
      asciiOffset = _getAsciiOffset(leftNibble, leftCaps);

      // add the converted character to the byte array.
      asciiBytes[2 * i] = byte(leftNibble + asciiOffset);

      // get the offset from nibble value to ascii character for right nibble.
      asciiOffset = _getAsciiOffset(rightNibble, rightCaps);

      // add the converted character to the byte array.
      asciiBytes[2 * i + 1] = byte(rightNibble + asciiOffset);
    }

   
    return concat(toSlice("0x"), toSlice(string(asciiBytes)));
  }

  function _toChecksumCapsFlags(address account) internal pure returns (
    bool[40] memory characterCapitalized
  ) {
    // convert the address to bytes.
    bytes20 a = bytes20(account);

    // hash the address (used to calculate checksum).
    bytes32 b = keccak256(abi.encodePacked(_toAsciiString(a)));

    // declare variable types.
    uint8 leftNibbleAddress;
    uint8 rightNibbleAddress;
    uint8 leftNibbleHash;
    uint8 rightNibbleHash;

    // iterate over bytes, processing left and right nibble in each iteration.
    for (uint256 i; i < a.length; i++) {
      // locate the byte and extract each nibble for the address and the hash.
      rightNibbleAddress = uint8(a[i]) % 16;
      leftNibbleAddress = (uint8(a[i]) - rightNibbleAddress) / 16;
      rightNibbleHash = uint8(b[i]) % 16;
      leftNibbleHash = (uint8(b[i]) - rightNibbleHash) / 16;

      characterCapitalized[2 * i] = (
        leftNibbleAddress > 9 &&
        leftNibbleHash > 7
      );
      characterCapitalized[2 * i + 1] = (
        rightNibbleAddress > 9 &&
        rightNibbleHash > 7
      );
    }
  }
  
  // based on https://ethereum.stackexchange.com/a/56499/48410
  function _toAsciiString(
    bytes20 data
  ) internal pure returns (string memory asciiString) {
    // create an in-memory fixed-size bytes array.
    bytes memory asciiBytes = new bytes(40);

    // declare variable types.
    uint8 b;
    uint8 leftNibble;
    uint8 rightNibble;

    // iterate over bytes, processing left and right nibble in each iteration.
    for (uint256 i = 0; i < data.length; i++) {
      // locate the byte and extract each nibble.
      b = uint8(uint160(data) / (2 ** (8 * (19 - i))));
      leftNibble = b / 16;
      rightNibble = b - 16 * leftNibble;

      // to convert to ascii characters, add 48 to 0-9 and 87 to a-f.
      asciiBytes[2 * i] = byte(leftNibble + (leftNibble < 10 ? 48 : 87));
      asciiBytes[2 * i + 1] = byte(rightNibble + (rightNibble < 10 ? 48 : 87));
    }

    return string(asciiBytes);
  }


    function char(bytes1 b) internal pure returns (bytes1 c) {
        if (uint8(b) < 10) return bytes1(uint8(b) + 0x30);
        else return bytes1(uint8(b) + 0x57);
    }
}