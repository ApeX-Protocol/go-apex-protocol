// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AmmFactoryMetaData contains all meta data concerning the AmmFactory contract.
var AmmFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"upperFactory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"config_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeToSetter_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"name\":\"AmmCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"createAmm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getAmm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"margin\",\"type\":\"address\"}],\"name\":\"initAmm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeTo_\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToSetter_\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AmmFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AmmFactoryMetaData.ABI instead.
var AmmFactoryABI = AmmFactoryMetaData.ABI

// AmmFactory is an auto generated Go binding around an Ethereum contract.
type AmmFactory struct {
	AmmFactoryCaller     // Read-only binding to the contract
	AmmFactoryTransactor // Write-only binding to the contract
	AmmFactoryFilterer   // Log filterer for contract events
}

// AmmFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmmFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmmFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmmFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmmFactorySession struct {
	Contract     *AmmFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmmFactoryCallerSession struct {
	Contract *AmmFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AmmFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmmFactoryTransactorSession struct {
	Contract     *AmmFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AmmFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmmFactoryRaw struct {
	Contract *AmmFactory // Generic contract binding to access the raw methods on
}

// AmmFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmmFactoryCallerRaw struct {
	Contract *AmmFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AmmFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmmFactoryTransactorRaw struct {
	Contract *AmmFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmmFactory creates a new instance of AmmFactory, bound to a specific deployed contract.
func NewAmmFactory(address common.Address, backend bind.ContractBackend) (*AmmFactory, error) {
	contract, err := bindAmmFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AmmFactory{AmmFactoryCaller: AmmFactoryCaller{contract: contract}, AmmFactoryTransactor: AmmFactoryTransactor{contract: contract}, AmmFactoryFilterer: AmmFactoryFilterer{contract: contract}}, nil
}

// NewAmmFactoryCaller creates a new read-only instance of AmmFactory, bound to a specific deployed contract.
func NewAmmFactoryCaller(address common.Address, caller bind.ContractCaller) (*AmmFactoryCaller, error) {
	contract, err := bindAmmFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmmFactoryCaller{contract: contract}, nil
}

// NewAmmFactoryTransactor creates a new write-only instance of AmmFactory, bound to a specific deployed contract.
func NewAmmFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AmmFactoryTransactor, error) {
	contract, err := bindAmmFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmmFactoryTransactor{contract: contract}, nil
}

// NewAmmFactoryFilterer creates a new log filterer instance of AmmFactory, bound to a specific deployed contract.
func NewAmmFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AmmFactoryFilterer, error) {
	contract, err := bindAmmFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmmFactoryFilterer{contract: contract}, nil
}

// bindAmmFactory binds a generic wrapper to an already deployed contract.
func bindAmmFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmmFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmmFactory *AmmFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmmFactory.Contract.AmmFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmmFactory *AmmFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmFactory.Contract.AmmFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmmFactory *AmmFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmmFactory.Contract.AmmFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmmFactory *AmmFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmmFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmmFactory *AmmFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmmFactory *AmmFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmmFactory.Contract.contract.Transact(opts, method, params...)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_AmmFactory *AmmFactoryCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmFactory.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_AmmFactory *AmmFactorySession) Config() (common.Address, error) {
	return _AmmFactory.Contract.Config(&_AmmFactory.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_AmmFactory *AmmFactoryCallerSession) Config() (common.Address, error) {
	return _AmmFactory.Contract.Config(&_AmmFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_AmmFactory *AmmFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_AmmFactory *AmmFactorySession) FeeTo() (common.Address, error) {
	return _AmmFactory.Contract.FeeTo(&_AmmFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_AmmFactory *AmmFactoryCallerSession) FeeTo() (common.Address, error) {
	return _AmmFactory.Contract.FeeTo(&_AmmFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_AmmFactory *AmmFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_AmmFactory *AmmFactorySession) FeeToSetter() (common.Address, error) {
	return _AmmFactory.Contract.FeeToSetter(&_AmmFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_AmmFactory *AmmFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _AmmFactory.Contract.FeeToSetter(&_AmmFactory.CallOpts)
}

// GetAmm is a free data retrieval call binding the contract method 0x65cc7a1d.
//
// Solidity: function getAmm(address , address ) view returns(address)
func (_AmmFactory *AmmFactoryCaller) GetAmm(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AmmFactory.contract.Call(opts, &out, "getAmm", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAmm is a free data retrieval call binding the contract method 0x65cc7a1d.
//
// Solidity: function getAmm(address , address ) view returns(address)
func (_AmmFactory *AmmFactorySession) GetAmm(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _AmmFactory.Contract.GetAmm(&_AmmFactory.CallOpts, arg0, arg1)
}

// GetAmm is a free data retrieval call binding the contract method 0x65cc7a1d.
//
// Solidity: function getAmm(address , address ) view returns(address)
func (_AmmFactory *AmmFactoryCallerSession) GetAmm(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _AmmFactory.Contract.GetAmm(&_AmmFactory.CallOpts, arg0, arg1)
}

// UpperFactory is a free data retrieval call binding the contract method 0xcedc12d8.
//
// Solidity: function upperFactory() view returns(address)
func (_AmmFactory *AmmFactoryCaller) UpperFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmFactory.contract.Call(opts, &out, "upperFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpperFactory is a free data retrieval call binding the contract method 0xcedc12d8.
//
// Solidity: function upperFactory() view returns(address)
func (_AmmFactory *AmmFactorySession) UpperFactory() (common.Address, error) {
	return _AmmFactory.Contract.UpperFactory(&_AmmFactory.CallOpts)
}

// UpperFactory is a free data retrieval call binding the contract method 0xcedc12d8.
//
// Solidity: function upperFactory() view returns(address)
func (_AmmFactory *AmmFactoryCallerSession) UpperFactory() (common.Address, error) {
	return _AmmFactory.Contract.UpperFactory(&_AmmFactory.CallOpts)
}

// CreateAmm is a paid mutator transaction binding the contract method 0x5131c85b.
//
// Solidity: function createAmm(address baseToken, address quoteToken) returns(address amm)
func (_AmmFactory *AmmFactoryTransactor) CreateAmm(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _AmmFactory.contract.Transact(opts, "createAmm", baseToken, quoteToken)
}

// CreateAmm is a paid mutator transaction binding the contract method 0x5131c85b.
//
// Solidity: function createAmm(address baseToken, address quoteToken) returns(address amm)
func (_AmmFactory *AmmFactorySession) CreateAmm(baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.CreateAmm(&_AmmFactory.TransactOpts, baseToken, quoteToken)
}

// CreateAmm is a paid mutator transaction binding the contract method 0x5131c85b.
//
// Solidity: function createAmm(address baseToken, address quoteToken) returns(address amm)
func (_AmmFactory *AmmFactoryTransactorSession) CreateAmm(baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.CreateAmm(&_AmmFactory.TransactOpts, baseToken, quoteToken)
}

// InitAmm is a paid mutator transaction binding the contract method 0x35e30d41.
//
// Solidity: function initAmm(address baseToken, address quoteToken, address margin) returns()
func (_AmmFactory *AmmFactoryTransactor) InitAmm(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, margin common.Address) (*types.Transaction, error) {
	return _AmmFactory.contract.Transact(opts, "initAmm", baseToken, quoteToken, margin)
}

// InitAmm is a paid mutator transaction binding the contract method 0x35e30d41.
//
// Solidity: function initAmm(address baseToken, address quoteToken, address margin) returns()
func (_AmmFactory *AmmFactorySession) InitAmm(baseToken common.Address, quoteToken common.Address, margin common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.InitAmm(&_AmmFactory.TransactOpts, baseToken, quoteToken, margin)
}

// InitAmm is a paid mutator transaction binding the contract method 0x35e30d41.
//
// Solidity: function initAmm(address baseToken, address quoteToken, address margin) returns()
func (_AmmFactory *AmmFactoryTransactorSession) InitAmm(baseToken common.Address, quoteToken common.Address, margin common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.InitAmm(&_AmmFactory.TransactOpts, baseToken, quoteToken, margin)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_AmmFactory *AmmFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, feeTo_ common.Address) (*types.Transaction, error) {
	return _AmmFactory.contract.Transact(opts, "setFeeTo", feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_AmmFactory *AmmFactorySession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.SetFeeTo(&_AmmFactory.TransactOpts, feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_AmmFactory *AmmFactoryTransactorSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.SetFeeTo(&_AmmFactory.TransactOpts, feeTo_)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address feeToSetter_) returns()
func (_AmmFactory *AmmFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, feeToSetter_ common.Address) (*types.Transaction, error) {
	return _AmmFactory.contract.Transact(opts, "setFeeToSetter", feeToSetter_)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address feeToSetter_) returns()
func (_AmmFactory *AmmFactorySession) SetFeeToSetter(feeToSetter_ common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.SetFeeToSetter(&_AmmFactory.TransactOpts, feeToSetter_)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address feeToSetter_) returns()
func (_AmmFactory *AmmFactoryTransactorSession) SetFeeToSetter(feeToSetter_ common.Address) (*types.Transaction, error) {
	return _AmmFactory.Contract.SetFeeToSetter(&_AmmFactory.TransactOpts, feeToSetter_)
}

// AmmFactoryAmmCreatedIterator is returned from FilterAmmCreated and is used to iterate over the raw logs and unpacked data for AmmCreated events raised by the AmmFactory contract.
type AmmFactoryAmmCreatedIterator struct {
	Event *AmmFactoryAmmCreated // Event containing the contract specifics and raw log

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
func (it *AmmFactoryAmmCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmFactoryAmmCreated)
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
		it.Event = new(AmmFactoryAmmCreated)
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
func (it *AmmFactoryAmmCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmFactoryAmmCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmFactoryAmmCreated represents a AmmCreated event raised by the AmmFactory contract.
type AmmFactoryAmmCreated struct {
	BaseToken  common.Address
	QuoteToken common.Address
	Amm        common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAmmCreated is a free log retrieval operation binding the contract event 0x8eef5cf520576c00176a77de7eeee6762c3f44c7231914ce6c5c01e03a703820.
//
// Solidity: event AmmCreated(address indexed baseToken, address indexed quoteToken, address amm)
func (_AmmFactory *AmmFactoryFilterer) FilterAmmCreated(opts *bind.FilterOpts, baseToken []common.Address, quoteToken []common.Address) (*AmmFactoryAmmCreatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _AmmFactory.contract.FilterLogs(opts, "AmmCreated", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &AmmFactoryAmmCreatedIterator{contract: _AmmFactory.contract, event: "AmmCreated", logs: logs, sub: sub}, nil
}

// WatchAmmCreated is a free log subscription operation binding the contract event 0x8eef5cf520576c00176a77de7eeee6762c3f44c7231914ce6c5c01e03a703820.
//
// Solidity: event AmmCreated(address indexed baseToken, address indexed quoteToken, address amm)
func (_AmmFactory *AmmFactoryFilterer) WatchAmmCreated(opts *bind.WatchOpts, sink chan<- *AmmFactoryAmmCreated, baseToken []common.Address, quoteToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _AmmFactory.contract.WatchLogs(opts, "AmmCreated", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmFactoryAmmCreated)
				if err := _AmmFactory.contract.UnpackLog(event, "AmmCreated", log); err != nil {
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

// ParseAmmCreated is a log parse operation binding the contract event 0x8eef5cf520576c00176a77de7eeee6762c3f44c7231914ce6c5c01e03a703820.
//
// Solidity: event AmmCreated(address indexed baseToken, address indexed quoteToken, address amm)
func (_AmmFactory *AmmFactoryFilterer) ParseAmmCreated(log types.Log) (*AmmFactoryAmmCreated, error) {
	event := new(AmmFactoryAmmCreated)
	if err := _AmmFactory.contract.UnpackLog(event, "AmmCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
