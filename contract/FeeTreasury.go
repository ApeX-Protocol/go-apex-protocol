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

// FeeTreasuryMetaData contains all meta data concerning the FeeTreasury contract.
var FeeTreasuryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"contractISwapRouter\",\"name\":\"v3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"USDC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextSettleTime_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardForCashback\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DistributeToCashback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardForStaking\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DistributeToStaking\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldRatio\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newRatio\",\"type\":\"uint8\"}],\"name\":\"RatioForStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldReward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newReward\",\"type\":\"address\"}],\"name\":\"RewardForCashbackChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldReward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newReward\",\"type\":\"address\"}],\"name\":\"RewardForStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInterval\",\"type\":\"uint256\"}],\"name\":\"SettlementIntervalChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amms\",\"type\":\"address[]\"}],\"name\":\"batchRemoveLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"batchSwapToETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSettleTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratioForStaking\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardForCashback\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardForStaking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newrRatio\",\"type\":\"uint8\"}],\"name\":\"setRatioForStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReward\",\"type\":\"address\"}],\"name\":\"setRewardForCashback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReward\",\"type\":\"address\"}],\"name\":\"setRewardForStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInterval\",\"type\":\"uint256\"}],\"name\":\"setSettlementInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"v3Factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"v3Fees\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"v3Router\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FeeTreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeTreasuryMetaData.ABI instead.
var FeeTreasuryABI = FeeTreasuryMetaData.ABI

// FeeTreasury is an auto generated Go binding around an Ethereum contract.
type FeeTreasury struct {
	FeeTreasuryCaller     // Read-only binding to the contract
	FeeTreasuryTransactor // Write-only binding to the contract
	FeeTreasuryFilterer   // Log filterer for contract events
}

// FeeTreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeTreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeTreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeTreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeTreasurySession struct {
	Contract     *FeeTreasury      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeTreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeTreasuryCallerSession struct {
	Contract *FeeTreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FeeTreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeTreasuryTransactorSession struct {
	Contract     *FeeTreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FeeTreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeTreasuryRaw struct {
	Contract *FeeTreasury // Generic contract binding to access the raw methods on
}

// FeeTreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeTreasuryCallerRaw struct {
	Contract *FeeTreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// FeeTreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeTreasuryTransactorRaw struct {
	Contract *FeeTreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeTreasury creates a new instance of FeeTreasury, bound to a specific deployed contract.
func NewFeeTreasury(address common.Address, backend bind.ContractBackend) (*FeeTreasury, error) {
	contract, err := bindFeeTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeTreasury{FeeTreasuryCaller: FeeTreasuryCaller{contract: contract}, FeeTreasuryTransactor: FeeTreasuryTransactor{contract: contract}, FeeTreasuryFilterer: FeeTreasuryFilterer{contract: contract}}, nil
}

// NewFeeTreasuryCaller creates a new read-only instance of FeeTreasury, bound to a specific deployed contract.
func NewFeeTreasuryCaller(address common.Address, caller bind.ContractCaller) (*FeeTreasuryCaller, error) {
	contract, err := bindFeeTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryCaller{contract: contract}, nil
}

// NewFeeTreasuryTransactor creates a new write-only instance of FeeTreasury, bound to a specific deployed contract.
func NewFeeTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeTreasuryTransactor, error) {
	contract, err := bindFeeTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryTransactor{contract: contract}, nil
}

// NewFeeTreasuryFilterer creates a new log filterer instance of FeeTreasury, bound to a specific deployed contract.
func NewFeeTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeTreasuryFilterer, error) {
	contract, err := bindFeeTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryFilterer{contract: contract}, nil
}

// bindFeeTreasury binds a generic wrapper to an already deployed contract.
func bindFeeTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeTreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeTreasury *FeeTreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeTreasury.Contract.FeeTreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeTreasury *FeeTreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTreasury.Contract.FeeTreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeTreasury *FeeTreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeTreasury.Contract.FeeTreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeTreasury *FeeTreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeTreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeTreasury *FeeTreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeTreasury *FeeTreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeTreasury.Contract.contract.Transact(opts, method, params...)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_FeeTreasury *FeeTreasurySession) USDC() (common.Address, error) {
	return _FeeTreasury.Contract.USDC(&_FeeTreasury.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) USDC() (common.Address, error) {
	return _FeeTreasury.Contract.USDC(&_FeeTreasury.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_FeeTreasury *FeeTreasurySession) WETH() (common.Address, error) {
	return _FeeTreasury.Contract.WETH(&_FeeTreasury.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) WETH() (common.Address, error) {
	return _FeeTreasury.Contract.WETH(&_FeeTreasury.CallOpts)
}

// NextSettleTime is a free data retrieval call binding the contract method 0xcd823e50.
//
// Solidity: function nextSettleTime() view returns(uint256)
func (_FeeTreasury *FeeTreasuryCaller) NextSettleTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "nextSettleTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSettleTime is a free data retrieval call binding the contract method 0xcd823e50.
//
// Solidity: function nextSettleTime() view returns(uint256)
func (_FeeTreasury *FeeTreasurySession) NextSettleTime() (*big.Int, error) {
	return _FeeTreasury.Contract.NextSettleTime(&_FeeTreasury.CallOpts)
}

// NextSettleTime is a free data retrieval call binding the contract method 0xcd823e50.
//
// Solidity: function nextSettleTime() view returns(uint256)
func (_FeeTreasury *FeeTreasuryCallerSession) NextSettleTime() (*big.Int, error) {
	return _FeeTreasury.Contract.NextSettleTime(&_FeeTreasury.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_FeeTreasury *FeeTreasurySession) Operator() (common.Address, error) {
	return _FeeTreasury.Contract.Operator(&_FeeTreasury.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) Operator() (common.Address, error) {
	return _FeeTreasury.Contract.Operator(&_FeeTreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeTreasury *FeeTreasurySession) Owner() (common.Address, error) {
	return _FeeTreasury.Contract.Owner(&_FeeTreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) Owner() (common.Address, error) {
	return _FeeTreasury.Contract.Owner(&_FeeTreasury.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FeeTreasury *FeeTreasurySession) PendingOwner() (common.Address, error) {
	return _FeeTreasury.Contract.PendingOwner(&_FeeTreasury.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) PendingOwner() (common.Address, error) {
	return _FeeTreasury.Contract.PendingOwner(&_FeeTreasury.CallOpts)
}

// RatioForStaking is a free data retrieval call binding the contract method 0x42dddc96.
//
// Solidity: function ratioForStaking() view returns(uint8)
func (_FeeTreasury *FeeTreasuryCaller) RatioForStaking(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "ratioForStaking")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RatioForStaking is a free data retrieval call binding the contract method 0x42dddc96.
//
// Solidity: function ratioForStaking() view returns(uint8)
func (_FeeTreasury *FeeTreasurySession) RatioForStaking() (uint8, error) {
	return _FeeTreasury.Contract.RatioForStaking(&_FeeTreasury.CallOpts)
}

// RatioForStaking is a free data retrieval call binding the contract method 0x42dddc96.
//
// Solidity: function ratioForStaking() view returns(uint8)
func (_FeeTreasury *FeeTreasuryCallerSession) RatioForStaking() (uint8, error) {
	return _FeeTreasury.Contract.RatioForStaking(&_FeeTreasury.CallOpts)
}

// RewardForCashback is a free data retrieval call binding the contract method 0xb56dc36b.
//
// Solidity: function rewardForCashback() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) RewardForCashback(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "rewardForCashback")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardForCashback is a free data retrieval call binding the contract method 0xb56dc36b.
//
// Solidity: function rewardForCashback() view returns(address)
func (_FeeTreasury *FeeTreasurySession) RewardForCashback() (common.Address, error) {
	return _FeeTreasury.Contract.RewardForCashback(&_FeeTreasury.CallOpts)
}

// RewardForCashback is a free data retrieval call binding the contract method 0xb56dc36b.
//
// Solidity: function rewardForCashback() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) RewardForCashback() (common.Address, error) {
	return _FeeTreasury.Contract.RewardForCashback(&_FeeTreasury.CallOpts)
}

// RewardForStaking is a free data retrieval call binding the contract method 0xa7d67691.
//
// Solidity: function rewardForStaking() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) RewardForStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "rewardForStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardForStaking is a free data retrieval call binding the contract method 0xa7d67691.
//
// Solidity: function rewardForStaking() view returns(address)
func (_FeeTreasury *FeeTreasurySession) RewardForStaking() (common.Address, error) {
	return _FeeTreasury.Contract.RewardForStaking(&_FeeTreasury.CallOpts)
}

// RewardForStaking is a free data retrieval call binding the contract method 0xa7d67691.
//
// Solidity: function rewardForStaking() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) RewardForStaking() (common.Address, error) {
	return _FeeTreasury.Contract.RewardForStaking(&_FeeTreasury.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FeeTreasury *FeeTreasurySession) Router() (common.Address, error) {
	return _FeeTreasury.Contract.Router(&_FeeTreasury.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) Router() (common.Address, error) {
	return _FeeTreasury.Contract.Router(&_FeeTreasury.CallOpts)
}

// SettlementInterval is a free data retrieval call binding the contract method 0x7ec4142c.
//
// Solidity: function settlementInterval() view returns(uint256)
func (_FeeTreasury *FeeTreasuryCaller) SettlementInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "settlementInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettlementInterval is a free data retrieval call binding the contract method 0x7ec4142c.
//
// Solidity: function settlementInterval() view returns(uint256)
func (_FeeTreasury *FeeTreasurySession) SettlementInterval() (*big.Int, error) {
	return _FeeTreasury.Contract.SettlementInterval(&_FeeTreasury.CallOpts)
}

// SettlementInterval is a free data retrieval call binding the contract method 0x7ec4142c.
//
// Solidity: function settlementInterval() view returns(uint256)
func (_FeeTreasury *FeeTreasuryCallerSession) SettlementInterval() (*big.Int, error) {
	return _FeeTreasury.Contract.SettlementInterval(&_FeeTreasury.CallOpts)
}

// V3Factory is a free data retrieval call binding the contract method 0x7c887c59.
//
// Solidity: function v3Factory() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) V3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "v3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V3Factory is a free data retrieval call binding the contract method 0x7c887c59.
//
// Solidity: function v3Factory() view returns(address)
func (_FeeTreasury *FeeTreasurySession) V3Factory() (common.Address, error) {
	return _FeeTreasury.Contract.V3Factory(&_FeeTreasury.CallOpts)
}

// V3Factory is a free data retrieval call binding the contract method 0x7c887c59.
//
// Solidity: function v3Factory() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) V3Factory() (common.Address, error) {
	return _FeeTreasury.Contract.V3Factory(&_FeeTreasury.CallOpts)
}

// V3Fees is a free data retrieval call binding the contract method 0xa11b4528.
//
// Solidity: function v3Fees(uint256 ) view returns(uint24)
func (_FeeTreasury *FeeTreasuryCaller) V3Fees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "v3Fees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// V3Fees is a free data retrieval call binding the contract method 0xa11b4528.
//
// Solidity: function v3Fees(uint256 ) view returns(uint24)
func (_FeeTreasury *FeeTreasurySession) V3Fees(arg0 *big.Int) (*big.Int, error) {
	return _FeeTreasury.Contract.V3Fees(&_FeeTreasury.CallOpts, arg0)
}

// V3Fees is a free data retrieval call binding the contract method 0xa11b4528.
//
// Solidity: function v3Fees(uint256 ) view returns(uint24)
func (_FeeTreasury *FeeTreasuryCallerSession) V3Fees(arg0 *big.Int) (*big.Int, error) {
	return _FeeTreasury.Contract.V3Fees(&_FeeTreasury.CallOpts, arg0)
}

// V3Router is a free data retrieval call binding the contract method 0x0dc91306.
//
// Solidity: function v3Router() view returns(address)
func (_FeeTreasury *FeeTreasuryCaller) V3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeTreasury.contract.Call(opts, &out, "v3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V3Router is a free data retrieval call binding the contract method 0x0dc91306.
//
// Solidity: function v3Router() view returns(address)
func (_FeeTreasury *FeeTreasurySession) V3Router() (common.Address, error) {
	return _FeeTreasury.Contract.V3Router(&_FeeTreasury.CallOpts)
}

// V3Router is a free data retrieval call binding the contract method 0x0dc91306.
//
// Solidity: function v3Router() view returns(address)
func (_FeeTreasury *FeeTreasuryCallerSession) V3Router() (common.Address, error) {
	return _FeeTreasury.Contract.V3Router(&_FeeTreasury.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_FeeTreasury *FeeTreasuryTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_FeeTreasury *FeeTreasurySession) AcceptOwner() (*types.Transaction, error) {
	return _FeeTreasury.Contract.AcceptOwner(&_FeeTreasury.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _FeeTreasury.Contract.AcceptOwner(&_FeeTreasury.TransactOpts)
}

// BatchRemoveLiquidity is a paid mutator transaction binding the contract method 0xff0e3f24.
//
// Solidity: function batchRemoveLiquidity(address[] amms) returns()
func (_FeeTreasury *FeeTreasuryTransactor) BatchRemoveLiquidity(opts *bind.TransactOpts, amms []common.Address) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "batchRemoveLiquidity", amms)
}

// BatchRemoveLiquidity is a paid mutator transaction binding the contract method 0xff0e3f24.
//
// Solidity: function batchRemoveLiquidity(address[] amms) returns()
func (_FeeTreasury *FeeTreasurySession) BatchRemoveLiquidity(amms []common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.BatchRemoveLiquidity(&_FeeTreasury.TransactOpts, amms)
}

// BatchRemoveLiquidity is a paid mutator transaction binding the contract method 0xff0e3f24.
//
// Solidity: function batchRemoveLiquidity(address[] amms) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) BatchRemoveLiquidity(amms []common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.BatchRemoveLiquidity(&_FeeTreasury.TransactOpts, amms)
}

// BatchSwapToETH is a paid mutator transaction binding the contract method 0x3cf38586.
//
// Solidity: function batchSwapToETH(address[] tokens) returns()
func (_FeeTreasury *FeeTreasuryTransactor) BatchSwapToETH(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "batchSwapToETH", tokens)
}

// BatchSwapToETH is a paid mutator transaction binding the contract method 0x3cf38586.
//
// Solidity: function batchSwapToETH(address[] tokens) returns()
func (_FeeTreasury *FeeTreasurySession) BatchSwapToETH(tokens []common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.BatchSwapToETH(&_FeeTreasury.TransactOpts, tokens)
}

// BatchSwapToETH is a paid mutator transaction binding the contract method 0x3cf38586.
//
// Solidity: function batchSwapToETH(address[] tokens) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) BatchSwapToETH(tokens []common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.BatchSwapToETH(&_FeeTreasury.TransactOpts, tokens)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_FeeTreasury *FeeTreasuryTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_FeeTreasury *FeeTreasurySession) Distribute() (*types.Transaction, error) {
	return _FeeTreasury.Contract.Distribute(&_FeeTreasury.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) Distribute() (*types.Transaction, error) {
	return _FeeTreasury.Contract.Distribute(&_FeeTreasury.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_FeeTreasury *FeeTreasuryTransactor) SetOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "setOperator", newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_FeeTreasury *FeeTreasurySession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetOperator(&_FeeTreasury.TransactOpts, newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetOperator(&_FeeTreasury.TransactOpts, newOperator)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_FeeTreasury *FeeTreasuryTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_FeeTreasury *FeeTreasurySession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetPendingOwner(&_FeeTreasury.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetPendingOwner(&_FeeTreasury.TransactOpts, newPendingOwner)
}

// SetRatioForStaking is a paid mutator transaction binding the contract method 0xc1a7c967.
//
// Solidity: function setRatioForStaking(uint8 newrRatio) returns()
func (_FeeTreasury *FeeTreasuryTransactor) SetRatioForStaking(opts *bind.TransactOpts, newrRatio uint8) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "setRatioForStaking", newrRatio)
}

// SetRatioForStaking is a paid mutator transaction binding the contract method 0xc1a7c967.
//
// Solidity: function setRatioForStaking(uint8 newrRatio) returns()
func (_FeeTreasury *FeeTreasurySession) SetRatioForStaking(newrRatio uint8) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetRatioForStaking(&_FeeTreasury.TransactOpts, newrRatio)
}

// SetRatioForStaking is a paid mutator transaction binding the contract method 0xc1a7c967.
//
// Solidity: function setRatioForStaking(uint8 newrRatio) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) SetRatioForStaking(newrRatio uint8) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetRatioForStaking(&_FeeTreasury.TransactOpts, newrRatio)
}

// SetRewardForCashback is a paid mutator transaction binding the contract method 0xc6d79c93.
//
// Solidity: function setRewardForCashback(address newReward) returns()
func (_FeeTreasury *FeeTreasuryTransactor) SetRewardForCashback(opts *bind.TransactOpts, newReward common.Address) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "setRewardForCashback", newReward)
}

// SetRewardForCashback is a paid mutator transaction binding the contract method 0xc6d79c93.
//
// Solidity: function setRewardForCashback(address newReward) returns()
func (_FeeTreasury *FeeTreasurySession) SetRewardForCashback(newReward common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetRewardForCashback(&_FeeTreasury.TransactOpts, newReward)
}

// SetRewardForCashback is a paid mutator transaction binding the contract method 0xc6d79c93.
//
// Solidity: function setRewardForCashback(address newReward) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) SetRewardForCashback(newReward common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetRewardForCashback(&_FeeTreasury.TransactOpts, newReward)
}

// SetRewardForStaking is a paid mutator transaction binding the contract method 0x2edba9fe.
//
// Solidity: function setRewardForStaking(address newReward) returns()
func (_FeeTreasury *FeeTreasuryTransactor) SetRewardForStaking(opts *bind.TransactOpts, newReward common.Address) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "setRewardForStaking", newReward)
}

// SetRewardForStaking is a paid mutator transaction binding the contract method 0x2edba9fe.
//
// Solidity: function setRewardForStaking(address newReward) returns()
func (_FeeTreasury *FeeTreasurySession) SetRewardForStaking(newReward common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetRewardForStaking(&_FeeTreasury.TransactOpts, newReward)
}

// SetRewardForStaking is a paid mutator transaction binding the contract method 0x2edba9fe.
//
// Solidity: function setRewardForStaking(address newReward) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) SetRewardForStaking(newReward common.Address) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetRewardForStaking(&_FeeTreasury.TransactOpts, newReward)
}

// SetSettlementInterval is a paid mutator transaction binding the contract method 0x5a683367.
//
// Solidity: function setSettlementInterval(uint256 newInterval) returns()
func (_FeeTreasury *FeeTreasuryTransactor) SetSettlementInterval(opts *bind.TransactOpts, newInterval *big.Int) (*types.Transaction, error) {
	return _FeeTreasury.contract.Transact(opts, "setSettlementInterval", newInterval)
}

// SetSettlementInterval is a paid mutator transaction binding the contract method 0x5a683367.
//
// Solidity: function setSettlementInterval(uint256 newInterval) returns()
func (_FeeTreasury *FeeTreasurySession) SetSettlementInterval(newInterval *big.Int) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetSettlementInterval(&_FeeTreasury.TransactOpts, newInterval)
}

// SetSettlementInterval is a paid mutator transaction binding the contract method 0x5a683367.
//
// Solidity: function setSettlementInterval(uint256 newInterval) returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) SetSettlementInterval(newInterval *big.Int) (*types.Transaction, error) {
	return _FeeTreasury.Contract.SetSettlementInterval(&_FeeTreasury.TransactOpts, newInterval)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeTreasury *FeeTreasuryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTreasury.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeTreasury *FeeTreasurySession) Receive() (*types.Transaction, error) {
	return _FeeTreasury.Contract.Receive(&_FeeTreasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeTreasury *FeeTreasuryTransactorSession) Receive() (*types.Transaction, error) {
	return _FeeTreasury.Contract.Receive(&_FeeTreasury.TransactOpts)
}

// FeeTreasuryDistributeToCashbackIterator is returned from FilterDistributeToCashback and is used to iterate over the raw logs and unpacked data for DistributeToCashback events raised by the FeeTreasury contract.
type FeeTreasuryDistributeToCashbackIterator struct {
	Event *FeeTreasuryDistributeToCashback // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryDistributeToCashbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryDistributeToCashback)
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
		it.Event = new(FeeTreasuryDistributeToCashback)
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
func (it *FeeTreasuryDistributeToCashbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryDistributeToCashbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryDistributeToCashback represents a DistributeToCashback event raised by the FeeTreasury contract.
type FeeTreasuryDistributeToCashback struct {
	RewardForCashback common.Address
	EthAmount         *big.Int
	UsdcAmount        *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDistributeToCashback is a free log retrieval operation binding the contract event 0x8c62aed53732e50473281886ba2b519cb08ee7b823e098fbe25d6e120d712c48.
//
// Solidity: event DistributeToCashback(address indexed rewardForCashback, uint256 ethAmount, uint256 usdcAmount, uint256 timestamp)
func (_FeeTreasury *FeeTreasuryFilterer) FilterDistributeToCashback(opts *bind.FilterOpts, rewardForCashback []common.Address) (*FeeTreasuryDistributeToCashbackIterator, error) {

	var rewardForCashbackRule []interface{}
	for _, rewardForCashbackItem := range rewardForCashback {
		rewardForCashbackRule = append(rewardForCashbackRule, rewardForCashbackItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "DistributeToCashback", rewardForCashbackRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryDistributeToCashbackIterator{contract: _FeeTreasury.contract, event: "DistributeToCashback", logs: logs, sub: sub}, nil
}

// WatchDistributeToCashback is a free log subscription operation binding the contract event 0x8c62aed53732e50473281886ba2b519cb08ee7b823e098fbe25d6e120d712c48.
//
// Solidity: event DistributeToCashback(address indexed rewardForCashback, uint256 ethAmount, uint256 usdcAmount, uint256 timestamp)
func (_FeeTreasury *FeeTreasuryFilterer) WatchDistributeToCashback(opts *bind.WatchOpts, sink chan<- *FeeTreasuryDistributeToCashback, rewardForCashback []common.Address) (event.Subscription, error) {

	var rewardForCashbackRule []interface{}
	for _, rewardForCashbackItem := range rewardForCashback {
		rewardForCashbackRule = append(rewardForCashbackRule, rewardForCashbackItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "DistributeToCashback", rewardForCashbackRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryDistributeToCashback)
				if err := _FeeTreasury.contract.UnpackLog(event, "DistributeToCashback", log); err != nil {
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

// ParseDistributeToCashback is a log parse operation binding the contract event 0x8c62aed53732e50473281886ba2b519cb08ee7b823e098fbe25d6e120d712c48.
//
// Solidity: event DistributeToCashback(address indexed rewardForCashback, uint256 ethAmount, uint256 usdcAmount, uint256 timestamp)
func (_FeeTreasury *FeeTreasuryFilterer) ParseDistributeToCashback(log types.Log) (*FeeTreasuryDistributeToCashback, error) {
	event := new(FeeTreasuryDistributeToCashback)
	if err := _FeeTreasury.contract.UnpackLog(event, "DistributeToCashback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryDistributeToStakingIterator is returned from FilterDistributeToStaking and is used to iterate over the raw logs and unpacked data for DistributeToStaking events raised by the FeeTreasury contract.
type FeeTreasuryDistributeToStakingIterator struct {
	Event *FeeTreasuryDistributeToStaking // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryDistributeToStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryDistributeToStaking)
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
		it.Event = new(FeeTreasuryDistributeToStaking)
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
func (it *FeeTreasuryDistributeToStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryDistributeToStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryDistributeToStaking represents a DistributeToStaking event raised by the FeeTreasury contract.
type FeeTreasuryDistributeToStaking struct {
	RewardForStaking common.Address
	EthAmount        *big.Int
	UsdcAmount       *big.Int
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeToStaking is a free log retrieval operation binding the contract event 0xf1f2c9fad9eecdb191575e1228d94c746321c25924daf61dca61e8e9932f2ca4.
//
// Solidity: event DistributeToStaking(address indexed rewardForStaking, uint256 ethAmount, uint256 usdcAmount, uint256 timestamp)
func (_FeeTreasury *FeeTreasuryFilterer) FilterDistributeToStaking(opts *bind.FilterOpts, rewardForStaking []common.Address) (*FeeTreasuryDistributeToStakingIterator, error) {

	var rewardForStakingRule []interface{}
	for _, rewardForStakingItem := range rewardForStaking {
		rewardForStakingRule = append(rewardForStakingRule, rewardForStakingItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "DistributeToStaking", rewardForStakingRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryDistributeToStakingIterator{contract: _FeeTreasury.contract, event: "DistributeToStaking", logs: logs, sub: sub}, nil
}

// WatchDistributeToStaking is a free log subscription operation binding the contract event 0xf1f2c9fad9eecdb191575e1228d94c746321c25924daf61dca61e8e9932f2ca4.
//
// Solidity: event DistributeToStaking(address indexed rewardForStaking, uint256 ethAmount, uint256 usdcAmount, uint256 timestamp)
func (_FeeTreasury *FeeTreasuryFilterer) WatchDistributeToStaking(opts *bind.WatchOpts, sink chan<- *FeeTreasuryDistributeToStaking, rewardForStaking []common.Address) (event.Subscription, error) {

	var rewardForStakingRule []interface{}
	for _, rewardForStakingItem := range rewardForStaking {
		rewardForStakingRule = append(rewardForStakingRule, rewardForStakingItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "DistributeToStaking", rewardForStakingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryDistributeToStaking)
				if err := _FeeTreasury.contract.UnpackLog(event, "DistributeToStaking", log); err != nil {
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

// ParseDistributeToStaking is a log parse operation binding the contract event 0xf1f2c9fad9eecdb191575e1228d94c746321c25924daf61dca61e8e9932f2ca4.
//
// Solidity: event DistributeToStaking(address indexed rewardForStaking, uint256 ethAmount, uint256 usdcAmount, uint256 timestamp)
func (_FeeTreasury *FeeTreasuryFilterer) ParseDistributeToStaking(log types.Log) (*FeeTreasuryDistributeToStaking, error) {
	event := new(FeeTreasuryDistributeToStaking)
	if err := _FeeTreasury.contract.UnpackLog(event, "DistributeToStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the FeeTreasury contract.
type FeeTreasuryNewOwnerIterator struct {
	Event *FeeTreasuryNewOwner // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryNewOwner)
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
		it.Event = new(FeeTreasuryNewOwner)
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
func (it *FeeTreasuryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryNewOwner represents a NewOwner event raised by the FeeTreasury contract.
type FeeTreasuryNewOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_FeeTreasury *FeeTreasuryFilterer) FilterNewOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*FeeTreasuryNewOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryNewOwnerIterator{contract: _FeeTreasury.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_FeeTreasury *FeeTreasuryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *FeeTreasuryNewOwner, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryNewOwner)
				if err := _FeeTreasury.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
func (_FeeTreasury *FeeTreasuryFilterer) ParseNewOwner(log types.Log) (*FeeTreasuryNewOwner, error) {
	event := new(FeeTreasuryNewOwner)
	if err := _FeeTreasury.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the FeeTreasury contract.
type FeeTreasuryNewPendingOwnerIterator struct {
	Event *FeeTreasuryNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryNewPendingOwner)
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
		it.Event = new(FeeTreasuryNewPendingOwner)
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
func (it *FeeTreasuryNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryNewPendingOwner represents a NewPendingOwner event raised by the FeeTreasury contract.
type FeeTreasuryNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_FeeTreasury *FeeTreasuryFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*FeeTreasuryNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryNewPendingOwnerIterator{contract: _FeeTreasury.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_FeeTreasury *FeeTreasuryFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *FeeTreasuryNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryNewPendingOwner)
				if err := _FeeTreasury.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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
func (_FeeTreasury *FeeTreasuryFilterer) ParseNewPendingOwner(log types.Log) (*FeeTreasuryNewPendingOwner, error) {
	event := new(FeeTreasuryNewPendingOwner)
	if err := _FeeTreasury.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the FeeTreasury contract.
type FeeTreasuryOperatorChangedIterator struct {
	Event *FeeTreasuryOperatorChanged // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryOperatorChanged)
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
		it.Event = new(FeeTreasuryOperatorChanged)
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
func (it *FeeTreasuryOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryOperatorChanged represents a OperatorChanged event raised by the FeeTreasury contract.
type FeeTreasuryOperatorChanged struct {
	OldOperator common.Address
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_FeeTreasury *FeeTreasuryFilterer) FilterOperatorChanged(opts *bind.FilterOpts, oldOperator []common.Address, newOperator []common.Address) (*FeeTreasuryOperatorChangedIterator, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "OperatorChanged", oldOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryOperatorChangedIterator{contract: _FeeTreasury.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_FeeTreasury *FeeTreasuryFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *FeeTreasuryOperatorChanged, oldOperator []common.Address, newOperator []common.Address) (event.Subscription, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "OperatorChanged", oldOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryOperatorChanged)
				if err := _FeeTreasury.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_FeeTreasury *FeeTreasuryFilterer) ParseOperatorChanged(log types.Log) (*FeeTreasuryOperatorChanged, error) {
	event := new(FeeTreasuryOperatorChanged)
	if err := _FeeTreasury.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryRatioForStakingChangedIterator is returned from FilterRatioForStakingChanged and is used to iterate over the raw logs and unpacked data for RatioForStakingChanged events raised by the FeeTreasury contract.
type FeeTreasuryRatioForStakingChangedIterator struct {
	Event *FeeTreasuryRatioForStakingChanged // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryRatioForStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryRatioForStakingChanged)
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
		it.Event = new(FeeTreasuryRatioForStakingChanged)
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
func (it *FeeTreasuryRatioForStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryRatioForStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryRatioForStakingChanged represents a RatioForStakingChanged event raised by the FeeTreasury contract.
type FeeTreasuryRatioForStakingChanged struct {
	OldRatio uint8
	NewRatio uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRatioForStakingChanged is a free log retrieval operation binding the contract event 0xc77f8dd2877f5ac909d42359b8c2dca5a7af08132dfde4b20187d48cf9c388de.
//
// Solidity: event RatioForStakingChanged(uint8 oldRatio, uint8 newRatio)
func (_FeeTreasury *FeeTreasuryFilterer) FilterRatioForStakingChanged(opts *bind.FilterOpts) (*FeeTreasuryRatioForStakingChangedIterator, error) {

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "RatioForStakingChanged")
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryRatioForStakingChangedIterator{contract: _FeeTreasury.contract, event: "RatioForStakingChanged", logs: logs, sub: sub}, nil
}

// WatchRatioForStakingChanged is a free log subscription operation binding the contract event 0xc77f8dd2877f5ac909d42359b8c2dca5a7af08132dfde4b20187d48cf9c388de.
//
// Solidity: event RatioForStakingChanged(uint8 oldRatio, uint8 newRatio)
func (_FeeTreasury *FeeTreasuryFilterer) WatchRatioForStakingChanged(opts *bind.WatchOpts, sink chan<- *FeeTreasuryRatioForStakingChanged) (event.Subscription, error) {

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "RatioForStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryRatioForStakingChanged)
				if err := _FeeTreasury.contract.UnpackLog(event, "RatioForStakingChanged", log); err != nil {
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

// ParseRatioForStakingChanged is a log parse operation binding the contract event 0xc77f8dd2877f5ac909d42359b8c2dca5a7af08132dfde4b20187d48cf9c388de.
//
// Solidity: event RatioForStakingChanged(uint8 oldRatio, uint8 newRatio)
func (_FeeTreasury *FeeTreasuryFilterer) ParseRatioForStakingChanged(log types.Log) (*FeeTreasuryRatioForStakingChanged, error) {
	event := new(FeeTreasuryRatioForStakingChanged)
	if err := _FeeTreasury.contract.UnpackLog(event, "RatioForStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryRewardForCashbackChangedIterator is returned from FilterRewardForCashbackChanged and is used to iterate over the raw logs and unpacked data for RewardForCashbackChanged events raised by the FeeTreasury contract.
type FeeTreasuryRewardForCashbackChangedIterator struct {
	Event *FeeTreasuryRewardForCashbackChanged // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryRewardForCashbackChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryRewardForCashbackChanged)
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
		it.Event = new(FeeTreasuryRewardForCashbackChanged)
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
func (it *FeeTreasuryRewardForCashbackChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryRewardForCashbackChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryRewardForCashbackChanged represents a RewardForCashbackChanged event raised by the FeeTreasury contract.
type FeeTreasuryRewardForCashbackChanged struct {
	OldReward common.Address
	NewReward common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardForCashbackChanged is a free log retrieval operation binding the contract event 0x2f2ee672c26f678335773f44e0597a0b9cd1999b65e5b8ac1cdead69316fd687.
//
// Solidity: event RewardForCashbackChanged(address indexed oldReward, address indexed newReward)
func (_FeeTreasury *FeeTreasuryFilterer) FilterRewardForCashbackChanged(opts *bind.FilterOpts, oldReward []common.Address, newReward []common.Address) (*FeeTreasuryRewardForCashbackChangedIterator, error) {

	var oldRewardRule []interface{}
	for _, oldRewardItem := range oldReward {
		oldRewardRule = append(oldRewardRule, oldRewardItem)
	}
	var newRewardRule []interface{}
	for _, newRewardItem := range newReward {
		newRewardRule = append(newRewardRule, newRewardItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "RewardForCashbackChanged", oldRewardRule, newRewardRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryRewardForCashbackChangedIterator{contract: _FeeTreasury.contract, event: "RewardForCashbackChanged", logs: logs, sub: sub}, nil
}

// WatchRewardForCashbackChanged is a free log subscription operation binding the contract event 0x2f2ee672c26f678335773f44e0597a0b9cd1999b65e5b8ac1cdead69316fd687.
//
// Solidity: event RewardForCashbackChanged(address indexed oldReward, address indexed newReward)
func (_FeeTreasury *FeeTreasuryFilterer) WatchRewardForCashbackChanged(opts *bind.WatchOpts, sink chan<- *FeeTreasuryRewardForCashbackChanged, oldReward []common.Address, newReward []common.Address) (event.Subscription, error) {

	var oldRewardRule []interface{}
	for _, oldRewardItem := range oldReward {
		oldRewardRule = append(oldRewardRule, oldRewardItem)
	}
	var newRewardRule []interface{}
	for _, newRewardItem := range newReward {
		newRewardRule = append(newRewardRule, newRewardItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "RewardForCashbackChanged", oldRewardRule, newRewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryRewardForCashbackChanged)
				if err := _FeeTreasury.contract.UnpackLog(event, "RewardForCashbackChanged", log); err != nil {
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

// ParseRewardForCashbackChanged is a log parse operation binding the contract event 0x2f2ee672c26f678335773f44e0597a0b9cd1999b65e5b8ac1cdead69316fd687.
//
// Solidity: event RewardForCashbackChanged(address indexed oldReward, address indexed newReward)
func (_FeeTreasury *FeeTreasuryFilterer) ParseRewardForCashbackChanged(log types.Log) (*FeeTreasuryRewardForCashbackChanged, error) {
	event := new(FeeTreasuryRewardForCashbackChanged)
	if err := _FeeTreasury.contract.UnpackLog(event, "RewardForCashbackChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasuryRewardForStakingChangedIterator is returned from FilterRewardForStakingChanged and is used to iterate over the raw logs and unpacked data for RewardForStakingChanged events raised by the FeeTreasury contract.
type FeeTreasuryRewardForStakingChangedIterator struct {
	Event *FeeTreasuryRewardForStakingChanged // Event containing the contract specifics and raw log

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
func (it *FeeTreasuryRewardForStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasuryRewardForStakingChanged)
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
		it.Event = new(FeeTreasuryRewardForStakingChanged)
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
func (it *FeeTreasuryRewardForStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasuryRewardForStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasuryRewardForStakingChanged represents a RewardForStakingChanged event raised by the FeeTreasury contract.
type FeeTreasuryRewardForStakingChanged struct {
	OldReward common.Address
	NewReward common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardForStakingChanged is a free log retrieval operation binding the contract event 0x2241d917e510bcd2b27506cc60a28cb86047ec1018cb6c6a51e0dc8a7f9ba861.
//
// Solidity: event RewardForStakingChanged(address indexed oldReward, address indexed newReward)
func (_FeeTreasury *FeeTreasuryFilterer) FilterRewardForStakingChanged(opts *bind.FilterOpts, oldReward []common.Address, newReward []common.Address) (*FeeTreasuryRewardForStakingChangedIterator, error) {

	var oldRewardRule []interface{}
	for _, oldRewardItem := range oldReward {
		oldRewardRule = append(oldRewardRule, oldRewardItem)
	}
	var newRewardRule []interface{}
	for _, newRewardItem := range newReward {
		newRewardRule = append(newRewardRule, newRewardItem)
	}

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "RewardForStakingChanged", oldRewardRule, newRewardRule)
	if err != nil {
		return nil, err
	}
	return &FeeTreasuryRewardForStakingChangedIterator{contract: _FeeTreasury.contract, event: "RewardForStakingChanged", logs: logs, sub: sub}, nil
}

// WatchRewardForStakingChanged is a free log subscription operation binding the contract event 0x2241d917e510bcd2b27506cc60a28cb86047ec1018cb6c6a51e0dc8a7f9ba861.
//
// Solidity: event RewardForStakingChanged(address indexed oldReward, address indexed newReward)
func (_FeeTreasury *FeeTreasuryFilterer) WatchRewardForStakingChanged(opts *bind.WatchOpts, sink chan<- *FeeTreasuryRewardForStakingChanged, oldReward []common.Address, newReward []common.Address) (event.Subscription, error) {

	var oldRewardRule []interface{}
	for _, oldRewardItem := range oldReward {
		oldRewardRule = append(oldRewardRule, oldRewardItem)
	}
	var newRewardRule []interface{}
	for _, newRewardItem := range newReward {
		newRewardRule = append(newRewardRule, newRewardItem)
	}

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "RewardForStakingChanged", oldRewardRule, newRewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasuryRewardForStakingChanged)
				if err := _FeeTreasury.contract.UnpackLog(event, "RewardForStakingChanged", log); err != nil {
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

// ParseRewardForStakingChanged is a log parse operation binding the contract event 0x2241d917e510bcd2b27506cc60a28cb86047ec1018cb6c6a51e0dc8a7f9ba861.
//
// Solidity: event RewardForStakingChanged(address indexed oldReward, address indexed newReward)
func (_FeeTreasury *FeeTreasuryFilterer) ParseRewardForStakingChanged(log types.Log) (*FeeTreasuryRewardForStakingChanged, error) {
	event := new(FeeTreasuryRewardForStakingChanged)
	if err := _FeeTreasury.contract.UnpackLog(event, "RewardForStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeTreasurySettlementIntervalChangedIterator is returned from FilterSettlementIntervalChanged and is used to iterate over the raw logs and unpacked data for SettlementIntervalChanged events raised by the FeeTreasury contract.
type FeeTreasurySettlementIntervalChangedIterator struct {
	Event *FeeTreasurySettlementIntervalChanged // Event containing the contract specifics and raw log

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
func (it *FeeTreasurySettlementIntervalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTreasurySettlementIntervalChanged)
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
		it.Event = new(FeeTreasurySettlementIntervalChanged)
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
func (it *FeeTreasurySettlementIntervalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTreasurySettlementIntervalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTreasurySettlementIntervalChanged represents a SettlementIntervalChanged event raised by the FeeTreasury contract.
type FeeTreasurySettlementIntervalChanged struct {
	OldInterval *big.Int
	NewInterval *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettlementIntervalChanged is a free log retrieval operation binding the contract event 0xa16e1dddb457564fbe13a889e5651ee325052b009c5524fa3f5b6ebdd7a3581a.
//
// Solidity: event SettlementIntervalChanged(uint256 oldInterval, uint256 newInterval)
func (_FeeTreasury *FeeTreasuryFilterer) FilterSettlementIntervalChanged(opts *bind.FilterOpts) (*FeeTreasurySettlementIntervalChangedIterator, error) {

	logs, sub, err := _FeeTreasury.contract.FilterLogs(opts, "SettlementIntervalChanged")
	if err != nil {
		return nil, err
	}
	return &FeeTreasurySettlementIntervalChangedIterator{contract: _FeeTreasury.contract, event: "SettlementIntervalChanged", logs: logs, sub: sub}, nil
}

// WatchSettlementIntervalChanged is a free log subscription operation binding the contract event 0xa16e1dddb457564fbe13a889e5651ee325052b009c5524fa3f5b6ebdd7a3581a.
//
// Solidity: event SettlementIntervalChanged(uint256 oldInterval, uint256 newInterval)
func (_FeeTreasury *FeeTreasuryFilterer) WatchSettlementIntervalChanged(opts *bind.WatchOpts, sink chan<- *FeeTreasurySettlementIntervalChanged) (event.Subscription, error) {

	logs, sub, err := _FeeTreasury.contract.WatchLogs(opts, "SettlementIntervalChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTreasurySettlementIntervalChanged)
				if err := _FeeTreasury.contract.UnpackLog(event, "SettlementIntervalChanged", log); err != nil {
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

// ParseSettlementIntervalChanged is a log parse operation binding the contract event 0xa16e1dddb457564fbe13a889e5651ee325052b009c5524fa3f5b6ebdd7a3581a.
//
// Solidity: event SettlementIntervalChanged(uint256 oldInterval, uint256 newInterval)
func (_FeeTreasury *FeeTreasuryFilterer) ParseSettlementIntervalChanged(log types.Log) (*FeeTreasurySettlementIntervalChanged, error) {
	event := new(FeeTreasurySettlementIntervalChanged)
	if err := _FeeTreasury.contract.UnpackLog(event, "SettlementIntervalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
