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

// MarginFactoryMetaData contains all meta data concerning the MarginFactory contract.
var MarginFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"upperFactory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"config_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"margin\",\"type\":\"address\"}],\"name\":\"MarginCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"createMargin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"margin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getMargin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"name\":\"initMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MarginFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MarginFactoryMetaData.ABI instead.
var MarginFactoryABI = MarginFactoryMetaData.ABI

// MarginFactory is an auto generated Go binding around an Ethereum contract.
type MarginFactory struct {
	MarginFactoryCaller     // Read-only binding to the contract
	MarginFactoryTransactor // Write-only binding to the contract
	MarginFactoryFilterer   // Log filterer for contract events
}

// MarginFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarginFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarginFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarginFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarginFactorySession struct {
	Contract     *MarginFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarginFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarginFactoryCallerSession struct {
	Contract *MarginFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MarginFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarginFactoryTransactorSession struct {
	Contract     *MarginFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MarginFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarginFactoryRaw struct {
	Contract *MarginFactory // Generic contract binding to access the raw methods on
}

// MarginFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarginFactoryCallerRaw struct {
	Contract *MarginFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MarginFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarginFactoryTransactorRaw struct {
	Contract *MarginFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarginFactory creates a new instance of MarginFactory, bound to a specific deployed contract.
func NewMarginFactory(address common.Address, backend bind.ContractBackend) (*MarginFactory, error) {
	contract, err := bindMarginFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarginFactory{MarginFactoryCaller: MarginFactoryCaller{contract: contract}, MarginFactoryTransactor: MarginFactoryTransactor{contract: contract}, MarginFactoryFilterer: MarginFactoryFilterer{contract: contract}}, nil
}

// NewMarginFactoryCaller creates a new read-only instance of MarginFactory, bound to a specific deployed contract.
func NewMarginFactoryCaller(address common.Address, caller bind.ContractCaller) (*MarginFactoryCaller, error) {
	contract, err := bindMarginFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarginFactoryCaller{contract: contract}, nil
}

// NewMarginFactoryTransactor creates a new write-only instance of MarginFactory, bound to a specific deployed contract.
func NewMarginFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MarginFactoryTransactor, error) {
	contract, err := bindMarginFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarginFactoryTransactor{contract: contract}, nil
}

// NewMarginFactoryFilterer creates a new log filterer instance of MarginFactory, bound to a specific deployed contract.
func NewMarginFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MarginFactoryFilterer, error) {
	contract, err := bindMarginFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarginFactoryFilterer{contract: contract}, nil
}

// bindMarginFactory binds a generic wrapper to an already deployed contract.
func bindMarginFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarginFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarginFactory *MarginFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarginFactory.Contract.MarginFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarginFactory *MarginFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarginFactory.Contract.MarginFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarginFactory *MarginFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarginFactory.Contract.MarginFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarginFactory *MarginFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarginFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarginFactory *MarginFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarginFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarginFactory *MarginFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarginFactory.Contract.contract.Transact(opts, method, params...)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_MarginFactory *MarginFactoryCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarginFactory.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_MarginFactory *MarginFactorySession) Config() (common.Address, error) {
	return _MarginFactory.Contract.Config(&_MarginFactory.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_MarginFactory *MarginFactoryCallerSession) Config() (common.Address, error) {
	return _MarginFactory.Contract.Config(&_MarginFactory.CallOpts)
}

// GetMargin is a free data retrieval call binding the contract method 0x0db58602.
//
// Solidity: function getMargin(address , address ) view returns(address)
func (_MarginFactory *MarginFactoryCaller) GetMargin(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MarginFactory.contract.Call(opts, &out, "getMargin", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMargin is a free data retrieval call binding the contract method 0x0db58602.
//
// Solidity: function getMargin(address , address ) view returns(address)
func (_MarginFactory *MarginFactorySession) GetMargin(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _MarginFactory.Contract.GetMargin(&_MarginFactory.CallOpts, arg0, arg1)
}

// GetMargin is a free data retrieval call binding the contract method 0x0db58602.
//
// Solidity: function getMargin(address , address ) view returns(address)
func (_MarginFactory *MarginFactoryCallerSession) GetMargin(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _MarginFactory.Contract.GetMargin(&_MarginFactory.CallOpts, arg0, arg1)
}

// UpperFactory is a free data retrieval call binding the contract method 0xcedc12d8.
//
// Solidity: function upperFactory() view returns(address)
func (_MarginFactory *MarginFactoryCaller) UpperFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarginFactory.contract.Call(opts, &out, "upperFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpperFactory is a free data retrieval call binding the contract method 0xcedc12d8.
//
// Solidity: function upperFactory() view returns(address)
func (_MarginFactory *MarginFactorySession) UpperFactory() (common.Address, error) {
	return _MarginFactory.Contract.UpperFactory(&_MarginFactory.CallOpts)
}

// UpperFactory is a free data retrieval call binding the contract method 0xcedc12d8.
//
// Solidity: function upperFactory() view returns(address)
func (_MarginFactory *MarginFactoryCallerSession) UpperFactory() (common.Address, error) {
	return _MarginFactory.Contract.UpperFactory(&_MarginFactory.CallOpts)
}

// CreateMargin is a paid mutator transaction binding the contract method 0x0211406b.
//
// Solidity: function createMargin(address baseToken, address quoteToken) returns(address margin)
func (_MarginFactory *MarginFactoryTransactor) CreateMargin(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _MarginFactory.contract.Transact(opts, "createMargin", baseToken, quoteToken)
}

// CreateMargin is a paid mutator transaction binding the contract method 0x0211406b.
//
// Solidity: function createMargin(address baseToken, address quoteToken) returns(address margin)
func (_MarginFactory *MarginFactorySession) CreateMargin(baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _MarginFactory.Contract.CreateMargin(&_MarginFactory.TransactOpts, baseToken, quoteToken)
}

// CreateMargin is a paid mutator transaction binding the contract method 0x0211406b.
//
// Solidity: function createMargin(address baseToken, address quoteToken) returns(address margin)
func (_MarginFactory *MarginFactoryTransactorSession) CreateMargin(baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _MarginFactory.Contract.CreateMargin(&_MarginFactory.TransactOpts, baseToken, quoteToken)
}

// InitMargin is a paid mutator transaction binding the contract method 0x9703dd94.
//
// Solidity: function initMargin(address baseToken, address quoteToken, address amm) returns()
func (_MarginFactory *MarginFactoryTransactor) InitMargin(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, amm common.Address) (*types.Transaction, error) {
	return _MarginFactory.contract.Transact(opts, "initMargin", baseToken, quoteToken, amm)
}

// InitMargin is a paid mutator transaction binding the contract method 0x9703dd94.
//
// Solidity: function initMargin(address baseToken, address quoteToken, address amm) returns()
func (_MarginFactory *MarginFactorySession) InitMargin(baseToken common.Address, quoteToken common.Address, amm common.Address) (*types.Transaction, error) {
	return _MarginFactory.Contract.InitMargin(&_MarginFactory.TransactOpts, baseToken, quoteToken, amm)
}

// InitMargin is a paid mutator transaction binding the contract method 0x9703dd94.
//
// Solidity: function initMargin(address baseToken, address quoteToken, address amm) returns()
func (_MarginFactory *MarginFactoryTransactorSession) InitMargin(baseToken common.Address, quoteToken common.Address, amm common.Address) (*types.Transaction, error) {
	return _MarginFactory.Contract.InitMargin(&_MarginFactory.TransactOpts, baseToken, quoteToken, amm)
}

// MarginFactoryMarginCreatedIterator is returned from FilterMarginCreated and is used to iterate over the raw logs and unpacked data for MarginCreated events raised by the MarginFactory contract.
type MarginFactoryMarginCreatedIterator struct {
	Event *MarginFactoryMarginCreated // Event containing the contract specifics and raw log

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
func (it *MarginFactoryMarginCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginFactoryMarginCreated)
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
		it.Event = new(MarginFactoryMarginCreated)
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
func (it *MarginFactoryMarginCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginFactoryMarginCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginFactoryMarginCreated represents a MarginCreated event raised by the MarginFactory contract.
type MarginFactoryMarginCreated struct {
	BaseToken  common.Address
	QuoteToken common.Address
	Margin     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMarginCreated is a free log retrieval operation binding the contract event 0x83346cf05e0d18b120f8744ef712d3d618c570274a2f60ff15138b7177235644.
//
// Solidity: event MarginCreated(address indexed baseToken, address indexed quoteToken, address margin)
func (_MarginFactory *MarginFactoryFilterer) FilterMarginCreated(opts *bind.FilterOpts, baseToken []common.Address, quoteToken []common.Address) (*MarginFactoryMarginCreatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _MarginFactory.contract.FilterLogs(opts, "MarginCreated", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &MarginFactoryMarginCreatedIterator{contract: _MarginFactory.contract, event: "MarginCreated", logs: logs, sub: sub}, nil
}

// WatchMarginCreated is a free log subscription operation binding the contract event 0x83346cf05e0d18b120f8744ef712d3d618c570274a2f60ff15138b7177235644.
//
// Solidity: event MarginCreated(address indexed baseToken, address indexed quoteToken, address margin)
func (_MarginFactory *MarginFactoryFilterer) WatchMarginCreated(opts *bind.WatchOpts, sink chan<- *MarginFactoryMarginCreated, baseToken []common.Address, quoteToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _MarginFactory.contract.WatchLogs(opts, "MarginCreated", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginFactoryMarginCreated)
				if err := _MarginFactory.contract.UnpackLog(event, "MarginCreated", log); err != nil {
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

// ParseMarginCreated is a log parse operation binding the contract event 0x83346cf05e0d18b120f8744ef712d3d618c570274a2f60ff15138b7177235644.
//
// Solidity: event MarginCreated(address indexed baseToken, address indexed quoteToken, address margin)
func (_MarginFactory *MarginFactoryFilterer) ParseMarginCreated(log types.Log) (*MarginFactoryMarginCreated, error) {
	event := new(MarginFactoryMarginCreated)
	if err := _MarginFactory.contract.UnpackLog(event, "MarginCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
