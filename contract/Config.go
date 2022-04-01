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

// ConfigMetaData contains all meta data concerning the Config contract.
var ConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"PriceOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInterval\",\"type\":\"uint256\"}],\"name\":\"RebaseIntervalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGap\",\"type\":\"uint256\"}],\"name\":\"RebasePriceGapChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBeta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beta\",\"type\":\"uint256\"}],\"name\":\"SetBeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"SetEmergency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeParameter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeParameter\",\"type\":\"uint256\"}],\"name\":\"SetFeeParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldInitMarginRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initMarginRatio\",\"type\":\"uint256\"}],\"name\":\"SetInitMarginRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidateFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidateFeeRatio\",\"type\":\"uint256\"}],\"name\":\"SetLiquidateFeeRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidateThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidateThreshold\",\"type\":\"uint256\"}],\"name\":\"SetLiquidateThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLpWithdrawThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpWithdrawThreshold\",\"type\":\"uint256\"}],\"name\":\"SetLpWithdrawThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxCPFBoost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxCPFBoost\",\"type\":\"uint256\"}],\"name\":\"SetMaxCPFBoost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTradingSlippage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTradingSlippage\",\"type\":\"uint256\"}],\"name\":\"TradingSlippageChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beta\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeParameter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inEmergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initMarginRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidateFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidateThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpWithdrawThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCPFBoost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebasePriceGap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"registerRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"routerMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newBeta\",\"type\":\"uint8\"}],\"name\":\"setBeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeParameter\",\"type\":\"uint256\"}],\"name\":\"setFeeParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marginRatio\",\"type\":\"uint256\"}],\"name\":\"setInitMarginRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeRatio\",\"type\":\"uint256\"}],\"name\":\"setLiquidateFeeRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setLiquidateThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLpWithdrawThreshold\",\"type\":\"uint256\"}],\"name\":\"setLpWithdrawThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxCPFBoost\",\"type\":\"uint256\"}],\"name\":\"setMaxCPFBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"setRebaseInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGap\",\"type\":\"uint256\"}],\"name\":\"setRebasePriceGap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTradingSlippage\",\"type\":\"uint256\"}],\"name\":\"setTradingSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingSlippage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"unregisterRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfigMetaData.ABI instead.
var ConfigABI = ConfigMetaData.ABI

// Config is an auto generated Go binding around an Ethereum contract.
type Config struct {
	ConfigCaller     // Read-only binding to the contract
	ConfigTransactor // Write-only binding to the contract
	ConfigFilterer   // Log filterer for contract events
}

// ConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigSession struct {
	Contract     *Config           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigCallerSession struct {
	Contract *ConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigTransactorSession struct {
	Contract     *ConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigRaw struct {
	Contract *Config // Generic contract binding to access the raw methods on
}

// ConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigCallerRaw struct {
	Contract *ConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigTransactorRaw struct {
	Contract *ConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfig creates a new instance of Config, bound to a specific deployed contract.
func NewConfig(address common.Address, backend bind.ContractBackend) (*Config, error) {
	contract, err := bindConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// NewConfigCaller creates a new read-only instance of Config, bound to a specific deployed contract.
func NewConfigCaller(address common.Address, caller bind.ContractCaller) (*ConfigCaller, error) {
	contract, err := bindConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigCaller{contract: contract}, nil
}

// NewConfigTransactor creates a new write-only instance of Config, bound to a specific deployed contract.
func NewConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigTransactor, error) {
	contract, err := bindConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigTransactor{contract: contract}, nil
}

// NewConfigFilterer creates a new log filterer instance of Config, bound to a specific deployed contract.
func NewConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigFilterer, error) {
	contract, err := bindConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigFilterer{contract: contract}, nil
}

// bindConfig binds a generic wrapper to an already deployed contract.
func bindConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.ConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.contract.Transact(opts, method, params...)
}

// Beta is a free data retrieval call binding the contract method 0x9faa3c91.
//
// Solidity: function beta() view returns(uint8)
func (_Config *ConfigCaller) Beta(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "beta")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Beta is a free data retrieval call binding the contract method 0x9faa3c91.
//
// Solidity: function beta() view returns(uint8)
func (_Config *ConfigSession) Beta() (uint8, error) {
	return _Config.Contract.Beta(&_Config.CallOpts)
}

// Beta is a free data retrieval call binding the contract method 0x9faa3c91.
//
// Solidity: function beta() view returns(uint8)
func (_Config *ConfigCallerSession) Beta() (uint8, error) {
	return _Config.Contract.Beta(&_Config.CallOpts)
}

// FeeParameter is a free data retrieval call binding the contract method 0xb85f9f0b.
//
// Solidity: function feeParameter() view returns(uint256)
func (_Config *ConfigCaller) FeeParameter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "feeParameter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeParameter is a free data retrieval call binding the contract method 0xb85f9f0b.
//
// Solidity: function feeParameter() view returns(uint256)
func (_Config *ConfigSession) FeeParameter() (*big.Int, error) {
	return _Config.Contract.FeeParameter(&_Config.CallOpts)
}

// FeeParameter is a free data retrieval call binding the contract method 0xb85f9f0b.
//
// Solidity: function feeParameter() view returns(uint256)
func (_Config *ConfigCallerSession) FeeParameter() (*big.Int, error) {
	return _Config.Contract.FeeParameter(&_Config.CallOpts)
}

// InEmergency is a free data retrieval call binding the contract method 0xec7b2151.
//
// Solidity: function inEmergency(address ) view returns(bool)
func (_Config *ConfigCaller) InEmergency(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "inEmergency", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InEmergency is a free data retrieval call binding the contract method 0xec7b2151.
//
// Solidity: function inEmergency(address ) view returns(bool)
func (_Config *ConfigSession) InEmergency(arg0 common.Address) (bool, error) {
	return _Config.Contract.InEmergency(&_Config.CallOpts, arg0)
}

// InEmergency is a free data retrieval call binding the contract method 0xec7b2151.
//
// Solidity: function inEmergency(address ) view returns(bool)
func (_Config *ConfigCallerSession) InEmergency(arg0 common.Address) (bool, error) {
	return _Config.Contract.InEmergency(&_Config.CallOpts, arg0)
}

// InitMarginRatio is a free data retrieval call binding the contract method 0xe57d5636.
//
// Solidity: function initMarginRatio() view returns(uint256)
func (_Config *ConfigCaller) InitMarginRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "initMarginRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitMarginRatio is a free data retrieval call binding the contract method 0xe57d5636.
//
// Solidity: function initMarginRatio() view returns(uint256)
func (_Config *ConfigSession) InitMarginRatio() (*big.Int, error) {
	return _Config.Contract.InitMarginRatio(&_Config.CallOpts)
}

// InitMarginRatio is a free data retrieval call binding the contract method 0xe57d5636.
//
// Solidity: function initMarginRatio() view returns(uint256)
func (_Config *ConfigCallerSession) InitMarginRatio() (*big.Int, error) {
	return _Config.Contract.InitMarginRatio(&_Config.CallOpts)
}

// LiquidateFeeRatio is a free data retrieval call binding the contract method 0x6dd26a07.
//
// Solidity: function liquidateFeeRatio() view returns(uint256)
func (_Config *ConfigCaller) LiquidateFeeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "liquidateFeeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidateFeeRatio is a free data retrieval call binding the contract method 0x6dd26a07.
//
// Solidity: function liquidateFeeRatio() view returns(uint256)
func (_Config *ConfigSession) LiquidateFeeRatio() (*big.Int, error) {
	return _Config.Contract.LiquidateFeeRatio(&_Config.CallOpts)
}

// LiquidateFeeRatio is a free data retrieval call binding the contract method 0x6dd26a07.
//
// Solidity: function liquidateFeeRatio() view returns(uint256)
func (_Config *ConfigCallerSession) LiquidateFeeRatio() (*big.Int, error) {
	return _Config.Contract.LiquidateFeeRatio(&_Config.CallOpts)
}

// LiquidateThreshold is a free data retrieval call binding the contract method 0xfdcb648c.
//
// Solidity: function liquidateThreshold() view returns(uint256)
func (_Config *ConfigCaller) LiquidateThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "liquidateThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidateThreshold is a free data retrieval call binding the contract method 0xfdcb648c.
//
// Solidity: function liquidateThreshold() view returns(uint256)
func (_Config *ConfigSession) LiquidateThreshold() (*big.Int, error) {
	return _Config.Contract.LiquidateThreshold(&_Config.CallOpts)
}

// LiquidateThreshold is a free data retrieval call binding the contract method 0xfdcb648c.
//
// Solidity: function liquidateThreshold() view returns(uint256)
func (_Config *ConfigCallerSession) LiquidateThreshold() (*big.Int, error) {
	return _Config.Contract.LiquidateThreshold(&_Config.CallOpts)
}

// LpWithdrawThreshold is a free data retrieval call binding the contract method 0xa8020418.
//
// Solidity: function lpWithdrawThreshold() view returns(uint256)
func (_Config *ConfigCaller) LpWithdrawThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "lpWithdrawThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpWithdrawThreshold is a free data retrieval call binding the contract method 0xa8020418.
//
// Solidity: function lpWithdrawThreshold() view returns(uint256)
func (_Config *ConfigSession) LpWithdrawThreshold() (*big.Int, error) {
	return _Config.Contract.LpWithdrawThreshold(&_Config.CallOpts)
}

// LpWithdrawThreshold is a free data retrieval call binding the contract method 0xa8020418.
//
// Solidity: function lpWithdrawThreshold() view returns(uint256)
func (_Config *ConfigCallerSession) LpWithdrawThreshold() (*big.Int, error) {
	return _Config.Contract.LpWithdrawThreshold(&_Config.CallOpts)
}

// MaxCPFBoost is a free data retrieval call binding the contract method 0x83866469.
//
// Solidity: function maxCPFBoost() view returns(uint256)
func (_Config *ConfigCaller) MaxCPFBoost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "maxCPFBoost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCPFBoost is a free data retrieval call binding the contract method 0x83866469.
//
// Solidity: function maxCPFBoost() view returns(uint256)
func (_Config *ConfigSession) MaxCPFBoost() (*big.Int, error) {
	return _Config.Contract.MaxCPFBoost(&_Config.CallOpts)
}

// MaxCPFBoost is a free data retrieval call binding the contract method 0x83866469.
//
// Solidity: function maxCPFBoost() view returns(uint256)
func (_Config *ConfigCallerSession) MaxCPFBoost() (*big.Int, error) {
	return _Config.Contract.MaxCPFBoost(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigCallerSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Config *ConfigCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Config *ConfigSession) PendingOwner() (common.Address, error) {
	return _Config.Contract.PendingOwner(&_Config.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Config *ConfigCallerSession) PendingOwner() (common.Address, error) {
	return _Config.Contract.PendingOwner(&_Config.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Config *ConfigCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Config *ConfigSession) PriceOracle() (common.Address, error) {
	return _Config.Contract.PriceOracle(&_Config.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Config *ConfigCallerSession) PriceOracle() (common.Address, error) {
	return _Config.Contract.PriceOracle(&_Config.CallOpts)
}

// RebaseInterval is a free data retrieval call binding the contract method 0x89edeb74.
//
// Solidity: function rebaseInterval() view returns(uint256)
func (_Config *ConfigCaller) RebaseInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "rebaseInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseInterval is a free data retrieval call binding the contract method 0x89edeb74.
//
// Solidity: function rebaseInterval() view returns(uint256)
func (_Config *ConfigSession) RebaseInterval() (*big.Int, error) {
	return _Config.Contract.RebaseInterval(&_Config.CallOpts)
}

// RebaseInterval is a free data retrieval call binding the contract method 0x89edeb74.
//
// Solidity: function rebaseInterval() view returns(uint256)
func (_Config *ConfigCallerSession) RebaseInterval() (*big.Int, error) {
	return _Config.Contract.RebaseInterval(&_Config.CallOpts)
}

// RebasePriceGap is a free data retrieval call binding the contract method 0x8ecb4551.
//
// Solidity: function rebasePriceGap() view returns(uint256)
func (_Config *ConfigCaller) RebasePriceGap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "rebasePriceGap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebasePriceGap is a free data retrieval call binding the contract method 0x8ecb4551.
//
// Solidity: function rebasePriceGap() view returns(uint256)
func (_Config *ConfigSession) RebasePriceGap() (*big.Int, error) {
	return _Config.Contract.RebasePriceGap(&_Config.CallOpts)
}

// RebasePriceGap is a free data retrieval call binding the contract method 0x8ecb4551.
//
// Solidity: function rebasePriceGap() view returns(uint256)
func (_Config *ConfigCallerSession) RebasePriceGap() (*big.Int, error) {
	return _Config.Contract.RebasePriceGap(&_Config.CallOpts)
}

// RouterMap is a free data retrieval call binding the contract method 0x9518332e.
//
// Solidity: function routerMap(address ) view returns(bool)
func (_Config *ConfigCaller) RouterMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "routerMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RouterMap is a free data retrieval call binding the contract method 0x9518332e.
//
// Solidity: function routerMap(address ) view returns(bool)
func (_Config *ConfigSession) RouterMap(arg0 common.Address) (bool, error) {
	return _Config.Contract.RouterMap(&_Config.CallOpts, arg0)
}

// RouterMap is a free data retrieval call binding the contract method 0x9518332e.
//
// Solidity: function routerMap(address ) view returns(bool)
func (_Config *ConfigCallerSession) RouterMap(arg0 common.Address) (bool, error) {
	return _Config.Contract.RouterMap(&_Config.CallOpts, arg0)
}

// TradingSlippage is a free data retrieval call binding the contract method 0x8ff471f7.
//
// Solidity: function tradingSlippage() view returns(uint256)
func (_Config *ConfigCaller) TradingSlippage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "tradingSlippage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradingSlippage is a free data retrieval call binding the contract method 0x8ff471f7.
//
// Solidity: function tradingSlippage() view returns(uint256)
func (_Config *ConfigSession) TradingSlippage() (*big.Int, error) {
	return _Config.Contract.TradingSlippage(&_Config.CallOpts)
}

// TradingSlippage is a free data retrieval call binding the contract method 0x8ff471f7.
//
// Solidity: function tradingSlippage() view returns(uint256)
func (_Config *ConfigCallerSession) TradingSlippage() (*big.Int, error) {
	return _Config.Contract.TradingSlippage(&_Config.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Config *ConfigTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Config *ConfigSession) AcceptOwner() (*types.Transaction, error) {
	return _Config.Contract.AcceptOwner(&_Config.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Config *ConfigTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _Config.Contract.AcceptOwner(&_Config.TransactOpts)
}

// RegisterRouter is a paid mutator transaction binding the contract method 0x22b19af7.
//
// Solidity: function registerRouter(address router) returns()
func (_Config *ConfigTransactor) RegisterRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "registerRouter", router)
}

// RegisterRouter is a paid mutator transaction binding the contract method 0x22b19af7.
//
// Solidity: function registerRouter(address router) returns()
func (_Config *ConfigSession) RegisterRouter(router common.Address) (*types.Transaction, error) {
	return _Config.Contract.RegisterRouter(&_Config.TransactOpts, router)
}

// RegisterRouter is a paid mutator transaction binding the contract method 0x22b19af7.
//
// Solidity: function registerRouter(address router) returns()
func (_Config *ConfigTransactorSession) RegisterRouter(router common.Address) (*types.Transaction, error) {
	return _Config.Contract.RegisterRouter(&_Config.TransactOpts, router)
}

// SetBeta is a paid mutator transaction binding the contract method 0x71fdd72d.
//
// Solidity: function setBeta(uint8 newBeta) returns()
func (_Config *ConfigTransactor) SetBeta(opts *bind.TransactOpts, newBeta uint8) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setBeta", newBeta)
}

// SetBeta is a paid mutator transaction binding the contract method 0x71fdd72d.
//
// Solidity: function setBeta(uint8 newBeta) returns()
func (_Config *ConfigSession) SetBeta(newBeta uint8) (*types.Transaction, error) {
	return _Config.Contract.SetBeta(&_Config.TransactOpts, newBeta)
}

// SetBeta is a paid mutator transaction binding the contract method 0x71fdd72d.
//
// Solidity: function setBeta(uint8 newBeta) returns()
func (_Config *ConfigTransactorSession) SetBeta(newBeta uint8) (*types.Transaction, error) {
	return _Config.Contract.SetBeta(&_Config.TransactOpts, newBeta)
}

// SetEmergency is a paid mutator transaction binding the contract method 0xeb02a115.
//
// Solidity: function setEmergency(address router) returns()
func (_Config *ConfigTransactor) SetEmergency(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setEmergency", router)
}

// SetEmergency is a paid mutator transaction binding the contract method 0xeb02a115.
//
// Solidity: function setEmergency(address router) returns()
func (_Config *ConfigSession) SetEmergency(router common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetEmergency(&_Config.TransactOpts, router)
}

// SetEmergency is a paid mutator transaction binding the contract method 0xeb02a115.
//
// Solidity: function setEmergency(address router) returns()
func (_Config *ConfigTransactorSession) SetEmergency(router common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetEmergency(&_Config.TransactOpts, router)
}

// SetFeeParameter is a paid mutator transaction binding the contract method 0x4b225189.
//
// Solidity: function setFeeParameter(uint256 newFeeParameter) returns()
func (_Config *ConfigTransactor) SetFeeParameter(opts *bind.TransactOpts, newFeeParameter *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setFeeParameter", newFeeParameter)
}

// SetFeeParameter is a paid mutator transaction binding the contract method 0x4b225189.
//
// Solidity: function setFeeParameter(uint256 newFeeParameter) returns()
func (_Config *ConfigSession) SetFeeParameter(newFeeParameter *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetFeeParameter(&_Config.TransactOpts, newFeeParameter)
}

// SetFeeParameter is a paid mutator transaction binding the contract method 0x4b225189.
//
// Solidity: function setFeeParameter(uint256 newFeeParameter) returns()
func (_Config *ConfigTransactorSession) SetFeeParameter(newFeeParameter *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetFeeParameter(&_Config.TransactOpts, newFeeParameter)
}

// SetInitMarginRatio is a paid mutator transaction binding the contract method 0xfbede224.
//
// Solidity: function setInitMarginRatio(uint256 marginRatio) returns()
func (_Config *ConfigTransactor) SetInitMarginRatio(opts *bind.TransactOpts, marginRatio *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setInitMarginRatio", marginRatio)
}

// SetInitMarginRatio is a paid mutator transaction binding the contract method 0xfbede224.
//
// Solidity: function setInitMarginRatio(uint256 marginRatio) returns()
func (_Config *ConfigSession) SetInitMarginRatio(marginRatio *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetInitMarginRatio(&_Config.TransactOpts, marginRatio)
}

// SetInitMarginRatio is a paid mutator transaction binding the contract method 0xfbede224.
//
// Solidity: function setInitMarginRatio(uint256 marginRatio) returns()
func (_Config *ConfigTransactorSession) SetInitMarginRatio(marginRatio *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetInitMarginRatio(&_Config.TransactOpts, marginRatio)
}

// SetLiquidateFeeRatio is a paid mutator transaction binding the contract method 0x711a8f52.
//
// Solidity: function setLiquidateFeeRatio(uint256 feeRatio) returns()
func (_Config *ConfigTransactor) SetLiquidateFeeRatio(opts *bind.TransactOpts, feeRatio *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setLiquidateFeeRatio", feeRatio)
}

// SetLiquidateFeeRatio is a paid mutator transaction binding the contract method 0x711a8f52.
//
// Solidity: function setLiquidateFeeRatio(uint256 feeRatio) returns()
func (_Config *ConfigSession) SetLiquidateFeeRatio(feeRatio *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetLiquidateFeeRatio(&_Config.TransactOpts, feeRatio)
}

// SetLiquidateFeeRatio is a paid mutator transaction binding the contract method 0x711a8f52.
//
// Solidity: function setLiquidateFeeRatio(uint256 feeRatio) returns()
func (_Config *ConfigTransactorSession) SetLiquidateFeeRatio(feeRatio *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetLiquidateFeeRatio(&_Config.TransactOpts, feeRatio)
}

// SetLiquidateThreshold is a paid mutator transaction binding the contract method 0x3ffd2ee2.
//
// Solidity: function setLiquidateThreshold(uint256 threshold) returns()
func (_Config *ConfigTransactor) SetLiquidateThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setLiquidateThreshold", threshold)
}

// SetLiquidateThreshold is a paid mutator transaction binding the contract method 0x3ffd2ee2.
//
// Solidity: function setLiquidateThreshold(uint256 threshold) returns()
func (_Config *ConfigSession) SetLiquidateThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetLiquidateThreshold(&_Config.TransactOpts, threshold)
}

// SetLiquidateThreshold is a paid mutator transaction binding the contract method 0x3ffd2ee2.
//
// Solidity: function setLiquidateThreshold(uint256 threshold) returns()
func (_Config *ConfigTransactorSession) SetLiquidateThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetLiquidateThreshold(&_Config.TransactOpts, threshold)
}

// SetLpWithdrawThreshold is a paid mutator transaction binding the contract method 0x7f48795e.
//
// Solidity: function setLpWithdrawThreshold(uint256 newLpWithdrawThreshold) returns()
func (_Config *ConfigTransactor) SetLpWithdrawThreshold(opts *bind.TransactOpts, newLpWithdrawThreshold *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setLpWithdrawThreshold", newLpWithdrawThreshold)
}

// SetLpWithdrawThreshold is a paid mutator transaction binding the contract method 0x7f48795e.
//
// Solidity: function setLpWithdrawThreshold(uint256 newLpWithdrawThreshold) returns()
func (_Config *ConfigSession) SetLpWithdrawThreshold(newLpWithdrawThreshold *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetLpWithdrawThreshold(&_Config.TransactOpts, newLpWithdrawThreshold)
}

// SetLpWithdrawThreshold is a paid mutator transaction binding the contract method 0x7f48795e.
//
// Solidity: function setLpWithdrawThreshold(uint256 newLpWithdrawThreshold) returns()
func (_Config *ConfigTransactorSession) SetLpWithdrawThreshold(newLpWithdrawThreshold *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetLpWithdrawThreshold(&_Config.TransactOpts, newLpWithdrawThreshold)
}

// SetMaxCPFBoost is a paid mutator transaction binding the contract method 0xd73e732e.
//
// Solidity: function setMaxCPFBoost(uint256 newMaxCPFBoost) returns()
func (_Config *ConfigTransactor) SetMaxCPFBoost(opts *bind.TransactOpts, newMaxCPFBoost *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setMaxCPFBoost", newMaxCPFBoost)
}

// SetMaxCPFBoost is a paid mutator transaction binding the contract method 0xd73e732e.
//
// Solidity: function setMaxCPFBoost(uint256 newMaxCPFBoost) returns()
func (_Config *ConfigSession) SetMaxCPFBoost(newMaxCPFBoost *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetMaxCPFBoost(&_Config.TransactOpts, newMaxCPFBoost)
}

// SetMaxCPFBoost is a paid mutator transaction binding the contract method 0xd73e732e.
//
// Solidity: function setMaxCPFBoost(uint256 newMaxCPFBoost) returns()
func (_Config *ConfigTransactorSession) SetMaxCPFBoost(newMaxCPFBoost *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetMaxCPFBoost(&_Config.TransactOpts, newMaxCPFBoost)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_Config *ConfigTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_Config *ConfigSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetPendingOwner(&_Config.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_Config *ConfigTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetPendingOwner(&_Config.TransactOpts, newPendingOwner)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newOracle) returns()
func (_Config *ConfigTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newOracle) returns()
func (_Config *ConfigSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetPriceOracle(&_Config.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newOracle) returns()
func (_Config *ConfigTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetPriceOracle(&_Config.TransactOpts, newOracle)
}

// SetRebaseInterval is a paid mutator transaction binding the contract method 0xef037fb9.
//
// Solidity: function setRebaseInterval(uint256 interval) returns()
func (_Config *ConfigTransactor) SetRebaseInterval(opts *bind.TransactOpts, interval *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setRebaseInterval", interval)
}

// SetRebaseInterval is a paid mutator transaction binding the contract method 0xef037fb9.
//
// Solidity: function setRebaseInterval(uint256 interval) returns()
func (_Config *ConfigSession) SetRebaseInterval(interval *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetRebaseInterval(&_Config.TransactOpts, interval)
}

// SetRebaseInterval is a paid mutator transaction binding the contract method 0xef037fb9.
//
// Solidity: function setRebaseInterval(uint256 interval) returns()
func (_Config *ConfigTransactorSession) SetRebaseInterval(interval *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetRebaseInterval(&_Config.TransactOpts, interval)
}

// SetRebasePriceGap is a paid mutator transaction binding the contract method 0x8392897d.
//
// Solidity: function setRebasePriceGap(uint256 newGap) returns()
func (_Config *ConfigTransactor) SetRebasePriceGap(opts *bind.TransactOpts, newGap *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setRebasePriceGap", newGap)
}

// SetRebasePriceGap is a paid mutator transaction binding the contract method 0x8392897d.
//
// Solidity: function setRebasePriceGap(uint256 newGap) returns()
func (_Config *ConfigSession) SetRebasePriceGap(newGap *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetRebasePriceGap(&_Config.TransactOpts, newGap)
}

// SetRebasePriceGap is a paid mutator transaction binding the contract method 0x8392897d.
//
// Solidity: function setRebasePriceGap(uint256 newGap) returns()
func (_Config *ConfigTransactorSession) SetRebasePriceGap(newGap *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetRebasePriceGap(&_Config.TransactOpts, newGap)
}

// SetTradingSlippage is a paid mutator transaction binding the contract method 0x683b41ef.
//
// Solidity: function setTradingSlippage(uint256 newTradingSlippage) returns()
func (_Config *ConfigTransactor) SetTradingSlippage(opts *bind.TransactOpts, newTradingSlippage *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setTradingSlippage", newTradingSlippage)
}

// SetTradingSlippage is a paid mutator transaction binding the contract method 0x683b41ef.
//
// Solidity: function setTradingSlippage(uint256 newTradingSlippage) returns()
func (_Config *ConfigSession) SetTradingSlippage(newTradingSlippage *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetTradingSlippage(&_Config.TransactOpts, newTradingSlippage)
}

// SetTradingSlippage is a paid mutator transaction binding the contract method 0x683b41ef.
//
// Solidity: function setTradingSlippage(uint256 newTradingSlippage) returns()
func (_Config *ConfigTransactorSession) SetTradingSlippage(newTradingSlippage *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetTradingSlippage(&_Config.TransactOpts, newTradingSlippage)
}

// UnregisterRouter is a paid mutator transaction binding the contract method 0x5012e0a4.
//
// Solidity: function unregisterRouter(address router) returns()
func (_Config *ConfigTransactor) UnregisterRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "unregisterRouter", router)
}

// UnregisterRouter is a paid mutator transaction binding the contract method 0x5012e0a4.
//
// Solidity: function unregisterRouter(address router) returns()
func (_Config *ConfigSession) UnregisterRouter(router common.Address) (*types.Transaction, error) {
	return _Config.Contract.UnregisterRouter(&_Config.TransactOpts, router)
}

// UnregisterRouter is a paid mutator transaction binding the contract method 0x5012e0a4.
//
// Solidity: function unregisterRouter(address router) returns()
func (_Config *ConfigTransactorSession) UnregisterRouter(router common.Address) (*types.Transaction, error) {
	return _Config.Contract.UnregisterRouter(&_Config.TransactOpts, router)
}

// ConfigNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Config contract.
type ConfigNewOwnerIterator struct {
	Event *ConfigNewOwner // Event containing the contract specifics and raw log

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
func (it *ConfigNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigNewOwner)
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
		it.Event = new(ConfigNewOwner)
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
func (it *ConfigNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigNewOwner represents a NewOwner event raised by the Config contract.
type ConfigNewOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_Config *ConfigFilterer) FilterNewOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ConfigNewOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigNewOwnerIterator{contract: _Config.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_Config *ConfigFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ConfigNewOwner, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigNewOwner)
				if err := _Config.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
func (_Config *ConfigFilterer) ParseNewOwner(log types.Log) (*ConfigNewOwner, error) {
	event := new(ConfigNewOwner)
	if err := _Config.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the Config contract.
type ConfigNewPendingOwnerIterator struct {
	Event *ConfigNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *ConfigNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigNewPendingOwner)
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
		it.Event = new(ConfigNewPendingOwner)
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
func (it *ConfigNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigNewPendingOwner represents a NewPendingOwner event raised by the Config contract.
type ConfigNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Config *ConfigFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*ConfigNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigNewPendingOwnerIterator{contract: _Config.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Config *ConfigFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *ConfigNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigNewPendingOwner)
				if err := _Config.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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
func (_Config *ConfigFilterer) ParseNewPendingOwner(log types.Log) (*ConfigNewPendingOwner, error) {
	event := new(ConfigNewPendingOwner)
	if err := _Config.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigPriceOracleChangedIterator is returned from FilterPriceOracleChanged and is used to iterate over the raw logs and unpacked data for PriceOracleChanged events raised by the Config contract.
type ConfigPriceOracleChangedIterator struct {
	Event *ConfigPriceOracleChanged // Event containing the contract specifics and raw log

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
func (it *ConfigPriceOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigPriceOracleChanged)
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
		it.Event = new(ConfigPriceOracleChanged)
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
func (it *ConfigPriceOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigPriceOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigPriceOracleChanged represents a PriceOracleChanged event raised by the Config contract.
type ConfigPriceOracleChanged struct {
	OldOracle common.Address
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleChanged is a free log retrieval operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed oldOracle, address indexed newOracle)
func (_Config *ConfigFilterer) FilterPriceOracleChanged(opts *bind.FilterOpts, oldOracle []common.Address, newOracle []common.Address) (*ConfigPriceOracleChangedIterator, error) {

	var oldOracleRule []interface{}
	for _, oldOracleItem := range oldOracle {
		oldOracleRule = append(oldOracleRule, oldOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "PriceOracleChanged", oldOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return &ConfigPriceOracleChangedIterator{contract: _Config.contract, event: "PriceOracleChanged", logs: logs, sub: sub}, nil
}

// WatchPriceOracleChanged is a free log subscription operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed oldOracle, address indexed newOracle)
func (_Config *ConfigFilterer) WatchPriceOracleChanged(opts *bind.WatchOpts, sink chan<- *ConfigPriceOracleChanged, oldOracle []common.Address, newOracle []common.Address) (event.Subscription, error) {

	var oldOracleRule []interface{}
	for _, oldOracleItem := range oldOracle {
		oldOracleRule = append(oldOracleRule, oldOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "PriceOracleChanged", oldOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigPriceOracleChanged)
				if err := _Config.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
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

// ParsePriceOracleChanged is a log parse operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed oldOracle, address indexed newOracle)
func (_Config *ConfigFilterer) ParsePriceOracleChanged(log types.Log) (*ConfigPriceOracleChanged, error) {
	event := new(ConfigPriceOracleChanged)
	if err := _Config.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRebaseIntervalChangedIterator is returned from FilterRebaseIntervalChanged and is used to iterate over the raw logs and unpacked data for RebaseIntervalChanged events raised by the Config contract.
type ConfigRebaseIntervalChangedIterator struct {
	Event *ConfigRebaseIntervalChanged // Event containing the contract specifics and raw log

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
func (it *ConfigRebaseIntervalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRebaseIntervalChanged)
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
		it.Event = new(ConfigRebaseIntervalChanged)
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
func (it *ConfigRebaseIntervalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRebaseIntervalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRebaseIntervalChanged represents a RebaseIntervalChanged event raised by the Config contract.
type ConfigRebaseIntervalChanged struct {
	OldInterval *big.Int
	NewInterval *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRebaseIntervalChanged is a free log retrieval operation binding the contract event 0xe53d8256da2697e24f3ec42a36b999071c1254abd9d21308a86ded54d1010fbe.
//
// Solidity: event RebaseIntervalChanged(uint256 oldInterval, uint256 newInterval)
func (_Config *ConfigFilterer) FilterRebaseIntervalChanged(opts *bind.FilterOpts) (*ConfigRebaseIntervalChangedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "RebaseIntervalChanged")
	if err != nil {
		return nil, err
	}
	return &ConfigRebaseIntervalChangedIterator{contract: _Config.contract, event: "RebaseIntervalChanged", logs: logs, sub: sub}, nil
}

// WatchRebaseIntervalChanged is a free log subscription operation binding the contract event 0xe53d8256da2697e24f3ec42a36b999071c1254abd9d21308a86ded54d1010fbe.
//
// Solidity: event RebaseIntervalChanged(uint256 oldInterval, uint256 newInterval)
func (_Config *ConfigFilterer) WatchRebaseIntervalChanged(opts *bind.WatchOpts, sink chan<- *ConfigRebaseIntervalChanged) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "RebaseIntervalChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRebaseIntervalChanged)
				if err := _Config.contract.UnpackLog(event, "RebaseIntervalChanged", log); err != nil {
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

// ParseRebaseIntervalChanged is a log parse operation binding the contract event 0xe53d8256da2697e24f3ec42a36b999071c1254abd9d21308a86ded54d1010fbe.
//
// Solidity: event RebaseIntervalChanged(uint256 oldInterval, uint256 newInterval)
func (_Config *ConfigFilterer) ParseRebaseIntervalChanged(log types.Log) (*ConfigRebaseIntervalChanged, error) {
	event := new(ConfigRebaseIntervalChanged)
	if err := _Config.contract.UnpackLog(event, "RebaseIntervalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRebasePriceGapChangedIterator is returned from FilterRebasePriceGapChanged and is used to iterate over the raw logs and unpacked data for RebasePriceGapChanged events raised by the Config contract.
type ConfigRebasePriceGapChangedIterator struct {
	Event *ConfigRebasePriceGapChanged // Event containing the contract specifics and raw log

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
func (it *ConfigRebasePriceGapChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRebasePriceGapChanged)
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
		it.Event = new(ConfigRebasePriceGapChanged)
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
func (it *ConfigRebasePriceGapChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRebasePriceGapChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRebasePriceGapChanged represents a RebasePriceGapChanged event raised by the Config contract.
type ConfigRebasePriceGapChanged struct {
	OldGap *big.Int
	NewGap *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRebasePriceGapChanged is a free log retrieval operation binding the contract event 0x3dcf61c33d20fa3fc1936a5cefe7c78cc9b97adaadcf393a558b922577b83ba4.
//
// Solidity: event RebasePriceGapChanged(uint256 oldGap, uint256 newGap)
func (_Config *ConfigFilterer) FilterRebasePriceGapChanged(opts *bind.FilterOpts) (*ConfigRebasePriceGapChangedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "RebasePriceGapChanged")
	if err != nil {
		return nil, err
	}
	return &ConfigRebasePriceGapChangedIterator{contract: _Config.contract, event: "RebasePriceGapChanged", logs: logs, sub: sub}, nil
}

// WatchRebasePriceGapChanged is a free log subscription operation binding the contract event 0x3dcf61c33d20fa3fc1936a5cefe7c78cc9b97adaadcf393a558b922577b83ba4.
//
// Solidity: event RebasePriceGapChanged(uint256 oldGap, uint256 newGap)
func (_Config *ConfigFilterer) WatchRebasePriceGapChanged(opts *bind.WatchOpts, sink chan<- *ConfigRebasePriceGapChanged) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "RebasePriceGapChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRebasePriceGapChanged)
				if err := _Config.contract.UnpackLog(event, "RebasePriceGapChanged", log); err != nil {
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

// ParseRebasePriceGapChanged is a log parse operation binding the contract event 0x3dcf61c33d20fa3fc1936a5cefe7c78cc9b97adaadcf393a558b922577b83ba4.
//
// Solidity: event RebasePriceGapChanged(uint256 oldGap, uint256 newGap)
func (_Config *ConfigFilterer) ParseRebasePriceGapChanged(log types.Log) (*ConfigRebasePriceGapChanged, error) {
	event := new(ConfigRebasePriceGapChanged)
	if err := _Config.contract.UnpackLog(event, "RebasePriceGapChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRouterRegisteredIterator is returned from FilterRouterRegistered and is used to iterate over the raw logs and unpacked data for RouterRegistered events raised by the Config contract.
type ConfigRouterRegisteredIterator struct {
	Event *ConfigRouterRegistered // Event containing the contract specifics and raw log

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
func (it *ConfigRouterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRouterRegistered)
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
		it.Event = new(ConfigRouterRegistered)
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
func (it *ConfigRouterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRouterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRouterRegistered represents a RouterRegistered event raised by the Config contract.
type ConfigRouterRegistered struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterRegistered is a free log retrieval operation binding the contract event 0x94df8c3a8087dce110e5fbc5acf380c83c94bbd31b2c8ed4c08e1396a696e1a8.
//
// Solidity: event RouterRegistered(address indexed router)
func (_Config *ConfigFilterer) FilterRouterRegistered(opts *bind.FilterOpts, router []common.Address) (*ConfigRouterRegisteredIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "RouterRegistered", routerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigRouterRegisteredIterator{contract: _Config.contract, event: "RouterRegistered", logs: logs, sub: sub}, nil
}

// WatchRouterRegistered is a free log subscription operation binding the contract event 0x94df8c3a8087dce110e5fbc5acf380c83c94bbd31b2c8ed4c08e1396a696e1a8.
//
// Solidity: event RouterRegistered(address indexed router)
func (_Config *ConfigFilterer) WatchRouterRegistered(opts *bind.WatchOpts, sink chan<- *ConfigRouterRegistered, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "RouterRegistered", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRouterRegistered)
				if err := _Config.contract.UnpackLog(event, "RouterRegistered", log); err != nil {
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

// ParseRouterRegistered is a log parse operation binding the contract event 0x94df8c3a8087dce110e5fbc5acf380c83c94bbd31b2c8ed4c08e1396a696e1a8.
//
// Solidity: event RouterRegistered(address indexed router)
func (_Config *ConfigFilterer) ParseRouterRegistered(log types.Log) (*ConfigRouterRegistered, error) {
	event := new(ConfigRouterRegistered)
	if err := _Config.contract.UnpackLog(event, "RouterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRouterUnregisteredIterator is returned from FilterRouterUnregistered and is used to iterate over the raw logs and unpacked data for RouterUnregistered events raised by the Config contract.
type ConfigRouterUnregisteredIterator struct {
	Event *ConfigRouterUnregistered // Event containing the contract specifics and raw log

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
func (it *ConfigRouterUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRouterUnregistered)
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
		it.Event = new(ConfigRouterUnregistered)
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
func (it *ConfigRouterUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRouterUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRouterUnregistered represents a RouterUnregistered event raised by the Config contract.
type ConfigRouterUnregistered struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterUnregistered is a free log retrieval operation binding the contract event 0x75651c5c467d4644a640a67bb1f11c3b5e6e05c68283a21a248c9b2685e4fa09.
//
// Solidity: event RouterUnregistered(address indexed router)
func (_Config *ConfigFilterer) FilterRouterUnregistered(opts *bind.FilterOpts, router []common.Address) (*ConfigRouterUnregisteredIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "RouterUnregistered", routerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigRouterUnregisteredIterator{contract: _Config.contract, event: "RouterUnregistered", logs: logs, sub: sub}, nil
}

// WatchRouterUnregistered is a free log subscription operation binding the contract event 0x75651c5c467d4644a640a67bb1f11c3b5e6e05c68283a21a248c9b2685e4fa09.
//
// Solidity: event RouterUnregistered(address indexed router)
func (_Config *ConfigFilterer) WatchRouterUnregistered(opts *bind.WatchOpts, sink chan<- *ConfigRouterUnregistered, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "RouterUnregistered", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRouterUnregistered)
				if err := _Config.contract.UnpackLog(event, "RouterUnregistered", log); err != nil {
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

// ParseRouterUnregistered is a log parse operation binding the contract event 0x75651c5c467d4644a640a67bb1f11c3b5e6e05c68283a21a248c9b2685e4fa09.
//
// Solidity: event RouterUnregistered(address indexed router)
func (_Config *ConfigFilterer) ParseRouterUnregistered(log types.Log) (*ConfigRouterUnregistered, error) {
	event := new(ConfigRouterUnregistered)
	if err := _Config.contract.UnpackLog(event, "RouterUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetBetaIterator is returned from FilterSetBeta and is used to iterate over the raw logs and unpacked data for SetBeta events raised by the Config contract.
type ConfigSetBetaIterator struct {
	Event *ConfigSetBeta // Event containing the contract specifics and raw log

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
func (it *ConfigSetBetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetBeta)
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
		it.Event = new(ConfigSetBeta)
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
func (it *ConfigSetBetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetBetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetBeta represents a SetBeta event raised by the Config contract.
type ConfigSetBeta struct {
	OldBeta *big.Int
	Beta    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBeta is a free log retrieval operation binding the contract event 0xa2545869d941985a4dda10a4646a6ca753ab7789f05cf35cb7fd1ca56f7ea412.
//
// Solidity: event SetBeta(uint256 oldBeta, uint256 beta)
func (_Config *ConfigFilterer) FilterSetBeta(opts *bind.FilterOpts) (*ConfigSetBetaIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetBeta")
	if err != nil {
		return nil, err
	}
	return &ConfigSetBetaIterator{contract: _Config.contract, event: "SetBeta", logs: logs, sub: sub}, nil
}

// WatchSetBeta is a free log subscription operation binding the contract event 0xa2545869d941985a4dda10a4646a6ca753ab7789f05cf35cb7fd1ca56f7ea412.
//
// Solidity: event SetBeta(uint256 oldBeta, uint256 beta)
func (_Config *ConfigFilterer) WatchSetBeta(opts *bind.WatchOpts, sink chan<- *ConfigSetBeta) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetBeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetBeta)
				if err := _Config.contract.UnpackLog(event, "SetBeta", log); err != nil {
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

// ParseSetBeta is a log parse operation binding the contract event 0xa2545869d941985a4dda10a4646a6ca753ab7789f05cf35cb7fd1ca56f7ea412.
//
// Solidity: event SetBeta(uint256 oldBeta, uint256 beta)
func (_Config *ConfigFilterer) ParseSetBeta(log types.Log) (*ConfigSetBeta, error) {
	event := new(ConfigSetBeta)
	if err := _Config.contract.UnpackLog(event, "SetBeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetEmergencyIterator is returned from FilterSetEmergency and is used to iterate over the raw logs and unpacked data for SetEmergency events raised by the Config contract.
type ConfigSetEmergencyIterator struct {
	Event *ConfigSetEmergency // Event containing the contract specifics and raw log

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
func (it *ConfigSetEmergencyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetEmergency)
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
		it.Event = new(ConfigSetEmergency)
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
func (it *ConfigSetEmergencyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetEmergencyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetEmergency represents a SetEmergency event raised by the Config contract.
type ConfigSetEmergency struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetEmergency is a free log retrieval operation binding the contract event 0xecd4fa7e5f855e908d57952c76af5585f7725f3524c7da4f3d336619e17b276d.
//
// Solidity: event SetEmergency(address indexed router)
func (_Config *ConfigFilterer) FilterSetEmergency(opts *bind.FilterOpts, router []common.Address) (*ConfigSetEmergencyIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetEmergency", routerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigSetEmergencyIterator{contract: _Config.contract, event: "SetEmergency", logs: logs, sub: sub}, nil
}

// WatchSetEmergency is a free log subscription operation binding the contract event 0xecd4fa7e5f855e908d57952c76af5585f7725f3524c7da4f3d336619e17b276d.
//
// Solidity: event SetEmergency(address indexed router)
func (_Config *ConfigFilterer) WatchSetEmergency(opts *bind.WatchOpts, sink chan<- *ConfigSetEmergency, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetEmergency", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetEmergency)
				if err := _Config.contract.UnpackLog(event, "SetEmergency", log); err != nil {
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

// ParseSetEmergency is a log parse operation binding the contract event 0xecd4fa7e5f855e908d57952c76af5585f7725f3524c7da4f3d336619e17b276d.
//
// Solidity: event SetEmergency(address indexed router)
func (_Config *ConfigFilterer) ParseSetEmergency(log types.Log) (*ConfigSetEmergency, error) {
	event := new(ConfigSetEmergency)
	if err := _Config.contract.UnpackLog(event, "SetEmergency", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetFeeParameterIterator is returned from FilterSetFeeParameter and is used to iterate over the raw logs and unpacked data for SetFeeParameter events raised by the Config contract.
type ConfigSetFeeParameterIterator struct {
	Event *ConfigSetFeeParameter // Event containing the contract specifics and raw log

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
func (it *ConfigSetFeeParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetFeeParameter)
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
		it.Event = new(ConfigSetFeeParameter)
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
func (it *ConfigSetFeeParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetFeeParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetFeeParameter represents a SetFeeParameter event raised by the Config contract.
type ConfigSetFeeParameter struct {
	OldFeeParameter *big.Int
	FeeParameter    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeParameter is a free log retrieval operation binding the contract event 0xd467fef04f1829519a852bd1ec06607466b3f22c9a92c9b0ec2efab69e98f0f2.
//
// Solidity: event SetFeeParameter(uint256 oldFeeParameter, uint256 feeParameter)
func (_Config *ConfigFilterer) FilterSetFeeParameter(opts *bind.FilterOpts) (*ConfigSetFeeParameterIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetFeeParameter")
	if err != nil {
		return nil, err
	}
	return &ConfigSetFeeParameterIterator{contract: _Config.contract, event: "SetFeeParameter", logs: logs, sub: sub}, nil
}

// WatchSetFeeParameter is a free log subscription operation binding the contract event 0xd467fef04f1829519a852bd1ec06607466b3f22c9a92c9b0ec2efab69e98f0f2.
//
// Solidity: event SetFeeParameter(uint256 oldFeeParameter, uint256 feeParameter)
func (_Config *ConfigFilterer) WatchSetFeeParameter(opts *bind.WatchOpts, sink chan<- *ConfigSetFeeParameter) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetFeeParameter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetFeeParameter)
				if err := _Config.contract.UnpackLog(event, "SetFeeParameter", log); err != nil {
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

// ParseSetFeeParameter is a log parse operation binding the contract event 0xd467fef04f1829519a852bd1ec06607466b3f22c9a92c9b0ec2efab69e98f0f2.
//
// Solidity: event SetFeeParameter(uint256 oldFeeParameter, uint256 feeParameter)
func (_Config *ConfigFilterer) ParseSetFeeParameter(log types.Log) (*ConfigSetFeeParameter, error) {
	event := new(ConfigSetFeeParameter)
	if err := _Config.contract.UnpackLog(event, "SetFeeParameter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetInitMarginRatioIterator is returned from FilterSetInitMarginRatio and is used to iterate over the raw logs and unpacked data for SetInitMarginRatio events raised by the Config contract.
type ConfigSetInitMarginRatioIterator struct {
	Event *ConfigSetInitMarginRatio // Event containing the contract specifics and raw log

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
func (it *ConfigSetInitMarginRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetInitMarginRatio)
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
		it.Event = new(ConfigSetInitMarginRatio)
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
func (it *ConfigSetInitMarginRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetInitMarginRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetInitMarginRatio represents a SetInitMarginRatio event raised by the Config contract.
type ConfigSetInitMarginRatio struct {
	OldInitMarginRatio *big.Int
	InitMarginRatio    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetInitMarginRatio is a free log retrieval operation binding the contract event 0x5edea59d534a04dd802be5d64f704cb2beda3fc7e16ee59646b655352202da80.
//
// Solidity: event SetInitMarginRatio(uint256 oldInitMarginRatio, uint256 initMarginRatio)
func (_Config *ConfigFilterer) FilterSetInitMarginRatio(opts *bind.FilterOpts) (*ConfigSetInitMarginRatioIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetInitMarginRatio")
	if err != nil {
		return nil, err
	}
	return &ConfigSetInitMarginRatioIterator{contract: _Config.contract, event: "SetInitMarginRatio", logs: logs, sub: sub}, nil
}

// WatchSetInitMarginRatio is a free log subscription operation binding the contract event 0x5edea59d534a04dd802be5d64f704cb2beda3fc7e16ee59646b655352202da80.
//
// Solidity: event SetInitMarginRatio(uint256 oldInitMarginRatio, uint256 initMarginRatio)
func (_Config *ConfigFilterer) WatchSetInitMarginRatio(opts *bind.WatchOpts, sink chan<- *ConfigSetInitMarginRatio) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetInitMarginRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetInitMarginRatio)
				if err := _Config.contract.UnpackLog(event, "SetInitMarginRatio", log); err != nil {
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

// ParseSetInitMarginRatio is a log parse operation binding the contract event 0x5edea59d534a04dd802be5d64f704cb2beda3fc7e16ee59646b655352202da80.
//
// Solidity: event SetInitMarginRatio(uint256 oldInitMarginRatio, uint256 initMarginRatio)
func (_Config *ConfigFilterer) ParseSetInitMarginRatio(log types.Log) (*ConfigSetInitMarginRatio, error) {
	event := new(ConfigSetInitMarginRatio)
	if err := _Config.contract.UnpackLog(event, "SetInitMarginRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetLiquidateFeeRatioIterator is returned from FilterSetLiquidateFeeRatio and is used to iterate over the raw logs and unpacked data for SetLiquidateFeeRatio events raised by the Config contract.
type ConfigSetLiquidateFeeRatioIterator struct {
	Event *ConfigSetLiquidateFeeRatio // Event containing the contract specifics and raw log

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
func (it *ConfigSetLiquidateFeeRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetLiquidateFeeRatio)
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
		it.Event = new(ConfigSetLiquidateFeeRatio)
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
func (it *ConfigSetLiquidateFeeRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetLiquidateFeeRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetLiquidateFeeRatio represents a SetLiquidateFeeRatio event raised by the Config contract.
type ConfigSetLiquidateFeeRatio struct {
	OldLiquidateFeeRatio *big.Int
	LiquidateFeeRatio    *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetLiquidateFeeRatio is a free log retrieval operation binding the contract event 0x6e8fb8d7ccbce74238212d18376842dde6f674e4404bdcc2762602939d0fb86c.
//
// Solidity: event SetLiquidateFeeRatio(uint256 oldLiquidateFeeRatio, uint256 liquidateFeeRatio)
func (_Config *ConfigFilterer) FilterSetLiquidateFeeRatio(opts *bind.FilterOpts) (*ConfigSetLiquidateFeeRatioIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetLiquidateFeeRatio")
	if err != nil {
		return nil, err
	}
	return &ConfigSetLiquidateFeeRatioIterator{contract: _Config.contract, event: "SetLiquidateFeeRatio", logs: logs, sub: sub}, nil
}

// WatchSetLiquidateFeeRatio is a free log subscription operation binding the contract event 0x6e8fb8d7ccbce74238212d18376842dde6f674e4404bdcc2762602939d0fb86c.
//
// Solidity: event SetLiquidateFeeRatio(uint256 oldLiquidateFeeRatio, uint256 liquidateFeeRatio)
func (_Config *ConfigFilterer) WatchSetLiquidateFeeRatio(opts *bind.WatchOpts, sink chan<- *ConfigSetLiquidateFeeRatio) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetLiquidateFeeRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetLiquidateFeeRatio)
				if err := _Config.contract.UnpackLog(event, "SetLiquidateFeeRatio", log); err != nil {
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

// ParseSetLiquidateFeeRatio is a log parse operation binding the contract event 0x6e8fb8d7ccbce74238212d18376842dde6f674e4404bdcc2762602939d0fb86c.
//
// Solidity: event SetLiquidateFeeRatio(uint256 oldLiquidateFeeRatio, uint256 liquidateFeeRatio)
func (_Config *ConfigFilterer) ParseSetLiquidateFeeRatio(log types.Log) (*ConfigSetLiquidateFeeRatio, error) {
	event := new(ConfigSetLiquidateFeeRatio)
	if err := _Config.contract.UnpackLog(event, "SetLiquidateFeeRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetLiquidateThresholdIterator is returned from FilterSetLiquidateThreshold and is used to iterate over the raw logs and unpacked data for SetLiquidateThreshold events raised by the Config contract.
type ConfigSetLiquidateThresholdIterator struct {
	Event *ConfigSetLiquidateThreshold // Event containing the contract specifics and raw log

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
func (it *ConfigSetLiquidateThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetLiquidateThreshold)
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
		it.Event = new(ConfigSetLiquidateThreshold)
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
func (it *ConfigSetLiquidateThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetLiquidateThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetLiquidateThreshold represents a SetLiquidateThreshold event raised by the Config contract.
type ConfigSetLiquidateThreshold struct {
	OldLiquidateThreshold *big.Int
	LiquidateThreshold    *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetLiquidateThreshold is a free log retrieval operation binding the contract event 0x33cf718637d6f1faf404b6902a2b551102381524c3b3c83f85cb3681e40758f7.
//
// Solidity: event SetLiquidateThreshold(uint256 oldLiquidateThreshold, uint256 liquidateThreshold)
func (_Config *ConfigFilterer) FilterSetLiquidateThreshold(opts *bind.FilterOpts) (*ConfigSetLiquidateThresholdIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetLiquidateThreshold")
	if err != nil {
		return nil, err
	}
	return &ConfigSetLiquidateThresholdIterator{contract: _Config.contract, event: "SetLiquidateThreshold", logs: logs, sub: sub}, nil
}

// WatchSetLiquidateThreshold is a free log subscription operation binding the contract event 0x33cf718637d6f1faf404b6902a2b551102381524c3b3c83f85cb3681e40758f7.
//
// Solidity: event SetLiquidateThreshold(uint256 oldLiquidateThreshold, uint256 liquidateThreshold)
func (_Config *ConfigFilterer) WatchSetLiquidateThreshold(opts *bind.WatchOpts, sink chan<- *ConfigSetLiquidateThreshold) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetLiquidateThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetLiquidateThreshold)
				if err := _Config.contract.UnpackLog(event, "SetLiquidateThreshold", log); err != nil {
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

// ParseSetLiquidateThreshold is a log parse operation binding the contract event 0x33cf718637d6f1faf404b6902a2b551102381524c3b3c83f85cb3681e40758f7.
//
// Solidity: event SetLiquidateThreshold(uint256 oldLiquidateThreshold, uint256 liquidateThreshold)
func (_Config *ConfigFilterer) ParseSetLiquidateThreshold(log types.Log) (*ConfigSetLiquidateThreshold, error) {
	event := new(ConfigSetLiquidateThreshold)
	if err := _Config.contract.UnpackLog(event, "SetLiquidateThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetLpWithdrawThresholdIterator is returned from FilterSetLpWithdrawThreshold and is used to iterate over the raw logs and unpacked data for SetLpWithdrawThreshold events raised by the Config contract.
type ConfigSetLpWithdrawThresholdIterator struct {
	Event *ConfigSetLpWithdrawThreshold // Event containing the contract specifics and raw log

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
func (it *ConfigSetLpWithdrawThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetLpWithdrawThreshold)
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
		it.Event = new(ConfigSetLpWithdrawThreshold)
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
func (it *ConfigSetLpWithdrawThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetLpWithdrawThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetLpWithdrawThreshold represents a SetLpWithdrawThreshold event raised by the Config contract.
type ConfigSetLpWithdrawThreshold struct {
	OldLpWithdrawThreshold *big.Int
	LpWithdrawThreshold    *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetLpWithdrawThreshold is a free log retrieval operation binding the contract event 0x325514ba4c843e515d3b79b4ed830f5a67fcbc43785d7ae07b01826768484471.
//
// Solidity: event SetLpWithdrawThreshold(uint256 oldLpWithdrawThreshold, uint256 lpWithdrawThreshold)
func (_Config *ConfigFilterer) FilterSetLpWithdrawThreshold(opts *bind.FilterOpts) (*ConfigSetLpWithdrawThresholdIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetLpWithdrawThreshold")
	if err != nil {
		return nil, err
	}
	return &ConfigSetLpWithdrawThresholdIterator{contract: _Config.contract, event: "SetLpWithdrawThreshold", logs: logs, sub: sub}, nil
}

// WatchSetLpWithdrawThreshold is a free log subscription operation binding the contract event 0x325514ba4c843e515d3b79b4ed830f5a67fcbc43785d7ae07b01826768484471.
//
// Solidity: event SetLpWithdrawThreshold(uint256 oldLpWithdrawThreshold, uint256 lpWithdrawThreshold)
func (_Config *ConfigFilterer) WatchSetLpWithdrawThreshold(opts *bind.WatchOpts, sink chan<- *ConfigSetLpWithdrawThreshold) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetLpWithdrawThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetLpWithdrawThreshold)
				if err := _Config.contract.UnpackLog(event, "SetLpWithdrawThreshold", log); err != nil {
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

// ParseSetLpWithdrawThreshold is a log parse operation binding the contract event 0x325514ba4c843e515d3b79b4ed830f5a67fcbc43785d7ae07b01826768484471.
//
// Solidity: event SetLpWithdrawThreshold(uint256 oldLpWithdrawThreshold, uint256 lpWithdrawThreshold)
func (_Config *ConfigFilterer) ParseSetLpWithdrawThreshold(log types.Log) (*ConfigSetLpWithdrawThreshold, error) {
	event := new(ConfigSetLpWithdrawThreshold)
	if err := _Config.contract.UnpackLog(event, "SetLpWithdrawThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetMaxCPFBoostIterator is returned from FilterSetMaxCPFBoost and is used to iterate over the raw logs and unpacked data for SetMaxCPFBoost events raised by the Config contract.
type ConfigSetMaxCPFBoostIterator struct {
	Event *ConfigSetMaxCPFBoost // Event containing the contract specifics and raw log

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
func (it *ConfigSetMaxCPFBoostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetMaxCPFBoost)
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
		it.Event = new(ConfigSetMaxCPFBoost)
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
func (it *ConfigSetMaxCPFBoostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetMaxCPFBoostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetMaxCPFBoost represents a SetMaxCPFBoost event raised by the Config contract.
type ConfigSetMaxCPFBoost struct {
	OldMaxCPFBoost *big.Int
	MaxCPFBoost    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetMaxCPFBoost is a free log retrieval operation binding the contract event 0x925aaf6ac64a831f734927a8208b1f475181ebdc9aa05aca723b5af4b419a3f8.
//
// Solidity: event SetMaxCPFBoost(uint256 oldMaxCPFBoost, uint256 maxCPFBoost)
func (_Config *ConfigFilterer) FilterSetMaxCPFBoost(opts *bind.FilterOpts) (*ConfigSetMaxCPFBoostIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetMaxCPFBoost")
	if err != nil {
		return nil, err
	}
	return &ConfigSetMaxCPFBoostIterator{contract: _Config.contract, event: "SetMaxCPFBoost", logs: logs, sub: sub}, nil
}

// WatchSetMaxCPFBoost is a free log subscription operation binding the contract event 0x925aaf6ac64a831f734927a8208b1f475181ebdc9aa05aca723b5af4b419a3f8.
//
// Solidity: event SetMaxCPFBoost(uint256 oldMaxCPFBoost, uint256 maxCPFBoost)
func (_Config *ConfigFilterer) WatchSetMaxCPFBoost(opts *bind.WatchOpts, sink chan<- *ConfigSetMaxCPFBoost) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetMaxCPFBoost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetMaxCPFBoost)
				if err := _Config.contract.UnpackLog(event, "SetMaxCPFBoost", log); err != nil {
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

// ParseSetMaxCPFBoost is a log parse operation binding the contract event 0x925aaf6ac64a831f734927a8208b1f475181ebdc9aa05aca723b5af4b419a3f8.
//
// Solidity: event SetMaxCPFBoost(uint256 oldMaxCPFBoost, uint256 maxCPFBoost)
func (_Config *ConfigFilterer) ParseSetMaxCPFBoost(log types.Log) (*ConfigSetMaxCPFBoost, error) {
	event := new(ConfigSetMaxCPFBoost)
	if err := _Config.contract.UnpackLog(event, "SetMaxCPFBoost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigTradingSlippageChangedIterator is returned from FilterTradingSlippageChanged and is used to iterate over the raw logs and unpacked data for TradingSlippageChanged events raised by the Config contract.
type ConfigTradingSlippageChangedIterator struct {
	Event *ConfigTradingSlippageChanged // Event containing the contract specifics and raw log

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
func (it *ConfigTradingSlippageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigTradingSlippageChanged)
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
		it.Event = new(ConfigTradingSlippageChanged)
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
func (it *ConfigTradingSlippageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigTradingSlippageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigTradingSlippageChanged represents a TradingSlippageChanged event raised by the Config contract.
type ConfigTradingSlippageChanged struct {
	OldTradingSlippage *big.Int
	NewTradingSlippage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTradingSlippageChanged is a free log retrieval operation binding the contract event 0x93301c5c8702bd9fef38d0e194884f0dcdceec480fbdac4620781ad65ca09dea.
//
// Solidity: event TradingSlippageChanged(uint256 oldTradingSlippage, uint256 newTradingSlippage)
func (_Config *ConfigFilterer) FilterTradingSlippageChanged(opts *bind.FilterOpts) (*ConfigTradingSlippageChangedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "TradingSlippageChanged")
	if err != nil {
		return nil, err
	}
	return &ConfigTradingSlippageChangedIterator{contract: _Config.contract, event: "TradingSlippageChanged", logs: logs, sub: sub}, nil
}

// WatchTradingSlippageChanged is a free log subscription operation binding the contract event 0x93301c5c8702bd9fef38d0e194884f0dcdceec480fbdac4620781ad65ca09dea.
//
// Solidity: event TradingSlippageChanged(uint256 oldTradingSlippage, uint256 newTradingSlippage)
func (_Config *ConfigFilterer) WatchTradingSlippageChanged(opts *bind.WatchOpts, sink chan<- *ConfigTradingSlippageChanged) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "TradingSlippageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigTradingSlippageChanged)
				if err := _Config.contract.UnpackLog(event, "TradingSlippageChanged", log); err != nil {
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

// ParseTradingSlippageChanged is a log parse operation binding the contract event 0x93301c5c8702bd9fef38d0e194884f0dcdceec480fbdac4620781ad65ca09dea.
//
// Solidity: event TradingSlippageChanged(uint256 oldTradingSlippage, uint256 newTradingSlippage)
func (_Config *ConfigFilterer) ParseTradingSlippageChanged(log types.Log) (*ConfigTradingSlippageChanged, error) {
	event := new(ConfigTradingSlippageChanged)
	if err := _Config.contract.UnpackLog(event, "TradingSlippageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
