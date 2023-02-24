pragma solidity >=0.6.9 <=0.7.6;
pragma experimental ABIEncoderV2;
import "./IBroker.sol";

contract Broker {
    IBroker public brokerData;
    struct Proposal {
        uint64 approve;
        uint64 reject;
        address[] votedAdmins;
        bool ordered;
        bool exist;
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

    event throwInterchainEvent(uint64 index, string dstFullID, string srcFullID, string func, bytes[] args, bytes32 hash, string[] group);
    event throwReceiptEvent(uint64 index, string dstFullID, string srcFullID, uint64 typ, bytes[][] results, bytes32 hash, bool[] multiStatus);
    event throwReceiptStatus(bool);

    address dataAddr;

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
        uint64 _adminThreshold,
        address _dataAddr) {
        bitxhubID = _bitxhubID;
        appchainID = _appchainID;
        validators = _validators;
        valThreshold = _valThreshold;
        admins = _admins;
        adminThreshold = _adminThreshold;
        dataAddr = _dataAddr;
        brokerData = IBroker(_dataAddr);
        brokerData.register();
    }

    function setAdmins(address[] memory _admins, uint64 _adminThreshold) public onlyAdmin {
        admins = _admins;
        adminThreshold = _adminThreshold;
    }

    function setValidators(address[] memory _validators, uint64 _valThreshold) public onlyAdmin {
        validators = _validators;
        valThreshold = _valThreshold;
    }


    function initialize() public onlyAdmin {
        for (uint n = 0; n < localServices.length; n++) {
            localWhiteList[localServices[n]] = false;
        }
        for (uint x = 0; x < proposalList.length; x++) {
            delete localServiceProposal[proposalList[x]];
        }
        delete localServices;

        brokerData.initialize();
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
            fullServiceIDList[i] = genFullServiceID(brokerData.addressToString(localServices[i]));
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
        bool[] memory isEncrypt) payable external
    {
        for (uint8 i = 0; i <  srcFullID.length; ++i) {
            if (serviceOrdered[brokerData.stringToAddress(destAddr[i])] == true) {
                string memory dstFullID = genFullServiceID(destAddr[i]);
                invokeIndexUpdateWithError(srcFullID[i], dstFullID, index[i], txStatus[i], isEncrypt[i], "service invoke batch must not ordered", uint64(1));
                continue;
            }
            // batch flag is true
            bool[] memory enAndBatch = new bool[](2);
            // flag 0 is isEncrypt, flag 1 is isBatch
            enAndBatch[0] = isEncrypt[i];
            enAndBatch[1] = true;
            invokeInterchainWithBatchFlag(srcFullID[i], destAddr[i], index[i], typ[i], callFunc[i], args[i], txStatus[i], signatures[i], enAndBatch);
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
        bool isEncrypt) payable public
    {
        bool[] memory enAndBatch = new bool[](2);
        // flag 0 is isEncrypt, flag 1 is isBatch
        enAndBatch[0] = isEncrypt;
        // batch flag is false
        enAndBatch[1] = false;
        invokeInterchainWithBatchFlag(srcFullID, destAddr, index, typ, callFunc, args, txStatus, signatures, enAndBatch);
    }

    // called on dest chain
    function invokeInterchainWithBatchFlag(
        string memory srcFullID,
    // 地址变为string格式，这样多签不会有问题，在验证多签之前使用checksum之前的合约地址
        string memory destAddr,
        uint64 index,
        uint64 typ,
        string memory callFunc,
        bytes[] memory args,
        uint64 txStatus,
        bytes[] memory signatures,
        bool[] memory enAndBatch) internal
    {
        string memory dstFullID = genFullServiceID(destAddr);
        string memory servicePair = genServicePair(srcFullID, dstFullID);
        {
            bool ok = brokerData.checkInterchainMultiSigns(srcFullID, dstFullID, index, typ, callFunc, args, txStatus, signatures, validators, valThreshold);
            if (!ok) {
                //这个地方broker无法判断有多少笔交易，所以无法给出 resultSize
                invokeIndexUpdateWithError(srcFullID, dstFullID, index, txStatus, enAndBatch[0], "invalid Interchain-multi-signature", uint64(1));
                return;
            }

            if (localWhiteList[brokerData.stringToAddress(destAddr)] == false) {
                invokeIndexUpdateWithError(srcFullID, dstFullID, index, txStatus, enAndBatch[0], "dest address is not in local white list", uint64(1));
                return;
            }
        }

        bool[] memory multiStatus;
        bytes[][] memory results;
        typ = 1;

        // INTERCHAIN && BEGIN
        if (txStatus == 0) {
            // check index when it is not batch
            if (!enAndBatch[1]) {
                if (brokerData.getInCounter(servicePair) < index) {
                    (multiStatus, results) = callService(brokerData.stringToAddress(destAddr), callFunc, args, false);
                    require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 0));
                    // if call service failed, set the receipt typ to receipt_failure
                    for (uint i = 0; i < multiStatus.length; i++){
                        if(!multiStatus[i]){
                            typ = 2;
                            break;
                        }
                    }
                }
            } else {
                // if batch is true, ignore index ordered
                (multiStatus, results) = callService(brokerData.stringToAddress(destAddr), callFunc, args, false);
                require(brokerData.invokeIndexUpdateForBatch(srcFullID, dstFullID, index, 0));
            }

        } else {
            // INTERCHAIN && FAILURE || INTERCHAIN && ROLLBACK, only happened in relay mode
            if (brokerData.getInCounter(servicePair) >= index) {
                (multiStatus, results) = callService(brokerData.stringToAddress(destAddr), callFunc, args, true);
            }
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 2));
            if (txStatus == 1) {
                typ = 2;
            } else {
                typ = 3;
            }
        }
        brokerData.setReceiptMessage(servicePair, index, enAndBatch[0], typ, results, multiStatus);
        if (enAndBatch[0]) {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, new bytes[][](0), computeHash(results), multiStatus);
        } else {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, results, computeHash(results), multiStatus);
        }
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

    function callService(address destAddr, string memory callFunc, bytes[] memory args, bool isRollback) private returns (bool[] memory, bytes[][] memory) {
        bool[] memory multiStatus;
        bytes[][] memory results;

        if (keccak256(abi.encodePacked(callFunc)) != keccak256(abi.encodePacked(""))) {
            (bool ok, bytes memory data) = address(destAddr).call(abi.encodeWithSignature(string(abi.encodePacked(callFunc, "(bytes[],bool)")), args, isRollback));
            if (ok) {
                (results, multiStatus) = abi.decode(data, (bytes[][],bool[]));
            }
        }

        return (multiStatus, results);
    }


    function invokeReceipts(
        string[] memory srcAddrs,
        string[] memory dstFullIDs,
        uint64[] memory indexs,
        uint64[] memory typs,
        bytes[][][] memory batchResults,
        bool[][] memory batchMultiStatus,
        uint64[] memory batchTxStatus,
        bytes[][] memory batchSignatures) payable external {

        require(srcAddrs.length == dstFullIDs.length && srcAddrs.length == indexs.length && srcAddrs.length == typs.length
        && srcAddrs.length == batchResults.length && srcAddrs.length == batchMultiStatus.length
        && srcAddrs.length == batchTxStatus.length && srcAddrs.length == batchSignatures.length, "invalid input length");

        for (uint8 i = 0; i < srcAddrs.length; ++i) {
            //0: interchain; 1: receipt_success; 2: receipt_fail; 3: receipt_Rollback
            // batch flag is true
            invokeReceiptWithBatchFlag(srcAddrs[i], dstFullIDs[i], indexs[i], typs[i], batchResults[i],
                batchMultiStatus[i], batchTxStatus[i], batchSignatures[i], true);
        }
    }

    // called on src chain
    function invokeReceipt(
        string memory srcAddr,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        bytes[][] memory results,
        bool[] memory multiStatus,
        uint64 txStatus,
        bytes[] memory signatures) payable external
    {
        // batch flag is false
        invokeReceiptWithBatchFlag(srcAddr, dstFullID, index, typ, results, multiStatus, txStatus, signatures, false);
    }

    // called on src chain
    function invokeReceiptWithBatchFlag(
        string memory srcAddr,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        bytes[][] memory results,
        bool[] memory multiStatus,
        uint64 txStatus,
        bytes[] memory signatures,
        bool batch) internal
    {
        string memory srcFullID = genFullServiceID(srcAddr);
        bool isRollback = false;

        if (txStatus != 0 && txStatus != 3) {
            isRollback = true;
        }
        {
            // 1. update index
            if (!batch) {
                require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 1));
            } else {
                require(brokerData.invokeIndexUpdateForBatch(srcFullID, dstFullID, index, 1));
            }
            // 2. check multiSign
            require(brokerData.checkReceiptMultiSigns(srcFullID, dstFullID, index, typ, results, txStatus, signatures, validators, valThreshold), "check receipt multiSign failed");
        }

        string memory outServicePair = genServicePair(srcFullID, dstFullID);
        // invoke other contract
        receiptCall(outServicePair, index, isRollback, srcAddr, results, multiStatus);
    }

    function receiptCall(string memory servicePair, uint64 index, bool isRollback, string memory srcAddr, bytes[][] memory results, bool[] memory multiStatus) private {
        string memory callFunc;
        bytes[] memory callArgs;
        bytes[] memory args;
        if (isRollback) {
            (callFunc, callArgs) = brokerData.getRollbackMessage(servicePair, index);
            args = new bytes[](callArgs.length);
            for (uint i = 0; i < callArgs.length; i++) {
                args[i] = callArgs[i];
            }
            if (keccak256(abi.encodePacked(callFunc)) != keccak256(abi.encodePacked(""))) {
                (bool ok,) = address(brokerData.stringToAddress(srcAddr)).call(abi.encodeWithSignature(string(abi.encodePacked(callFunc, "(bytes[],bool[])")), args, multiStatus));
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
            (callFunc, callArgs) = brokerData.getCallbackMessage(servicePair, index);
            args = new bytes[](callArgs.length);
            for (uint i = 0; i < callArgs.length; i++) {
                args[i] = callArgs[i];
            }
            if (keccak256(abi.encodePacked(callFunc)) != keccak256(abi.encodePacked(""))) {
                (bool ok,) = address(brokerData.stringToAddress(srcAddr)).call(abi.encodeWithSignature(string(abi.encodePacked(callFunc, "(bytes[],bool[],bytes[][])")), args, multiStatus, results));
                if (!ok) {
                    emit throwReceiptStatus(false);
                    return;
                }
            }
        }
        emit throwReceiptStatus(true);
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

        if(txStatus == 0) {
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 0));
            typ = 2;
        } else {
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 2));
            if(txStatus == 1) {
                typ = 2;
            } else {
                typ = 3;
            }
        }

        bool[] memory multiStatus = new bool[](resultsSize);
        for (uint64 i = 0; i < resultsSize; i++) {
            multiStatus[i] = false;
        }

        brokerData.setReceiptMessage(servicePair, index, isEncrypt, typ, results, multiStatus);

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
        // 不允许同broker服务自跨链
        require(!brokerData.checkAppchainIdContains(appchainID, destFullServiceID), "dest service is belong to current broker!");
        // 不允许输入不规范的fullServiceID
        require(brokerData.getSplitLength(destFullServiceID, ":") == 3, "dest service id is not correct");

        string memory curFullID = genFullServiceID(brokerData.addressToString(msg.sender));
        string memory outServicePair = genServicePair(curFullID, destFullServiceID);

        // Record the order of interchain contract which has been started.
        uint64 currentOutCounter = brokerData.markOutCounter(outServicePair);


        brokerData.setOutMessage(outServicePair, isEncrypt, group, funcCall, args, funcCb, argsCb, funcRb, argsRb);

        bytes32 hash = computeInvokeHash(funcCall, args);

        if (isEncrypt) {
            funcCall = "";
            args = new bytes[](0);
        }

        // Throw interchain event for listening of plugin.
        emit throwInterchainEvent(currentOutCounter, destFullServiceID, curFullID, funcCall, args, hash, group);
    }

    function computeInvokeHash(string memory funcCall, bytes[] memory args) private pure returns(bytes32) {
        bytes memory packed = abi.encodePacked(funcCall);
        for (uint i = 0; i < args.length; i++) {
            packed = abi.encodePacked(packed, args[i]);
        }
        return keccak256(packed);
    }

    // The helper functions that help plugin query.
    function getOuterMeta() public view returns (string[] memory, uint64[] memory) {
        return brokerData.getOuterMeta();
    }

    function getOutMessage(string memory outServicePair, uint64 idx) public view returns (string memory, bytes[] memory, bool, string[] memory) {
        return brokerData.getOutMessage(outServicePair, idx);
    }

    function getReceiptMessage(string memory inServicePair, uint64 idx) public view returns (bytes[][] memory, uint64, bool, bool[] memory)  {
        return brokerData.getReceiptMessage(inServicePair, idx);
    }

    function getInnerMeta() public view returns (string[] memory, uint64[] memory) {
        return brokerData.getInnerMeta();
    }

    function getCallbackMeta() public view returns (string[] memory, uint64[] memory) {
        return brokerData.getCallbackMeta();
    }

    function getDstRollbackMeta() public view returns (string[] memory, uint64[] memory) {
        return brokerData.getDstRollbackMeta();
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
}