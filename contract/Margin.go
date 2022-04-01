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

// IMarginPosition is an auto generated low-level Go binding around an user-defined struct.
type IMarginPosition struct {
	QuoteSize *big.Int
	BaseSize  *big.Int
	TradeSize *big.Int
}

// MarginMetaData contains all meta data concerning the Margin contract.
var MarginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMargin.Position\",\"name\":\"position\",\"type\":\"tuple\"}],\"name\":\"AddMargin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMargin.Position\",\"name\":\"position\",\"type\":\"tuple\"}],\"name\":\"ClosePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMargin.Position\",\"name\":\"position\",\"type\":\"tuple\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMargin.Position\",\"name\":\"position\",\"type\":\"tuple\"}],\"name\":\"OpenPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmountFromMargin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMargin.Position\",\"name\":\"position\",\"type\":\"tuple\"}],\"name\":\"RemoveMargin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cpf\",\"type\":\"int256\"}],\"name\":\"UpdateCPF\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"addMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"calDebtRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debtRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"calFundingFee\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"calUnrealizedPnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"unrealizedPnl\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"canLiquidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"closePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewLatestCPF\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawable\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"amm_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateCPF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"netPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"openPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySwapBaseWithAmm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"removeMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalQuoteLong\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalQuoteShort\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderCPF\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderPositionMap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCPF\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"newLatestCPF\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarginABI is the input ABI used to generate the binding from.
// Deprecated: Use MarginMetaData.ABI instead.
var MarginABI = MarginMetaData.ABI

// Margin is an auto generated Go binding around an Ethereum contract.
type Margin struct {
	MarginCaller     // Read-only binding to the contract
	MarginTransactor // Write-only binding to the contract
	MarginFilterer   // Log filterer for contract events
}

// MarginCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarginSession struct {
	Contract     *Margin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarginCallerSession struct {
	Contract *MarginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarginTransactorSession struct {
	Contract     *MarginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarginRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarginRaw struct {
	Contract *Margin // Generic contract binding to access the raw methods on
}

// MarginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarginCallerRaw struct {
	Contract *MarginCaller // Generic read-only contract binding to access the raw methods on
}

// MarginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarginTransactorRaw struct {
	Contract *MarginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMargin creates a new instance of Margin, bound to a specific deployed contract.
func NewMargin(address common.Address, backend bind.ContractBackend) (*Margin, error) {
	contract, err := bindMargin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Margin{MarginCaller: MarginCaller{contract: contract}, MarginTransactor: MarginTransactor{contract: contract}, MarginFilterer: MarginFilterer{contract: contract}}, nil
}

// NewMarginCaller creates a new read-only instance of Margin, bound to a specific deployed contract.
func NewMarginCaller(address common.Address, caller bind.ContractCaller) (*MarginCaller, error) {
	contract, err := bindMargin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarginCaller{contract: contract}, nil
}

// NewMarginTransactor creates a new write-only instance of Margin, bound to a specific deployed contract.
func NewMarginTransactor(address common.Address, transactor bind.ContractTransactor) (*MarginTransactor, error) {
	contract, err := bindMargin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarginTransactor{contract: contract}, nil
}

// NewMarginFilterer creates a new log filterer instance of Margin, bound to a specific deployed contract.
func NewMarginFilterer(address common.Address, filterer bind.ContractFilterer) (*MarginFilterer, error) {
	contract, err := bindMargin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarginFilterer{contract: contract}, nil
}

// bindMargin binds a generic wrapper to an already deployed contract.
func bindMargin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Margin *MarginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Margin.Contract.MarginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Margin *MarginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.Contract.MarginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Margin *MarginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Margin.Contract.MarginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Margin *MarginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Margin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Margin *MarginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Margin *MarginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Margin.Contract.contract.Transact(opts, method, params...)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_Margin *MarginCaller) Amm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "amm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_Margin *MarginSession) Amm() (common.Address, error) {
	return _Margin.Contract.Amm(&_Margin.CallOpts)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_Margin *MarginCallerSession) Amm() (common.Address, error) {
	return _Margin.Contract.Amm(&_Margin.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Margin *MarginCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Margin *MarginSession) BaseToken() (common.Address, error) {
	return _Margin.Contract.BaseToken(&_Margin.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Margin *MarginCallerSession) BaseToken() (common.Address, error) {
	return _Margin.Contract.BaseToken(&_Margin.CallOpts)
}

// CalDebtRatio is a free data retrieval call binding the contract method 0x763620ac.
//
// Solidity: function calDebtRatio(address trader) view returns(uint256 debtRatio)
func (_Margin *MarginCaller) CalDebtRatio(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "calDebtRatio", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalDebtRatio is a free data retrieval call binding the contract method 0x763620ac.
//
// Solidity: function calDebtRatio(address trader) view returns(uint256 debtRatio)
func (_Margin *MarginSession) CalDebtRatio(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.CalDebtRatio(&_Margin.CallOpts, trader)
}

// CalDebtRatio is a free data retrieval call binding the contract method 0x763620ac.
//
// Solidity: function calDebtRatio(address trader) view returns(uint256 debtRatio)
func (_Margin *MarginCallerSession) CalDebtRatio(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.CalDebtRatio(&_Margin.CallOpts, trader)
}

// CalFundingFee is a free data retrieval call binding the contract method 0x55e89810.
//
// Solidity: function calFundingFee(address trader) view returns(int256)
func (_Margin *MarginCaller) CalFundingFee(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "calFundingFee", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalFundingFee is a free data retrieval call binding the contract method 0x55e89810.
//
// Solidity: function calFundingFee(address trader) view returns(int256)
func (_Margin *MarginSession) CalFundingFee(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.CalFundingFee(&_Margin.CallOpts, trader)
}

// CalFundingFee is a free data retrieval call binding the contract method 0x55e89810.
//
// Solidity: function calFundingFee(address trader) view returns(int256)
func (_Margin *MarginCallerSession) CalFundingFee(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.CalFundingFee(&_Margin.CallOpts, trader)
}

// CalUnrealizedPnl is a free data retrieval call binding the contract method 0xee4424e5.
//
// Solidity: function calUnrealizedPnl(address trader) view returns(int256 unrealizedPnl)
func (_Margin *MarginCaller) CalUnrealizedPnl(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "calUnrealizedPnl", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalUnrealizedPnl is a free data retrieval call binding the contract method 0xee4424e5.
//
// Solidity: function calUnrealizedPnl(address trader) view returns(int256 unrealizedPnl)
func (_Margin *MarginSession) CalUnrealizedPnl(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.CalUnrealizedPnl(&_Margin.CallOpts, trader)
}

// CalUnrealizedPnl is a free data retrieval call binding the contract method 0xee4424e5.
//
// Solidity: function calUnrealizedPnl(address trader) view returns(int256 unrealizedPnl)
func (_Margin *MarginCallerSession) CalUnrealizedPnl(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.CalUnrealizedPnl(&_Margin.CallOpts, trader)
}

// CanLiquidate is a free data retrieval call binding the contract method 0xb9f4ff55.
//
// Solidity: function canLiquidate(address trader) view returns(bool)
func (_Margin *MarginCaller) CanLiquidate(opts *bind.CallOpts, trader common.Address) (bool, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "canLiquidate", trader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanLiquidate is a free data retrieval call binding the contract method 0xb9f4ff55.
//
// Solidity: function canLiquidate(address trader) view returns(bool)
func (_Margin *MarginSession) CanLiquidate(trader common.Address) (bool, error) {
	return _Margin.Contract.CanLiquidate(&_Margin.CallOpts, trader)
}

// CanLiquidate is a free data retrieval call binding the contract method 0xb9f4ff55.
//
// Solidity: function canLiquidate(address trader) view returns(bool)
func (_Margin *MarginCallerSession) CanLiquidate(trader common.Address) (bool, error) {
	return _Margin.Contract.CanLiquidate(&_Margin.CallOpts, trader)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Margin *MarginCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Margin *MarginSession) Config() (common.Address, error) {
	return _Margin.Contract.Config(&_Margin.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Margin *MarginCallerSession) Config() (common.Address, error) {
	return _Margin.Contract.Config(&_Margin.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Margin *MarginCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Margin *MarginSession) Factory() (common.Address, error) {
	return _Margin.Contract.Factory(&_Margin.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Margin *MarginCallerSession) Factory() (common.Address, error) {
	return _Margin.Contract.Factory(&_Margin.CallOpts)
}

// GetNewLatestCPF is a free data retrieval call binding the contract method 0xcab25eea.
//
// Solidity: function getNewLatestCPF() view returns(int256)
func (_Margin *MarginCaller) GetNewLatestCPF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getNewLatestCPF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewLatestCPF is a free data retrieval call binding the contract method 0xcab25eea.
//
// Solidity: function getNewLatestCPF() view returns(int256)
func (_Margin *MarginSession) GetNewLatestCPF() (*big.Int, error) {
	return _Margin.Contract.GetNewLatestCPF(&_Margin.CallOpts)
}

// GetNewLatestCPF is a free data retrieval call binding the contract method 0xcab25eea.
//
// Solidity: function getNewLatestCPF() view returns(int256)
func (_Margin *MarginCallerSession) GetNewLatestCPF() (*big.Int, error) {
	return _Margin.Contract.GetNewLatestCPF(&_Margin.CallOpts)
}

// GetPosition is a free data retrieval call binding the contract method 0x16c19739.
//
// Solidity: function getPosition(address trader) view returns(int256, int256, uint256)
func (_Margin *MarginCaller) GetPosition(opts *bind.CallOpts, trader common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getPosition", trader)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPosition is a free data retrieval call binding the contract method 0x16c19739.
//
// Solidity: function getPosition(address trader) view returns(int256, int256, uint256)
func (_Margin *MarginSession) GetPosition(trader common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Margin.Contract.GetPosition(&_Margin.CallOpts, trader)
}

// GetPosition is a free data retrieval call binding the contract method 0x16c19739.
//
// Solidity: function getPosition(address trader) view returns(int256, int256, uint256)
func (_Margin *MarginCallerSession) GetPosition(trader common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Margin.Contract.GetPosition(&_Margin.CallOpts, trader)
}

// GetWithdrawable is a free data retrieval call binding the contract method 0x32cc6ae6.
//
// Solidity: function getWithdrawable(address trader) view returns(uint256 withdrawable)
func (_Margin *MarginCaller) GetWithdrawable(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getWithdrawable", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawable is a free data retrieval call binding the contract method 0x32cc6ae6.
//
// Solidity: function getWithdrawable(address trader) view returns(uint256 withdrawable)
func (_Margin *MarginSession) GetWithdrawable(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.GetWithdrawable(&_Margin.CallOpts, trader)
}

// GetWithdrawable is a free data retrieval call binding the contract method 0x32cc6ae6.
//
// Solidity: function getWithdrawable(address trader) view returns(uint256 withdrawable)
func (_Margin *MarginCallerSession) GetWithdrawable(trader common.Address) (*big.Int, error) {
	return _Margin.Contract.GetWithdrawable(&_Margin.CallOpts, trader)
}

// LastUpdateCPF is a free data retrieval call binding the contract method 0x85a5bd82.
//
// Solidity: function lastUpdateCPF() view returns(uint256)
func (_Margin *MarginCaller) LastUpdateCPF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "lastUpdateCPF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateCPF is a free data retrieval call binding the contract method 0x85a5bd82.
//
// Solidity: function lastUpdateCPF() view returns(uint256)
func (_Margin *MarginSession) LastUpdateCPF() (*big.Int, error) {
	return _Margin.Contract.LastUpdateCPF(&_Margin.CallOpts)
}

// LastUpdateCPF is a free data retrieval call binding the contract method 0x85a5bd82.
//
// Solidity: function lastUpdateCPF() view returns(uint256)
func (_Margin *MarginCallerSession) LastUpdateCPF() (*big.Int, error) {
	return _Margin.Contract.LastUpdateCPF(&_Margin.CallOpts)
}

// NetPosition is a free data retrieval call binding the contract method 0x833b0acb.
//
// Solidity: function netPosition() view returns(int256)
func (_Margin *MarginCaller) NetPosition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "netPosition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetPosition is a free data retrieval call binding the contract method 0x833b0acb.
//
// Solidity: function netPosition() view returns(int256)
func (_Margin *MarginSession) NetPosition() (*big.Int, error) {
	return _Margin.Contract.NetPosition(&_Margin.CallOpts)
}

// NetPosition is a free data retrieval call binding the contract method 0x833b0acb.
//
// Solidity: function netPosition() view returns(int256)
func (_Margin *MarginCallerSession) NetPosition() (*big.Int, error) {
	return _Margin.Contract.NetPosition(&_Margin.CallOpts)
}

// QuerySwapBaseWithAmm is a free data retrieval call binding the contract method 0x76dc1c09.
//
// Solidity: function querySwapBaseWithAmm(bool isLong, uint256 quoteAmount) view returns(uint256)
func (_Margin *MarginCaller) QuerySwapBaseWithAmm(opts *bind.CallOpts, isLong bool, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "querySwapBaseWithAmm", isLong, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySwapBaseWithAmm is a free data retrieval call binding the contract method 0x76dc1c09.
//
// Solidity: function querySwapBaseWithAmm(bool isLong, uint256 quoteAmount) view returns(uint256)
func (_Margin *MarginSession) QuerySwapBaseWithAmm(isLong bool, quoteAmount *big.Int) (*big.Int, error) {
	return _Margin.Contract.QuerySwapBaseWithAmm(&_Margin.CallOpts, isLong, quoteAmount)
}

// QuerySwapBaseWithAmm is a free data retrieval call binding the contract method 0x76dc1c09.
//
// Solidity: function querySwapBaseWithAmm(bool isLong, uint256 quoteAmount) view returns(uint256)
func (_Margin *MarginCallerSession) QuerySwapBaseWithAmm(isLong bool, quoteAmount *big.Int) (*big.Int, error) {
	return _Margin.Contract.QuerySwapBaseWithAmm(&_Margin.CallOpts, isLong, quoteAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_Margin *MarginCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_Margin *MarginSession) QuoteToken() (common.Address, error) {
	return _Margin.Contract.QuoteToken(&_Margin.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_Margin *MarginCallerSession) QuoteToken() (common.Address, error) {
	return _Margin.Contract.QuoteToken(&_Margin.CallOpts)
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() view returns(uint256)
func (_Margin *MarginCaller) Reserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "reserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() view returns(uint256)
func (_Margin *MarginSession) Reserve() (*big.Int, error) {
	return _Margin.Contract.Reserve(&_Margin.CallOpts)
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() view returns(uint256)
func (_Margin *MarginCallerSession) Reserve() (*big.Int, error) {
	return _Margin.Contract.Reserve(&_Margin.CallOpts)
}

// TotalQuoteLong is a free data retrieval call binding the contract method 0x65049cb5.
//
// Solidity: function totalQuoteLong() view returns(uint256)
func (_Margin *MarginCaller) TotalQuoteLong(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "totalQuoteLong")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalQuoteLong is a free data retrieval call binding the contract method 0x65049cb5.
//
// Solidity: function totalQuoteLong() view returns(uint256)
func (_Margin *MarginSession) TotalQuoteLong() (*big.Int, error) {
	return _Margin.Contract.TotalQuoteLong(&_Margin.CallOpts)
}

// TotalQuoteLong is a free data retrieval call binding the contract method 0x65049cb5.
//
// Solidity: function totalQuoteLong() view returns(uint256)
func (_Margin *MarginCallerSession) TotalQuoteLong() (*big.Int, error) {
	return _Margin.Contract.TotalQuoteLong(&_Margin.CallOpts)
}

// TotalQuoteShort is a free data retrieval call binding the contract method 0x92f4c2f2.
//
// Solidity: function totalQuoteShort() view returns(uint256)
func (_Margin *MarginCaller) TotalQuoteShort(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "totalQuoteShort")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalQuoteShort is a free data retrieval call binding the contract method 0x92f4c2f2.
//
// Solidity: function totalQuoteShort() view returns(uint256)
func (_Margin *MarginSession) TotalQuoteShort() (*big.Int, error) {
	return _Margin.Contract.TotalQuoteShort(&_Margin.CallOpts)
}

// TotalQuoteShort is a free data retrieval call binding the contract method 0x92f4c2f2.
//
// Solidity: function totalQuoteShort() view returns(uint256)
func (_Margin *MarginCallerSession) TotalQuoteShort() (*big.Int, error) {
	return _Margin.Contract.TotalQuoteShort(&_Margin.CallOpts)
}

// TraderCPF is a free data retrieval call binding the contract method 0x58557bb5.
//
// Solidity: function traderCPF(address ) view returns(int256)
func (_Margin *MarginCaller) TraderCPF(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "traderCPF", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TraderCPF is a free data retrieval call binding the contract method 0x58557bb5.
//
// Solidity: function traderCPF(address ) view returns(int256)
func (_Margin *MarginSession) TraderCPF(arg0 common.Address) (*big.Int, error) {
	return _Margin.Contract.TraderCPF(&_Margin.CallOpts, arg0)
}

// TraderCPF is a free data retrieval call binding the contract method 0x58557bb5.
//
// Solidity: function traderCPF(address ) view returns(int256)
func (_Margin *MarginCallerSession) TraderCPF(arg0 common.Address) (*big.Int, error) {
	return _Margin.Contract.TraderCPF(&_Margin.CallOpts, arg0)
}

// TraderPositionMap is a free data retrieval call binding the contract method 0x52ed67dc.
//
// Solidity: function traderPositionMap(address ) view returns(int256 quoteSize, int256 baseSize, uint256 tradeSize)
func (_Margin *MarginCaller) TraderPositionMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	QuoteSize *big.Int
	BaseSize  *big.Int
	TradeSize *big.Int
}, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "traderPositionMap", arg0)

	outstruct := new(struct {
		QuoteSize *big.Int
		BaseSize  *big.Int
		TradeSize *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.QuoteSize = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TradeSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TraderPositionMap is a free data retrieval call binding the contract method 0x52ed67dc.
//
// Solidity: function traderPositionMap(address ) view returns(int256 quoteSize, int256 baseSize, uint256 tradeSize)
func (_Margin *MarginSession) TraderPositionMap(arg0 common.Address) (struct {
	QuoteSize *big.Int
	BaseSize  *big.Int
	TradeSize *big.Int
}, error) {
	return _Margin.Contract.TraderPositionMap(&_Margin.CallOpts, arg0)
}

// TraderPositionMap is a free data retrieval call binding the contract method 0x52ed67dc.
//
// Solidity: function traderPositionMap(address ) view returns(int256 quoteSize, int256 baseSize, uint256 tradeSize)
func (_Margin *MarginCallerSession) TraderPositionMap(arg0 common.Address) (struct {
	QuoteSize *big.Int
	BaseSize  *big.Int
	TradeSize *big.Int
}, error) {
	return _Margin.Contract.TraderPositionMap(&_Margin.CallOpts, arg0)
}

// AddMargin is a paid mutator transaction binding the contract method 0xcf70cb69.
//
// Solidity: function addMargin(address trader, uint256 depositAmount) returns()
func (_Margin *MarginTransactor) AddMargin(opts *bind.TransactOpts, trader common.Address, depositAmount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "addMargin", trader, depositAmount)
}

// AddMargin is a paid mutator transaction binding the contract method 0xcf70cb69.
//
// Solidity: function addMargin(address trader, uint256 depositAmount) returns()
func (_Margin *MarginSession) AddMargin(trader common.Address, depositAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.AddMargin(&_Margin.TransactOpts, trader, depositAmount)
}

// AddMargin is a paid mutator transaction binding the contract method 0xcf70cb69.
//
// Solidity: function addMargin(address trader, uint256 depositAmount) returns()
func (_Margin *MarginTransactorSession) AddMargin(trader common.Address, depositAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.AddMargin(&_Margin.TransactOpts, trader, depositAmount)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x742fe1f8.
//
// Solidity: function closePosition(address trader, uint256 quoteAmount) returns(uint256 baseAmount)
func (_Margin *MarginTransactor) ClosePosition(opts *bind.TransactOpts, trader common.Address, quoteAmount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "closePosition", trader, quoteAmount)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x742fe1f8.
//
// Solidity: function closePosition(address trader, uint256 quoteAmount) returns(uint256 baseAmount)
func (_Margin *MarginSession) ClosePosition(trader common.Address, quoteAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.ClosePosition(&_Margin.TransactOpts, trader, quoteAmount)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x742fe1f8.
//
// Solidity: function closePosition(address trader, uint256 quoteAmount) returns(uint256 baseAmount)
func (_Margin *MarginTransactorSession) ClosePosition(trader common.Address, quoteAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.ClosePosition(&_Margin.TransactOpts, trader, quoteAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Margin *MarginTransactor) Deposit(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "deposit", user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Margin *MarginSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Deposit(&_Margin.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Margin *MarginTransactorSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Deposit(&_Margin.TransactOpts, user, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address baseToken_, address quoteToken_, address amm_) returns()
func (_Margin *MarginTransactor) Initialize(opts *bind.TransactOpts, baseToken_ common.Address, quoteToken_ common.Address, amm_ common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "initialize", baseToken_, quoteToken_, amm_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address baseToken_, address quoteToken_, address amm_) returns()
func (_Margin *MarginSession) Initialize(baseToken_ common.Address, quoteToken_ common.Address, amm_ common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Initialize(&_Margin.TransactOpts, baseToken_, quoteToken_, amm_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address baseToken_, address quoteToken_, address amm_) returns()
func (_Margin *MarginTransactorSession) Initialize(baseToken_ common.Address, quoteToken_ common.Address, amm_ common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Initialize(&_Margin.TransactOpts, baseToken_, quoteToken_, amm_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address trader, address to) returns(uint256 quoteAmount, uint256 baseAmount, uint256 bonus)
func (_Margin *MarginTransactor) Liquidate(opts *bind.TransactOpts, trader common.Address, to common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "liquidate", trader, to)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address trader, address to) returns(uint256 quoteAmount, uint256 baseAmount, uint256 bonus)
func (_Margin *MarginSession) Liquidate(trader common.Address, to common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Liquidate(&_Margin.TransactOpts, trader, to)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address trader, address to) returns(uint256 quoteAmount, uint256 baseAmount, uint256 bonus)
func (_Margin *MarginTransactorSession) Liquidate(trader common.Address, to common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Liquidate(&_Margin.TransactOpts, trader, to)
}

// OpenPosition is a paid mutator transaction binding the contract method 0xd67fb553.
//
// Solidity: function openPosition(address trader, uint8 side, uint256 quoteAmount) returns(uint256 baseAmount)
func (_Margin *MarginTransactor) OpenPosition(opts *bind.TransactOpts, trader common.Address, side uint8, quoteAmount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "openPosition", trader, side, quoteAmount)
}

// OpenPosition is a paid mutator transaction binding the contract method 0xd67fb553.
//
// Solidity: function openPosition(address trader, uint8 side, uint256 quoteAmount) returns(uint256 baseAmount)
func (_Margin *MarginSession) OpenPosition(trader common.Address, side uint8, quoteAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.OpenPosition(&_Margin.TransactOpts, trader, side, quoteAmount)
}

// OpenPosition is a paid mutator transaction binding the contract method 0xd67fb553.
//
// Solidity: function openPosition(address trader, uint8 side, uint256 quoteAmount) returns(uint256 baseAmount)
func (_Margin *MarginTransactorSession) OpenPosition(trader common.Address, side uint8, quoteAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.OpenPosition(&_Margin.TransactOpts, trader, side, quoteAmount)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xa8fe8f81.
//
// Solidity: function removeMargin(address trader, address to, uint256 withdrawAmount) returns()
func (_Margin *MarginTransactor) RemoveMargin(opts *bind.TransactOpts, trader common.Address, to common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "removeMargin", trader, to, withdrawAmount)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xa8fe8f81.
//
// Solidity: function removeMargin(address trader, address to, uint256 withdrawAmount) returns()
func (_Margin *MarginSession) RemoveMargin(trader common.Address, to common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.RemoveMargin(&_Margin.TransactOpts, trader, to, withdrawAmount)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xa8fe8f81.
//
// Solidity: function removeMargin(address trader, address to, uint256 withdrawAmount) returns()
func (_Margin *MarginTransactorSession) RemoveMargin(trader common.Address, to common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.RemoveMargin(&_Margin.TransactOpts, trader, to, withdrawAmount)
}

// UpdateCPF is a paid mutator transaction binding the contract method 0x84b4ce78.
//
// Solidity: function updateCPF() returns(int256 newLatestCPF)
func (_Margin *MarginTransactor) UpdateCPF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "updateCPF")
}

// UpdateCPF is a paid mutator transaction binding the contract method 0x84b4ce78.
//
// Solidity: function updateCPF() returns(int256 newLatestCPF)
func (_Margin *MarginSession) UpdateCPF() (*types.Transaction, error) {
	return _Margin.Contract.UpdateCPF(&_Margin.TransactOpts)
}

// UpdateCPF is a paid mutator transaction binding the contract method 0x84b4ce78.
//
// Solidity: function updateCPF() returns(int256 newLatestCPF)
func (_Margin *MarginTransactorSession) UpdateCPF() (*types.Transaction, error) {
	return _Margin.Contract.UpdateCPF(&_Margin.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address user, address receiver, uint256 amount) returns()
func (_Margin *MarginTransactor) Withdraw(opts *bind.TransactOpts, user common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "withdraw", user, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address user, address receiver, uint256 amount) returns()
func (_Margin *MarginSession) Withdraw(user common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Withdraw(&_Margin.TransactOpts, user, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address user, address receiver, uint256 amount) returns()
func (_Margin *MarginTransactorSession) Withdraw(user common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Withdraw(&_Margin.TransactOpts, user, receiver, amount)
}

// MarginAddMarginIterator is returned from FilterAddMargin and is used to iterate over the raw logs and unpacked data for AddMargin events raised by the Margin contract.
type MarginAddMarginIterator struct {
	Event *MarginAddMargin // Event containing the contract specifics and raw log

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
func (it *MarginAddMarginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginAddMargin)
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
		it.Event = new(MarginAddMargin)
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
func (it *MarginAddMarginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginAddMarginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginAddMargin represents a AddMargin event raised by the Margin contract.
type MarginAddMargin struct {
	Trader        common.Address
	DepositAmount *big.Int
	Position      IMarginPosition
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddMargin is a free log retrieval operation binding the contract event 0x2c8c31a18cfdc333bcf992ef4f59b0a2a92084c08f0e20497957abcb6c10bfcb.
//
// Solidity: event AddMargin(address indexed trader, uint256 depositAmount, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) FilterAddMargin(opts *bind.FilterOpts, trader []common.Address) (*MarginAddMarginIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "AddMargin", traderRule)
	if err != nil {
		return nil, err
	}
	return &MarginAddMarginIterator{contract: _Margin.contract, event: "AddMargin", logs: logs, sub: sub}, nil
}

// WatchAddMargin is a free log subscription operation binding the contract event 0x2c8c31a18cfdc333bcf992ef4f59b0a2a92084c08f0e20497957abcb6c10bfcb.
//
// Solidity: event AddMargin(address indexed trader, uint256 depositAmount, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) WatchAddMargin(opts *bind.WatchOpts, sink chan<- *MarginAddMargin, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "AddMargin", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginAddMargin)
				if err := _Margin.contract.UnpackLog(event, "AddMargin", log); err != nil {
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

// ParseAddMargin is a log parse operation binding the contract event 0x2c8c31a18cfdc333bcf992ef4f59b0a2a92084c08f0e20497957abcb6c10bfcb.
//
// Solidity: event AddMargin(address indexed trader, uint256 depositAmount, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) ParseAddMargin(log types.Log) (*MarginAddMargin, error) {
	event := new(MarginAddMargin)
	if err := _Margin.contract.UnpackLog(event, "AddMargin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginClosePositionIterator is returned from FilterClosePosition and is used to iterate over the raw logs and unpacked data for ClosePosition events raised by the Margin contract.
type MarginClosePositionIterator struct {
	Event *MarginClosePosition // Event containing the contract specifics and raw log

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
func (it *MarginClosePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginClosePosition)
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
		it.Event = new(MarginClosePosition)
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
func (it *MarginClosePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginClosePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginClosePosition represents a ClosePosition event raised by the Margin contract.
type MarginClosePosition struct {
	Trader      common.Address
	QuoteAmount *big.Int
	BaseAmount  *big.Int
	FundingFee  *big.Int
	Position    IMarginPosition
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClosePosition is a free log retrieval operation binding the contract event 0x9b383459b14df5363d3d1327ae34de16258b94a771c13244410e22aa11787e25.
//
// Solidity: event ClosePosition(address indexed trader, uint256 quoteAmount, uint256 baseAmount, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) FilterClosePosition(opts *bind.FilterOpts, trader []common.Address) (*MarginClosePositionIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "ClosePosition", traderRule)
	if err != nil {
		return nil, err
	}
	return &MarginClosePositionIterator{contract: _Margin.contract, event: "ClosePosition", logs: logs, sub: sub}, nil
}

// WatchClosePosition is a free log subscription operation binding the contract event 0x9b383459b14df5363d3d1327ae34de16258b94a771c13244410e22aa11787e25.
//
// Solidity: event ClosePosition(address indexed trader, uint256 quoteAmount, uint256 baseAmount, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) WatchClosePosition(opts *bind.WatchOpts, sink chan<- *MarginClosePosition, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "ClosePosition", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginClosePosition)
				if err := _Margin.contract.UnpackLog(event, "ClosePosition", log); err != nil {
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

// ParseClosePosition is a log parse operation binding the contract event 0x9b383459b14df5363d3d1327ae34de16258b94a771c13244410e22aa11787e25.
//
// Solidity: event ClosePosition(address indexed trader, uint256 quoteAmount, uint256 baseAmount, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) ParseClosePosition(log types.Log) (*MarginClosePosition, error) {
	event := new(MarginClosePosition)
	if err := _Margin.contract.UnpackLog(event, "ClosePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Margin contract.
type MarginDepositIterator struct {
	Event *MarginDeposit // Event containing the contract specifics and raw log

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
func (it *MarginDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginDeposit)
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
		it.Event = new(MarginDeposit)
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
func (it *MarginDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginDeposit represents a Deposit event raised by the Margin contract.
type MarginDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Margin *MarginFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*MarginDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &MarginDepositIterator{contract: _Margin.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Margin *MarginFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MarginDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginDeposit)
				if err := _Margin.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Margin *MarginFilterer) ParseDeposit(log types.Log) (*MarginDeposit, error) {
	event := new(MarginDeposit)
	if err := _Margin.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Margin contract.
type MarginLiquidateIterator struct {
	Event *MarginLiquidate // Event containing the contract specifics and raw log

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
func (it *MarginLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginLiquidate)
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
		it.Event = new(MarginLiquidate)
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
func (it *MarginLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginLiquidate represents a Liquidate event raised by the Margin contract.
type MarginLiquidate struct {
	Liquidator  common.Address
	Trader      common.Address
	To          common.Address
	QuoteAmount *big.Int
	BaseAmount  *big.Int
	Bonus       *big.Int
	FundingFee  *big.Int
	Position    IMarginPosition
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x9d24c3dea12298e139fff56a08afeb2d81d8d41351677d2e21c4dad89eda693f.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed trader, address indexed to, uint256 quoteAmount, uint256 baseAmount, uint256 bonus, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, trader []common.Address, to []common.Address) (*MarginLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "Liquidate", liquidatorRule, traderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarginLiquidateIterator{contract: _Margin.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x9d24c3dea12298e139fff56a08afeb2d81d8d41351677d2e21c4dad89eda693f.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed trader, address indexed to, uint256 quoteAmount, uint256 baseAmount, uint256 bonus, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *MarginLiquidate, liquidator []common.Address, trader []common.Address, to []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "Liquidate", liquidatorRule, traderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginLiquidate)
				if err := _Margin.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x9d24c3dea12298e139fff56a08afeb2d81d8d41351677d2e21c4dad89eda693f.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed trader, address indexed to, uint256 quoteAmount, uint256 baseAmount, uint256 bonus, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) ParseLiquidate(log types.Log) (*MarginLiquidate, error) {
	event := new(MarginLiquidate)
	if err := _Margin.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginOpenPositionIterator is returned from FilterOpenPosition and is used to iterate over the raw logs and unpacked data for OpenPosition events raised by the Margin contract.
type MarginOpenPositionIterator struct {
	Event *MarginOpenPosition // Event containing the contract specifics and raw log

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
func (it *MarginOpenPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginOpenPosition)
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
		it.Event = new(MarginOpenPosition)
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
func (it *MarginOpenPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginOpenPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginOpenPosition represents a OpenPosition event raised by the Margin contract.
type MarginOpenPosition struct {
	Trader      common.Address
	Side        uint8
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	FundingFee  *big.Int
	Position    IMarginPosition
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpenPosition is a free log retrieval operation binding the contract event 0xcc7ddfa701c3ec1faad0e3717271cf26abfbf7d48444448921bb9fbc27b26fd7.
//
// Solidity: event OpenPosition(address indexed trader, uint8 side, uint256 baseAmount, uint256 quoteAmount, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) FilterOpenPosition(opts *bind.FilterOpts, trader []common.Address) (*MarginOpenPositionIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "OpenPosition", traderRule)
	if err != nil {
		return nil, err
	}
	return &MarginOpenPositionIterator{contract: _Margin.contract, event: "OpenPosition", logs: logs, sub: sub}, nil
}

// WatchOpenPosition is a free log subscription operation binding the contract event 0xcc7ddfa701c3ec1faad0e3717271cf26abfbf7d48444448921bb9fbc27b26fd7.
//
// Solidity: event OpenPosition(address indexed trader, uint8 side, uint256 baseAmount, uint256 quoteAmount, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) WatchOpenPosition(opts *bind.WatchOpts, sink chan<- *MarginOpenPosition, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "OpenPosition", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginOpenPosition)
				if err := _Margin.contract.UnpackLog(event, "OpenPosition", log); err != nil {
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

// ParseOpenPosition is a log parse operation binding the contract event 0xcc7ddfa701c3ec1faad0e3717271cf26abfbf7d48444448921bb9fbc27b26fd7.
//
// Solidity: event OpenPosition(address indexed trader, uint8 side, uint256 baseAmount, uint256 quoteAmount, int256 fundingFee, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) ParseOpenPosition(log types.Log) (*MarginOpenPosition, error) {
	event := new(MarginOpenPosition)
	if err := _Margin.contract.UnpackLog(event, "OpenPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginRemoveMarginIterator is returned from FilterRemoveMargin and is used to iterate over the raw logs and unpacked data for RemoveMargin events raised by the Margin contract.
type MarginRemoveMarginIterator struct {
	Event *MarginRemoveMargin // Event containing the contract specifics and raw log

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
func (it *MarginRemoveMarginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginRemoveMargin)
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
		it.Event = new(MarginRemoveMargin)
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
func (it *MarginRemoveMarginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginRemoveMarginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginRemoveMargin represents a RemoveMargin event raised by the Margin contract.
type MarginRemoveMargin struct {
	Trader                   common.Address
	To                       common.Address
	WithdrawAmount           *big.Int
	FundingFee               *big.Int
	WithdrawAmountFromMargin *big.Int
	Position                 IMarginPosition
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRemoveMargin is a free log retrieval operation binding the contract event 0xacbeac375d5865b32775ef1296dac3dc42068eb9098044daf3e121176f4364db.
//
// Solidity: event RemoveMargin(address indexed trader, address indexed to, uint256 withdrawAmount, int256 fundingFee, uint256 withdrawAmountFromMargin, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) FilterRemoveMargin(opts *bind.FilterOpts, trader []common.Address, to []common.Address) (*MarginRemoveMarginIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "RemoveMargin", traderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarginRemoveMarginIterator{contract: _Margin.contract, event: "RemoveMargin", logs: logs, sub: sub}, nil
}

// WatchRemoveMargin is a free log subscription operation binding the contract event 0xacbeac375d5865b32775ef1296dac3dc42068eb9098044daf3e121176f4364db.
//
// Solidity: event RemoveMargin(address indexed trader, address indexed to, uint256 withdrawAmount, int256 fundingFee, uint256 withdrawAmountFromMargin, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) WatchRemoveMargin(opts *bind.WatchOpts, sink chan<- *MarginRemoveMargin, trader []common.Address, to []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "RemoveMargin", traderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginRemoveMargin)
				if err := _Margin.contract.UnpackLog(event, "RemoveMargin", log); err != nil {
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

// ParseRemoveMargin is a log parse operation binding the contract event 0xacbeac375d5865b32775ef1296dac3dc42068eb9098044daf3e121176f4364db.
//
// Solidity: event RemoveMargin(address indexed trader, address indexed to, uint256 withdrawAmount, int256 fundingFee, uint256 withdrawAmountFromMargin, (int256,int256,uint256) position)
func (_Margin *MarginFilterer) ParseRemoveMargin(log types.Log) (*MarginRemoveMargin, error) {
	event := new(MarginRemoveMargin)
	if err := _Margin.contract.UnpackLog(event, "RemoveMargin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginUpdateCPFIterator is returned from FilterUpdateCPF and is used to iterate over the raw logs and unpacked data for UpdateCPF events raised by the Margin contract.
type MarginUpdateCPFIterator struct {
	Event *MarginUpdateCPF // Event containing the contract specifics and raw log

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
func (it *MarginUpdateCPFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginUpdateCPF)
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
		it.Event = new(MarginUpdateCPF)
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
func (it *MarginUpdateCPFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginUpdateCPFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginUpdateCPF represents a UpdateCPF event raised by the Margin contract.
type MarginUpdateCPF struct {
	TimeStamp *big.Int
	Cpf       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateCPF is a free log retrieval operation binding the contract event 0xbb9790c7d61d850cbc07d4e95659d1765b0907b4ed044959265472f776ec1a8c.
//
// Solidity: event UpdateCPF(uint256 timeStamp, int256 cpf)
func (_Margin *MarginFilterer) FilterUpdateCPF(opts *bind.FilterOpts) (*MarginUpdateCPFIterator, error) {

	logs, sub, err := _Margin.contract.FilterLogs(opts, "UpdateCPF")
	if err != nil {
		return nil, err
	}
	return &MarginUpdateCPFIterator{contract: _Margin.contract, event: "UpdateCPF", logs: logs, sub: sub}, nil
}

// WatchUpdateCPF is a free log subscription operation binding the contract event 0xbb9790c7d61d850cbc07d4e95659d1765b0907b4ed044959265472f776ec1a8c.
//
// Solidity: event UpdateCPF(uint256 timeStamp, int256 cpf)
func (_Margin *MarginFilterer) WatchUpdateCPF(opts *bind.WatchOpts, sink chan<- *MarginUpdateCPF) (event.Subscription, error) {

	logs, sub, err := _Margin.contract.WatchLogs(opts, "UpdateCPF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginUpdateCPF)
				if err := _Margin.contract.UnpackLog(event, "UpdateCPF", log); err != nil {
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

// ParseUpdateCPF is a log parse operation binding the contract event 0xbb9790c7d61d850cbc07d4e95659d1765b0907b4ed044959265472f776ec1a8c.
//
// Solidity: event UpdateCPF(uint256 timeStamp, int256 cpf)
func (_Margin *MarginFilterer) ParseUpdateCPF(log types.Log) (*MarginUpdateCPF, error) {
	event := new(MarginUpdateCPF)
	if err := _Margin.contract.UnpackLog(event, "UpdateCPF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Margin contract.
type MarginWithdrawIterator struct {
	Event *MarginWithdraw // Event containing the contract specifics and raw log

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
func (it *MarginWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginWithdraw)
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
		it.Event = new(MarginWithdraw)
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
func (it *MarginWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginWithdraw represents a Withdraw event raised by the Margin contract.
type MarginWithdraw struct {
	User     common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed receiver, uint256 amount)
func (_Margin *MarginFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, receiver []common.Address) (*MarginWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "Withdraw", userRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MarginWithdrawIterator{contract: _Margin.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed receiver, uint256 amount)
func (_Margin *MarginFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MarginWithdraw, user []common.Address, receiver []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "Withdraw", userRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginWithdraw)
				if err := _Margin.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed receiver, uint256 amount)
func (_Margin *MarginFilterer) ParseWithdraw(log types.Log) (*MarginWithdraw, error) {
	event := new(MarginWithdraw)
	if err := _Margin.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
