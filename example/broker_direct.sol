pragma solidity >=0.6.9 <=0.7.6;
pragma experimental ABIEncoderV2;
import "./IBroker.sol";
import "./ITransaction.sol";

contract BrokerDirect {
    IBroker brokerData;
    // transaction management contract in direct mode
    ITransacion transaction;

    string bitxhubID;
    string appchainID;
    address[] public admins;
    uint64 public adminThreshold;

    event throwInterchainEvent(uint64 index, string dstFullID, string srcFullID, string func, bytes[] args, bytes32 hash, string[] group);
    event throwReceiptEvent(uint64 index, string dstFullID, string srcFullID, uint64 typ, bytes[][] results, bytes32 hash, bool[] multiStatus);
    event throwReceiptStatus(bool);

    // Authority control. Contracts need to be registered.
    modifier onlyWhiteList {
        require(transaction.getLocalWhiteList(msg.sender) == true, "Invoker are not in white list");
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

    constructor(string memory _appchainID,
        address[] memory _admins,
        uint64 _adminThreshold,
        address _dataAddr,
        address _transactionAddr) {
        appchainID = _appchainID;
        admins = _admins;
        adminThreshold = _adminThreshold;
        brokerData = IBroker(_dataAddr);
        brokerData.register();
        transaction = ITransacion(_transactionAddr);
        transaction.registerBroker();
    }

    // update admin list and adminThreshold
    function setAdmins(address[] memory _admins, uint64 _adminThreshold) public onlyAdmin {
        admins = _admins;
        adminThreshold = _adminThreshold;
    }

    function initialize() public onlyAdmin {
        transaction.initialize();
        brokerData.initialize();
    }

    // register local service to Broker
    function register(bool ordered) public {
        require(tx.origin != msg.sender, "register not by contract");
        transaction.register(msg.sender, ordered, admins.length, adminThreshold);
    }

    function audit(address addr, int64 status) public onlyAdmin returns (bool) {
        return transaction.audit(addr, status, msg.sender, admins.length);
    }

    // register remote appchain ID in direct mode, invoked by appchain admin
    function registerAppchain(string memory chainID, string memory broker, address ruleAddr, bytes memory trustRoot) public onlyAdmin {
        transaction.registerAppchain(chainID, broker, ruleAddr, trustRoot);
    }

    // register service ID from counterparty appchain in direct mode, invoked by appchain admin
    // serviceID: the service from counterparty appchain which will call service on current appchain
    // whiteList：service list on current appchain which are allowed to be called by remote service
    function registerRemoteService(string memory chainID, string memory serviceID, address[] memory whiteList) public onlyAdmin {
        transaction.registerRemoteService(chainID, serviceID, whiteList);
    }

    function getAppchainInfo(string memory chainID) public view returns (string memory, bytes memory, address) {
        return transaction.getAppchainInfo(chainID);
    }

    // get the registered local service list
    function getLocalServiceList() public view returns (string[] memory) {
        return transaction.getLocalServiceList(bitxhubID, appchainID);
    }

    // get the registered counterparty service list
    function getRemoteServiceList() public view returns (string[] memory) {
        return transaction.getRemoteServiceList();
    }

    // get the registered counterparty service list
    function getRSWhiteList(string memory remoteAddr) public view returns (address[] memory) {
        return transaction.getRSWhiteList(remoteAddr);
    }

    // get the registered counterparty service list
    function getLocalWhiteList(address addr) public view returns (bool) {
        return transaction.getLocalWhiteList(addr);
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
        bool isEncrypt) payable external {
        // bool isRollback = false;
        string memory dstFullID = genFullServiceID(destAddr);
        string memory servicePair = genServicePair(srcFullID, dstFullID);

        bool[] memory multiStatus;
        bytes[][] memory results;

        if (!checkService(srcFullID, brokerData.stringToAddress(destAddr), index, dstFullID, typ, isEncrypt, uint64(1))) {
            return;
        }

        if (txStatus == 0) {
            typ = 1;
            // INTERCHAIN && BEGIN
            if (brokerData.getInCounter(servicePair) < index) {
                (multiStatus, results) = callService(brokerData.stringToAddress(destAddr), callFunc, args, false);
            }
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 0));
            // if call service failed, set the receipt typ to receipt_failure
            for (uint i = 0; i < multiStatus.length; i++){
                if(!multiStatus[i]){
                    typ = 2;
                    break;
                }
            }
        } else {
            // INTERCHAIN && FAILURE || INTERCHAIN && ROLLBACK
            // rollback only if RECEIPT_SUCCESS
            if (brokerData.getInCounter(servicePair) >= index && brokerData.getReceiptStatus(servicePair, index)) {
                (multiStatus, results) = callService(brokerData.stringToAddress(destAddr), callFunc, args, true);
            }
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 2));
            // ROLLBACK -> ROLLBACK_END
            typ = 4;
        }

        brokerData.setReceiptMessage(servicePair, index, isEncrypt, typ, results, multiStatus);

        if (isEncrypt) {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, new bytes[][](0), computeHash(results), multiStatus);
        } else {
            emit throwReceiptEvent(index, dstFullID, srcFullID, typ, results, computeHash(results), multiStatus);
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
        bytes[] memory signatures) payable external {
        string memory srcFullID = genFullServiceID(srcAddr);
        bool isRollback = false;
        // IBTP_RECEIPT_SUCCESS || IBTP_RECEIPT_FAILURE || IBTP_RECEIPT_ROLLBACK || IBTP_RECEIPT_ROLLBACK_END
        require(typ == 1 || typ == 2 || typ == 3 || typ == 4, "IBTP type is not correct in direct mode");
        if (typ == 1) {
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 1));
            transaction.endTransactionSuccess(srcFullID, dstFullID, index);
        }
        if (typ == 2) {
            isRollback = true;
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 1));
            transaction.endTransactionFail(srcFullID, dstFullID, index);
        }
        if (typ == 3) {
            isRollback = true;
            transaction.rollbackTransaction(srcFullID, dstFullID, index);
        }
        if (typ == 4) {
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 1));
            transaction.endTransactionRollback(srcFullID, dstFullID, index);
            return;
        }

        string memory outServicePair = genServicePair(srcFullID, dstFullID);

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
        require(!brokerData.checkAppchainIdContains(appchainID, destFullServiceID), "dest service is belong to current broker!");
        string memory curFullID = genFullServiceID(brokerData.addressToString(msg.sender));
        string memory outServicePair = genServicePair(curFullID, destFullServiceID);

        {
            // 直连模式下未注册的remoteService无法发出跨链交易
            bool flag = false;
            string[] memory remoteServices = transaction.getRemoteServiceList();
            for (uint i = 0; i < remoteServices.length; i++) {
                if (keccak256(abi.encodePacked(destFullServiceID)) == keccak256(abi.encodePacked(remoteServices[i]))) {
                    flag = true;
                    break;
                }
            }
            require(flag == true, "remote service is not registered");
            flag = false;
            address[] memory banList = transaction.getRSWhiteList(destFullServiceID);
            for (uint i = 0; i < banList.length; i++) {
                if (msg.sender == banList[i]) {
                    flag = true;
                    break;
                }
            }
            require(flag == false, "remote service is not allowed to call dest address");
        }


        // Record the order of interchain contract which has been started.
        uint64 currentOutCounter = brokerData.markOutCounter(outServicePair);

        brokerData.setOutMessage(outServicePair, isEncrypt, group, funcCall, args, funcCb, argsCb, funcRb, argsRb);

        bytes32 hash = computeInvokeHash(funcCall, args);

        if (isEncrypt) {
            funcCall = "";
            args = new bytes[](0);
        }

        // Start transaction and record current block number in direct mode
        transaction.startTransaction(curFullID, destFullServiceID, currentOutCounter);

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

    // get transaction start timestamp and transaction status in direct mode
    function getDirectTransactionMeta(string memory id) public view returns (uint, uint64) {
        return (transaction.getStartTimestamp(id), transaction.getTransactionStatus(id));
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

    // checkService: 1. check if dst service is registered;
    // 2. check if src service is registered;
    // 3. check if src service is in banList;
    function checkService(string memory remoteService, address destAddr, uint64 index, string memory dstFullID, uint64 typ, bool isEncrypt, uint64 resultsSize) private returns(bool) {
        //require(localWhiteList[destAddr] == true, "dest address is not in local white list");
        if (!transaction.getLocalWhiteList(destAddr)) {
            invokeIndexUpdateWithError(remoteService, dstFullID, index, typ, isEncrypt, "dest address is not in local white list", resultsSize);
            return false;
        }

        bool flag = false;
        string[] memory remoteServices = transaction.getRemoteServiceList();
        for (uint i = 0; i < remoteServices.length; i++) {
            if (keccak256(abi.encodePacked(remoteService)) == keccak256(abi.encodePacked(remoteServices[i]))) {
                flag = true;
                break;
            }
        }
        //require(flag == true, "remote service is not registered");
        if (!flag) {
            invokeIndexUpdateWithError(remoteService, dstFullID, index, typ, isEncrypt, "remote service is not registered", resultsSize);
            return false;
        }

        flag = false;
        address[] memory banList = transaction.getRSWhiteList(remoteService);
        for (uint i = 0; i < banList.length; i++) {
            if (destAddr == banList[i]) {
                flag = true;
                break;
            }
        }
        if (flag) {
            invokeIndexUpdateWithError(remoteService, dstFullID, index, typ, isEncrypt, "remote service is not allowed to call dest address", resultsSize);
            return false;
        }
        return true;
    }

    function invokeIndexUpdateWithError(string memory srcFullID, string memory dstFullID, uint64 index, uint64 typ, bool isEncrypt, string memory errorMsg, uint64 resultsSize) private {
        string memory servicePair = genServicePair(srcFullID, dstFullID);
        bytes[][] memory results = new bytes[][](resultsSize);
        bytes[] memory result = new bytes[](1);
        result[0] = bytes(errorMsg);
        for (uint64 i = 0; i < resultsSize; i++) {
            results[i] = result;
        }

        // INTERCHAIN => RECEIPT_FAIL
        // RECEIPT_ROLLBACK => RECEIPT_ROLLBACK_END
        if (typ == 0) {
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 0));
            typ = 2;
        } else {
            require(brokerData.invokeIndexUpdate(srcFullID, dstFullID, index, 2));
            typ = 4;
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
}