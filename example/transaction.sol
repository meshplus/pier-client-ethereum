pragma solidity >= 0.5.7;
pragma experimental ABIEncoderV2;

contract Transaction {
    // begin-1, begin_rollback-2, success-3, fail-4, rollback-5
    mapping(string => uint64) transactionStatus;

    string[] transactionId;

    address brokerAddr;

    struct Appchain {
        string id;
        string broker;
        bytes trustRoot;
        address ruleAddr;
        uint64 status;
        bool exist;
    }

    mapping(string => Appchain) appchains;
    mapping(string => address[]) remoteWhiteList;
    string[] remoteServices;
    string[] appchainIDs;

    //AccessControl
    modifier onlyBroker{
        require(msg.sender == brokerAddr, "Invoker are not the Broker");
        _;
    }

    constructor(address _brokerAddr) public {
        brokerAddr = _brokerAddr;
    }

    // register remote appchain ID in direct mode, invoked by appchain admin
    function registerAppchain(string memory chainID, string memory broker, address ruleAddr, bytes memory trustRoot) public onlyBroker {
        require(appchains[chainID].exist == false, "this appchain has already been registered");
        // require(rule.length != 0, "validate rule should not be empty");

        appchains[chainID] = Appchain(chainID, broker, trustRoot, ruleAddr, 1, true);
        appchainIDs.push(chainID);
    }

    // register service ID from counterparty appchain in direct mode, invoked by appchain admin
    // serviceID: the service from counterparty appchain which will call service on current appchain
    // whiteList：service list on current appchain which are allowed to be called by remote service
    function registerRemoteService(string memory chainID, string memory serviceID, address[] memory whiteList) public onlyBroker {
        require(appchains[chainID].exist == true, "this appchain is not registered");
        require(appchains[chainID].status == 1, "the appchain's status is not available");

        string memory fullServiceID = genRemoteFullServiceID(chainID, serviceID);
        // todo whiteList是否存在于当前audit中
        remoteWhiteList[fullServiceID] = whiteList;
        remoteServices.push(fullServiceID);
    }

    function getAppchainInfo(string memory chainID) public view returns (string memory, bytes memory, address) {
        Appchain memory appchain = appchains[chainID];

        require(appchain.exist == true, "this appchain is not registered");

        return (appchain.broker, appchain.trustRoot, appchain.ruleAddr);
    }

    // get the registered counterparty service list
    function getRSWhiteList(string memory remoteAddr) public view returns (address[] memory) {
        return remoteWhiteList[remoteAddr];
    }

    // get the registered counterparty service list
    function getRemoteServiceList() public view returns (string[] memory) {
        return remoteServices;
    }

    function genRemoteFullServiceID(string memory chainID, string memory serviceID) public view returns (string memory) {
        return string(abi.encodePacked(":", chainID, ":", serviceID));
    }

    function initialize() public onlyBroker {
        for (uint i = 0; i < transactionId.length; ++i) {
            transactionStatus[transactionId[i]] = 0;
        }
        for (uint x = 0; x < remoteServices.length; x++) {
            delete remoteWhiteList[remoteServices[x]];
        }
        for (uint z = 0; z < appchainIDs.length; z++) {
            delete appchains[appchainIDs[z]];
        }
        delete transactionId;
        delete remoteServices;
        delete appchainIDs;
    }

    function startTransaction(string memory from, string memory to, uint64 index) public onlyBroker {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 0, "Transaction is recorded.");
        transactionStatus[IBTPid] = 1;
        // record IBTPid
        transactionId.push(IBTPid);
    }

    function rollbackTransaction(string memory from, string memory to, uint64 index) public onlyBroker {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 1, "Transaction status is not begin.");
        transactionStatus[IBTPid] = 2;
    }

    // begin => success
    function endTransactionSuccess(string memory from, string memory to, uint64 index) public onlyBroker {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 1, "transaction status is not begin.");
        transactionStatus[IBTPid] = 3;
    }

    // begin => fail
    function endTransactionFail(string memory from, string memory to, uint64 index) public onlyBroker {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 1, "Transaction status is not begin.");
        transactionStatus[IBTPid] = 4;
    }

    // begin_rollback => rollback
    function endTransactionRollback(string memory from, string memory to, uint64 index) public onlyBroker {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 2, "Transaction status is not begin_rollback.");
        transactionStatus[IBTPid] = 5;
    }

    function getTransactionStatus(string memory IBTPid) public view returns (uint64) {
        return transactionStatus[IBTPid];
    }

    function genIBTPid(string memory from, string memory to, uint64 index) public pure returns (string memory) {
        string memory id = uint2str(uint(index));
        return string(abi.encodePacked(from, "-", to, "-", id));
    }

    function uint2str(uint _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len;
        while (_i != 0) {
            k = k-1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
