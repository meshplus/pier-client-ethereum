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

// AssetExchangeABI is the input ABI used to generate the binding from.
const AssetExchangeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"string\"},{\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeFinish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeConfirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcChainID\",\"type\":\"address\"},{\"name\":\"srcAddr\",\"type\":\"string\"},{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"senderOnSrcChain\",\"type\":\"string\"},{\"name\":\"receiverOnSrcChain\",\"type\":\"string\"},{\"name\":\"assetOnSrcChain\",\"type\":\"uint64\"},{\"name\":\"senderOnDstChain\",\"type\":\"string\"},{\"name\":\"receiverOnDstChain\",\"type\":\"string\"},{\"name\":\"assetOnDstChain\",\"type\":\"uint64\"}],\"name\":\"interchainAssetExchangeInit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AssetExchangeFuncSigs maps the 4-byte function signature to its string representation.
var AssetExchangeFuncSigs = map[string]string{
	"3a503948": "interchainAssetExchangeConfirm(string,string)",
	"25b90965": "interchainAssetExchangeFinish(string,string,string)",
	"413210a7": "interchainAssetExchangeInit(address,string,string,string,string,uint64,string,string,uint64)",
}

// AssetExchange is an auto generated Go binding around an Ethereum contract.
type AssetExchange struct {
	AssetExchangeCaller     // Read-only binding to the contract
	AssetExchangeTransactor // Write-only binding to the contract
	AssetExchangeFilterer   // Log filterer for contract events
}

// AssetExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetExchangeSession struct {
	Contract     *AssetExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetExchangeCallerSession struct {
	Contract *AssetExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AssetExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetExchangeTransactorSession struct {
	Contract     *AssetExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AssetExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetExchangeRaw struct {
	Contract *AssetExchange // Generic contract binding to access the raw methods on
}

// AssetExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetExchangeCallerRaw struct {
	Contract *AssetExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// AssetExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetExchangeTransactorRaw struct {
	Contract *AssetExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetExchange creates a new instance of AssetExchange, bound to a specific deployed contract.
func NewAssetExchange(address common.Address, backend bind.ContractBackend) (*AssetExchange, error) {
	contract, err := bindAssetExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetExchange{AssetExchangeCaller: AssetExchangeCaller{contract: contract}, AssetExchangeTransactor: AssetExchangeTransactor{contract: contract}, AssetExchangeFilterer: AssetExchangeFilterer{contract: contract}}, nil
}

// NewAssetExchangeCaller creates a new read-only instance of AssetExchange, bound to a specific deployed contract.
func NewAssetExchangeCaller(address common.Address, caller bind.ContractCaller) (*AssetExchangeCaller, error) {
	contract, err := bindAssetExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetExchangeCaller{contract: contract}, nil
}

// NewAssetExchangeTransactor creates a new write-only instance of AssetExchange, bound to a specific deployed contract.
func NewAssetExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetExchangeTransactor, error) {
	contract, err := bindAssetExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetExchangeTransactor{contract: contract}, nil
}

// NewAssetExchangeFilterer creates a new log filterer instance of AssetExchange, bound to a specific deployed contract.
func NewAssetExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetExchangeFilterer, error) {
	contract, err := bindAssetExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetExchangeFilterer{contract: contract}, nil
}

// bindAssetExchange binds a generic wrapper to an already deployed contract.
func bindAssetExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetExchange *AssetExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetExchange.Contract.AssetExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetExchange *AssetExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetExchange.Contract.AssetExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetExchange *AssetExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetExchange.Contract.AssetExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetExchange *AssetExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetExchange *AssetExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetExchange *AssetExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetExchange.Contract.contract.Transact(opts, method, params...)
}

// InterchainAssetExchangeConfirm is a paid mutator transaction binding the contract method 0x3a503948.
//
// Solidity: function interchainAssetExchangeConfirm(string assetExchangeId, string signatures) returns(bool)
func (_AssetExchange *AssetExchangeTransactor) InterchainAssetExchangeConfirm(opts *bind.TransactOpts, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _AssetExchange.contract.Transact(opts, "interchainAssetExchangeConfirm", assetExchangeId, signatures)
}

// InterchainAssetExchangeConfirm is a paid mutator transaction binding the contract method 0x3a503948.
//
// Solidity: function interchainAssetExchangeConfirm(string assetExchangeId, string signatures) returns(bool)
func (_AssetExchange *AssetExchangeSession) InterchainAssetExchangeConfirm(assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _AssetExchange.Contract.InterchainAssetExchangeConfirm(&_AssetExchange.TransactOpts, assetExchangeId, signatures)
}

// InterchainAssetExchangeConfirm is a paid mutator transaction binding the contract method 0x3a503948.
//
// Solidity: function interchainAssetExchangeConfirm(string assetExchangeId, string signatures) returns(bool)
func (_AssetExchange *AssetExchangeTransactorSession) InterchainAssetExchangeConfirm(assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _AssetExchange.Contract.InterchainAssetExchangeConfirm(&_AssetExchange.TransactOpts, assetExchangeId, signatures)
}

// InterchainAssetExchangeFinish is a paid mutator transaction binding the contract method 0x25b90965.
//
// Solidity: function interchainAssetExchangeFinish(string assetExchangeId, string status, string signatures) returns(bool)
func (_AssetExchange *AssetExchangeTransactor) InterchainAssetExchangeFinish(opts *bind.TransactOpts, assetExchangeId string, status string, signatures string) (*types.Transaction, error) {
	return _AssetExchange.contract.Transact(opts, "interchainAssetExchangeFinish", assetExchangeId, status, signatures)
}

// InterchainAssetExchangeFinish is a paid mutator transaction binding the contract method 0x25b90965.
//
// Solidity: function interchainAssetExchangeFinish(string assetExchangeId, string status, string signatures) returns(bool)
func (_AssetExchange *AssetExchangeSession) InterchainAssetExchangeFinish(assetExchangeId string, status string, signatures string) (*types.Transaction, error) {
	return _AssetExchange.Contract.InterchainAssetExchangeFinish(&_AssetExchange.TransactOpts, assetExchangeId, status, signatures)
}

// InterchainAssetExchangeFinish is a paid mutator transaction binding the contract method 0x25b90965.
//
// Solidity: function interchainAssetExchangeFinish(string assetExchangeId, string status, string signatures) returns(bool)
func (_AssetExchange *AssetExchangeTransactorSession) InterchainAssetExchangeFinish(assetExchangeId string, status string, signatures string) (*types.Transaction, error) {
	return _AssetExchange.Contract.InterchainAssetExchangeFinish(&_AssetExchange.TransactOpts, assetExchangeId, status, signatures)
}

// InterchainAssetExchangeInit is a paid mutator transaction binding the contract method 0x413210a7.
//
// Solidity: function interchainAssetExchangeInit(address srcChainID, string srcAddr, string assetExchangeId, string senderOnSrcChain, string receiverOnSrcChain, uint64 assetOnSrcChain, string senderOnDstChain, string receiverOnDstChain, uint64 assetOnDstChain) returns(bool)
func (_AssetExchange *AssetExchangeTransactor) InterchainAssetExchangeInit(opts *bind.TransactOpts, srcChainID common.Address, srcAddr string, assetExchangeId string, senderOnSrcChain string, receiverOnSrcChain string, assetOnSrcChain uint64, senderOnDstChain string, receiverOnDstChain string, assetOnDstChain uint64) (*types.Transaction, error) {
	return _AssetExchange.contract.Transact(opts, "interchainAssetExchangeInit", srcChainID, srcAddr, assetExchangeId, senderOnSrcChain, receiverOnSrcChain, assetOnSrcChain, senderOnDstChain, receiverOnDstChain, assetOnDstChain)
}

// InterchainAssetExchangeInit is a paid mutator transaction binding the contract method 0x413210a7.
//
// Solidity: function interchainAssetExchangeInit(address srcChainID, string srcAddr, string assetExchangeId, string senderOnSrcChain, string receiverOnSrcChain, uint64 assetOnSrcChain, string senderOnDstChain, string receiverOnDstChain, uint64 assetOnDstChain) returns(bool)
func (_AssetExchange *AssetExchangeSession) InterchainAssetExchangeInit(srcChainID common.Address, srcAddr string, assetExchangeId string, senderOnSrcChain string, receiverOnSrcChain string, assetOnSrcChain uint64, senderOnDstChain string, receiverOnDstChain string, assetOnDstChain uint64) (*types.Transaction, error) {
	return _AssetExchange.Contract.InterchainAssetExchangeInit(&_AssetExchange.TransactOpts, srcChainID, srcAddr, assetExchangeId, senderOnSrcChain, receiverOnSrcChain, assetOnSrcChain, senderOnDstChain, receiverOnDstChain, assetOnDstChain)
}

// InterchainAssetExchangeInit is a paid mutator transaction binding the contract method 0x413210a7.
//
// Solidity: function interchainAssetExchangeInit(address srcChainID, string srcAddr, string assetExchangeId, string senderOnSrcChain, string receiverOnSrcChain, uint64 assetOnSrcChain, string senderOnDstChain, string receiverOnDstChain, uint64 assetOnDstChain) returns(bool)
func (_AssetExchange *AssetExchangeTransactorSession) InterchainAssetExchangeInit(srcChainID common.Address, srcAddr string, assetExchangeId string, senderOnSrcChain string, receiverOnSrcChain string, assetOnSrcChain uint64, senderOnDstChain string, receiverOnDstChain string, assetOnDstChain uint64) (*types.Transaction, error) {
	return _AssetExchange.Contract.InterchainAssetExchangeInit(&_AssetExchange.TransactOpts, srcChainID, srcAddr, assetExchangeId, senderOnSrcChain, receiverOnSrcChain, assetOnSrcChain, senderOnDstChain, receiverOnDstChain, assetOnDstChain)
}

// BrokerABI is the input ABI used to generate the binding from.
const BrokerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"destChainID\",\"type\":\"address\"},{\"name\":\"destAddr\",\"type\":\"string\"},{\"name\":\"args\",\"type\":\"string\"}],\"name\":\"InterchainTransferInvoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"interchainSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"string\"}],\"name\":\"addAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCallbackMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"interchainGet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"string\"}],\"name\":\"removeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeConfirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destChainID\",\"type\":\"address\"},{\"name\":\"destAddr\",\"type\":\"string\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"InterchainDataSwapInvoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destChainID\",\"type\":\"address\"},{\"name\":\"destAddr\",\"type\":\"string\"},{\"name\":\"args\",\"type\":\"string\"},{\"name\":\"typ\",\"type\":\"uint64\"}],\"name\":\"InterchainAssetExchangeInvoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInnerMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint64\"},{\"name\":\"isRollback\",\"type\":\"bool\"}],\"name\":\"interchainCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getInMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSrcRollbackMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getOutMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"int64\"}],\"name\":\"audit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"},{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"interchainConfirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOuterMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeRefund\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"destAddr\",\"type\":\"address\"},{\"name\":\"srcAddr\",\"type\":\"string\"},{\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"name\":\"senderOnSrcChain\",\"type\":\"string\"},{\"name\":\"receiverOnSrcChain\",\"type\":\"string\"},{\"name\":\"assetOnSrcChain\",\"type\":\"uint64\"},{\"name\":\"senderOnDstChain\",\"type\":\"string\"},{\"name\":\"receiverOnDstChain\",\"type\":\"string\"},{\"name\":\"assetOnDstChain\",\"type\":\"uint64\"}],\"name\":\"interchainAssetExchangeInit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDstRollbackMeta\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fid\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tid\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"func\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"args\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"callback\",\"type\":\"string\"}],\"name\":\"throwEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogInterchainData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LogInterchainStatus\",\"type\":\"event\"}]"

// BrokerFuncSigs maps the 4-byte function signature to its string representation.
var BrokerFuncSigs = map[string]string{
	"66af2193": "InterchainAssetExchangeInvoke(address,string,string,uint64)",
	"5e7d7c4c": "InterchainDataSwapInvoke(address,string,string)",
	"06bef67c": "InterchainTransferInvoke(address,string,string)",
	"2ad26a04": "addAccount(string)",
	"b38ff85f": "audit(address,int64)",
	"3b6bbe4a": "getCallbackMeta()",
	"ed63513f": "getDstRollbackMeta()",
	"83c44c27": "getInMessage(address,uint64)",
	"67b9fa3b": "getInnerMeta()",
	"a0342a3f": "getOutMessage(address,uint64)",
	"c20cab50": "getOuterMeta()",
	"9edb89ea": "getSrcRollbackMeta()",
	"8129fc1c": "initialize()",
	"47c6ff2b": "interchainAssetExchangeConfirm(address,uint64,address,string,string)",
	"e8662659": "interchainAssetExchangeInit(address,uint64,address,string,string,string,string,uint64,string,string,uint64)",
	"c07c0a53": "interchainAssetExchangeRedeem(address,uint64,address,string,string)",
	"d89cecd0": "interchainAssetExchangeRefund(address,uint64,address,string,string)",
	"6e3c95d4": "interchainCharge(address,uint64,address,string,string,uint64,bool)",
	"be7c4222": "interchainConfirm(address,uint64,address,bool,string,uint64)",
	"3c25819a": "interchainGet(address,uint64,address,string)",
	"19ba2f2f": "interchainSet(address,uint64,address,string,string)",
	"4420e486": "register(address)",
	"46038ddc": "removeAccount(string)",
}

// BrokerBin is the compiled bytecode used for deploying new contracts.
var BrokerBin = "0x600560c09081527f416c69636500000000000000000000000000000000000000000000000000000060e052608090815261014060405260036101008181527f426f6200000000000000000000000000000000000000000000000000000000006101205260a052620000729160026200013e565b503480156200008057600080fd5b5060005b60035481101562000137576001600460038381548110620000a157fe5b9060005260206000200160405180828054600181600116156101000203166002900480156200010a5780601f10620000e75761010080835404028352918201916200010a565b820191906000526020600020905b815481529060010190602001808311620000f5575b50509283525050604051908190036020019020805491151560ff1990921691909117905560010162000084565b50620002b6565b82805482825590600052602060002090810192821562000190579160200282015b828111156200019057825180516200017f918491602090910190620001a2565b50916020019190600101906200015f565b506200019e92915062000223565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001e557805160ff191683800117855562000215565b8280016001018555821562000215579182015b8281111562000215578251825591602001919060010190620001f8565b506200019e9291506200024e565b6200024b91905b808211156200019e5760006200024182826200026b565b506001016200022a565b90565b6200024b91905b808211156200019e576000815560010162000255565b50805460018160011615610100020316600290046000825580601f10620002935750620002b3565b601f016020900490600052602060002090810190620002b391906200024e565b50565b613d5180620002c66000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80636e3c95d4116100c3578063be7c42221161007c578063be7c422214610d42578063c07c0a5314610e22578063c20cab5014610f73578063d89cecd014610f7b578063e8662659146110cc578063ed63513f1461144e5761014d565b80636e3c95d414610b265780638129fc1c14610c8757806383c44c2714610c8f5780639edb89ea14610cd6578063a0342a3f14610cde578063b38ff85f14610d135761014d565b80634420e486116101155780634420e4861461068657806346038ddc146106ac57806347c6ff2b146107505780635e7d7c4c146108a157806366af2193146109da57806367b9fa3b14610b1e5761014d565b806306bef67c1461015257806319ba2f2f1461030e5780632ad26a04146104735780633b6bbe4a146105195780633c25819a146105ba575b600080fd5b61028b6004803603606081101561016857600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561019257600080fd5b8201836020820111156101a457600080fd5b803590602001918460018302840111600160201b831117156101c557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561021757600080fd5b82018360208201111561022957600080fd5b803590602001918460018302840111600160201b8311171561024a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611456945050505050565b604051808315151515815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102d25781810151838201526020016102ba565b50505050905090810190601f1680156102ff5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61045f600480360360a081101561032457600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b81111561036657600080fd5b82018360208201111561037857600080fd5b803590602001918460018302840111600160201b8311171561039957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103eb57600080fd5b8201836020820111156103fd57600080fd5b803590602001918460018302840111600160201b8311171561041e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061165b945050505050565b604080519115158252519081900360200190f35b6105176004803603602081101561048957600080fd5b810190602081018135600160201b8111156104a357600080fd5b8201836020820111156104b557600080fd5b803590602001918460018302840111600160201b831117156104d657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061181e945050505050565b005b6105216118e3565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561056557818101518382015260200161054d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156105a457818101518382015260200161058c565b5050505090500194505050505060405180910390f35b61028b600480360360808110156105d057600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b81111561061257600080fd5b82018360208201111561062457600080fd5b803590602001918460018302840111600160201b8311171561064557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611a14945050505050565b6105176004803603602081101561069c57600080fd5b50356001600160a01b0316611c56565b610517600480360360208110156106c257600080fd5b810190602081018135600160201b8111156106dc57600080fd5b8201836020820111156106ee57600080fd5b803590602001918460018302840111600160201b8311171561070f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c7d945050505050565b61045f600480360360a081101561076657600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b8111156107a857600080fd5b8201836020820111156107ba57600080fd5b803590602001918460018302840111600160201b831117156107db57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561082d57600080fd5b82018360208201111561083f57600080fd5b803590602001918460018302840111600160201b8311171561086057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611d00945050505050565b61045f600480360360608110156108b757600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156108e157600080fd5b8201836020820111156108f357600080fd5b803590602001918460018302840111600160201b8311171561091457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561096657600080fd5b82018360208201111561097857600080fd5b803590602001918460018302840111600160201b8311171561099957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611f2c945050505050565b61045f600480360360808110156109f057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b811115610a1a57600080fd5b820183602082011115610a2c57600080fd5b803590602001918460018302840111600160201b83111715610a4d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610a9f57600080fd5b820183602082011115610ab157600080fd5b803590602001918460018302840111600160201b83111715610ad257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b0316915061201c9050565b610521612281565b61045f600480360360e0811015610b3c57600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b811115610b7e57600080fd5b820183602082011115610b9057600080fd5b803590602001918460018302840111600160201b83111715610bb157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610c0357600080fd5b820183602082011115610c1557600080fd5b803590602001918460018302840111600160201b83111715610c3657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550506001600160401b038335169350505060200135151561239c565b61051761267d565b610cc460048036036040811015610ca557600080fd5b5080356001600160a01b031690602001356001600160401b0316612934565b60408051918252519081900360200190f35b610521612968565b610cc460048036036040811015610cf457600080fd5b5080356001600160a01b031690602001356001600160401b0316612a83565b61045f60048036036040811015610d2957600080fd5b506001600160a01b03813516906020013560070b612ab6565b61045f600480360360c0811015610d5857600080fd5b6001600160a01b0382358116926001600160401b036020820135169260408201359092169160608201351515919081019060a081016080820135600160201b811115610da357600080fd5b820183602082011115610db557600080fd5b803590602001918460018302840111600160201b83111715610dd657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b03169150612b859050565b61045f600480360360a0811015610e3857600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b811115610e7a57600080fd5b820183602082011115610e8c57600080fd5b803590602001918460018302840111600160201b83111715610ead57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610eff57600080fd5b820183602082011115610f1157600080fd5b803590602001918460018302840111600160201b83111715610f3257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612dee945050505050565b610521612e1a565b61045f600480360360a0811015610f9157600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b811115610fd357600080fd5b820183602082011115610fe557600080fd5b803590602001918460018302840111600160201b8311171561100657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561105857600080fd5b82018360208201111561106a57600080fd5b803590602001918460018302840111600160201b8311171561108b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612ee2945050505050565b61045f60048036036101608110156110e357600080fd5b6001600160a01b0382358116926001600160401b0360208201351692604082013590921691810190608081016060820135600160201b81111561112557600080fd5b82018360208201111561113757600080fd5b803590602001918460018302840111600160201b8311171561115857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156111aa57600080fd5b8201836020820111156111bc57600080fd5b803590602001918460018302840111600160201b831117156111dd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561122f57600080fd5b82018360208201111561124157600080fd5b803590602001918460018302840111600160201b8311171561126257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156112b457600080fd5b8201836020820111156112c657600080fd5b803590602001918460018302840111600160201b831117156112e757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092956001600160401b03853516959094909350604081019250602001359050600160201b81111561134a57600080fd5b82018360208201111561135c57600080fd5b803590602001918460018302840111600160201b8311171561137d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156113cf57600080fd5b8201836020820111156113e157600080fd5b803590602001918460018302840111600160201b8311171561140257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b03169150612f0e9050565b61052161332c565b33600090815260208190526040812054606090600790810b900b6001146114c75760408051600160e51b62461bcd02815260206004820152601d60248201527f496e766f6b657220617265206e6f7420696e207768697465206c697374000000604482015290519081900360640190fd5b60606114ef84604051806040016040528060018152602001600160fa1b600b028152506133e0565b9050805160001415611520576000604051806060016040528060238152602001613cc8602391399250925050611653565b60048160008151811061152f57fe5b60200260200101516040518082805190602001908083835b602083106115665780518252601f199092019160209182019101611547565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506115c290505760006040518060600160405280603b8152602001613ceb603b91399250925050611653565b600061163c8733886040518060400160405280601081526020017f696e746572636861696e43686172676500000000000000000000000000000000815250896040518060400160405280601181526020017f696e746572636861696e436f6e6669726d000000000000000000000000000000815250613547565b604080516020810190915260008152909450925050505b935093915050565b6001600160a01b0385166000908152600c60205260408120546001600160401b039081166001018116908616146116b65760408051600081529051600080516020613ca88339815191529181900360200190a1506000611815565b836116c1878761381d565b60408051600160e41b63093c00ab028152600481019182528551604482015285516001600160a01b038416926393c00ab09288928892918291602482019160640190602087019080838360005b8381101561172657818101518382015260200161170e565b50505050905090810190601f1680156117535780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561178657818101518382015260200161176e565b50505050905090810190601f1680156117b35780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156117d457600080fd5b505af11580156117e8573d6000803e3d6000fd5b505060408051600181529051600080516020613ca88339815191529350908190036020019150a160019150505b95945050505050565b6000805b600254811015611862576002818154811061183957fe5b6000918252602090912001546001600160a01b031633141561185a57600191505b600101611822565b50801561186e57600080fd5b60016004836040518082805190602001908083835b602083106118a25780518252601f199092019160209182019101611883565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220805460ff19169315159390931790925550505050565b6060806060600780549050604051908082528060200260200182016040528015611917578160200160208202803883390190505b50905060005b6007546001600160401b03821610156119ab57600c60006007836001600160401b03168154811061194a57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b0391821691849190841690811061198b57fe5b6001600160401b039092166020928302919091019091015260010161191d565b5060078181805480602002602001604051908101604052809291908181526020018280548015611a0457602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116119e6575b5050505050915092509250509091565b6000606083611a22876138dc565b6001600160a01b038516600090815260208190526040902054600790810b900b600114611a62575050604080516020810190915260008082529150611c4d565b60606000826001600160a01b0316636079cf2a876040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ac2578181015183820152602001611aaa565b50505050905090810190601f168015611aef5780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b158015611b0c57600080fd5b505afa158015611b20573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015611b4957600080fd5b815160208301805191939283019291600160201b811115611b6957600080fd5b82016020810184811115611b7c57600080fd5b8151600160201b811182820187101715611b9557600080fd5b5050604080518615158152602080820183815284519383019390935283519399509697507f436160f7c24c5f31561ec9422a629accdbbd4e9e8ce21e86e634f497997769a8968896508995509093509091606084019185019080838360005b83811015611c0c578181015183820152602001611bf4565b50505050905090810190601f168015611c395780820380516001836020036101000a031916815260200191505b50935050505060405180910390a193509150505b94509492505050565b6001600160a01b0316600090815260208190526040902080546001600160401b0319169055565b6000805b600254811015611cc15760028181548110611c9857fe5b6000918252602090912001546001600160a01b0316331415611cb957600191505b600101611c81565b508015611ccd57600080fd5b6000600483604051808280519060200190808383602083106118a25780518252601f199092019160209182019101611883565b6001600160a01b0385166000908152600c60205260408120546001600160401b03908116600101811690861614611d5b5760408051600081529051600080516020613ca88339815191529181900360200190a1506000611815565b611d65868661381d565b6001600160a01b038416600090815260208190526040902054600790810b900b600114611db65760408051600081529051600080516020613ca88339815191529181900360200190a1506000611815565b60408051600160e31b63074a07290281526004810191825284516044820152845186926000926001600160a01b03851692633a50394892899289928291602481019160649091019060208701908083838d5b83811015611e20578181015183820152602001611e08565b50505050905090810190601f168015611e4d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611e80578181015183820152602001611e68565b50505050905090810190601f168015611ead5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015611ece57600080fd5b505af1158015611ee2573d6000803e3d6000fd5b505050506040513d6020811015611ef857600080fd5b50516040805182151581529051919250600080516020613ca8833981519152919081900360200190a1979650505050505050565b33600090815260208190526040812054600790810b900b600114611f9a5760408051600160e51b62461bcd02815260206004820152601d60248201527f496e766f6b657220617265206e6f7420696e207768697465206c697374000000604482015290519081900360640190fd5b6120128433856040518060400160405280600d81526020017f696e746572636861696e47657400000000000000000000000000000000000000815250866040518060400160405280600d81526020017f696e746572636861696e53657400000000000000000000000000000000000000815250613547565b90505b9392505050565b60006060806001600160401b03841661207e576040518060400160405280601b81526020017f696e746572636861696e417373657445786368616e6765496e697400000000008152509150604051806020016040528060008152509050612188565b836001600160401b031660011415612105576040518060400160405280601d81526020017f696e746572636861696e417373657445786368616e676552656465656d00000081525091506040518060400160405280601e81526020017f696e746572636861696e417373657445786368616e6765436f6e6669726d00008152509050612188565b836001600160401b031660021415612188576040518060400160405280601d81526020017f696e746572636861696e417373657445786368616e6765526566756e6400000081525091506040518060400160405280601e81526020017f696e746572636861696e417373657445786368616e6765436f6e6669726d000081525090505b6000612198883389868a87613547565b905060018115151415612273576001600160a01b0388166000908152600a60209081526040808320546008909252909120546001600160401b039182169116111561222a576001600160a01b038816600090815260086020908152604080832054600c90925290912080546001600160401b0319166001600160401b0392831660001901909216919091179055612273565b6001600160a01b0388166000908152600a6020908152604080832054600c90925290912080546001600160401b0319166001600160401b03928316600019019092169190911790555b506001979650505050505050565b60608060606006805490506040519080825280602002602001820160405280156122b5578160200160208202803883390190505b50905060005b60065481101561233557600a6000600683815481106122d657fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b039091169083908390811061231557fe5b6001600160401b03909216602092830291909101909101526001016122bb565b5060068181805480602002602001604051908101604052809291908181526020018280548015611a04576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116119e6575050505050915092509250509091565b600081612403576001600160a01b0388166000908152600a60205260409020546001600160401b039081166001018116908816146123fe5760408051600081529051600080516020613ca88339815191529181900360200190a1506000612672565b61248a565b6001600160a01b0388166000908152600e60205260409020546001600160401b038089169116106124585760408051600081529051600080516020613ca88339815191529181900360200190a1506000612672565b6001600160a01b0388166000908152600e6020526040902080546001600160401b0319166001600160401b0389161790555b612493886138dc565b6001600160a01b038616600090815260208190526040902054600790810b900b6001146124e45760408051600081529051600080516020613ca88339815191529181900360200190a1506000612672565b604051600160e01b63e5b2a5ed0281526001600160401b0384166044820152821515606482015260806004820190815286516084830152865188926000926001600160a01b0385169263e5b2a5ed928b928b928b928b92918291602481019160a49091019060208901908083838f5b8381101561256b578181015183820152602001612553565b50505050905090810190601f1680156125985780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b838110156125cb5781810151838201526020016125b3565b50505050905090810190601f1680156125f85780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561261b57600080fd5b505af115801561262f573d6000803e3d6000fd5b505050506040513d602081101561264557600080fd5b50516040805182151581529051919250600080516020613ca8833981519152919081900360200190a19150505b979650505050505050565b60005b6006548110156126e8576000600a60006006848154811061269d57fe5b6000918252602080832091909101546001600160a01b03168352820192909252604001902080546001600160401b0319166001600160401b0392909216919091179055600101612680565b5060005b600554811015612754576000600860006005848154811061270957fe5b6000918252602080832091909101546001600160a01b03168352820192909252604001902080546001600160401b0319166001600160401b03929092169190911790556001016126ec565b5060005b6007548110156127c0576000600c60006007848154811061277557fe5b6000918252602080832091909101546001600160a01b03168352820192909252604001902080546001600160401b0319166001600160401b0392909216919091179055600101612758565b5060005b60065481101561282c576000600e6000600684815481106127e157fe5b6000918252602080832091909101546001600160a01b03168352820192909252604001902080546001600160401b0319166001600160401b03929092169190911790556001016127c4565b5060005b600554811015612898576000600d60006005848154811061284d57fe5b6000918252602080832091909101546001600160a01b03168352820192909252604001902080546001600160401b0319166001600160401b0392909216919091179055600101612830565b5060005b600154811015612907576000806000600184815481106128b857fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020805460079290920b6001600160401b03166001600160401b031990921691909117905560010161289c565b506000612915600582613c5d565b506000612923600682613c5d565b506000612931600782613c5d565b50565b6001600160a01b0382166000908152600b602090815260408083206001600160401b03851684529091529020545b92915050565b606080606060058054905060405190808252806020026020018201604052801561299c578160200160208202803883390190505b50905060005b600554811015612a1c57600d6000600583815481106129bd57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b03909116908390839081106129fc57fe5b6001600160401b03909216602092830291909101909101526001016129a2565b5060058181805480602002602001604051908101604052809291908181526020018280548015611a04576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116119e6575050505050915092509250509091565b6001600160a01b03821660009081526009602090815260408083206001600160401b038516845290915290205492915050565b60008160070b60001914158015612ad157508160070b600014155b8015612ae157508160070b600114155b15612aee57506000612962565b6001600160a01b038316600090815260208190526040902080546001600160401b0319166001600160401b03600785900b9081169190911790915560011415612b7c576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0385161790555b50600192915050565b60008315612ca5576001600160a01b0387166000908152600c60205260409020546001600160401b03908116600101811690871614801590612bf157506001600160a01b0387166000908152600d60205260409020546001600160401b03908116600101811690871614155b15612c205760408051600081529051600080516020613ca88339815191529181900360200190a1506000612de4565b612c2a878761381d565b6001600160a01b038516600090815260208190526040902054600790810b900b600114612c7b5760408051600081529051600080516020613ca88339815191529181900360200190a1506000612de4565b60408051600181529051600080516020613ca88339815191529181900360200190a1506001612de4565b6001600160a01b038781166000908152600d6020908152604080832080546001600160401b0319166001600160401b038c8116919091179091558151600160e21b6311a1f28d0281529087166024820152600481019182528751604482015287518a95861693634687ca34938a938a9391928392606490920191908601908083838c5b83811015612d40578181015183820152602001612d28565b50505050905090810190601f168015612d6d5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015612d8d57600080fd5b505af1158015612da1573d6000803e3d6000fd5b505050506040513d6020811015612db757600080fd5b50516040805182151581529051919250600080516020613ca8833981519152919081900360200190a19150505b9695505050505050565b6000612de486868686604051806040016040528060018152602001600160f81b603102815250876139a5565b6060806060600580549050604051908082528060200260200182016040528015612e4e578160200160208202803883390190505b50905060005b6005546001600160401b0382161015612a1c57600860006005836001600160401b031681548110612e8157fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b03918216918491908416908110612ec257fe5b6001600160401b0390921660209283029190910190910152600101612e54565b6000612de486868686604051806040016040528060018152602001600160f91b601902815250876139a5565b6001600160a01b038b166000908152600a60205260408120546001600160401b039081166001018116908c1614612f695760408051600081529051600080516020613ca88339815191529181900360200190a150600061331d565b612f728c6138dc565b6001600160a01b038a16600090815260208190526040902054600790810b900b600114612fc35760408051600081529051600080516020613ca88339815191529181900360200190a150600061331d565b60008a90506000816001600160a01b031663413210a78f8d8d8d8d8d8d8d8d6040518a63ffffffff1660e01b8152600401808a6001600160a01b03166001600160a01b0316815260200180602001806020018060200180602001896001600160401b03166001600160401b031681526020018060200180602001886001600160401b03166001600160401b0316815260200187810387528f818151815260200191508051906020019080838360005b8381101561308a578181015183820152602001613072565b50505050905090810190601f1680156130b75780820380516001836020036101000a031916815260200191505b5087810386528e818151815260200191508051906020019080838360005b838110156130ed5781810151838201526020016130d5565b50505050905090810190601f16801561311a5780820380516001836020036101000a031916815260200191505b5087810385528d5181528d516020918201918f019080838360005b8381101561314d578181015183820152602001613135565b50505050905090810190601f16801561317a5780820380516001836020036101000a031916815260200191505b5087810384528c5181528c516020918201918e019080838360005b838110156131ad578181015183820152602001613195565b50505050905090810190601f1680156131da5780820380516001836020036101000a031916815260200191505b5087810383528a5181528a516020918201918c019080838360005b8381101561320d5781810151838201526020016131f5565b50505050905090810190601f16801561323a5780820380516001836020036101000a031916815260200191505b5087810382528951815289516020918201918b019080838360005b8381101561326d578181015183820152602001613255565b50505050905090810190601f16801561329a5780820380516001836020036101000a031916815260200191505b509f50505050505050505050505050505050602060405180830381600087803b1580156132c657600080fd5b505af11580156132da573d6000803e3d6000fd5b505050506040513d60208110156132f057600080fd5b50516040805182151581529051919250600080516020613ca8833981519152919081900360200190a19150505b9b9a5050505050505050505050565b6060806060600680549050604051908082528060200260200182016040528015613360578160200160208202803883390190505b50905060005b60065481101561233557600e60006006838154811061338157fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205482516001600160401b03909116908390839081106133c057fe5b6001600160401b0390921660209283029190910190910152600101613366565b606082600060015b6001835103821015613420576000613401878785613bd7565b90508060001914156134135750613420565b60019081019250016133e8565b8060405190808252806020026020018201604052801561345457816020015b606081526020019060019003908161343f5790505b50935060009150600090505b600183510382101561353e576000613479878785613bd7565b9050806000191415613489575082515b60608382036040519080825280601f01601f1916602001820160405280156134b8576020820181803883390190505b509050806000855b84811015613510578781815181106134d457fe5b602001015160f81c60f81b8383806001019450815181106134f157fe5b60200101906001600160f81b031916908160001a9053506001016134c0565b508360010195508188868060010197508151811061352a57fe5b602002602001018190525050505050613460565b50505092915050565b6001600160a01b038616600090815260086020526040812080546001600160401b0319811660016001600160401b03928316810183169190911792839055911614156135d957600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319166001600160a01b0389161790555b6001600160a01b038088166000818152600960209081526040808320600880845282852080546001600160401b039081168752928552838620439055868652908452548251911680825281840195909552948b169085015260e0606085018181528a519186019190915289517fad89cfa05a757be8d2179bb6609bf9034971b2427bd49d48e79552d3e8493e99958d948d948d948d948d948d949093608085019260a086019260c0870192610100880192908c01918190849084905b838110156136ad578181015183820152602001613695565b50505050905090810190601f1680156136da5780820380516001836020036101000a031916815260200191505b5085810384528851815288516020918201918a019080838360005b8381101561370d5781810151838201526020016136f5565b50505050905090810190601f16801561373a5780820380516001836020036101000a031916815260200191505b50858103835287518152875160209182019189019080838360005b8381101561376d578181015183820152602001613755565b50505050905090810190601f16801561379a5780820380516001836020036101000a031916815260200191505b50858103825286518152865160209182019188019080838360005b838110156137cd5781810151838201526020016137b5565b50505050905090810190601f1680156137fa5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390a15060019695505050505050565b6001600160a01b0382166000908152600c60205260409020546001600160401b031661388f57600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b0319166001600160a01b0384161790555b6001600160a01b03919091166000908152600c6020908152604080832080546001600160401b0319166001600160401b039586161790819055600b83528184209416835292905220439055565b6001600160a01b0381166000908152600a6020526040902080546001600160401b0319811660016001600160401b039283168101831691909117928390559116141561396e57600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319166001600160a01b0383161790555b6001600160a01b03166000908152600b60209081526040808320600a8352818420546001600160401b031684529091529020439055565b6001600160a01b0386166000908152600a60205260408120546001600160401b03908116600101811690871614613a005760408051600081529051600080516020613ca88339815191529181900360200190a1506000612de4565b613a09876138dc565b6001600160a01b038516600090815260208190526040902054600790810b900b600114613a5a5760408051600081529051600080516020613ca88339815191529181900360200190a1506000612de4565b60008590506000816001600160a01b03166325b909658787876040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015613ac7578181015183820152602001613aaf565b50505050905090810190601f168015613af45780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b83811015613b27578181015183820152602001613b0f565b50505050905090810190601f168015613b545780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015613b87578181015183820152602001613b6f565b50505050905090810190601f168015613bb45780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b158015612d8d57600080fd5b815160009084908490600114613be957fe5b835b8251811015613c4f5781600081518110613c0157fe5b602001015160f81c60f81b6001600160f81b031916838281518110613c2257fe5b602001015160f81c60f81b6001600160f81b0319161415613c47579250612015915050565b600101613beb565b506000199695505050505050565b815481835581811115613c8157600083815260209020613c81918101908301613c86565b505050565b613ca491905b80821115613ca05760008155600101613c8c565b5090565b9056fe23de11857b4338b8e6ccaec81162b447b44040ff3cfdd1174d548975eb5c1c3e6172677320666f7220696e766f6b696e67207472616e7366657220697320656d70747973656e646572206163636f756e74206973206e6f7420616c6c6f77656420746f20696e766f6b6520696e746572636861696e207472616e73666572a165627a7a723058202708b9dfb46a57fd1cd71645599897d3cdb7cde448908b05ce08668a7ae2f49d0029"

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

// GetDstRollbackMeta is a free data retrieval call binding the contract method 0xed63513f.
//
// Solidity: function getDstRollbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerCaller) GetDstRollbackMeta(opts *bind.CallOpts) ([]common.Address, []uint64, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Broker.contract.Call(opts, out, "getDstRollbackMeta")
	return *ret0, *ret1, err
}

// GetDstRollbackMeta is a free data retrieval call binding the contract method 0xed63513f.
//
// Solidity: function getDstRollbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerSession) GetDstRollbackMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetDstRollbackMeta(&_Broker.CallOpts)
}

// GetDstRollbackMeta is a free data retrieval call binding the contract method 0xed63513f.
//
// Solidity: function getDstRollbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerCallerSession) GetDstRollbackMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetDstRollbackMeta(&_Broker.CallOpts)
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

// GetSrcRollbackMeta is a free data retrieval call binding the contract method 0x9edb89ea.
//
// Solidity: function getSrcRollbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerCaller) GetSrcRollbackMeta(opts *bind.CallOpts) ([]common.Address, []uint64, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Broker.contract.Call(opts, out, "getSrcRollbackMeta")
	return *ret0, *ret1, err
}

// GetSrcRollbackMeta is a free data retrieval call binding the contract method 0x9edb89ea.
//
// Solidity: function getSrcRollbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerSession) GetSrcRollbackMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetSrcRollbackMeta(&_Broker.CallOpts)
}

// GetSrcRollbackMeta is a free data retrieval call binding the contract method 0x9edb89ea.
//
// Solidity: function getSrcRollbackMeta() view returns(address[], uint64[])
func (_Broker *BrokerCallerSession) GetSrcRollbackMeta() ([]common.Address, []uint64, error) {
	return _Broker.Contract.GetSrcRollbackMeta(&_Broker.CallOpts)
}

// InterchainAssetExchangeInvoke is a paid mutator transaction binding the contract method 0x66af2193.
//
// Solidity: function InterchainAssetExchangeInvoke(address destChainID, string destAddr, string args, uint64 typ) returns(bool)
func (_Broker *BrokerTransactor) InterchainAssetExchangeInvoke(opts *bind.TransactOpts, destChainID common.Address, destAddr string, args string, typ uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "InterchainAssetExchangeInvoke", destChainID, destAddr, args, typ)
}

// InterchainAssetExchangeInvoke is a paid mutator transaction binding the contract method 0x66af2193.
//
// Solidity: function InterchainAssetExchangeInvoke(address destChainID, string destAddr, string args, uint64 typ) returns(bool)
func (_Broker *BrokerSession) InterchainAssetExchangeInvoke(destChainID common.Address, destAddr string, args string, typ uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeInvoke(&_Broker.TransactOpts, destChainID, destAddr, args, typ)
}

// InterchainAssetExchangeInvoke is a paid mutator transaction binding the contract method 0x66af2193.
//
// Solidity: function InterchainAssetExchangeInvoke(address destChainID, string destAddr, string args, uint64 typ) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainAssetExchangeInvoke(destChainID common.Address, destAddr string, args string, typ uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeInvoke(&_Broker.TransactOpts, destChainID, destAddr, args, typ)
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
// Solidity: function InterchainTransferInvoke(address destChainID, string destAddr, string args) returns(bool, string)
func (_Broker *BrokerTransactor) InterchainTransferInvoke(opts *bind.TransactOpts, destChainID common.Address, destAddr string, args string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "InterchainTransferInvoke", destChainID, destAddr, args)
}

// InterchainTransferInvoke is a paid mutator transaction binding the contract method 0x06bef67c.
//
// Solidity: function InterchainTransferInvoke(address destChainID, string destAddr, string args) returns(bool, string)
func (_Broker *BrokerSession) InterchainTransferInvoke(destChainID common.Address, destAddr string, args string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainTransferInvoke(&_Broker.TransactOpts, destChainID, destAddr, args)
}

// InterchainTransferInvoke is a paid mutator transaction binding the contract method 0x06bef67c.
//
// Solidity: function InterchainTransferInvoke(address destChainID, string destAddr, string args) returns(bool, string)
func (_Broker *BrokerTransactorSession) InterchainTransferInvoke(destChainID common.Address, destAddr string, args string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainTransferInvoke(&_Broker.TransactOpts, destChainID, destAddr, args)
}

// AddAccount is a paid mutator transaction binding the contract method 0x2ad26a04.
//
// Solidity: function addAccount(string account) returns()
func (_Broker *BrokerTransactor) AddAccount(opts *bind.TransactOpts, account string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "addAccount", account)
}

// AddAccount is a paid mutator transaction binding the contract method 0x2ad26a04.
//
// Solidity: function addAccount(string account) returns()
func (_Broker *BrokerSession) AddAccount(account string) (*types.Transaction, error) {
	return _Broker.Contract.AddAccount(&_Broker.TransactOpts, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0x2ad26a04.
//
// Solidity: function addAccount(string account) returns()
func (_Broker *BrokerTransactorSession) AddAccount(account string) (*types.Transaction, error) {
	return _Broker.Contract.AddAccount(&_Broker.TransactOpts, account)
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

// InterchainAssetExchangeConfirm is a paid mutator transaction binding the contract method 0x47c6ff2b.
//
// Solidity: function interchainAssetExchangeConfirm(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerTransactor) InterchainAssetExchangeConfirm(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainAssetExchangeConfirm", sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeConfirm is a paid mutator transaction binding the contract method 0x47c6ff2b.
//
// Solidity: function interchainAssetExchangeConfirm(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerSession) InterchainAssetExchangeConfirm(sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeConfirm(&_Broker.TransactOpts, sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeConfirm is a paid mutator transaction binding the contract method 0x47c6ff2b.
//
// Solidity: function interchainAssetExchangeConfirm(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainAssetExchangeConfirm(sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeConfirm(&_Broker.TransactOpts, sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeInit is a paid mutator transaction binding the contract method 0xe8662659.
//
// Solidity: function interchainAssetExchangeInit(address sourceChainID, uint64 index, address destAddr, string srcAddr, string assetExchangeId, string senderOnSrcChain, string receiverOnSrcChain, uint64 assetOnSrcChain, string senderOnDstChain, string receiverOnDstChain, uint64 assetOnDstChain) returns(bool)
func (_Broker *BrokerTransactor) InterchainAssetExchangeInit(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, srcAddr string, assetExchangeId string, senderOnSrcChain string, receiverOnSrcChain string, assetOnSrcChain uint64, senderOnDstChain string, receiverOnDstChain string, assetOnDstChain uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainAssetExchangeInit", sourceChainID, index, destAddr, srcAddr, assetExchangeId, senderOnSrcChain, receiverOnSrcChain, assetOnSrcChain, senderOnDstChain, receiverOnDstChain, assetOnDstChain)
}

// InterchainAssetExchangeInit is a paid mutator transaction binding the contract method 0xe8662659.
//
// Solidity: function interchainAssetExchangeInit(address sourceChainID, uint64 index, address destAddr, string srcAddr, string assetExchangeId, string senderOnSrcChain, string receiverOnSrcChain, uint64 assetOnSrcChain, string senderOnDstChain, string receiverOnDstChain, uint64 assetOnDstChain) returns(bool)
func (_Broker *BrokerSession) InterchainAssetExchangeInit(sourceChainID common.Address, index uint64, destAddr common.Address, srcAddr string, assetExchangeId string, senderOnSrcChain string, receiverOnSrcChain string, assetOnSrcChain uint64, senderOnDstChain string, receiverOnDstChain string, assetOnDstChain uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeInit(&_Broker.TransactOpts, sourceChainID, index, destAddr, srcAddr, assetExchangeId, senderOnSrcChain, receiverOnSrcChain, assetOnSrcChain, senderOnDstChain, receiverOnDstChain, assetOnDstChain)
}

// InterchainAssetExchangeInit is a paid mutator transaction binding the contract method 0xe8662659.
//
// Solidity: function interchainAssetExchangeInit(address sourceChainID, uint64 index, address destAddr, string srcAddr, string assetExchangeId, string senderOnSrcChain, string receiverOnSrcChain, uint64 assetOnSrcChain, string senderOnDstChain, string receiverOnDstChain, uint64 assetOnDstChain) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainAssetExchangeInit(sourceChainID common.Address, index uint64, destAddr common.Address, srcAddr string, assetExchangeId string, senderOnSrcChain string, receiverOnSrcChain string, assetOnSrcChain uint64, senderOnDstChain string, receiverOnDstChain string, assetOnDstChain uint64) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeInit(&_Broker.TransactOpts, sourceChainID, index, destAddr, srcAddr, assetExchangeId, senderOnSrcChain, receiverOnSrcChain, assetOnSrcChain, senderOnDstChain, receiverOnDstChain, assetOnDstChain)
}

// InterchainAssetExchangeRedeem is a paid mutator transaction binding the contract method 0xc07c0a53.
//
// Solidity: function interchainAssetExchangeRedeem(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerTransactor) InterchainAssetExchangeRedeem(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainAssetExchangeRedeem", sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeRedeem is a paid mutator transaction binding the contract method 0xc07c0a53.
//
// Solidity: function interchainAssetExchangeRedeem(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerSession) InterchainAssetExchangeRedeem(sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeRedeem(&_Broker.TransactOpts, sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeRedeem is a paid mutator transaction binding the contract method 0xc07c0a53.
//
// Solidity: function interchainAssetExchangeRedeem(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainAssetExchangeRedeem(sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeRedeem(&_Broker.TransactOpts, sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeRefund is a paid mutator transaction binding the contract method 0xd89cecd0.
//
// Solidity: function interchainAssetExchangeRefund(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerTransactor) InterchainAssetExchangeRefund(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainAssetExchangeRefund", sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeRefund is a paid mutator transaction binding the contract method 0xd89cecd0.
//
// Solidity: function interchainAssetExchangeRefund(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerSession) InterchainAssetExchangeRefund(sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeRefund(&_Broker.TransactOpts, sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainAssetExchangeRefund is a paid mutator transaction binding the contract method 0xd89cecd0.
//
// Solidity: function interchainAssetExchangeRefund(address sourceChainID, uint64 index, address destAddr, string assetExchangeId, string signatures) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainAssetExchangeRefund(sourceChainID common.Address, index uint64, destAddr common.Address, assetExchangeId string, signatures string) (*types.Transaction, error) {
	return _Broker.Contract.InterchainAssetExchangeRefund(&_Broker.TransactOpts, sourceChainID, index, destAddr, assetExchangeId, signatures)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0x6e3c95d4.
//
// Solidity: function interchainCharge(address sourceChainID, uint64 index, address destAddr, string sender, string receiver, uint64 amount, bool isRollback) returns(bool)
func (_Broker *BrokerTransactor) InterchainCharge(opts *bind.TransactOpts, sourceChainID common.Address, index uint64, destAddr common.Address, sender string, receiver string, amount uint64, isRollback bool) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "interchainCharge", sourceChainID, index, destAddr, sender, receiver, amount, isRollback)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0x6e3c95d4.
//
// Solidity: function interchainCharge(address sourceChainID, uint64 index, address destAddr, string sender, string receiver, uint64 amount, bool isRollback) returns(bool)
func (_Broker *BrokerSession) InterchainCharge(sourceChainID common.Address, index uint64, destAddr common.Address, sender string, receiver string, amount uint64, isRollback bool) (*types.Transaction, error) {
	return _Broker.Contract.InterchainCharge(&_Broker.TransactOpts, sourceChainID, index, destAddr, sender, receiver, amount, isRollback)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0x6e3c95d4.
//
// Solidity: function interchainCharge(address sourceChainID, uint64 index, address destAddr, string sender, string receiver, uint64 amount, bool isRollback) returns(bool)
func (_Broker *BrokerTransactorSession) InterchainCharge(sourceChainID common.Address, index uint64, destAddr common.Address, sender string, receiver string, amount uint64, isRollback bool) (*types.Transaction, error) {
	return _Broker.Contract.InterchainCharge(&_Broker.TransactOpts, sourceChainID, index, destAddr, sender, receiver, amount, isRollback)
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

// RemoveAccount is a paid mutator transaction binding the contract method 0x46038ddc.
//
// Solidity: function removeAccount(string account) returns()
func (_Broker *BrokerTransactor) RemoveAccount(opts *bind.TransactOpts, account string) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "removeAccount", account)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0x46038ddc.
//
// Solidity: function removeAccount(string account) returns()
func (_Broker *BrokerSession) RemoveAccount(account string) (*types.Transaction, error) {
	return _Broker.Contract.RemoveAccount(&_Broker.TransactOpts, account)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0x46038ddc.
//
// Solidity: function removeAccount(string account) returns()
func (_Broker *BrokerTransactorSession) RemoveAccount(account string) (*types.Transaction, error) {
	return _Broker.Contract.RemoveAccount(&_Broker.TransactOpts, account)
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
const TransferABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"val\",\"type\":\"uint64\"}],\"name\":\"interchainRollback\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"string\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"val\",\"type\":\"uint64\"},{\"name\":\"isRollback\",\"type\":\"bool\"}],\"name\":\"interchainCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransferFuncSigs maps the 4-byte function signature to its string representation.
var TransferFuncSigs = map[string]string{
	"e5b2a5ed": "interchainCharge(string,string,uint64,bool)",
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

// InterchainCharge is a paid mutator transaction binding the contract method 0xe5b2a5ed.
//
// Solidity: function interchainCharge(string sender, string receiver, uint64 val, bool isRollback) returns(bool)
func (_Transfer *TransferTransactor) InterchainCharge(opts *bind.TransactOpts, sender string, receiver string, val uint64, isRollback bool) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "interchainCharge", sender, receiver, val, isRollback)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xe5b2a5ed.
//
// Solidity: function interchainCharge(string sender, string receiver, uint64 val, bool isRollback) returns(bool)
func (_Transfer *TransferSession) InterchainCharge(sender string, receiver string, val uint64, isRollback bool) (*types.Transaction, error) {
	return _Transfer.Contract.InterchainCharge(&_Transfer.TransactOpts, sender, receiver, val, isRollback)
}

// InterchainCharge is a paid mutator transaction binding the contract method 0xe5b2a5ed.
//
// Solidity: function interchainCharge(string sender, string receiver, uint64 val, bool isRollback) returns(bool)
func (_Transfer *TransferTransactorSession) InterchainCharge(sender string, receiver string, val uint64, isRollback bool) (*types.Transaction, error) {
	return _Transfer.Contract.InterchainCharge(&_Transfer.TransactOpts, sender, receiver, val, isRollback)
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
