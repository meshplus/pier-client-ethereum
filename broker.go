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
const BrokerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogInterchainData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LogInterchainStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fid\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"func\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callback\",\"type\":\"string\"}],\"name\":\"throwEvent\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destChainID\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"typ\",\"type\":\"uint64\"}],\"name\":\"InterchainAssetExchangeInvoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destChainID\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"InterchainDataSwapInvoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destChainID\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"args\",\"type\":\"string\"}],\"name\":\"InterchainTransferInvoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"status\",\"type\":\"int64\"}],\"name\":\"audit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCallbackMeta\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getInMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInnerMeta\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"getOutMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOuterMeta\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeConfirm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"srcAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"senderOnSrcChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiverOnSrcChain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"assetOnSrcChain\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"senderOnDstChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiverOnDstChain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"assetOnDstChain\",\"type\":\"uint64\"}],\"name\":\"interchainAssetExchangeInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetExchangeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatures\",\"type\":\"string\"}],\"name\":\"interchainAssetExchangeRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sender\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"interchainCharge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"sender\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"interchainConfirm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"interchainGet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceChainID\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"interchainSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
