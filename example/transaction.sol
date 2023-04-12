// SPDX-License-Identifier: Unlicensed
pragma solidity >= 0.5.7;
pragma experimental ABIEncoderV2;
import "./ITransaction.sol";

contract Transaction is ITransacion {
    struct Proposal {
        uint64 approve;
        uint64 reject;
        address[] votedAdmins;
        bool ordered;
        bool exist;
    }

    struct Appchain {
        string id;
        string broker;
        bytes trustRoot;
        address ruleAddr;
        uint64 status;
        bool exist;
    }

    // Only the contract in the whitelist can invoke the Broker for interchain operations.
    mapping(address => bool) localWhiteList;
    address[] localServices;
    mapping(address => Proposal) localServiceProposal;
    address[] proposalList;
    mapping(address => bool) serviceOrdered;
    uint64 adminThreshold;

    mapping(string => Appchain) appchains;
    mapping(string => address[]) remoteWhiteList;
    string[] remoteServices;
    string[] appchainIDs;

    // begin-1, begin_rollback-2, success-3, fail-4, rollback-5
    mapping(string => uint64) transactionStatus;
    mapping(string => uint) startTimestamp;
    string[] transactionId;

    mapping(address => Proposal) localProposal;
    address public BrokerAddr;
    address[] public txAdmins;
    uint64 public txAdminThreshold;

    //AccessControl
    modifier onlyTxAdmin {
        bool flag = false;
        for (uint i = 0; i < txAdmins.length; i++) {
            if (msg.sender == txAdmins[i]) {flag = true;}
        }

        require(flag == true, "Invoker are not in admin list");
        _;
    }

    modifier onlyBroker{
        require(msg.sender == BrokerAddr, "Invoker are not the Broker");
        _;
    }

    constructor(address[] memory _txAdmins,
        uint64 _txAdminThreshold) {
        txAdmins = _txAdmins;
        txAdminThreshold = _txAdminThreshold;
    }

    function initialize() external override onlyBroker {
        for (uint n = 0; n < localServices.length; n++) {
            localWhiteList[localServices[n]] = false;
        }
        for (uint x = 0; x < proposalList.length; x++) {
            delete localServiceProposal[proposalList[x]];
        }
        delete localServices;

        for (uint x = 0; x < transactionId.length; ++x) {
            transactionStatus[transactionId[x]] = 0;
        }
        for (uint y = 0; y < remoteServices.length; y++) {
            delete remoteWhiteList[remoteServices[y]];
        }
        for (uint z = 0; z < appchainIDs.length; z++) {
            delete appchains[appchainIDs[z]];
        }
        delete transactionId;
        delete remoteServices;
        delete appchainIDs;
    }

    function registerBroker() external override {
        require(tx.origin != msg.sender, "register not by contract");
        if (BrokerAddr == msg.sender || localProposal[msg.sender].exist) {
            return;
        }

        localProposal[msg.sender] = Proposal(0, 0, new address[](txAdmins.length), false, true);
    }

    function auditBroker(address addr, int64 status) external onlyTxAdmin returns (bool) {
        uint result = voteBroker(addr, status);

        if (result == 0) {
            return false;
        }

        if (result == 1) {
            delete localProposal[addr];
            BrokerAddr = addr;
        } else {
            delete localProposal[addr];
        }

        return true;
    }

    function voteBroker(address addr, int64 status) private returns (uint) {
        require(localProposal[addr].exist, "the proposal does not exist");
        require(status == 0 || status == 1, "vote status should be 0 or 1");

        for (uint i = 0; i < localProposal[addr].votedAdmins.length; i++) {
            require(localProposal[addr].votedAdmins[i] != msg.sender, "current use has voted the proposal");
        }

        localProposal[addr].votedAdmins[localProposal[addr].reject + localProposal[addr].approve] = msg.sender;
        if (status == 0) {
            localProposal[addr].reject++;
            if (localProposal[addr].reject == txAdmins.length - txAdminThreshold + 1) {
                return 2;
            }
        } else {
            localProposal[addr].approve++;
            if (localProposal[addr].approve == txAdminThreshold) {
                return 1;
            }
        }

        return 0;
    }

    // register local service to Broker
    function register(address addr, bool ordered, uint admins, uint64 _adminThreshold) external onlyBroker override {
        if (localWhiteList[addr] || localServiceProposal[addr].exist) {
            return;
        }

        adminThreshold = _adminThreshold;
        localServiceProposal[addr] = Proposal(0, 0, new address[](admins), ordered, true);
    }

    function audit(address addr, int64 status, address voter, uint admins) external onlyBroker override returns (bool) {
        uint result = vote(addr, status, voter, admins);

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
    function vote(address addr, int64 status, address voter, uint admins) private returns (uint) {
        require(localServiceProposal[addr].exist, "the proposal does not exist");
        require(status == 0 || status == 1, "vote status should be 0 or 1");

        for (uint i = 0; i < localServiceProposal[addr].votedAdmins.length; i++) {
            require(localServiceProposal[addr].votedAdmins[i] != voter, "current use has voted the proposal");
        }

        localServiceProposal[addr].votedAdmins[localServiceProposal[addr].reject + localServiceProposal[addr].approve] = voter;
        if (status == 0) {
            localServiceProposal[addr].reject++;
            if (localServiceProposal[addr].reject == admins - adminThreshold + 1) {
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

    // register remote appchain ID in direct mode, invoked by appchain admin
    function registerAppchain(string memory chainID, string memory broker, address ruleAddr, bytes memory trustRoot) external onlyBroker override {
        require(appchains[chainID].exist == false, "this appchain has already been registered");
        // require(rule.length != 0, "validate rule should not be empty");

        appchains[chainID] = Appchain(chainID, broker, trustRoot, ruleAddr, 1, true);
        appchainIDs.push(chainID);
    }

    // register service ID from counterparty appchain in direct mode, invoked by appchain admin
    // serviceID: the service from counterparty appchain which will call service on current appchain
    // whiteList：service list on current appchain which are allowed to be called by remote service
    function registerRemoteService(string memory chainID, string memory serviceID, address[] memory whiteList) external onlyBroker override {
        require(appchains[chainID].exist == true, "this appchain is not registered");
        require(appchains[chainID].status == 1, "the appchain's status is not available");

        string memory fullServiceID = genRemoteFullServiceID(chainID, serviceID);
        // todo whiteList是否存在于当前audit中
        remoteWhiteList[fullServiceID] = whiteList;
        remoteServices.push(fullServiceID);
    }

    function getAppchainInfo(string memory chainID) external view override returns (string memory, bytes memory, address) {
        Appchain memory appchain = appchains[chainID];

        require(appchain.exist == true, "this appchain is not registered");

        return (appchain.broker, appchain.trustRoot, appchain.ruleAddr);
    }

    // get the registered counterparty service list
    function getRSWhiteList(string memory remoteAddr) external view override returns (address[] memory) {
        return remoteWhiteList[remoteAddr];
    }

    // get the registered counterparty service list
    function getRemoteServiceList() external view override returns (string[] memory) {
        return remoteServices;
    }

    function genRemoteFullServiceID(string memory chainID, string memory serviceID) private pure returns (string memory) {
        return string(abi.encodePacked(":", chainID, ":", serviceID));
    }

    // get the registered local service list
    function getLocalServiceList(string memory bitxhubID, string memory appchainID) external view override returns (string[] memory) {
        string[] memory fullServiceIDList = new string[](localServices.length);
        for (uint i = 0; i < localServices.length; i++) {
            fullServiceIDList[i] = genFullServiceID(bitxhubID, appchainID, addressToString(localServices[i]));
        }

        return fullServiceIDList;
    }

    // get the registered counterparty service list
    function getLocalWhiteList(address addr) external view override returns (bool) {
        return localWhiteList[addr];
    }

    function startTransaction(string memory from, string memory to, uint64 index) external onlyBroker override {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 0, "Transaction is recorded.");
        transactionStatus[IBTPid] = 1;
        startTimestamp[IBTPid] = block.timestamp;
        // record IBTPid
        transactionId.push(IBTPid);
    }

    function rollbackTransaction(string memory from, string memory to, uint64 index) external onlyBroker override {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 1, "Transaction status is not begin.");
        transactionStatus[IBTPid] = 2;
    }

    // begin => success
    function endTransactionSuccess(string memory from, string memory to, uint64 index) external onlyBroker override {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 1, "transaction status is not begin.");
        transactionStatus[IBTPid] = 3;
    }

    // begin => fail
    function endTransactionFail(string memory from, string memory to, uint64 index) external onlyBroker override {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 1, "Transaction status is not begin.");
        transactionStatus[IBTPid] = 4;
    }

    // begin_rollback => rollback
    function endTransactionRollback(string memory from, string memory to, uint64 index) external onlyBroker override {
        string memory IBTPid = genIBTPid(from, to, index);
        require(transactionStatus[IBTPid] == 2, "Transaction status is not begin_rollback.");
        transactionStatus[IBTPid] = 5;
    }

    function getTransactionStatus(string memory IBTPid) external view override returns (uint64) {
        return transactionStatus[IBTPid];
    }

    function getStartTimestamp(string memory IBTPid) external view override returns (uint) {
        return startTimestamp[IBTPid];
    }

    function genIBTPid(string memory from, string memory to, uint64 index) private pure returns (string memory) {
        string memory id = uint2str(uint(index));
        return string(abi.encodePacked(from, "-", to, "-", id));
    }

    function genFullServiceID(string memory bitxhubID, string memory appchainID, string memory serviceID) private pure returns (string memory) {
        return string(abi.encodePacked(bitxhubID, ":", appchainID, ":", serviceID));
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
