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
const BrokerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getOutMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"string\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getInMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCallbackMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInnerMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcChainMethod\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"req\",\"type\":\"bool\"},{\"name\":\"bizCallData\",\"type\":\"bytes\"}],\"name\":\"invokeInterchain\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcChainMethod\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"req\",\"type\":\"bool\"},{\"name\":\"err\",\"type\":\"string\"}],\"name\":\"invokeIndexUpdateWithError\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"int64\"}],\"name\":\"audit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destContractDID\",\"type\":\"string\"},{\"name\":\"funcs\",\"type\":\"string\"},{\"name\":\"args\",\"type\":\"string\"},{\"name\":\"argscb\",\"type\":\"string\"},{\"name\":\"argsrb\",\"type\":\"string\"}],\"name\":\"emitInterchainEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"string\"},{\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOuterMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"destDID\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"fid\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"funcs\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"args\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"argscb\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"argsrb\",\"type\":\"string\"}],\"name\":\"throwEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogInterchainData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LogInterchainStatus\",\"type\":\"event\"}]"

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
var BrokerBin = "0x608060405234801561001057600080fd5b5061213e806100206000396000f3fe6080604052600436106100a75760003560e01c80638129fc1c116100645780638129fc1c1461016f57806382e2384e14610184578063b38ff85f146101a4578063b508c8d4146101d1578063bed34bba146101f1578063c20cab5014610211576100a7565b806329793e6e146100ac57806331942306146100e25780633b6bbe4a146101025780634420e4861461012557806367b9fa3b146101475780637018c8c81461015c575b600080fd5b3480156100b857600080fd5b506100cc6100c7366004611adf565b610226565b6040516100d99190611f7a565b60405180910390f35b3480156100ee57600080fd5b506100cc6100fd366004611adf565b61026f565b34801561010e57600080fd5b5061011761028a565b6040516100d9929190611ef7565b34801561013157600080fd5b50610145610140366004611898565b610431565b005b34801561015357600080fd5b50610117610459565b61014561016a3660046118f0565b6105dd565b34801561017b57600080fd5b50610145610696565b34801561019057600080fd5b5061014561019f366004611b25565b610876565b3480156101b057600080fd5b506101c46101bf3660046118b6565b610888565b6040516100d99190611f1c565b3480156101dd57600080fd5b506101456101ec3660046119f8565b610958565b3480156101fd57600080fd5b506101c461020c36600461199b565b610b2d565b34801561021d57600080fd5b50610117610b86565b600061023183610d1e565b6007836040516102419190611ed4565b90815260408051602092819003830190206001600160401b0385166000908152925290205490505b92915050565b600061027a83610d1e565b6009836040516102419190611ed4565b60608060606005805490506040519080825280602002602001820160405280156102be578160200160208202803883390190505b50905060005b6005546001600160401b038216101561035257600a6005826001600160401b0316815481106102ef57fe5b906000526020600020016040516103069190611ee0565b9081526040519081900360200190205482516001600160401b0391821691849190841690811061033257fe5b6001600160401b03909216602092830291909101909101526001016102c4565b5060058181805480602002602001604051908101604052809291908181526020016000905b828210156104225760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561040e5780601f106103e35761010080835404028352916020019161040e565b820191906000526020600020905b8154815290600101906020018083116103f157829003601f168201915b505050505081526020019060010190610377565b50505050915092509250509091565b6001600160a01b03166000908152602081905260409020805467ffffffffffffffff19169055565b606080606060048054905060405190808252806020026020018201604052801561048d578160200160208202803883390190505b50905060005b60045481101561050d576008600482815481106104ac57fe5b906000526020600020016040516104c39190611ee0565b9081526040519081900360200190205482516001600160401b03909116908390839081106104ed57fe5b6001600160401b0390921660209283029190910190910152600101610493565b5060048181805480602002602001604051908101604052809291908181526020016000905b828210156104225760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156105c95780601f1061059e576101008083540402835291602001916105c9565b820191906000526020600020905b8154815290600101906020018083116105ac57829003601f168201915b505050505081526020019060010190610532565b6001600160a01b038416600090815260208190526040902054600790810b900b60011461060957600080fd5b61065987878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602081019091529081528a93508892509050610e1e565b604051602060848237805160208160040183378151808260240184376000808285348b5af13d806000863e818015610692578186a08186f35b8186fd5b60005b6004548110156107035760006008600483815481106106b457fe5b906000526020600020016040516106cb9190611ee0565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff19909216919091179055600101610699565b5060005b60035481101561077157600060066003838154811061072257fe5b906000526020600020016040516107399190611ee0565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff19909216919091179055600101610707565b5060005b6005548110156107df576000600a6005838154811061079057fe5b906000526020600020016040516107a79190611ee0565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff19909216919091179055600101610775565b5060005b60015481101561084f576000806000600184815481106107ff57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020805460079290920b6001600160401b031667ffffffffffffffff199092169190911790556001016107e3565b5061085c60036000611698565b61086860046000611698565b61087460056000611698565b565b61088284848484610e1e565b50505050565b60008160070b600019141580156108a357508160070b600014155b80156108b357508160070b600114155b156108c057506000610269565b6001600160a01b0383166000908152602081905260409020805467ffffffffffffffff19166001600160401b03600785900b908116919091179091556001141561094f576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0385161790555b50600192915050565b33600090815260208190526040902054600790810b900b60011461099a57604051600160e51b62461bcd02815260040161099190611f2a565b60405180910390fd5b60606109a586610ffb565b90506006816040516109b79190611ed4565b90815260405190819003602001812080546001600160401b038082166001011667ffffffffffffffff199091161790556006906109f5908390611ed4565b908152604051908190036020019020546001600160401b031660011415610a5d5760038054600181018083556000929092528251610a5a917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208501906116b9565b50505b43600782604051610a6e9190611ed4565b90815260200160405180910390206000600684604051610a8e9190611ed4565b9081526040805160209281900383019020546001600160401b0316835290820192909252810160002091909155517f85026d9efdca77567284833e9e5fa5ab0d412e8d0d9869942def88510971977c90600690610aec908490611ed4565b90815260405190819003602001812054610b1d916001600160401b0390911690899033908a908a908a908a90611f88565b60405180910390a1505050505050565b600081604051602001610b409190611ed4565b6040516020818303038152906040528051906020012083604051602001610b679190611ed4565b6040516020818303038152906040528051906020012014905092915050565b6060806060600380549050604051908082528060200260200182016040528015610bba578160200160208202803883390190505b50905060005b6003546001600160401b0382161015610c4e5760066003826001600160401b031681548110610beb57fe5b90600052602060002001604051610c029190611ee0565b9081526040519081900360200190205482516001600160401b03918216918491908416908110610c2e57fe5b6001600160401b0390921660209283029190910190910152600101610bc0565b5060038181805480602002602001604051908101604052809291908181526020016000905b828210156104225760008481526020908190208301805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610d0a5780601f10610cdf57610100808354040283529160200191610d0a565b820191906000526020600020905b815481529060010190602001808311610ced57829003601f168201915b505050505081526020019060010190610c73565b6060610d4682604051806040016040528060018152602001600160f91b601d028152506111ab565b90508051600414610d6c57604051600160e51b62461bcd02815260040161099190611f4a565b610da881600081518110610d7c57fe5b6020026020010151604051806040016040528060038152602001600160ea1b62191a5902815250610b2d565b610dc757604051600160e51b62461bcd02815260040161099190611f5a565b60015b6004811015610e19576060828281518110610de157fe5b60200260200101519050805160001415610e1057604051600160e51b62461bcd02815260040161099190611f6a565b50600101610dca565b505050565b8115610f0f57826001600160401b0316600885604051610e3e9190611ed4565b908152604051908190036020019020546001600160401b039081166001011614610e6757600080fd5b610e7084611311565b604051602001610e7f90611eec565b6040516020818303038152906040528051906020012081604051602001610ea69190611ed4565b6040516020818303038152906040528051906020012014610f0a5780600b85604051610ed29190611ed4565b90815260408051602092819003830190206001600160401b0387166000908152908352208251610f0893919291909101906116b9565b505b610882565b826001600160401b0316600a85604051610f299190611ed4565b908152604051908190036020019020546001600160401b039081166001011614610f5257600080fd5b610f5c8484611424565b604051602001610f6b90611eec565b6040516020818303038152906040528051906020012081604051602001610f929190611ed4565b60405160208183030381529060405280519060200120146108825780600c85604051610fbe9190611ed4565b90815260408051602092819003830190206001600160401b0387166000908152908352208251610ff493919291909101906116b9565b5050505050565b60608061102483604051806040016040528060018152602001600160f91b601d028152506111ab565b9050805160041461104a57604051600160e51b62461bcd02815260040161099190611f3a565b61105a81600081518110610d7c57fe5b61107957604051600160e51b62461bcd02815260040161099190611f3a565b60015b60048110156110cb57606082828151811061109357fe5b602002602001015190508051600014156110c257604051600160e51b62461bcd02815260040161099190611f6a565b5060010161107c565b5060606111186110ee836000815181106110e157fe5b6020026020010151611539565b611113604051806040016040528060018152602001600160f91b601d02815250611539565b61155e565b905061113661112682611539565b611113846001815181106110e157fe5b90506111446110ee82611539565b905061116261115282611539565b611113846002815181106110e157fe5b90506111706110ee82611539565b90506111a361117e82611539565b611113604051806040016040528060018152602001600160f91b601702815250611539565b949350505050565b606082600060015b60018351038210156111eb5760006111cc8787856115d2565b90508060001914156111de57506111eb565b60019081019250016111b3565b8060405190808252806020026020018201604052801561121f57816020015b606081526020019060019003908161120a5790505b50935060009150600090505b600183510382116113085760006112438787856115d2565b9050806000191415611253575082515b60608382036040519080825280601f01601f191660200182016040528015611282576020820181803883390190505b509050806000855b848110156112da5787818151811061129e57fe5b602001015160f81c60f81b8383806001019450815181106112bb57fe5b60200101906001600160f81b031916908160001a90535060010161128a565b50836001019550818886806001019750815181106112f457fe5b60200260200101819052505050505061122b565b50505092915050565b6008816040516113219190611ed4565b90815260405190819003602001812080546001600160401b038082166001011667ffffffffffffffff1990911617905560089061135f908390611ed4565b908152604051908190036020019020546001600160401b0316600114156113c757600480546001810180835560009290925282516113c4917f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b019060208501906116b9565b50505b436009826040516113d89190611ed4565b908152602001604051809103902060006008846040516113f89190611ed4565b9081526040805160209281900383019020546001600160401b0316835290820192909252016000205550565b600a826040516114349190611ed4565b908152604051908190036020019020546001600160401b03166114985760058054600181018083556000929092528351611495917f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0019060208601906116b9565b50505b80600a836040516114a99190611ed4565b908152602001604051809103902060006101000a8154816001600160401b0302191690836001600160401b03160217905550436009836040516114ec9190611ed4565b90815260200160405180910390206000600a8560405161150c9190611ed4565b9081526040805160209281900383019020546001600160401b031683529082019290925201600020555050565b611541611737565b506040805180820190915281518152602082810190820152919050565b60608082600001518460000151016040519080825280601f01601f191660200182016040528015611596576020820181803883390190505b50905060006020820190506115b4818660200151876000015161165a565b8451602085015185516115ca928401919061165a565b509392505050565b8151600090849084906001146115e457fe5b835b825181101561164a57816000815181106115fc57fe5b602001015160f81c60f81b6001600160f81b03191683828151811061161d57fe5b602001015160f81c60f81b6001600160f81b0319161415611642579250611653915050565b6001016115e6565b50600019925050505b9392505050565b5b6020811061167a578151835260209283019290910190601f190161165b565b905182516020929092036101000a6000190180199091169116179052565b50805460008255906000526020600020908101906116b69190611751565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116fa57805160ff1916838001178555611727565b82800160010185558215611727579182015b8281111561172757825182559160200191906001019061170c565b50611733929150611777565b5090565b604051806040016040528060008152602001600081525090565b61177491905b8082111561173357600061176b8282611791565b50600101611757565b90565b61177491905b80821115611733576000815560010161177d565b50805460018160011615610100020316600290046000825580601f106117b757506116b6565b601f0160209004906000526020600020908101906116b69190611777565b600061165382356120a0565b60006116538235612083565b60008083601f8401126117ff57600080fd5b5081356001600160401b0381111561181657600080fd5b60208301915083600182028301111561182e57600080fd5b9250929050565b600061165382356120ab565b600082601f83011261185257600080fd5b813561186561186082612038565b612012565b9150808252602083016020830185838301111561188157600080fd5b6113088382846120c2565b60006116538235612094565b6000602082840312156118aa57600080fd5b60006111a384846117d5565b600080604083850312156118c957600080fd5b60006118d585856117d5565b92505060206118e685828601611835565b9150509250929050565b600080600080600080600060a0888a03121561190b57600080fd5b87356001600160401b0381111561192157600080fd5b61192d8a828b016117ed565b975097505060206119408a828b0161188c565b95505060406119518a828b016117d5565b94505060606119628a828b016117e1565b93505060808801356001600160401b0381111561197e57600080fd5b61198a8a828b016117ed565b925092505092959891949750929550565b600080604083850312156119ae57600080fd5b82356001600160401b038111156119c457600080fd5b6119d085828601611841565b92505060208301356001600160401b038111156119ec57600080fd5b6118e685828601611841565b600080600080600060a08688031215611a1057600080fd5b85356001600160401b03811115611a2657600080fd5b611a3288828901611841565b95505060208601356001600160401b03811115611a4e57600080fd5b611a5a88828901611841565b94505060408601356001600160401b03811115611a7657600080fd5b611a8288828901611841565b93505060608601356001600160401b03811115611a9e57600080fd5b611aaa88828901611841565b92505060808601356001600160401b03811115611ac657600080fd5b611ad288828901611841565b9150509295509295909350565b60008060408385031215611af257600080fd5b82356001600160401b03811115611b0857600080fd5b611b1485828601611841565b92505060206118e68582860161188c565b60008060008060808587031215611b3b57600080fd5b84356001600160401b03811115611b5157600080fd5b611b5d87828801611841565b9450506020611b6e8782880161188c565b9350506040611b7f878288016117e1565b92505060608501356001600160401b03811115611b9b57600080fd5b611ba787828801611841565b91505092959194509250565b60006116538383611caa565b6000611bcb8383611ecb565b505060200190565b611bdc816120b1565b82525050565b6000611bed82612071565b611bf78185612075565b935083602082028501611c098561205f565b60005b84811015611c40578383038852611c24838351611bb3565b9250611c2f8261205f565b602098909801979150600101611c0c565b50909695505050505050565b6000611c5782612071565b611c618185612075565b9350611c6c8361205f565b60005b82811015611c9757611c82868351611bbf565b9550611c8d8261205f565b9150600101611c6f565b5093949350505050565b611bdc81612083565b6000611cb582612071565b611cbf8185612075565b9350611ccf8185602086016120ce565b611cd8816120fa565b9093019392505050565b6000611ced82612071565b611cf7818561207e565b9350611d078185602086016120ce565b9290920192915050565b600081546001811660008114611d2e5760018114611d5157611d90565b607f6002830416611d3f818761207e565b60ff1984168152955085019250611d90565b60028204611d5f818761207e565b9550611d6a85612065565b60005b82811015611d8957815488820152600190910190602001611d6d565b5050850192505b505092915050565b6000611da5601d83612075565b7f496e766f6b657220617265206e6f7420696e207768697465206c697374000000815260200192915050565b6000611dde601783612075565b7f646964206973206e6f74206c6567616c20666f726d6174000000000000000000815260200192915050565b6000611e17601783612075565b7f646964206973206e6f7420696e20666f75722070617274000000000000000000815260200192915050565b6000611e50601183612075565b7f707265666978206973206e6f7420646964000000000000000000000000000000815260200192915050565b6000611e89601383612075565b7f6469642073756273657420697320656d70747900000000000000000000000000815260200192915050565b600061026960008361207e565b611bdc81611774565b611bdc81612094565b60006116538284611ce2565b60006116538284611d11565b600061026982611eb5565b60408082528101611f088185611be2565b905081810360208301526111a38184611c4c565b602081016102698284611ca1565b6020808252810161026981611d98565b6020808252810161026981611dd1565b6020808252810161026981611e0a565b6020808252810161026981611e43565b6020808252810161026981611e7c565b602081016102698284611ec2565b60e08101611f96828a611ecb565b8181036020830152611fa88189611caa565b9050611fb76040830188611bd3565b8181036060830152611fc98187611caa565b90508181036080830152611fdd8186611caa565b905081810360a0830152611ff18185611caa565b905081810360c08301526120058184611caa565b9998505050505050505050565b6040518181016001600160401b038111828210171561203057600080fd5b604052919050565b60006001600160401b0382111561204e57600080fd5b506020601f91909101601f19160190565b60200190565b60009081526020902090565b5190565b90815260200190565b919050565b151590565b6001600160a01b031690565b6001600160401b031690565b600061026982612088565b60070b90565b6000610269826000610269826120a0565b82818337506000910152565b60005b838110156120e95781810151838201526020016120d1565b838111156108825750506000910152565b601f01601f19169056fea265627a7a72305820e5a523f99246934b4657eaa635739cfab06b29deca26dd804a59e7335a16b2866c6578706572696d656e74616cf50037"

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
