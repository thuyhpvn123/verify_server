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
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"internalType\":\"enumAuthOTP.TypeMethod\",\"name\":\"_typeMethod\",\"type\":\"uint8\"}],\"name\":\"addBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"AuthenticationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"AuthenticationHashStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"otp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chatbotPhone\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumAuthOTP.TypeMethod\",\"name\":\"typeMethod\",\"type\":\"uint8\"}],\"name\":\"BotAuthenticationRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_identifier\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_encryptedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_encryptedSecretKey\",\"type\":\"bytes\"}],\"name\":\"completeAuthentication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"otp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userEmail\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"targetMailServerEmail\",\"type\":\"string\"}],\"name\":\"EmailAuthenticationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"primaryEmail\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subEmail\",\"type\":\"string\"}],\"name\":\"EmailSubCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromEmail\",\"type\":\"string\"}],\"name\":\"EmailVerified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_identifier\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"enumAuthOTP.TypeMethod\",\"name\":\"_typeMethod\",\"type\":\"uint8\"}],\"name\":\"requestAuthentication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_agreed\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"}],\"name\":\"setDomainEmail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"StepVerified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_botId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"internalType\":\"enumAuthOTP.TypeMethod\",\"name\":\"_typeMethod\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_otp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_identifier\",\"type\":\"string\"}],\"name\":\"validateOTP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authenticatedWallets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authenticationHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"detailBots\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"phoneNumber\",\"type\":\"string\"},{\"internalType\":\"enumAuthOTP.TypeMethod\",\"name\":\"typeMethod\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"busy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeOccupied\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"detailBotsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"emailIdToWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getSubEmail\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"identifierToWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OTPs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"OTP\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"botId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeRequest\",\"type\":\"uint256\"},{\"internalType\":\"enumAuthOTP.TypeMethod\",\"name\":\"method\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"phoneIdToWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"primaryEmails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"publicKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"subEmails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verificationStates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"emailVerified\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"phoneVerified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_encryptedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_encryptedSecretKey\",\"type\":\"bytes\"}],\"name\":\"verifyAuthenticationHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// OTPs is a free data retrieval call binding the contract method 0x6d956877.
//
// Solidity: function OTPs(string ) view returns(uint256 OTP, string publicKey, bool verified, uint256 botId, uint256 timeRequest, uint8 method)
func (_Contract *ContractCaller) OTPs(opts *bind.CallOpts, arg0 string) (struct {
	OTP         *big.Int
	PublicKey   string
	Verified    bool
	BotId       *big.Int
	TimeRequest *big.Int
	Method      uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "OTPs", arg0)

	outstruct := new(struct {
		OTP         *big.Int
		PublicKey   string
		Verified    bool
		BotId       *big.Int
		TimeRequest *big.Int
		Method      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OTP = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PublicKey = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Verified = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.BotId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TimeRequest = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Method = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// OTPs is a free data retrieval call binding the contract method 0x6d956877.
//
// Solidity: function OTPs(string ) view returns(uint256 OTP, string publicKey, bool verified, uint256 botId, uint256 timeRequest, uint8 method)
func (_Contract *ContractSession) OTPs(arg0 string) (struct {
	OTP         *big.Int
	PublicKey   string
	Verified    bool
	BotId       *big.Int
	TimeRequest *big.Int
	Method      uint8
}, error) {
	return _Contract.Contract.OTPs(&_Contract.CallOpts, arg0)
}

// OTPs is a free data retrieval call binding the contract method 0x6d956877.
//
// Solidity: function OTPs(string ) view returns(uint256 OTP, string publicKey, bool verified, uint256 botId, uint256 timeRequest, uint8 method)
func (_Contract *ContractCallerSession) OTPs(arg0 string) (struct {
	OTP         *big.Int
	PublicKey   string
	Verified    bool
	BotId       *big.Int
	TimeRequest *big.Int
	Method      uint8
}, error) {
	return _Contract.Contract.OTPs(&_Contract.CallOpts, arg0)
}

// AuthenticatedWallets is a free data retrieval call binding the contract method 0x81d0e9a3.
//
// Solidity: function authenticatedWallets(address ) view returns(bool)
func (_Contract *ContractCaller) AuthenticatedWallets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "authenticatedWallets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthenticatedWallets is a free data retrieval call binding the contract method 0x81d0e9a3.
//
// Solidity: function authenticatedWallets(address ) view returns(bool)
func (_Contract *ContractSession) AuthenticatedWallets(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AuthenticatedWallets(&_Contract.CallOpts, arg0)
}

// AuthenticatedWallets is a free data retrieval call binding the contract method 0x81d0e9a3.
//
// Solidity: function authenticatedWallets(address ) view returns(bool)
func (_Contract *ContractCallerSession) AuthenticatedWallets(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AuthenticatedWallets(&_Contract.CallOpts, arg0)
}

// AuthenticationHashes is a free data retrieval call binding the contract method 0xb5de2544.
//
// Solidity: function authenticationHashes(address ) view returns(bytes32)
func (_Contract *ContractCaller) AuthenticationHashes(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "authenticationHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AuthenticationHashes is a free data retrieval call binding the contract method 0xb5de2544.
//
// Solidity: function authenticationHashes(address ) view returns(bytes32)
func (_Contract *ContractSession) AuthenticationHashes(arg0 common.Address) ([32]byte, error) {
	return _Contract.Contract.AuthenticationHashes(&_Contract.CallOpts, arg0)
}

// AuthenticationHashes is a free data retrieval call binding the contract method 0xb5de2544.
//
// Solidity: function authenticationHashes(address ) view returns(bytes32)
func (_Contract *ContractCallerSession) AuthenticationHashes(arg0 common.Address) ([32]byte, error) {
	return _Contract.Contract.AuthenticationHashes(&_Contract.CallOpts, arg0)
}

// DetailBots is a free data retrieval call binding the contract method 0xc50e17f7.
//
// Solidity: function detailBots(uint256 ) view returns(string phoneNumber, uint8 typeMethod, bool busy, uint256 timeOccupied, bool status)
func (_Contract *ContractCaller) DetailBots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PhoneNumber  string
	TypeMethod   uint8
	Busy         bool
	TimeOccupied *big.Int
	Status       bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "detailBots", arg0)

	outstruct := new(struct {
		PhoneNumber  string
		TypeMethod   uint8
		Busy         bool
		TimeOccupied *big.Int
		Status       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PhoneNumber = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TypeMethod = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Busy = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.TimeOccupied = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// DetailBots is a free data retrieval call binding the contract method 0xc50e17f7.
//
// Solidity: function detailBots(uint256 ) view returns(string phoneNumber, uint8 typeMethod, bool busy, uint256 timeOccupied, bool status)
func (_Contract *ContractSession) DetailBots(arg0 *big.Int) (struct {
	PhoneNumber  string
	TypeMethod   uint8
	Busy         bool
	TimeOccupied *big.Int
	Status       bool
}, error) {
	return _Contract.Contract.DetailBots(&_Contract.CallOpts, arg0)
}

// DetailBots is a free data retrieval call binding the contract method 0xc50e17f7.
//
// Solidity: function detailBots(uint256 ) view returns(string phoneNumber, uint8 typeMethod, bool busy, uint256 timeOccupied, bool status)
func (_Contract *ContractCallerSession) DetailBots(arg0 *big.Int) (struct {
	PhoneNumber  string
	TypeMethod   uint8
	Busy         bool
	TimeOccupied *big.Int
	Status       bool
}, error) {
	return _Contract.Contract.DetailBots(&_Contract.CallOpts, arg0)
}

// DetailBotsCount is a free data retrieval call binding the contract method 0x9d5e5fa2.
//
// Solidity: function detailBotsCount() view returns(uint256)
func (_Contract *ContractCaller) DetailBotsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "detailBotsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetailBotsCount is a free data retrieval call binding the contract method 0x9d5e5fa2.
//
// Solidity: function detailBotsCount() view returns(uint256)
func (_Contract *ContractSession) DetailBotsCount() (*big.Int, error) {
	return _Contract.Contract.DetailBotsCount(&_Contract.CallOpts)
}

// DetailBotsCount is a free data retrieval call binding the contract method 0x9d5e5fa2.
//
// Solidity: function detailBotsCount() view returns(uint256)
func (_Contract *ContractCallerSession) DetailBotsCount() (*big.Int, error) {
	return _Contract.Contract.DetailBotsCount(&_Contract.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(string)
func (_Contract *ContractCaller) Domain(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "domain")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(string)
func (_Contract *ContractSession) Domain() (string, error) {
	return _Contract.Contract.Domain(&_Contract.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(string)
func (_Contract *ContractCallerSession) Domain() (string, error) {
	return _Contract.Contract.Domain(&_Contract.CallOpts)
}

// EmailIdToWallet is a free data retrieval call binding the contract method 0x2dd48791.
//
// Solidity: function emailIdToWallet(string ) view returns(address)
func (_Contract *ContractCaller) EmailIdToWallet(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "emailIdToWallet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmailIdToWallet is a free data retrieval call binding the contract method 0x2dd48791.
//
// Solidity: function emailIdToWallet(string ) view returns(address)
func (_Contract *ContractSession) EmailIdToWallet(arg0 string) (common.Address, error) {
	return _Contract.Contract.EmailIdToWallet(&_Contract.CallOpts, arg0)
}

// EmailIdToWallet is a free data retrieval call binding the contract method 0x2dd48791.
//
// Solidity: function emailIdToWallet(string ) view returns(address)
func (_Contract *ContractCallerSession) EmailIdToWallet(arg0 string) (common.Address, error) {
	return _Contract.Contract.EmailIdToWallet(&_Contract.CallOpts, arg0)
}

// GetSubEmail is a free data retrieval call binding the contract method 0x65fbe46d.
//
// Solidity: function getSubEmail(address wallet) view returns(string)
func (_Contract *ContractCaller) GetSubEmail(opts *bind.CallOpts, wallet common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSubEmail", wallet)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSubEmail is a free data retrieval call binding the contract method 0x65fbe46d.
//
// Solidity: function getSubEmail(address wallet) view returns(string)
func (_Contract *ContractSession) GetSubEmail(wallet common.Address) (string, error) {
	return _Contract.Contract.GetSubEmail(&_Contract.CallOpts, wallet)
}

// GetSubEmail is a free data retrieval call binding the contract method 0x65fbe46d.
//
// Solidity: function getSubEmail(address wallet) view returns(string)
func (_Contract *ContractCallerSession) GetSubEmail(wallet common.Address) (string, error) {
	return _Contract.Contract.GetSubEmail(&_Contract.CallOpts, wallet)
}

// IdentifierToWallet is a free data retrieval call binding the contract method 0x0843cc7d.
//
// Solidity: function identifierToWallet(string ) view returns(address)
func (_Contract *ContractCaller) IdentifierToWallet(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "identifierToWallet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentifierToWallet is a free data retrieval call binding the contract method 0x0843cc7d.
//
// Solidity: function identifierToWallet(string ) view returns(address)
func (_Contract *ContractSession) IdentifierToWallet(arg0 string) (common.Address, error) {
	return _Contract.Contract.IdentifierToWallet(&_Contract.CallOpts, arg0)
}

// IdentifierToWallet is a free data retrieval call binding the contract method 0x0843cc7d.
//
// Solidity: function identifierToWallet(string ) view returns(address)
func (_Contract *ContractCallerSession) IdentifierToWallet(arg0 string) (common.Address, error) {
	return _Contract.Contract.IdentifierToWallet(&_Contract.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Contract *ContractCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Contract *ContractSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Contract.Contract.IsAdmin(&_Contract.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Contract *ContractCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Contract.Contract.IsAdmin(&_Contract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PhoneIdToWallet is a free data retrieval call binding the contract method 0x33c6129f.
//
// Solidity: function phoneIdToWallet(string ) view returns(address)
func (_Contract *ContractCaller) PhoneIdToWallet(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "phoneIdToWallet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PhoneIdToWallet is a free data retrieval call binding the contract method 0x33c6129f.
//
// Solidity: function phoneIdToWallet(string ) view returns(address)
func (_Contract *ContractSession) PhoneIdToWallet(arg0 string) (common.Address, error) {
	return _Contract.Contract.PhoneIdToWallet(&_Contract.CallOpts, arg0)
}

// PhoneIdToWallet is a free data retrieval call binding the contract method 0x33c6129f.
//
// Solidity: function phoneIdToWallet(string ) view returns(address)
func (_Contract *ContractCallerSession) PhoneIdToWallet(arg0 string) (common.Address, error) {
	return _Contract.Contract.PhoneIdToWallet(&_Contract.CallOpts, arg0)
}

// PrimaryEmails is a free data retrieval call binding the contract method 0x18419204.
//
// Solidity: function primaryEmails(address ) view returns(string)
func (_Contract *ContractCaller) PrimaryEmails(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "primaryEmails", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PrimaryEmails is a free data retrieval call binding the contract method 0x18419204.
//
// Solidity: function primaryEmails(address ) view returns(string)
func (_Contract *ContractSession) PrimaryEmails(arg0 common.Address) (string, error) {
	return _Contract.Contract.PrimaryEmails(&_Contract.CallOpts, arg0)
}

// PrimaryEmails is a free data retrieval call binding the contract method 0x18419204.
//
// Solidity: function primaryEmails(address ) view returns(string)
func (_Contract *ContractCallerSession) PrimaryEmails(arg0 common.Address) (string, error) {
	return _Contract.Contract.PrimaryEmails(&_Contract.CallOpts, arg0)
}

// PublicKeyHashes is a free data retrieval call binding the contract method 0x8c82cc1d.
//
// Solidity: function publicKeyHashes(string ) view returns(bytes32 dataHash, uint256 timestamp)
func (_Contract *ContractCaller) PublicKeyHashes(opts *bind.CallOpts, arg0 string) (struct {
	DataHash  [32]byte
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "publicKeyHashes", arg0)

	outstruct := new(struct {
		DataHash  [32]byte
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PublicKeyHashes is a free data retrieval call binding the contract method 0x8c82cc1d.
//
// Solidity: function publicKeyHashes(string ) view returns(bytes32 dataHash, uint256 timestamp)
func (_Contract *ContractSession) PublicKeyHashes(arg0 string) (struct {
	DataHash  [32]byte
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.PublicKeyHashes(&_Contract.CallOpts, arg0)
}

// PublicKeyHashes is a free data retrieval call binding the contract method 0x8c82cc1d.
//
// Solidity: function publicKeyHashes(string ) view returns(bytes32 dataHash, uint256 timestamp)
func (_Contract *ContractCallerSession) PublicKeyHashes(arg0 string) (struct {
	DataHash  [32]byte
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.PublicKeyHashes(&_Contract.CallOpts, arg0)
}

// SubEmails is a free data retrieval call binding the contract method 0x80cc9230.
//
// Solidity: function subEmails(address ) view returns(string)
func (_Contract *ContractCaller) SubEmails(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "subEmails", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubEmails is a free data retrieval call binding the contract method 0x80cc9230.
//
// Solidity: function subEmails(address ) view returns(string)
func (_Contract *ContractSession) SubEmails(arg0 common.Address) (string, error) {
	return _Contract.Contract.SubEmails(&_Contract.CallOpts, arg0)
}

// SubEmails is a free data retrieval call binding the contract method 0x80cc9230.
//
// Solidity: function subEmails(address ) view returns(string)
func (_Contract *ContractCallerSession) SubEmails(arg0 common.Address) (string, error) {
	return _Contract.Contract.SubEmails(&_Contract.CallOpts, arg0)
}

// VerificationStates is a free data retrieval call binding the contract method 0xdfd87439.
//
// Solidity: function verificationStates(address ) view returns(bool emailVerified, bool phoneVerified)
func (_Contract *ContractCaller) VerificationStates(opts *bind.CallOpts, arg0 common.Address) (struct {
	EmailVerified bool
	PhoneVerified bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verificationStates", arg0)

	outstruct := new(struct {
		EmailVerified bool
		PhoneVerified bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmailVerified = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PhoneVerified = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// VerificationStates is a free data retrieval call binding the contract method 0xdfd87439.
//
// Solidity: function verificationStates(address ) view returns(bool emailVerified, bool phoneVerified)
func (_Contract *ContractSession) VerificationStates(arg0 common.Address) (struct {
	EmailVerified bool
	PhoneVerified bool
}, error) {
	return _Contract.Contract.VerificationStates(&_Contract.CallOpts, arg0)
}

// VerificationStates is a free data retrieval call binding the contract method 0xdfd87439.
//
// Solidity: function verificationStates(address ) view returns(bool emailVerified, bool phoneVerified)
func (_Contract *ContractCallerSession) VerificationStates(arg0 common.Address) (struct {
	EmailVerified bool
	PhoneVerified bool
}, error) {
	return _Contract.Contract.VerificationStates(&_Contract.CallOpts, arg0)
}

// VerifyAuthenticationHash is a free data retrieval call binding the contract method 0x1d1584e5.
//
// Solidity: function verifyAuthenticationHash(address _userWallet, bytes _encryptedMessage, bytes _encryptedSecretKey) view returns(bool)
func (_Contract *ContractCaller) VerifyAuthenticationHash(opts *bind.CallOpts, _userWallet common.Address, _encryptedMessage []byte, _encryptedSecretKey []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifyAuthenticationHash", _userWallet, _encryptedMessage, _encryptedSecretKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAuthenticationHash is a free data retrieval call binding the contract method 0x1d1584e5.
//
// Solidity: function verifyAuthenticationHash(address _userWallet, bytes _encryptedMessage, bytes _encryptedSecretKey) view returns(bool)
func (_Contract *ContractSession) VerifyAuthenticationHash(_userWallet common.Address, _encryptedMessage []byte, _encryptedSecretKey []byte) (bool, error) {
	return _Contract.Contract.VerifyAuthenticationHash(&_Contract.CallOpts, _userWallet, _encryptedMessage, _encryptedSecretKey)
}

// VerifyAuthenticationHash is a free data retrieval call binding the contract method 0x1d1584e5.
//
// Solidity: function verifyAuthenticationHash(address _userWallet, bytes _encryptedMessage, bytes _encryptedSecretKey) view returns(bool)
func (_Contract *ContractCallerSession) VerifyAuthenticationHash(_userWallet common.Address, _encryptedMessage []byte, _encryptedSecretKey []byte) (bool, error) {
	return _Contract.Contract.VerifyAuthenticationHash(&_Contract.CallOpts, _userWallet, _encryptedMessage, _encryptedSecretKey)
}

// WalletCooldown is a free data retrieval call binding the contract method 0x18e7ea2a.
//
// Solidity: function walletCooldown(address ) view returns(uint256)
func (_Contract *ContractCaller) WalletCooldown(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "walletCooldown", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WalletCooldown is a free data retrieval call binding the contract method 0x18e7ea2a.
//
// Solidity: function walletCooldown(address ) view returns(uint256)
func (_Contract *ContractSession) WalletCooldown(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.WalletCooldown(&_Contract.CallOpts, arg0)
}

// WalletCooldown is a free data retrieval call binding the contract method 0x18e7ea2a.
//
// Solidity: function walletCooldown(address ) view returns(uint256)
func (_Contract *ContractCallerSession) WalletCooldown(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.WalletCooldown(&_Contract.CallOpts, arg0)
}

// AddBot is a paid mutator transaction binding the contract method 0x3cf2d7b0.
//
// Solidity: function addBot(string _phoneNumber, uint8 _typeMethod) returns()
func (_Contract *ContractTransactor) AddBot(opts *bind.TransactOpts, _phoneNumber string, _typeMethod uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addBot", _phoneNumber, _typeMethod)
}

// AddBot is a paid mutator transaction binding the contract method 0x3cf2d7b0.
//
// Solidity: function addBot(string _phoneNumber, uint8 _typeMethod) returns()
func (_Contract *ContractSession) AddBot(_phoneNumber string, _typeMethod uint8) (*types.Transaction, error) {
	return _Contract.Contract.AddBot(&_Contract.TransactOpts, _phoneNumber, _typeMethod)
}

// AddBot is a paid mutator transaction binding the contract method 0x3cf2d7b0.
//
// Solidity: function addBot(string _phoneNumber, uint8 _typeMethod) returns()
func (_Contract *ContractTransactorSession) AddBot(_phoneNumber string, _typeMethod uint8) (*types.Transaction, error) {
	return _Contract.Contract.AddBot(&_Contract.TransactOpts, _phoneNumber, _typeMethod)
}

// CompleteAuthentication is a paid mutator transaction binding the contract method 0xf0f78ec9.
//
// Solidity: function completeAuthentication(string _identifier, bytes _encryptedMessage, bytes _encryptedSecretKey) returns()
func (_Contract *ContractTransactor) CompleteAuthentication(opts *bind.TransactOpts, _identifier string, _encryptedMessage []byte, _encryptedSecretKey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "completeAuthentication", _identifier, _encryptedMessage, _encryptedSecretKey)
}

// CompleteAuthentication is a paid mutator transaction binding the contract method 0xf0f78ec9.
//
// Solidity: function completeAuthentication(string _identifier, bytes _encryptedMessage, bytes _encryptedSecretKey) returns()
func (_Contract *ContractSession) CompleteAuthentication(_identifier string, _encryptedMessage []byte, _encryptedSecretKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CompleteAuthentication(&_Contract.TransactOpts, _identifier, _encryptedMessage, _encryptedSecretKey)
}

// CompleteAuthentication is a paid mutator transaction binding the contract method 0xf0f78ec9.
//
// Solidity: function completeAuthentication(string _identifier, bytes _encryptedMessage, bytes _encryptedSecretKey) returns()
func (_Contract *ContractTransactorSession) CompleteAuthentication(_identifier string, _encryptedMessage []byte, _encryptedSecretKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CompleteAuthentication(&_Contract.TransactOpts, _identifier, _encryptedMessage, _encryptedSecretKey)
}

// RequestAuthentication is a paid mutator transaction binding the contract method 0x9b3e74c8.
//
// Solidity: function requestAuthentication(string _identifier, address _walletAddress, string _publicKey, uint8 _typeMethod) returns()
func (_Contract *ContractTransactor) RequestAuthentication(opts *bind.TransactOpts, _identifier string, _walletAddress common.Address, _publicKey string, _typeMethod uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestAuthentication", _identifier, _walletAddress, _publicKey, _typeMethod)
}

// RequestAuthentication is a paid mutator transaction binding the contract method 0x9b3e74c8.
//
// Solidity: function requestAuthentication(string _identifier, address _walletAddress, string _publicKey, uint8 _typeMethod) returns()
func (_Contract *ContractSession) RequestAuthentication(_identifier string, _walletAddress common.Address, _publicKey string, _typeMethod uint8) (*types.Transaction, error) {
	return _Contract.Contract.RequestAuthentication(&_Contract.TransactOpts, _identifier, _walletAddress, _publicKey, _typeMethod)
}

// RequestAuthentication is a paid mutator transaction binding the contract method 0x9b3e74c8.
//
// Solidity: function requestAuthentication(string _identifier, address _walletAddress, string _publicKey, uint8 _typeMethod) returns()
func (_Contract *ContractTransactorSession) RequestAuthentication(_identifier string, _walletAddress common.Address, _publicKey string, _typeMethod uint8) (*types.Transaction, error) {
	return _Contract.Contract.RequestAuthentication(&_Contract.TransactOpts, _identifier, _walletAddress, _publicKey, _typeMethod)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _admin, bool _agreed) returns()
func (_Contract *ContractTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address, _agreed bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAdmin", _admin, _agreed)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _admin, bool _agreed) returns()
func (_Contract *ContractSession) SetAdmin(_admin common.Address, _agreed bool) (*types.Transaction, error) {
	return _Contract.Contract.SetAdmin(&_Contract.TransactOpts, _admin, _agreed)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _admin, bool _agreed) returns()
func (_Contract *ContractTransactorSession) SetAdmin(_admin common.Address, _agreed bool) (*types.Transaction, error) {
	return _Contract.Contract.SetAdmin(&_Contract.TransactOpts, _admin, _agreed)
}

// SetDomainEmail is a paid mutator transaction binding the contract method 0x72354866.
//
// Solidity: function setDomainEmail(string _domain) returns()
func (_Contract *ContractTransactor) SetDomainEmail(opts *bind.TransactOpts, _domain string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDomainEmail", _domain)
}

// SetDomainEmail is a paid mutator transaction binding the contract method 0x72354866.
//
// Solidity: function setDomainEmail(string _domain) returns()
func (_Contract *ContractSession) SetDomainEmail(_domain string) (*types.Transaction, error) {
	return _Contract.Contract.SetDomainEmail(&_Contract.TransactOpts, _domain)
}

// SetDomainEmail is a paid mutator transaction binding the contract method 0x72354866.
//
// Solidity: function setDomainEmail(string _domain) returns()
func (_Contract *ContractTransactorSession) SetDomainEmail(_domain string) (*types.Transaction, error) {
	return _Contract.Contract.SetDomainEmail(&_Contract.TransactOpts, _domain)
}

// UpdateBot is a paid mutator transaction binding the contract method 0xdd3d6f6c.
//
// Solidity: function updateBot(uint256 _botId, string _phoneNumber, uint8 _typeMethod, bool _status) returns()
func (_Contract *ContractTransactor) UpdateBot(opts *bind.TransactOpts, _botId *big.Int, _phoneNumber string, _typeMethod uint8, _status bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateBot", _botId, _phoneNumber, _typeMethod, _status)
}

// UpdateBot is a paid mutator transaction binding the contract method 0xdd3d6f6c.
//
// Solidity: function updateBot(uint256 _botId, string _phoneNumber, uint8 _typeMethod, bool _status) returns()
func (_Contract *ContractSession) UpdateBot(_botId *big.Int, _phoneNumber string, _typeMethod uint8, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBot(&_Contract.TransactOpts, _botId, _phoneNumber, _typeMethod, _status)
}

// UpdateBot is a paid mutator transaction binding the contract method 0xdd3d6f6c.
//
// Solidity: function updateBot(uint256 _botId, string _phoneNumber, uint8 _typeMethod, bool _status) returns()
func (_Contract *ContractTransactorSession) UpdateBot(_botId *big.Int, _phoneNumber string, _typeMethod uint8, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBot(&_Contract.TransactOpts, _botId, _phoneNumber, _typeMethod, _status)
}

// ValidateOTP is a paid mutator transaction binding the contract method 0x8e81ed21.
//
// Solidity: function validateOTP(uint256 _otp, string _identifier) returns(string publicKey, address wallet)
func (_Contract *ContractTransactor) ValidateOTP(opts *bind.TransactOpts, _otp *big.Int, _identifier string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "validateOTP", _otp, _identifier)
}

// ValidateOTP is a paid mutator transaction binding the contract method 0x8e81ed21.
//
// Solidity: function validateOTP(uint256 _otp, string _identifier) returns(string publicKey, address wallet)
func (_Contract *ContractSession) ValidateOTP(_otp *big.Int, _identifier string) (*types.Transaction, error) {
	return _Contract.Contract.ValidateOTP(&_Contract.TransactOpts, _otp, _identifier)
}

// ValidateOTP is a paid mutator transaction binding the contract method 0x8e81ed21.
//
// Solidity: function validateOTP(uint256 _otp, string _identifier) returns(string publicKey, address wallet)
func (_Contract *ContractTransactorSession) ValidateOTP(_otp *big.Int, _identifier string) (*types.Transaction, error) {
	return _Contract.Contract.ValidateOTP(&_Contract.TransactOpts, _otp, _identifier)
}

// ContractAuthenticationCompletedIterator is returned from FilterAuthenticationCompleted and is used to iterate over the raw logs and unpacked data for AuthenticationCompleted events raised by the Contract contract.
type ContractAuthenticationCompletedIterator struct {
	Event *ContractAuthenticationCompleted // Event containing the contract specifics and raw log

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
func (it *ContractAuthenticationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAuthenticationCompleted)
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
		it.Event = new(ContractAuthenticationCompleted)
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
func (it *ContractAuthenticationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAuthenticationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAuthenticationCompleted represents a AuthenticationCompleted event raised by the Contract contract.
type ContractAuthenticationCompleted struct {
	Wallet     common.Address
	Identifier common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthenticationCompleted is a free log retrieval operation binding the contract event 0xedaa5becc81ce9adeaa52cf4aa4f0152f2ec6e56d323ce09c9d702ff8ca1edc1.
//
// Solidity: event AuthenticationCompleted(address indexed wallet, string indexed identifier)
func (_Contract *ContractFilterer) FilterAuthenticationCompleted(opts *bind.FilterOpts, wallet []common.Address, identifier []string) (*ContractAuthenticationCompletedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AuthenticationCompleted", walletRule, identifierRule)
	if err != nil {
		return nil, err
	}
	return &ContractAuthenticationCompletedIterator{contract: _Contract.contract, event: "AuthenticationCompleted", logs: logs, sub: sub}, nil
}

// WatchAuthenticationCompleted is a free log subscription operation binding the contract event 0xedaa5becc81ce9adeaa52cf4aa4f0152f2ec6e56d323ce09c9d702ff8ca1edc1.
//
// Solidity: event AuthenticationCompleted(address indexed wallet, string indexed identifier)
func (_Contract *ContractFilterer) WatchAuthenticationCompleted(opts *bind.WatchOpts, sink chan<- *ContractAuthenticationCompleted, wallet []common.Address, identifier []string) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AuthenticationCompleted", walletRule, identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAuthenticationCompleted)
				if err := _Contract.contract.UnpackLog(event, "AuthenticationCompleted", log); err != nil {
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

// ParseAuthenticationCompleted is a log parse operation binding the contract event 0xedaa5becc81ce9adeaa52cf4aa4f0152f2ec6e56d323ce09c9d702ff8ca1edc1.
//
// Solidity: event AuthenticationCompleted(address indexed wallet, string indexed identifier)
func (_Contract *ContractFilterer) ParseAuthenticationCompleted(log types.Log) (*ContractAuthenticationCompleted, error) {
	event := new(ContractAuthenticationCompleted)
	if err := _Contract.contract.UnpackLog(event, "AuthenticationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAuthenticationHashStoredIterator is returned from FilterAuthenticationHashStored and is used to iterate over the raw logs and unpacked data for AuthenticationHashStored events raised by the Contract contract.
type ContractAuthenticationHashStoredIterator struct {
	Event *ContractAuthenticationHashStored // Event containing the contract specifics and raw log

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
func (it *ContractAuthenticationHashStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAuthenticationHashStored)
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
		it.Event = new(ContractAuthenticationHashStored)
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
func (it *ContractAuthenticationHashStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAuthenticationHashStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAuthenticationHashStored represents a AuthenticationHashStored event raised by the Contract contract.
type ContractAuthenticationHashStored struct {
	Wallet   common.Address
	DataHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuthenticationHashStored is a free log retrieval operation binding the contract event 0xde435a6d944a30c4364c8c1bafe5d5ce57a5a7c5278be2e4c41a5f6315d3e75e.
//
// Solidity: event AuthenticationHashStored(address indexed wallet, bytes32 dataHash)
func (_Contract *ContractFilterer) FilterAuthenticationHashStored(opts *bind.FilterOpts, wallet []common.Address) (*ContractAuthenticationHashStoredIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AuthenticationHashStored", walletRule)
	if err != nil {
		return nil, err
	}
	return &ContractAuthenticationHashStoredIterator{contract: _Contract.contract, event: "AuthenticationHashStored", logs: logs, sub: sub}, nil
}

// WatchAuthenticationHashStored is a free log subscription operation binding the contract event 0xde435a6d944a30c4364c8c1bafe5d5ce57a5a7c5278be2e4c41a5f6315d3e75e.
//
// Solidity: event AuthenticationHashStored(address indexed wallet, bytes32 dataHash)
func (_Contract *ContractFilterer) WatchAuthenticationHashStored(opts *bind.WatchOpts, sink chan<- *ContractAuthenticationHashStored, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AuthenticationHashStored", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAuthenticationHashStored)
				if err := _Contract.contract.UnpackLog(event, "AuthenticationHashStored", log); err != nil {
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

// ParseAuthenticationHashStored is a log parse operation binding the contract event 0xde435a6d944a30c4364c8c1bafe5d5ce57a5a7c5278be2e4c41a5f6315d3e75e.
//
// Solidity: event AuthenticationHashStored(address indexed wallet, bytes32 dataHash)
func (_Contract *ContractFilterer) ParseAuthenticationHashStored(log types.Log) (*ContractAuthenticationHashStored, error) {
	event := new(ContractAuthenticationHashStored)
	if err := _Contract.contract.UnpackLog(event, "AuthenticationHashStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBotAuthenticationRequestedIterator is returned from FilterBotAuthenticationRequested and is used to iterate over the raw logs and unpacked data for BotAuthenticationRequested events raised by the Contract contract.
type ContractBotAuthenticationRequestedIterator struct {
	Event *ContractBotAuthenticationRequested // Event containing the contract specifics and raw log

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
func (it *ContractBotAuthenticationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBotAuthenticationRequested)
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
		it.Event = new(ContractBotAuthenticationRequested)
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
func (it *ContractBotAuthenticationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBotAuthenticationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBotAuthenticationRequested represents a BotAuthenticationRequested event raised by the Contract contract.
type ContractBotAuthenticationRequested struct {
	Wallet       common.Address
	Otp          *big.Int
	ChatbotPhone string
	TypeMethod   uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBotAuthenticationRequested is a free log retrieval operation binding the contract event 0xde36d4fecccb0b9c355c224b65c500912aa8be614338608c0db0fc6aaef5c1b1.
//
// Solidity: event BotAuthenticationRequested(address indexed wallet, uint256 otp, string chatbotPhone, uint8 typeMethod)
func (_Contract *ContractFilterer) FilterBotAuthenticationRequested(opts *bind.FilterOpts, wallet []common.Address) (*ContractBotAuthenticationRequestedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BotAuthenticationRequested", walletRule)
	if err != nil {
		return nil, err
	}
	return &ContractBotAuthenticationRequestedIterator{contract: _Contract.contract, event: "BotAuthenticationRequested", logs: logs, sub: sub}, nil
}

// WatchBotAuthenticationRequested is a free log subscription operation binding the contract event 0xde36d4fecccb0b9c355c224b65c500912aa8be614338608c0db0fc6aaef5c1b1.
//
// Solidity: event BotAuthenticationRequested(address indexed wallet, uint256 otp, string chatbotPhone, uint8 typeMethod)
func (_Contract *ContractFilterer) WatchBotAuthenticationRequested(opts *bind.WatchOpts, sink chan<- *ContractBotAuthenticationRequested, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BotAuthenticationRequested", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBotAuthenticationRequested)
				if err := _Contract.contract.UnpackLog(event, "BotAuthenticationRequested", log); err != nil {
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

// ParseBotAuthenticationRequested is a log parse operation binding the contract event 0xde36d4fecccb0b9c355c224b65c500912aa8be614338608c0db0fc6aaef5c1b1.
//
// Solidity: event BotAuthenticationRequested(address indexed wallet, uint256 otp, string chatbotPhone, uint8 typeMethod)
func (_Contract *ContractFilterer) ParseBotAuthenticationRequested(log types.Log) (*ContractBotAuthenticationRequested, error) {
	event := new(ContractBotAuthenticationRequested)
	if err := _Contract.contract.UnpackLog(event, "BotAuthenticationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEmailAuthenticationRequestedIterator is returned from FilterEmailAuthenticationRequested and is used to iterate over the raw logs and unpacked data for EmailAuthenticationRequested events raised by the Contract contract.
type ContractEmailAuthenticationRequestedIterator struct {
	Event *ContractEmailAuthenticationRequested // Event containing the contract specifics and raw log

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
func (it *ContractEmailAuthenticationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEmailAuthenticationRequested)
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
		it.Event = new(ContractEmailAuthenticationRequested)
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
func (it *ContractEmailAuthenticationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEmailAuthenticationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEmailAuthenticationRequested represents a EmailAuthenticationRequested event raised by the Contract contract.
type ContractEmailAuthenticationRequested struct {
	Wallet                common.Address
	Otp                   *big.Int
	UserEmail             string
	TargetMailServerEmail string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterEmailAuthenticationRequested is a free log retrieval operation binding the contract event 0x353051016f669ad783bd246c4556ca65afc8ab12eea784fc026dacd257c1f532.
//
// Solidity: event EmailAuthenticationRequested(address indexed wallet, uint256 otp, string userEmail, string targetMailServerEmail)
func (_Contract *ContractFilterer) FilterEmailAuthenticationRequested(opts *bind.FilterOpts, wallet []common.Address) (*ContractEmailAuthenticationRequestedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EmailAuthenticationRequested", walletRule)
	if err != nil {
		return nil, err
	}
	return &ContractEmailAuthenticationRequestedIterator{contract: _Contract.contract, event: "EmailAuthenticationRequested", logs: logs, sub: sub}, nil
}

// WatchEmailAuthenticationRequested is a free log subscription operation binding the contract event 0x353051016f669ad783bd246c4556ca65afc8ab12eea784fc026dacd257c1f532.
//
// Solidity: event EmailAuthenticationRequested(address indexed wallet, uint256 otp, string userEmail, string targetMailServerEmail)
func (_Contract *ContractFilterer) WatchEmailAuthenticationRequested(opts *bind.WatchOpts, sink chan<- *ContractEmailAuthenticationRequested, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EmailAuthenticationRequested", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEmailAuthenticationRequested)
				if err := _Contract.contract.UnpackLog(event, "EmailAuthenticationRequested", log); err != nil {
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

// ParseEmailAuthenticationRequested is a log parse operation binding the contract event 0x353051016f669ad783bd246c4556ca65afc8ab12eea784fc026dacd257c1f532.
//
// Solidity: event EmailAuthenticationRequested(address indexed wallet, uint256 otp, string userEmail, string targetMailServerEmail)
func (_Contract *ContractFilterer) ParseEmailAuthenticationRequested(log types.Log) (*ContractEmailAuthenticationRequested, error) {
	event := new(ContractEmailAuthenticationRequested)
	if err := _Contract.contract.UnpackLog(event, "EmailAuthenticationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEmailSubCreatedIterator is returned from FilterEmailSubCreated and is used to iterate over the raw logs and unpacked data for EmailSubCreated events raised by the Contract contract.
type ContractEmailSubCreatedIterator struct {
	Event *ContractEmailSubCreated // Event containing the contract specifics and raw log

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
func (it *ContractEmailSubCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEmailSubCreated)
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
		it.Event = new(ContractEmailSubCreated)
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
func (it *ContractEmailSubCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEmailSubCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEmailSubCreated represents a EmailSubCreated event raised by the Contract contract.
type ContractEmailSubCreated struct {
	Wallet       common.Address
	PrimaryEmail string
	SubEmail     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmailSubCreated is a free log retrieval operation binding the contract event 0x38fdf5814d779f17488bd9a2022ecde6fe53d5e49ce368adb3e489f761700eb6.
//
// Solidity: event EmailSubCreated(address indexed wallet, string primaryEmail, string subEmail)
func (_Contract *ContractFilterer) FilterEmailSubCreated(opts *bind.FilterOpts, wallet []common.Address) (*ContractEmailSubCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EmailSubCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return &ContractEmailSubCreatedIterator{contract: _Contract.contract, event: "EmailSubCreated", logs: logs, sub: sub}, nil
}

// WatchEmailSubCreated is a free log subscription operation binding the contract event 0x38fdf5814d779f17488bd9a2022ecde6fe53d5e49ce368adb3e489f761700eb6.
//
// Solidity: event EmailSubCreated(address indexed wallet, string primaryEmail, string subEmail)
func (_Contract *ContractFilterer) WatchEmailSubCreated(opts *bind.WatchOpts, sink chan<- *ContractEmailSubCreated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EmailSubCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEmailSubCreated)
				if err := _Contract.contract.UnpackLog(event, "EmailSubCreated", log); err != nil {
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

// ParseEmailSubCreated is a log parse operation binding the contract event 0x38fdf5814d779f17488bd9a2022ecde6fe53d5e49ce368adb3e489f761700eb6.
//
// Solidity: event EmailSubCreated(address indexed wallet, string primaryEmail, string subEmail)
func (_Contract *ContractFilterer) ParseEmailSubCreated(log types.Log) (*ContractEmailSubCreated, error) {
	event := new(ContractEmailSubCreated)
	if err := _Contract.contract.UnpackLog(event, "EmailSubCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEmailVerifiedIterator is returned from FilterEmailVerified and is used to iterate over the raw logs and unpacked data for EmailVerified events raised by the Contract contract.
type ContractEmailVerifiedIterator struct {
	Event *ContractEmailVerified // Event containing the contract specifics and raw log

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
func (it *ContractEmailVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEmailVerified)
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
		it.Event = new(ContractEmailVerified)
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
func (it *ContractEmailVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEmailVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEmailVerified represents a EmailVerified event raised by the Contract contract.
type ContractEmailVerified struct {
	Wallet    common.Address
	FromEmail string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmailVerified is a free log retrieval operation binding the contract event 0x270b24c750778eee371dba85d0896f844ae74331a3ab04959266d8df778289c9.
//
// Solidity: event EmailVerified(address indexed wallet, string fromEmail)
func (_Contract *ContractFilterer) FilterEmailVerified(opts *bind.FilterOpts, wallet []common.Address) (*ContractEmailVerifiedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EmailVerified", walletRule)
	if err != nil {
		return nil, err
	}
	return &ContractEmailVerifiedIterator{contract: _Contract.contract, event: "EmailVerified", logs: logs, sub: sub}, nil
}

// WatchEmailVerified is a free log subscription operation binding the contract event 0x270b24c750778eee371dba85d0896f844ae74331a3ab04959266d8df778289c9.
//
// Solidity: event EmailVerified(address indexed wallet, string fromEmail)
func (_Contract *ContractFilterer) WatchEmailVerified(opts *bind.WatchOpts, sink chan<- *ContractEmailVerified, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EmailVerified", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEmailVerified)
				if err := _Contract.contract.UnpackLog(event, "EmailVerified", log); err != nil {
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

// ParseEmailVerified is a log parse operation binding the contract event 0x270b24c750778eee371dba85d0896f844ae74331a3ab04959266d8df778289c9.
//
// Solidity: event EmailVerified(address indexed wallet, string fromEmail)
func (_Contract *ContractFilterer) ParseEmailVerified(log types.Log) (*ContractEmailVerified, error) {
	event := new(ContractEmailVerified)
	if err := _Contract.contract.UnpackLog(event, "EmailVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStepVerifiedIterator is returned from FilterStepVerified and is used to iterate over the raw logs and unpacked data for StepVerified events raised by the Contract contract.
type ContractStepVerifiedIterator struct {
	Event *ContractStepVerified // Event containing the contract specifics and raw log

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
func (it *ContractStepVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStepVerified)
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
		it.Event = new(ContractStepVerified)
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
func (it *ContractStepVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStepVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStepVerified represents a StepVerified event raised by the Contract contract.
type ContractStepVerified struct {
	Wallet     common.Address
	Identifier common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStepVerified is a free log retrieval operation binding the contract event 0x2b1cc8eeea57ec5bd08351a0a9af3137ad236574e2577fdbaf2cdd401e80b396.
//
// Solidity: event StepVerified(address indexed wallet, string indexed identifier)
func (_Contract *ContractFilterer) FilterStepVerified(opts *bind.FilterOpts, wallet []common.Address, identifier []string) (*ContractStepVerifiedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "StepVerified", walletRule, identifierRule)
	if err != nil {
		return nil, err
	}
	return &ContractStepVerifiedIterator{contract: _Contract.contract, event: "StepVerified", logs: logs, sub: sub}, nil
}

// WatchStepVerified is a free log subscription operation binding the contract event 0x2b1cc8eeea57ec5bd08351a0a9af3137ad236574e2577fdbaf2cdd401e80b396.
//
// Solidity: event StepVerified(address indexed wallet, string indexed identifier)
func (_Contract *ContractFilterer) WatchStepVerified(opts *bind.WatchOpts, sink chan<- *ContractStepVerified, wallet []common.Address, identifier []string) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "StepVerified", walletRule, identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStepVerified)
				if err := _Contract.contract.UnpackLog(event, "StepVerified", log); err != nil {
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

// ParseStepVerified is a log parse operation binding the contract event 0x2b1cc8eeea57ec5bd08351a0a9af3137ad236574e2577fdbaf2cdd401e80b396.
//
// Solidity: event StepVerified(address indexed wallet, string indexed identifier)
func (_Contract *ContractFilterer) ParseStepVerified(log types.Log) (*ContractStepVerified, error) {
	event := new(ContractStepVerified)
	if err := _Contract.contract.UnpackLog(event, "StepVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
