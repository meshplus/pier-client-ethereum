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
const BrokerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogInterchainData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LogInterchainStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fid\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"funcs\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"argscb\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"argsrb\",\"type\":\"string\"}],\"name\":\"throwEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"status\",\"type\":\"int64\"}],\"name\":\"audit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destChainID\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funcs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argscb\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argsrb\",\"type\":\"string\"}],\"name\":\"emitInterchainEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallbackMeta\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getInMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInnerMeta\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getOutMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOuterMeta\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"req\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"err\",\"type\":\"string\"}],\"name\":\"invokeIndexUpdateWithError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"req\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"bizCallData\",\"type\":\"bytes\"}],\"name\":\"invokeInterchain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BrokerFuncSigs maps the 4-byte function signature to its string representation.
var BrokerFuncSigs = map[string]string{
	"b38ff85f": "audit(address,int64)",
	"50b0b902": "emitInterchainEvent(address,string,string,string,string,string)",
	"3b6bbe4a": "getCallbackMeta()",
	"83c44c27": "getInMessage(address,uint64)",
	"67b9fa3b": "getInnerMeta()",
	"a0342a3f": "getOutMessage(address,uint64)",
	"c20cab50": "getOuterMeta()",
	"8129fc1c": "initialize()",
	"7cf636ce": "invokeIndexUpdateWithError(address,uint64,bool,string)",
	"3aabe619": "invokeInterchain(address,uint64,address,bool,bytes)",
	"4420e486": "register(address)",
}

// BrokerBin is the compiled bytecode used for deploying new contracts.
var BrokerBin = "0x608060405234801561001057600080fd5b506116bd806100206000396000f3fe60806040526004361061009c5760003560e01c80637cf636ce116100645780637cf636ce1461050d5780638129fc1c146105e557806383c44c27146105fa578063a0342a3f1461064e578063b38ff85f14610690578063c20cab50146106e05761009c565b80633aabe619146100a15780633b6bbe4a146101425780634420e486146101f057806350b0b9021461022357806367b9fa3b146104f8575b600080fd5b610140600480360360a08110156100b757600080fd5b6001600160a01b0382358116926001600160401b036020820135169260408201359092169160608201351515919081019060a081016080820135600160201b81111561010257600080fd5b82018360208201111561011457600080fd5b803590602001918460018302840111600160201b8311171561013557600080fd5b5090925090506106f5565b005b34801561014e57600080fd5b5061015761077c565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561019b578181015183820152602001610183565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156101da5781810151838201526020016101c2565b5050505090500194505050505060405180910390f35b3480156101fc57600080fd5b506101406004803603602081101561021357600080fd5b50356001600160a01b03166108c3565b34801561022f57600080fd5b50610140600480360360c081101561024657600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460018302840111600160201b831117156102a357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102f557600080fd5b82018360208201111561030757600080fd5b803590602001918460018302840111600160201b8311171561032857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561037a57600080fd5b82018360208201111561038c57600080fd5b803590602001918460018302840111600160201b831117156103ad57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103ff57600080fd5b82018360208201111561041157600080fd5b803590602001918460018302840111600160201b8311171561043257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561048457600080fd5b82018360208201111561049657600080fd5b803590602001918460018302840111600160201b831117156104b757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108eb945050505050565b34801561050457600080fd5b50610157610c8e565b34801561051957600080fd5b506101406004803603608081101561053057600080fd5b6001600160a01b03823516916001600160401b0360208201351691604082013515159190810190608081016060820135600160201b81111561057157600080fd5b82018360208201111561058357600080fd5b803590602001918460018302840111600160201b831117156105a457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dbf945050505050565b3480156105f157600080fd5b50610140610dd1565b34801561060657600080fd5b5061063c6004803603604081101561061d57600080fd5b5080356001600160a01b031690602001356001600160401b0316610fae565b60408051918252519081900360200190f35b34801561065a57600080fd5b5061063c6004803603604081101561067157600080fd5b5080356001600160a01b031690602001356001600160401b0316610fe2565b34801561069c57600080fd5b506106cc600480360360408110156106b357600080fd5b506001600160a01b03813516906020013560070b611015565b604080519115158252519081900360200190f35b3480156106ec57600080fd5b506101576110e5565b6001600160a01b038416600090815260208190526040902054600790810b900b60011461072157600080fd5b61073c8686856040518060200160405280600081525061122a565b604051602060848237805160208160040183378151808260240184376000808285348b5af19150503d806000843e818015610778578184a08184f35b8184fd5b60608060006005805490506001600160401b038111801561079c57600080fd5b506040519080825280602002602001820160405280156107c6578160200160208202803683370190505b50905060005b6005546001600160401b038216101561085a57600a60006005836001600160401b0316815481106107f957fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b0391821691849190841690811061083a57fe5b6001600160401b03909216602092830291909101909101526001016107cc565b50600581818054806020026020016040519081016040528092919081815260200182805480156108b357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610895575b5050505050915092509250509091565b6001600160a01b03166000908152602081905260409020805467ffffffffffffffff19169055565b33600090815260208190526040902054600790810b900b600114610956576040805162461bcd60e51b815260206004820152601d60248201527f496e766f6b657220617265206e6f7420696e207768697465206c697374000000604482015290519081900360640190fd5b6001600160a01b0386166000908152600660205260409020805467ffffffffffffffff19811660016001600160401b03928316810183169190911792839055911614156109e957600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b0388161790555b6001600160a01b0386166000818152600760209081526040808320600680845282852080546001600160401b03908116875292855283862043905586865290845254825191168082528184019590955233918101829052610100606082018181528b51918301919091528a517fdde0d454bdf1d147a0842ac1864ecc133506af30efc60d34dabc910267c4e40a96958d958d948d948d948d948d94608085019260a086019260c087019260e0880192610120890192918e01918190849084905b83811015610ac1578181015183820152602001610aa9565b50505050905090810190601f168015610aee5780820380516001836020036101000a031916815260200191505b5086810385528a5181528a516020918201918c019080838360005b83811015610b21578181015183820152602001610b09565b50505050905090810190601f168015610b4e5780820380516001836020036101000a031916815260200191505b5086810384528951815289516020918201918b019080838360005b83811015610b81578181015183820152602001610b69565b50505050905090810190601f168015610bae5780820380516001836020036101000a031916815260200191505b5086810383528851815288516020918201918a019080838360005b83811015610be1578181015183820152602001610bc9565b50505050905090810190601f168015610c0e5780820380516001836020036101000a031916815260200191505b50868103825287518152875160209182019189019080838360005b83811015610c41578181015183820152602001610c29565b50505050905090810190601f168015610c6e5780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060405180910390a1505050505050565b60608060006004805490506001600160401b0381118015610cae57600080fd5b50604051908082528060200260200182016040528015610cd8578160200160208202803683370190505b50905060005b600454811015610d58576008600060048381548110610cf957fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b0390911690839083908110610d3857fe5b6001600160401b0390921660209283029190910190910152600101610cde565b50600481818054806020026020016040519081016040528092919081815260200182805480156108b3576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610895575050505050915092509250509091565b610dcb8484848461122a565b50505050565b60005b600454811015610e3d5760006008600060048481548110610df157fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805467ffffffffffffffff19166001600160401b0392909216919091179055600101610dd4565b5060005b600354811015610eaa5760006006600060038481548110610e5e57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805467ffffffffffffffff19166001600160401b0392909216919091179055600101610e41565b5060005b600554811015610f17576000600a600060058481548110610ecb57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805467ffffffffffffffff19166001600160401b0392909216919091179055600101610eae565b5060005b600154811015610f8757600080600060018481548110610f3757fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020805460079290920b6001600160401b031667ffffffffffffffff19909216919091179055600101610f1b565b50610f94600360006115c5565b610fa0600460006115c5565b610fac600560006115c5565b565b6001600160a01b03821660009081526009602090815260408083206001600160401b03851684529091529020545b92915050565b6001600160a01b03821660009081526007602090815260408083206001600160401b038516845290915290205492915050565b60008160070b6000191415801561103057508160070b600014155b801561104057508160070b600114155b1561104d57506000610fdc565b6001600160a01b0383166000908152602081905260409020805467ffffffffffffffff19166001600160401b03600785900b90811691909117909155600114156110dc576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0385161790555b50600192915050565b60608060006003805490506001600160401b038111801561110557600080fd5b5060405190808252806020026020018201604052801561112f578160200160208202803683370190505b50905060005b6003546001600160401b03821610156111c357600660006003836001600160401b03168154811061116257fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b039182169184919084169081106111a357fe5b6001600160401b0390921660209283029190910190910152600101611135565b50600381818054806020026020016040519081016040528092919081815260200182805480156108b3576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610895575050505050915092509250509091565b8115611335576001600160a01b0384166000908152600860205260409020546001600160401b0390811660010181169084161461126657600080fd5b61126f8461143b565b604080516000815260208082018084528251902084519093859301918291908401908083835b602083106112b45780518252601f199092019160209182019101611295565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014611330576001600160a01b0384166000908152600b602090815260408083206001600160401b03871684528252909120825161132e928401906115e6565b505b610dcb565b6001600160a01b0384166000908152600a60205260409020546001600160401b0390811660010181169084161461136b57600080fd5b6113758484611505565b604080516000815260208082018084528251902084519093859301918291908401908083835b602083106113ba5780518252601f19909201916020918201910161139b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014610dcb576001600160a01b0384166000908152600c602090815260408083206001600160401b038716845282529091208251611434928401906115e6565b5050505050565b6001600160a01b0381166000908152600860205260409020805467ffffffffffffffff19811660016001600160401b03928316810183169190911792839055911614156114ce57600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319166001600160a01b0383161790555b6001600160a01b0316600090815260096020908152604080832060088352818420546001600160401b031684529091529020439055565b6001600160a01b0382166000908152600a60205260409020546001600160401b031661157757600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319166001600160a01b0384161790555b6001600160a01b03919091166000908152600a60209081526040808320805467ffffffffffffffff19166001600160401b039586161790819055600983528184209416835292905220439055565b50805460008255906000526020600020908101906115e39190611672565b50565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928261161c5760008555611662565b82601f1061163557805160ff1916838001178555611662565b82800160010185558215611662579182015b82811115611662578251825591602001919060010190611647565b5061166e929150611672565b5090565b5b8082111561166e576000815560010161167356fea264697066735822122049275507c26b3ece6e129c5a51778c31e2c03d16894143088b81a74d1dc4444d64736f6c63430007060033"

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

// GetCallbackMeta is a free data retrieval call binding the contract method 0x3b6bbe4a.
//
// Solidity: function getCallbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerCaller) GetCallbackMeta(opts *bind.CallOpts) ([]common.Address, []uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCallbackMeta")

	if err != nil {
		return *new([]common.Address), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetCallbackMeta is a free data retrieval call binding the contract method 0x3b6bbe4a.
//
// Solidity: function getCallbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerSession) GetCallbackMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetCallbackMeta(&_Broker.CallOpts)
}

// GetCallbackMeta is a free data retrieval call binding the contract method 0x3b6bbe4a.
//
// Solidity: function getCallbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerCallerSession) GetCallbackMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetCallbackMeta(&_Broker.CallOpts)
}

// GetInMessage is a free data retrieval call binding the contract method 0x83c44c27.
//
// Solidity: function getInMessage(address from, uint64 idx) view returns(uint256)
func (_Broker *BrokerCaller) GetInMessage(opts *bind.CallOpts, from common.Address, idx uint64) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getInMessage", from, idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInMessage is a free data retrieval call binding the contract method 0x83c44c27.
//
// Solidity: function getInMessage(address from, uint64 idx) view returns(uint256)
func (_Broker *BrokerSession) GetInMessage(from common.Address, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetInMessage(&_Broker.CallOpts, from, idx)
}

// GetInMessage is a free data retrieval call binding the contract method 0x83c44c27.
//
// Solidity: function getInMessage(address from, uint64 idx) view returns(uint256)
func (_Broker *BrokerCallerSession) GetInMessage(from common.Address, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetInMessage(&_Broker.CallOpts, from, idx)
}

// GetInnerMeta is a free data retrieval call binding the contract method 0x67b9fa3b.
//
// Solidity: function getInnerMeta() view returns(address[], uint64[])
func (_Broker *BrokerCaller) GetInnerMeta(opts *bind.CallOpts) ([]common.Address, []uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getInnerMeta")

	if err != nil {
		return *new([]common.Address), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetInnerMeta is a free data retrieval call binding the contract method 0x67b9fa3b.
//
// Solidity: function getInnerMeta() view returns(address[], uint64[])
func (_Broker *BrokerSession) GetInnerMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetInnerMeta(&_Broker.CallOpts)
}

// GetInnerMeta is a free data retrieval call binding the contract method 0x67b9fa3b.
//
// Solidity: function getInnerMeta() view returns(address[], uint64[])
func (_Broker *BrokerCallerSession) GetInnerMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetInnerMeta(&_Broker.CallOpts)
}

// GetOutMessage is a free data retrieval call binding the contract method 0xa0342a3f.
//
// Solidity: function getOutMessage(address to, uint64 idx) view returns(uint256)
func (_Broker *BrokerCaller) GetOutMessage(opts *bind.CallOpts, to common.Address, idx uint64) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getOutMessage", to, idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutMessage is a free data retrieval call binding the contract method 0xa0342a3f.
//
// Solidity: function getOutMessage(address to, uint64 idx) view returns(uint256)
func (_Broker *BrokerSession) GetOutMessage(to common.Address, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetOutMessage(&_Broker.CallOpts, to, idx)
}

// GetOutMessage is a free data retrieval call binding the contract method 0xa0342a3f.
//
// Solidity: function getOutMessage(address to, uint64 idx) view returns(uint256)
func (_Broker *BrokerCallerSession) GetOutMessage(to common.Address, idx uint64) (*big.Int, error) {
	return _Broker.Contract.GetOutMessage(&_Broker.CallOpts, to, idx)
}

// GetOuterMeta is a free data retrieval call binding the contract method 0xc20cab50.
//
// Solidity: function getOuterMeta() view returns(address[], uint64[])
func (_Broker *BrokerCaller) GetOuterMeta(opts *bind.CallOpts) ([]common.Address, []uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getOuterMeta")

	if err != nil {
		return *new([]common.Address), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetOuterMeta is a free data retrieval call binding the contract method 0xc20cab50.
//
// Solidity: function getOuterMeta() view returns(address[], uint64[])
func (_Broker *BrokerSession) GetOuterMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetOuterMeta(&_Broker.CallOpts)
}

// GetOuterMeta is a free data retrieval call binding the contract method 0xc20cab50.
//
// Solidity: function getOuterMeta() view returns(address[], uint64[])
func (_Broker *BrokerCallerSession) GetOuterMeta() ([]common.Address, []uint64, error) {
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

// EmitInterchainEvent is a paid mutator transaction binding the contract method 0x50b0b902.
//
// Solidity: function emitInterchainEvent(address destChainID, string destAddr, string funcs, string args, string argscb, string argsrb) returns()
func (_Broker *BrokerTransactor) EmitInterchainEvent(opts *bind.TransactOpts, destChainID common.Address, destAddr string, funcs string, args string, argscb string, argsrb string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "emitInterchainEvent", destChainID, destAddr, funcs, args, argscb, argsrb)
}

// EmitInterchainEvent is a paid mutator transaction binding the contract method 0x50b0b902.
//
// Solidity: function emitInterchainEvent(address destChainID, string destAddr, string funcs, string args, string argscb, string argsrb) returns()
func (_Broker *BrokerSession) EmitInterchainEvent(destChainID common.Address, destAddr string, funcs string, args string, argscb string, argsrb string) (*types.Transaction, error) {
	return _Broker.Contract.EmitInterchainEvent(&_Broker.TransactOpts, destChainID, destAddr, funcs, args, argscb, argsrb)
}

// EmitInterchainEvent is a paid mutator transaction binding the contract method 0x50b0b902.
//
// Solidity: function emitInterchainEvent(address destChainID, string destAddr, string funcs, string args, string argscb, string argsrb) returns()
func (_Broker *BrokerTransactorSession) EmitInterchainEvent(destChainID common.Address, destAddr string, funcs string, args string, argscb string, argsrb string) (*types.Transaction, error) {
	return _Broker.Contract.EmitInterchainEvent(&_Broker.TransactOpts, destChainID, destAddr, funcs, args, argscb, argsrb)
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

// InvokeIndexUpdateWithError is a paid mutator transaction binding the contract method 0x7cf636ce.
//
// Solidity: function invokeIndexUpdateWithError(address srcChainID, uint64 index, bool req, string err) returns()
func (_Broker *BrokerTransactor) InvokeIndexUpdateWithError(opts *bind.TransactOpts, srcChainID common.Address, index uint64, req bool, err string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "invokeIndexUpdateWithError", srcChainID, index, req, err)
}

// InvokeIndexUpdateWithError is a paid mutator transaction binding the contract method 0x7cf636ce.
//
// Solidity: function invokeIndexUpdateWithError(address srcChainID, uint64 index, bool req, string err) returns()
func (_Broker *BrokerSession) InvokeIndexUpdateWithError(srcChainID common.Address, index uint64, req bool, err string) (*types.Transaction, error) {
	return _Broker.Contract.InvokeIndexUpdateWithError(&_Broker.TransactOpts, srcChainID, index, req, err)
}

// InvokeIndexUpdateWithError is a paid mutator transaction binding the contract method 0x7cf636ce.
//
// Solidity: function invokeIndexUpdateWithError(address srcChainID, uint64 index, bool req, string err) returns()
func (_Broker *BrokerTransactorSession) InvokeIndexUpdateWithError(srcChainID common.Address, index uint64, req bool, err string) (*types.Transaction, error) {
	return _Broker.Contract.InvokeIndexUpdateWithError(&_Broker.TransactOpts, srcChainID, index, req, err)
}

// InvokeInterchain is a paid mutator transaction binding the contract method 0x3aabe619.
//
// Solidity: function invokeInterchain(address srcChainID, uint64 index, address destAddr, bool req, bytes bizCallData) payable returns()
func (_Broker *BrokerTransactor) InvokeInterchain(opts *bind.TransactOpts, srcChainID common.Address, index uint64, destAddr common.Address, req bool, bizCallData []byte) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "invokeInterchain", srcChainID, index, destAddr, req, bizCallData)
}

// InvokeInterchain is a paid mutator transaction binding the contract method 0x3aabe619.
//
// Solidity: function invokeInterchain(address srcChainID, uint64 index, address destAddr, bool req, bytes bizCallData) payable returns()
func (_Broker *BrokerSession) InvokeInterchain(srcChainID common.Address, index uint64, destAddr common.Address, req bool, bizCallData []byte) (*types.Transaction, error) {
	return _Broker.Contract.InvokeInterchain(&_Broker.TransactOpts, srcChainID, index, destAddr, req, bizCallData)
}

// InvokeInterchain is a paid mutator transaction binding the contract method 0x3aabe619.
//
// Solidity: function invokeInterchain(address srcChainID, uint64 index, address destAddr, bool req, bytes bizCallData) payable returns()
func (_Broker *BrokerTransactorSession) InvokeInterchain(srcChainID common.Address, index uint64, destAddr common.Address, req bool, bizCallData []byte) (*types.Transaction, error) {
	return _Broker.Contract.InvokeInterchain(&_Broker.TransactOpts, srcChainID, index, destAddr, req, bizCallData)
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
	Index  uint64
	To     common.Address
	Fid    common.Address
	Tid    string
	Funcs  string
	Args   string
	Argscb string
	Argsrb string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterThrowEvent is a free log retrieval operation binding the contract event 0xdde0d454bdf1d147a0842ac1864ecc133506af30efc60d34dabc910267c4e40a.
//
// Solidity: event throwEvent(uint64 index, address to, address fid, string tid, string funcs, string args, string argscb, string argsrb)
func (_Broker *BrokerFilterer) FilterThrowEvent(opts *bind.FilterOpts) (*BrokerThrowEventIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "throwEvent")
	if err != nil {
		return nil, err
	}
	return &BrokerThrowEventIterator{contract: _Broker.contract, event: "throwEvent", logs: logs, sub: sub}, nil
}

// WatchThrowEvent is a free log subscription operation binding the contract event 0xdde0d454bdf1d147a0842ac1864ecc133506af30efc60d34dabc910267c4e40a.
//
// Solidity: event throwEvent(uint64 index, address to, address fid, string tid, string funcs, string args, string argscb, string argsrb)
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

// ParseThrowEvent is a log parse operation binding the contract event 0xdde0d454bdf1d147a0842ac1864ecc133506af30efc60d34dabc910267c4e40a.
//
// Solidity: event throwEvent(uint64 index, address to, address fid, string tid, string funcs, string args, string argscb, string argsrb)
func (_Broker *BrokerFilterer) ParseThrowEvent(log types.Log) (*BrokerThrowEvent, error) {
	event := new(BrokerThrowEvent)
	if err := _Broker.contract.UnpackLog(event, "throwEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
