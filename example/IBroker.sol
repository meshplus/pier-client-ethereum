// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.6;
pragma experimental ABIEncoderV2;

interface IBroker {
    function register() external;

    function initialize() external;

    function checkInterchainMultiSigns(string memory srcFullID,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        string memory callFunc,
        bytes[] memory args,
        uint64 txStatus,
        bytes[] memory multiSignatures,
        address[] memory validators,
        uint64 valThreshold) external returns(bool);

    function checkReceiptMultiSigns(string memory srcFullID,
        string memory dstFullID,
        uint64 index,
        uint64 typ,
        bytes[][] memory result,
        uint64 txStatus,
        bytes[] memory multiSignatures,
        address[] memory validators,
        uint64 valThreshold) external returns(bool);

    function setOutMessage(string memory servicePair,
        bool isEncrypt,
        string[] memory group,
        string memory funcCall,
        bytes[] memory args,
        string memory funcCb,
        bytes[] memory argsCb,
        string memory funcRb,
        bytes[] memory argsRb) external;

    function invokeIndexUpdate(string memory srcFullID, string memory dstFullID, uint64 index, uint64 reqType) external returns(bool);

    function invokeIndexUpdateForBatch(string memory srcFullID, string memory dstFullID, uint64 index, uint64 reqType) external returns(bool);

    function getInCounter(string memory servicePair) external view returns(uint64);

    function getCallbackMessage(string memory servicePair, uint64 index) external view returns(string memory, bytes[] memory);

    function getRollbackMessage(string memory servicePair, uint64 index) external view returns(string memory, bytes[] memory);

    function setReceiptMessage(string memory servicePair, uint64 index, bool isEncrypt, uint64 typ, bytes[][] memory results, bool[] memory multiStatus) external;

    function markOutCounter(string memory servicePair) external returns(uint64);

    function stringToAddress(string memory _address) external pure returns (address);

    function addressToString(address account) external pure returns (string memory asciiString);

    function getSplitLength(string memory _str, string memory _delimiter) external pure returns (uint8);

    function checkAppchainIdContains (string memory appchainId, string memory destFullService) external pure returns(bool);

    function getOuterMeta() external view returns (string[] memory, uint64[] memory);

    function getOutMessage(string memory outServicePair, uint64 idx) external view returns (string memory, bytes[] memory, bool, string[] memory);

    function getReceiptMessage(string memory inServicePair, uint64 idx) external view returns (bytes[][] memory, uint64, bool, bool[] memory);

    function getReceiptStatus(string memory inServicePair, uint64 idx) external view returns (bool);

    function getInnerMeta() external view returns (string[] memory, uint64[] memory);

    function getCallbackMeta() external view returns (string[] memory, uint64[] memory);

    function getDstRollbackMeta() external view returns (string[] memory, uint64[] memory);
}
