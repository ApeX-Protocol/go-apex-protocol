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

// Multicall2Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall2Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Result struct {
	Success    bool
	ReturnData []byte
}

// Multicall2MetaData contains all meta data concerning the Multicall2 contract.
var Multicall2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Multicall2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Multicall2MetaData.ABI instead.
var Multicall2ABI = Multicall2MetaData.ABI

// Multicall2 is an auto generated Go binding around an Ethereum contract.
type Multicall2 struct {
	Multicall2Caller     // Read-only binding to the contract
	Multicall2Transactor // Write-only binding to the contract
	Multicall2Filterer   // Log filterer for contract events
}

// Multicall2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Multicall2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Multicall2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Multicall2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Multicall2Session struct {
	Contract     *Multicall2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Multicall2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Multicall2CallerSession struct {
	Contract *Multicall2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Multicall2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Multicall2TransactorSession struct {
	Contract     *Multicall2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Multicall2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Multicall2Raw struct {
	Contract *Multicall2 // Generic contract binding to access the raw methods on
}

// Multicall2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Multicall2CallerRaw struct {
	Contract *Multicall2Caller // Generic read-only contract binding to access the raw methods on
}

// Multicall2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Multicall2TransactorRaw struct {
	Contract *Multicall2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall2 creates a new instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2(address common.Address, backend bind.ContractBackend) (*Multicall2, error) {
	contract, err := bindMulticall2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall2{Multicall2Caller: Multicall2Caller{contract: contract}, Multicall2Transactor: Multicall2Transactor{contract: contract}, Multicall2Filterer: Multicall2Filterer{contract: contract}}, nil
}

// NewMulticall2Caller creates a new read-only instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2Caller(address common.Address, caller bind.ContractCaller) (*Multicall2Caller, error) {
	contract, err := bindMulticall2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Multicall2Caller{contract: contract}, nil
}

// NewMulticall2Transactor creates a new write-only instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2Transactor(address common.Address, transactor bind.ContractTransactor) (*Multicall2Transactor, error) {
	contract, err := bindMulticall2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Multicall2Transactor{contract: contract}, nil
}

// NewMulticall2Filterer creates a new log filterer instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2Filterer(address common.Address, filterer bind.ContractFilterer) (*Multicall2Filterer, error) {
	contract, err := bindMulticall2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Multicall2Filterer{contract: contract}, nil
}

// bindMulticall2 binds a generic wrapper to an already deployed contract.
func bindMulticall2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Multicall2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall2 *Multicall2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall2.Contract.Multicall2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall2 *Multicall2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall2.Contract.Multicall2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall2 *Multicall2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall2.Contract.Multicall2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall2 *Multicall2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall2 *Multicall2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall2 *Multicall2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall2.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Caller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Session) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicall2.Contract.GetBlockHash(&_Multicall2.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2CallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicall2.Contract.GetBlockHash(&_Multicall2.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall2 *Multicall2Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall2 *Multicall2Session) GetBlockNumber() (*big.Int, error) {
	return _Multicall2.Contract.GetBlockNumber(&_Multicall2.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall2 *Multicall2CallerSession) GetBlockNumber() (*big.Int, error) {
	return _Multicall2.Contract.GetBlockNumber(&_Multicall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall2 *Multicall2Session) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicall2.Contract.GetCurrentBlockCoinbase(&_Multicall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicall2.Contract.GetCurrentBlockCoinbase(&_Multicall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall2 *Multicall2Session) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockDifficulty(&_Multicall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockDifficulty(&_Multicall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall2 *Multicall2Session) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockGasLimit(&_Multicall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockGasLimit(&_Multicall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall2 *Multicall2Session) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockTimestamp(&_Multicall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockTimestamp(&_Multicall2.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall2 *Multicall2Caller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall2 *Multicall2Session) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicall2.Contract.GetEthBalance(&_Multicall2.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall2 *Multicall2CallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicall2.Contract.GetEthBalance(&_Multicall2.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Caller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Session) GetLastBlockHash() ([32]byte, error) {
	return _Multicall2.Contract.GetLastBlockHash(&_Multicall2.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2CallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicall2.Contract.GetLastBlockHash(&_Multicall2.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall2 *Multicall2Transactor) Aggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall2 *Multicall2Session) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.Aggregate(&_Multicall2.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall2 *Multicall2TransactorSession) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.Aggregate(&_Multicall2.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Transactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Session) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.BlockAndAggregate(&_Multicall2.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2TransactorSession) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.BlockAndAggregate(&_Multicall2.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Transactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Session) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_Multicall2 *Multicall2TransactorSession) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Transactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Session) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryBlockAndAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2TransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryBlockAndAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}
