pragma solidity >=0.5.6;

contract DataSwapper {
    mapping(string => string) dataM; // map for accounts
    // change the address of Broker accordingly
    address BrokerAddr = 0x97135d4d2578dd2347FF5382db77553bE50bff3f;
    Broker broker = Broker(BrokerAddr);

    // AccessControl
    modifier onlyBroker {
        require(msg.sender == BrokerAddr, "Invoker are not the Broker");
        _;
    }

    // contract for data exchange
    function getData(string memory key) public returns(string memory) {
        return dataM[key];
    }

    function get(address destChainID, string memory destAddr, string memory key) public {
        broker.emitInterchainEvent(destChainID, destAddr, "interchainGet,interchainSet,", key, key, "");
    }

    function set(string memory key, string memory value) public {
        dataM[key] = value;
    }

    function interchainSet(string memory key, string memory value) public onlyBroker {
        set(key, value);
    }

    function interchainGet(string memory key) public onlyBroker returns(bool, string memory) {
        return (true, dataM[key]);
    }
}

abstract contract Broker {
    function emitInterchainEvent(
        address destChainID,
        string memory destAddr,
        string memory funcs,
        string memory args,
        string memory argscb,
        string memory argsrb) virtual public;
}