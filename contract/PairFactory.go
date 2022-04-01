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

// PairFactoryMetaData contains all meta data concerning the PairFactory contract.
var PairFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"margin\",\"type\":\"address\"}],\"name\":\"NewPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"margin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getAmm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getMargin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ammFactory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marginFactory_\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PairFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PairFactoryMetaData.ABI instead.
var PairFactoryABI = PairFactoryMetaData.ABI

// PairFactory is an auto generated Go binding around an Ethereum contract.
type PairFactory struct {
	PairFactoryCaller     // Read-only binding to the contract
	PairFactoryTransactor // Write-only binding to the contract
	PairFactoryFilterer   // Log filterer for contract events
}

// PairFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairFactorySession struct {
	Contract     *PairFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairFactoryCallerSession struct {
	Contract *PairFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PairFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairFactoryTransactorSession struct {
	Contract     *PairFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PairFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairFactoryRaw struct {
	Contract *PairFactory // Generic contract binding to access the raw methods on
}

// PairFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairFactoryCallerRaw struct {
	Contract *PairFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PairFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairFactoryTransactorRaw struct {
	Contract *PairFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairFactory creates a new instance of PairFactory, bound to a specific deployed contract.
func NewPairFactory(address common.Address, backend bind.ContractBackend) (*PairFactory, error) {
	contract, err := bindPairFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairFactory{PairFactoryCaller: PairFactoryCaller{contract: contract}, PairFactoryTransactor: PairFactoryTransactor{contract: contract}, PairFactoryFilterer: PairFactoryFilterer{contract: contract}}, nil
}

// NewPairFactoryCaller creates a new read-only instance of PairFactory, bound to a specific deployed contract.
func NewPairFactoryCaller(address common.Address, caller bind.ContractCaller) (*PairFactoryCaller, error) {
	contract, err := bindPairFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairFactoryCaller{contract: contract}, nil
}

// NewPairFactoryTransactor creates a new write-only instance of PairFactory, bound to a specific deployed contract.
func NewPairFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PairFactoryTransactor, error) {
	contract, err := bindPairFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairFactoryTransactor{contract: contract}, nil
}

// NewPairFactoryFilterer creates a new log filterer instance of PairFactory, bound to a specific deployed contract.
func NewPairFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PairFactoryFilterer, error) {
	contract, err := bindPairFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairFactoryFilterer{contract: contract}, nil
}

// bindPairFactory binds a generic wrapper to an already deployed contract.
func bindPairFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairFactory *PairFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairFactory.Contract.PairFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairFactory *PairFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairFactory.Contract.PairFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairFactory *PairFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairFactory.Contract.PairFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairFactory *PairFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairFactory *PairFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairFactory *PairFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairFactory.Contract.contract.Transact(opts, method, params...)
}

// AmmFactory is a free data retrieval call binding the contract method 0xdacda92f.
//
// Solidity: function ammFactory() view returns(address)
func (_PairFactory *PairFactoryCaller) AmmFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairFactory.contract.Call(opts, &out, "ammFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmFactory is a free data retrieval call binding the contract method 0xdacda92f.
//
// Solidity: function ammFactory() view returns(address)
func (_PairFactory *PairFactorySession) AmmFactory() (common.Address, error) {
	return _PairFactory.Contract.AmmFactory(&_PairFactory.CallOpts)
}

// AmmFactory is a free data retrieval call binding the contract method 0xdacda92f.
//
// Solidity: function ammFactory() view returns(address)
func (_PairFactory *PairFactoryCallerSession) AmmFactory() (common.Address, error) {
	return _PairFactory.Contract.AmmFactory(&_PairFactory.CallOpts)
}

// GetAmm is a free data retrieval call binding the contract method 0x65cc7a1d.
//
// Solidity: function getAmm(address baseToken, address quoteToken) view returns(address)
func (_PairFactory *PairFactoryCaller) GetAmm(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _PairFactory.contract.Call(opts, &out, "getAmm", baseToken, quoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAmm is a free data retrieval call binding the contract method 0x65cc7a1d.
//
// Solidity: function getAmm(address baseToken, address quoteToken) view returns(address)
func (_PairFactory *PairFactorySession) GetAmm(baseToken common.Address, quoteToken common.Address) (common.Address, error) {
	return _PairFactory.Contract.GetAmm(&_PairFactory.CallOpts, baseToken, quoteToken)
}

// GetAmm is a free data retrieval call binding the contract method 0x65cc7a1d.
//
// Solidity: function getAmm(address baseToken, address quoteToken) view returns(address)
func (_PairFactory *PairFactoryCallerSession) GetAmm(baseToken common.Address, quoteToken common.Address) (common.Address, error) {
	return _PairFactory.Contract.GetAmm(&_PairFactory.CallOpts, baseToken, quoteToken)
}

// GetMargin is a free data retrieval call binding the contract method 0x0db58602.
//
// Solidity: function getMargin(address baseToken, address quoteToken) view returns(address)
func (_PairFactory *PairFactoryCaller) GetMargin(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _PairFactory.contract.Call(opts, &out, "getMargin", baseToken, quoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMargin is a free data retrieval call binding the contract method 0x0db58602.
//
// Solidity: function getMargin(address baseToken, address quoteToken) view returns(address)
func (_PairFactory *PairFactorySession) GetMargin(baseToken common.Address, quoteToken common.Address) (common.Address, error) {
	return _PairFactory.Contract.GetMargin(&_PairFactory.CallOpts, baseToken, quoteToken)
}

// GetMargin is a free data retrieval call binding the contract method 0x0db58602.
//
// Solidity: function getMargin(address baseToken, address quoteToken) view returns(address)
func (_PairFactory *PairFactoryCallerSession) GetMargin(baseToken common.Address, quoteToken common.Address) (common.Address, error) {
	return _PairFactory.Contract.GetMargin(&_PairFactory.CallOpts, baseToken, quoteToken)
}

// MarginFactory is a free data retrieval call binding the contract method 0xbacf68ed.
//
// Solidity: function marginFactory() view returns(address)
func (_PairFactory *PairFactoryCaller) MarginFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairFactory.contract.Call(opts, &out, "marginFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarginFactory is a free data retrieval call binding the contract method 0xbacf68ed.
//
// Solidity: function marginFactory() view returns(address)
func (_PairFactory *PairFactorySession) MarginFactory() (common.Address, error) {
	return _PairFactory.Contract.MarginFactory(&_PairFactory.CallOpts)
}

// MarginFactory is a free data retrieval call binding the contract method 0xbacf68ed.
//
// Solidity: function marginFactory() view returns(address)
func (_PairFactory *PairFactoryCallerSession) MarginFactory() (common.Address, error) {
	return _PairFactory.Contract.MarginFactory(&_PairFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairFactory *PairFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairFactory *PairFactorySession) Owner() (common.Address, error) {
	return _PairFactory.Contract.Owner(&_PairFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairFactory *PairFactoryCallerSession) Owner() (common.Address, error) {
	return _PairFactory.Contract.Owner(&_PairFactory.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PairFactory *PairFactoryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairFactory.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PairFactory *PairFactorySession) PendingOwner() (common.Address, error) {
	return _PairFactory.Contract.PendingOwner(&_PairFactory.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PairFactory *PairFactoryCallerSession) PendingOwner() (common.Address, error) {
	return _PairFactory.Contract.PendingOwner(&_PairFactory.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_PairFactory *PairFactoryTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairFactory.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_PairFactory *PairFactorySession) AcceptOwner() (*types.Transaction, error) {
	return _PairFactory.Contract.AcceptOwner(&_PairFactory.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_PairFactory *PairFactoryTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _PairFactory.Contract.AcceptOwner(&_PairFactory.TransactOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address baseToken, address quoteToken) returns(address amm, address margin)
func (_PairFactory *PairFactoryTransactor) CreatePair(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _PairFactory.contract.Transact(opts, "createPair", baseToken, quoteToken)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address baseToken, address quoteToken) returns(address amm, address margin)
func (_PairFactory *PairFactorySession) CreatePair(baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _PairFactory.Contract.CreatePair(&_PairFactory.TransactOpts, baseToken, quoteToken)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address baseToken, address quoteToken) returns(address amm, address margin)
func (_PairFactory *PairFactoryTransactorSession) CreatePair(baseToken common.Address, quoteToken common.Address) (*types.Transaction, error) {
	return _PairFactory.Contract.CreatePair(&_PairFactory.TransactOpts, baseToken, quoteToken)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address ammFactory_, address marginFactory_) returns()
func (_PairFactory *PairFactoryTransactor) Init(opts *bind.TransactOpts, ammFactory_ common.Address, marginFactory_ common.Address) (*types.Transaction, error) {
	return _PairFactory.contract.Transact(opts, "init", ammFactory_, marginFactory_)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address ammFactory_, address marginFactory_) returns()
func (_PairFactory *PairFactorySession) Init(ammFactory_ common.Address, marginFactory_ common.Address) (*types.Transaction, error) {
	return _PairFactory.Contract.Init(&_PairFactory.TransactOpts, ammFactory_, marginFactory_)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address ammFactory_, address marginFactory_) returns()
func (_PairFactory *PairFactoryTransactorSession) Init(ammFactory_ common.Address, marginFactory_ common.Address) (*types.Transaction, error) {
	return _PairFactory.Contract.Init(&_PairFactory.TransactOpts, ammFactory_, marginFactory_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_PairFactory *PairFactoryTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _PairFactory.contract.Transact(opts, "setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_PairFactory *PairFactorySession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _PairFactory.Contract.SetPendingOwner(&_PairFactory.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_PairFactory *PairFactoryTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _PairFactory.Contract.SetPendingOwner(&_PairFactory.TransactOpts, newPendingOwner)
}

// PairFactoryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the PairFactory contract.
type PairFactoryNewOwnerIterator struct {
	Event *PairFactoryNewOwner // Event containing the contract specifics and raw log

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
func (it *PairFactoryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairFactoryNewOwner)
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
		it.Event = new(PairFactoryNewOwner)
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
func (it *PairFactoryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairFactoryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairFactoryNewOwner represents a NewOwner event raised by the PairFactory contract.
type PairFactoryNewOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_PairFactory *PairFactoryFilterer) FilterNewOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PairFactoryNewOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PairFactory.contract.FilterLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PairFactoryNewOwnerIterator{contract: _PairFactory.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_PairFactory *PairFactoryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *PairFactoryNewOwner, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PairFactory.contract.WatchLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairFactoryNewOwner)
				if err := _PairFactory.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_PairFactory *PairFactoryFilterer) ParseNewOwner(log types.Log) (*PairFactoryNewOwner, error) {
	event := new(PairFactoryNewOwner)
	if err := _PairFactory.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairFactoryNewPairIterator is returned from FilterNewPair and is used to iterate over the raw logs and unpacked data for NewPair events raised by the PairFactory contract.
type PairFactoryNewPairIterator struct {
	Event *PairFactoryNewPair // Event containing the contract specifics and raw log

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
func (it *PairFactoryNewPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairFactoryNewPair)
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
		it.Event = new(PairFactoryNewPair)
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
func (it *PairFactoryNewPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairFactoryNewPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairFactoryNewPair represents a NewPair event raised by the PairFactory contract.
type PairFactoryNewPair struct {
	BaseToken  common.Address
	QuoteToken common.Address
	Amm        common.Address
	Margin     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewPair is a free log retrieval operation binding the contract event 0xb231fe977931114284ce64d77cdfe4bcb8121c500a095d281b55216f7e225854.
//
// Solidity: event NewPair(address indexed baseToken, address indexed quoteToken, address amm, address margin)
func (_PairFactory *PairFactoryFilterer) FilterNewPair(opts *bind.FilterOpts, baseToken []common.Address, quoteToken []common.Address) (*PairFactoryNewPairIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _PairFactory.contract.FilterLogs(opts, "NewPair", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &PairFactoryNewPairIterator{contract: _PairFactory.contract, event: "NewPair", logs: logs, sub: sub}, nil
}

// WatchNewPair is a free log subscription operation binding the contract event 0xb231fe977931114284ce64d77cdfe4bcb8121c500a095d281b55216f7e225854.
//
// Solidity: event NewPair(address indexed baseToken, address indexed quoteToken, address amm, address margin)
func (_PairFactory *PairFactoryFilterer) WatchNewPair(opts *bind.WatchOpts, sink chan<- *PairFactoryNewPair, baseToken []common.Address, quoteToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _PairFactory.contract.WatchLogs(opts, "NewPair", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairFactoryNewPair)
				if err := _PairFactory.contract.UnpackLog(event, "NewPair", log); err != nil {
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

// ParseNewPair is a log parse operation binding the contract event 0xb231fe977931114284ce64d77cdfe4bcb8121c500a095d281b55216f7e225854.
//
// Solidity: event NewPair(address indexed baseToken, address indexed quoteToken, address amm, address margin)
func (_PairFactory *PairFactoryFilterer) ParseNewPair(log types.Log) (*PairFactoryNewPair, error) {
	event := new(PairFactoryNewPair)
	if err := _PairFactory.contract.UnpackLog(event, "NewPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairFactoryNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the PairFactory contract.
type PairFactoryNewPendingOwnerIterator struct {
	Event *PairFactoryNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *PairFactoryNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairFactoryNewPendingOwner)
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
		it.Event = new(PairFactoryNewPendingOwner)
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
func (it *PairFactoryNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairFactoryNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairFactoryNewPendingOwner represents a NewPendingOwner event raised by the PairFactory contract.
type PairFactoryNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_PairFactory *PairFactoryFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*PairFactoryNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _PairFactory.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PairFactoryNewPendingOwnerIterator{contract: _PairFactory.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_PairFactory *PairFactoryFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *PairFactoryNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _PairFactory.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairFactoryNewPendingOwner)
				if err := _PairFactory.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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

// ParseNewPendingOwner is a log parse operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_PairFactory *PairFactoryFilterer) ParseNewPendingOwner(log types.Log) (*PairFactoryNewPendingOwner, error) {
	event := new(PairFactoryNewPendingOwner)
	if err := _PairFactory.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
