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
const BrokerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"destChainID\",\"type\":\"address\"},{\"name\":\"destAddr\",\"type\":\"string\"},{\"name\":\"args\",\"type\":\"string\"}],\"name\":\"InterchainTransferInvoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"interchainSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCallbackMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"interchainGet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destChainID\",\"type\":\"address\"},{\"name\":\"destAddr\",\"type\":\"string\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"InterchainDataSwapInvoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInnerMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getInMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getOutMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"int64\"}],\"name\":\"audit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"},{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"interchainConfirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"interchainCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOuterMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fid\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tid\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"func\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"args\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"callback\",\"type\":\"string\"}],\"name\":\"throwEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogInterchainData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LogInterchainStatus\",\"type\":\"event\"}]"

// BrokerFuncSigs maps the 4-byte function signature to its string representation.
var BrokerFuncSigs = map[string]string{
	"5e7d7c4c": "InterchainDataSwapInvoke(address,string,string)",
	"06bef67c": "InterchainTransferInvoke(address,string,string)",
	"b38ff85f": "audit(address,int64)",
	"3b6bbe4a": "getCallbackMeta()",
	"83c44c27": "getInMessage(address,uint64)",
	"67b9fa3b": "getInnerMeta()",
	"a0342a3f": "getOutMessage(address,uint64)",
	"c20cab50": "getOuterMeta()",
	"8129fc1c": "initialize()",
	"befbf664": "interchainCharge(address,uint64,address,string,string,uint64)",
	"be7c4222": "interchainConfirm(address,uint64,address,bool,string,uint64)",
	"3c25819a": "interchainGet(address,uint64,address,string)",
	"19ba2f2f": "interchainSet(address,uint64,address,string,string)",
	"4420e486": "register(address)",
}

// BrokerBin is the compiled bytecode used for deploying new contracts.
var BrokerBin = "0x608060405234801561001057600080fd5b50611f4e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638129fc1c1161008c578063b38ff85f11610066578063b38ff85f1461076a578063be7c422214610799578063befbf66414610879578063c20cab50146109d5576100ea565b80638129fc1c146106e657806383c44c27146106ee578063a0342a3f14610735576100ea565b80633c25819a116100c85780633c25819a1461042e5780634420e4861461057d5780635e7d7c4c146105a557806367b9fa3b146106de576100ea565b806306bef67c146100ef57806319ba2f2f1461023c5780633b6bbe4a1461038d575b600080fd5b6102286004803603606081101561010557600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561012f57600080fd5b82018360208201111561014157600080fd5b803590602001918460018302840111600160201b8311171561016257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101b457600080fd5b8201836020820111156101c657600080fd5b803590602001918460018302840111600160201b831117156101e757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109dd945050505050565b604080519115158252519081900360200190f35b610228600480360360a081101561025257600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b81111561029457600080fd5b8201836020820111156102a657600080fd5b803590602001918460018302840111600160201b831117156102c757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561031957600080fd5b82018360208201111561032b57600080fd5b803590602001918460018302840111600160201b8311171561034c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610acb945050505050565b610395610c8e565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103d95781810151838201526020016103c1565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610418578181015183820152602001610400565b5050505090500194505050505060405180910390f35b6104fa6004803603608081101561044457600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b81111561048657600080fd5b82018360208201111561049857600080fd5b803590602001918460018302840111600160201b831117156104b957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dbf945050505050565b604051808315151515815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610541578181015183820152602001610529565b50505050905090810190601f16801561056e5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6105a36004803603602081101561059357600080fd5b50356001600160a01b0316611001565b005b610228600480360360608110156105bb57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105e557600080fd5b8201836020820111156105f757600080fd5b803590602001918460018302840111600160201b8311171561061857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561066a57600080fd5b82018360208201111561067c57600080fd5b803590602001918460018302840111600160201b8311171561069d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611029945050505050565b61039561110f565b6105a361122a565b6107236004803603604081101561070457600080fd5b5080356001600160a01b031690602001356001600160401b031661140d565b60408051918252519081900360200190f35b6107236004803603604081101561074b57600080fd5b5080356001600160a01b031690602001356001600160401b0316611441565b6102286004803603604081101561078057600080fd5b506001600160a01b03813516906020013560070b611474565b610228600480360360c08110156107af57600080fd5b6001600160a01b0382358116926001600160401b036020820135169260408201359092169160608201351515919081019060a081016080820135600160201b8111156107fa57600080fd5b82018360208201111561080c57600080fd5b803590602001918460018302840111600160201b8311171561082d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506115449050565b610228600480360360c081101561088f57600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b8111156108d157600080fd5b8201836020820111156108e357600080fd5b803590602001918460018302840111600160201b8311171561090457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561095657600080fd5b82018360208201111561096857600080fd5b803590602001918460018302840111600160201b8311171561098957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b0316915061174a9050565b610395611928565b33600090815260208190526040812054600790810b900b600114610a4b5760408051600160e51b62461bcd02815260206004820152601d60248201527f496e766f6b657220617265206e6f7420696e207768697465206c697374000000604482015290519081900360640190fd5b610ac38433856040518060400160405280601081526020017f696e746572636861696e43686172676500000000000000000000000000000000815250866040518060400160405280601181526020017f696e746572636861696e436f6e6669726d000000000000000000000000000000815250611a57565b949350505050565b6001600160a01b0385166000908152600a60205260408120546001600160401b03908116600101811690861614610b265760408051600081529051600080516020611f038339815191529181900360200190a1506000610c85565b83610b318787611d2e565b60408051600160e41b63093c00ab028152600481019182528551604482015285516001600160a01b038416926393c00ab09288928892918291602482019160640190602087019080838360005b83811015610b96578181015183820152602001610b7e565b50505050905090810190601f168015610bc35780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610bf6578181015183820152602001610bde565b50505050905090810190601f168015610c235780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610c4457600080fd5b505af1158015610c58573d6000803e3d6000fd5b505060408051600181529051600080516020611f038339815191529350908190036020019150a160019150505b95945050505050565b6060806060600580549050604051908082528060200260200182016040528015610cc2578160200160208202803883390190505b50905060005b6005546001600160401b0382161015610d5657600a60006005836001600160401b031681548110610cf557fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b03918216918491908416908110610d3657fe5b6001600160401b0390921660209283029190910190910152600101610cc8565b5060058181805480602002602001604051908101604052809291908181526020018280548015610daf57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d91575b5050505050915092509250509091565b6000606083610dcd87611dee565b6001600160a01b038516600090815260208190526040902054600790810b900b600114610e0d575050604080516020810190915260008082529150610ff8565b60606000826001600160a01b0316636079cf2a876040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e6d578181015183820152602001610e55565b50505050905090810190601f168015610e9a5780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b158015610eb757600080fd5b505afa158015610ecb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610ef457600080fd5b815160208301805191939283019291600160201b811115610f1457600080fd5b82016020810184811115610f2757600080fd5b8151600160201b811182820187101715610f4057600080fd5b5050604080518615158152602080820183815284519383019390935283519399509697507f436160f7c24c5f31561ec9422a629accdbbd4e9e8ce21e86e634f497997769a8968896508995509093509091606084019185019080838360005b83811015610fb7578181015183820152602001610f9f565b50505050905090810190601f168015610fe45780820380516001836020036101000a031916815260200191505b50935050505060405180910390a193509150505b94509492505050565b6001600160a01b03166000908152602081905260409020805467ffffffffffffffff19169055565b33600090815260208190526040812054600790810b900b6001146110975760408051600160e51b62461bcd02815260206004820152601d60248201527f496e766f6b657220617265206e6f7420696e207768697465206c697374000000604482015290519081900360640190fd5b610ac38433856040518060400160405280600d81526020017f696e746572636861696e47657400000000000000000000000000000000000000815250866040518060400160405280600d81526020017f696e746572636861696e53657400000000000000000000000000000000000000815250611a57565b6060806060600480549050604051908082528060200260200182016040528015611143578160200160208202803883390190505b50905060005b6004548110156111c357600860006004838154811061116457fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b03909116908390839081106111a357fe5b6001600160401b0390921660209283029190910190910152600101611149565b5060048181805480602002602001604051908101604052809291908181526020018280548015610daf576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610d91575050505050915092509250509091565b60005b600454811015611296576000600860006004848154811061124a57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805467ffffffffffffffff19166001600160401b039290921691909117905560010161122d565b5060005b60035481101561130357600060066000600384815481106112b757fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805467ffffffffffffffff19166001600160401b039290921691909117905560010161129a565b5060005b600554811015611370576000600a60006005848154811061132457fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805467ffffffffffffffff19166001600160401b0392909216919091179055600101611307565b5060005b6001548110156113e05760008060006001848154811061139057fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020805460079290920b6001600160401b031667ffffffffffffffff19909216919091179055600101611374565b5060006113ee600382611eb8565b5060006113fc600482611eb8565b50600061140a600582611eb8565b50565b6001600160a01b03821660009081526009602090815260408083206001600160401b03851684529091529020545b92915050565b6001600160a01b03821660009081526007602090815260408083206001600160401b038516845290915290205492915050565b60008160070b6000191415801561148f57508160070b600014155b801561149f57508160070b600114155b156114ac5750600061143b565b6001600160a01b0383166000908152602081905260409020805467ffffffffffffffff19166001600160401b03600785900b908116919091179091556001141561153b576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0385161790555b50600192915050565b6001600160a01b0386166000908152600a60205260408120546001600160401b0390811660010181169087161461159f5760408051600081529051600080516020611f038339815191529181900360200190a1506000611740565b6115a98787611d2e565b6001600160a01b038516600090815260208190526040902054600790810b900b6001146115fa5760408051600081529051600080516020611f038339815191529181900360200190a1506000611740565b831561162a5760408051600181529051600080516020611f038339815191529181900360200190a1506001611740565b60408051600160e21b6311a1f28d0281526001600160401b03841660248201526004810191825284516044820152845187926000926001600160a01b03851692634687ca34928992899282916064019060208601908083838c5b8381101561169c578181015183820152602001611684565b50505050905090810190601f1680156116c95780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156116e957600080fd5b505af11580156116fd573d6000803e3d6000fd5b505050506040513d602081101561171357600080fd5b50516040805182151581529051919250600080516020611f03833981519152919081900360200190a19150505b9695505050505050565b6001600160a01b0386166000908152600860205260408120546001600160401b039081166001018116908716146117a55760408051600081529051600080516020611f038339815191529181900360200190a1506000611740565b6117ae87611dee565b6001600160a01b038516600090815260208190526040902054600790810b900b6001146117ff5760408051600081529051600080516020611f038339815191529181900360200190a1506000611740565b604051600160e01b63e01b35170281526001600160401b038316604482015260606004820190815285516064830152855187926000926001600160a01b0385169263e01b3517928a928a928a92829160248201916084019060208801908083838e5b83811015611879578181015183820152602001611861565b50505050905090810190601f1680156118a65780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156118d95781810151838201526020016118c1565b50505050905090810190601f1680156119065780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156116e957600080fd5b606080606060038054905060405190808252806020026020018201604052801561195c578160200160208202803883390190505b50905060005b6003546001600160401b03821610156119f057600660006003836001600160401b03168154811061198f57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b039182169184919084169081106119d057fe5b6001600160401b0390921660209283029190910190910152600101611962565b5060038181805480602002602001604051908101604052809291908181526020018280548015610daf576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610d91575050505050915092509250509091565b6001600160a01b0386166000908152600660205260408120805467ffffffffffffffff19811660016001600160401b0392831681018316919091179283905591161415611aea57600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b0389161790555b6001600160a01b038088166000818152600760209081526040808320600680845282852080546001600160401b039081168752928552838620439055868652908452548251911680825281840195909552948b169085015260e0606085018181528a519186019190915289517fad89cfa05a757be8d2179bb6609bf9034971b2427bd49d48e79552d3e8493e99958d948d948d948d948d948d949093608085019260a086019260c0870192610100880192908c01918190849084905b83811015611bbe578181015183820152602001611ba6565b50505050905090810190601f168015611beb5780820380516001836020036101000a031916815260200191505b5085810384528851815288516020918201918a019080838360005b83811015611c1e578181015183820152602001611c06565b50505050905090810190601f168015611c4b5780820380516001836020036101000a031916815260200191505b50858103835287518152875160209182019189019080838360005b83811015611c7e578181015183820152602001611c66565b50505050905090810190601f168015611cab5780820380516001836020036101000a031916815260200191505b50858103825286518152865160209182019188019080838360005b83811015611cde578181015183820152602001611cc6565b50505050905090810190601f168015611d0b5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390a15060019695505050505050565b6001600160a01b0382166000908152600a60205260409020546001600160401b0316611da057600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319166001600160a01b0384161790555b6001600160a01b03919091166000908152600a60209081526040808320805467ffffffffffffffff19166001600160401b039586161790819055600983528184209416835292905220439055565b6001600160a01b0381166000908152600860205260409020805467ffffffffffffffff19811660016001600160401b0392831681018316919091179283905591161415611e8157600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319166001600160a01b0383161790555b6001600160a01b0316600090815260096020908152604080832060088352818420546001600160401b031684529091529020439055565b815481835581811115611edc57600083815260209020611edc918101908301611ee1565b505050565b611eff91905b80821115611efb5760008155600101611ee7565b5090565b9056fe23de11857b4338b8e6ccaec81162b447b44040ff3cfdd1174d548975eb5c1c3ea165627a7a723058201bacc383d28a591bd4e8491afeffb90ff4b349a00f24168591ec0eb16e2e444a0029"

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
func (_Broker *BrokerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Broker *BrokerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new([]common.Address)
		ret1 = new([]uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Broker.contract.Call(opts, out, "getCallbackMeta")
	return *ret0, *ret1, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Broker.contract.Call(opts, out, "getInMessage", from, idx)
	return *ret0, err
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
	var (
		ret0 = new([]common.Address)
		ret1 = new([]uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Broker.contract.Call(opts, out, "getInnerMeta")
	return *ret0, *ret1, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Broker.contract.Call(opts, out, "getOutMessage", to, idx)
	return *ret0, err
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
	var (
		ret0 = new([]common.Address)
		ret1 = new([]uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Broker.contract.Call(opts, out, "getOuterMeta")
	return *ret0, *ret1, err
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

// InterchainDataSwapInvoke is a paid mutator transaction binding the contract method 0x5e7d7c4c.
//
// Solidity: function InterchainDataSwapInvoke(address destChainID, string destAddr, string key) returns(bool)
func (_Broker *BrokerTransactor) InterchainDataSwapInvoke(opts *bind.TransactOpts, destChainID common.Address, destAddr string, key string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "InterchainDataSwapInvoke", destChainID, destAddr, key)
}

// InterchainDataSwapInvoke is a paid mutator transaction binding the contract method 0x5e7d7c4c.
//
// Solidity: function InterchainDataSwapInvoke(address destChainID, string destAddr, string key) returns(bool)
func (_Broker *BrokerSession) InterchainDataSwapInvoke(destChainID common.Address, destAddr string, key string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainDataSwapInvoke(&_Broker.TransactOpts, destChainID, destAddr, key)
}

// InterchainDataSwapInvoke is a paid mutator transaction binding the contract method 0x5e7d7c4c.
//
// Solidity: function InterchainDataSwapInvoke(address destChainID, string destAddr, string key) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainDataSwapInvoke(destChainID common.Address, destAddr string, key string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainDataSwapInvoke(&_Broker.TransactOpts, destChainID, destAddr, key)
}

// InterchainTransferInvoke is a paid mutator transaction binding the contract method 0x06bef67c.
//
// Solidity: function InterchainTransferInvoke(address destChainID, string destAddr, string args) returns(bool)
func (_Broker *BrokerTransactor) InterchainTransferInvoke(opts *bind.TransactOpts, destChainID common.Address, destAddr string, args string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "InterchainTransferInvoke", destChainID, destAddr, args)
}

// InterchainTransferInvoke is a paid mutator transaction binding the contract method 0x06bef67c.
//
// Solidity: function InterchainTransferInvoke(address destChainID, string destAddr, string args) returns(bool)
func (_Broker *BrokerSession) InterchainTransferInvoke(destChainID common.Address, destAddr string, args string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainTransferInvoke(&_Broker.TransactOpts, destChainID, destAddr, args)
}

// InterchainTransferInvoke is a paid mutator transaction binding the contract method 0x06bef67c.
//
// Solidity: function InterchainTransferInvoke(address destChainID, string destAddr, string args) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainTransferInvoke(destChainID common.Address, destAddr string, args string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainTransferInvoke(&_Broker.TransactOpts, destChainID, destAddr, args)
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

// InterchainCharge is a paid mutator transaction binding the contract method 0xbefbf664.
//
// Solidity: function interchainCharge(address sourceChainID, uint64 index, address destAddr, string sender, string receiver, uint64 amount) returns(bool)
func (_Broker *BrokerTransactor) InterchainCharge(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, sender string, receiver string, amount uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainCharge", sourceChainID, index, destAddr, sender, receiver, amount)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xbefbf664.
//
// Solidity: function interchainCharge(address sourceChainID, uint64 index, address destAddr, string sender, string receiver, uint64 amount) returns(bool)
func (_Broker *BrokerSession) InterchainCharge(sourceChainID common.Address, index uint64, destAddr common.Address, sender string, receiver string, amount uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainCharge(&_Broker.TransactOpts, sourceChainID, index, destAddr, sender, receiver, amount)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xbefbf664.
//
// Solidity: function interchainCharge(address sourceChainID, uint64 index, address destAddr, string sender, string receiver, uint64 amount) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainCharge(sourceChainID common.Address, index uint64, destAddr common.Address, sender string, receiver string, amount uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainCharge(&_Broker.TransactOpts, sourceChainID, index, destAddr, sender, receiver, amount)
}

// InterchainConfirm is a paid mutator transaction binding the contract method 0xbe7c4222.
//
// Solidity: function interchainConfirm(address sourceChainID, uint64 index, address destAddr, bool status, string sender, uint64 amount) returns(bool)
func (_Broker *BrokerTransactor) InterchainConfirm(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, status bool, sender string, amount uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainConfirm", sourceChainID, index, destAddr, status, sender, amount)
}

// InterchainConfirm is a paid mutator transaction binding the contract method 0xbe7c4222.
//
// Solidity: function interchainConfirm(address sourceChainID, uint64 index, address destAddr, bool status, string sender, uint64 amount) returns(bool)
func (_Broker *BrokerSession) InterchainConfirm(sourceChainID common.Address, index uint64, destAddr common.Address, status bool, sender string, amount uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainConfirm(&_Broker.TransactOpts, sourceChainID, index, destAddr, status, sender, amount)
}

// InterchainConfirm is a paid mutator transaction binding the contract method 0xbe7c4222.
//
// Solidity: function interchainConfirm(address sourceChainID, uint64 index, address destAddr, bool status, string sender, uint64 amount) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainConfirm(sourceChainID common.Address, index uint64, destAddr common.Address, status bool, sender string, amount uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainConfirm(&_Broker.TransactOpts, sourceChainID, index, destAddr, status, sender, amount)
}

// InterchainGet is a paid mutator transaction binding the contract method 0x3c25819a.
//
// Solidity: function interchainGet(address sourceChainID, uint64 index, address destAddr, string key) returns(bool, string)
func (_Broker *BrokerTransactor) InterchainGet(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, key string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainGet", sourceChainID, index, destAddr, key)
}

// InterchainGet is a paid mutator transaction binding the contract method 0x3c25819a.
//
// Solidity: function interchainGet(address sourceChainID, uint64 index, address destAddr, string key) returns(bool, string)
func (_Broker *BrokerSession) InterchainGet(sourceChainID common.Address, index uint64, destAddr common.Address, key string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainGet(&_Broker.TransactOpts, sourceChainID, index, destAddr, key)
}

// InterchainGet is a paid mutator transaction binding the contract method 0x3c25819a.
//
// Solidity: function interchainGet(address sourceChainID, uint64 index, address destAddr, string key) returns(bool, string)
func (_Broker *BrokerTransactorSession) InterchainGet(sourceChainID common.Address, index uint64, destAddr common.Address, key string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainGet(&_Broker.TransactOpts, sourceChainID, index, destAddr, key)
}

// InterchainSet is a paid mutator transaction binding the contract method 0x19ba2f2f.
//
// Solidity: function interchainSet(address sourceChainID, uint64 index, address destAddr, string key, string value) returns(bool)
func (_Broker *BrokerTransactor) InterchainSet(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, key string, value string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainSet", sourceChainID, index, destAddr, key, value)
}

// InterchainSet is a paid mutator transaction binding the contract method 0x19ba2f2f.
//
// Solidity: function interchainSet(address sourceChainID, uint64 index, address destAddr, string key, string value) returns(bool)
func (_Broker *BrokerSession) InterchainSet(sourceChainID common.Address, index uint64, destAddr common.Address, key string, value string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainSet(&_Broker.TransactOpts, sourceChainID, index, destAddr, key, value)
}

// InterchainSet is a paid mutator transaction binding the contract method 0x19ba2f2f.
//
// Solidity: function interchainSet(address sourceChainID, uint64 index, address destAddr, string key, string value) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainSet(sourceChainID common.Address, index uint64, destAddr common.Address, key string, value string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainSet(&_Broker.TransactOpts, sourceChainID, index, destAddr, key, value)
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
	Index    uint64
	To       common.Address
	Fid      common.Address
	Tid      string
	Func     string
	Args     string
	Callback string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterThrowEvent is a free log retrieval operation binding the contract event 0xad89cfa05a757be8d2179bb6609bf9034971b2427bd49d48e79552d3e8493e99.
//
// Solidity: event throwEvent(uint64 index, address to, address fid, string tid, string func, string args, string callback)
func (_Broker *BrokerFilterer) FilterThrowEvent(opts *bind.FilterOpts) (*BrokerThrowEventIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "throwEvent")
	if err != nil {
		return nil, err
	}
	return &BrokerThrowEventIterator{contract: _Broker.contract, event: "throwEvent", logs: logs, sub: sub}, nil
}

// WatchThrowEvent is a free log subscription operation binding the contract event 0xad89cfa05a757be8d2179bb6609bf9034971b2427bd49d48e79552d3e8493e99.
//
// Solidity: event throwEvent(uint64 index, address to, address fid, string tid, string func, string args, string callback)
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

// ParseThrowEvent is a log parse operation binding the contract event 0xad89cfa05a757be8d2179bb6609bf9034971b2427bd49d48e79552d3e8493e99.
//
// Solidity: event throwEvent(uint64 index, address to, address fid, string tid, string func, string args, string callback)
func (_Broker *BrokerFilterer) ParseThrowEvent(log types.Log) (*BrokerThrowEvent, error) {
	event := new(BrokerThrowEvent)
	if err := _Broker.contract.UnpackLog(event, "throwEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataSwapperABI is the input ABI used to generate the binding from.
const DataSwapperABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"interchainGet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"interchainSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DataSwapperFuncSigs maps the 4-byte function signature to its string representation.
var DataSwapperFuncSigs = map[string]string{
	"6079cf2a": "interchainGet(string)",
	"93c00ab0": "interchainSet(string,string)",
}

// DataSwapper is an auto generated Go binding around an Ethereum contract.
type DataSwapper struct {
	DataSwapperCaller     // Read-only binding to the contract
	DataSwapperTransactor // Write-only binding to the contract
	DataSwapperFilterer   // Log filterer for contract events
}

// DataSwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataSwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataSwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataSwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSwapperSession struct {
	Contract     *DataSwapper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataSwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataSwapperCallerSession struct {
	Contract *DataSwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DataSwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataSwapperTransactorSession struct {
	Contract     *DataSwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DataSwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataSwapperRaw struct {
	Contract *DataSwapper // Generic contract binding to access the raw methods on
}

// DataSwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataSwapperCallerRaw struct {
	Contract *DataSwapperCaller // Generic read-only contract binding to access the raw methods on
}

// DataSwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataSwapperTransactorRaw struct {
	Contract *DataSwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataSwapper creates a new instance of DataSwapper, bound to a specific deployed contract.
func NewDataSwapper(address common.Address, backend bind.ContractBackend) (*DataSwapper, error) {
	contract, err := bindDataSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataSwapper{DataSwapperCaller: DataSwapperCaller{contract: contract}, DataSwapperTransactor: DataSwapperTransactor{contract: contract}, DataSwapperFilterer: DataSwapperFilterer{contract: contract}}, nil
}

// NewDataSwapperCaller creates a new read-only instance of DataSwapper, bound to a specific deployed contract.
func NewDataSwapperCaller(address common.Address, caller bind.ContractCaller) (*DataSwapperCaller, error) {
	contract, err := bindDataSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataSwapperCaller{contract: contract}, nil
}

// NewDataSwapperTransactor creates a new write-only instance of DataSwapper, bound to a specific deployed contract.
func NewDataSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*DataSwapperTransactor, error) {
	contract, err := bindDataSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataSwapperTransactor{contract: contract}, nil
}

// NewDataSwapperFilterer creates a new log filterer instance of DataSwapper, bound to a specific deployed contract.
func NewDataSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*DataSwapperFilterer, error) {
	contract, err := bindDataSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataSwapperFilterer{contract: contract}, nil
}

// bindDataSwapper binds a generic wrapper to an already deployed contract.
func bindDataSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataSwapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataSwapper *DataSwapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataSwapper.Contract.DataSwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataSwapper *DataSwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataSwapper.Contract.DataSwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataSwapper *DataSwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataSwapper.Contract.DataSwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataSwapper *DataSwapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataSwapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataSwapper *DataSwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataSwapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataSwapper *DataSwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataSwapper.Contract.contract.Transact(opts, method, params...)
}

// InterchainGet is a free data retrieval call binding the contract method 0x6079cf2a.
//
// Solidity: function interchainGet(string key) view returns(bool, string)
func (_DataSwapper *DataSwapperCaller) InterchainGet(opts *bind.CallOpts, key string) (bool, string, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DataSwapper.contract.Call(opts, out, "interchainGet", key)
	return *ret0, *ret1, err
}

// InterchainGet is a free data retrieval call binding the contract method 0x6079cf2a.
//
// Solidity: function interchainGet(string key) view returns(bool, string)
func (_DataSwapper *DataSwapperSession) InterchainGet(key string) (bool, string, error) {
	return _DataSwapper.Contract.InterchainGet(&_DataSwapper.CallOpts, key)
}

// InterchainGet is a free data retrieval call binding the contract method 0x6079cf2a.
//
// Solidity: function interchainGet(string key) view returns(bool, string)
func (_DataSwapper *DataSwapperCallerSession) InterchainGet(key string) (bool, string, error) {
	return _DataSwapper.Contract.InterchainGet(&_DataSwapper.CallOpts, key)
}

// InterchainSet is a paid mutator transaction binding the contract method 0x93c00ab0.
//
// Solidity: function interchainSet(string key, string value) returns()
func (_DataSwapper *DataSwapperTransactor) InterchainSet(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _DataSwapper.contract.Transact(opts, "interchainSet", key, value)
}

// InterchainSet is a paid mutator transaction binding the contract method 0x93c00ab0.
//
// Solidity: function interchainSet(string key, string value) returns()
func (_DataSwapper *DataSwapperSession) InterchainSet(key string, value string) (*types.Transaction, error) {
	return _DataSwapper.Contract.InterchainSet(&_DataSwapper.TransactOpts, key, value)
}

// InterchainSet is a paid mutator transaction binding the contract method 0x93c00ab0.
//
// Solidity: function interchainSet(string key, string value) returns()
func (_DataSwapper *DataSwapperTransactorSession) InterchainSet(key string, value string) (*types.Transaction, error) {
	return _DataSwapper.Contract.InterchainSet(&_DataSwapper.TransactOpts, key, value)
}

// TransferABI is the input ABI used to generate the binding from.
const TransferABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"val\",\"type\":\"uint64\"}],\"name\":\"interchainRollback\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"val\",\"type\":\"uint64\"}],\"name\":\"interchainCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransferFuncSigs maps the 4-byte function signature to its string representation.
var TransferFuncSigs = map[string]string{
	"e01b3517": "interchainCharge(string,string,uint64)",
	"4687ca34": "interchainRollback(string,uint64)",
}

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
	TransferFilterer   // Log filterer for contract events
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// NewTransferFilterer creates a new log filterer instance of Transfer, bound to a specific deployed contract.
func NewTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferFilterer, error) {
	contract, err := bindTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferFilterer{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xe01b3517.
//
// Solidity: function interchainCharge(string sender, string receiver, uint64 val) returns(bool)
func (_Transfer *TransferTransactor) InterchainCharge(opts *bind.TransactOpts, sender string, receiver string, val uint64) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "interchainCharge", sender, receiver, val)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xe01b3517.
//
// Solidity: function interchainCharge(string sender, string receiver, uint64 val) returns(bool)
func (_Transfer *TransferSession) InterchainCharge(sender string, receiver string, val uint64) (*types.Transaction, error) {
	return _Transfer.Contract.InterchainCharge(&_Transfer.TransactOpts, sender, receiver, val)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xe01b3517.
//
// Solidity: function interchainCharge(string sender, string receiver, uint64 val) returns(bool)
func (_Transfer *TransferTransactorSession) InterchainCharge(sender string, receiver string, val uint64) (*types.Transaction, error) {
	return _Transfer.Contract.InterchainCharge(&_Transfer.TransactOpts, sender, receiver, val)
}

// InterchainRollback is a paid mutator transaction binding the contract method 0x4687ca34.
//
// Solidity: function interchainRollback(string sender, uint64 val) returns(bool)
func (_Transfer *TransferTransactor) InterchainRollback(opts *bind.TransactOpts, sender string, val uint64) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "interchainRollback", sender, val)
}

// InterchainRollback is a paid mutator transaction binding the contract method 0x4687ca34.
//
// Solidity: function interchainRollback(string sender, uint64 val) returns(bool)
func (_Transfer *TransferSession) InterchainRollback(sender string, val uint64) (*types.Transaction, error) {
	return _Transfer.Contract.InterchainRollback(&_Transfer.TransactOpts, sender, val)
}

// InterchainRollback is a paid mutator transaction binding the contract method 0x4687ca34.
//
// Solidity: function interchainRollback(string sender, uint64 val) returns(bool)
func (_Transfer *TransferTransactorSession) InterchainRollback(sender string, val uint64) (*types.Transaction, error) {
	return _Transfer.Contract.InterchainRollback(&_Transfer.TransactOpts, sender, val)
}
