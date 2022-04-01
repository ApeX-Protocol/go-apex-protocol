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

// AmmMetaData contains all meta data concerning the Amm contract.
var AmmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"ForceSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteReserveBefore\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteReserveAfter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_baseReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteReserveFromInternal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteReserveFromExternal\",\"type\":\"uint256\"}],\"name\":\"Rebase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserveBase\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserveQuote\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"estimateSwap\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"forceSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRealBaseReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realBaseReserve\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserveBase\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserveQuote\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTheMaxBurnLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxLiquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"margin_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"margin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteReserveAfter\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AmmABI is the input ABI used to generate the binding from.
// Deprecated: Use AmmMetaData.ABI instead.
var AmmABI = AmmMetaData.ABI

// Amm is an auto generated Go binding around an Ethereum contract.
type Amm struct {
	AmmCaller     // Read-only binding to the contract
	AmmTransactor // Write-only binding to the contract
	AmmFilterer   // Log filterer for contract events
}

// AmmCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmmSession struct {
	Contract     *Amm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmmCallerSession struct {
	Contract *AmmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AmmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmmTransactorSession struct {
	Contract     *AmmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmmRaw struct {
	Contract *Amm // Generic contract binding to access the raw methods on
}

// AmmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmmCallerRaw struct {
	Contract *AmmCaller // Generic read-only contract binding to access the raw methods on
}

// AmmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmmTransactorRaw struct {
	Contract *AmmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmm creates a new instance of Amm, bound to a specific deployed contract.
func NewAmm(address common.Address, backend bind.ContractBackend) (*Amm, error) {
	contract, err := bindAmm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Amm{AmmCaller: AmmCaller{contract: contract}, AmmTransactor: AmmTransactor{contract: contract}, AmmFilterer: AmmFilterer{contract: contract}}, nil
}

// NewAmmCaller creates a new read-only instance of Amm, bound to a specific deployed contract.
func NewAmmCaller(address common.Address, caller bind.ContractCaller) (*AmmCaller, error) {
	contract, err := bindAmm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmmCaller{contract: contract}, nil
}

// NewAmmTransactor creates a new write-only instance of Amm, bound to a specific deployed contract.
func NewAmmTransactor(address common.Address, transactor bind.ContractTransactor) (*AmmTransactor, error) {
	contract, err := bindAmm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmmTransactor{contract: contract}, nil
}

// NewAmmFilterer creates a new log filterer instance of Amm, bound to a specific deployed contract.
func NewAmmFilterer(address common.Address, filterer bind.ContractFilterer) (*AmmFilterer, error) {
	contract, err := bindAmm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmmFilterer{contract: contract}, nil
}

// bindAmm binds a generic wrapper to an already deployed contract.
func bindAmm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Amm *AmmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Amm.Contract.AmmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Amm *AmmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.Contract.AmmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Amm *AmmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Amm.Contract.AmmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Amm *AmmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Amm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Amm *AmmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Amm *AmmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Amm.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Amm *AmmCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Amm *AmmSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Amm.Contract.DOMAINSEPARATOR(&_Amm.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Amm *AmmCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Amm.Contract.DOMAINSEPARATOR(&_Amm.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Amm *AmmCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Amm *AmmSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Amm.Contract.MINIMUMLIQUIDITY(&_Amm.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Amm *AmmCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Amm.Contract.MINIMUMLIQUIDITY(&_Amm.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Amm *AmmCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Amm *AmmSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Amm.Contract.PERMITTYPEHASH(&_Amm.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Amm *AmmCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Amm.Contract.PERMITTYPEHASH(&_Amm.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Amm *AmmCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Amm *AmmSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Amm.Contract.Allowance(&_Amm.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Amm *AmmCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Amm.Contract.Allowance(&_Amm.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Amm *AmmCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Amm *AmmSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Amm.Contract.BalanceOf(&_Amm.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Amm *AmmCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Amm.Contract.BalanceOf(&_Amm.CallOpts, arg0)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Amm *AmmCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Amm *AmmSession) BaseToken() (common.Address, error) {
	return _Amm.Contract.BaseToken(&_Amm.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Amm *AmmCallerSession) BaseToken() (common.Address, error) {
	return _Amm.Contract.BaseToken(&_Amm.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Amm *AmmCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Amm *AmmSession) Config() (common.Address, error) {
	return _Amm.Contract.Config(&_Amm.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Amm *AmmCallerSession) Config() (common.Address, error) {
	return _Amm.Contract.Config(&_Amm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Amm *AmmCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Amm *AmmSession) Decimals() (uint8, error) {
	return _Amm.Contract.Decimals(&_Amm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Amm *AmmCallerSession) Decimals() (uint8, error) {
	return _Amm.Contract.Decimals(&_Amm.CallOpts)
}

// EstimateSwap is a free data retrieval call binding the contract method 0xbc936576.
//
// Solidity: function estimateSwap(address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) view returns(uint256[2] amounts)
func (_Amm *AmmCaller) EstimateSwap(opts *bind.CallOpts, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "estimateSwap", inputToken, outputToken, inputAmount, outputAmount)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// EstimateSwap is a free data retrieval call binding the contract method 0xbc936576.
//
// Solidity: function estimateSwap(address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) view returns(uint256[2] amounts)
func (_Amm *AmmSession) EstimateSwap(inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) ([2]*big.Int, error) {
	return _Amm.Contract.EstimateSwap(&_Amm.CallOpts, inputToken, outputToken, inputAmount, outputAmount)
}

// EstimateSwap is a free data retrieval call binding the contract method 0xbc936576.
//
// Solidity: function estimateSwap(address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) view returns(uint256[2] amounts)
func (_Amm *AmmCallerSession) EstimateSwap(inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) ([2]*big.Int, error) {
	return _Amm.Contract.EstimateSwap(&_Amm.CallOpts, inputToken, outputToken, inputAmount, outputAmount)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Amm *AmmCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Amm *AmmSession) Factory() (common.Address, error) {
	return _Amm.Contract.Factory(&_Amm.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Amm *AmmCallerSession) Factory() (common.Address, error) {
	return _Amm.Contract.Factory(&_Amm.CallOpts)
}

// GetFeeLiquidity is a free data retrieval call binding the contract method 0x6f316f18.
//
// Solidity: function getFeeLiquidity() view returns(uint256)
func (_Amm *AmmCaller) GetFeeLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getFeeLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeLiquidity is a free data retrieval call binding the contract method 0x6f316f18.
//
// Solidity: function getFeeLiquidity() view returns(uint256)
func (_Amm *AmmSession) GetFeeLiquidity() (*big.Int, error) {
	return _Amm.Contract.GetFeeLiquidity(&_Amm.CallOpts)
}

// GetFeeLiquidity is a free data retrieval call binding the contract method 0x6f316f18.
//
// Solidity: function getFeeLiquidity() view returns(uint256)
func (_Amm *AmmCallerSession) GetFeeLiquidity() (*big.Int, error) {
	return _Amm.Contract.GetFeeLiquidity(&_Amm.CallOpts)
}

// GetRealBaseReserve is a free data retrieval call binding the contract method 0x92c00dbc.
//
// Solidity: function getRealBaseReserve() view returns(uint256 realBaseReserve)
func (_Amm *AmmCaller) GetRealBaseReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getRealBaseReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRealBaseReserve is a free data retrieval call binding the contract method 0x92c00dbc.
//
// Solidity: function getRealBaseReserve() view returns(uint256 realBaseReserve)
func (_Amm *AmmSession) GetRealBaseReserve() (*big.Int, error) {
	return _Amm.Contract.GetRealBaseReserve(&_Amm.CallOpts)
}

// GetRealBaseReserve is a free data retrieval call binding the contract method 0x92c00dbc.
//
// Solidity: function getRealBaseReserve() view returns(uint256 realBaseReserve)
func (_Amm *AmmCallerSession) GetRealBaseReserve() (*big.Int, error) {
	return _Amm.Contract.GetRealBaseReserve(&_Amm.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserveBase, uint112 reserveQuote, uint32 blockTimestamp)
func (_Amm *AmmCaller) GetReserves(opts *bind.CallOpts) (struct {
	ReserveBase    *big.Int
	ReserveQuote   *big.Int
	BlockTimestamp uint32
}, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		ReserveBase    *big.Int
		ReserveQuote   *big.Int
		BlockTimestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveQuote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserveBase, uint112 reserveQuote, uint32 blockTimestamp)
func (_Amm *AmmSession) GetReserves() (struct {
	ReserveBase    *big.Int
	ReserveQuote   *big.Int
	BlockTimestamp uint32
}, error) {
	return _Amm.Contract.GetReserves(&_Amm.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserveBase, uint112 reserveQuote, uint32 blockTimestamp)
func (_Amm *AmmCallerSession) GetReserves() (struct {
	ReserveBase    *big.Int
	ReserveQuote   *big.Int
	BlockTimestamp uint32
}, error) {
	return _Amm.Contract.GetReserves(&_Amm.CallOpts)
}

// GetTheMaxBurnLiquidity is a free data retrieval call binding the contract method 0x09daf930.
//
// Solidity: function getTheMaxBurnLiquidity() view returns(uint256 maxLiquidity)
func (_Amm *AmmCaller) GetTheMaxBurnLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getTheMaxBurnLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTheMaxBurnLiquidity is a free data retrieval call binding the contract method 0x09daf930.
//
// Solidity: function getTheMaxBurnLiquidity() view returns(uint256 maxLiquidity)
func (_Amm *AmmSession) GetTheMaxBurnLiquidity() (*big.Int, error) {
	return _Amm.Contract.GetTheMaxBurnLiquidity(&_Amm.CallOpts)
}

// GetTheMaxBurnLiquidity is a free data retrieval call binding the contract method 0x09daf930.
//
// Solidity: function getTheMaxBurnLiquidity() view returns(uint256 maxLiquidity)
func (_Amm *AmmCallerSession) GetTheMaxBurnLiquidity() (*big.Int, error) {
	return _Amm.Contract.GetTheMaxBurnLiquidity(&_Amm.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Amm *AmmCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Amm *AmmSession) KLast() (*big.Int, error) {
	return _Amm.Contract.KLast(&_Amm.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Amm *AmmCallerSession) KLast() (*big.Int, error) {
	return _Amm.Contract.KLast(&_Amm.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256)
func (_Amm *AmmCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "lastPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256)
func (_Amm *AmmSession) LastPrice() (*big.Int, error) {
	return _Amm.Contract.LastPrice(&_Amm.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256)
func (_Amm *AmmCallerSession) LastPrice() (*big.Int, error) {
	return _Amm.Contract.LastPrice(&_Amm.CallOpts)
}

// Margin is a free data retrieval call binding the contract method 0x8f76691a.
//
// Solidity: function margin() view returns(address)
func (_Amm *AmmCaller) Margin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "margin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Margin is a free data retrieval call binding the contract method 0x8f76691a.
//
// Solidity: function margin() view returns(address)
func (_Amm *AmmSession) Margin() (common.Address, error) {
	return _Amm.Contract.Margin(&_Amm.CallOpts)
}

// Margin is a free data retrieval call binding the contract method 0x8f76691a.
//
// Solidity: function margin() view returns(address)
func (_Amm *AmmCallerSession) Margin() (common.Address, error) {
	return _Amm.Contract.Margin(&_Amm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Amm *AmmCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Amm *AmmSession) Name() (string, error) {
	return _Amm.Contract.Name(&_Amm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Amm *AmmCallerSession) Name() (string, error) {
	return _Amm.Contract.Name(&_Amm.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Amm *AmmCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Amm *AmmSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Amm.Contract.Nonces(&_Amm.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Amm *AmmCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Amm.Contract.Nonces(&_Amm.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Amm *AmmCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Amm *AmmSession) Price0CumulativeLast() (*big.Int, error) {
	return _Amm.Contract.Price0CumulativeLast(&_Amm.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Amm *AmmCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Amm.Contract.Price0CumulativeLast(&_Amm.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Amm *AmmCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Amm *AmmSession) Price1CumulativeLast() (*big.Int, error) {
	return _Amm.Contract.Price1CumulativeLast(&_Amm.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Amm *AmmCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Amm.Contract.Price1CumulativeLast(&_Amm.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_Amm *AmmCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_Amm *AmmSession) QuoteToken() (common.Address, error) {
	return _Amm.Contract.QuoteToken(&_Amm.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_Amm *AmmCallerSession) QuoteToken() (common.Address, error) {
	return _Amm.Contract.QuoteToken(&_Amm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Amm *AmmCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Amm *AmmSession) Symbol() (string, error) {
	return _Amm.Contract.Symbol(&_Amm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Amm *AmmCallerSession) Symbol() (string, error) {
	return _Amm.Contract.Symbol(&_Amm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Amm *AmmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Amm *AmmSession) TotalSupply() (*big.Int, error) {
	return _Amm.Contract.TotalSupply(&_Amm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Amm *AmmCallerSession) TotalSupply() (*big.Int, error) {
	return _Amm.Contract.TotalSupply(&_Amm.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Amm *AmmTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Amm *AmmSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Approve(&_Amm.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Amm *AmmTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Approve(&_Amm.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Amm.Contract.Burn(&_Amm.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Amm.Contract.Burn(&_Amm.TransactOpts, to)
}

// ForceSwap is a paid mutator transaction binding the contract method 0x988370a3.
//
// Solidity: function forceSwap(address trader, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) returns()
func (_Amm *AmmTransactor) ForceSwap(opts *bind.TransactOpts, trader common.Address, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "forceSwap", trader, inputToken, outputToken, inputAmount, outputAmount)
}

// ForceSwap is a paid mutator transaction binding the contract method 0x988370a3.
//
// Solidity: function forceSwap(address trader, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) returns()
func (_Amm *AmmSession) ForceSwap(trader common.Address, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.ForceSwap(&_Amm.TransactOpts, trader, inputToken, outputToken, inputAmount, outputAmount)
}

// ForceSwap is a paid mutator transaction binding the contract method 0x988370a3.
//
// Solidity: function forceSwap(address trader, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) returns()
func (_Amm *AmmTransactorSession) ForceSwap(trader common.Address, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.ForceSwap(&_Amm.TransactOpts, trader, inputToken, outputToken, inputAmount, outputAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address baseToken_, address quoteToken_, address margin_) returns()
func (_Amm *AmmTransactor) Initialize(opts *bind.TransactOpts, baseToken_ common.Address, quoteToken_ common.Address, margin_ common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "initialize", baseToken_, quoteToken_, margin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address baseToken_, address quoteToken_, address margin_) returns()
func (_Amm *AmmSession) Initialize(baseToken_ common.Address, quoteToken_ common.Address, margin_ common.Address) (*types.Transaction, error) {
	return _Amm.Contract.Initialize(&_Amm.TransactOpts, baseToken_, quoteToken_, margin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address baseToken_, address quoteToken_, address margin_) returns()
func (_Amm *AmmTransactorSession) Initialize(baseToken_ common.Address, quoteToken_ common.Address, margin_ common.Address) (*types.Transaction, error) {
	return _Amm.Contract.Initialize(&_Amm.TransactOpts, baseToken_, quoteToken_, margin_)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Amm.Contract.Mint(&_Amm.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Amm.Contract.Mint(&_Amm.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Amm *AmmTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Amm *AmmSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Amm.Contract.Permit(&_Amm.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Amm *AmmTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Amm.Contract.Permit(&_Amm.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns(uint256 quoteReserveAfter)
func (_Amm *AmmTransactor) Rebase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "rebase")
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns(uint256 quoteReserveAfter)
func (_Amm *AmmSession) Rebase() (*types.Transaction, error) {
	return _Amm.Contract.Rebase(&_Amm.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns(uint256 quoteReserveAfter)
func (_Amm *AmmTransactorSession) Rebase() (*types.Transaction, error) {
	return _Amm.Contract.Rebase(&_Amm.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address trader, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) returns(uint256[2] amounts)
func (_Amm *AmmTransactor) Swap(opts *bind.TransactOpts, trader common.Address, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "swap", trader, inputToken, outputToken, inputAmount, outputAmount)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address trader, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) returns(uint256[2] amounts)
func (_Amm *AmmSession) Swap(trader common.Address, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Swap(&_Amm.TransactOpts, trader, inputToken, outputToken, inputAmount, outputAmount)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address trader, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount) returns(uint256[2] amounts)
func (_Amm *AmmTransactorSession) Swap(trader common.Address, inputToken common.Address, outputToken common.Address, inputAmount *big.Int, outputAmount *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Swap(&_Amm.TransactOpts, trader, inputToken, outputToken, inputAmount, outputAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Amm *AmmTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Amm *AmmSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Transfer(&_Amm.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Amm *AmmTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Transfer(&_Amm.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Amm *AmmTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Amm *AmmSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.TransferFrom(&_Amm.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Amm *AmmTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.TransferFrom(&_Amm.TransactOpts, from, to, value)
}

// AmmApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Amm contract.
type AmmApprovalIterator struct {
	Event *AmmApproval // Event containing the contract specifics and raw log

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
func (it *AmmApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmApproval)
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
		it.Event = new(AmmApproval)
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
func (it *AmmApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmApproval represents a Approval event raised by the Amm contract.
type AmmApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Amm *AmmFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AmmApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AmmApprovalIterator{contract: _Amm.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Amm *AmmFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AmmApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmApproval)
				if err := _Amm.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Amm *AmmFilterer) ParseApproval(log types.Log) (*AmmApproval, error) {
	event := new(AmmApproval)
	if err := _Amm.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Amm contract.
type AmmBurnIterator struct {
	Event *AmmBurn // Event containing the contract specifics and raw log

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
func (it *AmmBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmBurn)
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
		it.Event = new(AmmBurn)
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
func (it *AmmBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmBurn represents a Burn event raised by the Amm contract.
type AmmBurn struct {
	Sender      common.Address
	To          common.Address
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	Liquidity   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed sender, address indexed to, uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AmmBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AmmBurnIterator{contract: _Amm.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed sender, address indexed to, uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AmmBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmBurn)
				if err := _Amm.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed sender, address indexed to, uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmFilterer) ParseBurn(log types.Log) (*AmmBurn, error) {
	event := new(AmmBurn)
	if err := _Amm.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmForceSwapIterator is returned from FilterForceSwap and is used to iterate over the raw logs and unpacked data for ForceSwap events raised by the Amm contract.
type AmmForceSwapIterator struct {
	Event *AmmForceSwap // Event containing the contract specifics and raw log

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
func (it *AmmForceSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmForceSwap)
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
		it.Event = new(AmmForceSwap)
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
func (it *AmmForceSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmForceSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmForceSwap represents a ForceSwap event raised by the Amm contract.
type AmmForceSwap struct {
	Trader       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForceSwap is a free log retrieval operation binding the contract event 0xa36e6e2023c9daa81ac16fbc0d2499822bc9617326793b1ec3ab50c90bcfe872.
//
// Solidity: event ForceSwap(address indexed trader, address indexed inputToken, address indexed outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Amm *AmmFilterer) FilterForceSwap(opts *bind.FilterOpts, trader []common.Address, inputToken []common.Address, outputToken []common.Address) (*AmmForceSwapIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var inputTokenRule []interface{}
	for _, inputTokenItem := range inputToken {
		inputTokenRule = append(inputTokenRule, inputTokenItem)
	}
	var outputTokenRule []interface{}
	for _, outputTokenItem := range outputToken {
		outputTokenRule = append(outputTokenRule, outputTokenItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "ForceSwap", traderRule, inputTokenRule, outputTokenRule)
	if err != nil {
		return nil, err
	}
	return &AmmForceSwapIterator{contract: _Amm.contract, event: "ForceSwap", logs: logs, sub: sub}, nil
}

// WatchForceSwap is a free log subscription operation binding the contract event 0xa36e6e2023c9daa81ac16fbc0d2499822bc9617326793b1ec3ab50c90bcfe872.
//
// Solidity: event ForceSwap(address indexed trader, address indexed inputToken, address indexed outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Amm *AmmFilterer) WatchForceSwap(opts *bind.WatchOpts, sink chan<- *AmmForceSwap, trader []common.Address, inputToken []common.Address, outputToken []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var inputTokenRule []interface{}
	for _, inputTokenItem := range inputToken {
		inputTokenRule = append(inputTokenRule, inputTokenItem)
	}
	var outputTokenRule []interface{}
	for _, outputTokenItem := range outputToken {
		outputTokenRule = append(outputTokenRule, outputTokenItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "ForceSwap", traderRule, inputTokenRule, outputTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmForceSwap)
				if err := _Amm.contract.UnpackLog(event, "ForceSwap", log); err != nil {
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

// ParseForceSwap is a log parse operation binding the contract event 0xa36e6e2023c9daa81ac16fbc0d2499822bc9617326793b1ec3ab50c90bcfe872.
//
// Solidity: event ForceSwap(address indexed trader, address indexed inputToken, address indexed outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Amm *AmmFilterer) ParseForceSwap(log types.Log) (*AmmForceSwap, error) {
	event := new(AmmForceSwap)
	if err := _Amm.contract.UnpackLog(event, "ForceSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Amm contract.
type AmmMintIterator struct {
	Event *AmmMint // Event containing the contract specifics and raw log

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
func (it *AmmMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmMint)
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
		it.Event = new(AmmMint)
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
func (it *AmmMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmMint represents a Mint event raised by the Amm contract.
type AmmMint struct {
	Sender      common.Address
	To          common.Address
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	Liquidity   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed sender, address indexed to, uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AmmMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AmmMintIterator{contract: _Amm.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed sender, address indexed to, uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AmmMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmMint)
				if err := _Amm.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed sender, address indexed to, uint256 baseAmount, uint256 quoteAmount, uint256 liquidity)
func (_Amm *AmmFilterer) ParseMint(log types.Log) (*AmmMint, error) {
	event := new(AmmMint)
	if err := _Amm.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmRebaseIterator is returned from FilterRebase and is used to iterate over the raw logs and unpacked data for Rebase events raised by the Amm contract.
type AmmRebaseIterator struct {
	Event *AmmRebase // Event containing the contract specifics and raw log

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
func (it *AmmRebaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmRebase)
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
		it.Event = new(AmmRebase)
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
func (it *AmmRebaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmRebaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmRebase represents a Rebase event raised by the Amm contract.
type AmmRebase struct {
	QuoteReserveBefore       *big.Int
	QuoteReserveAfter        *big.Int
	BaseReserve              *big.Int
	QuoteReserveFromInternal *big.Int
	QuoteReserveFromExternal *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRebase is a free log retrieval operation binding the contract event 0xa904dcf42a3461c90173c0b966672b2cb4350e529ea12d44e562fcec43836324.
//
// Solidity: event Rebase(uint256 quoteReserveBefore, uint256 quoteReserveAfter, uint256 _baseReserve, uint256 quoteReserveFromInternal, uint256 quoteReserveFromExternal)
func (_Amm *AmmFilterer) FilterRebase(opts *bind.FilterOpts) (*AmmRebaseIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Rebase")
	if err != nil {
		return nil, err
	}
	return &AmmRebaseIterator{contract: _Amm.contract, event: "Rebase", logs: logs, sub: sub}, nil
}

// WatchRebase is a free log subscription operation binding the contract event 0xa904dcf42a3461c90173c0b966672b2cb4350e529ea12d44e562fcec43836324.
//
// Solidity: event Rebase(uint256 quoteReserveBefore, uint256 quoteReserveAfter, uint256 _baseReserve, uint256 quoteReserveFromInternal, uint256 quoteReserveFromExternal)
func (_Amm *AmmFilterer) WatchRebase(opts *bind.WatchOpts, sink chan<- *AmmRebase) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Rebase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmRebase)
				if err := _Amm.contract.UnpackLog(event, "Rebase", log); err != nil {
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

// ParseRebase is a log parse operation binding the contract event 0xa904dcf42a3461c90173c0b966672b2cb4350e529ea12d44e562fcec43836324.
//
// Solidity: event Rebase(uint256 quoteReserveBefore, uint256 quoteReserveAfter, uint256 _baseReserve, uint256 quoteReserveFromInternal, uint256 quoteReserveFromExternal)
func (_Amm *AmmFilterer) ParseRebase(log types.Log) (*AmmRebase, error) {
	event := new(AmmRebase)
	if err := _Amm.contract.UnpackLog(event, "Rebase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Amm contract.
type AmmSwapIterator struct {
	Event *AmmSwap // Event containing the contract specifics and raw log

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
func (it *AmmSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmSwap)
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
		it.Event = new(AmmSwap)
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
func (it *AmmSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmSwap represents a Swap event raised by the Amm contract.
type AmmSwap struct {
	Trader       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address indexed trader, address indexed inputToken, address indexed outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Amm *AmmFilterer) FilterSwap(opts *bind.FilterOpts, trader []common.Address, inputToken []common.Address, outputToken []common.Address) (*AmmSwapIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var inputTokenRule []interface{}
	for _, inputTokenItem := range inputToken {
		inputTokenRule = append(inputTokenRule, inputTokenItem)
	}
	var outputTokenRule []interface{}
	for _, outputTokenItem := range outputToken {
		outputTokenRule = append(outputTokenRule, outputTokenItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Swap", traderRule, inputTokenRule, outputTokenRule)
	if err != nil {
		return nil, err
	}
	return &AmmSwapIterator{contract: _Amm.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address indexed trader, address indexed inputToken, address indexed outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Amm *AmmFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *AmmSwap, trader []common.Address, inputToken []common.Address, outputToken []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var inputTokenRule []interface{}
	for _, inputTokenItem := range inputToken {
		inputTokenRule = append(inputTokenRule, inputTokenItem)
	}
	var outputTokenRule []interface{}
	for _, outputTokenItem := range outputToken {
		outputTokenRule = append(outputTokenRule, outputTokenItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Swap", traderRule, inputTokenRule, outputTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmSwap)
				if err := _Amm.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address indexed trader, address indexed inputToken, address indexed outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Amm *AmmFilterer) ParseSwap(log types.Log) (*AmmSwap, error) {
	event := new(AmmSwap)
	if err := _Amm.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Amm contract.
type AmmSyncIterator struct {
	Event *AmmSync // Event containing the contract specifics and raw log

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
func (it *AmmSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmSync)
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
		it.Event = new(AmmSync)
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
func (it *AmmSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmSync represents a Sync event raised by the Amm contract.
type AmmSync struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserveBase, uint112 reserveQuote)
func (_Amm *AmmFilterer) FilterSync(opts *bind.FilterOpts) (*AmmSyncIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &AmmSyncIterator{contract: _Amm.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserveBase, uint112 reserveQuote)
func (_Amm *AmmFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *AmmSync) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmSync)
				if err := _Amm.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserveBase, uint112 reserveQuote)
func (_Amm *AmmFilterer) ParseSync(log types.Log) (*AmmSync, error) {
	event := new(AmmSync)
	if err := _Amm.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Amm contract.
type AmmTransferIterator struct {
	Event *AmmTransfer // Event containing the contract specifics and raw log

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
func (it *AmmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmTransfer)
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
		it.Event = new(AmmTransfer)
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
func (it *AmmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmTransfer represents a Transfer event raised by the Amm contract.
type AmmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Amm *AmmFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AmmTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AmmTransferIterator{contract: _Amm.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Amm *AmmFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AmmTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmTransfer)
				if err := _Amm.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Amm *AmmFilterer) ParseTransfer(log types.Log) (*AmmTransfer, error) {
	event := new(AmmTransfer)
	if err := _Amm.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
