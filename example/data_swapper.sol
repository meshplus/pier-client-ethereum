pragma solidity >=0.5.6;
pragma experimental ABIEncoderV2;

contract DataSwapper {
    mapping(string => string) dataM; // map for accounts
    // change the address of Broker accordingly
    address BrokerAddr;

    // AccessControl
    modifier onlyBroker {
        require(msg.sender == BrokerAddr, "Invoker are not the Broker");
        _;
    }

    constructor(address _brokerAddr, bool _ordered) {
        BrokerAddr = _brokerAddr;
        Broker(BrokerAddr).register(_ordered);
    }

    function register(bool _ordered) public {
        Broker(BrokerAddr).register(_ordered);
    }

    // contract for data exchange
    function getData(string memory key) public view returns(string memory) {
        return dataM[key];
    }

    function get(string memory destChainServiceID, string[] memory keyArr) public {
        uint len = keyArr.length;
        bytes[] memory args = new bytes[](len);
        for (uint i = 0; i < len; i++) {
            args[i] = abi.encodePacked(keyArr[i]);
        }
        bytes[] memory argsCb = new bytes[](len);
        for (uint i = 0; i < len; i++) {
            argsCb[i] = abi.encodePacked(keyArr[i]);
        }
        Broker(BrokerAddr).emitInterchainEvent(destChainServiceID, "interchainGet", args, "interchainSet", argsCb, "", new bytes[](0), false, new string[](0));
    }

    function set(string memory key, string memory value) public {
        dataM[key] = value;
    }

    function interchainSet(bytes[] memory args,bool[] memory multiStatus,bytes[][] memory results) public onlyBroker {
        require(args.length == results.length, "interchainSet args' length must equals results' length");
        for (uint i = 0;i < args.length; i++){
            string memory key = string(args[i]);
            string memory value = string(results[i][0]);
            set(key, value);
        }
    }

    function interchainGet(bytes[] memory args, bool isRollback) public view onlyBroker returns(bytes[][] memory,bool[] memory) {
        uint len=args.length;
        bool[] memory multiStatus = new bool[](len);
        bytes[][] memory result=new bytes[][](len);
        for(uint i=0;i<len;i++){
            string memory key = string(args[i]);
            bytes[] memory tmp = new bytes[](1);
            tmp[0] = abi.encodePacked(dataM[key]);
            result[i] = tmp;
            multiStatus[i] = true;
        }
        return (result,multiStatus);
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
        bool isEncrypt,
        string[] memory group) public virtual;

    function register(bool ordered) public virtual;
}
