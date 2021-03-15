pragma solidity >=0.5.6;
pragma experimental ABIEncoderV2;

contract Broker {
    // Only the contract in the whitelist can invoke the Broker for interchain operations.
    mapping(address => int64) whiteList;
    address[] contracts;
    address[] admins;

    event throwEvent(uint64 index, string destDID, address fid, string funcs, string args, string argscb, string argsrb);
    event LogInterchainData(bool status, string data);
    event LogInterchainStatus(bool status);

    string[] outChains;
    string[] inChains;
    string[] callbackChains;

    mapping(string => uint64) outCounter; // mapping from contract address to out event last index
    mapping(string => mapping(uint64 => uint)) outMessages;
    mapping(string => uint64) inCounter;
    mapping(string => mapping(uint64 => uint)) inMessages;
    mapping(string => uint64) callbackCounter;
    mapping(string => mapping(uint64 => string)) invokeError;
    mapping(string => mapping(uint64 => string)) callbackError;

    // Authority control. Contracts need to be registered.
    modifier onlyWhiteList {
        require(whiteList[msg.sender]==1, "Invoker are not in white list");
        _;
    }

    // Authority control. Only the administrator can audit the contract
    modifier onlyAdmin {
        bool flag = false;
        for (uint i = 0; i < admins.length; i++) {
            if (msg.sender == admins[i]) { flag = true; }
        }
        if (flag) { revert(); }
        _;
    }

    function initialize() public {
        for (uint i = 0; i < inChains.length; i++) {
            inCounter[inChains[i]] = 0;
        }
        for (uint i = 0; i < outChains.length; i++) {
            outCounter[outChains[i]] = 0;
        }
        for (uint i = 0; i < callbackChains.length; i++) {
            callbackCounter[callbackChains[i]] = 0;
        }
        for (uint i = 0; i < contracts.length; i++) {
            whiteList[contracts[i]] = 0;
        }
        delete outChains;
        delete inChains;
        delete callbackChains;
    }

    // 0: auditting  1: approved  -1: refused
    function register(address addr) public {
        whiteList[addr] = 0;
    }

    function audit(address addr, int64 status) public returns(bool) {
        if (status != -1 && status != 0 && status != 1) { return false; }
        whiteList[addr] = status;
        // Only approved contracts can be recorded
        if (status == 1) {
            contracts.push(addr);
        }
        return true;
    }

    function invokeInterchain(string calldata srcChainMethod, uint64 index, address destAddr, bool req, bytes calldata bizCallData) payable external {
        require(whiteList[destAddr] == 1);
        invokeIndexUpdate(srcChainMethod, index, req, "");

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

    function invokeIndexUpdate(string memory srcChainMethod, uint64 index, bool req, string memory err) private {
        if (req) {
            require(inCounter[srcChainMethod] + 1 == index);
            markInCounter(srcChainMethod);
            if (keccak256(abi.encodePacked(err)) != keccak256(abi.encodePacked(""))) {
                invokeError[srcChainMethod][index] = err;
            }
        } else {
            // invoke callback or rollback
            require(callbackCounter[srcChainMethod] + 1 == index);
            markCallbackCounter(srcChainMethod, index);
            if (keccak256(abi.encodePacked(err)) != keccak256(abi.encodePacked(""))) {
                callbackError[srcChainMethod][index] = err;
            }
        }
    }

    function invokeIndexUpdateWithError(string memory srcChainMethod, uint64 index, bool req, string memory err) public {
        invokeIndexUpdate(srcChainMethod, index, req, err);
    }

    function emitInterchainEvent(
        string memory destContractDID,
        string memory funcs,
        string memory args,
        string memory argscb,
        string memory argsrb)
    public onlyWhiteList {
        // Record the order of interchain contract which has been started.
        string memory destChainMethod = parseMethod(destContractDID);
        outCounter[destChainMethod]++;
        if (outCounter[destChainMethod] == 1) {
            outChains.push(destChainMethod);
        }
        outMessages[destChainMethod][outCounter[destChainMethod]] = block.number;

        // Throw interchain event for listening of plugin.
        emit throwEvent(outCounter[destChainMethod], destContractDID, msg.sender, funcs, args, argscb, argsrb);
    }

    // The helper functions that help document Meta information.
    function markCallbackCounter(string memory from, uint64 index) private {
        if (callbackCounter[from] == 0) {
            callbackChains.push(from);
        }
        callbackCounter[from] = index;
        inMessages[from][callbackCounter[from]] = block.number;
    }

    function markInCounter(string memory from) private {
        inCounter[from]++;
        if (inCounter[from] == 1) {
            inChains.push(from);
        }

        inMessages[from][inCounter[from]] = block.number;
    }

    // The helper functions that help plugin query.
    function getOuterMeta() public view returns(string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](outChains.length);
        for (uint64 i = 0; i < outChains.length; i++) {
            indices[i] = outCounter[outChains[i]];
        }

        return (outChains, indices);
    }

    function getOutMessage(string memory to, uint64 idx) public view returns (uint) {
        validDID(to);
        return outMessages[to][idx];
    }

    function getInMessage(string memory from, uint64 idx) public view returns (uint)  {
        validDID(from);
        return inMessages[from][idx];
    }

    function getInnerMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](inChains.length);
        for (uint i = 0; i < inChains.length; i++) {
            indices[i] = inCounter[inChains[i]];
        }

        return (inChains, indices);
    }

    function getCallbackMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](callbackChains.length);
        for (uint64 i = 0; i < callbackChains.length; i++) {
            indices[i] = callbackCounter[callbackChains[i]];
        }

        return (callbackChains, indices);
    }

    function validDID(string memory did) internal pure {
        string[] memory splitArr = split(did, ":");
        require(splitArr.length==4, "did is not in four part");
        require(compareStrings(splitArr[0], "did"), "prefix is not did");
        for (uint256 i=1;i<4;i++) {
            bytes memory tmp = bytes(splitArr[i]);
            require(tmp.length!=0, "did subset is empty");
        }
    }

    function parseMethod(string memory did) internal pure returns(string memory) {
        string[] memory splitArr = split(did, ":");
        // check did format
        require(splitArr.length==4, "did is not legal format");
        require(compareStrings(splitArr[0], "did"), "did is not legal format");
        for (uint256 i=1;i<4;i++) {
            bytes memory tmp = bytes(splitArr[i]);
            require(tmp.length!=0, "did subset is empty");
        }

        string memory method = concat(toSlice(splitArr[0]), toSlice(":"));
        method = concat(toSlice(method), toSlice(splitArr[1]));
        method = concat(toSlice(method), toSlice(":"));
        method = concat(toSlice(method), toSlice(splitArr[2]));
        method = concat(toSlice(method), toSlice(":"));
        method = concat(toSlice(method), toSlice("."));
        return method;
    }

    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
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

    function split(string memory _base, string memory _delimiter) internal pure returns (string[] memory splitArr) {
        bytes memory _baseBytes = bytes(_base);

        uint _offset = 0;
        uint _splitsCount = 1;
        while (_offset < _baseBytes.length - 1) {
            int _limit = _indexOf(_base, _delimiter, _offset);
            if (_limit == - 1)
                break;
            else {
                _splitsCount++;
                _offset = uint(_limit) + 1;
            }
        }

        splitArr = new string[](_splitsCount);

        _offset = 0;
        _splitsCount = 0;
        while (_offset <= _baseBytes.length - 1) {
            int _limit = _indexOf(_base, _delimiter, _offset);
            if (_limit == - 1) {
                _limit = int(_baseBytes.length);
            }

            string memory _tmp = new string(uint(_limit) - _offset);
            bytes memory _tmpBytes = bytes(_tmp);

            uint j = 0;
            for (uint i = _offset; i < uint(_limit); i++) {
                _tmpBytes[j++] = _baseBytes[i];
            }
            _offset = uint(_limit) + 1;
            splitArr[_splitsCount++] = string(_tmpBytes);
        }
        return splitArr;
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

        return -1;
    }
}