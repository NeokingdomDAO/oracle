// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// TelediskoTokenMetaData contains all meta data concerning the TelediskoToken contract.
var TelediskoTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OfferExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OfferMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"VestingSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFER_EXPIRATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSnapshotId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"matchOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintVesting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"offeredBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShareholderRegistry\",\"name\":\"shareholderRegistry\",\"type\":\"address\"}],\"name\":\"setShareholderRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setVesting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVoting\",\"name\":\"voting\",\"type\":\"address\"}],\"name\":\"setVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unlockedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"vestingBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TelediskoTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TelediskoTokenMetaData.ABI instead.
var TelediskoTokenABI = TelediskoTokenMetaData.ABI

// TelediskoToken is an auto generated Go binding around an Ethereum contract.
type TelediskoToken struct {
	TelediskoTokenCaller     // Read-only binding to the contract
	TelediskoTokenTransactor // Write-only binding to the contract
	TelediskoTokenFilterer   // Log filterer for contract events
}

// TelediskoTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TelediskoTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TelediskoTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TelediskoTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TelediskoTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TelediskoTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TelediskoTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TelediskoTokenSession struct {
	Contract     *TelediskoToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TelediskoTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TelediskoTokenCallerSession struct {
	Contract *TelediskoTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TelediskoTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TelediskoTokenTransactorSession struct {
	Contract     *TelediskoTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TelediskoTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TelediskoTokenRaw struct {
	Contract *TelediskoToken // Generic contract binding to access the raw methods on
}

// TelediskoTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TelediskoTokenCallerRaw struct {
	Contract *TelediskoTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TelediskoTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TelediskoTokenTransactorRaw struct {
	Contract *TelediskoTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTelediskoToken creates a new instance of TelediskoToken, bound to a specific deployed contract.
func NewTelediskoToken(address common.Address, backend bind.ContractBackend) (*TelediskoToken, error) {
	contract, err := bindTelediskoToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TelediskoToken{TelediskoTokenCaller: TelediskoTokenCaller{contract: contract}, TelediskoTokenTransactor: TelediskoTokenTransactor{contract: contract}, TelediskoTokenFilterer: TelediskoTokenFilterer{contract: contract}}, nil
}

// NewTelediskoTokenCaller creates a new read-only instance of TelediskoToken, bound to a specific deployed contract.
func NewTelediskoTokenCaller(address common.Address, caller bind.ContractCaller) (*TelediskoTokenCaller, error) {
	contract, err := bindTelediskoToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenCaller{contract: contract}, nil
}

// NewTelediskoTokenTransactor creates a new write-only instance of TelediskoToken, bound to a specific deployed contract.
func NewTelediskoTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TelediskoTokenTransactor, error) {
	contract, err := bindTelediskoToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenTransactor{contract: contract}, nil
}

// NewTelediskoTokenFilterer creates a new log filterer instance of TelediskoToken, bound to a specific deployed contract.
func NewTelediskoTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TelediskoTokenFilterer, error) {
	contract, err := bindTelediskoToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenFilterer{contract: contract}, nil
}

// bindTelediskoToken binds a generic wrapper to an already deployed contract.
func bindTelediskoToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TelediskoTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TelediskoToken *TelediskoTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TelediskoToken.Contract.TelediskoTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TelediskoToken *TelediskoTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TelediskoToken.Contract.TelediskoTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TelediskoToken *TelediskoTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TelediskoToken.Contract.TelediskoTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TelediskoToken *TelediskoTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TelediskoToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TelediskoToken *TelediskoTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TelediskoToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TelediskoToken *TelediskoTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TelediskoToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TelediskoToken *TelediskoTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TelediskoToken *TelediskoTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TelediskoToken.Contract.DEFAULTADMINROLE(&_TelediskoToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TelediskoToken *TelediskoTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TelediskoToken.Contract.DEFAULTADMINROLE(&_TelediskoToken.CallOpts)
}

// OFFEREXPIRATION is a free data retrieval call binding the contract method 0x3fe57a1d.
//
// Solidity: function OFFER_EXPIRATION() view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) OFFEREXPIRATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "OFFER_EXPIRATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OFFEREXPIRATION is a free data retrieval call binding the contract method 0x3fe57a1d.
//
// Solidity: function OFFER_EXPIRATION() view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) OFFEREXPIRATION() (*big.Int, error) {
	return _TelediskoToken.Contract.OFFEREXPIRATION(&_TelediskoToken.CallOpts)
}

// OFFEREXPIRATION is a free data retrieval call binding the contract method 0x3fe57a1d.
//
// Solidity: function OFFER_EXPIRATION() view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) OFFEREXPIRATION() (*big.Int, error) {
	return _TelediskoToken.Contract.OFFEREXPIRATION(&_TelediskoToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.Allowance(&_TelediskoToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.Allowance(&_TelediskoToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.BalanceOf(&_TelediskoToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.BalanceOf(&_TelediskoToken.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _TelediskoToken.Contract.BalanceOfAt(&_TelediskoToken.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _TelediskoToken.Contract.BalanceOfAt(&_TelediskoToken.CallOpts, account, snapshotId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TelediskoToken *TelediskoTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TelediskoToken *TelediskoTokenSession) Decimals() (uint8, error) {
	return _TelediskoToken.Contract.Decimals(&_TelediskoToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TelediskoToken *TelediskoTokenCallerSession) Decimals() (uint8, error) {
	return _TelediskoToken.Contract.Decimals(&_TelediskoToken.CallOpts)
}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) GetCurrentSnapshotId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "getCurrentSnapshotId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) GetCurrentSnapshotId() (*big.Int, error) {
	return _TelediskoToken.Contract.GetCurrentSnapshotId(&_TelediskoToken.CallOpts)
}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) GetCurrentSnapshotId() (*big.Int, error) {
	return _TelediskoToken.Contract.GetCurrentSnapshotId(&_TelediskoToken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TelediskoToken *TelediskoTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TelediskoToken *TelediskoTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TelediskoToken.Contract.GetRoleAdmin(&_TelediskoToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TelediskoToken *TelediskoTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TelediskoToken.Contract.GetRoleAdmin(&_TelediskoToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TelediskoToken *TelediskoTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TelediskoToken *TelediskoTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TelediskoToken.Contract.HasRole(&_TelediskoToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TelediskoToken *TelediskoTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TelediskoToken.Contract.HasRole(&_TelediskoToken.CallOpts, role, account)
}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) LockedBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "lockedBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) LockedBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.LockedBalanceOf(&_TelediskoToken.CallOpts, account)
}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) LockedBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.LockedBalanceOf(&_TelediskoToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TelediskoToken *TelediskoTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TelediskoToken *TelediskoTokenSession) Name() (string, error) {
	return _TelediskoToken.Contract.Name(&_TelediskoToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TelediskoToken *TelediskoTokenCallerSession) Name() (string, error) {
	return _TelediskoToken.Contract.Name(&_TelediskoToken.CallOpts)
}

// OfferedBalanceOf is a free data retrieval call binding the contract method 0xcfe63ddf.
//
// Solidity: function offeredBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) OfferedBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "offeredBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferedBalanceOf is a free data retrieval call binding the contract method 0xcfe63ddf.
//
// Solidity: function offeredBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) OfferedBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.OfferedBalanceOf(&_TelediskoToken.CallOpts, account)
}

// OfferedBalanceOf is a free data retrieval call binding the contract method 0xcfe63ddf.
//
// Solidity: function offeredBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) OfferedBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.OfferedBalanceOf(&_TelediskoToken.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TelediskoToken *TelediskoTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TelediskoToken *TelediskoTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TelediskoToken.Contract.SupportsInterface(&_TelediskoToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TelediskoToken *TelediskoTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TelediskoToken.Contract.SupportsInterface(&_TelediskoToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TelediskoToken *TelediskoTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TelediskoToken *TelediskoTokenSession) Symbol() (string, error) {
	return _TelediskoToken.Contract.Symbol(&_TelediskoToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TelediskoToken *TelediskoTokenCallerSession) Symbol() (string, error) {
	return _TelediskoToken.Contract.Symbol(&_TelediskoToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) TotalSupply() (*big.Int, error) {
	return _TelediskoToken.Contract.TotalSupply(&_TelediskoToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TelediskoToken.Contract.TotalSupply(&_TelediskoToken.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _TelediskoToken.Contract.TotalSupplyAt(&_TelediskoToken.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _TelediskoToken.Contract.TotalSupplyAt(&_TelediskoToken.CallOpts, snapshotId)
}

// UnlockedBalanceOf is a free data retrieval call binding the contract method 0x84955c88.
//
// Solidity: function unlockedBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) UnlockedBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "unlockedBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedBalanceOf is a free data retrieval call binding the contract method 0x84955c88.
//
// Solidity: function unlockedBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) UnlockedBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.UnlockedBalanceOf(&_TelediskoToken.CallOpts, account)
}

// UnlockedBalanceOf is a free data retrieval call binding the contract method 0x84955c88.
//
// Solidity: function unlockedBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) UnlockedBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.UnlockedBalanceOf(&_TelediskoToken.CallOpts, account)
}

// VestingBalanceOf is a free data retrieval call binding the contract method 0x36ca0365.
//
// Solidity: function vestingBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCaller) VestingBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TelediskoToken.contract.Call(opts, &out, "vestingBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingBalanceOf is a free data retrieval call binding the contract method 0x36ca0365.
//
// Solidity: function vestingBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) VestingBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.VestingBalanceOf(&_TelediskoToken.CallOpts, account)
}

// VestingBalanceOf is a free data retrieval call binding the contract method 0x36ca0365.
//
// Solidity: function vestingBalanceOf(address account) view returns(uint256)
func (_TelediskoToken *TelediskoTokenCallerSession) VestingBalanceOf(account common.Address) (*big.Int, error) {
	return _TelediskoToken.Contract.VestingBalanceOf(&_TelediskoToken.CallOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Approve(&_TelediskoToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Approve(&_TelediskoToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Burn(&_TelediskoToken.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Burn(&_TelediskoToken.TransactOpts, account, amount)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xa269993d.
//
// Solidity: function createOffer(uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactor) CreateOffer(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "createOffer", amount)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xa269993d.
//
// Solidity: function createOffer(uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenSession) CreateOffer(amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.CreateOffer(&_TelediskoToken.TransactOpts, amount)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xa269993d.
//
// Solidity: function createOffer(uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) CreateOffer(amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.CreateOffer(&_TelediskoToken.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TelediskoToken *TelediskoTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.DecreaseAllowance(&_TelediskoToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.DecreaseAllowance(&_TelediskoToken.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.GrantRole(&_TelediskoToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.GrantRole(&_TelediskoToken.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TelediskoToken *TelediskoTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.IncreaseAllowance(&_TelediskoToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.IncreaseAllowance(&_TelediskoToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns()
func (_TelediskoToken *TelediskoTokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "initialize", name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns()
func (_TelediskoToken *TelediskoTokenSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Initialize(&_TelediskoToken.TransactOpts, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Initialize(&_TelediskoToken.TransactOpts, name, symbol)
}

// MatchOffer is a paid mutator transaction binding the contract method 0x573b643b.
//
// Solidity: function matchOffer(address from, address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactor) MatchOffer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "matchOffer", from, to, amount)
}

// MatchOffer is a paid mutator transaction binding the contract method 0x573b643b.
//
// Solidity: function matchOffer(address from, address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenSession) MatchOffer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.MatchOffer(&_TelediskoToken.TransactOpts, from, to, amount)
}

// MatchOffer is a paid mutator transaction binding the contract method 0x573b643b.
//
// Solidity: function matchOffer(address from, address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) MatchOffer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.MatchOffer(&_TelediskoToken.TransactOpts, from, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Mint(&_TelediskoToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Mint(&_TelediskoToken.TransactOpts, to, amount)
}

// MintVesting is a paid mutator transaction binding the contract method 0x44292600.
//
// Solidity: function mintVesting(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactor) MintVesting(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "mintVesting", to, amount)
}

// MintVesting is a paid mutator transaction binding the contract method 0x44292600.
//
// Solidity: function mintVesting(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenSession) MintVesting(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.MintVesting(&_TelediskoToken.TransactOpts, to, amount)
}

// MintVesting is a paid mutator transaction binding the contract method 0x44292600.
//
// Solidity: function mintVesting(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) MintVesting(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.MintVesting(&_TelediskoToken.TransactOpts, to, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.RenounceRole(&_TelediskoToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.RenounceRole(&_TelediskoToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.RevokeRole(&_TelediskoToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.RevokeRole(&_TelediskoToken.TransactOpts, role, account)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_TelediskoToken *TelediskoTokenTransactor) SetShareholderRegistry(opts *bind.TransactOpts, shareholderRegistry common.Address) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "setShareholderRegistry", shareholderRegistry)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_TelediskoToken *TelediskoTokenSession) SetShareholderRegistry(shareholderRegistry common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.SetShareholderRegistry(&_TelediskoToken.TransactOpts, shareholderRegistry)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) SetShareholderRegistry(shareholderRegistry common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.SetShareholderRegistry(&_TelediskoToken.TransactOpts, shareholderRegistry)
}

// SetVesting is a paid mutator transaction binding the contract method 0x96c9ee62.
//
// Solidity: function setVesting(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactor) SetVesting(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "setVesting", to, amount)
}

// SetVesting is a paid mutator transaction binding the contract method 0x96c9ee62.
//
// Solidity: function setVesting(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenSession) SetVesting(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.SetVesting(&_TelediskoToken.TransactOpts, to, amount)
}

// SetVesting is a paid mutator transaction binding the contract method 0x96c9ee62.
//
// Solidity: function setVesting(address to, uint256 amount) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) SetVesting(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.SetVesting(&_TelediskoToken.TransactOpts, to, amount)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_TelediskoToken *TelediskoTokenTransactor) SetVoting(opts *bind.TransactOpts, voting common.Address) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "setVoting", voting)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_TelediskoToken *TelediskoTokenSession) SetVoting(voting common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.SetVoting(&_TelediskoToken.TransactOpts, voting)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_TelediskoToken *TelediskoTokenTransactorSession) SetVoting(voting common.Address) (*types.Transaction, error) {
	return _TelediskoToken.Contract.SetVoting(&_TelediskoToken.TransactOpts, voting)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_TelediskoToken *TelediskoTokenTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_TelediskoToken *TelediskoTokenSession) Snapshot() (*types.Transaction, error) {
	return _TelediskoToken.Contract.Snapshot(&_TelediskoToken.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_TelediskoToken *TelediskoTokenTransactorSession) Snapshot() (*types.Transaction, error) {
	return _TelediskoToken.Contract.Snapshot(&_TelediskoToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Transfer(&_TelediskoToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.Transfer(&_TelediskoToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.TransferFrom(&_TelediskoToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TelediskoToken *TelediskoTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TelediskoToken.Contract.TransferFrom(&_TelediskoToken.TransactOpts, from, to, amount)
}

// TelediskoTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TelediskoToken contract.
type TelediskoTokenApprovalIterator struct {
	Event *TelediskoTokenApproval // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenApproval)
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
		it.Event = new(TelediskoTokenApproval)
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
func (it *TelediskoTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenApproval represents a Approval event raised by the TelediskoToken contract.
type TelediskoTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TelediskoToken *TelediskoTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TelediskoTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenApprovalIterator{contract: _TelediskoToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TelediskoToken *TelediskoTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TelediskoTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenApproval)
				if err := _TelediskoToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TelediskoToken *TelediskoTokenFilterer) ParseApproval(log types.Log) (*TelediskoTokenApproval, error) {
	event := new(TelediskoTokenApproval)
	if err := _TelediskoToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TelediskoToken contract.
type TelediskoTokenInitializedIterator struct {
	Event *TelediskoTokenInitialized // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenInitialized)
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
		it.Event = new(TelediskoTokenInitialized)
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
func (it *TelediskoTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenInitialized represents a Initialized event raised by the TelediskoToken contract.
type TelediskoTokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TelediskoToken *TelediskoTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*TelediskoTokenInitializedIterator, error) {

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenInitializedIterator{contract: _TelediskoToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TelediskoToken *TelediskoTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TelediskoTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenInitialized)
				if err := _TelediskoToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TelediskoToken *TelediskoTokenFilterer) ParseInitialized(log types.Log) (*TelediskoTokenInitialized, error) {
	event := new(TelediskoTokenInitialized)
	if err := _TelediskoToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the TelediskoToken contract.
type TelediskoTokenOfferCreatedIterator struct {
	Event *TelediskoTokenOfferCreated // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenOfferCreated)
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
		it.Event = new(TelediskoTokenOfferCreated)
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
func (it *TelediskoTokenOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenOfferCreated represents a OfferCreated event raised by the TelediskoToken contract.
type TelediskoTokenOfferCreated struct {
	Id         *big.Int
	From       common.Address
	Amount     *big.Int
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0xbc9e56cae6351f41cd4be50fc6e8ff4e5c6994d27010448a584ee9f845e89e10.
//
// Solidity: event OfferCreated(uint128 id, address from, uint256 amount, uint256 expiration)
func (_TelediskoToken *TelediskoTokenFilterer) FilterOfferCreated(opts *bind.FilterOpts) (*TelediskoTokenOfferCreatedIterator, error) {

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "OfferCreated")
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenOfferCreatedIterator{contract: _TelediskoToken.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0xbc9e56cae6351f41cd4be50fc6e8ff4e5c6994d27010448a584ee9f845e89e10.
//
// Solidity: event OfferCreated(uint128 id, address from, uint256 amount, uint256 expiration)
func (_TelediskoToken *TelediskoTokenFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *TelediskoTokenOfferCreated) (event.Subscription, error) {

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "OfferCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenOfferCreated)
				if err := _TelediskoToken.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0xbc9e56cae6351f41cd4be50fc6e8ff4e5c6994d27010448a584ee9f845e89e10.
//
// Solidity: event OfferCreated(uint128 id, address from, uint256 amount, uint256 expiration)
func (_TelediskoToken *TelediskoTokenFilterer) ParseOfferCreated(log types.Log) (*TelediskoTokenOfferCreated, error) {
	event := new(TelediskoTokenOfferCreated)
	if err := _TelediskoToken.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenOfferExpiredIterator is returned from FilterOfferExpired and is used to iterate over the raw logs and unpacked data for OfferExpired events raised by the TelediskoToken contract.
type TelediskoTokenOfferExpiredIterator struct {
	Event *TelediskoTokenOfferExpired // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenOfferExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenOfferExpired)
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
		it.Event = new(TelediskoTokenOfferExpired)
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
func (it *TelediskoTokenOfferExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenOfferExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenOfferExpired represents a OfferExpired event raised by the TelediskoToken contract.
type TelediskoTokenOfferExpired struct {
	Id     *big.Int
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOfferExpired is a free log retrieval operation binding the contract event 0x1045242aa9cb5120416a010de00bec9810d7e5dc1796c7d33f69df3f444aaf37.
//
// Solidity: event OfferExpired(uint128 id, address from, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) FilterOfferExpired(opts *bind.FilterOpts) (*TelediskoTokenOfferExpiredIterator, error) {

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "OfferExpired")
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenOfferExpiredIterator{contract: _TelediskoToken.contract, event: "OfferExpired", logs: logs, sub: sub}, nil
}

// WatchOfferExpired is a free log subscription operation binding the contract event 0x1045242aa9cb5120416a010de00bec9810d7e5dc1796c7d33f69df3f444aaf37.
//
// Solidity: event OfferExpired(uint128 id, address from, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) WatchOfferExpired(opts *bind.WatchOpts, sink chan<- *TelediskoTokenOfferExpired) (event.Subscription, error) {

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "OfferExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenOfferExpired)
				if err := _TelediskoToken.contract.UnpackLog(event, "OfferExpired", log); err != nil {
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

// ParseOfferExpired is a log parse operation binding the contract event 0x1045242aa9cb5120416a010de00bec9810d7e5dc1796c7d33f69df3f444aaf37.
//
// Solidity: event OfferExpired(uint128 id, address from, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) ParseOfferExpired(log types.Log) (*TelediskoTokenOfferExpired, error) {
	event := new(TelediskoTokenOfferExpired)
	if err := _TelediskoToken.contract.UnpackLog(event, "OfferExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenOfferMatchedIterator is returned from FilterOfferMatched and is used to iterate over the raw logs and unpacked data for OfferMatched events raised by the TelediskoToken contract.
type TelediskoTokenOfferMatchedIterator struct {
	Event *TelediskoTokenOfferMatched // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenOfferMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenOfferMatched)
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
		it.Event = new(TelediskoTokenOfferMatched)
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
func (it *TelediskoTokenOfferMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenOfferMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenOfferMatched represents a OfferMatched event raised by the TelediskoToken contract.
type TelediskoTokenOfferMatched struct {
	Id     *big.Int
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOfferMatched is a free log retrieval operation binding the contract event 0xd0d70de2baa99922b89d5f189e78194a3d56c616aaf58fa941a374228aec3d08.
//
// Solidity: event OfferMatched(uint128 id, address from, address to, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) FilterOfferMatched(opts *bind.FilterOpts) (*TelediskoTokenOfferMatchedIterator, error) {

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "OfferMatched")
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenOfferMatchedIterator{contract: _TelediskoToken.contract, event: "OfferMatched", logs: logs, sub: sub}, nil
}

// WatchOfferMatched is a free log subscription operation binding the contract event 0xd0d70de2baa99922b89d5f189e78194a3d56c616aaf58fa941a374228aec3d08.
//
// Solidity: event OfferMatched(uint128 id, address from, address to, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) WatchOfferMatched(opts *bind.WatchOpts, sink chan<- *TelediskoTokenOfferMatched) (event.Subscription, error) {

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "OfferMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenOfferMatched)
				if err := _TelediskoToken.contract.UnpackLog(event, "OfferMatched", log); err != nil {
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

// ParseOfferMatched is a log parse operation binding the contract event 0xd0d70de2baa99922b89d5f189e78194a3d56c616aaf58fa941a374228aec3d08.
//
// Solidity: event OfferMatched(uint128 id, address from, address to, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) ParseOfferMatched(log types.Log) (*TelediskoTokenOfferMatched, error) {
	event := new(TelediskoTokenOfferMatched)
	if err := _TelediskoToken.contract.UnpackLog(event, "OfferMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TelediskoToken contract.
type TelediskoTokenRoleAdminChangedIterator struct {
	Event *TelediskoTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenRoleAdminChanged)
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
		it.Event = new(TelediskoTokenRoleAdminChanged)
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
func (it *TelediskoTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenRoleAdminChanged represents a RoleAdminChanged event raised by the TelediskoToken contract.
type TelediskoTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TelediskoToken *TelediskoTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TelediskoTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenRoleAdminChangedIterator{contract: _TelediskoToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TelediskoToken *TelediskoTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TelediskoTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenRoleAdminChanged)
				if err := _TelediskoToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TelediskoToken *TelediskoTokenFilterer) ParseRoleAdminChanged(log types.Log) (*TelediskoTokenRoleAdminChanged, error) {
	event := new(TelediskoTokenRoleAdminChanged)
	if err := _TelediskoToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TelediskoToken contract.
type TelediskoTokenRoleGrantedIterator struct {
	Event *TelediskoTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenRoleGranted)
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
		it.Event = new(TelediskoTokenRoleGranted)
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
func (it *TelediskoTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenRoleGranted represents a RoleGranted event raised by the TelediskoToken contract.
type TelediskoTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TelediskoToken *TelediskoTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TelediskoTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenRoleGrantedIterator{contract: _TelediskoToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TelediskoToken *TelediskoTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TelediskoTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenRoleGranted)
				if err := _TelediskoToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TelediskoToken *TelediskoTokenFilterer) ParseRoleGranted(log types.Log) (*TelediskoTokenRoleGranted, error) {
	event := new(TelediskoTokenRoleGranted)
	if err := _TelediskoToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TelediskoToken contract.
type TelediskoTokenRoleRevokedIterator struct {
	Event *TelediskoTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenRoleRevoked)
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
		it.Event = new(TelediskoTokenRoleRevoked)
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
func (it *TelediskoTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenRoleRevoked represents a RoleRevoked event raised by the TelediskoToken contract.
type TelediskoTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TelediskoToken *TelediskoTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TelediskoTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenRoleRevokedIterator{contract: _TelediskoToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TelediskoToken *TelediskoTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TelediskoTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenRoleRevoked)
				if err := _TelediskoToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TelediskoToken *TelediskoTokenFilterer) ParseRoleRevoked(log types.Log) (*TelediskoTokenRoleRevoked, error) {
	event := new(TelediskoTokenRoleRevoked)
	if err := _TelediskoToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the TelediskoToken contract.
type TelediskoTokenSnapshotIterator struct {
	Event *TelediskoTokenSnapshot // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenSnapshot)
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
		it.Event = new(TelediskoTokenSnapshot)
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
func (it *TelediskoTokenSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenSnapshot represents a Snapshot event raised by the TelediskoToken contract.
type TelediskoTokenSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_TelediskoToken *TelediskoTokenFilterer) FilterSnapshot(opts *bind.FilterOpts) (*TelediskoTokenSnapshotIterator, error) {

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenSnapshotIterator{contract: _TelediskoToken.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_TelediskoToken *TelediskoTokenFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *TelediskoTokenSnapshot) (event.Subscription, error) {

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenSnapshot)
				if err := _TelediskoToken.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_TelediskoToken *TelediskoTokenFilterer) ParseSnapshot(log types.Log) (*TelediskoTokenSnapshot, error) {
	event := new(TelediskoTokenSnapshot)
	if err := _TelediskoToken.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TelediskoToken contract.
type TelediskoTokenTransferIterator struct {
	Event *TelediskoTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenTransfer)
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
		it.Event = new(TelediskoTokenTransfer)
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
func (it *TelediskoTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenTransfer represents a Transfer event raised by the TelediskoToken contract.
type TelediskoTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TelediskoToken *TelediskoTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TelediskoTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenTransferIterator{contract: _TelediskoToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TelediskoToken *TelediskoTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TelediskoTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenTransfer)
				if err := _TelediskoToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TelediskoToken *TelediskoTokenFilterer) ParseTransfer(log types.Log) (*TelediskoTokenTransfer, error) {
	event := new(TelediskoTokenTransfer)
	if err := _TelediskoToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TelediskoTokenVestingSetIterator is returned from FilterVestingSet and is used to iterate over the raw logs and unpacked data for VestingSet events raised by the TelediskoToken contract.
type TelediskoTokenVestingSetIterator struct {
	Event *TelediskoTokenVestingSet // Event containing the contract specifics and raw log

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
func (it *TelediskoTokenVestingSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelediskoTokenVestingSet)
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
		it.Event = new(TelediskoTokenVestingSet)
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
func (it *TelediskoTokenVestingSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelediskoTokenVestingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelediskoTokenVestingSet represents a VestingSet event raised by the TelediskoToken contract.
type TelediskoTokenVestingSet struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVestingSet is a free log retrieval operation binding the contract event 0x90782417eb6da69d0b671119cc8e9c2063c7b0987ef1378141a54770a8c91fb3.
//
// Solidity: event VestingSet(address to, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) FilterVestingSet(opts *bind.FilterOpts) (*TelediskoTokenVestingSetIterator, error) {

	logs, sub, err := _TelediskoToken.contract.FilterLogs(opts, "VestingSet")
	if err != nil {
		return nil, err
	}
	return &TelediskoTokenVestingSetIterator{contract: _TelediskoToken.contract, event: "VestingSet", logs: logs, sub: sub}, nil
}

// WatchVestingSet is a free log subscription operation binding the contract event 0x90782417eb6da69d0b671119cc8e9c2063c7b0987ef1378141a54770a8c91fb3.
//
// Solidity: event VestingSet(address to, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) WatchVestingSet(opts *bind.WatchOpts, sink chan<- *TelediskoTokenVestingSet) (event.Subscription, error) {

	logs, sub, err := _TelediskoToken.contract.WatchLogs(opts, "VestingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelediskoTokenVestingSet)
				if err := _TelediskoToken.contract.UnpackLog(event, "VestingSet", log); err != nil {
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

// ParseVestingSet is a log parse operation binding the contract event 0x90782417eb6da69d0b671119cc8e9c2063c7b0987ef1378141a54770a8c91fb3.
//
// Solidity: event VestingSet(address to, uint256 amount)
func (_TelediskoToken *TelediskoTokenFilterer) ParseVestingSet(log types.Log) (*TelediskoTokenVestingSet, error) {
	event := new(TelediskoTokenVestingSet)
	if err := _TelediskoToken.contract.UnpackLog(event, "VestingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
