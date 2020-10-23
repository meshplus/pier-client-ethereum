pragma solidity >=0.5.7;

contract AssetExchange {
    event InterchainAssetExchangeInitNotify(
        address srcChain,
        string srcAddr,
        string assetExchangeID,
        string senderOnSrcChain,
        string receiverOnSrcChain,
        uint64 assetOnSrcChain,
        string senderOnDstChain,
        string receiverOnDstChain,
        uint64 assetOnDstChain);

    struct ExchangeData {
        string sender;
        string receiver;
        uint64 ammount;
        bool finished;
    }

    mapping(address => bool) addrs;
    address[] addrList;
    mapping(address => bool) bxhValidators;
    uint32 validatorSize = 0;
    mapping(string => uint64) accountM; // map for accounts
    mapping(string => ExchangeData) exchangeDataM;
    uint64 assetExchangeIndex = 0;

    address AdminAddr; // admin address
    // change the address of Broker accordingly
    //    address BrokerAddr = 0x9E0901D698E854F6CFE9e478C38d20A01908768a;
    address BrokerAddr = 0xD3880ea40670eD51C3e3C0ea089fDbDc9e3FBBb4;
    Broker broker = Broker(BrokerAddr);

    // AccessControl
    modifier onlyBroker {
        require(msg.sender == BrokerAddr, "Invoker are not the Broker");
        _;
    }

    modifier onlyAdmin {
        require(msg.sender == AdminAddr, "Invoker are not the Admin");
        _;
    }

    constructor() public {
        AdminAddr = msg.sender;
    }

    // 资产交换类的业务合约
    function assetExchangeInit(address destChainID, string memory destAddr, string memory srcAddr,
        string memory senderOnSrcChain, string memory receiverOnSrcChain, string memory assetOnSrcChain,
        string memory senderOnDstChain, string memory receiverOnDstChain, string memory assetOnDstChain) public {
        uint64 am = uint64(parseInt(assetOnSrcChain));
        require(accountM[senderOnSrcChain] >= am, "not sufficient funds");

        accountM[senderOnSrcChain] -= am;
        accountM[toString(BrokerAddr)] += am;

        string memory assetExchangeID = calcassetExchangeID();
        require(exchangeDataM[assetExchangeID].finished == false, "this asset exchange is finished");
        exchangeDataM[assetExchangeID] = ExchangeData(senderOnSrcChain, receiverOnSrcChain, am, false);

        // 拼接参数
        string memory args = concat(toSlice(srcAddr), toSlice(","));
        args = concat(toSlice(args), toSlice(assetExchangeID));
        args = concat(toSlice(args), toSlice(","));
        args = concat(toSlice(args), toSlice(senderOnSrcChain));
        args = concat(toSlice(args), toSlice(","));
        args = concat(toSlice(args), toSlice(receiverOnSrcChain));
        args = concat(toSlice(args), toSlice(","));
        args = concat(toSlice(args), toSlice(assetOnSrcChain));
        args = concat(toSlice(args), toSlice(","));
        args = concat(toSlice(args), toSlice(senderOnDstChain));
        args = concat(toSlice(args), toSlice(","));
        args = concat(toSlice(args), toSlice(receiverOnDstChain));
        args = concat(toSlice(args), toSlice(","));
        args = concat(toSlice(args), toSlice(assetOnDstChain));

        bool ok = broker.InterchainAssetExchangeInvoke(destChainID, destAddr, args, 0);
        require(ok);
    }

    function assetExchangeRedeem(address destChainID, string memory destAddr, string memory assetExchangeID,
        string memory senderOnSrcChain, string memory receiverOnSrcChain, string memory assetOnSrcChain) public {
        uint64 am = uint64(parseInt(assetOnSrcChain));
        require(accountM[senderOnSrcChain] >= am, "not sufficient funds");
        require(exchangeDataM[assetExchangeID].finished == false, "this asset exchange is finished");

        accountM[senderOnSrcChain] -= am;
        accountM[toString(BrokerAddr)] += am;
        exchangeDataM[assetExchangeID] = ExchangeData(senderOnSrcChain, receiverOnSrcChain, am, false);

        bool ok = broker.InterchainAssetExchangeInvoke(destChainID, destAddr, assetExchangeID, 1);
        require(ok);
    }

    function assetExchangeRefund(address destChainID, string memory destAddr, string memory assetExchangeID) public {
        require(exchangeDataM[assetExchangeID].finished == false, "this asset exchange is finished");
        bool ok = broker.InterchainAssetExchangeInvoke(destChainID, destAddr, assetExchangeID, 2);
        require(ok);
    }

    function interchainAssetExchangeInit(
        address srcChainID,
        string memory srcAddr,
        string memory assetExchangeID,
        string memory senderOnSrcChain,
        string memory receiverOnSrcChain,
        uint64 assetOnSrcChain,
        string memory senderOnDstChain,
        string memory receiverOnDstChain,
        uint64 assetOnDstChain) public onlyBroker returns (bool) {
        emit InterchainAssetExchangeInitNotify(
            srcChainID,
            srcAddr,
            assetExchangeID,
            senderOnSrcChain,
            receiverOnSrcChain,
            assetOnSrcChain,
            senderOnDstChain,
            receiverOnDstChain,
            assetOnDstChain);

        return true;
    }

    function interchainAssetExchangeFinish(string memory assetExchangeID, string memory status, string memory signatures) public onlyBroker returns (bool) {
        bool result = false;

        if (!exchangeDataM[assetExchangeID].finished) {

            if (verifySignatures(assetExchangeID, status, signatures)) {
                ExchangeData memory data = exchangeDataM[assetExchangeID];

                if (keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("1"))) {
                    result = interchainRedeem(data.receiver, data.ammount);
                } else if (keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("2"))) {
                    result = interchainRefund(data.sender, data.ammount);
                }

                exchangeDataM[assetExchangeID].finished = true;
            }
        }

        return result;
    }

    function interchainAssetExchangeConfirm(string memory assetExchangeID, string memory signatures) public onlyBroker returns (bool) {
        bool result = false;

        if (!exchangeDataM[assetExchangeID].finished) {
            ExchangeData memory data = exchangeDataM[assetExchangeID];

            if (verifySignatures(assetExchangeID, "1", signatures)) {
                result = interchainRedeem(data.receiver, data.ammount);
                exchangeDataM[assetExchangeID].finished = true;
            } else if (verifySignatures(assetExchangeID, "2", signatures)) {
                result = interchainRefund(data.sender, data.ammount);
                exchangeDataM[assetExchangeID].finished = true;
            }
        }

        return result;
    }

    function interchainRedeem(string memory receiver, uint64 amount) internal returns (bool) {
        accountM[toString(BrokerAddr)] -= amount;
        accountM[receiver] += amount;
        return true;
    }

    function interchainRefund(string memory sender, uint64 amount) internal returns (bool) {
        accountM[toString(BrokerAddr)] -= amount;
        accountM[sender] += amount;
        return true;
    }

    function getBalance(string memory id) public view returns (uint64) {
        return accountM[id];
    }

    function setBalance(string memory id, uint64 amount) public {
        accountM[id] = amount;
    }

    function setBxhValidator(address validator) public onlyAdmin {
        if (bxhValidators[validator] == false) {
            bxhValidators[validator] = true;
            validatorSize++;
        }
    }

    function getBxhValidatorSize() public view returns (uint64) {
        return validatorSize;
    }

    function parseInt(string memory self) internal pure returns (uint _ret) {
        bytes memory _bytesValue = bytes(self);
        uint j = 1;
        for (uint i = _bytesValue.length - 1; i >= 0 && i < _bytesValue.length; i--) {
            assert(uint8(_bytesValue[i]) >= 48 && uint8(_bytesValue[i]) <= 57);
            _ret += (uint8(_bytesValue[i]) - 48) * j;
            j *= 10;
        }
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
        assembly {retptr := add(ret, 32)}
        memcpy(retptr, self._ptr, self._len);
        memcpy(retptr + self._len, other._ptr, other._len);
        return ret;
    }

    function memcpy(uint dest, uint src, uint len) private pure {
        // Copy word-length chunks while possible
        for (; len >= 32; len -= 32) {
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

    function verifySignatures(string memory assetExchangeID, string memory status, string memory signatures) internal  returns (bool){
        uint32 sigLen = 65;
        string memory message = concat(toSlice(assetExchangeID), toSlice("-"));
        message = concat(toSlice(message), toSlice(status));
        bytes memory bytesMsg = bytes(message);
        bytes32 sha256Msg = sha256(bytesMsg);
        bytes memory bytesSignatures = bytes(signatures);
        uint32 threshold = (validatorSize - 1) / 3;
        uint32 count = 0;

        for (uint i = 0; i * sigLen < bytesSignatures.length; i++) {
            bytes memory bytesSign = new bytes(sigLen);
            for (uint j = 0; j < sigLen; j++) {
                bytesSign[j] = bytesSignatures[i * sigLen + j];
            }

            (uint8 v, bytes32 r, bytes32 s) = splitSignature(bytesSign);

            address addr = ecrecover(sha256Msg, v, r, s);
            if (addrs[addr]) {
                continue;
            }

            if (bxhValidators[addr]) {
                addrs[addr] = true;
                count++;
                addrList.push(addr);
            }
        }

        for (uint i = 0; i < addrList.length; i++) {
            addrs[addrList[i]] = false;
        }
        delete addrList;

        if (count > threshold) {
            return true;
        }

        return false;
    }

    function splitSignature(bytes memory sig) internal pure returns (bytes32 r, bytes32 s) {
        require(sig.length == 65);

        assembly {
            // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
            // second 32 bytes
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        return (v, r, s);
    }

    function toString(address _addr) public pure returns (string memory) {
        bytes32 value = bytes32(uint256(_addr));
        bytes memory alphabet = "0123456789abcdef";

        bytes memory str = new bytes(42);
        str[0] = '0';
        str[1] = 'x';
        for (uint256 i = 0; i < 20; i++) {
            str[2+i*2] = alphabet[uint8(value[i + 12] >> 4)];
            str[3+i*2] = alphabet[uint8(value[i + 12] & 0x0f)];
        }
        return string(str);
    }

    function calcassetExchangeID() public returns (string memory) {
        string memory curAddr = toString(address(this));
        bytes memory m = new bytes(bytes(curAddr).length + 32 + 8);

        uint counter = 0;
        for (uint64 i = 0; i < bytes(curAddr).length; i++) {
            m[counter] = bytes(curAddr)[i];
            counter++;
        }

        for (uint64 i = 0; i < 32; i++) {
            m[counter] = byte(uint8(now >> (8 * i) & uint(255)));
            counter++;
        }

        for (uint64 i = 0; i < 8; i++) {
            m[counter] = byte(uint8(assetExchangeIndex >> (8 * i) & uint64(255)));
            counter++;
        }

        bytes32 hash = sha256(m);

        assetExchangeIndex++;

        return bytes32ToHexString(hash);

    }

    function bytes32ToHexString(bytes32 x) private pure returns (string memory) {
        bytes memory str = new bytes(64);
        bytes memory hextable = bytes("0123456789abcdef");
        for (uint i = 0; i < 32; i++) {
            uint8 b = uint8(x[i]);
            str[i * 2] = hextable[b >> 4];
            str[i * 2 + 1] = hextable[b & 0x0f];
        }

        return string(str);
    }
}

contract Broker {
    function InterchainAssetExchangeInvoke(address destChainID, string memory destAddr, string memory args, uint64 typ) public returns (bool);
}
