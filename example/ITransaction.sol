// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.6;
pragma experimental ABIEncoderV2;

interface ITransacion {
    function registerBroker() external;

    function initialize() external;

    function registerAppchain(string memory chainID, string memory broker, address ruleAddr, bytes memory trustRoot) external;

    function getAppchainInfo(string memory chainID) external view returns (string memory, bytes memory, address);

    function registerRemoteService(string memory chainID, string memory serviceID, address[] memory whiteList) external;

    function getRSWhiteList(string memory remoteAddr) external view returns (address[] memory);

    function getRemoteServiceList() external view returns (string[] memory);

    function startTransaction(string memory from, string memory to, uint64 index) external;

    function rollbackTransaction(string memory from, string memory to, uint64 index) external;

    function endTransactionSuccess(string memory from, string memory to, uint64 index) external;

    function endTransactionFail(string memory from, string memory to, uint64 index) external;

    function endTransactionRollback(string memory from, string memory to, uint64 index) external;

    function getTransactionStatus(string memory IBTPid) external view returns (uint64);

    function getStartTimestamp(string memory IBTPid) external view returns (uint);

    function register(address addr, bool ordered, uint admins, uint64 _adminThreshold) external;

    function audit(address addr, int64 status, address voter, uint admins) external returns (bool);

    function getLocalServiceList(string memory bitxhubID, string memory appchainID) external view returns (string[] memory);

    function getLocalWhiteList(address addr) external view returns (bool);
}