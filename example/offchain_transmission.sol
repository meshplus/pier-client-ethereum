pragma solidity >=0.5.6;
pragma experimental ABIEncoderV2;

// 1.支持结构化数据的push和get
// 2.支持链下数据的push和get
contract OffChainTransmission {
    // 只有本链存在的数据才写在当前
    struct Data {
        // 该数据的key
        string key;
        // 假如是结构化数据，就是明文；假如是非结构化，就是链下数据唯一id
        string value;
        // 标记链上还是链下 '1':链下，'0':链上
        string isStruct;
    }
    mapping(string => Data) dataM; // map for accounts
    // change the address of Broker accordingly
    address BrokerAddr;

    // AccessControl
    modifier onlyBroker {
        require(msg.sender == BrokerAddr, "Invoker are not the Broker");
        _;
    }

    constructor(address _brokerAddr) public {
        BrokerAddr = _brokerAddr;
        Broker(BrokerAddr).register();
    }

    function register() public {
        Broker(BrokerAddr).register();
    }

    // contract for data exchange
    function getData(string memory key) public view returns(string memory ,string memory, string memory) {
        return (dataM[key].key, dataM[key].value, dataM[key].isStruct);
    }

    function set(string memory key, string memory value, string memory isStruct) public {
        require(keccak256(abi.encodePacked(isStruct)) == keccak256("1") || keccak256(abi.encodePacked(isStruct)) == keccak256("0"), "struct type error!");
        dataM[key] = Data(key, value, isStruct);
    }


    // 扩展成为支持链下文件传输的合约
    function push(string memory destChainServiceID, string memory key) public {
        bytes memory _key = abi.encodePacked(dataM[key].key);
        require(keccak256(_key) == keccak256(abi.encodePacked(key)), "set data first!");
        bytes[] memory args = new bytes[](3);
        args[0] = abi.encodePacked(dataM[key].key);
        args[1] = abi.encodePacked(dataM[key].value);
        args[2] = abi.encodePacked(dataM[key].isStruct);


        bytes[] memory argsCb = new bytes[](1);
        argsCb[0] = abi.encodePacked(key);

        Broker(BrokerAddr).emitInterchainEvent(destChainServiceID, "interchainPush", args, "interchainConfirmed", argsCb, "", new bytes[](0), false);
    }

    function interchainPush(bytes[] memory args, bool isRollback) public onlyBroker returns(bytes[] memory){
        require(args.length == 3, "interchainPush args' length is not correct, expect 3");
        string memory key = string(args[0]);
        string memory value = string(args[1]);
        string memory isStruct = string(args[2]);
        set(key, value, isStruct);

        return new bytes[](0);
    }

    function interchainConfirmed(bytes[] memory args) public onlyBroker {
        require(args.length == 1, "interchainSet args' length is not correct, expect 1");
        // business logic todo
    }


    function get(string memory destChainServiceID, string memory key) public {
        bytes[] memory args = new bytes[](1);
        args[0] = abi.encodePacked(key);

        bytes[] memory argsCb = new bytes[](1);
        argsCb[0] = abi.encodePacked(key);

        Broker(BrokerAddr).emitInterchainEvent(destChainServiceID, "interchainGet", args, "interchainSet", argsCb, "", new bytes[](0), false);
    }


    function interchainSet(bytes[] memory args) public onlyBroker {
        require(args.length == 3, "interchainSet args' length is not correct, expect 3");
        string memory key = string(args[0]);
        string memory value = string(args[1]);
        string memory isStruct = string(args[2]);
        set(key, value, isStruct);
    }

    function interchainGet(bytes[] memory args, bool isRollback) public onlyBroker returns(bytes[] memory) {
        require(args.length == 1, "interchainGet args' length is not correct, expect 1");
        string memory key = string(args[0]);

        bytes[] memory result = new bytes[](2);
        result[0] = abi.encodePacked(dataM[key].value);
        result[1] = abi.encodePacked(dataM[key].isStruct);

        return result;
    }
}

abstract contract Broker {
    function emitInterchainEvent(
        string memory destFullServiceID,
        string memory func,
        bytes[] memory args,
        string memory funcCb,
        bytes[] memory argsCb,
        string memory funcRb,
        bytes[] memory argsRb,
        bool isEncrypt) public virtual;

    function register() public virtual;
}
