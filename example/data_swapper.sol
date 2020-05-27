pragma solidity >=0.5.7;

contract DataSwapper {
    mapping(string => string) dataM; // map for accounts
    Hasher hasher = Hasher(0x00000000000000000000000000000000000000fa);
    // change the address of Broker accordingly
//    address BrokerAddr = 0x9E0901D698E854F6CFE9e478C38d20A01908768a;
    address BrokerAddr = 0x2346f3BA3F0B6676aa711595daB8A27d0317DB57;
    Broker broker = Broker(BrokerAddr);

    // AccessControl
    modifier onlyBroker {
        require(msg.sender == BrokerAddr, "Invoker are not the Broker");
        _;
    }

    // 数据交换类的业务合约
    function getData(string memory key) public returns(string memory) {
        return dataM[key];
    }

    function get(address destChainID, string memory destAddr, string memory key) public {
        bool ok = broker.InterchainDataSwapInvoke(destChainID, destAddr, key);
        require(ok);
    }

    function set(string memory key, string memory value) public {
        dataM[key] = value;
    }

    function interchainSet(string memory key, string memory value) public onlyBroker {
        set(key, value);
    }

    function interchainGet(string memory key) public onlyBroker view returns(bool, string memory) {
        return (true, dataM[key]);
    }
}

contract Broker {
    function InterchainDataSwapInvoke(address destChainID, string memory destAddr, string memory key) public returns(bool);
}

contract Hasher {
    function getHash() public returns(bytes32);
}