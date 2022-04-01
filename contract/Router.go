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

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"config_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pairFactory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pcvTreasury_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pcv\",\"type\":\"bool\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pcv\",\"type\":\"bool\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"autoWithdraw\",\"type\":\"bool\"}],\"name\":\"closePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"closePositionETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"baseSize\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"quoteSize\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"getQuoteAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"getWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmountLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openPositionETHWithWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmountLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openPositionWithMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"marginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmountLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openPositionWithWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pcvTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLastOperation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Router *RouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Router *RouterSession) WETH() (common.Address, error) {
	return _Router.Contract.WETH(&_Router.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Router *RouterCallerSession) WETH() (common.Address, error) {
	return _Router.Contract.WETH(&_Router.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Router *RouterCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Router *RouterSession) Config() (common.Address, error) {
	return _Router.Contract.Config(&_Router.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Router *RouterCallerSession) Config() (common.Address, error) {
	return _Router.Contract.Config(&_Router.CallOpts)
}

// GetPosition is a free data retrieval call binding the contract method 0x713390f5.
//
// Solidity: function getPosition(address baseToken, address quoteToken, address holder) view returns(int256 baseSize, int256 quoteSize, uint256 tradeSize)
func (_Router *RouterCaller) GetPosition(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, holder common.Address) (struct {
	BaseSize  *big.Int
	QuoteSize *big.Int
	TradeSize *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getPosition", baseToken, quoteToken, holder)

	outstruct := new(struct {
		BaseSize  *big.Int
		QuoteSize *big.Int
		TradeSize *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseSize = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuoteSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TradeSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPosition is a free data retrieval call binding the contract method 0x713390f5.
//
// Solidity: function getPosition(address baseToken, address quoteToken, address holder) view returns(int256 baseSize, int256 quoteSize, uint256 tradeSize)
func (_Router *RouterSession) GetPosition(baseToken common.Address, quoteToken common.Address, holder common.Address) (struct {
	BaseSize  *big.Int
	QuoteSize *big.Int
	TradeSize *big.Int
}, error) {
	return _Router.Contract.GetPosition(&_Router.CallOpts, baseToken, quoteToken, holder)
}

// GetPosition is a free data retrieval call binding the contract method 0x713390f5.
//
// Solidity: function getPosition(address baseToken, address quoteToken, address holder) view returns(int256 baseSize, int256 quoteSize, uint256 tradeSize)
func (_Router *RouterCallerSession) GetPosition(baseToken common.Address, quoteToken common.Address, holder common.Address) (struct {
	BaseSize  *big.Int
	QuoteSize *big.Int
	TradeSize *big.Int
}, error) {
	return _Router.Contract.GetPosition(&_Router.CallOpts, baseToken, quoteToken, holder)
}

// GetQuoteAmount is a free data retrieval call binding the contract method 0xab46b4a6.
//
// Solidity: function getQuoteAmount(address baseToken, address quoteToken, uint8 side, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_Router *RouterCaller) GetQuoteAmount(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, side uint8, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getQuoteAmount", baseToken, quoteToken, side, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteAmount is a free data retrieval call binding the contract method 0xab46b4a6.
//
// Solidity: function getQuoteAmount(address baseToken, address quoteToken, uint8 side, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_Router *RouterSession) GetQuoteAmount(baseToken common.Address, quoteToken common.Address, side uint8, baseAmount *big.Int) (*big.Int, error) {
	return _Router.Contract.GetQuoteAmount(&_Router.CallOpts, baseToken, quoteToken, side, baseAmount)
}

// GetQuoteAmount is a free data retrieval call binding the contract method 0xab46b4a6.
//
// Solidity: function getQuoteAmount(address baseToken, address quoteToken, uint8 side, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_Router *RouterCallerSession) GetQuoteAmount(baseToken common.Address, quoteToken common.Address, side uint8, baseAmount *big.Int) (*big.Int, error) {
	return _Router.Contract.GetQuoteAmount(&_Router.CallOpts, baseToken, quoteToken, side, baseAmount)
}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address baseToken, address quoteToken) view returns(uint256 reserveBase, uint256 reserveQuote)
func (_Router *RouterCaller) GetReserves(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address) (struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getReserves", baseToken, quoteToken)

	outstruct := new(struct {
		ReserveBase  *big.Int
		ReserveQuote *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveQuote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address baseToken, address quoteToken) view returns(uint256 reserveBase, uint256 reserveQuote)
func (_Router *RouterSession) GetReserves(baseToken common.Address, quoteToken common.Address) (struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
}, error) {
	return _Router.Contract.GetReserves(&_Router.CallOpts, baseToken, quoteToken)
}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address baseToken, address quoteToken) view returns(uint256 reserveBase, uint256 reserveQuote)
func (_Router *RouterCallerSession) GetReserves(baseToken common.Address, quoteToken common.Address) (struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
}, error) {
	return _Router.Contract.GetReserves(&_Router.CallOpts, baseToken, quoteToken)
}

// GetWithdrawable is a free data retrieval call binding the contract method 0x05ea8223.
//
// Solidity: function getWithdrawable(address baseToken, address quoteToken, address holder) view returns(uint256 amount)
func (_Router *RouterCaller) GetWithdrawable(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, holder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getWithdrawable", baseToken, quoteToken, holder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawable is a free data retrieval call binding the contract method 0x05ea8223.
//
// Solidity: function getWithdrawable(address baseToken, address quoteToken, address holder) view returns(uint256 amount)
func (_Router *RouterSession) GetWithdrawable(baseToken common.Address, quoteToken common.Address, holder common.Address) (*big.Int, error) {
	return _Router.Contract.GetWithdrawable(&_Router.CallOpts, baseToken, quoteToken, holder)
}

// GetWithdrawable is a free data retrieval call binding the contract method 0x05ea8223.
//
// Solidity: function getWithdrawable(address baseToken, address quoteToken, address holder) view returns(uint256 amount)
func (_Router *RouterCallerSession) GetWithdrawable(baseToken common.Address, quoteToken common.Address, holder common.Address) (*big.Int, error) {
	return _Router.Contract.GetWithdrawable(&_Router.CallOpts, baseToken, quoteToken, holder)
}

// PairFactory is a free data retrieval call binding the contract method 0xe14f870d.
//
// Solidity: function pairFactory() view returns(address)
func (_Router *RouterCaller) PairFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pairFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFactory is a free data retrieval call binding the contract method 0xe14f870d.
//
// Solidity: function pairFactory() view returns(address)
func (_Router *RouterSession) PairFactory() (common.Address, error) {
	return _Router.Contract.PairFactory(&_Router.CallOpts)
}

// PairFactory is a free data retrieval call binding the contract method 0xe14f870d.
//
// Solidity: function pairFactory() view returns(address)
func (_Router *RouterCallerSession) PairFactory() (common.Address, error) {
	return _Router.Contract.PairFactory(&_Router.CallOpts)
}

// PcvTreasury is a free data retrieval call binding the contract method 0x4e9cacd9.
//
// Solidity: function pcvTreasury() view returns(address)
func (_Router *RouterCaller) PcvTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pcvTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PcvTreasury is a free data retrieval call binding the contract method 0x4e9cacd9.
//
// Solidity: function pcvTreasury() view returns(address)
func (_Router *RouterSession) PcvTreasury() (common.Address, error) {
	return _Router.Contract.PcvTreasury(&_Router.CallOpts)
}

// PcvTreasury is a free data retrieval call binding the contract method 0x4e9cacd9.
//
// Solidity: function pcvTreasury() view returns(address)
func (_Router *RouterCallerSession) PcvTreasury() (common.Address, error) {
	return _Router.Contract.PcvTreasury(&_Router.CallOpts)
}

// UserLastOperation is a free data retrieval call binding the contract method 0x9f9e28c6.
//
// Solidity: function userLastOperation(address , address ) view returns(uint256)
func (_Router *RouterCaller) UserLastOperation(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "userLastOperation", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastOperation is a free data retrieval call binding the contract method 0x9f9e28c6.
//
// Solidity: function userLastOperation(address , address ) view returns(uint256)
func (_Router *RouterSession) UserLastOperation(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Router.Contract.UserLastOperation(&_Router.CallOpts, arg0, arg1)
}

// UserLastOperation is a free data retrieval call binding the contract method 0x9f9e28c6.
//
// Solidity: function userLastOperation(address , address ) view returns(uint256)
func (_Router *RouterCallerSession) UserLastOperation(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Router.Contract.UserLastOperation(&_Router.CallOpts, arg0, arg1)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xfbf0f0e3.
//
// Solidity: function addLiquidity(address baseToken, address quoteToken, uint256 baseAmount, uint256 quoteAmountMin, uint256 deadline, bool pcv) returns(uint256 quoteAmount, uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidity(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, baseAmount *big.Int, quoteAmountMin *big.Int, deadline *big.Int, pcv bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity", baseToken, quoteToken, baseAmount, quoteAmountMin, deadline, pcv)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xfbf0f0e3.
//
// Solidity: function addLiquidity(address baseToken, address quoteToken, uint256 baseAmount, uint256 quoteAmountMin, uint256 deadline, bool pcv) returns(uint256 quoteAmount, uint256 liquidity)
func (_Router *RouterSession) AddLiquidity(baseToken common.Address, quoteToken common.Address, baseAmount *big.Int, quoteAmountMin *big.Int, deadline *big.Int, pcv bool) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, baseToken, quoteToken, baseAmount, quoteAmountMin, deadline, pcv)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xfbf0f0e3.
//
// Solidity: function addLiquidity(address baseToken, address quoteToken, uint256 baseAmount, uint256 quoteAmountMin, uint256 deadline, bool pcv) returns(uint256 quoteAmount, uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidity(baseToken common.Address, quoteToken common.Address, baseAmount *big.Int, quoteAmountMin *big.Int, deadline *big.Int, pcv bool) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, baseToken, quoteToken, baseAmount, quoteAmountMin, deadline, pcv)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x0568af12.
//
// Solidity: function addLiquidityETH(address quoteToken, uint256 quoteAmountMin, uint256 deadline, bool pcv) payable returns(uint256 ethAmount, uint256 quoteAmount, uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, quoteToken common.Address, quoteAmountMin *big.Int, deadline *big.Int, pcv bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidityETH", quoteToken, quoteAmountMin, deadline, pcv)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x0568af12.
//
// Solidity: function addLiquidityETH(address quoteToken, uint256 quoteAmountMin, uint256 deadline, bool pcv) payable returns(uint256 ethAmount, uint256 quoteAmount, uint256 liquidity)
func (_Router *RouterSession) AddLiquidityETH(quoteToken common.Address, quoteAmountMin *big.Int, deadline *big.Int, pcv bool) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityETH(&_Router.TransactOpts, quoteToken, quoteAmountMin, deadline, pcv)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x0568af12.
//
// Solidity: function addLiquidityETH(address quoteToken, uint256 quoteAmountMin, uint256 deadline, bool pcv) payable returns(uint256 ethAmount, uint256 quoteAmount, uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidityETH(quoteToken common.Address, quoteAmountMin *big.Int, deadline *big.Int, pcv bool) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityETH(&_Router.TransactOpts, quoteToken, quoteAmountMin, deadline, pcv)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x7dbdc226.
//
// Solidity: function closePosition(address baseToken, address quoteToken, uint256 quoteAmount, uint256 deadline, bool autoWithdraw) returns(uint256 baseAmount, uint256 withdrawAmount)
func (_Router *RouterTransactor) ClosePosition(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, quoteAmount *big.Int, deadline *big.Int, autoWithdraw bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "closePosition", baseToken, quoteToken, quoteAmount, deadline, autoWithdraw)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x7dbdc226.
//
// Solidity: function closePosition(address baseToken, address quoteToken, uint256 quoteAmount, uint256 deadline, bool autoWithdraw) returns(uint256 baseAmount, uint256 withdrawAmount)
func (_Router *RouterSession) ClosePosition(baseToken common.Address, quoteToken common.Address, quoteAmount *big.Int, deadline *big.Int, autoWithdraw bool) (*types.Transaction, error) {
	return _Router.Contract.ClosePosition(&_Router.TransactOpts, baseToken, quoteToken, quoteAmount, deadline, autoWithdraw)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x7dbdc226.
//
// Solidity: function closePosition(address baseToken, address quoteToken, uint256 quoteAmount, uint256 deadline, bool autoWithdraw) returns(uint256 baseAmount, uint256 withdrawAmount)
func (_Router *RouterTransactorSession) ClosePosition(baseToken common.Address, quoteToken common.Address, quoteAmount *big.Int, deadline *big.Int, autoWithdraw bool) (*types.Transaction, error) {
	return _Router.Contract.ClosePosition(&_Router.TransactOpts, baseToken, quoteToken, quoteAmount, deadline, autoWithdraw)
}

// ClosePositionETH is a paid mutator transaction binding the contract method 0xa23e1aee.
//
// Solidity: function closePositionETH(address quoteToken, uint256 quoteAmount, uint256 deadline) returns(uint256 baseAmount, uint256 withdrawAmount)
func (_Router *RouterTransactor) ClosePositionETH(opts *bind.TransactOpts, quoteToken common.Address, quoteAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "closePositionETH", quoteToken, quoteAmount, deadline)
}

// ClosePositionETH is a paid mutator transaction binding the contract method 0xa23e1aee.
//
// Solidity: function closePositionETH(address quoteToken, uint256 quoteAmount, uint256 deadline) returns(uint256 baseAmount, uint256 withdrawAmount)
func (_Router *RouterSession) ClosePositionETH(quoteToken common.Address, quoteAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.ClosePositionETH(&_Router.TransactOpts, quoteToken, quoteAmount, deadline)
}

// ClosePositionETH is a paid mutator transaction binding the contract method 0xa23e1aee.
//
// Solidity: function closePositionETH(address quoteToken, uint256 quoteAmount, uint256 deadline) returns(uint256 baseAmount, uint256 withdrawAmount)
func (_Router *RouterTransactorSession) ClosePositionETH(quoteToken common.Address, quoteAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.ClosePositionETH(&_Router.TransactOpts, quoteToken, quoteAmount, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x0284c3f5.
//
// Solidity: function deposit(address baseToken, address quoteToken, address holder, uint256 amount) returns()
func (_Router *RouterTransactor) Deposit(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, holder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "deposit", baseToken, quoteToken, holder, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x0284c3f5.
//
// Solidity: function deposit(address baseToken, address quoteToken, address holder, uint256 amount) returns()
func (_Router *RouterSession) Deposit(baseToken common.Address, quoteToken common.Address, holder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Deposit(&_Router.TransactOpts, baseToken, quoteToken, holder, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x0284c3f5.
//
// Solidity: function deposit(address baseToken, address quoteToken, address holder, uint256 amount) returns()
func (_Router *RouterTransactorSession) Deposit(baseToken common.Address, quoteToken common.Address, holder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Deposit(&_Router.TransactOpts, baseToken, quoteToken, holder, amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x734029bf.
//
// Solidity: function depositETH(address quoteToken, address holder) payable returns()
func (_Router *RouterTransactor) DepositETH(opts *bind.TransactOpts, quoteToken common.Address, holder common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "depositETH", quoteToken, holder)
}

// DepositETH is a paid mutator transaction binding the contract method 0x734029bf.
//
// Solidity: function depositETH(address quoteToken, address holder) payable returns()
func (_Router *RouterSession) DepositETH(quoteToken common.Address, holder common.Address) (*types.Transaction, error) {
	return _Router.Contract.DepositETH(&_Router.TransactOpts, quoteToken, holder)
}

// DepositETH is a paid mutator transaction binding the contract method 0x734029bf.
//
// Solidity: function depositETH(address quoteToken, address holder) payable returns()
func (_Router *RouterTransactorSession) DepositETH(quoteToken common.Address, holder common.Address) (*types.Transaction, error) {
	return _Router.Contract.DepositETH(&_Router.TransactOpts, quoteToken, holder)
}

// Liquidate is a paid mutator transaction binding the contract method 0x4c9fdb4c.
//
// Solidity: function liquidate(address baseToken, address quoteToken, address trader, address to) returns(uint256 quoteAmount, uint256 baseAmount, uint256 bonus)
func (_Router *RouterTransactor) Liquidate(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, trader common.Address, to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "liquidate", baseToken, quoteToken, trader, to)
}

// Liquidate is a paid mutator transaction binding the contract method 0x4c9fdb4c.
//
// Solidity: function liquidate(address baseToken, address quoteToken, address trader, address to) returns(uint256 quoteAmount, uint256 baseAmount, uint256 bonus)
func (_Router *RouterSession) Liquidate(baseToken common.Address, quoteToken common.Address, trader common.Address, to common.Address) (*types.Transaction, error) {
	return _Router.Contract.Liquidate(&_Router.TransactOpts, baseToken, quoteToken, trader, to)
}

// Liquidate is a paid mutator transaction binding the contract method 0x4c9fdb4c.
//
// Solidity: function liquidate(address baseToken, address quoteToken, address trader, address to) returns(uint256 quoteAmount, uint256 baseAmount, uint256 bonus)
func (_Router *RouterTransactorSession) Liquidate(baseToken common.Address, quoteToken common.Address, trader common.Address, to common.Address) (*types.Transaction, error) {
	return _Router.Contract.Liquidate(&_Router.TransactOpts, baseToken, quoteToken, trader, to)
}

// OpenPositionETHWithWallet is a paid mutator transaction binding the contract method 0x374adbcc.
//
// Solidity: function openPositionETHWithWallet(address quoteToken, uint8 side, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) payable returns(uint256 baseAmount)
func (_Router *RouterTransactor) OpenPositionETHWithWallet(opts *bind.TransactOpts, quoteToken common.Address, side uint8, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "openPositionETHWithWallet", quoteToken, side, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionETHWithWallet is a paid mutator transaction binding the contract method 0x374adbcc.
//
// Solidity: function openPositionETHWithWallet(address quoteToken, uint8 side, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) payable returns(uint256 baseAmount)
func (_Router *RouterSession) OpenPositionETHWithWallet(quoteToken common.Address, side uint8, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.OpenPositionETHWithWallet(&_Router.TransactOpts, quoteToken, side, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionETHWithWallet is a paid mutator transaction binding the contract method 0x374adbcc.
//
// Solidity: function openPositionETHWithWallet(address quoteToken, uint8 side, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) payable returns(uint256 baseAmount)
func (_Router *RouterTransactorSession) OpenPositionETHWithWallet(quoteToken common.Address, side uint8, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.OpenPositionETHWithWallet(&_Router.TransactOpts, quoteToken, side, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionWithMargin is a paid mutator transaction binding the contract method 0x43c8d33b.
//
// Solidity: function openPositionWithMargin(address baseToken, address quoteToken, uint8 side, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) returns(uint256 baseAmount)
func (_Router *RouterTransactor) OpenPositionWithMargin(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, side uint8, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "openPositionWithMargin", baseToken, quoteToken, side, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionWithMargin is a paid mutator transaction binding the contract method 0x43c8d33b.
//
// Solidity: function openPositionWithMargin(address baseToken, address quoteToken, uint8 side, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) returns(uint256 baseAmount)
func (_Router *RouterSession) OpenPositionWithMargin(baseToken common.Address, quoteToken common.Address, side uint8, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.OpenPositionWithMargin(&_Router.TransactOpts, baseToken, quoteToken, side, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionWithMargin is a paid mutator transaction binding the contract method 0x43c8d33b.
//
// Solidity: function openPositionWithMargin(address baseToken, address quoteToken, uint8 side, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) returns(uint256 baseAmount)
func (_Router *RouterTransactorSession) OpenPositionWithMargin(baseToken common.Address, quoteToken common.Address, side uint8, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.OpenPositionWithMargin(&_Router.TransactOpts, baseToken, quoteToken, side, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionWithWallet is a paid mutator transaction binding the contract method 0x3dff89a7.
//
// Solidity: function openPositionWithWallet(address baseToken, address quoteToken, uint8 side, uint256 marginAmount, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) returns(uint256 baseAmount)
func (_Router *RouterTransactor) OpenPositionWithWallet(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, side uint8, marginAmount *big.Int, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "openPositionWithWallet", baseToken, quoteToken, side, marginAmount, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionWithWallet is a paid mutator transaction binding the contract method 0x3dff89a7.
//
// Solidity: function openPositionWithWallet(address baseToken, address quoteToken, uint8 side, uint256 marginAmount, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) returns(uint256 baseAmount)
func (_Router *RouterSession) OpenPositionWithWallet(baseToken common.Address, quoteToken common.Address, side uint8, marginAmount *big.Int, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.OpenPositionWithWallet(&_Router.TransactOpts, baseToken, quoteToken, side, marginAmount, quoteAmount, baseAmountLimit, deadline)
}

// OpenPositionWithWallet is a paid mutator transaction binding the contract method 0x3dff89a7.
//
// Solidity: function openPositionWithWallet(address baseToken, address quoteToken, uint8 side, uint256 marginAmount, uint256 quoteAmount, uint256 baseAmountLimit, uint256 deadline) returns(uint256 baseAmount)
func (_Router *RouterTransactorSession) OpenPositionWithWallet(baseToken common.Address, quoteToken common.Address, side uint8, marginAmount *big.Int, quoteAmount *big.Int, baseAmountLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.OpenPositionWithWallet(&_Router.TransactOpts, baseToken, quoteToken, side, marginAmount, quoteAmount, baseAmountLimit, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address baseToken, address quoteToken, uint256 liquidity, uint256 baseAmountMin, uint256 deadline) returns(uint256 baseAmount, uint256 quoteAmount)
func (_Router *RouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, liquidity *big.Int, baseAmountMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidity", baseToken, quoteToken, liquidity, baseAmountMin, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address baseToken, address quoteToken, uint256 liquidity, uint256 baseAmountMin, uint256 deadline) returns(uint256 baseAmount, uint256 quoteAmount)
func (_Router *RouterSession) RemoveLiquidity(baseToken common.Address, quoteToken common.Address, liquidity *big.Int, baseAmountMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, baseToken, quoteToken, liquidity, baseAmountMin, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address baseToken, address quoteToken, uint256 liquidity, uint256 baseAmountMin, uint256 deadline) returns(uint256 baseAmount, uint256 quoteAmount)
func (_Router *RouterTransactorSession) RemoveLiquidity(baseToken common.Address, quoteToken common.Address, liquidity *big.Int, baseAmountMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, baseToken, quoteToken, liquidity, baseAmountMin, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xa1cfacde.
//
// Solidity: function removeLiquidityETH(address quoteToken, uint256 liquidity, uint256 ethAmountMin, uint256 deadline) returns(uint256 ethAmount, uint256 quoteAmount)
func (_Router *RouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, quoteToken common.Address, liquidity *big.Int, ethAmountMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidityETH", quoteToken, liquidity, ethAmountMin, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xa1cfacde.
//
// Solidity: function removeLiquidityETH(address quoteToken, uint256 liquidity, uint256 ethAmountMin, uint256 deadline) returns(uint256 ethAmount, uint256 quoteAmount)
func (_Router *RouterSession) RemoveLiquidityETH(quoteToken common.Address, liquidity *big.Int, ethAmountMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityETH(&_Router.TransactOpts, quoteToken, liquidity, ethAmountMin, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xa1cfacde.
//
// Solidity: function removeLiquidityETH(address quoteToken, uint256 liquidity, uint256 ethAmountMin, uint256 deadline) returns(uint256 ethAmount, uint256 quoteAmount)
func (_Router *RouterTransactorSession) RemoveLiquidityETH(quoteToken common.Address, liquidity *big.Int, ethAmountMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityETH(&_Router.TransactOpts, quoteToken, liquidity, ethAmountMin, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address baseToken, address quoteToken, uint256 amount) returns()
func (_Router *RouterTransactor) Withdraw(opts *bind.TransactOpts, baseToken common.Address, quoteToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdraw", baseToken, quoteToken, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address baseToken, address quoteToken, uint256 amount) returns()
func (_Router *RouterSession) Withdraw(baseToken common.Address, quoteToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Withdraw(&_Router.TransactOpts, baseToken, quoteToken, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address baseToken, address quoteToken, uint256 amount) returns()
func (_Router *RouterTransactorSession) Withdraw(baseToken common.Address, quoteToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Withdraw(&_Router.TransactOpts, baseToken, quoteToken, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address quoteToken, uint256 amount) returns()
func (_Router *RouterTransactor) WithdrawETH(opts *bind.TransactOpts, quoteToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdrawETH", quoteToken, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address quoteToken, uint256 amount) returns()
func (_Router *RouterSession) WithdrawETH(quoteToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.WithdrawETH(&_Router.TransactOpts, quoteToken, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address quoteToken, uint256 amount) returns()
func (_Router *RouterTransactorSession) WithdrawETH(quoteToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.WithdrawETH(&_Router.TransactOpts, quoteToken, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}
