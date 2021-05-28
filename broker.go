// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BrokerABI is the input ABI used to generate the binding from.
const BrokerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogInterchainData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LogInterchainStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destDID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fid\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"funcs\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"argscb\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"argsrb\",\"type\":\"string\"}],\"name\":\"throwEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"status\",\"type\":\"int64\"}],\"name\":\"audit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"destContractDID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funcs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argscb\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argsrb\",\"type\":\"string\"}],\"name\":\"emitInterchainEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallbackMeta\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getInMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInnerMeta\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getOutMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOuterMeta\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcChainMethod\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"req\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"err\",\"type\":\"string\"}],\"name\":\"invokeIndexUpdateWithError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcChainMethod\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"req\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"bizCallData\",\"type\":\"bytes\"}],\"name\":\"invokeInterchain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BrokerFuncSigs maps the 4-byte function signature to its string representation.
var BrokerFuncSigs = map[string]string{
	"b38ff85f": "audit(address,int64)",
	"bed34bba": "compareStrings(string,string)",
	"b508c8d4": "emitInterchainEvent(string,string,string,string,string)",
	"3b6bbe4a": "getCallbackMeta()",
	"31942306": "getInMessage(string,uint64)",
	"67b9fa3b": "getInnerMeta()",
	"29793e6e": "getOutMessage(string,uint64)",
	"c20cab50": "getOuterMeta()",
	"8129fc1c": "initialize()",
	"82e2384e": "invokeIndexUpdateWithError(string,uint64,bool,string)",
	"7018c8c8": "invokeInterchain(string,uint64,address,bool,bytes)",
	"4420e486": "register(address)",
}

// BrokerBin is the compiled bytecode used for deploying new contracts.
var BrokerBin = "0x608060405234801561001057600080fd5b50611f8c806100206000396000f3fe6080604052600436106100a75760003560e01c80638129fc1c116100645780638129fc1c1461016f57806382e2384e14610184578063b38ff85f146101a4578063b508c8d4146101d1578063bed34bba146101f1578063c20cab5014610211576100a7565b806329793e6e146100ac57806331942306146100e25780633b6bbe4a146101025780634420e4861461012557806367b9fa3b146101475780637018c8c81461015c575b600080fd5b3480156100b857600080fd5b506100cc6100c7366004611b17565b610226565b6040516100d99190611e5c565b60405180910390f35b3480156100ee57600080fd5b506100cc6100fd366004611b17565b61026f565b34801561010e57600080fd5b5061011761028a565b6040516100d9929190611cb0565b34801561013157600080fd5b506101456101403660046118dd565b610447565b005b34801561015357600080fd5b5061011761046f565b61014561016a366004611942565b610609565b34801561017b57600080fd5b506101456106c5565b34801561019057600080fd5b5061014561019f366004611b5b565b6108a5565b3480156101b057600080fd5b506101c46101bf366004611904565b6108b7565b6040516100d99190611d54565b3480156101dd57600080fd5b506101456101ec366004611a4c565b610987565b3480156101fd57600080fd5b506101c461020c3660046119ec565b610b57565b34801561021d57600080fd5b50610117610bb0565b600061023183610d5e565b6007836040516102419190611c21565b90815260408051602092819003830190206001600160401b0385166000908152925290205490505b92915050565b600061027a83610d5e565b6009836040516102419190611c21565b60608060606005805490506001600160401b03811180156102aa57600080fd5b506040519080825280602002602001820160405280156102d4578160200160208202803683370190505b50905060005b6005546001600160401b038216101561036857600a6005826001600160401b03168154811061030557fe5b9060005260206000200160405161031c9190611c3d565b9081526040519081900360200190205482516001600160401b0391821691849190841690811061034857fe5b6001600160401b03909216602092830291909101909101526001016102da565b5060058181805480602002602001604051908101604052809291908181526020016000905b828210156104385760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156104245780601f106103f957610100808354040283529160200191610424565b820191906000526020600020905b81548152906001019060200180831161040757829003601f168201915b50505050508152602001906001019061038d565b50505050915092509250509091565b6001600160a01b03166000908152602081905260409020805467ffffffffffffffff19169055565b60608060606004805490506001600160401b038111801561048f57600080fd5b506040519080825280602002602001820160405280156104b9578160200160208202803683370190505b50905060005b600454811015610539576008600482815481106104d857fe5b906000526020600020016040516104ef9190611c3d565b9081526040519081900360200190205482516001600160401b039091169083908390811061051957fe5b6001600160401b03909216602092830291909101909101526001016104bf565b5060048181805480602002602001604051908101604052809291908181526020016000905b828210156104385760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156105f55780601f106105ca576101008083540402835291602001916105f5565b820191906000526020600020905b8154815290600101906020018083116105d857829003601f168201915b50505050508152602001906001019061055e565b6001600160a01b038416600090815260208190526040902054600790810b900b60011461063557600080fd5b61068587878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602081019091529081528a93508892509050610e4f565b604051602060848237805160208160040183378151808260240184376000808285348b5af19150503d806000843e8180156106c1578184a08184f35b8184fd5b60005b6004548110156107325760006008600483815481106106e357fe5b906000526020600020016040516106fa9190611c3d565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff199092169190911790556001016106c8565b5060005b6003548110156107a057600060066003838154811061075157fe5b906000526020600020016040516107689190611c3d565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff19909216919091179055600101610736565b5060005b60055481101561080e576000600a600583815481106107bf57fe5b906000526020600020016040516107d69190611c3d565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff199092169190911790556001016107a4565b5060005b60015481101561087e5760008060006001848154811061082e57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020805460079290920b6001600160401b031667ffffffffffffffff19909216919091179055600101610812565b5061088b600360006116eb565b610897600460006116eb565b6108a3600560006116eb565b565b6108b184848484610e4f565b50505050565b60008160070b600019141580156108d257508160070b600014155b80156108e257508160070b600114155b156108ef57506000610269565b6001600160a01b0383166000908152602081905260409020805467ffffffffffffffff19166001600160401b03600785900b908116919091179091556001141561097e576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0385161790555b50600192915050565b33600090815260208190526040902054600790810b900b6001146109c65760405162461bcd60e51b81526004016109bd90611d5f565b60405180910390fd5b60606109d18661102c565b90506006816040516109e39190611c21565b90815260405190819003602001812080546001600160401b038082166001011667ffffffffffffffff19909116179055600690610a21908390611c21565b908152604051908190036020019020546001600160401b031660011415610a8757600380546001810182556000919091528151610a85917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0190602084019061170c565b505b43600782604051610a989190611c21565b90815260200160405180910390206000600684604051610ab89190611c21565b9081526040805160209281900383019020546001600160401b0316835290820192909252810160002091909155517f85026d9efdca77567284833e9e5fa5ab0d412e8d0d9869942def88510971977c90600690610b16908490611c21565b90815260405190819003602001812054610b47916001600160401b0390911690899033908a908a908a908a90611e65565b60405180910390a1505050505050565b600081604051602001610b6a9190611c21565b6040516020818303038152906040528051906020012083604051602001610b919190611c21565b6040516020818303038152906040528051906020012014905092915050565b60608060606003805490506001600160401b0381118015610bd057600080fd5b50604051908082528060200260200182016040528015610bfa578160200160208202803683370190505b50905060005b6003546001600160401b0382161015610c8e5760066003826001600160401b031681548110610c2b57fe5b90600052602060002001604051610c429190611c3d565b9081526040519081900360200190205482516001600160401b03918216918491908416908110610c6e57fe5b6001600160401b0390921660209283029190910190910152600101610c00565b5060038181805480602002602001604051908101604052809291908181526020016000905b828210156104385760008481526020908190208301805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610d4a5780601f10610d1f57610100808354040283529160200191610d4a565b820191906000526020600020905b815481529060010190602001808311610d2d57829003601f168201915b505050505081526020019060010190610cb3565b6060610d8382604051806040016040528060018152602001601d60f91b8152506111ca565b90508051600414610da65760405162461bcd60e51b81526004016109bd90611dcd565b610ddf81600081518110610db657fe5b602002602001015160405180604001604052806003815260200162191a5960ea1b815250610b57565b610dfb5760405162461bcd60e51b81526004016109bd90611e04565b60015b6004811015610e4a576060828281518110610e1557fe5b60200260200101519050805160001415610e415760405162461bcd60e51b81526004016109bd90611e2f565b50600101610dfe565b505050565b8115610f4057826001600160401b0316600885604051610e6f9190611c21565b908152604051908190036020019020546001600160401b039081166001011614610e9857600080fd5b610ea18461135c565b604051602001610eb090611cad565b6040516020818303038152906040528051906020012081604051602001610ed79190611c21565b6040516020818303038152906040528051906020012014610f3b5780600b85604051610f039190611c21565b90815260408051602092819003830190206001600160401b0387166000908152908352208251610f39939192919091019061170c565b505b6108b1565b826001600160401b0316600a85604051610f5a9190611c21565b908152604051908190036020019020546001600160401b039081166001011614610f8357600080fd5b610f8d848461146d565b604051602001610f9c90611cad565b6040516020818303038152906040528051906020012081604051602001610fc39190611c21565b60405160208183030381529060405280519060200120146108b15780600c85604051610fef9190611c21565b90815260408051602092819003830190206001600160401b0387166000908152908352208251611025939192919091019061170c565b5050505050565b60608061105283604051806040016040528060018152602001601d60f91b8152506111ca565b905080516004146110755760405162461bcd60e51b81526004016109bd90611d96565b61108581600081518110610db657fe5b6110a15760405162461bcd60e51b81526004016109bd90611d96565b60015b60048110156110f05760608282815181106110bb57fe5b602002602001015190508051600014156110e75760405162461bcd60e51b81526004016109bd90611e2f565b506001016110a4565b50606061113a6111138360008151811061110657fe5b6020026020010151611580565b611135604051806040016040528060018152602001601d60f91b815250611580565b6115a5565b905061115861114882611580565b6111358460018151811061110657fe5b905061116661111382611580565b905061118461117482611580565b6111358460028151811061110657fe5b905061119261111382611580565b90506111c26111a082611580565b611135604051806040016040528060018152602001601760f91b815250611580565b949350505050565b606082600060015b600183510382101561120a5760006111eb87878561162b565b90508060001914156111fd575061120a565b60019081019250016111d2565b806001600160401b038111801561122057600080fd5b5060405190808252806020026020018201604052801561125457816020015b606081526020019060019003908161123f5790505b50935060009150600090505b6001835103821161135357600061127887878561162b565b9050806000191415611288575082515b60608382036001600160401b03811180156112a257600080fd5b506040519080825280601f01601f1916602001820160405280156112cd576020820181803683370190505b509050806000855b84811015611325578781815181106112e957fe5b602001015160f81c60f81b83838060010194508151811061130657fe5b60200101906001600160f81b031916908160001a9053506001016112d5565b508360010195508188868060010197508151811061133f57fe5b602002602001018190525050505050611260565b50505092915050565b60088160405161136c9190611c21565b90815260405190819003602001812080546001600160401b038082166001011667ffffffffffffffff199091161790556008906113aa908390611c21565b908152604051908190036020019020546001600160401b0316600114156114105760048054600181018255600091909152815161140e917f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0190602084019061170c565b505b436009826040516114219190611c21565b908152602001604051809103902060006008846040516114419190611c21565b9081526040805160209281900383019020546001600160401b0316835290820192909252016000205550565b600a8260405161147d9190611c21565b908152604051908190036020019020546001600160401b03166114df576005805460018101825560009190915282516114dd917f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00190602085019061170c565b505b80600a836040516114f09190611c21565b908152602001604051809103902060006101000a8154816001600160401b0302191690836001600160401b03160217905550436009836040516115339190611c21565b90815260200160405180910390206000600a856040516115539190611c21565b9081526040805160209281900383019020546001600160401b031683529082019290925201600020555050565b61158861178a565b506040805180820190915281518152602082810190820152919050565b805182516060918291016001600160401b03811180156115c457600080fd5b506040519080825280601f01601f1916602001820160405280156115ef576020820181803683370190505b509050600060208201905061160d81866020015187600001516116ad565b84516020850151855161162392840191906116ad565b509392505050565b81516000908490849060011461163d57fe5b835b825181101561169d578160008151811061165557fe5b602001015160f81c60f81b6001600160f81b03191683828151811061167657fe5b01602001516001600160f81b03191614156116955792506116a6915050565b60010161163f565b50600019925050505b9392505050565b5b602081106116cd578151835260209283019290910190601f19016116ae565b905182516020929092036101000a6000190180199091169116179052565b508054600082559060005260206000209081019061170991906117a4565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061174d57805160ff191683800117855561177a565b8280016001018555821561177a579182015b8281111561177a57825182559160200191906001019061175f565b506117869291506117c1565b5090565b604051806040016040528060008152602001600081525090565b808211156117865760006117b882826117d6565b506001016117a4565b5b8082111561178657600081556001016117c2565b50805460018160011615610100020316600290046000825580601f106117fc5750611709565b601f01602090049060005260206000209081019061170991906117c1565b60008083601f84011261182b578182fd5b5081356001600160401b03811115611841578182fd5b60208301915083602082850101111561185957600080fd5b9250929050565b600082601f830112611870578081fd5b81356001600160401b0380821115611886578283fd5b604051601f8301601f1916810160200182811182821017156118a6578485fd5b6040528281529250828483016020018610156118c157600080fd5b8260208601602083013760006020848301015250505092915050565b6000602082840312156118ee578081fd5b81356001600160a01b03811681146116a6578182fd5b60008060408385031215611916578081fd5b823561192181611f1e565b91506020830135600781900b8114611937578182fd5b809150509250929050565b600080600080600080600060a0888a03121561195c578283fd5b87356001600160401b0380821115611972578485fd5b61197e8b838c0161181a565b909950975060208a0135915061199382611f41565b9095506040890135906119a582611f1e565b9094506060890135906119b782611f33565b909350608089013590808211156119cc578384fd5b506119d98a828b0161181a565b989b979a50959850939692959293505050565b600080604083850312156119fe578182fd5b82356001600160401b0380821115611a14578384fd5b611a2086838701611860565b93506020850135915080821115611a35578283fd5b50611a4285828601611860565b9150509250929050565b600080600080600060a08688031215611a63578081fd5b85356001600160401b0380821115611a79578283fd5b611a8589838a01611860565b96506020880135915080821115611a9a578283fd5b611aa689838a01611860565b95506040880135915080821115611abb578283fd5b611ac789838a01611860565b94506060880135915080821115611adc578283fd5b611ae889838a01611860565b93506080880135915080821115611afd578283fd5b50611b0a88828901611860565b9150509295509295909350565b60008060408385031215611b29578182fd5b82356001600160401b03811115611b3e578283fd5b611b4a85828601611860565b925050602083013561193781611f41565b60008060008060808587031215611b70578384fd5b84356001600160401b0380821115611b86578586fd5b611b9288838901611860565b955060208701359150611ba482611f41565b909350604086013590611bb682611f33565b90925060608601359080821115611bcb578283fd5b50611bd887828801611860565b91505092959194509250565b6001600160401b0316815260200190565b60008151808452611c0d816020860160208601611ef2565b601f01601f19169290920160200192915050565b60008251611c33818460208701611ef2565b9190910192915050565b6000808354600180821660008114611c5c5760018114611c7357611ca2565b60ff198316865260028304607f1686019350611ca2565b600283048786526020808720875b83811015611c9a5781548a820152908501908201611c81565b505050860193505b509195945050505050565b90565b60006040820160408352808551611cc78184611e5c565b915081925060208082028301818901865b84811015611d02578683038652611cf0838351611bf5565b95840195925090830190600101611cd8565b50508681038288015280945087519350611d1c8482611e5c565b9450508087019150845b83811015611d4757611d39858451611be4565b945091810191600101611d26565b5092979650505050505050565b901515815260200190565b6020808252601d908201527f496e766f6b657220617265206e6f7420696e207768697465206c697374000000604082015260600190565b60208082526017908201527f646964206973206e6f74206c6567616c20666f726d6174000000000000000000604082015260600190565b60208082526017908201527f646964206973206e6f7420696e20666f75722070617274000000000000000000604082015260600190565b6020808252601190820152701c1c99599a5e081a5cc81b9bdd08191a59607a1b604082015260600190565b6020808252601390820152726469642073756273657420697320656d70747960681b604082015260600190565b90815260200190565b60006001600160401b038916825260e06020830152611e8760e0830189611bf5565b6001600160a01b03881660408401528281036060840152611ea88188611bf5565b90508281036080840152611ebc8187611bf5565b905082810360a0840152611ed08186611bf5565b905082810360c0840152611ee48185611bf5565b9a9950505050505050505050565b60005b83811015611f0d578181015183820152602001611ef5565b838111156108b15750506000910152565b6001600160a01b038116811461170957600080fd5b801515811461170957600080fd5b6001600160401b038116811461170957600080fdfea26469706673582212203599fc85c8098309286da82a412120d935d85200eb29c418b6aa0ae41be0244164736f6c634300060c0033"

// DeployBroker deploys a new Ethereum contract, binding an instance of Broker to it.
func DeployBroker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Broker, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BrokerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// Broker is an auto generated Go binding around an Ethereum contract.
type Broker struct {
	BrokerCaller     // Read-only binding to the contract
	BrokerTransactor // Write-only binding to the contract
	BrokerFilterer   // Log filterer for contract events
}

// BrokerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokerSession struct {
	Contract     *Broker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokerCallerSession struct {
	Contract *BrokerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrokerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokerTransactorSession struct {
	Contract     *BrokerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokerRaw struct {
	Contract *Broker // Generic contract binding to access the raw methods on
}

// BrokerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokerCallerRaw struct {
	Contract *BrokerCaller // Generic read-only contract binding to access the raw methods on
}

// BrokerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokerTransactorRaw struct {
	Contract *BrokerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBroker creates a new instance of Broker, bound to a specific deployed contract.
func NewBroker(address common.Address, backend bind.ContractBackend) (*Broker, error) {
	contract, err := bindBroker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// NewBrokerCaller creates a new read-only instance of Broker, bound to a specific deployed contract.
func NewBrokerCaller(address common.Address, caller bind.ContractCaller) (*BrokerCaller, error) {
	contract, err := bindBroker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerCaller{contract: contract}, nil
}

// NewBrokerTransactor creates a new write-only instance of Broker, bound to a specific deployed contract.
func NewBrokerTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokerTransactor, error) {
	contract, err := bindBroker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerTransactor{contract: contract}, nil
}

// NewBrokerFilterer creates a new log filterer instance of Broker, bound to a specific deployed contract.
func NewBrokerFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokerFilterer, error) {
	contract, err := bindBroker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokerFilterer{contract: contract}, nil
}

// bindBroker binds a generic wrapper to an already deployed contract.
func bindBroker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.BrokerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transact(opts, method, params...)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Broker *BrokerCaller) CompareStrings(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "compareStrings", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Broker *BrokerSession) CompareStrings(a string, b string) (bool, error) {
	return _Broker.Contract.CompareStrings(&_Broker.CallOpts, a, b)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Broker *BrokerCallerSession) CompareStrings(a string, b string) (bool, error) {
	return _Broker.Contract.CompareStrings(&_Broker.CallOpts, a, b)
}

// GetCallbackMeta is a free data retrieval call binding the contract method 0x3b6bbe4a.
//
// Solidity: function getCallbackMeta() view returns(string[], uint64[])
func (_Broker *BrokerCaller) GetCallbackMeta(opts *bind.CallOpts) ([]string, []uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCallbackMeta")

	if err != nil {
		return *new([]string), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetCallbackMeta is a free data retrieval call binding the contract method 0x3b6bbe4a.
//
// Solidity: function getCallbackMeta() view returns(string[], uint64[])
func (_Broker *BrokerSession) GetCallbackMeta() ([]string, []uint64, error) {
	return _Broker.Contract.GetCallbackMeta(&_Broker.CallOpts)
}

// GetCallbackMeta is a free data retrieval call binding the contract method 0x3b6bbe4a.
//
// Solidity: function getCallbackMeta() view returns(string[], uint64[])
func (_Broker *BrokerCallerSession) GetCallbackMeta() ([]string, []uint64, error) {
	return _Broker.Contract.GetCallbackMeta(&_Broker.CallOpts)
}

// GetInMessage is a free data retrieval call binding the contract method 0x31942306.
//
// Solidity: function getInMessage(string from, uint64 idx) view returns(uint256)
func (_Broker *BrokerCaller) GetInMessage(opts *bind.CallOpts, from string, idx uint64) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getInMessage", from, idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInMessage is a free data retrieval call binding the contract method 0x31942306.
//
// Solidity: function getInMessage(string from, uint64 idx) view returns(uint256)
func (_Broker *BrokerSession) GetInMessage(from string, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetInMessage(&_Broker.CallOpts, from, idx)
}

// GetInMessage is a free data retrieval call binding the contract method 0x31942306.
//
// Solidity: function getInMessage(string from, uint64 idx) view returns(uint256)
func (_Broker *BrokerCallerSession) GetInMessage(from string, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetInMessage(&_Broker.CallOpts, from, idx)
}

// GetInnerMeta is a free data retrieval call binding the contract method 0x67b9fa3b.
//
// Solidity: function getInnerMeta() view returns(string[], uint64[])
func (_Broker *BrokerCaller) GetInnerMeta(opts *bind.CallOpts) ([]string, []uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getInnerMeta")

	if err != nil {
		return *new([]string), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetInnerMeta is a free data retrieval call binding the contract method 0x67b9fa3b.
//
// Solidity: function getInnerMeta() view returns(string[], uint64[])
func (_Broker *BrokerSession) GetInnerMeta() ([]string, []uint64, error) {
	return _Broker.Contract.GetInnerMeta(&_Broker.CallOpts)
}

// GetInnerMeta is a free data retrieval call binding the contract method 0x67b9fa3b.
//
// Solidity: function getInnerMeta() view returns(string[], uint64[])
func (_Broker *BrokerCallerSession) GetInnerMeta() ([]string, []uint64, error) {
	return _Broker.Contract.GetInnerMeta(&_Broker.CallOpts)
}

// GetOutMessage is a free data retrieval call binding the contract method 0x29793e6e.
//
// Solidity: function getOutMessage(string to, uint64 idx) view returns(uint256)
func (_Broker *BrokerCaller) GetOutMessage(opts *bind.CallOpts, to string, idx uint64) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getOutMessage", to, idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutMessage is a free data retrieval call binding the contract method 0x29793e6e.
//
// Solidity: function getOutMessage(string to, uint64 idx) view returns(uint256)
func (_Broker *BrokerSession) GetOutMessage(to string, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetOutMessage(&_Broker.CallOpts, to, idx)
}

// GetOutMessage is a free data retrieval call binding the contract method 0x29793e6e.
//
// Solidity: function getOutMessage(string to, uint64 idx) view returns(uint256)
func (_Broker *BrokerCallerSession) GetOutMessage(to string, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetOutMessage(&_Broker.CallOpts, to, idx)
}

// GetOuterMeta is a free data retrieval call binding the contract method 0xc20cab50.
//
// Solidity: function getOuterMeta() view returns(string[], uint64[])
func (_Broker *BrokerCaller) GetOuterMeta(opts *bind.CallOpts) ([]string, []uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getOuterMeta")

	if err != nil {
		return *new([]string), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetOuterMeta is a free data retrieval call binding the contract method 0xc20cab50.
//
// Solidity: function getOuterMeta() view returns(string[], uint64[])
func (_Broker *BrokerSession) GetOuterMeta() ([]string, []uint64, error) {
	return _Broker.Contract.GetOuterMeta(&_Broker.CallOpts)
}

// GetOuterMeta is a free data retrieval call binding the contract method 0xc20cab50.
//
// Solidity: function getOuterMeta() view returns(string[], uint64[])
func (_Broker *BrokerCallerSession) GetOuterMeta() ([]string, []uint64, error) {
	return _Broker.Contract.GetOuterMeta(&_Broker.CallOpts)
}

// Audit is a paid mutator transaction binding the contract method 0xb38ff85f.
//
// Solidity: function audit(address addr, int64 status) returns(bool)
func (_Broker *BrokerTransactor) Audit(opts *bind.TransactOpts, addr common.Address, status int64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "audit", addr, status)
}

// Audit is a paid mutator transaction binding the contract method 0xb38ff85f.
//
// Solidity: function audit(address addr, int64 status) returns(bool)
func (_Broker *BrokerSession) Audit(addr common.Address, status int64) (*types.Transaction, error) {
	return _Broker.Contract.Audit(&_Broker.TransactOpts, addr, status)
}

// Audit is a paid mutator transaction binding the contract method 0xb38ff85f.
//
// Solidity: function audit(address addr, int64 status) returns(bool)
func (_Broker *BrokerTransactorSession) Audit(addr common.Address, status int64) (*types.Transaction, error) {
	return _Broker.Contract.Audit(&_Broker.TransactOpts, addr, status)
}

// EmitInterchainEvent is a paid mutator transaction binding the contract method 0xb508c8d4.
//
// Solidity: function emitInterchainEvent(string destContractDID, string funcs, string args, string argscb, string argsrb) returns()
func (_Broker *BrokerTransactor) EmitInterchainEvent(opts *bind.TransactOpts, destContractDID string, funcs string, args string, argscb string, argsrb string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "emitInterchainEvent", destContractDID, funcs, args, argscb, argsrb)
}

// EmitInterchainEvent is a paid mutator transaction binding the contract method 0xb508c8d4.
//
// Solidity: function emitInterchainEvent(string destContractDID, string funcs, string args, string argscb, string argsrb) returns()
func (_Broker *BrokerSession) EmitInterchainEvent(destContractDID string, funcs string, args string, argscb string, argsrb string) (*types.Transaction, error) {
	return _Broker.Contract.EmitInterchainEvent(&_Broker.TransactOpts, destContractDID, funcs, args, argscb, argsrb)
}

// EmitInterchainEvent is a paid mutator transaction binding the contract method 0xb508c8d4.
//
// Solidity: function emitInterchainEvent(string destContractDID, string funcs, string args, string argscb, string argsrb) returns()
func (_Broker *BrokerTransactorSession) EmitInterchainEvent(destContractDID string, funcs string, args string, argscb string, argsrb string) (*types.Transaction, error) {
	return _Broker.Contract.EmitInterchainEvent(&_Broker.TransactOpts, destContractDID, funcs, args, argscb, argsrb)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Broker *BrokerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Broker *BrokerSession) Initialize() (*types.Transaction, error) {
	return _Broker.Contract.Initialize(&_Broker.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Broker *BrokerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Broker.Contract.Initialize(&_Broker.TransactOpts)
}

// InvokeIndexUpdateWithError is a paid mutator transaction binding the contract method 0x82e2384e.
//
// Solidity: function invokeIndexUpdateWithError(string srcChainMethod, uint64 index, bool req, string err) returns()
func (_Broker *BrokerTransactor) InvokeIndexUpdateWithError(opts *bind.TransactOpts, srcChainMethod string, index uint64, req bool, err string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "invokeIndexUpdateWithError", srcChainMethod, index, req, err)
}

// InvokeIndexUpdateWithError is a paid mutator transaction binding the contract method 0x82e2384e.
//
// Solidity: function invokeIndexUpdateWithError(string srcChainMethod, uint64 index, bool req, string err) returns()
func (_Broker *BrokerSession) InvokeIndexUpdateWithError(srcChainMethod string, index uint64, req bool, err string) (*types.Transaction, error) {
	return _Broker.Contract.InvokeIndexUpdateWithError(&_Broker.TransactOpts, srcChainMethod, index, req, err)
}

// InvokeIndexUpdateWithError is a paid mutator transaction binding the contract method 0x82e2384e.
//
// Solidity: function invokeIndexUpdateWithError(string srcChainMethod, uint64 index, bool req, string err) returns()
func (_Broker *BrokerTransactorSession) InvokeIndexUpdateWithError(srcChainMethod string, index uint64, req bool, err string) (*types.Transaction, error) {
	return _Broker.Contract.InvokeIndexUpdateWithError(&_Broker.TransactOpts, srcChainMethod, index, req, err)
}

// InvokeInterchain is a paid mutator transaction binding the contract method 0x7018c8c8.
//
// Solidity: function invokeInterchain(string srcChainMethod, uint64 index, address destAddr, bool req, bytes bizCallData) payable returns()
func (_Broker *BrokerTransactor) InvokeInterchain(opts *bind.TransactOpts, srcChainMethod string, index uint64, destAddr common.Address, req bool, bizCallData []byte) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "invokeInterchain", srcChainMethod, index, destAddr, req, bizCallData)
}

// InvokeInterchain is a paid mutator transaction binding the contract method 0x7018c8c8.
//
// Solidity: function invokeInterchain(string srcChainMethod, uint64 index, address destAddr, bool req, bytes bizCallData) payable returns()
func (_Broker *BrokerSession) InvokeInterchain(srcChainMethod string, index uint64, destAddr common.Address, req bool, bizCallData []byte) (*types.Transaction, error) {
	return _Broker.Contract.InvokeInterchain(&_Broker.TransactOpts, srcChainMethod, index, destAddr, req, bizCallData)
}

// InvokeInterchain is a paid mutator transaction binding the contract method 0x7018c8c8.
//
// Solidity: function invokeInterchain(string srcChainMethod, uint64 index, address destAddr, bool req, bytes bizCallData) payable returns()
func (_Broker *BrokerTransactorSession) InvokeInterchain(srcChainMethod string, index uint64, destAddr common.Address, req bool, bizCallData []byte) (*types.Transaction, error) {
	return _Broker.Contract.InvokeInterchain(&_Broker.TransactOpts, srcChainMethod, index, destAddr, req, bizCallData)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Broker *BrokerTransactor) Register(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "register", addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Broker *BrokerSession) Register(addr common.Address) (*types.Transaction, error) {
	return _Broker.Contract.Register(&_Broker.TransactOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Broker *BrokerTransactorSession) Register(addr common.Address) (*types.Transaction, error) {
	return _Broker.Contract.Register(&_Broker.TransactOpts, addr)
}

// BrokerLogInterchainDataIterator is returned from FilterLogInterchainData and is used to iterate over the raw logs and unpacked data for LogInterchainData events raised by the Broker contract.
type BrokerLogInterchainDataIterator struct {
	Event *BrokerLogInterchainData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerLogInterchainDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerLogInterchainData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerLogInterchainData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerLogInterchainDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerLogInterchainDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerLogInterchainData represents a LogInterchainData event raised by the Broker contract.
type BrokerLogInterchainData struct {
	Status bool
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogInterchainData is a free log retrieval operation binding the contract event 0x436160f7c24c5f31561ec9422a629accdbbd4e9e8ce21e86e634f497997769a8.
//
// Solidity: event LogInterchainData(bool status, string data)
func (_Broker *BrokerFilterer) FilterLogInterchainData(opts *bind.FilterOpts) (*BrokerLogInterchainDataIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "LogInterchainData")
	if err != nil {
		return nil, err
	}
	return &BrokerLogInterchainDataIterator{contract: _Broker.contract, event: "LogInterchainData", logs: logs, sub: sub}, nil
}

// WatchLogInterchainData is a free log subscription operation binding the contract event 0x436160f7c24c5f31561ec9422a629accdbbd4e9e8ce21e86e634f497997769a8.
//
// Solidity: event LogInterchainData(bool status, string data)
func (_Broker *BrokerFilterer) WatchLogInterchainData(opts *bind.WatchOpts, sink chan<- *BrokerLogInterchainData) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "LogInterchainData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerLogInterchainData)
				if err := _Broker.contract.UnpackLog(event, "LogInterchainData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogInterchainData is a log parse operation binding the contract event 0x436160f7c24c5f31561ec9422a629accdbbd4e9e8ce21e86e634f497997769a8.
//
// Solidity: event LogInterchainData(bool status, string data)
func (_Broker *BrokerFilterer) ParseLogInterchainData(log types.Log) (*BrokerLogInterchainData, error) {
	event := new(BrokerLogInterchainData)
	if err := _Broker.contract.UnpackLog(event, "LogInterchainData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerLogInterchainStatusIterator is returned from FilterLogInterchainStatus and is used to iterate over the raw logs and unpacked data for LogInterchainStatus events raised by the Broker contract.
type BrokerLogInterchainStatusIterator struct {
	Event *BrokerLogInterchainStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerLogInterchainStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerLogInterchainStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerLogInterchainStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerLogInterchainStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerLogInterchainStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerLogInterchainStatus represents a LogInterchainStatus event raised by the Broker contract.
type BrokerLogInterchainStatus struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogInterchainStatus is a free log retrieval operation binding the contract event 0x23de11857b4338b8e6ccaec81162b447b44040ff3cfdd1174d548975eb5c1c3e.
//
// Solidity: event LogInterchainStatus(bool status)
func (_Broker *BrokerFilterer) FilterLogInterchainStatus(opts *bind.FilterOpts) (*BrokerLogInterchainStatusIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "LogInterchainStatus")
	if err != nil {
		return nil, err
	}
	return &BrokerLogInterchainStatusIterator{contract: _Broker.contract, event: "LogInterchainStatus", logs: logs, sub: sub}, nil
}

// WatchLogInterchainStatus is a free log subscription operation binding the contract event 0x23de11857b4338b8e6ccaec81162b447b44040ff3cfdd1174d548975eb5c1c3e.
//
// Solidity: event LogInterchainStatus(bool status)
func (_Broker *BrokerFilterer) WatchLogInterchainStatus(opts *bind.WatchOpts, sink chan<- *BrokerLogInterchainStatus) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "LogInterchainStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerLogInterchainStatus)
				if err := _Broker.contract.UnpackLog(event, "LogInterchainStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogInterchainStatus is a log parse operation binding the contract event 0x23de11857b4338b8e6ccaec81162b447b44040ff3cfdd1174d548975eb5c1c3e.
//
// Solidity: event LogInterchainStatus(bool status)
func (_Broker *BrokerFilterer) ParseLogInterchainStatus(log types.Log) (*BrokerLogInterchainStatus, error) {
	event := new(BrokerLogInterchainStatus)
	if err := _Broker.contract.UnpackLog(event, "LogInterchainStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerThrowEventIterator is returned from FilterThrowEvent and is used to iterate over the raw logs and unpacked data for ThrowEvent events raised by the Broker contract.
type BrokerThrowEventIterator struct {
	Event *BrokerThrowEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerThrowEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerThrowEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerThrowEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerThrowEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerThrowEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerThrowEvent represents a ThrowEvent event raised by the Broker contract.
type BrokerThrowEvent struct {
	Index   uint64
	DestDID string
	Fid     common.Address
	Funcs   string
	Args    string
	Argscb  string
	Argsrb  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterThrowEvent is a free log retrieval operation binding the contract event 0x85026d9efdca77567284833e9e5fa5ab0d412e8d0d9869942def88510971977c.
//
// Solidity: event throwEvent(uint64 index, string destDID, address fid, string funcs, string args, string argscb, string argsrb)
func (_Broker *BrokerFilterer) FilterThrowEvent(opts *bind.FilterOpts) (*BrokerThrowEventIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "throwEvent")
	if err != nil {
		return nil, err
	}
	return &BrokerThrowEventIterator{contract: _Broker.contract, event: "throwEvent", logs: logs, sub: sub}, nil
}

// WatchThrowEvent is a free log subscription operation binding the contract event 0x85026d9efdca77567284833e9e5fa5ab0d412e8d0d9869942def88510971977c.
//
// Solidity: event throwEvent(uint64 index, string destDID, address fid, string funcs, string args, string argscb, string argsrb)
func (_Broker *BrokerFilterer) WatchThrowEvent(opts *bind.WatchOpts, sink chan<- *BrokerThrowEvent) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "throwEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerThrowEvent)
				if err := _Broker.contract.UnpackLog(event, "throwEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThrowEvent is a log parse operation binding the contract event 0x85026d9efdca77567284833e9e5fa5ab0d412e8d0d9869942def88510971977c.
//
// Solidity: event throwEvent(uint64 index, string destDID, address fid, string funcs, string args, string argscb, string argsrb)
func (_Broker *BrokerFilterer) ParseThrowEvent(log types.Log) (*BrokerThrowEvent, error) {
	event := new(BrokerThrowEvent)
	if err := _Broker.contract.UnpackLog(event, "throwEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
