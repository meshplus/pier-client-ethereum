pragma solidity >=0.6.9 <=0.7.6;
pragma experimental ABIEncoderV2;

contract Broker {
    struct Proposal {
        uint64 approve;
        uint64 reject;
        address[] votedAdmins;
        bool ordered;
        bool exist;
    }

    struct CallFunc {
        string func;
        bytes[] args;
    }

    struct InterchainInvoke {
        bool encrypt;
        string[] group;
        CallFunc callFunc;
        CallFunc callback;
        CallFunc rollback;
    }

    struct Receipt {
        bool encrypt;
        uint64 typ;
        bytes[][] results;
        bool[] multiStatus;
    }

    struct multiInvokeArgs {
        string contractAddr;
        CallFunc invokeFunc;
        bytes[] arg;
    }

    // Only the contract in the whitelist can invoke the Broker for interchain operations.
    mapping(address => bool) localWhiteList;
    address[] localServices;
    mapping(address => Proposal) localServiceProposal;
    address[] proposalList;
    mapping(address => bool) serviceOrdered;

    string bitxhubID;
    string appchainID;
    address[] public validators;
    uint64 public valThreshold;
    address[] public admins;
    uint64 public adminThreshold;

    address[] bxhSigners;

    event throwInterchainEvent(uint64 index, string dstFullID, string srcFullID, string func, bytes[] args, bytes32 hash, string[] group);
    event throwReceiptEvent(uint64 index, string dstFullID, string srcFullID, uint64 typ, bytes[][] results, bytes32 hash, bool[] multiStatus);
    event throwReceiptStatus(bool);

    string[] outServicePairs;
    string[] inServicePairs;
    string[] callbackServicePairs;

    mapping(string => uint64) outCounter;
    mapping(string => uint64) callbackCounter;
    mapping(string => uint64) inCounter;
    mapping(string => uint64) dstRollbackCounter;

    mapping(string => mapping(uint64 => InterchainInvoke)) outMessages;
    mapping(string => mapping(uint64 => Receipt)) receiptMessages;

    // Authority control. Contracts need to be registered.
    modifier onlyWhiteList {
        require(localWhiteList[msg.sender] == true, "Invoker are not in white list");
        _;
    }

    // Authority control. Only the administrator can audit the contract
    modifier onlyAdmin {
        bool flag = false;
        for (uint i = 0; i < admins.length; i++) {
            if (msg.sender == admins[i]) {flag = true;}
        }

        require(flag == true, "Invoker are not in admin list");
        _;
    }

    constructor(string memory _bitxhubID,
        string memory _appchainID,
        address[] memory _validators,
        uint64 _valThreshold,
        address[] memory _admins,
        uint64 _adminThreshold) {
        bitxhubID = _bitxhubID;
        appchainID = _appchainID;
        validators = _validators;
        valThreshold = _valThreshold;
        admins = _admins;
        adminThreshold = _adminThreshold;
    }

    function setAdmins(address[] memory _admins, uint64 _adminThreshold) public onlyAdmin {
        admins = _admins;
        adminThreshold = _adminThreshold;
    }

    function setValidators(address[] memory _validators, uint64 _valThreshold) public onlyAdmin {
        validators = _validators;
        valThreshold = _valThreshold;
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
        for (uint n = 0; n < localServices.length; n++) {
            localWhiteList[localServices[n]] = false;
        }
        for (uint x = 0; x < proposalList.length; x++) {
            delete localServiceProposal[proposalList[x]];
        }
        delete outServicePairs;
        delete inServicePairs;
        delete callbackServicePairs;
        delete localServices;
    }

    // register local service to Broker
    function register(bool ordered) public {
        require(tx.origin != msg.sender, "register not by contract");
        if (localWhiteList[msg.sender] || localServiceProposal[msg.sender].exist) {
            return;
        }

        localServiceProposal[msg.sender] = Proposal(0, 0, new address[](admins.length), ordered, true);
    }

    function audit(address addr, int64 status) public onlyAdmin returns (bool) {
        uint result = vote(addr, status);

        if (result == 0) {
            return false;
        }

        if (result == 1) {
            bool ordered = localServiceProposal[addr].ordered;
            delete localServiceProposal[addr];
            localWhiteList[addr] = true;
            serviceOrdered[addr] = ordered;
            localServices.push(addr);
        } else {
            delete localServiceProposal[addr];
        }

        return true;
    }

    // return value explain:
    // 0: vote is not finished
    // 1: approve the proposal
    // 2: reject the proposal
    function vote(address addr, int64 status) private returns (uint) {
        require(localServiceProposal[addr].exist, "the proposal does not exist");
        require(status == 0 || status == 1, "vote status should be 0 or 1");

        for (uint i = 0; i < localServiceProposal[addr].votedAdmins.length; i++) {
            require(localServiceProposal[addr].votedAdmins[i] != msg.sender, "current use has voted the proposal");
        }

        localServiceProposal[addr].votedAdmins[localServiceProposal[addr].reject + localServiceProposal[addr].approve] = msg.sender;
        if (status == 0) {
            localServiceProposal[addr].reject++;
            if (localServiceProposal[addr].reject == admins.length - adminThreshold + 1) {
                return 2;
            }
        } else {
            localServiceProposal[addr].approve++;
            if (localServiceProposal[addr].approve == adminThreshold) {
                return 1;
            }
        }

        return 0;
    }

    // get the registered local service list
    function getLocalServiceList() public view returns (string[] memory) {
        string[] memory fullServiceIDList = new string[](localServices.length);
        for (uint i = 0; i < localServices.length; i++) {
            fullServiceIDList[i] = genFullServiceID(addressToString(localServices[i]));
        }

        return fullServiceIDList;
    }

    // get the registered counterparty service list
    function getLocalWhiteList(address addr) public view returns (bool) {
        return localWhiteList[addr];
    }

    function invokeInterchains(
        string[] memory srcFullID,
        string[] memory destAddr,
        uint64[] memory index,
        uint64[] memory typ,
        string[] memory callFunc,
        bytes[][] memory args,
        uint64[] memory txStatus,
        bytes[][] memory signatures,
        bool[] memory isEncrypt) payable external {
        for (uint8 i = 0; i < srcFullID.length; ++i) {
            if (serviceOrdered[stringToAddress(destAddr[i])] == true) {
                string memory dstFullID = genFullServiceID(destAddr[i]);
                invokeIndexUpdateWithError(srcFullID[i], dstFullID, index[i], txStatus[i], isEncrypt[i], "dst service is not ordered", uint64(1));
                continue;
            }
            invokeInterchain(srcFullID[i], destAddr[i], index[i], typ[i], callFunc[i], args[i], txStatus[i], signatures[i], isEncrypt[i]);
        }
    }

    // called on dest chain
    function invokeInterchain(
        string memory srcFullID,
    // 地址变为string格式，这样多签不会有问题，在验证多签之前使用checksum之前的合约地址
        string memory destAddr,
        uint64 index,
        uint64 typ,
        string memory callFunc,
        bytes[] memory args,
        uint64 txStatus,
        bytes[] memory signatures,
        bool isEncrypt) payable public {
        string memory dstFullID = genFullServiceID(destAddr);
        string memory servicePair = genServicePair(srcFullID, dstFullID);
        {
            bool ok = checkInterchainMultiSigns(srcFullID, dstFullID, index, typ, callFunc, args, txStatus, signatures);
            if (!ok) {
                invokeIndexUpdateWithError(srcFullID, dstFullID, index, txStatus, isEncrypt, "invalid interchain-multi-signature", uint64(1));
                return;
            }

            if (localWhiteList[stringToAddress(destAddr)] == false) {
                invokeIndexUpdateWithError(srcFullID, dstFullID, index, txStatus, isEncrypt, "dest address is not in local white list", uint64(1));
                return;
            }
        }

        //        bool status = true;
        bool[] memory status = new bool[](1);
        status[0] = true;
        bytes[][] memory results = new bytes[][](1);
        if (txStatus == 0) {
            // INTERCHAIN && BEGIN
            if (inCounter[servicePair] < index) {
                (status[0], results[0]) = callService(stringToAddress(destAddr), callFunc, args, false);
            }
            invokeIndexUpdate(srcFullID, dstFullID, index, 0);
            if (status[0]) {
                typ = 1;
            } else {
                typ = 2;
            }
        } else {
            // INTERCHAIN && FAILURE || INTERCHAIN && ROLLBACK, only happened in relay mode
            if (inCounter[servicePair] >= index) {
                (status[0], results[0]) = callService(stringToAddress(destAddr), callFunc, args, true);
            }
            invokeIndexUpdate(srcFullID, dstFullID, index, 2);
            if (txStatus == 1) {
                typ = 2;
            } else {
                typ = 3;
            }
        }


        receiptMessages[servicePair][index] = Receipt(isEncrypt, typ, results, status);

        if (isEncrypt) {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, new bytes[][](0), computeHash(results), status);
        } else {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, results, computeHash(results), status);
        }
    }

    // called on dest chain
    function invokeMultiInterchain(
        string memory srcFullID,
    // 地址变为string格式，这样多签不会有问题，在验证多签之前使用checksum之前的合约地址
        string memory destAddr,
        uint64 index,
        uint64 typ,
        string memory callFunc,
        bytes[][] memory args,
        uint64 txStatus,
        bytes[] memory signatures,
        bool isEncrypt) payable public {
        string memory dstFullID = genFullServiceID(destAddr);
        string memory servicePair = genServicePair(srcFullID, dstFullID);
        {
            bool ok = checkMultiInterchainMultiSigns(srcFullID, dstFullID, index, typ, callFunc, args, txStatus, signatures);
            if (!ok) {
                invokeIndexUpdateWithError(srcFullID, dstFullID, index, txStatus, isEncrypt, "invalid multiInterchain-multi-signature", uint64(args.length));
                return;
            }

            if (localWhiteList[stringToAddress(destAddr)] == false) {
                invokeIndexUpdateWithError(srcFullID, dstFullID, index, txStatus, isEncrypt, "dest address is not in local white list", uint64(args.length));
                return;
            }
        }

        bytes[][] memory results = new bytes[][](args.length);
        bool[] memory multiStatus = new bool[](args.length);
        typ = 1;
        if (txStatus == 0) {
            // INTERCHAIN && BEGIN
            if (inCounter[servicePair] < index) {
                (multiStatus, results) = callMultiService(stringToAddress(destAddr), callFunc, args, false);
                for (uint i = 0; i < multiStatus.length; i++){
                    if(!multiStatus[i]){
                        typ = 2;
                        break;
                    }
                }
            }
            invokeIndexUpdate(srcFullID, dstFullID, index, 0);
        } else {
            // INTERCHAIN && FAILURE || INTERCHAIN && ROLLBACK, only happened in relay mode
            if (inCounter[servicePair] >= index) {
                (multiStatus, results) = callMultiService(stringToAddress(destAddr), callFunc, args, true);
            }
            invokeIndexUpdate(srcFullID, dstFullID, index, 2);
            if (txStatus == 1) {
                typ = 2;
            } else {
                typ = 3;
            }
        }


        receiptMessages[servicePair][index] = Receipt(isEncrypt, typ, results, multiStatus);

        if (isEncrypt) {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, new bytes[][](0), computeHash(results), multiStatus);
        } else {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, results, computeHash(results), multiStatus);
        }
    }

    function callService(address destAddr, string memory callFunc, bytes[] memory args, bool isRollback) private returns (bool, bytes[] memory) {
        bool status = true;
        bytes[] memory result;

        if (keccak256(abi.encodePacked(callFunc)) != keccak256(abi.encodePacked(""))) {
            (bool ok, bytes memory data) = address(destAddr).call(abi.encodeWithSignature(string(abi.encodePacked(callFunc, "(bytes[],bool)")), args, isRollback));
            status = ok;
            if (status) {
                result = abi.decode(data, (bytes[]));
            }
        }

        return (status, result);
    }


    function callMultiService(address destAddr, string memory callFunc, bytes[][] memory args, bool isRollback) private returns (bool[] memory, bytes[][] memory) {
        bool status = true;
        bytes[][] memory Results;
        bool[] memory MultiStatus;

        if (keccak256(abi.encodePacked(callFunc)) != keccak256(abi.encodePacked(""))) {
            (bool ok, bytes memory data) = address(destAddr).call(abi.encodeWithSignature(string(abi.encodePacked(callFunc, "(bytes[][],bool)")), args, isRollback));
            status = ok;
            if (status) {
                (Results, MultiStatus) = abi.decode(data, (bytes[][],bool[]));
            }
        }

        return (MultiStatus, Results);
    }

    function computeHash(bytes[][] memory args) internal pure returns (bytes32) {
        bytes memory packed;
        for (uint i = 0; i < args.length; i++) {
            bytes[] memory arg = args[i];
            for (uint j = 0; j < arg.length; j++) {
                packed = abi.encodePacked(packed, arg[j]);
            }
        }
        return keccak256(packed);
    }

    // called on src chain
    function invokeReceipt(
        string memory srcAddr,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        bytes[][] memory results,
        uint64 txStatus,
        bytes[] memory signatures) payable external {
        string memory srcFullID = genFullServiceID(srcAddr);
        bool isRollback = false;
        if (txStatus != 0 && txStatus != 3) {
            isRollback = true;
        }

        invokeIndexUpdate(srcFullID, dstFullID, index, 1);
        checkReceiptMultiSigns(srcFullID, dstFullID, index, typ, results, txStatus, signatures);

        string memory outServicePair = genServicePair(srcFullID, dstFullID);
        CallFunc memory invokeFunc = outMessages[outServicePair][index].callback;
        bytes[] memory args = new bytes[](invokeFunc.args.length);

        if (isRollback) {
            invokeFunc = outMessages[outServicePair][index].rollback;
            args = new bytes[](invokeFunc.args.length);
            for (uint i = 0; i < invokeFunc.args.length; i++) {
                args[i] = invokeFunc.args[i];
            }
        }

        if (!isRollback) {
            args = new bytes[](invokeFunc.args.length + results[0].length);
            for (uint i = 0; i < invokeFunc.args.length; i++) {
                args[i] = invokeFunc.args[i];
            }
            for (uint i = 0; i < results[0].length; i++) {
                args[invokeFunc.args.length + i] = results[0][i];
            }
        }

        if (keccak256(abi.encodePacked(invokeFunc.func)) != keccak256(abi.encodePacked(""))) {
            string memory method = string(abi.encodePacked(invokeFunc.func, "(bytes[])"));
            (bool ok,) = address(stringToAddress(srcAddr)).call(abi.encodeWithSignature(method, args));
            emit throwReceiptStatus(ok);
            return;
        }

        emit throwReceiptStatus(true);
    }

    // called on src chain
    function invokeMultiReceipt(
        string memory srcAddr,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        bytes[][] memory results,
        bool[] memory multiStatus,
        uint64 txStatus,
        bytes[] memory signatures) payable external {
        string memory srcFullID = genFullServiceID(srcAddr);
        bool isRollback = false;

        if (txStatus != 0 && txStatus != 3) {
            isRollback = true;
        }
        {
            invokeIndexUpdate(srcFullID, dstFullID, index, 1);
            checkReceiptMultiSigns(srcFullID, dstFullID, index, typ, results, txStatus, signatures);
        }

        string memory outServicePair = genServicePair(srcFullID, dstFullID);


        if (isRollback) {
            CallFunc memory invokeFunc = outMessages[outServicePair][index].rollback;
            bytes[] memory args = new bytes[](invokeFunc.args.length);
            for (uint i = 0; i < invokeFunc.args.length; i++) {
                args[i] = invokeFunc.args[i];
            }
            if (keccak256(abi.encodePacked(invokeFunc.func)) != keccak256(abi.encodePacked(""))) {
                (bool ok,) = address(stringToAddress(srcAddr)).call(abi.encodeWithSignature(string(abi.encodePacked(invokeFunc.func, "(bytes[],bool[])")), args, multiStatus));
                if (!ok){
                    emit throwReceiptStatus(false);
                    return;
                }
            }
        }

        bool flag = false;
        for (uint i = 0; i < multiStatus.length; i++) {
            if (multiStatus[i] == true){
                flag = true;
                break;
            }
        }

        if (flag) {
            CallFunc memory invokeFunc = outMessages[outServicePair][index].callback;
            bytes[] memory args = new bytes[](invokeFunc.args.length);
            for (uint i = 0; i < invokeFunc.args.length; i++) {
                args[i] = invokeFunc.args[i];
            }
            if (keccak256(abi.encodePacked(invokeFunc.func)) != keccak256(abi.encodePacked(""))) {
                (bool ok,) = address(stringToAddress(srcAddr)).call(abi.encodeWithSignature(string(abi.encodePacked(invokeFunc.func, "(bytes[],bool[],bytes[][])")), args, multiStatus, results));
                if (!ok) {
                    emit throwReceiptStatus(false);
                    return;
                }
            }
        }
        emit throwReceiptStatus(true);
    }


    function invokeIndexUpdate(string memory srcFullID, string memory dstFullID, uint64 index, uint64 reqType) private {
        string memory servicePair = genServicePair(srcFullID, dstFullID);
        if (reqType == 0) {
            require(inCounter[servicePair] + 1 == index);
            markInCounter(servicePair);
        } else if (reqType == 1) {
            // invoke src callback or rollback
            require(callbackCounter[servicePair] + 1 == index);
            markCallbackCounter(servicePair, index);
        } else if (reqType == 2) {
            // invoke dst rollback
            // if one to multi, maybe invoke rollback for smaller index
            if (dstRollbackCounter[servicePair] + 1 > index) {
                Receipt memory receipt = receiptMessages[servicePair][index];
                require(receipt.typ == 1);
            }

            markDstRollbackCounter(servicePair, index);
            if (inCounter[servicePair] + 1 == index) {
                markInCounter(servicePair);
            }
        }
    }

    function invokeIndexUpdateWithError(string memory srcFullID, string memory dstFullID, uint64 index, uint64 txStatus, bool isEncrypt, string memory errorMsg, uint64 resultsSize) private {
        string memory servicePair = genServicePair(srcFullID, dstFullID);
        uint64 typ;
        bytes[][] memory results = new bytes[][](resultsSize);
        bytes[] memory result = new bytes[](1);
        for (uint64 i = 0; i < resultsSize; i++) {
            result[0] = bytes(errorMsg);
            results[i] = result;
        }

        if (txStatus == 0) {
            invokeIndexUpdate(srcFullID, dstFullID, index, 0);
            typ = 2;
        } else {
            invokeIndexUpdate(srcFullID, dstFullID, index, 2);
            if (txStatus == 1) {
                typ = 2;
            } else {
                typ = 3;
            }
        }

        bool[] memory multiStatus = new bool[](resultsSize);
        for (uint64 i = 0; i < resultsSize; i++) {
            multiStatus[i] = false;
        }

        receiptMessages[servicePair][index] = Receipt(isEncrypt, typ, results, multiStatus);

        if (isEncrypt) {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, new bytes[][](0), computeHash(results), multiStatus);
        } else {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, results, computeHash(results), multiStatus);
        }
    }

    function emitInterchainEvent(
        string memory destFullServiceID,
        string memory funcCall,
        bytes[] memory args,
        string memory funcCb,
        bytes[] memory argsCb,
        string memory funcRb,
        bytes[] memory argsRb,
        bool isEncrypt,
        string[] memory group)
    public onlyWhiteList {
        checkAppchainIdContains(appchainID, destFullServiceID);
        string memory curFullID = genFullServiceID(addressToString(msg.sender));
        string memory outServicePair = genServicePair(curFullID, destFullServiceID);

        // Record the order of interchain contract which has been started.
        outCounter[outServicePair]++;
        if (outCounter[outServicePair] == 1) {
            outServicePairs.push(outServicePair);
        }

        outMessages[outServicePair][outCounter[outServicePair]] = InterchainInvoke(isEncrypt, group,
            CallFunc(funcCall, args),
            CallFunc(funcCb, argsCb),
            CallFunc(funcRb, argsRb));

        bytes memory packed = abi.encodePacked(funcCall);
        for (uint i = 0; i < args.length; i++) {
            packed = abi.encodePacked(packed, args[i]);
        }
        bytes32 hash = keccak256(packed);

        if (isEncrypt) {
            funcCall = "";
            args = new bytes[](0);
        }

        // Throw interchain event for listening of plugin.
        emit throwInterchainEvent(outCounter[outServicePair], destFullServiceID, curFullID, funcCall, args, hash, group);
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
    }

    // The helper functions that help plugin query.
    function getOuterMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](outServicePairs.length);
        for (uint64 i = 0; i < outServicePairs.length; i++) {
            indices[i] = outCounter[outServicePairs[i]];
        }

        return (outServicePairs, indices);
    }

    function getOutMessage(string memory outServicePair, uint64 idx) public view returns (string memory, bytes[] memory, bool, string[] memory) {
        InterchainInvoke memory invoke = outMessages[outServicePair][idx];
        return (invoke.callFunc.func, invoke.callFunc.args, invoke.encrypt, invoke.group);
    }

    function getReceiptMessage(string memory inServicePair, uint64 idx) public view returns (bytes[][] memory, uint64, bool, bool[] memory)  {
        Receipt memory receipt = receiptMessages[inServicePair][idx];
        return (receipt.results, receipt.typ, receipt.encrypt, receipt.multiStatus);
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

    function getDstRollbackMeta() public view returns (string[] memory, uint64[] memory) {
        uint64[] memory indices = new uint64[](inServicePairs.length);
        for (uint i = 0; i < inServicePairs.length; i++) {
            indices[i] = dstRollbackCounter[inServicePairs[i]];
        }

        return (inServicePairs, indices);
    }

    function genFullServiceID(string memory serviceID) private view returns (string memory) {
        return string(abi.encodePacked(bitxhubID, ":", appchainID, ":", serviceID));
    }

    function genServicePair(string memory from, string memory to) private pure returns (string memory) {
        return string(abi.encodePacked(from, "-", to));
    }

    function getChainID() public view returns (string memory, string memory) {
        return (bitxhubID, appchainID);
    }

    function checkInterchainMultiSigns(string memory srcFullID,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        string memory callFunc,
        bytes[] memory args,
        uint64 txStatus,
        bytes[] memory multiSignatures) private returns (bool) {
        bytes memory packed = abi.encodePacked(srcFullID, dstFullID, index, typ);
        bytes memory funcPacked = abi.encodePacked(callFunc);

        funcPacked = abi.encodePacked(funcPacked, uint64(0));
        for (uint i = 0; i < args.length; i++) {
            funcPacked = abi.encodePacked(funcPacked, args[i]);
        }
        packed = abi.encodePacked(packed, keccak256(funcPacked), txStatus);
        bytes32 hash = keccak256(packed);

        return checkMultiSigns(hash, multiSignatures);
    }

    function checkMultiInterchainMultiSigns(string memory srcFullID,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        string memory callFunc,
        bytes[][] memory args,
        uint64 txStatus,
        bytes[] memory multiSignatures) private returns (bool) {
        bytes memory packed = abi.encodePacked(srcFullID, dstFullID, index, typ);
        bytes memory funcPacked = abi.encodePacked(callFunc);
        funcPacked = abi.encodePacked(funcPacked, uint64(1));
        if (args.length == 0) {
            funcPacked = abi.encodePacked(funcPacked, uint64(0));
        } else {
            funcPacked = abi.encodePacked(funcPacked, uint64(args[0].length));
        }
        for (uint i = 0; i < args.length; i++) {
            bytes[] memory arg = args[i];
            for (uint j = 0; j < arg.length; j++) {
                funcPacked = abi.encodePacked(funcPacked, arg[j]);
            }
        }
        packed = abi.encodePacked(packed, keccak256(funcPacked), txStatus);
        bytes32 hash = keccak256(packed);

        //        require(checkMultiSigns(hash, multiSignatures), "invalid MultiInterchain-multi-signature");
        return checkMultiSigns(hash, multiSignatures);
    }

    function checkReceiptMultiSigns(string memory srcFullID,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        bytes[][] memory results,
        uint64 txStatus,
        bytes[] memory multiSignatures) private {
        bytes memory packed = abi.encodePacked(srcFullID, dstFullID, index, typ);
        bytes memory data;
        if (typ == 0) {
            string memory outServicePair = genServicePair(srcFullID, dstFullID);
            CallFunc memory callFunc = outMessages[outServicePair][index].callFunc;
            data = abi.encodePacked(data, callFunc.func);
            for (uint i = 0; i < callFunc.args.length; i++) {
                data = abi.encodePacked(data, callFunc.args[i]);
            }
        } else {
            for (uint i = 0; i < results.length; i++) {
                bytes[] memory result = results[i];
                for (uint j = 0; j < result.length; i++) {
                    data = abi.encodePacked(data, result[j]);
                }
            }
        }
        packed = abi.encodePacked(packed, keccak256(data), txStatus);
        bytes32 hash = keccak256(packed);

        require(checkMultiSigns(hash, multiSignatures), "invalid Receipt-multi-signature");
    }

    function checkMultiSigns(bytes32 hash, bytes[] memory multiSignatures) private returns (bool) {
        for (uint i = 0; i < multiSignatures.length; i++) {
            bytes memory sig = multiSignatures[i];
            if (sig.length != 65) {
                continue;
            }

            (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);

            address addr = ecrecover(hash, v, r, s);

            if (addressArrayContains(validators, addr)) {
                if (addressArrayContains(bxhSigners, addr)) {
                    continue;
                }
                bxhSigners.push(addr);
                if (bxhSigners.length == valThreshold) {
                    delete bxhSigners;
                    return true;
                }
            }
        }
        delete bxhSigners;
        return false;
    }

    function addressArrayContains(address[] memory addrs, address addr) private pure returns (bool) {
        for (uint i = 0; i < addrs.length; i++) {
            if (addrs[i] == addr) {
                return true;
            }
        }

        return false;
    }

    function splitSignature(bytes memory sig) internal pure returns (uint8 v, bytes32 r, bytes32 s) {
        assembly {
        // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
        // second 32 bytes
            s := mload(add(sig, 64))
        // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        return (v + 27, r, s);
    }

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
    ) internal pure returns (string memory asciiString) {
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
            b = uint8(uint160(data) / (2 ** (8 * (19 - i))));
            leftNibble = b / 16;
            rightNibble = b - 16 * leftNibble;

            // locate and extract each capitalization status.
            leftCaps = caps[2 * i];
            rightCaps = caps[2 * i + 1];

            // get the offset from nibble value to ascii character for left nibble.
            asciiOffset = _getAsciiOffset(leftNibble, leftCaps);

            // add the converted character to the byte array.
            asciiBytes[2 * i] = bytes1(leftNibble + asciiOffset);

            // get the offset from nibble value to ascii character for right nibble.
            asciiOffset = _getAsciiOffset(rightNibble, rightCaps);

            // add the converted character to the byte array.
            asciiBytes[2 * i + 1] = bytes1(rightNibble + asciiOffset);
        }


        return string(abi.encodePacked("0x", asciiBytes));
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
            asciiBytes[2 * i] = bytes1(leftNibble + (leftNibble < 10 ? 48 : 87));
            asciiBytes[2 * i + 1] = bytes1(rightNibble + (rightNibble < 10 ? 48 : 87));
        }

        return string(asciiBytes);
    }

    function stringToAddress(string memory _address) internal pure returns (address) {
        bytes memory temp = bytes(_address);
        if (temp.length != 42) {
            revert(string(abi.encodePacked(_address, " is not a valid address")));
        }

        uint160 result = 0;
        uint160 b1;
        uint160 b2;
        for (uint256 i = 2; i < 2 + 2 * 20; i += 2) {
            result *= 256;
            b1 = uint160(uint8(temp[i]));
            b2 = uint160(uint8(temp[i + 1]));
            if ((b1 >= 97) && (b1 <= 102)) {
                b1 -= 87;
            } else if ((b1 >= 65) && (b1 <= 70)) {
                b1 -= 55;
            } else if ((b1 >= 48) && (b1 <= 57)) {
                b1 -= 48;
            }

            if ((b2 >= 97) && (b2 <= 102)) {
                b2 -= 87;
            } else if ((b2 >= 65) && (b2 <= 70)) {
                b2 -= 55;
            } else if ((b2 >= 48) && (b2 <= 57)) {
                b2 -= 48;
            }
            result += (b1 * 16 + b2);
        }
        return address(result);
    }

    function checkAppchainIdContains(string memory appchainId, string memory destFullService) private pure {
        bytes memory whatBytes = bytes(appchainId);
        bytes memory whereBytes = bytes(destFullService);

        require(whereBytes.length >= whatBytes.length);

        bool found = false;
        for (uint i = 0; i <= whereBytes.length - whatBytes.length; i++) {
            bool flag = true;
            for (uint j = 0; j < whatBytes.length; j++)
                if (whereBytes [i + j] != whatBytes [j]) {
                    flag = false;
                    break;
                }
            if (flag) {
                found = true;
                break;
            }
        }
        // 不允许同broker服务自跨链
        require(!found, "dest service is belong to current broker!");
    }
}