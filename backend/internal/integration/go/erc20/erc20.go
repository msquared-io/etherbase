// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Argument is an auto generated low-level Go binding around an user-defined struct.
type Argument struct {
	Name      string
	ArgType   string
	IsIndexed bool
}

// EtherDatabaseLibBatchSetValue is an auto generated low-level Go binding around an user-defined struct.
type EtherDatabaseLibBatchSetValue struct {
	Segments []string
	DataType uint8
	Data     []byte
}

// EtherDatabaseLibNode is an auto generated low-level Go binding around an user-defined struct.
type EtherDatabaseLibNode struct {
	Path     []byte
	DataType uint8
	Data     []byte
	Exists   bool
}

// EventEmitterBatchEmitEvent is an auto generated low-level Go binding around an user-defined struct.
type EventEmitterBatchEmitEvent struct {
	Name           string
	ArgumentTopics [][32]byte
	Data           []byte
}

// EventEmitterFactoryEmitterInfo is an auto generated low-level Go binding around an user-defined struct.
type EventEmitterFactoryEmitterInfo struct {
	EmitterAddress common.Address
	Owner          common.Address
}

// EventSchema is an auto generated low-level Go binding around an user-defined struct.
type EventSchema struct {
	Name           string
	Id             string
	Args           []Argument
	EventTopic     [32]byte
	NumIndexedArgs uint8
}

// Multicall3Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// Multicall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// Multicall3Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Result struct {
	Success    bool
	ReturnData []byte
}

// RoleControlIdentityView is an auto generated low-level Go binding around an user-defined struct.
type RoleControlIdentityView struct {
	WalletAddress common.Address
	Roles         []uint8
}

// DataValidatorMetaData contains all meta data concerning the DataValidator contract.
var DataValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"validateDataEncoding\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102f98061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063477b6e6a1461002d575b5f5ffd5b61004061003b36600461016b565b610042565b005b6001816005811115610056576100566101fd565b03610074578180602001905181019061006f9190610211565b505050565b6003816005811115610088576100886101fd565b036100a1578180602001905181019061006f9190610286565b60048160058111156100b5576100b56101fd565b036100ce578180602001905181019061006f9190610286565b60028160058111156100e2576100e26101fd565b036100fb578180602001905181019061006f919061029d565b5050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561013c5761013c6100ff565b604052919050565b5f67ffffffffffffffff82111561015d5761015d6100ff565b50601f01601f191660200190565b5f5f6040838503121561017c575f5ffd5b823567ffffffffffffffff811115610192575f5ffd5b8301601f810185136101a2575f5ffd5b80356101b56101b082610144565b610113565b8181528660208385010111156101c9575f5ffd5b816020840160208301375f602083830101528094505050506020830135600681106101f2575f5ffd5b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215610221575f5ffd5b815167ffffffffffffffff811115610237575f5ffd5b8201601f81018413610247575f5ffd5b80516102556101b082610144565b818152856020838501011115610269575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f60208284031215610296575f5ffd5b5051919050565b5f602082840312156102ad575f5ffd5b815180151581146102bc575f5ffd5b939250505056fea2646970667358221220c765b345f58c570266774e2485021da716518b3726f4aeae6d3c14bcc74ec29d64736f6c634300081c0033",
}

// DataValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use DataValidatorMetaData.ABI instead.
var DataValidatorABI = DataValidatorMetaData.ABI

// DataValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataValidatorMetaData.Bin instead.
var DataValidatorBin = DataValidatorMetaData.Bin

// DeployDataValidator deploys a new Ethereum contract, binding an instance of DataValidator to it.
func DeployDataValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataValidator, error) {
	parsed, err := DataValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataValidator{DataValidatorCaller: DataValidatorCaller{contract: contract}, DataValidatorTransactor: DataValidatorTransactor{contract: contract}, DataValidatorFilterer: DataValidatorFilterer{contract: contract}}, nil
}

// DataValidator is an auto generated Go binding around an Ethereum contract.
type DataValidator struct {
	DataValidatorCaller     // Read-only binding to the contract
	DataValidatorTransactor // Write-only binding to the contract
	DataValidatorFilterer   // Log filterer for contract events
}

// DataValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataValidatorSession struct {
	Contract     *DataValidator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataValidatorCallerSession struct {
	Contract *DataValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DataValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataValidatorTransactorSession struct {
	Contract     *DataValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DataValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataValidatorRaw struct {
	Contract *DataValidator // Generic contract binding to access the raw methods on
}

// DataValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataValidatorCallerRaw struct {
	Contract *DataValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// DataValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataValidatorTransactorRaw struct {
	Contract *DataValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataValidator creates a new instance of DataValidator, bound to a specific deployed contract.
func NewDataValidator(address common.Address, backend bind.ContractBackend) (*DataValidator, error) {
	contract, err := bindDataValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataValidator{DataValidatorCaller: DataValidatorCaller{contract: contract}, DataValidatorTransactor: DataValidatorTransactor{contract: contract}, DataValidatorFilterer: DataValidatorFilterer{contract: contract}}, nil
}

// NewDataValidatorCaller creates a new read-only instance of DataValidator, bound to a specific deployed contract.
func NewDataValidatorCaller(address common.Address, caller bind.ContractCaller) (*DataValidatorCaller, error) {
	contract, err := bindDataValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataValidatorCaller{contract: contract}, nil
}

// NewDataValidatorTransactor creates a new write-only instance of DataValidator, bound to a specific deployed contract.
func NewDataValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DataValidatorTransactor, error) {
	contract, err := bindDataValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataValidatorTransactor{contract: contract}, nil
}

// NewDataValidatorFilterer creates a new log filterer instance of DataValidator, bound to a specific deployed contract.
func NewDataValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DataValidatorFilterer, error) {
	contract, err := bindDataValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataValidatorFilterer{contract: contract}, nil
}

// bindDataValidator binds a generic wrapper to an already deployed contract.
func bindDataValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataValidator *DataValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataValidator.Contract.DataValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataValidator *DataValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataValidator.Contract.DataValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataValidator *DataValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataValidator.Contract.DataValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataValidator *DataValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataValidator *DataValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataValidator *DataValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataValidator.Contract.contract.Transact(opts, method, params...)
}

// ValidateDataEncoding is a free data retrieval call binding the contract method 0x477b6e6a.
//
// Solidity: function validateDataEncoding(bytes data, uint8 dataType) pure returns()
func (_DataValidator *DataValidatorCaller) ValidateDataEncoding(opts *bind.CallOpts, data []byte, dataType uint8) error {
	var out []interface{}
	err := _DataValidator.contract.Call(opts, &out, "validateDataEncoding", data, dataType)

	if err != nil {
		return err
	}

	return err

}

// ValidateDataEncoding is a free data retrieval call binding the contract method 0x477b6e6a.
//
// Solidity: function validateDataEncoding(bytes data, uint8 dataType) pure returns()
func (_DataValidator *DataValidatorSession) ValidateDataEncoding(data []byte, dataType uint8) error {
	return _DataValidator.Contract.ValidateDataEncoding(&_DataValidator.CallOpts, data, dataType)
}

// ValidateDataEncoding is a free data retrieval call binding the contract method 0x477b6e6a.
//
// Solidity: function validateDataEncoding(bytes data, uint8 dataType) pure returns()
func (_DataValidator *DataValidatorCallerSession) ValidateDataEncoding(data []byte, dataType uint8) error {
	return _DataValidator.Contract.ValidateDataEncoding(&_DataValidator.CallOpts, data, dataType)
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051610409380380610409833981016040819052602b916055565b600280546001600160a01b031916339081179091555f82815590815260016020526040902055606b565b5f602082840312156064575f5ffd5b5051919050565b610391806100785f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806318160ddd1461005957806327e235e31461007457806340c10f19146100935780638da5cb5b146100a8578063a9059cbb146100d3575b5f5ffd5b6100615f5481565b6040519081526020015b60405180910390f35b6100616100823660046102d9565b60016020525f908152604090205481565b6100a66100a13660046102f9565b6100f6565b005b6002546100bb906001600160a01b031681565b6040516001600160a01b03909116815260200161006b565b6100e66100e13660046102f9565b6101d0565b604051901515815260200161006b565b6002546001600160a01b0316331461014b5760405162461bcd60e51b815260206004820152601360248201527213db9b1e481bdddb995c8818d85b881b5a5b9d606a1b60448201526064015b60405180910390fd5b805f5f82825461015b9190610335565b90915550506001600160a01b0382165f9081526001602052604081208054839290610187908490610335565b90915550506040518181526001600160a01b038316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b335f908152600160205260408120548211156102255760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606401610142565b335f9081526001602052604081208054849290610243908490610348565b90915550506001600160a01b0383165f908152600160205260408120805484929061026f908490610335565b90915550506040518281526001600160a01b0384169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35060015b92915050565b80356001600160a01b03811681146102d4575f5ffd5b919050565b5f602082840312156102e9575f5ffd5b6102f2826102be565b9392505050565b5f5f6040838503121561030a575f5ffd5b610313836102be565b946020939093013593505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156102b8576102b8610321565b818103818111156102b8576102b861032156fea2646970667358221220ce2dde0318895a7cdf467c22a517fe552730df10e63a0efaef12599831b2e57464736f6c634300081c0033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _initialSupply *big.Int) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ERC20 *ERC20Caller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ERC20 *ERC20Session) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20.Contract.Balances(&_ERC20.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20.Contract.Balances(&_ERC20.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Session) Owner() (common.Address, error) {
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20CallerSession) Owner() (common.Address, error) {
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_ERC20 *ERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_ERC20 *ERC20Session) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Mint(&_ERC20.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_ERC20 *ERC20TransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Mint(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthDBUpdaterMetaData contains all meta data concerning the EthDBUpdater contract.
var EthDBUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"}]",
}

// EthDBUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use EthDBUpdaterMetaData.ABI instead.
var EthDBUpdaterABI = EthDBUpdaterMetaData.ABI

// EthDBUpdater is an auto generated Go binding around an Ethereum contract.
type EthDBUpdater struct {
	EthDBUpdaterCaller     // Read-only binding to the contract
	EthDBUpdaterTransactor // Write-only binding to the contract
	EthDBUpdaterFilterer   // Log filterer for contract events
}

// EthDBUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthDBUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDBUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthDBUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDBUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthDBUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDBUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthDBUpdaterSession struct {
	Contract     *EthDBUpdater     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthDBUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthDBUpdaterCallerSession struct {
	Contract *EthDBUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EthDBUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthDBUpdaterTransactorSession struct {
	Contract     *EthDBUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EthDBUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthDBUpdaterRaw struct {
	Contract *EthDBUpdater // Generic contract binding to access the raw methods on
}

// EthDBUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthDBUpdaterCallerRaw struct {
	Contract *EthDBUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// EthDBUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthDBUpdaterTransactorRaw struct {
	Contract *EthDBUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthDBUpdater creates a new instance of EthDBUpdater, bound to a specific deployed contract.
func NewEthDBUpdater(address common.Address, backend bind.ContractBackend) (*EthDBUpdater, error) {
	contract, err := bindEthDBUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthDBUpdater{EthDBUpdaterCaller: EthDBUpdaterCaller{contract: contract}, EthDBUpdaterTransactor: EthDBUpdaterTransactor{contract: contract}, EthDBUpdaterFilterer: EthDBUpdaterFilterer{contract: contract}}, nil
}

// NewEthDBUpdaterCaller creates a new read-only instance of EthDBUpdater, bound to a specific deployed contract.
func NewEthDBUpdaterCaller(address common.Address, caller bind.ContractCaller) (*EthDBUpdaterCaller, error) {
	contract, err := bindEthDBUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthDBUpdaterCaller{contract: contract}, nil
}

// NewEthDBUpdaterTransactor creates a new write-only instance of EthDBUpdater, bound to a specific deployed contract.
func NewEthDBUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*EthDBUpdaterTransactor, error) {
	contract, err := bindEthDBUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthDBUpdaterTransactor{contract: contract}, nil
}

// NewEthDBUpdaterFilterer creates a new log filterer instance of EthDBUpdater, bound to a specific deployed contract.
func NewEthDBUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*EthDBUpdaterFilterer, error) {
	contract, err := bindEthDBUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthDBUpdaterFilterer{contract: contract}, nil
}

// bindEthDBUpdater binds a generic wrapper to an already deployed contract.
func bindEthDBUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthDBUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDBUpdater *EthDBUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthDBUpdater.Contract.EthDBUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDBUpdater *EthDBUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDBUpdater.Contract.EthDBUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDBUpdater *EthDBUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDBUpdater.Contract.EthDBUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDBUpdater *EthDBUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthDBUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDBUpdater *EthDBUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDBUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDBUpdater *EthDBUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDBUpdater.Contract.contract.Transact(opts, method, params...)
}

// EthDBUpdaterEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the EthDBUpdater contract.
type EthDBUpdaterEthDBPathUpdateIterator struct {
	Event *EthDBUpdaterEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *EthDBUpdaterEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDBUpdaterEthDBPathUpdate)
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
		it.Event = new(EthDBUpdaterEthDBPathUpdate)
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
func (it *EthDBUpdaterEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDBUpdaterEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDBUpdaterEthDBPathUpdate represents a EthDBPathUpdate event raised by the EthDBUpdater contract.
type EthDBUpdaterEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EthDBUpdater *EthDBUpdaterFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*EthDBUpdaterEthDBPathUpdateIterator, error) {

	logs, sub, err := _EthDBUpdater.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &EthDBUpdaterEthDBPathUpdateIterator{contract: _EthDBUpdater.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EthDBUpdater *EthDBUpdaterFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *EthDBUpdaterEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _EthDBUpdater.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDBUpdaterEthDBPathUpdate)
				if err := _EthDBUpdater.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EthDBUpdater *EthDBUpdaterFilterer) ParseEthDBPathUpdate(log types.Log) (*EthDBUpdaterEthDBPathUpdate, error) {
	event := new(EthDBUpdaterEthDBPathUpdate)
	if err := _EthDBUpdater.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthDBUpdaterEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the EthDBUpdater contract.
type EthDBUpdaterEthDBUpdateIterator struct {
	Event *EthDBUpdaterEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *EthDBUpdaterEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDBUpdaterEthDBUpdate)
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
		it.Event = new(EthDBUpdaterEthDBUpdate)
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
func (it *EthDBUpdaterEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDBUpdaterEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDBUpdaterEthDBUpdate represents a EthDBUpdate event raised by the EthDBUpdater contract.
type EthDBUpdaterEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_EthDBUpdater *EthDBUpdaterFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*EthDBUpdaterEthDBUpdateIterator, error) {

	logs, sub, err := _EthDBUpdater.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &EthDBUpdaterEthDBUpdateIterator{contract: _EthDBUpdater.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_EthDBUpdater *EthDBUpdaterFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *EthDBUpdaterEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _EthDBUpdater.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDBUpdaterEthDBUpdate)
				if err := _EthDBUpdater.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_EthDBUpdater *EthDBUpdaterFilterer) ParseEthDBUpdate(log types.Log) (*EthDBUpdaterEthDBUpdate, error) {
	event := new(EthDBUpdaterEthDBUpdate)
	if err := _EthDBUpdater.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherDatabaseLibMetaData contains all meta data concerning the EtherDatabaseLib contract.
var EtherDatabaseLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"InvalidDataEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDataType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PathNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NewSegment\",\"type\":\"event\"}]",
	Bin: "0x611e6b610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610610127575f3560e01c80639a7044d7116100b4578063dea25e9f11610079578063dea25e9f146102d7578063e3fa71ec146102ea578063e41d307214610309578063e955702c14610328578063f2bffc131461033b575f5ffd5b80639a7044d7146102385780639e72831114610257578063b64ee37d14610278578063c49fed3214610298578063c837f232146102b8575f5ffd5b806334ee1c56116100fa57806334ee1c561461019857806340efd129146101b757806355517745146101d65780635d0aad23146101f65780638b7b2e1714610219575f5ffd5b80631b33bcaf1461012b5780631c74ab571461014c5780632519093f1461017257806332a1c4bb14610185575b5f5ffd5b818015610136575f5ffd5b5061014a6101453660046113e2565b61035a565b005b61015f61015a366004611430565b61038f565b6040519081526020015b60405180910390f35b61015f610180366004611430565b6103ba565b61015f6101933660046114b4565b6103c8565b8180156101a3575f5ffd5b5061015f6101b23660046114b4565b6103f5565b8180156101c2575f5ffd5b5061014a6101d1366004611430565b610401565b6101e96101e4366004611430565b61051b565b604051610169919061151c565b610209610204366004611430565b61053d565b6040519015158152602001610169565b818015610224575f5ffd5b5061014a610233366004611430565b61057b565b818015610243575f5ffd5b5061014a61025236600461152e565b61065f565b61026a610265366004611430565b61067a565b6040516101699291906115d9565b61028b6102863660046115f8565b610734565b604051610169919061160f565b6102ab6102a63660046115f8565b610a11565b60405161016991906116b9565b8180156102c3575f5ffd5b5061014a6102d2366004611723565b610ae9565b6102096102e5366004611430565b610b33565b8180156102f5575f5ffd5b5061014a61030436600461152e565b610b54565b818015610314575f5ffd5b5061014a6103233660046117bb565b610b97565b6101e9610336366004611430565b610bb4565b818015610346575f5ffd5b5061014a6103553660046113e2565b610bc3565b61038984848460038560405160200161037591815260200190565b604051602081830303815290604052610bde565b50505050565b5f61039d8484846004610dd4565b8060200190518101906103b09190611815565b90505b9392505050565b5f61039d8484846003610dd4565b5f8360020183836040516103dd92919061182c565b90815260200160405180910390205490509392505050565b5f6103b084848461105a565b5f61040d848484611162565b905083600701816040516104219190611852565b9081526040519081900360200190205460ff16610451576040516357509b2760e11b815260040160405180910390fd5b5f84600601826040516104649190611852565b90815260200160405180910390205490505f85600501828154811061048b5761048b61185d565b905f5260205f2090600402016003015f6101000a81548160ff02191690831515021790555084600601826040516104c29190611852565b90815260200160405180910390205f905584600701826040516104e59190611852565b908152604080519182900360209081018320805460ff19169055820190525f8082526105149186918691611228565b5050505050565b606061052a8484846001610dd4565b8060200190518101906103b09190611885565b5f5f61054a858585611162565b9050846007018160405161055e9190611852565b9081526040519081900360200190205460ff169150509392505050565b5f5b81811015610389576106578484848481811061059b5761059b61185d565b90506020028101906105ad9190611935565b6105b79080611953565b8686868181106105c9576105c961185d565b90506020028101906105db9190611935565b6105ec906040810190602001611998565b8787878181106105fe576105fe61185d565b90506020028101906106109190611935565b61061e9060408101906119b1565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bde92505050565b60010161057d565b61051485858560018686604051602001610375929190611a1b565b5f60605f610689868686611162565b9050856007018160405161069d9190611852565b9081526040519081900360200190205460ff166106cc57505060408051602081019091525f808252915061072c565b5f6106d98787875f610dd4565b90508660050187600601836040516106f19190611852565b908152602001604051809103902054815481106107105761071061185d565b5f91825260209091206001600490920201015460ff1693509150505b935093915050565b60605f805b600584015481101561078b5783600501818154811061075a5761075a61185d565b5f91825260209091206003600490920201015460ff1615610783578161077f81611a42565b9250505b600101610739565b505f816001600160401b038111156107a5576107a5611871565b60405190808252806020026020018201604052801561080057816020015b6107ed604080516080810190915260608152602081015f8152606060208201525f60409091015290565b8152602001906001900390816107c35790505b5090505f805b6005860154811015610a07578560050181815481106108275761082761185d565b5f91825260209091206003600490920201015460ff16156109ff578560050181815481106108575761085761185d565b905f5260205f2090600402016040518060800160405290815f8201805461087d90611a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546108a990611a5a565b80156108f45780601f106108cb576101008083540402835291602001916108f4565b820191905f5260205f20905b8154815290600101906020018083116108d757829003601f168201915b5050509183525050600182015460209091019060ff16600581111561091b5761091b6115a5565b600581111561092c5761092c6115a5565b815260200160028201805461094090611a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461096c90611a5a565b80156109b75780601f1061098e576101008083540402835291602001916109b7565b820191905f5260205f20905b81548152906001019060200180831161099a57829003601f168201915b50505091835250506003919091015460ff16151560209091015283518490849081106109e5576109e561185d565b602002602001018190525081806109fb90611a42565b9250505b600101610806565b5090949350505050565b606081600301805480602002602001604051908101604052809291908181526020015f905b82821015610ade578382905f5260205f20018054610a5390611a5a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7f90611a5a565b8015610aca5780601f10610aa157610100808354040283529160200191610aca565b820191905f5260205f20905b815481529060010190602001808311610aad57829003601f168201915b505050505081526020019060010190610a36565b505050509050919050565b610b2b8686868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bde92505050565b505050505050565b5f610b418484846002610dd4565b8060200190518101906103b09190611a92565b610514858585600586868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bde92505050565b610389848484600285604051602001610375911515815260200190565b60606103b08484846005610dd4565b61038984848460048560405160200161037591815260200190565b5f610bea86868661127c565b90508560070181604051610bfe9190611852565b9081526040519081900360200190205460ff1615610c9e575f8660060182604051610c299190611852565b90815260200160405180910390205490505f876005018281548110610c5057610c5061185d565b905f5260205f209060040201905083816002019081610c6f9190611af6565b5060018082018054879260ff1990911690836005811115610c9257610c926115a5565b02179055505050610dc8565b5f6040518060800160405280838152602001856005811115610cc257610cc26115a5565b815260208082018690526001604090920182905260058a01805492830181555f9081522082519293508392600490920201908190610d009082611af6565b50602082015160018083018054909160ff1990911690836005811115610d2857610d286115a5565b021790555060408201516002820190610d419082611af6565b50606091909101516003909101805460ff19169115159190911790556005870154610d6e90600190611bb0565b8760060183604051610d809190611852565b90815260200160405180910390208190555060018760070183604051610da69190611852565b908152604051908190036020019020805491151560ff19909216919091179055505b610b2b85858585611228565b60605f610de2868686611162565b90508560070181604051610df69190611852565b9081526040519081900360200190205460ff16610e26576040516357509b2760e11b815260040160405180910390fd5b5f866005018760060183604051610e3d9190611852565b90815260200160405180910390205481548110610e5c57610e5c61185d565b5f9182526020909120600490910201600381015490915060ff16610e93576040516357509b2760e11b815260040160405180910390fd5b5f846005811115610ea657610ea66115a5565b14158015610edd5750836005811115610ec157610ec16115a5565b600182015460ff166005811115610eda57610eda6115a5565b14155b15610f1157600181015460405163254c060560e21b8152610f0891869160ff90911690600401611bc9565b60405180910390fd5b5f846005811115610f2457610f246115a5565b14158015610f4457506005846005811115610f4157610f416115a5565b14155b15610fc35786546040516323bdb73560e11b81526001600160a01b039091169063477b6e6a90610f7d9060028501908890600401611be4565b5f6040518083038186803b158015610f93575f5ffd5b505afa925050508015610fa4575060015b610fc35783604051631e9c993560e21b8152600401610f089190611c74565b806002018054610fd290611a5a565b80601f0160208091040260200160405190810160405280929190818152602001828054610ffe90611a5a565b80156110495780601f1061102057610100808354040283529160200191611049565b820191905f5260205f20905b81548152906001019060200180831161102c57829003601f168201915b505050505092505050949350505050565b5f5f84600201848460405161107092919061182c565b9081526040805160209281900383019020545f8181526004890190935291205490915060ff16156110a25790506103b3565b6001850180549081905f6110b583611a42565b9190505550808660020186866040516110cf92919061182c565b90815260405160209181900382019020919091556003870180546001810182555f91825291902001611102858783611c82565b505f81815260048701602052604090819020805460ff191660011790555181907f4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778906111519088908890611a1b565b60405180910390a291506103b39050565b604080515f808252602082019092526060915b838110156111fb575f866002018686848181106111945761119461185d565b90506020028101906111a691906119b1565b6040516111b492919061182c565b9081526020016040518091039020549050826111cf826112fd565b6040516020016111e0929190611d3b565b60408051601f19818403018152919052925050600101611175565b5060405161120f9082905f90602001611d4f565b60408051808303601f1901815291905295945050505050565b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a84848385600581111561125e5761125e6115a5565b60405161126e9493929190611d73565b60405180910390a150505050565b604080515f808252602082019092526060915b838110156111fb575f6112c5878787858181106112ae576112ae61185d565b90506020028101906112c091906119b1565b61105a565b9050826112d1826112fd565b6040516020016112e2929190611d3b565b60408051601f1981840301815291905292505060010161128f565b6060815f03611328576040515f60208201526021016040516020818303038152906040529050919050565b6060825b6080811061136a578181607f1660801760f81b604051602001611350929190611d4f565b60408051601f19818403018152919052915060071c61132c565b8181607f1660f81b604051602001611383929190611d4f565b60408051601f19818403018152919052949350505050565b5f5f83601f8401126113ab575f5ffd5b5081356001600160401b038111156113c1575f5ffd5b6020830191508360208260051b85010111156113db575f5ffd5b9250929050565b5f5f5f5f606085870312156113f5575f5ffd5b8435935060208501356001600160401b03811115611411575f5ffd5b61141d8782880161139b565b9598909750949560400135949350505050565b5f5f5f60408486031215611442575f5ffd5b8335925060208401356001600160401b0381111561145e575f5ffd5b61146a8682870161139b565b9497909650939450505050565b5f5f83601f840112611487575f5ffd5b5081356001600160401b0381111561149d575f5ffd5b6020830191508360208285010111156113db575f5ffd5b5f5f5f604084860312156114c6575f5ffd5b8335925060208401356001600160401b038111156114e2575f5ffd5b61146a86828701611477565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6103b360208301846114ee565b5f5f5f5f5f60608688031215611542575f5ffd5b8535945060208601356001600160401b0381111561155e575f5ffd5b61156a8882890161139b565b90955093505060408601356001600160401b03811115611588575f5ffd5b61159488828901611477565b969995985093965092949392505050565b634e487b7160e01b5f52602160045260245ffd5b600681106115d557634e487b7160e01b5f52602160045260245ffd5b9052565b6115e381846115b9565b604060208201525f6103b060408301846114ee565b5f60208284031215611608575f5ffd5b5035919050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156116ad57603f19878603018452815180516080875261165b60808801826114ee565b9050602082015161166f60208901826115b9565b506040820151878203604089015261168782826114ee565b606093840151151598909301979097525094506020938401939190910190600101611635565b50929695505050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156116ad57603f198786030184526116fb8583516114ee565b945060209384019391909101906001016116df565b80356006811061171e575f5ffd5b919050565b5f5f5f5f5f5f60808789031215611738575f5ffd5b8635955060208701356001600160401b03811115611754575f5ffd5b61176089828a0161139b565b9096509450611773905060408801611710565b925060608701356001600160401b0381111561178d575f5ffd5b61179989828a01611477565b979a9699509497509295939492505050565b80151581146117b8575f5ffd5b50565b5f5f5f5f606085870312156117ce575f5ffd5b8435935060208501356001600160401b038111156117ea575f5ffd5b6117f68782880161139b565b909450925050604085013561180a816117ab565b939692955090935050565b5f60208284031215611825575f5ffd5b5051919050565b818382375f9101908152919050565b5f81518060208401855e5f93019283525090919050565b5f6103b3828461183b565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215611895575f5ffd5b81516001600160401b038111156118aa575f5ffd5b8201601f810184136118ba575f5ffd5b80516001600160401b038111156118d3576118d3611871565b604051601f8201601f19908116603f011681016001600160401b038111828210171561190157611901611871565b604052818152828201602001861015611918575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f8235605e19833603018112611949575f5ffd5b9190910192915050565b5f5f8335601e19843603018112611968575f5ffd5b8301803591506001600160401b03821115611981575f5ffd5b6020019150600581901b36038213156113db575f5ffd5b5f602082840312156119a8575f5ffd5b6103b382611710565b5f5f8335601e198436030181126119c6575f5ffd5b8301803591506001600160401b038211156119df575f5ffd5b6020019150368190038213156113db575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f6103b06020830184866119f3565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611a5357611a53611a2e565b5060010190565b600181811c90821680611a6e57607f821691505b602082108103611a8c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215611aa2575f5ffd5b81516103b3816117ab565b601f821115611af157805f5260205f20601f840160051c81016020851015611ad25750805b601f840160051c820191505b81811015610514575f8155600101611ade565b505050565b81516001600160401b03811115611b0f57611b0f611871565b611b2381611b1d8454611a5a565b84611aad565b6020601f821160018114611b55575f8315611b3e5750848201515b5f19600385901b1c1916600184901b178455610514565b5f84815260208120601f198516915b82811015611b845787850151825560209485019460019092019101611b64565b5084821015611ba157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b81810381811115611bc357611bc3611a2e565b92915050565b60408101611bd782856115b9565b6103b360208301846115b9565b604081525f5f8454611bf581611a5a565b806040860152600182165f8114611c135760018114611c2f57611c60565b60ff1983166060870152606082151560051b8701019350611c60565b875f5260205f205f5b83811015611c5757815488820160600152600190910190602001611c38565b87016060019450505b505050809150506103b360208301846115b9565b60208101611bc382846115b9565b6001600160401b03831115611c9957611c99611871565b611cad83611ca78354611a5a565b83611aad565b5f601f841160018114611cde575f8515611cc75750838201355b5f19600387901b1c1916600186901b178355610514565b5f83815260208120601f198716915b82811015611d0d5786850135825560209485019460019092019101611ced565b5086821015611d29575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f6103b0611d49838661183b565b8461183b565b5f611d5a828561183b565b6001600160f81b03199390931683525050600101919050565b606080825281018490525f6080600586901b830181019083018783601e1936839003015b89821015611e0457868503607f190184528235818112611db5575f5ffd5b8b016020810190356001600160401b03811115611dd0575f5ffd5b803603821315611dde575f5ffd5b611de98782846119f3565b96505050602083019250602084019350600182019150611d97565b505050508281036020840152611e1a81866114ee565b915050611e2c604083018460ff169052565b9594505050505056fea264697066735822122069b99780d252f54e5a7fcc24ee7040d3e50f06eb4dac5a54ff9268030e6c575264736f6c634300081c0033",
}

// EtherDatabaseLibABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherDatabaseLibMetaData.ABI instead.
var EtherDatabaseLibABI = EtherDatabaseLibMetaData.ABI

// EtherDatabaseLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherDatabaseLibMetaData.Bin instead.
var EtherDatabaseLibBin = EtherDatabaseLibMetaData.Bin

// DeployEtherDatabaseLib deploys a new Ethereum contract, binding an instance of EtherDatabaseLib to it.
func DeployEtherDatabaseLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EtherDatabaseLib, error) {
	parsed, err := EtherDatabaseLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherDatabaseLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherDatabaseLib{EtherDatabaseLibCaller: EtherDatabaseLibCaller{contract: contract}, EtherDatabaseLibTransactor: EtherDatabaseLibTransactor{contract: contract}, EtherDatabaseLibFilterer: EtherDatabaseLibFilterer{contract: contract}}, nil
}

// EtherDatabaseLib is an auto generated Go binding around an Ethereum contract.
type EtherDatabaseLib struct {
	EtherDatabaseLibCaller     // Read-only binding to the contract
	EtherDatabaseLibTransactor // Write-only binding to the contract
	EtherDatabaseLibFilterer   // Log filterer for contract events
}

// EtherDatabaseLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherDatabaseLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDatabaseLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherDatabaseLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDatabaseLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherDatabaseLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDatabaseLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherDatabaseLibSession struct {
	Contract     *EtherDatabaseLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherDatabaseLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherDatabaseLibCallerSession struct {
	Contract *EtherDatabaseLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EtherDatabaseLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherDatabaseLibTransactorSession struct {
	Contract     *EtherDatabaseLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EtherDatabaseLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherDatabaseLibRaw struct {
	Contract *EtherDatabaseLib // Generic contract binding to access the raw methods on
}

// EtherDatabaseLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherDatabaseLibCallerRaw struct {
	Contract *EtherDatabaseLibCaller // Generic read-only contract binding to access the raw methods on
}

// EtherDatabaseLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherDatabaseLibTransactorRaw struct {
	Contract *EtherDatabaseLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherDatabaseLib creates a new instance of EtherDatabaseLib, bound to a specific deployed contract.
func NewEtherDatabaseLib(address common.Address, backend bind.ContractBackend) (*EtherDatabaseLib, error) {
	contract, err := bindEtherDatabaseLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseLib{EtherDatabaseLibCaller: EtherDatabaseLibCaller{contract: contract}, EtherDatabaseLibTransactor: EtherDatabaseLibTransactor{contract: contract}, EtherDatabaseLibFilterer: EtherDatabaseLibFilterer{contract: contract}}, nil
}

// NewEtherDatabaseLibCaller creates a new read-only instance of EtherDatabaseLib, bound to a specific deployed contract.
func NewEtherDatabaseLibCaller(address common.Address, caller bind.ContractCaller) (*EtherDatabaseLibCaller, error) {
	contract, err := bindEtherDatabaseLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseLibCaller{contract: contract}, nil
}

// NewEtherDatabaseLibTransactor creates a new write-only instance of EtherDatabaseLib, bound to a specific deployed contract.
func NewEtherDatabaseLibTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherDatabaseLibTransactor, error) {
	contract, err := bindEtherDatabaseLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseLibTransactor{contract: contract}, nil
}

// NewEtherDatabaseLibFilterer creates a new log filterer instance of EtherDatabaseLib, bound to a specific deployed contract.
func NewEtherDatabaseLibFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherDatabaseLibFilterer, error) {
	contract, err := bindEtherDatabaseLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseLibFilterer{contract: contract}, nil
}

// bindEtherDatabaseLib binds a generic wrapper to an already deployed contract.
func bindEtherDatabaseLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherDatabaseLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDatabaseLib *EtherDatabaseLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherDatabaseLib.Contract.EtherDatabaseLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDatabaseLib *EtherDatabaseLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.EtherDatabaseLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDatabaseLib *EtherDatabaseLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.EtherDatabaseLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDatabaseLib *EtherDatabaseLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherDatabaseLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDatabaseLib *EtherDatabaseLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDatabaseLib *EtherDatabaseLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.contract.Transact(opts, method, params...)
}

// EtherDatabaseLibEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the EtherDatabaseLib contract.
type EtherDatabaseLibEthDBPathUpdateIterator struct {
	Event *EtherDatabaseLibEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *EtherDatabaseLibEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDatabaseLibEthDBPathUpdate)
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
		it.Event = new(EtherDatabaseLibEthDBPathUpdate)
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
func (it *EtherDatabaseLibEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDatabaseLibEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDatabaseLibEthDBPathUpdate represents a EthDBPathUpdate event raised by the EtherDatabaseLib contract.
type EtherDatabaseLibEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherDatabaseLib *EtherDatabaseLibFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*EtherDatabaseLibEthDBPathUpdateIterator, error) {

	logs, sub, err := _EtherDatabaseLib.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseLibEthDBPathUpdateIterator{contract: _EtherDatabaseLib.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherDatabaseLib *EtherDatabaseLibFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *EtherDatabaseLibEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _EtherDatabaseLib.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDatabaseLibEthDBPathUpdate)
				if err := _EtherDatabaseLib.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherDatabaseLib *EtherDatabaseLibFilterer) ParseEthDBPathUpdate(log types.Log) (*EtherDatabaseLibEthDBPathUpdate, error) {
	event := new(EtherDatabaseLibEthDBPathUpdate)
	if err := _EtherDatabaseLib.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherDatabaseLibNewSegmentIterator is returned from FilterNewSegment and is used to iterate over the raw logs and unpacked data for NewSegment events raised by the EtherDatabaseLib contract.
type EtherDatabaseLibNewSegmentIterator struct {
	Event *EtherDatabaseLibNewSegment // Event containing the contract specifics and raw log

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
func (it *EtherDatabaseLibNewSegmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDatabaseLibNewSegment)
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
		it.Event = new(EtherDatabaseLibNewSegment)
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
func (it *EtherDatabaseLibNewSegmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDatabaseLibNewSegmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDatabaseLibNewSegment represents a NewSegment event raised by the EtherDatabaseLib contract.
type EtherDatabaseLibNewSegment struct {
	Segment string
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSegment is a free log retrieval operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherDatabaseLib *EtherDatabaseLibFilterer) FilterNewSegment(opts *bind.FilterOpts, id []*big.Int) (*EtherDatabaseLibNewSegmentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherDatabaseLib.contract.FilterLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseLibNewSegmentIterator{contract: _EtherDatabaseLib.contract, event: "NewSegment", logs: logs, sub: sub}, nil
}

// WatchNewSegment is a free log subscription operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherDatabaseLib *EtherDatabaseLibFilterer) WatchNewSegment(opts *bind.WatchOpts, sink chan<- *EtherDatabaseLibNewSegment, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherDatabaseLib.contract.WatchLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDatabaseLibNewSegment)
				if err := _EtherDatabaseLib.contract.UnpackLog(event, "NewSegment", log); err != nil {
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

// ParseNewSegment is a log parse operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherDatabaseLib *EtherDatabaseLibFilterer) ParseNewSegment(log types.Log) (*EtherDatabaseLibNewSegment, error) {
	event := new(EtherDatabaseLibNewSegment)
	if err := _EtherDatabaseLib.contract.UnpackLog(event, "NewSegment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterMetaData contains all meta data concerning the EventEmitter contract.
var EventEmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EventNameAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EventNameNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfTopics\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyIndexedargs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NewSegment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"SchemaRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"initialRoles\",\"type\":\"uint8[]\"}],\"name\":\"createWalletIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"deleteIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"ArgumentTopics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitEvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"argumentTopics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEventEmitter.BatchEmitEvent[]\",\"name\":\"events\",\"type\":\"tuple[]\"}],\"name\":\"emitEvents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"encodePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventSchemas\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllIdentities\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structRoleControl.IdentityView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSegments\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structEtherDatabaseLib.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getEventSchemaFromId\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getEventSchemaFromName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getOrCreateSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredEventNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"hasEntry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"idToName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"registerEventSchema\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"removeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEtherDatabaseLib.BatchSetValue[]\",\"name\":\"values\",\"type\":\"tuple[]\"}],\"name\":\"setValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051613abf380380613abf83398101604081905261002e916103ae565b61003960088361005e565b600180555f80546001600160a01b0319166001600160a01b0383161790555050610407565b6002820180546001600160a01b0319166001600160a01b03831617905560408051600480825260a082019092525f91602082016080803683370190505090506002815f815181106100b1576100b16103df565b602002602001019060038111156100ca576100ca6103f3565b908160038111156100dd576100dd6103f3565b815250505f816001815181106100f5576100f56103df565b6020026020010190600381111561010e5761010e6103f3565b90816003811115610121576101216103f3565b8152505060018160028151811061013a5761013a6103df565b60200260200101906003811115610153576101536103f3565b90816003811115610166576101666103f3565b8152505060038160038151811061017f5761017f6103df565b60200260200101906003811115610198576101986103f3565b908160038111156101ab576101ab6103f3565b9052506101b98383836101be565b505050565b6001600160a01b0382165f9081526020849052604090205460ff16156101f7576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602085815260408220805460ff191660019081179091558681018054918201815583529082200180546001600160a01b0319169092179091555b815181101561027a57610272848484848151811061025f5761025f6103df565b60200260200101516102b360201b60201c565b60010161023f565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a2505050565b6001600160a01b0382165f9081526020849052604090205460ff166102eb5760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208481526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a918202191690836003811115610343576103436103f3565b021790555081600381111561035a5761035a6103f3565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a350505050565b80516001600160a01b03811681146103a9575f5ffd5b919050565b5f5f604083850312156103bf575f5ffd5b6103c883610393565b91506103d660208401610393565b90509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b6136ab806104145f395ff3fe608060405234801561000f575f5ffd5b5060043610610208575f3560e01c806369da55d21161011f578063a4dbe8a0116100a9578063c8740be111610079578063c8740be1146104b3578063d50d542e146104c6578063dced4783146104d9578063e6149cd0146104ec578063f30826f0146104ff575f5ffd5b8063a4dbe8a014610467578063a8d29d1d1461047a578063ab8412831461048d578063c76853b8146104a0575f5ffd5b80638ae0b3fa116100ef5780638ae0b3fa146104005780638c710456146104135780638da5cb5b14610426578063960a9678146104415780639e85b3ea14610454575f5ffd5b806369da55d2146103b257806371934ef0146103c55780637eeede35146103d85780638998fbe8146103ed575f5ffd5b80633aa311dd116101a057806349ab123d1161017057806349ab123d146103385780634e0a38dc14610358578063562b7be91461036b57806360778d951461037e57806364ad6c881461039f575f5ffd5b80633aa311dd146102dc5780633e840236146102fd5780633f54d8bf146103105780634151df6114610325575f5ffd5b80631a66ac39116101db5780631a66ac39146102715780631bfa8601146102845780632f2545d2146102995780632fc37787146102bc575f5ffd5b80630277e0231461020c5780630761b8d014610234578063077d3c031461024957806317be85c31461025c575b5f5ffd5b61021f61021a366004612208565b610507565b60405190151581526020015b60405180910390f35b6102476102423660046122c0565b61052d565b005b610247610257366004612360565b6105b0565b61026461062b565b60405161022b91906123ed565b61024761027f366004612497565b6106a6565b61028c610723565b60405161022b91906124de565b6102ac6102a7366004612529565b61079a565b60405161022b949392919061256d565b6102cf6102ca3660046125ac565b6108dd565b60405161022b91906125ea565b6102ef6102ea3660046125ac565b6108f1565b60405190815260200161022b565b61024761030b366004612360565b61096d565b6103186109b7565b60405161022b91906125fc565b6102ef610333366004612653565b610a2d565b61034b610346366004612653565b610a6a565b60405161022b9190612685565b61021f6103663660046125ac565b610e5d565b610247610379366004612772565b610ed9565b61039161038c3660046125ac565b610f59565b60405161022b9291906127dc565b61021f6103ad3660046125ac565b610fe6565b6102476103c03660046125ac565b611023565b6102ef6103d3366004612653565b61106c565b6103e06110a9565b60405161022b919061280b565b6102476103fb3660046125ac565b611120565b61024761040e3660046128c1565b61118b565b61021f6104213660046125ac565b6111d6565b600a546040516001600160a01b03909116815260200161022b565b61024761044f366004612772565b61128a565b610247610462366004612497565b6112d7565b610247610475366004612913565b611322565b6102476104883660046129b8565b61157d565b6102cf61049b3660046125ac565b6115d2565b6102cf6104ae366004612529565b611651565b6102ef6104c13660046125ac565b6116f4565b61034b6104d4366004612653565b611731565b6102476104e73660046129d3565b611aa7565b6102cf6104fa3660046125ac565b611af3565b610318611b31565b5f8061051560083383611c05565b6105228787878787611c69565b979650505050505050565b600361053b60083383611c05565b60405163641bf91960e11b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063c837f2329061057c905f908a908a908a908a908a90600401612aed565b5f6040518083038186803b158015610592575f5ffd5b505af41580156105a4573d5f5f3e3d5ffd5b50505050505050505050565b60026105be60083383611c05565b604051630d3d280360e21b815273__$5a92de94e3c11b56e4ee4f365f96da5afc$__906334f4a00c906105fa9060089087908790600401612b33565b5f6040518083038186803b158015610610575f5ffd5b505af4158015610622573d5f5f3e3d5ffd5b50505050505050565b60405163b64ee37d60e01b81525f600482015260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063b64ee37d906024015f60405180830381865af415801561067a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a19190810190612bc7565b905090565b60036106b460083383611c05565b60405163f2bffc1360e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063f2bffc13906106f1905f90889088908890600401612d05565b5f6040518083038186803b158015610707575f5ffd5b505af4158015610719573d5f5f3e3d5ffd5b5050505050505050565b604051630cf1f1fb60e21b81526008600482015260609073__$5a92de94e3c11b56e4ee4f365f96da5afc$__906333c7c7ec906024015f60405180830381865af4158015610773573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a19190810190612d2f565b8051602081830181018051600b825292820191909301209152805481906107c090612dc8565b80601f01602080910402602001604051908101604052809291908181526020018280546107ec90612dc8565b80156108375780601f1061080e57610100808354040283529160200191610837565b820191905f5260205f20905b81548152906001019060200180831161081a57829003601f168201915b50505050509080600101805461084c90612dc8565b80601f016020809104026020016040519081016040528092919081815260200182805461087890612dc8565b80156108c35780601f1061089a576101008083540402835291602001916108c3565b820191905f5260205f20905b8154815290600101906020018083116108a657829003601f168201915b50505050600383015460049093015491929160ff16905084565b60606108ea5f8484611d98565b9392505050565b604051632519093f60e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__90632519093f9061092e90849087908790600401612e00565b602060405180830381865af4158015610949573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ea9190612e22565b600261097b60083383611c05565b60405163310adb3d60e11b815273__$5a92de94e3c11b56e4ee4f365f96da5afc$__90636215b67a906105fa9060089087908790600401612b33565b60405163624ff69960e11b81525f600482015260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063c49fed32906024015f60405180830381865af4158015610a06573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a19190810190612e39565b604051631a770e2b60e11b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__906334ee1c569061092e90849087908790600401612edd565b610a9d6040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b5f600c8484604051610ab0929190612ef6565b90815260200160405180910390208054610ac990612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610af590612dc8565b8015610b405780601f10610b1757610100808354040283529160200191610b40565b820191905f5260205f20905b815481529060010190602001808311610b2357829003601f168201915b5050505050905080515f03610b6857604051634fde869d60e01b815260040160405180910390fd5b600b81604051610b789190612f1c565b90815260200160405180910390206040518060a00160405290815f82018054610ba090612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610bcc90612dc8565b8015610c175780601f10610bee57610100808354040283529160200191610c17565b820191905f5260205f20905b815481529060010190602001808311610bfa57829003601f168201915b50505050508152602001600182018054610c3090612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5c90612dc8565b8015610ca75780601f10610c7e57610100808354040283529160200191610ca7565b820191905f5260205f20905b815481529060010190602001808311610c8a57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015610e35578382905f5260205f2090600302016040518060600160405290815f82018054610d0490612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3090612dc8565b8015610d7b5780601f10610d5257610100808354040283529160200191610d7b565b820191905f5260205f20905b815481529060010190602001808311610d5e57829003601f168201915b50505050508152602001600182018054610d9490612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610dc090612dc8565b8015610e0b5780601f10610de257610100808354040283529160200191610e0b565b820191905f5260205f20905b815481529060010190602001808311610dee57829003601f168201915b50505091835250506002919091015460ff1615156020918201529082526001929092019101610cd4565b505050908252506003820154602082015260049091015460ff16604090910152949350505050565b60405163dea25e9f60e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063dea25e9f90610e9a90849087908790600401612e00565b602060405180830381865af4158015610eb5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ea9190612f27565b6003610ee760083383611c05565b604051639a7044d760e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__90639a7044d790610f26905f908990899089908990600401612f42565b5f6040518083038186803b158015610f3c575f5ffd5b505af4158015610f4e573d5f5f3e3d5ffd5b505050505050505050565b604051639e72831160e01b81525f9060609073__$d451bf733f00e36b69819fe95d7d8eab5b$__90639e72831190610f9990859088908890600401612e00565b5f60405180830381865af4158015610fb3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610fda9190810190612f7a565b915091505b9250929050565b604051635d0aad2360e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__90635d0aad2390610e9a90849087908790600401612e00565b600361103160083383611c05565b604051638b7b2e1760e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__90638b7b2e17906105fa905f9087908790600401612fc7565b6040516332a1c4bb60e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__906332a1c4bb9061092e90849087908790600401612edd565b604051639b2642df60e01b81526008600482015260609073__$5a92de94e3c11b56e4ee4f365f96da5afc$__90639b2642df906024015f60405180830381865af41580156110f9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a191908101906130cf565b6040516340efd12960e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__906340efd1299061115b905f9086908690600401612e00565b5f6040518083038186803b158015611171575f5ffd5b505af4158015611183573d5f5f3e3d5ffd5b505050505050565b600361119960083383611c05565b60405163720e983960e11b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063e41d3072906106f1905f9088908890889060040161322a565b5f806111e460083383611c05565b5f5b8381101561127f573685858381811061120157611201613256565b9050602002810190611213919061326a565b90506112756112228280613288565b61122f60208501856132ca565b61123c6040870187613288565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611c6992505050565b50506001016111e6565b506001949350505050565b600361129860083383611c05565b6040516338fe9c7b60e21b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063e3fa71ec90610f26905f908990899089908990600401612f42565b60036112e560083383611c05565b604051631b33bcaf60e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__90631b33bcaf906106f1905f90889088908890600401612d05565b600161133060083383611c05565b600b8888604051611342929190612ef6565b908152604051908190036020019020805461135c90612dc8565b15905061137c57604051630268c2d760e51b815260040160405180910390fd5b5f805b838110156113d75784848281811061139957611399613256565b90506020028101906113ab919061326a565b6113bc90606081019060400161330f565b156113cf57816113cb8161333e565b9250505b60010161137f565b5060038160ff1611156113fc57604051626f3b7d60e31b815260040160405180910390fd5b5f600b8a8a60405161140f929190612ef6565b90815260405190819003602001902090508061142c8a8c836133a7565b506001810161143c888a836133a7565b50600381018690555f5b848110156114a1578160020186868381811061146457611464613256565b9050602002810190611476919061326a565b81546001810183555f92835260209092209091600302016114978282613472565b5050600101611446565b5060048101805460ff191660ff8416179055600d80546001810182555f919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5016114ef8a8c836133a7565b508989600c8a8a604051611504929190612ef6565b9081526020016040518091039020918261151f9291906133a7565b508989604051611530929190612ef6565b60405180910390207f4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de898960405161156992919061357c565b60405180910390a250505050505050505050565b600261158b60083383611c05565b60405163f76888e760e01b8152600860048201526001600160a01b038316602482015273__$5a92de94e3c11b56e4ee4f365f96da5afc$__9063f76888e79060440161115b565b604051633a555c0b60e21b815260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063e955702c90611610905f9087908790600401612e00565b5f60405180830381865af415801561162a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108ea919081019061358f565b8051602081830181018051600c825292820191909301209152805461167590612dc8565b80601f01602080910402602001604051908101604052809291908181526020018280546116a190612dc8565b80156116ec5780601f106116c3576101008083540402835291602001916116ec565b820191905f5260205f20905b8154815290600101906020018083116116cf57829003601f168201915b505050505081565b604051631c74ab5760e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__90631c74ab579061092e90849087908790600401612e00565b6117646040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b600b8383604051611776929190612ef6565b908152604051908190036020019020805461179090612dc8565b90505f036117b157604051634fde869d60e01b815260040160405180910390fd5b600b83836040516117c3929190612ef6565b90815260200160405180910390206040518060a00160405290815f820180546117eb90612dc8565b80601f016020809104026020016040519081016040528092919081815260200182805461181790612dc8565b80156118625780601f1061183957610100808354040283529160200191611862565b820191905f5260205f20905b81548152906001019060200180831161184557829003601f168201915b5050505050815260200160018201805461187b90612dc8565b80601f01602080910402602001604051908101604052809291908181526020018280546118a790612dc8565b80156118f25780601f106118c9576101008083540402835291602001916118f2565b820191905f5260205f20905b8154815290600101906020018083116118d557829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015611a80578382905f5260205f2090600302016040518060600160405290815f8201805461194f90612dc8565b80601f016020809104026020016040519081016040528092919081815260200182805461197b90612dc8565b80156119c65780601f1061199d576101008083540402835291602001916119c6565b820191905f5260205f20905b8154815290600101906020018083116119a957829003601f168201915b505050505081526020016001820180546119df90612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054611a0b90612dc8565b8015611a565780601f10611a2d57610100808354040283529160200191611a56565b820191905f5260205f20905b815481529060010190602001808311611a3957829003601f168201915b50505091835250506002919091015460ff161515602091820152908252600192909201910161191f565b505050908252506003820154602082015260049091015460ff166040909101529392505050565b6002611ab560083383611c05565b604051631fbbcd3760e31b815273__$5a92de94e3c11b56e4ee4f365f96da5afc$__9063fdde69b8906106f1906008908890889088906004016135c0565b604051635551774560e01b815260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__90635551774590611610905f9087908790600401612e00565b6060600d805480602002602001604051908101604052809291908181526020015f905b82821015611bfc578382905f5260205f20018054611b7190612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054611b9d90612dc8565b8015611be85780601f10611bbf57610100808354040283529160200191611be8565b820191905f5260205f20905b815481529060010190602001808311611bcb57829003601f168201915b505050505081526020019060010190611b54565b50505050905090565b6001600160a01b0382165f9081526020849052604090205460ff16611c3d5760405163c597feeb60e01b815260040160405180910390fd5b611c48838383611dad565b611c64576040516282b42960e81b815260040160405180910390fd5b505050565b5f600b8686604051611c7c929190612ef6565b9081526040519081900360200190208054611c9690612dc8565b90505f03611cb757604051634fde869d60e01b815260040160405180910390fd5b600b8686604051611cc9929190612ef6565b9081526040519081900360200190206004015460ff168314611cfe57604051631f73a41360e01b815260040160405180910390fd5b5f600b8787604051611d11929190612ef6565b90815260200160405180910390206003015490505f8351905060208401855f8114611d535760018114611d5c5760028114611d675760038114611d7757611d88565b838383a1611d88565b8735848484a2611d88565b60208801358835858585a3611d88565b604088013560208901358935868686a45b5060019998505050505050505050565b6060611da5848484611e48565b949350505050565b6001600160a01b0382165f908152602084905260408120600101815b8154811015611e3d57836003811115611de457611de46123c5565b828281548110611df657611df6613256565b905f5260205f2090602091828204019190069054906101000a900460ff166003811115611e2557611e256123c5565b03611e35576001925050506108ea565b600101611dc9565b505f95945050505050565b604080515f808252602082019092526060915b83811015611ec9575f611e9187878785818110611e7a57611e7a613256565b9050602002810190611e8c9190613288565b611ef6565b905082611e9d82611ffe565b604051602001611eae929190613625565b60408051601f19818403018152919052925050600101611e5b565b50604051611edd9082905f90602001613639565b60408051808303601f1901815291905295945050505050565b5f5f846002018484604051611f0c929190612ef6565b9081526040805160209281900383019020545f8181526004890190935291205490915060ff1615611f3e5790506108ea565b6001850180549081905f611f518361365d565b919050555080866002018686604051611f6b929190612ef6565b90815260405160209181900382019020919091556003870180546001810182555f91825291902001611f9e8587836133a7565b505f81815260048701602052604090819020805460ff191660011790555181907f4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e77890611fed908890889061357c565b60405180910390a291506108ea9050565b6060815f03612029576040515f60208201526021016040516020818303038152906040529050919050565b6060825b6080811061206b578181607f1660801760f81b604051602001612051929190613639565b60408051601f19818403018152919052915060071c61202d565b8181607f1660f81b604051602001612084929190613639565b60408051601f19818403018152919052949350505050565b5f5f83601f8401126120ac575f5ffd5b5081356001600160401b038111156120c2575f5ffd5b602083019150836020828501011115610fdf575f5ffd5b5f5f83601f8401126120e9575f5ffd5b5081356001600160401b038111156120ff575f5ffd5b6020830191508360208260051b8501011115610fdf575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b038111828210171561214f5761214f612119565b60405290565b604080519081016001600160401b038111828210171561214f5761214f612119565b604051601f8201601f191681016001600160401b038111828210171561219f5761219f612119565b604052919050565b5f6001600160401b038211156121bf576121bf612119565b50601f01601f191660200190565b5f6121df6121da846121a7565b612177565b90508281528383830111156121f2575f5ffd5b828260208301375f602084830101529392505050565b5f5f5f5f5f6060868803121561221c575f5ffd5b85356001600160401b03811115612231575f5ffd5b61223d8882890161209c565b90965094505060208601356001600160401b0381111561225b575f5ffd5b612267888289016120d9565b90945092505060408601356001600160401b03811115612285575f5ffd5b8601601f81018813612295575f5ffd5b6122a4888235602084016121cd565b9150509295509295909350565b600681106122bd575f5ffd5b50565b5f5f5f5f5f606086880312156122d4575f5ffd5b85356001600160401b038111156122e9575f5ffd5b6122f5888289016120d9565b9096509450506020860135612309816122b1565b925060408601356001600160401b03811115612323575f5ffd5b61232f8882890161209c565b969995985093965092949392505050565b6001600160a01b03811681146122bd575f5ffd5b600481106122bd575f5ffd5b5f5f60408385031215612371575f5ffd5b823561237c81612340565b9150602083013561238c81612354565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b5f52602160045260245ffd5b600681106123e9576123e96123c5565b9052565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561248b57603f1987860301845281518051608087526124396080880182612397565b9050602082015161244d60208901826123d9565b50604082015187820360408901526124658282612397565b606093840151151598909301979097525094506020938401939190910190600101612413565b50929695505050505050565b5f5f5f604084860312156124a9575f5ffd5b83356001600160401b038111156124be575f5ffd5b6124ca868287016120d9565b909790965060209590950135949350505050565b602080825282518282018190525f918401906040840190835b8181101561251e5783516001600160a01b03168352602093840193909201916001016124f7565b509095945050505050565b5f60208284031215612539575f5ffd5b81356001600160401b0381111561254e575f5ffd5b8201601f8101841361255e575f5ffd5b611da5848235602084016121cd565b608081525f61257f6080830187612397565b82810360208401526125918187612397565b91505083604083015260ff8316606083015295945050505050565b5f5f602083850312156125bd575f5ffd5b82356001600160401b038111156125d2575f5ffd5b6125de858286016120d9565b90969095509350505050565b602081525f6108ea6020830184612397565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561248b57603f1987860301845261263e858351612397565b94506020938401939190910190600101612622565b5f5f60208385031215612664575f5ffd5b82356001600160401b03811115612679575f5ffd5b6125de8582860161209c565b602081525f825160a060208401526126a060c0840182612397565b90506020840151601f198483030160408501526126bd8282612397565b6040860151858203601f190160608701528051808352919350602090810192508084019190600582901b8501015f5b8281101561275157601f1986830301845284518051606084526127126060850182612397565b90506020820151848203602086015261272b8282612397565b6040938401511515959093019490945250602095860195949094019391506001016126ec565b50606088015160808801526080880151945061052260a088018660ff169052565b5f5f5f5f60408587031215612785575f5ffd5b84356001600160401b0381111561279a575f5ffd5b6127a6878288016120d9565b90955093505060208501356001600160401b038111156127c4575f5ffd5b6127d08782880161209c565b95989497509550505050565b6127e681846123d9565b604060208201525f611da56040830184612397565b600481106123e9576123e96123c5565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561248b57868503603f19018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101905f9060608801905b8083101561289c576128858285516127fb565b602082019150602084019350600183019250612872565b50965050506020938401939190910190600101612831565b80151581146122bd575f5ffd5b5f5f5f604084860312156128d3575f5ffd5b83356001600160401b038111156128e8575f5ffd5b6128f4868287016120d9565b9094509250506020840135612908816128b4565b809150509250925092565b5f5f5f5f5f5f5f6080888a031215612929575f5ffd5b87356001600160401b0381111561293e575f5ffd5b61294a8a828b0161209c565b90985096505060208801356001600160401b03811115612968575f5ffd5b6129748a828b0161209c565b9096509450506040880135925060608801356001600160401b03811115612999575f5ffd5b6129a58a828b016120d9565b989b979a50959850939692959293505050565b5f602082840312156129c8575f5ffd5b81356108ea81612340565b5f5f5f604084860312156129e5575f5ffd5b83356129f081612340565b925060208401356001600160401b03811115612a0a575f5ffd5b612a16868287016120d9565b9497909650939450505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f5f8335601e19843603018112612a60575f5ffd5b83016020810192503590506001600160401b03811115612a7e575f5ffd5b803603821315610fdf575f5ffd5b5f8383855260208501945060208460051b820101835f5b86811015612ae157838303601f19018852612abe8287612a4b565b612ac9858284612a23565b60209a8b019a90955093909301925050600101612aa3565b50909695505050505050565b868152608060208201525f612b06608083018789612a8c565b612b1360408401876123d9565b8281036060840152612b26818587612a23565b9998505050505050505050565b8381526001600160a01b038316602082015260608101611da560408301846127fb565b5f6001600160401b03821115612b6e57612b6e612119565b5060051b60200190565b5f82601f830112612b87575f5ffd5b8151602083015f612b9a6121da846121a7565b9050828152858383011115612bad575f5ffd5b8282602083015e5f92810160200192909252509392505050565b5f60208284031215612bd7575f5ffd5b81516001600160401b03811115612bec575f5ffd5b8201601f81018413612bfc575f5ffd5b8051612c0a6121da82612b56565b8082825260208201915060208360051b850101925086831115612c2b575f5ffd5b602084015b83811015612cfa5780516001600160401b03811115612c4d575f5ffd5b85016080818a03601f19011215612c62575f5ffd5b612c6a61212d565b60208201516001600160401b03811115612c82575f5ffd5b612c918b602083860101612b78565b8252506040820151612ca2816122b1565b602082015260608201516001600160401b03811115612cbf575f5ffd5b612cce8b602083860101612b78565b60408301525060808201519150612ce4826128b4565b6060810191909152835260209283019201612c30565b509695505050505050565b848152606060208201525f612d1e606083018587612a8c565b905082604083015295945050505050565b5f60208284031215612d3f575f5ffd5b81516001600160401b03811115612d54575f5ffd5b8201601f81018413612d64575f5ffd5b8051612d726121da82612b56565b8082825260208201915060208360051b850101925086831115612d93575f5ffd5b6020840193505b82841015612dbe578351612dad81612340565b825260209384019390910190612d9a565b9695505050505050565b600181811c90821680612ddc57607f821691505b602082108103612dfa57634e487b7160e01b5f52602260045260245ffd5b50919050565b838152604060208201525f612e19604083018486612a8c565b95945050505050565b5f60208284031215612e32575f5ffd5b5051919050565b5f60208284031215612e49575f5ffd5b81516001600160401b03811115612e5e575f5ffd5b8201601f81018413612e6e575f5ffd5b8051612e7c6121da82612b56565b8082825260208201915060208360051b850101925086831115612e9d575f5ffd5b602084015b83811015612cfa5780516001600160401b03811115612ebf575f5ffd5b612ece89602083890101612b78565b84525060209283019201612ea2565b838152604060208201525f612e19604083018486612a23565b818382375f9101908152919050565b5f81518060208401855e5f93019283525090919050565b5f6108ea8284612f05565b5f60208284031215612f37575f5ffd5b81516108ea816128b4565b858152606060208201525f612f5b606083018688612a8c565b8281036040840152612f6e818587612a23565b98975050505050505050565b5f5f60408385031215612f8b575f5ffd5b8251612f96816122b1565b60208401519092506001600160401b03811115612fb1575f5ffd5b612fbd85828601612b78565b9150509250929050565b83815260406020820181905281018290525f6060600584901b830181019083018583605e1936839003015b878210156130c157868503605f190184528235818112613010575f5ffd5b8901803536829003601e19018112613026575f5ffd5b81016020810190356001600160401b03811115613041575f5ffd5b8060051b3603821315613052575f5ffd5b60608852613064606089018284612a8c565b9150506020820135613075816122b1565b61308260208901826123d9565b506130906040830183612a4b565b925087820360408901526130a5828483612a23565b9750505050602083019250602084019350600182019150612ff2565b509298975050505050505050565b5f602082840312156130df575f5ffd5b81516001600160401b038111156130f4575f5ffd5b8201601f81018413613104575f5ffd5b80516131126121da82612b56565b8082825260208201915060208360051b850101925086831115613133575f5ffd5b602084015b83811015612cfa5780516001600160401b03811115613155575f5ffd5b85016040818a03601f1901121561316a575f5ffd5b613172612155565b602082015161318081612340565b815260408201516001600160401b0381111561319a575f5ffd5b60208184010192505089601f8301126131b1575f5ffd5b81516131bf6121da82612b56565b8082825260208201915060208360051b86010192508c8311156131e0575f5ffd5b6020850194505b8285101561320b5784516131fa81612354565b8252602094850194909101906131e7565b8060208501525050508085525050602083019250602081019050613138565b848152606060208201525f613243606083018587612a8c565b9050821515604083015295945050505050565b634e487b7160e01b5f52603260045260245ffd5b5f8235605e1983360301811261327e575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261329d575f5ffd5b8301803591506001600160401b038211156132b6575f5ffd5b602001915036819003821315610fdf575f5ffd5b5f5f8335601e198436030181126132df575f5ffd5b8301803591506001600160401b038211156132f8575f5ffd5b6020019150600581901b3603821315610fdf575f5ffd5b5f6020828403121561331f575f5ffd5b81356108ea816128b4565b634e487b7160e01b5f52601160045260245ffd5b5f60ff821660ff81036133535761335361332a565b60010192915050565b601f821115611c6457805f5260205f20601f840160051c810160208510156133815750805b601f840160051c820191505b818110156133a0575f815560010161338d565b5050505050565b6001600160401b038311156133be576133be612119565b6133d2836133cc8354612dc8565b8361335c565b5f601f841160018114613403575f85156133ec5750838201355b5f19600387901b1c1916600186901b1783556133a0565b5f83815260208120601f198716915b828110156134325786850135825560209485019460019092019101613412565b508682101561344e575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f813561346c816128b4565b92915050565b61347c8283613288565b6001600160401b0381111561349357613493612119565b6134a7816134a18554612dc8565b8561335c565b5f601f8211600181146134d8575f83156134c15750838201355b5f19600385901b1c1916600184901b17855561352f565b5f85815260208120601f198516915b8281101561350757868501358255602094850194600190920191016134e7565b5084821015613523575f1960f88660031b161c19848701351681555b505060018360011b0185555b505050506135406020830183613288565b61354e8183600186016133a7565b505061357861355f60408401613460565b6002830160ff1981541660ff8315151681178255505050565b5050565b602081525f611da5602083018486612a23565b5f6020828403121561359f575f5ffd5b81516001600160401b038111156135b4575f5ffd5b611da584828501612b78565b8481526001600160a01b038416602082015260606040820181905281018290525f8360808301825b858110156136195782356135fb81612354565b61360583826127fb565b5060209283019291909101906001016135e8565b50979650505050505050565b5f611da56136338386612f05565b84612f05565b5f6136448285612f05565b6001600160f81b03199390931683525050600101919050565b5f6001820161366e5761366e61332a565b506001019056fea2646970667358221220f8b15555cb7cdff5a8b4a6abc5fd5c705beaed31fd5f7dedb88877476213dfb764736f6c634300081c0033",
}

// EventEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use EventEmitterMetaData.ABI instead.
var EventEmitterABI = EventEmitterMetaData.ABI

// EventEmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventEmitterMetaData.Bin instead.
var EventEmitterBin = EventEmitterMetaData.Bin

// DeployEventEmitter deploys a new Ethereum contract, binding an instance of EventEmitter to it.
func DeployEventEmitter(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, validator common.Address) (common.Address, *types.Transaction, *EventEmitter, error) {
	parsed, err := EventEmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	roleControlAddr, _, _, _ := DeployRoleControl(auth, backend)
	EventEmitterBin = strings.ReplaceAll(EventEmitterBin, "__$5a92de94e3c11b56e4ee4f365f96da5afc$__", roleControlAddr.String()[2:])

	etherDatabaseLibAddr, _, _, _ := DeployEtherDatabaseLib(auth, backend)
	EventEmitterBin = strings.ReplaceAll(EventEmitterBin, "__$d451bf733f00e36b69819fe95d7d8eab5b$__", etherDatabaseLibAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventEmitterBin), backend, _owner, validator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventEmitter{EventEmitterCaller: EventEmitterCaller{contract: contract}, EventEmitterTransactor: EventEmitterTransactor{contract: contract}, EventEmitterFilterer: EventEmitterFilterer{contract: contract}}, nil
}

// EventEmitter is an auto generated Go binding around an Ethereum contract.
type EventEmitter struct {
	EventEmitterCaller     // Read-only binding to the contract
	EventEmitterTransactor // Write-only binding to the contract
	EventEmitterFilterer   // Log filterer for contract events
}

// EventEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventEmitterSession struct {
	Contract     *EventEmitter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventEmitterCallerSession struct {
	Contract *EventEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventEmitterTransactorSession struct {
	Contract     *EventEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventEmitterRaw struct {
	Contract *EventEmitter // Generic contract binding to access the raw methods on
}

// EventEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventEmitterCallerRaw struct {
	Contract *EventEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// EventEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventEmitterTransactorRaw struct {
	Contract *EventEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventEmitter creates a new instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitter(address common.Address, backend bind.ContractBackend) (*EventEmitter, error) {
	contract, err := bindEventEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventEmitter{EventEmitterCaller: EventEmitterCaller{contract: contract}, EventEmitterTransactor: EventEmitterTransactor{contract: contract}, EventEmitterFilterer: EventEmitterFilterer{contract: contract}}, nil
}

// NewEventEmitterCaller creates a new read-only instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitterCaller(address common.Address, caller bind.ContractCaller) (*EventEmitterCaller, error) {
	contract, err := bindEventEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventEmitterCaller{contract: contract}, nil
}

// NewEventEmitterTransactor creates a new write-only instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*EventEmitterTransactor, error) {
	contract, err := bindEventEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventEmitterTransactor{contract: contract}, nil
}

// NewEventEmitterFilterer creates a new log filterer instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*EventEmitterFilterer, error) {
	contract, err := bindEventEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFilterer{contract: contract}, nil
}

// bindEventEmitter binds a generic wrapper to an already deployed contract.
func bindEventEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventEmitter *EventEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventEmitter.Contract.EventEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventEmitter *EventEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitter.Contract.EventEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventEmitter *EventEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventEmitter.Contract.EventEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventEmitter *EventEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventEmitter *EventEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventEmitter *EventEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventEmitter.Contract.contract.Transact(opts, method, params...)
}

// EventSchemas is a free data retrieval call binding the contract method 0x2f2545d2.
//
// Solidity: function eventSchemas(string ) view returns(string name, string id, bytes32 eventTopic, uint8 numIndexedArgs)
func (_EventEmitter *EventEmitterCaller) EventSchemas(opts *bind.CallOpts, arg0 string) (struct {
	Name           string
	Id             string
	EventTopic     [32]byte
	NumIndexedArgs uint8
}, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "eventSchemas", arg0)

	outstruct := new(struct {
		Name           string
		Id             string
		EventTopic     [32]byte
		NumIndexedArgs uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Id = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.EventTopic = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.NumIndexedArgs = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// EventSchemas is a free data retrieval call binding the contract method 0x2f2545d2.
//
// Solidity: function eventSchemas(string ) view returns(string name, string id, bytes32 eventTopic, uint8 numIndexedArgs)
func (_EventEmitter *EventEmitterSession) EventSchemas(arg0 string) (struct {
	Name           string
	Id             string
	EventTopic     [32]byte
	NumIndexedArgs uint8
}, error) {
	return _EventEmitter.Contract.EventSchemas(&_EventEmitter.CallOpts, arg0)
}

// EventSchemas is a free data retrieval call binding the contract method 0x2f2545d2.
//
// Solidity: function eventSchemas(string ) view returns(string name, string id, bytes32 eventTopic, uint8 numIndexedArgs)
func (_EventEmitter *EventEmitterCallerSession) EventSchemas(arg0 string) (struct {
	Name           string
	Id             string
	EventTopic     [32]byte
	NumIndexedArgs uint8
}, error) {
	return _EventEmitter.Contract.EventSchemas(&_EventEmitter.CallOpts, arg0)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_EventEmitter *EventEmitterCaller) GetAllIdentities(opts *bind.CallOpts) ([]RoleControlIdentityView, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getAllIdentities")

	if err != nil {
		return *new([]RoleControlIdentityView), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleControlIdentityView)).(*[]RoleControlIdentityView)

	return out0, err

}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_EventEmitter *EventEmitterSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _EventEmitter.Contract.GetAllIdentities(&_EventEmitter.CallOpts)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_EventEmitter *EventEmitterCallerSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _EventEmitter.Contract.GetAllIdentities(&_EventEmitter.CallOpts)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EventEmitter *EventEmitterCaller) GetAllSegments(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getAllSegments")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EventEmitter *EventEmitterSession) GetAllSegments() ([]string, error) {
	return _EventEmitter.Contract.GetAllSegments(&_EventEmitter.CallOpts)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EventEmitter *EventEmitterCallerSession) GetAllSegments() ([]string, error) {
	return _EventEmitter.Contract.GetAllSegments(&_EventEmitter.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_EventEmitter *EventEmitterCaller) GetAllWallets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getAllWallets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_EventEmitter *EventEmitterSession) GetAllWallets() ([]common.Address, error) {
	return _EventEmitter.Contract.GetAllWallets(&_EventEmitter.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_EventEmitter *EventEmitterCallerSession) GetAllWallets() ([]common.Address, error) {
	return _EventEmitter.Contract.GetAllWallets(&_EventEmitter.CallOpts)
}

// GetBool is a free data retrieval call binding the contract method 0x4e0a38dc.
//
// Solidity: function getBool(string[] segments) view returns(bool)
func (_EventEmitter *EventEmitterCaller) GetBool(opts *bind.CallOpts, segments []string) (bool, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getBool", segments)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x4e0a38dc.
//
// Solidity: function getBool(string[] segments) view returns(bool)
func (_EventEmitter *EventEmitterSession) GetBool(segments []string) (bool, error) {
	return _EventEmitter.Contract.GetBool(&_EventEmitter.CallOpts, segments)
}

// GetBool is a free data retrieval call binding the contract method 0x4e0a38dc.
//
// Solidity: function getBool(string[] segments) view returns(bool)
func (_EventEmitter *EventEmitterCallerSession) GetBool(segments []string) (bool, error) {
	return _EventEmitter.Contract.GetBool(&_EventEmitter.CallOpts, segments)
}

// GetBytes is a free data retrieval call binding the contract method 0xab841283.
//
// Solidity: function getBytes(string[] segments) view returns(bytes)
func (_EventEmitter *EventEmitterCaller) GetBytes(opts *bind.CallOpts, segments []string) ([]byte, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getBytes", segments)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0xab841283.
//
// Solidity: function getBytes(string[] segments) view returns(bytes)
func (_EventEmitter *EventEmitterSession) GetBytes(segments []string) ([]byte, error) {
	return _EventEmitter.Contract.GetBytes(&_EventEmitter.CallOpts, segments)
}

// GetBytes is a free data retrieval call binding the contract method 0xab841283.
//
// Solidity: function getBytes(string[] segments) view returns(bytes)
func (_EventEmitter *EventEmitterCallerSession) GetBytes(segments []string) ([]byte, error) {
	return _EventEmitter.Contract.GetBytes(&_EventEmitter.CallOpts, segments)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EventEmitter *EventEmitterCaller) GetEntries(opts *bind.CallOpts) ([]EtherDatabaseLibNode, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]EtherDatabaseLibNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]EtherDatabaseLibNode)).(*[]EtherDatabaseLibNode)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EventEmitter *EventEmitterSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EventEmitter.Contract.GetEntries(&_EventEmitter.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EventEmitter *EventEmitterCallerSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EventEmitter.Contract.GetEntries(&_EventEmitter.CallOpts)
}

// GetEventSchemaFromId is a free data retrieval call binding the contract method 0x49ab123d.
//
// Solidity: function getEventSchemaFromId(string id) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitter *EventEmitterCaller) GetEventSchemaFromId(opts *bind.CallOpts, id string) (EventSchema, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getEventSchemaFromId", id)

	if err != nil {
		return *new(EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new(EventSchema)).(*EventSchema)

	return out0, err

}

// GetEventSchemaFromId is a free data retrieval call binding the contract method 0x49ab123d.
//
// Solidity: function getEventSchemaFromId(string id) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitter *EventEmitterSession) GetEventSchemaFromId(id string) (EventSchema, error) {
	return _EventEmitter.Contract.GetEventSchemaFromId(&_EventEmitter.CallOpts, id)
}

// GetEventSchemaFromId is a free data retrieval call binding the contract method 0x49ab123d.
//
// Solidity: function getEventSchemaFromId(string id) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitter *EventEmitterCallerSession) GetEventSchemaFromId(id string) (EventSchema, error) {
	return _EventEmitter.Contract.GetEventSchemaFromId(&_EventEmitter.CallOpts, id)
}

// GetEventSchemaFromName is a free data retrieval call binding the contract method 0xd50d542e.
//
// Solidity: function getEventSchemaFromName(string name) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitter *EventEmitterCaller) GetEventSchemaFromName(opts *bind.CallOpts, name string) (EventSchema, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getEventSchemaFromName", name)

	if err != nil {
		return *new(EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new(EventSchema)).(*EventSchema)

	return out0, err

}

// GetEventSchemaFromName is a free data retrieval call binding the contract method 0xd50d542e.
//
// Solidity: function getEventSchemaFromName(string name) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitter *EventEmitterSession) GetEventSchemaFromName(name string) (EventSchema, error) {
	return _EventEmitter.Contract.GetEventSchemaFromName(&_EventEmitter.CallOpts, name)
}

// GetEventSchemaFromName is a free data retrieval call binding the contract method 0xd50d542e.
//
// Solidity: function getEventSchemaFromName(string name) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitter *EventEmitterCallerSession) GetEventSchemaFromName(name string) (EventSchema, error) {
	return _EventEmitter.Contract.GetEventSchemaFromName(&_EventEmitter.CallOpts, name)
}

// GetInt is a free data retrieval call binding the contract method 0xc8740be1.
//
// Solidity: function getInt(string[] segments) view returns(int256)
func (_EventEmitter *EventEmitterCaller) GetInt(opts *bind.CallOpts, segments []string) (*big.Int, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getInt", segments)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0xc8740be1.
//
// Solidity: function getInt(string[] segments) view returns(int256)
func (_EventEmitter *EventEmitterSession) GetInt(segments []string) (*big.Int, error) {
	return _EventEmitter.Contract.GetInt(&_EventEmitter.CallOpts, segments)
}

// GetInt is a free data retrieval call binding the contract method 0xc8740be1.
//
// Solidity: function getInt(string[] segments) view returns(int256)
func (_EventEmitter *EventEmitterCallerSession) GetInt(segments []string) (*big.Int, error) {
	return _EventEmitter.Contract.GetInt(&_EventEmitter.CallOpts, segments)
}

// GetRegisteredEventNames is a free data retrieval call binding the contract method 0xf30826f0.
//
// Solidity: function getRegisteredEventNames() view returns(string[])
func (_EventEmitter *EventEmitterCaller) GetRegisteredEventNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getRegisteredEventNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetRegisteredEventNames is a free data retrieval call binding the contract method 0xf30826f0.
//
// Solidity: function getRegisteredEventNames() view returns(string[])
func (_EventEmitter *EventEmitterSession) GetRegisteredEventNames() ([]string, error) {
	return _EventEmitter.Contract.GetRegisteredEventNames(&_EventEmitter.CallOpts)
}

// GetRegisteredEventNames is a free data retrieval call binding the contract method 0xf30826f0.
//
// Solidity: function getRegisteredEventNames() view returns(string[])
func (_EventEmitter *EventEmitterCallerSession) GetRegisteredEventNames() ([]string, error) {
	return _EventEmitter.Contract.GetRegisteredEventNames(&_EventEmitter.CallOpts)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EventEmitter *EventEmitterCaller) GetSegmentId(opts *bind.CallOpts, segment string) (*big.Int, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getSegmentId", segment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EventEmitter *EventEmitterSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EventEmitter.Contract.GetSegmentId(&_EventEmitter.CallOpts, segment)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EventEmitter *EventEmitterCallerSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EventEmitter.Contract.GetSegmentId(&_EventEmitter.CallOpts, segment)
}

// GetString is a free data retrieval call binding the contract method 0xe6149cd0.
//
// Solidity: function getString(string[] segments) view returns(string)
func (_EventEmitter *EventEmitterCaller) GetString(opts *bind.CallOpts, segments []string) (string, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getString", segments)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0xe6149cd0.
//
// Solidity: function getString(string[] segments) view returns(string)
func (_EventEmitter *EventEmitterSession) GetString(segments []string) (string, error) {
	return _EventEmitter.Contract.GetString(&_EventEmitter.CallOpts, segments)
}

// GetString is a free data retrieval call binding the contract method 0xe6149cd0.
//
// Solidity: function getString(string[] segments) view returns(string)
func (_EventEmitter *EventEmitterCallerSession) GetString(segments []string) (string, error) {
	return _EventEmitter.Contract.GetString(&_EventEmitter.CallOpts, segments)
}

// GetUint is a free data retrieval call binding the contract method 0x3aa311dd.
//
// Solidity: function getUint(string[] segments) view returns(uint256)
func (_EventEmitter *EventEmitterCaller) GetUint(opts *bind.CallOpts, segments []string) (*big.Int, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getUint", segments)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0x3aa311dd.
//
// Solidity: function getUint(string[] segments) view returns(uint256)
func (_EventEmitter *EventEmitterSession) GetUint(segments []string) (*big.Int, error) {
	return _EventEmitter.Contract.GetUint(&_EventEmitter.CallOpts, segments)
}

// GetUint is a free data retrieval call binding the contract method 0x3aa311dd.
//
// Solidity: function getUint(string[] segments) view returns(uint256)
func (_EventEmitter *EventEmitterCallerSession) GetUint(segments []string) (*big.Int, error) {
	return _EventEmitter.Contract.GetUint(&_EventEmitter.CallOpts, segments)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EventEmitter *EventEmitterCaller) GetValue(opts *bind.CallOpts, segments []string) (uint8, []byte, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "getValue", segments)

	if err != nil {
		return *new(uint8), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EventEmitter *EventEmitterSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EventEmitter.Contract.GetValue(&_EventEmitter.CallOpts, segments)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EventEmitter *EventEmitterCallerSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EventEmitter.Contract.GetValue(&_EventEmitter.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EventEmitter *EventEmitterCaller) HasEntry(opts *bind.CallOpts, segments []string) (bool, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "hasEntry", segments)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EventEmitter *EventEmitterSession) HasEntry(segments []string) (bool, error) {
	return _EventEmitter.Contract.HasEntry(&_EventEmitter.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EventEmitter *EventEmitterCallerSession) HasEntry(segments []string) (bool, error) {
	return _EventEmitter.Contract.HasEntry(&_EventEmitter.CallOpts, segments)
}

// IdToName is a free data retrieval call binding the contract method 0xc76853b8.
//
// Solidity: function idToName(string ) view returns(string)
func (_EventEmitter *EventEmitterCaller) IdToName(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "idToName", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IdToName is a free data retrieval call binding the contract method 0xc76853b8.
//
// Solidity: function idToName(string ) view returns(string)
func (_EventEmitter *EventEmitterSession) IdToName(arg0 string) (string, error) {
	return _EventEmitter.Contract.IdToName(&_EventEmitter.CallOpts, arg0)
}

// IdToName is a free data retrieval call binding the contract method 0xc76853b8.
//
// Solidity: function idToName(string ) view returns(string)
func (_EventEmitter *EventEmitterCallerSession) IdToName(arg0 string) (string, error) {
	return _EventEmitter.Contract.IdToName(&_EventEmitter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventEmitter *EventEmitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventEmitter *EventEmitterSession) Owner() (common.Address, error) {
	return _EventEmitter.Contract.Owner(&_EventEmitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventEmitter *EventEmitterCallerSession) Owner() (common.Address, error) {
	return _EventEmitter.Contract.Owner(&_EventEmitter.CallOpts)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_EventEmitter *EventEmitterTransactor) CreateWalletIdentity(opts *bind.TransactOpts, walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "createWalletIdentity", walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_EventEmitter *EventEmitterSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _EventEmitter.Contract.CreateWalletIdentity(&_EventEmitter.TransactOpts, walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_EventEmitter *EventEmitterTransactorSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _EventEmitter.Contract.CreateWalletIdentity(&_EventEmitter.TransactOpts, walletAddress, initialRoles)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_EventEmitter *EventEmitterTransactor) DeleteIdentity(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "deleteIdentity", wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_EventEmitter *EventEmitterSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _EventEmitter.Contract.DeleteIdentity(&_EventEmitter.TransactOpts, wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_EventEmitter *EventEmitterTransactorSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _EventEmitter.Contract.DeleteIdentity(&_EventEmitter.TransactOpts, wallet)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x0277e023.
//
// Solidity: function emitEvent(string name, bytes32[] ArgumentTopics, bytes data) returns(bool success)
func (_EventEmitter *EventEmitterTransactor) EmitEvent(opts *bind.TransactOpts, name string, ArgumentTopics [][32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitEvent", name, ArgumentTopics, data)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x0277e023.
//
// Solidity: function emitEvent(string name, bytes32[] ArgumentTopics, bytes data) returns(bool success)
func (_EventEmitter *EventEmitterSession) EmitEvent(name string, ArgumentTopics [][32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEvent(&_EventEmitter.TransactOpts, name, ArgumentTopics, data)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x0277e023.
//
// Solidity: function emitEvent(string name, bytes32[] ArgumentTopics, bytes data) returns(bool success)
func (_EventEmitter *EventEmitterTransactorSession) EmitEvent(name string, ArgumentTopics [][32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEvent(&_EventEmitter.TransactOpts, name, ArgumentTopics, data)
}

// EmitEvents is a paid mutator transaction binding the contract method 0x8c710456.
//
// Solidity: function emitEvents((string,bytes32[],bytes)[] events) returns(bool success)
func (_EventEmitter *EventEmitterTransactor) EmitEvents(opts *bind.TransactOpts, events []EventEmitterBatchEmitEvent) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitEvents", events)
}

// EmitEvents is a paid mutator transaction binding the contract method 0x8c710456.
//
// Solidity: function emitEvents((string,bytes32[],bytes)[] events) returns(bool success)
func (_EventEmitter *EventEmitterSession) EmitEvents(events []EventEmitterBatchEmitEvent) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEvents(&_EventEmitter.TransactOpts, events)
}

// EmitEvents is a paid mutator transaction binding the contract method 0x8c710456.
//
// Solidity: function emitEvents((string,bytes32[],bytes)[] events) returns(bool success)
func (_EventEmitter *EventEmitterTransactorSession) EmitEvents(events []EventEmitterBatchEmitEvent) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEvents(&_EventEmitter.TransactOpts, events)
}

// EncodePath is a paid mutator transaction binding the contract method 0x2fc37787.
//
// Solidity: function encodePath(string[] segments) returns(bytes)
func (_EventEmitter *EventEmitterTransactor) EncodePath(opts *bind.TransactOpts, segments []string) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "encodePath", segments)
}

// EncodePath is a paid mutator transaction binding the contract method 0x2fc37787.
//
// Solidity: function encodePath(string[] segments) returns(bytes)
func (_EventEmitter *EventEmitterSession) EncodePath(segments []string) (*types.Transaction, error) {
	return _EventEmitter.Contract.EncodePath(&_EventEmitter.TransactOpts, segments)
}

// EncodePath is a paid mutator transaction binding the contract method 0x2fc37787.
//
// Solidity: function encodePath(string[] segments) returns(bytes)
func (_EventEmitter *EventEmitterTransactorSession) EncodePath(segments []string) (*types.Transaction, error) {
	return _EventEmitter.Contract.EncodePath(&_EventEmitter.TransactOpts, segments)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EventEmitter *EventEmitterTransactor) GetOrCreateSegmentId(opts *bind.TransactOpts, segment string) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "getOrCreateSegmentId", segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EventEmitter *EventEmitterSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EventEmitter.Contract.GetOrCreateSegmentId(&_EventEmitter.TransactOpts, segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EventEmitter *EventEmitterTransactorSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EventEmitter.Contract.GetOrCreateSegmentId(&_EventEmitter.TransactOpts, segment)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_EventEmitter *EventEmitterTransactor) GrantRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "grantRole", wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_EventEmitter *EventEmitterSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EventEmitter.Contract.GrantRole(&_EventEmitter.TransactOpts, wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_EventEmitter *EventEmitterTransactorSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EventEmitter.Contract.GrantRole(&_EventEmitter.TransactOpts, wallet, role)
}

// RegisterEventSchema is a paid mutator transaction binding the contract method 0xa4dbe8a0.
//
// Solidity: function registerEventSchema(string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EventEmitter *EventEmitterTransactor) RegisterEventSchema(opts *bind.TransactOpts, name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "registerEventSchema", name, id, eventTopic, args)
}

// RegisterEventSchema is a paid mutator transaction binding the contract method 0xa4dbe8a0.
//
// Solidity: function registerEventSchema(string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EventEmitter *EventEmitterSession) RegisterEventSchema(name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EventEmitter.Contract.RegisterEventSchema(&_EventEmitter.TransactOpts, name, id, eventTopic, args)
}

// RegisterEventSchema is a paid mutator transaction binding the contract method 0xa4dbe8a0.
//
// Solidity: function registerEventSchema(string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EventEmitter *EventEmitterTransactorSession) RegisterEventSchema(name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EventEmitter.Contract.RegisterEventSchema(&_EventEmitter.TransactOpts, name, id, eventTopic, args)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EventEmitter *EventEmitterTransactor) RemoveValue(opts *bind.TransactOpts, segments []string) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "removeValue", segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EventEmitter *EventEmitterSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EventEmitter.Contract.RemoveValue(&_EventEmitter.TransactOpts, segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EventEmitter *EventEmitterTransactorSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EventEmitter.Contract.RemoveValue(&_EventEmitter.TransactOpts, segments)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_EventEmitter *EventEmitterTransactor) RevokeRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "revokeRole", wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_EventEmitter *EventEmitterSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EventEmitter.Contract.RevokeRole(&_EventEmitter.TransactOpts, wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_EventEmitter *EventEmitterTransactorSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EventEmitter.Contract.RevokeRole(&_EventEmitter.TransactOpts, wallet, role)
}

// SetBool is a paid mutator transaction binding the contract method 0x8ae0b3fa.
//
// Solidity: function setBool(string[] segments, bool value) returns()
func (_EventEmitter *EventEmitterTransactor) SetBool(opts *bind.TransactOpts, segments []string, value bool) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setBool", segments, value)
}

// SetBool is a paid mutator transaction binding the contract method 0x8ae0b3fa.
//
// Solidity: function setBool(string[] segments, bool value) returns()
func (_EventEmitter *EventEmitterSession) SetBool(segments []string, value bool) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetBool(&_EventEmitter.TransactOpts, segments, value)
}

// SetBool is a paid mutator transaction binding the contract method 0x8ae0b3fa.
//
// Solidity: function setBool(string[] segments, bool value) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetBool(segments []string, value bool) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetBool(&_EventEmitter.TransactOpts, segments, value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x960a9678.
//
// Solidity: function setBytes(string[] segments, bytes data) returns()
func (_EventEmitter *EventEmitterTransactor) SetBytes(opts *bind.TransactOpts, segments []string, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setBytes", segments, data)
}

// SetBytes is a paid mutator transaction binding the contract method 0x960a9678.
//
// Solidity: function setBytes(string[] segments, bytes data) returns()
func (_EventEmitter *EventEmitterSession) SetBytes(segments []string, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetBytes(&_EventEmitter.TransactOpts, segments, data)
}

// SetBytes is a paid mutator transaction binding the contract method 0x960a9678.
//
// Solidity: function setBytes(string[] segments, bytes data) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetBytes(segments []string, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetBytes(&_EventEmitter.TransactOpts, segments, data)
}

// SetInt is a paid mutator transaction binding the contract method 0x1a66ac39.
//
// Solidity: function setInt(string[] segments, int256 value) returns()
func (_EventEmitter *EventEmitterTransactor) SetInt(opts *bind.TransactOpts, segments []string, value *big.Int) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setInt", segments, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x1a66ac39.
//
// Solidity: function setInt(string[] segments, int256 value) returns()
func (_EventEmitter *EventEmitterSession) SetInt(segments []string, value *big.Int) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetInt(&_EventEmitter.TransactOpts, segments, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x1a66ac39.
//
// Solidity: function setInt(string[] segments, int256 value) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetInt(segments []string, value *big.Int) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetInt(&_EventEmitter.TransactOpts, segments, value)
}

// SetString is a paid mutator transaction binding the contract method 0x562b7be9.
//
// Solidity: function setString(string[] segments, string value) returns()
func (_EventEmitter *EventEmitterTransactor) SetString(opts *bind.TransactOpts, segments []string, value string) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setString", segments, value)
}

// SetString is a paid mutator transaction binding the contract method 0x562b7be9.
//
// Solidity: function setString(string[] segments, string value) returns()
func (_EventEmitter *EventEmitterSession) SetString(segments []string, value string) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetString(&_EventEmitter.TransactOpts, segments, value)
}

// SetString is a paid mutator transaction binding the contract method 0x562b7be9.
//
// Solidity: function setString(string[] segments, string value) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetString(segments []string, value string) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetString(&_EventEmitter.TransactOpts, segments, value)
}

// SetUint is a paid mutator transaction binding the contract method 0x9e85b3ea.
//
// Solidity: function setUint(string[] segments, uint256 value) returns()
func (_EventEmitter *EventEmitterTransactor) SetUint(opts *bind.TransactOpts, segments []string, value *big.Int) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setUint", segments, value)
}

// SetUint is a paid mutator transaction binding the contract method 0x9e85b3ea.
//
// Solidity: function setUint(string[] segments, uint256 value) returns()
func (_EventEmitter *EventEmitterSession) SetUint(segments []string, value *big.Int) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetUint(&_EventEmitter.TransactOpts, segments, value)
}

// SetUint is a paid mutator transaction binding the contract method 0x9e85b3ea.
//
// Solidity: function setUint(string[] segments, uint256 value) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetUint(segments []string, value *big.Int) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetUint(&_EventEmitter.TransactOpts, segments, value)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EventEmitter *EventEmitterTransactor) SetValue(opts *bind.TransactOpts, segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setValue", segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EventEmitter *EventEmitterSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetValue(&_EventEmitter.TransactOpts, segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetValue(&_EventEmitter.TransactOpts, segments, dataType, data)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EventEmitter *EventEmitterTransactor) SetValues(opts *bind.TransactOpts, values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "setValues", values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EventEmitter *EventEmitterSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetValues(&_EventEmitter.TransactOpts, values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EventEmitter *EventEmitterTransactorSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EventEmitter.Contract.SetValues(&_EventEmitter.TransactOpts, values)
}

// EventEmitterIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the EventEmitter contract.
type EventEmitterIdentityCreatedIterator struct {
	Event *EventEmitterIdentityCreated // Event containing the contract specifics and raw log

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
func (it *EventEmitterIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterIdentityCreated)
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
		it.Event = new(EventEmitterIdentityCreated)
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
func (it *EventEmitterIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterIdentityCreated represents a IdentityCreated event raised by the EventEmitter contract.
type EventEmitterIdentityCreated struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_EventEmitter *EventEmitterFilterer) FilterIdentityCreated(opts *bind.FilterOpts, wallet []common.Address) (*EventEmitterIdentityCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterIdentityCreatedIterator{contract: _EventEmitter.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_EventEmitter *EventEmitterFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *EventEmitterIdentityCreated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterIdentityCreated)
				if err := _EventEmitter.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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

// ParseIdentityCreated is a log parse operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_EventEmitter *EventEmitterFilterer) ParseIdentityCreated(log types.Log) (*EventEmitterIdentityCreated, error) {
	event := new(EventEmitterIdentityCreated)
	if err := _EventEmitter.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterNewSegmentIterator is returned from FilterNewSegment and is used to iterate over the raw logs and unpacked data for NewSegment events raised by the EventEmitter contract.
type EventEmitterNewSegmentIterator struct {
	Event *EventEmitterNewSegment // Event containing the contract specifics and raw log

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
func (it *EventEmitterNewSegmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterNewSegment)
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
		it.Event = new(EventEmitterNewSegment)
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
func (it *EventEmitterNewSegmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterNewSegmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterNewSegment represents a NewSegment event raised by the EventEmitter contract.
type EventEmitterNewSegment struct {
	Segment string
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSegment is a free log retrieval operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EventEmitter *EventEmitterFilterer) FilterNewSegment(opts *bind.FilterOpts, id []*big.Int) (*EventEmitterNewSegmentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterNewSegmentIterator{contract: _EventEmitter.contract, event: "NewSegment", logs: logs, sub: sub}, nil
}

// WatchNewSegment is a free log subscription operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EventEmitter *EventEmitterFilterer) WatchNewSegment(opts *bind.WatchOpts, sink chan<- *EventEmitterNewSegment, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterNewSegment)
				if err := _EventEmitter.contract.UnpackLog(event, "NewSegment", log); err != nil {
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

// ParseNewSegment is a log parse operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EventEmitter *EventEmitterFilterer) ParseNewSegment(log types.Log) (*EventEmitterNewSegment, error) {
	event := new(EventEmitterNewSegment)
	if err := _EventEmitter.contract.UnpackLog(event, "NewSegment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EventEmitter contract.
type EventEmitterRoleGrantedIterator struct {
	Event *EventEmitterRoleGranted // Event containing the contract specifics and raw log

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
func (it *EventEmitterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterRoleGranted)
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
		it.Event = new(EventEmitterRoleGranted)
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
func (it *EventEmitterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterRoleGranted represents a RoleGranted event raised by the EventEmitter contract.
type EventEmitterRoleGranted struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_EventEmitter *EventEmitterFilterer) FilterRoleGranted(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*EventEmitterRoleGrantedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterRoleGrantedIterator{contract: _EventEmitter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_EventEmitter *EventEmitterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EventEmitterRoleGranted, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterRoleGranted)
				if err := _EventEmitter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_EventEmitter *EventEmitterFilterer) ParseRoleGranted(log types.Log) (*EventEmitterRoleGranted, error) {
	event := new(EventEmitterRoleGranted)
	if err := _EventEmitter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterSchemaRegisteredIterator is returned from FilterSchemaRegistered and is used to iterate over the raw logs and unpacked data for SchemaRegistered events raised by the EventEmitter contract.
type EventEmitterSchemaRegisteredIterator struct {
	Event *EventEmitterSchemaRegistered // Event containing the contract specifics and raw log

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
func (it *EventEmitterSchemaRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterSchemaRegistered)
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
		it.Event = new(EventEmitterSchemaRegistered)
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
func (it *EventEmitterSchemaRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterSchemaRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterSchemaRegistered represents a SchemaRegistered event raised by the EventEmitter contract.
type EventEmitterSchemaRegistered struct {
	Name common.Hash
	Id   string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSchemaRegistered is a free log retrieval operation binding the contract event 0x4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de.
//
// Solidity: event SchemaRegistered(string indexed name, string id)
func (_EventEmitter *EventEmitterFilterer) FilterSchemaRegistered(opts *bind.FilterOpts, name []string) (*EventEmitterSchemaRegisteredIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "SchemaRegistered", nameRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterSchemaRegisteredIterator{contract: _EventEmitter.contract, event: "SchemaRegistered", logs: logs, sub: sub}, nil
}

// WatchSchemaRegistered is a free log subscription operation binding the contract event 0x4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de.
//
// Solidity: event SchemaRegistered(string indexed name, string id)
func (_EventEmitter *EventEmitterFilterer) WatchSchemaRegistered(opts *bind.WatchOpts, sink chan<- *EventEmitterSchemaRegistered, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "SchemaRegistered", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterSchemaRegistered)
				if err := _EventEmitter.contract.UnpackLog(event, "SchemaRegistered", log); err != nil {
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

// ParseSchemaRegistered is a log parse operation binding the contract event 0x4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de.
//
// Solidity: event SchemaRegistered(string indexed name, string id)
func (_EventEmitter *EventEmitterFilterer) ParseSchemaRegistered(log types.Log) (*EventEmitterSchemaRegistered, error) {
	event := new(EventEmitterSchemaRegistered)
	if err := _EventEmitter.contract.UnpackLog(event, "SchemaRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterFactoryMetaData contains all meta data concerning the EventEmitterFactory contract.
var EventEmitterFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ContractAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyContractABI\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CustomContractAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"}],\"name\":\"CustomContractSchemaAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"emitterAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EmitterCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractABI\",\"type\":\"string\"}],\"name\":\"addCustomContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"addCustomContractSchema\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createEmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customContractAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractABI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"emitters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"emitterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getCustomContractABI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getCustomContractEventNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"}],\"name\":\"getCustomContractSchema\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCustomContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmitters\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"emitterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structEventEmitterFactory.EmitterInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalSchemas\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051615949380380615949833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b6158c3806100865f395ff3fe608060405234801561000f575f5ffd5b50600436106100ca575f3560e01c8063a66ec86811610088578063bef863a611610063578063bef863a6146101db578063bfa33281146101f0578063ccfaf63c14610210578063e22c27ae14610223575f5ffd5b8063a66ec86814610193578063abd60974146101a8578063bd68dd10146101bb575f5ffd5b80621f75d9146100ce578063176fa407146100f95780631c32f662146101195780633a5381b5146101395780635a0911c41461014b57806380bd201d1461017e575b5f5ffd5b6100e16100dc366004611569565b610238565b6040516100f0939291906115b7565b60405180910390f35b610101610370565b6040516001600160a01b0390911681526020016100f0565b61012c610127366004611569565b61048d565b6040516100f091906115f6565b5f54610101906001600160a01b031681565b61015e610159366004611659565b610579565b604080516001600160a01b039384168152929091166020830152016100f0565b61019161018c366004611747565b6105b0565b005b61019b61071b565b6040516100f091906118a4565b6101016101b6366004611659565b610c08565b6101ce6101c9366004611569565b610c30565b6040516100f091906118fb565b6101e3610d12565b6040516100f0919061190d565b6102036101fe366004611958565b610d72565b6040516100f091906119a2565b61019161021e3660046119b4565b6110bd565b61022b61139a565b6040516100f09190611b6b565b60026020525f9081526040902080546001820180546001600160a01b03909216929161026390611bba565b80601f016020809104026020016040519081016040528092919081815260200182805461028f90611bba565b80156102da5780601f106102b1576101008083540402835291602001916102da565b820191905f5260205f20905b8154815290600101906020018083116102bd57829003601f168201915b5050505050908060030180546102ef90611bba565b80601f016020809104026020016040519081016040528092919081815260200182805461031b90611bba565b80156103665780601f1061033d57610100808354040283529160200191610366565b820191905f5260205f20905b81548152906001019060200180831161034957829003601f168201915b5050505050905083565b5f8054604051829133916001600160a01b039091169061038f90611483565b6001600160a01b03928316815291166020820152604001604051809103905ff0801580156103bf573d5f5f3e3d5ffd5b506040805180820182526001600160a01b0380841680835233602084018181526001805480820182555f918252955160029096027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6810180549787166001600160a01b031998891617905591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf790920180549290951691909516179092559251939450927fd69887947906502e8670345f1af75cd1bc3be116978564188dda62fad0aed1ba9190a3919050565b6001600160a01b0381165f9081526002602081815260408084209092018054835181840281018401909452808452606094919290919084015b8282101561056e578382905f5260205f200180546104e390611bba565b80601f016020809104026020016040519081016040528092919081815260200182805461050f90611bba565b801561055a5780601f106105315761010080835404028352916020019161055a565b820191905f5260205f20905b81548152906001019060200180831161053d57829003601f168201915b5050505050815260200190600101906104c6565b505050509050919050565b60018181548110610588575f80fd5b5f918252602090912060029091020180546001909101546001600160a01b0391821692501682565b6001600160a01b038381165f9081526002602052604090205416156105e85760405163e8dc2ba560e01b815260040160405180910390fd5b80515f036106085760405162e32e1d60e21b815260040160405180910390fd5b6001600160a01b0383165f81815260026020526040902080546001600160a01b03191690911781556001810161063e8482611c38565b506003810161064d8382611c38565b50604080515f8082526020820190925290610678565b60608152602001906001900390816106635790505b50805161068f916002840191602090910190611490565b50600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b0386169081179091556040517f994774dc478a0a24cbad071ebe255a31e5868d3bce654611faa96b645a940f609061070d9086906118fb565b60405180910390a250505050565b60408051600580825260c082019092526060915f9190816020015b6107696040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b815260200190600190039081610736575050604080516003808252608082019092529192505f9190602082015b604080516060808201835280825260208201525f918101919091528152602001906001900390816107965750506040805160a0810182526004606082019081526366726f6d60e01b608083015281528151808301835260078152666164647265737360c81b60208281019190915282015260019181019190915281519192509082905f9061082657610826611cf2565b6020026020010181905250604051806060016040528060405180604001604052806002815260200161746f60f01b8152508152602001604051806040016040528060078152602001666164647265737360c81b8152508152602001600115158152508160018151811061089b5761089b611cf2565b602002602001018190525060405180606001604052806040518060400160405280600581526020016476616c756560d81b8152508152602001604051806040016040528060078152602001663ab4b73a191a9b60c91b81525081526020015f15158152508160028151811061091257610912611cf2565b60200260200101819052506040518060a00160405280604051806040016040528060088152602001672a3930b739b332b960c11b815250815260200160405180608001604052806044815260200161584a6044913981526020018281526020017fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8152602001600260ff16815250825f815181106109b2576109b2611cf2565b6020908102919091010152604080516003808252608082019092525f91816020015b604080516060808201835280825260208201525f918101919091528152602001906001900390816109d45750506040805160a0810182526004606082019081526366726f6d60e01b608083015281528151808301835260078152666164647265737360c81b60208281019190915282015260019181019190915281519192509082905f90610a6457610a64611cf2565b6020026020010181905250604051806060016040528060405180604001604052806002815260200161746f60f01b8152508152602001604051806040016040528060078152602001666164647265737360c81b81525081526020016001151581525081600181518110610ad957610ad9611cf2565b60200260200101819052506040518060600160405280604051806040016040528060078152602001661d1bdad95b925960ca1b8152508152602001604051806040016040528060078152602001663ab4b73a191a9b60c91b81525081526020016001151581525081600281518110610b5357610b53611cf2565b60200260200101819052506040518060a00160405280604051806040016040528060088152602001672a3930b739b332b960c11b81525081526020016040518060800160405280604481526020016158066044913981526020018281526020017fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8152602001600360ff1681525083600281518110610bf457610bf4611cf2565b602002602001018190525082935050505090565b60038181548110610c17575f80fd5b5f918252602090912001546001600160a01b0316905081565b6001600160a01b038181165f9081526002602052604090205460609116610c6a57604051633b56498960e21b815260040160405180910390fd5b6001600160a01b0382165f9081526002602052604090206003018054610c8f90611bba565b80601f0160208091040260200160405190810160405280929190818152602001828054610cbb90611bba565b8015610d065780601f10610cdd57610100808354040283529160200191610d06565b820191905f5260205f20905b815481529060010190602001808311610ce957829003601f168201915b50505050509050919050565b60606003805480602002602001604051908101604052809291908181526020018280548015610d6857602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610d4a575b5050505050905090565b610da56040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b60025f846001600160a01b03166001600160a01b031681526020019081526020015f2060040182604051610dd99190611d06565b90815260200160405180910390206040518060a00160405290815f82018054610e0190611bba565b80601f0160208091040260200160405190810160405280929190818152602001828054610e2d90611bba565b8015610e785780601f10610e4f57610100808354040283529160200191610e78565b820191905f5260205f20905b815481529060010190602001808311610e5b57829003601f168201915b50505050508152602001600182018054610e9190611bba565b80601f0160208091040260200160405190810160405280929190818152602001828054610ebd90611bba565b8015610f085780601f10610edf57610100808354040283529160200191610f08565b820191905f5260205f20905b815481529060010190602001808311610eeb57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015611096578382905f5260205f2090600302016040518060600160405290815f82018054610f6590611bba565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9190611bba565b8015610fdc5780601f10610fb357610100808354040283529160200191610fdc565b820191905f5260205f20905b815481529060010190602001808311610fbf57829003601f168201915b50505050508152602001600182018054610ff590611bba565b80601f016020809104026020016040519081016040528092919081815260200182805461102190611bba565b801561106c5780601f106110435761010080835404028352916020019161106c565b820191905f5260205f20905b81548152906001019060200180831161104f57829003601f168201915b50505091835250506002919091015460ff1615156020918201529082526001929092019101610f35565b505050908252506003820154602082015260049091015460ff166040909101529392505050565b6001600160a01b038581165f908152600260205260409020541661111d5760405162461bcd60e51b815260206004820152601260248201527110dbdb9d1c9858dd081b9bdd08199bdd5b9960721b60448201526064015b60405180910390fd5b6001600160a01b0385165f908152600260205260409081902090516004820190611148908790611d06565b908152604051908190036020019020805461116290611bba565b1590506111b15760405162461bcd60e51b815260206004820152601d60248201527f4576656e74206e616d6520616c726561647920726567697374657265640000006044820152606401611114565b5f805b83518110156111f6578381815181106111cf576111cf611cf2565b602002602001015160400151156111ee57816111ea81611d1c565b9250505b6001016111b4565b5060038160ff16111561124b5760405162461bcd60e51b815260206004820152601a60248201527f546f6f206d616e7920696e646578656420617267756d656e74730000000000006044820152606401611114565b5f826004018760405161125e9190611d06565b90815260405190819003602001902090508061127a8882611c38565b50600181016112898782611c38565b506003810185905560048101805460ff191660ff84161790555f5b845181101561132d57816002018582815181106112c3576112c3611cf2565b60209081029190910181015182546001810184555f9384529190922082516003909202019081906112f49082611c38565b50602082015160018201906113099082611c38565b50604091909101516002909101805460ff19169115159190911790556001016112a4565b506002830180546001810182555f91825260209091200161134e8882611c38565b50876001600160a01b03167f8c7d8ad02950ea5bd57f54af9c47ca84706c82e280cada0c04a2da9f2bb4818c8860405161138891906118fb565b60405180910390a25050505050505050565b6001546060905f906001600160401b038111156113b9576113b9611670565b6040519080825280602002602001820160405280156113fd57816020015b604080518082019091525f80825260208201528152602001906001900390816113d75790505b5090505f5b60015481101561147d576001818154811061141f5761141f611cf2565b5f9182526020918290206040805180820190915260029092020180546001600160a01b0390811683526001909101541691810191909152825183908390811061146a5761146a611cf2565b6020908102919091010152600101611402565b50919050565b613abf80611d4783390190565b828054828255905f5260205f209081019282156114d4579160200282015b828111156114d457825182906114c49082611c38565b50916020019190600101906114ae565b506114e09291506114e4565b5090565b808211156114e0575f6114f78282611500565b506001016114e4565b50805461150c90611bba565b5f825580601f1061151b575050565b601f0160209004905f5260205f2090810190611537919061153a565b50565b5b808211156114e0575f815560010161153b565b80356001600160a01b0381168114611564575f5ffd5b919050565b5f60208284031215611579575f5ffd5b6115828261154e565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03841681526060602082018190525f906115da90830185611589565b82810360408401526115ec8185611589565b9695505050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561164d57603f19878603018452611638858351611589565b9450602093840193919091019060010161161c565b50929695505050505050565b5f60208284031215611669575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156116a6576116a6611670565b60405290565b604051601f8201601f191681016001600160401b03811182821017156116d4576116d4611670565b604052919050565b5f82601f8301126116eb575f5ffd5b81356001600160401b0381111561170457611704611670565b611717601f8201601f19166020016116ac565b81815284602083860101111561172b575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f60608486031215611759575f5ffd5b6117628461154e565b925060208401356001600160401b0381111561177c575f5ffd5b611788868287016116dc565b92505060408401356001600160401b038111156117a3575f5ffd5b6117af868287016116dc565b9150509250925092565b5f815160a084526117cd60a0850182611589565b9050602083015184820360208601526117e68282611589565b9150506040830151848203604086015281815180845260208401915060208160051b8501016020840193505f5b8281101561187857601f1986830301845284518051606084526118396060850182611589565b9050602082015184820360208601526118528282611589565b604093840151151595909301949094525060209586019594909401939150600101611813565b506060870151606089015260808701519450611899608089018660ff169052565b979650505050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561164d57603f198786030184526118e68583516117b9565b945060209384019391909101906001016118ca565b602081525f6115826020830184611589565b602080825282518282018190525f918401906040840190835b8181101561194d5783516001600160a01b0316835260209384019390920191600101611926565b509095945050505050565b5f5f60408385031215611969575f5ffd5b6119728361154e565b915060208301356001600160401b0381111561198c575f5ffd5b611998858286016116dc565b9150509250929050565b602081525f61158260208301846117b9565b5f5f5f5f5f60a086880312156119c8575f5ffd5b6119d18661154e565b945060208601356001600160401b038111156119eb575f5ffd5b6119f7888289016116dc565b94505060408601356001600160401b03811115611a12575f5ffd5b611a1e888289016116dc565b9350506060860135915060808601356001600160401b03811115611a40575f5ffd5b8601601f81018813611a50575f5ffd5b80356001600160401b03811115611a6957611a69611670565b8060051b611a79602082016116ac565b9182526020818401810192908101908b841115611a94575f5ffd5b6020850192505b83831015611b595782356001600160401b03811115611ab8575f5ffd5b85016060818e03601f19011215611acd575f5ffd5b611ad5611684565b60208201356001600160401b03811115611aed575f5ffd5b611afc8f6020838601016116dc565b82525060408201356001600160401b03811115611b17575f5ffd5b611b268f6020838601016116dc565b602083015250606082013591508115158214611b40575f5ffd5b6040810191909152825260209283019290910190611a9b565b80955050505050509295509295909350565b602080825282518282018190525f918401906040840190835b8181101561194d57835180516001600160a01b039081168552602091820151168185015290930192604090920191600101611b84565b600181811c90821680611bce57607f821691505b60208210810361147d57634e487b7160e01b5f52602260045260245ffd5b601f821115611c3357805f5260205f20601f840160051c81016020851015611c115750805b601f840160051c820191505b81811015611c30575f8155600101611c1d565b50505b505050565b81516001600160401b03811115611c5157611c51611670565b611c6581611c5f8454611bba565b84611bec565b6020601f821160018114611c97575f8315611c805750848201515b5f19600385901b1c1916600184901b178455611c30565b5f84815260208120601f198516915b82811015611cc65787850151825560209485019460019092019101611ca6565b5084821015611ce357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffd5b5f82518060208501845e5f920191825250919050565b5f60ff821660ff8103611d3d57634e487b7160e01b5f52601160045260245ffd5b6001019291505056fe608060405234801561000f575f5ffd5b50604051613abf380380613abf83398101604081905261002e916103ae565b61003960088361005e565b600180555f80546001600160a01b0319166001600160a01b0383161790555050610407565b6002820180546001600160a01b0319166001600160a01b03831617905560408051600480825260a082019092525f91602082016080803683370190505090506002815f815181106100b1576100b16103df565b602002602001019060038111156100ca576100ca6103f3565b908160038111156100dd576100dd6103f3565b815250505f816001815181106100f5576100f56103df565b6020026020010190600381111561010e5761010e6103f3565b90816003811115610121576101216103f3565b8152505060018160028151811061013a5761013a6103df565b60200260200101906003811115610153576101536103f3565b90816003811115610166576101666103f3565b8152505060038160038151811061017f5761017f6103df565b60200260200101906003811115610198576101986103f3565b908160038111156101ab576101ab6103f3565b9052506101b98383836101be565b505050565b6001600160a01b0382165f9081526020849052604090205460ff16156101f7576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602085815260408220805460ff191660019081179091558681018054918201815583529082200180546001600160a01b0319169092179091555b815181101561027a57610272848484848151811061025f5761025f6103df565b60200260200101516102b360201b60201c565b60010161023f565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a2505050565b6001600160a01b0382165f9081526020849052604090205460ff166102eb5760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208481526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a918202191690836003811115610343576103436103f3565b021790555081600381111561035a5761035a6103f3565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a350505050565b80516001600160a01b03811681146103a9575f5ffd5b919050565b5f5f604083850312156103bf575f5ffd5b6103c883610393565b91506103d660208401610393565b90509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b6136ab806104145f395ff3fe608060405234801561000f575f5ffd5b5060043610610208575f3560e01c806369da55d21161011f578063a4dbe8a0116100a9578063c8740be111610079578063c8740be1146104b3578063d50d542e146104c6578063dced4783146104d9578063e6149cd0146104ec578063f30826f0146104ff575f5ffd5b8063a4dbe8a014610467578063a8d29d1d1461047a578063ab8412831461048d578063c76853b8146104a0575f5ffd5b80638ae0b3fa116100ef5780638ae0b3fa146104005780638c710456146104135780638da5cb5b14610426578063960a9678146104415780639e85b3ea14610454575f5ffd5b806369da55d2146103b257806371934ef0146103c55780637eeede35146103d85780638998fbe8146103ed575f5ffd5b80633aa311dd116101a057806349ab123d1161017057806349ab123d146103385780634e0a38dc14610358578063562b7be91461036b57806360778d951461037e57806364ad6c881461039f575f5ffd5b80633aa311dd146102dc5780633e840236146102fd5780633f54d8bf146103105780634151df6114610325575f5ffd5b80631a66ac39116101db5780631a66ac39146102715780631bfa8601146102845780632f2545d2146102995780632fc37787146102bc575f5ffd5b80630277e0231461020c5780630761b8d014610234578063077d3c031461024957806317be85c31461025c575b5f5ffd5b61021f61021a366004612208565b610507565b60405190151581526020015b60405180910390f35b6102476102423660046122c0565b61052d565b005b610247610257366004612360565b6105b0565b61026461062b565b60405161022b91906123ed565b61024761027f366004612497565b6106a6565b61028c610723565b60405161022b91906124de565b6102ac6102a7366004612529565b61079a565b60405161022b949392919061256d565b6102cf6102ca3660046125ac565b6108dd565b60405161022b91906125ea565b6102ef6102ea3660046125ac565b6108f1565b60405190815260200161022b565b61024761030b366004612360565b61096d565b6103186109b7565b60405161022b91906125fc565b6102ef610333366004612653565b610a2d565b61034b610346366004612653565b610a6a565b60405161022b9190612685565b61021f6103663660046125ac565b610e5d565b610247610379366004612772565b610ed9565b61039161038c3660046125ac565b610f59565b60405161022b9291906127dc565b61021f6103ad3660046125ac565b610fe6565b6102476103c03660046125ac565b611023565b6102ef6103d3366004612653565b61106c565b6103e06110a9565b60405161022b919061280b565b6102476103fb3660046125ac565b611120565b61024761040e3660046128c1565b61118b565b61021f6104213660046125ac565b6111d6565b600a546040516001600160a01b03909116815260200161022b565b61024761044f366004612772565b61128a565b610247610462366004612497565b6112d7565b610247610475366004612913565b611322565b6102476104883660046129b8565b61157d565b6102cf61049b3660046125ac565b6115d2565b6102cf6104ae366004612529565b611651565b6102ef6104c13660046125ac565b6116f4565b61034b6104d4366004612653565b611731565b6102476104e73660046129d3565b611aa7565b6102cf6104fa3660046125ac565b611af3565b610318611b31565b5f8061051560083383611c05565b6105228787878787611c69565b979650505050505050565b600361053b60083383611c05565b60405163641bf91960e11b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063c837f2329061057c905f908a908a908a908a908a90600401612aed565b5f6040518083038186803b158015610592575f5ffd5b505af41580156105a4573d5f5f3e3d5ffd5b50505050505050505050565b60026105be60083383611c05565b604051630d3d280360e21b815273__$5a92de94e3c11b56e4ee4f365f96da5afc$__906334f4a00c906105fa9060089087908790600401612b33565b5f6040518083038186803b158015610610575f5ffd5b505af4158015610622573d5f5f3e3d5ffd5b50505050505050565b60405163b64ee37d60e01b81525f600482015260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063b64ee37d906024015f60405180830381865af415801561067a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a19190810190612bc7565b905090565b60036106b460083383611c05565b60405163f2bffc1360e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063f2bffc13906106f1905f90889088908890600401612d05565b5f6040518083038186803b158015610707575f5ffd5b505af4158015610719573d5f5f3e3d5ffd5b5050505050505050565b604051630cf1f1fb60e21b81526008600482015260609073__$5a92de94e3c11b56e4ee4f365f96da5afc$__906333c7c7ec906024015f60405180830381865af4158015610773573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a19190810190612d2f565b8051602081830181018051600b825292820191909301209152805481906107c090612dc8565b80601f01602080910402602001604051908101604052809291908181526020018280546107ec90612dc8565b80156108375780601f1061080e57610100808354040283529160200191610837565b820191905f5260205f20905b81548152906001019060200180831161081a57829003601f168201915b50505050509080600101805461084c90612dc8565b80601f016020809104026020016040519081016040528092919081815260200182805461087890612dc8565b80156108c35780601f1061089a576101008083540402835291602001916108c3565b820191905f5260205f20905b8154815290600101906020018083116108a657829003601f168201915b50505050600383015460049093015491929160ff16905084565b60606108ea5f8484611d98565b9392505050565b604051632519093f60e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__90632519093f9061092e90849087908790600401612e00565b602060405180830381865af4158015610949573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ea9190612e22565b600261097b60083383611c05565b60405163310adb3d60e11b815273__$5a92de94e3c11b56e4ee4f365f96da5afc$__90636215b67a906105fa9060089087908790600401612b33565b60405163624ff69960e11b81525f600482015260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063c49fed32906024015f60405180830381865af4158015610a06573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a19190810190612e39565b604051631a770e2b60e11b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__906334ee1c569061092e90849087908790600401612edd565b610a9d6040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b5f600c8484604051610ab0929190612ef6565b90815260200160405180910390208054610ac990612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610af590612dc8565b8015610b405780601f10610b1757610100808354040283529160200191610b40565b820191905f5260205f20905b815481529060010190602001808311610b2357829003601f168201915b5050505050905080515f03610b6857604051634fde869d60e01b815260040160405180910390fd5b600b81604051610b789190612f1c565b90815260200160405180910390206040518060a00160405290815f82018054610ba090612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610bcc90612dc8565b8015610c175780601f10610bee57610100808354040283529160200191610c17565b820191905f5260205f20905b815481529060010190602001808311610bfa57829003601f168201915b50505050508152602001600182018054610c3090612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5c90612dc8565b8015610ca75780601f10610c7e57610100808354040283529160200191610ca7565b820191905f5260205f20905b815481529060010190602001808311610c8a57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015610e35578382905f5260205f2090600302016040518060600160405290815f82018054610d0490612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3090612dc8565b8015610d7b5780601f10610d5257610100808354040283529160200191610d7b565b820191905f5260205f20905b815481529060010190602001808311610d5e57829003601f168201915b50505050508152602001600182018054610d9490612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054610dc090612dc8565b8015610e0b5780601f10610de257610100808354040283529160200191610e0b565b820191905f5260205f20905b815481529060010190602001808311610dee57829003601f168201915b50505091835250506002919091015460ff1615156020918201529082526001929092019101610cd4565b505050908252506003820154602082015260049091015460ff16604090910152949350505050565b60405163dea25e9f60e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063dea25e9f90610e9a90849087908790600401612e00565b602060405180830381865af4158015610eb5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ea9190612f27565b6003610ee760083383611c05565b604051639a7044d760e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__90639a7044d790610f26905f908990899089908990600401612f42565b5f6040518083038186803b158015610f3c575f5ffd5b505af4158015610f4e573d5f5f3e3d5ffd5b505050505050505050565b604051639e72831160e01b81525f9060609073__$d451bf733f00e36b69819fe95d7d8eab5b$__90639e72831190610f9990859088908890600401612e00565b5f60405180830381865af4158015610fb3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610fda9190810190612f7a565b915091505b9250929050565b604051635d0aad2360e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__90635d0aad2390610e9a90849087908790600401612e00565b600361103160083383611c05565b604051638b7b2e1760e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__90638b7b2e17906105fa905f9087908790600401612fc7565b6040516332a1c4bb60e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__906332a1c4bb9061092e90849087908790600401612edd565b604051639b2642df60e01b81526008600482015260609073__$5a92de94e3c11b56e4ee4f365f96da5afc$__90639b2642df906024015f60405180830381865af41580156110f9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a191908101906130cf565b6040516340efd12960e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__906340efd1299061115b905f9086908690600401612e00565b5f6040518083038186803b158015611171575f5ffd5b505af4158015611183573d5f5f3e3d5ffd5b505050505050565b600361119960083383611c05565b60405163720e983960e11b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063e41d3072906106f1905f9088908890889060040161322a565b5f806111e460083383611c05565b5f5b8381101561127f573685858381811061120157611201613256565b9050602002810190611213919061326a565b90506112756112228280613288565b61122f60208501856132ca565b61123c6040870187613288565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611c6992505050565b50506001016111e6565b506001949350505050565b600361129860083383611c05565b6040516338fe9c7b60e21b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__9063e3fa71ec90610f26905f908990899089908990600401612f42565b60036112e560083383611c05565b604051631b33bcaf60e01b815273__$d451bf733f00e36b69819fe95d7d8eab5b$__90631b33bcaf906106f1905f90889088908890600401612d05565b600161133060083383611c05565b600b8888604051611342929190612ef6565b908152604051908190036020019020805461135c90612dc8565b15905061137c57604051630268c2d760e51b815260040160405180910390fd5b5f805b838110156113d75784848281811061139957611399613256565b90506020028101906113ab919061326a565b6113bc90606081019060400161330f565b156113cf57816113cb8161333e565b9250505b60010161137f565b5060038160ff1611156113fc57604051626f3b7d60e31b815260040160405180910390fd5b5f600b8a8a60405161140f929190612ef6565b90815260405190819003602001902090508061142c8a8c836133a7565b506001810161143c888a836133a7565b50600381018690555f5b848110156114a1578160020186868381811061146457611464613256565b9050602002810190611476919061326a565b81546001810183555f92835260209092209091600302016114978282613472565b5050600101611446565b5060048101805460ff191660ff8416179055600d80546001810182555f919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5016114ef8a8c836133a7565b508989600c8a8a604051611504929190612ef6565b9081526020016040518091039020918261151f9291906133a7565b508989604051611530929190612ef6565b60405180910390207f4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de898960405161156992919061357c565b60405180910390a250505050505050505050565b600261158b60083383611c05565b60405163f76888e760e01b8152600860048201526001600160a01b038316602482015273__$5a92de94e3c11b56e4ee4f365f96da5afc$__9063f76888e79060440161115b565b604051633a555c0b60e21b815260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__9063e955702c90611610905f9087908790600401612e00565b5f60405180830381865af415801561162a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108ea919081019061358f565b8051602081830181018051600c825292820191909301209152805461167590612dc8565b80601f01602080910402602001604051908101604052809291908181526020018280546116a190612dc8565b80156116ec5780601f106116c3576101008083540402835291602001916116ec565b820191905f5260205f20905b8154815290600101906020018083116116cf57829003601f168201915b505050505081565b604051631c74ab5760e01b81525f9073__$d451bf733f00e36b69819fe95d7d8eab5b$__90631c74ab579061092e90849087908790600401612e00565b6117646040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b600b8383604051611776929190612ef6565b908152604051908190036020019020805461179090612dc8565b90505f036117b157604051634fde869d60e01b815260040160405180910390fd5b600b83836040516117c3929190612ef6565b90815260200160405180910390206040518060a00160405290815f820180546117eb90612dc8565b80601f016020809104026020016040519081016040528092919081815260200182805461181790612dc8565b80156118625780601f1061183957610100808354040283529160200191611862565b820191905f5260205f20905b81548152906001019060200180831161184557829003601f168201915b5050505050815260200160018201805461187b90612dc8565b80601f01602080910402602001604051908101604052809291908181526020018280546118a790612dc8565b80156118f25780601f106118c9576101008083540402835291602001916118f2565b820191905f5260205f20905b8154815290600101906020018083116118d557829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015611a80578382905f5260205f2090600302016040518060600160405290815f8201805461194f90612dc8565b80601f016020809104026020016040519081016040528092919081815260200182805461197b90612dc8565b80156119c65780601f1061199d576101008083540402835291602001916119c6565b820191905f5260205f20905b8154815290600101906020018083116119a957829003601f168201915b505050505081526020016001820180546119df90612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054611a0b90612dc8565b8015611a565780601f10611a2d57610100808354040283529160200191611a56565b820191905f5260205f20905b815481529060010190602001808311611a3957829003601f168201915b50505091835250506002919091015460ff161515602091820152908252600192909201910161191f565b505050908252506003820154602082015260049091015460ff166040909101529392505050565b6002611ab560083383611c05565b604051631fbbcd3760e31b815273__$5a92de94e3c11b56e4ee4f365f96da5afc$__9063fdde69b8906106f1906008908890889088906004016135c0565b604051635551774560e01b815260609073__$d451bf733f00e36b69819fe95d7d8eab5b$__90635551774590611610905f9087908790600401612e00565b6060600d805480602002602001604051908101604052809291908181526020015f905b82821015611bfc578382905f5260205f20018054611b7190612dc8565b80601f0160208091040260200160405190810160405280929190818152602001828054611b9d90612dc8565b8015611be85780601f10611bbf57610100808354040283529160200191611be8565b820191905f5260205f20905b815481529060010190602001808311611bcb57829003601f168201915b505050505081526020019060010190611b54565b50505050905090565b6001600160a01b0382165f9081526020849052604090205460ff16611c3d5760405163c597feeb60e01b815260040160405180910390fd5b611c48838383611dad565b611c64576040516282b42960e81b815260040160405180910390fd5b505050565b5f600b8686604051611c7c929190612ef6565b9081526040519081900360200190208054611c9690612dc8565b90505f03611cb757604051634fde869d60e01b815260040160405180910390fd5b600b8686604051611cc9929190612ef6565b9081526040519081900360200190206004015460ff168314611cfe57604051631f73a41360e01b815260040160405180910390fd5b5f600b8787604051611d11929190612ef6565b90815260200160405180910390206003015490505f8351905060208401855f8114611d535760018114611d5c5760028114611d675760038114611d7757611d88565b838383a1611d88565b8735848484a2611d88565b60208801358835858585a3611d88565b604088013560208901358935868686a45b5060019998505050505050505050565b6060611da5848484611e48565b949350505050565b6001600160a01b0382165f908152602084905260408120600101815b8154811015611e3d57836003811115611de457611de46123c5565b828281548110611df657611df6613256565b905f5260205f2090602091828204019190069054906101000a900460ff166003811115611e2557611e256123c5565b03611e35576001925050506108ea565b600101611dc9565b505f95945050505050565b604080515f808252602082019092526060915b83811015611ec9575f611e9187878785818110611e7a57611e7a613256565b9050602002810190611e8c9190613288565b611ef6565b905082611e9d82611ffe565b604051602001611eae929190613625565b60408051601f19818403018152919052925050600101611e5b565b50604051611edd9082905f90602001613639565b60408051808303601f1901815291905295945050505050565b5f5f846002018484604051611f0c929190612ef6565b9081526040805160209281900383019020545f8181526004890190935291205490915060ff1615611f3e5790506108ea565b6001850180549081905f611f518361365d565b919050555080866002018686604051611f6b929190612ef6565b90815260405160209181900382019020919091556003870180546001810182555f91825291902001611f9e8587836133a7565b505f81815260048701602052604090819020805460ff191660011790555181907f4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e77890611fed908890889061357c565b60405180910390a291506108ea9050565b6060815f03612029576040515f60208201526021016040516020818303038152906040529050919050565b6060825b6080811061206b578181607f1660801760f81b604051602001612051929190613639565b60408051601f19818403018152919052915060071c61202d565b8181607f1660f81b604051602001612084929190613639565b60408051601f19818403018152919052949350505050565b5f5f83601f8401126120ac575f5ffd5b5081356001600160401b038111156120c2575f5ffd5b602083019150836020828501011115610fdf575f5ffd5b5f5f83601f8401126120e9575f5ffd5b5081356001600160401b038111156120ff575f5ffd5b6020830191508360208260051b8501011115610fdf575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b038111828210171561214f5761214f612119565b60405290565b604080519081016001600160401b038111828210171561214f5761214f612119565b604051601f8201601f191681016001600160401b038111828210171561219f5761219f612119565b604052919050565b5f6001600160401b038211156121bf576121bf612119565b50601f01601f191660200190565b5f6121df6121da846121a7565b612177565b90508281528383830111156121f2575f5ffd5b828260208301375f602084830101529392505050565b5f5f5f5f5f6060868803121561221c575f5ffd5b85356001600160401b03811115612231575f5ffd5b61223d8882890161209c565b90965094505060208601356001600160401b0381111561225b575f5ffd5b612267888289016120d9565b90945092505060408601356001600160401b03811115612285575f5ffd5b8601601f81018813612295575f5ffd5b6122a4888235602084016121cd565b9150509295509295909350565b600681106122bd575f5ffd5b50565b5f5f5f5f5f606086880312156122d4575f5ffd5b85356001600160401b038111156122e9575f5ffd5b6122f5888289016120d9565b9096509450506020860135612309816122b1565b925060408601356001600160401b03811115612323575f5ffd5b61232f8882890161209c565b969995985093965092949392505050565b6001600160a01b03811681146122bd575f5ffd5b600481106122bd575f5ffd5b5f5f60408385031215612371575f5ffd5b823561237c81612340565b9150602083013561238c81612354565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b5f52602160045260245ffd5b600681106123e9576123e96123c5565b9052565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561248b57603f1987860301845281518051608087526124396080880182612397565b9050602082015161244d60208901826123d9565b50604082015187820360408901526124658282612397565b606093840151151598909301979097525094506020938401939190910190600101612413565b50929695505050505050565b5f5f5f604084860312156124a9575f5ffd5b83356001600160401b038111156124be575f5ffd5b6124ca868287016120d9565b909790965060209590950135949350505050565b602080825282518282018190525f918401906040840190835b8181101561251e5783516001600160a01b03168352602093840193909201916001016124f7565b509095945050505050565b5f60208284031215612539575f5ffd5b81356001600160401b0381111561254e575f5ffd5b8201601f8101841361255e575f5ffd5b611da5848235602084016121cd565b608081525f61257f6080830187612397565b82810360208401526125918187612397565b91505083604083015260ff8316606083015295945050505050565b5f5f602083850312156125bd575f5ffd5b82356001600160401b038111156125d2575f5ffd5b6125de858286016120d9565b90969095509350505050565b602081525f6108ea6020830184612397565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561248b57603f1987860301845261263e858351612397565b94506020938401939190910190600101612622565b5f5f60208385031215612664575f5ffd5b82356001600160401b03811115612679575f5ffd5b6125de8582860161209c565b602081525f825160a060208401526126a060c0840182612397565b90506020840151601f198483030160408501526126bd8282612397565b6040860151858203601f190160608701528051808352919350602090810192508084019190600582901b8501015f5b8281101561275157601f1986830301845284518051606084526127126060850182612397565b90506020820151848203602086015261272b8282612397565b6040938401511515959093019490945250602095860195949094019391506001016126ec565b50606088015160808801526080880151945061052260a088018660ff169052565b5f5f5f5f60408587031215612785575f5ffd5b84356001600160401b0381111561279a575f5ffd5b6127a6878288016120d9565b90955093505060208501356001600160401b038111156127c4575f5ffd5b6127d08782880161209c565b95989497509550505050565b6127e681846123d9565b604060208201525f611da56040830184612397565b600481106123e9576123e96123c5565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561248b57868503603f19018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101905f9060608801905b8083101561289c576128858285516127fb565b602082019150602084019350600183019250612872565b50965050506020938401939190910190600101612831565b80151581146122bd575f5ffd5b5f5f5f604084860312156128d3575f5ffd5b83356001600160401b038111156128e8575f5ffd5b6128f4868287016120d9565b9094509250506020840135612908816128b4565b809150509250925092565b5f5f5f5f5f5f5f6080888a031215612929575f5ffd5b87356001600160401b0381111561293e575f5ffd5b61294a8a828b0161209c565b90985096505060208801356001600160401b03811115612968575f5ffd5b6129748a828b0161209c565b9096509450506040880135925060608801356001600160401b03811115612999575f5ffd5b6129a58a828b016120d9565b989b979a50959850939692959293505050565b5f602082840312156129c8575f5ffd5b81356108ea81612340565b5f5f5f604084860312156129e5575f5ffd5b83356129f081612340565b925060208401356001600160401b03811115612a0a575f5ffd5b612a16868287016120d9565b9497909650939450505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f5f8335601e19843603018112612a60575f5ffd5b83016020810192503590506001600160401b03811115612a7e575f5ffd5b803603821315610fdf575f5ffd5b5f8383855260208501945060208460051b820101835f5b86811015612ae157838303601f19018852612abe8287612a4b565b612ac9858284612a23565b60209a8b019a90955093909301925050600101612aa3565b50909695505050505050565b868152608060208201525f612b06608083018789612a8c565b612b1360408401876123d9565b8281036060840152612b26818587612a23565b9998505050505050505050565b8381526001600160a01b038316602082015260608101611da560408301846127fb565b5f6001600160401b03821115612b6e57612b6e612119565b5060051b60200190565b5f82601f830112612b87575f5ffd5b8151602083015f612b9a6121da846121a7565b9050828152858383011115612bad575f5ffd5b8282602083015e5f92810160200192909252509392505050565b5f60208284031215612bd7575f5ffd5b81516001600160401b03811115612bec575f5ffd5b8201601f81018413612bfc575f5ffd5b8051612c0a6121da82612b56565b8082825260208201915060208360051b850101925086831115612c2b575f5ffd5b602084015b83811015612cfa5780516001600160401b03811115612c4d575f5ffd5b85016080818a03601f19011215612c62575f5ffd5b612c6a61212d565b60208201516001600160401b03811115612c82575f5ffd5b612c918b602083860101612b78565b8252506040820151612ca2816122b1565b602082015260608201516001600160401b03811115612cbf575f5ffd5b612cce8b602083860101612b78565b60408301525060808201519150612ce4826128b4565b6060810191909152835260209283019201612c30565b509695505050505050565b848152606060208201525f612d1e606083018587612a8c565b905082604083015295945050505050565b5f60208284031215612d3f575f5ffd5b81516001600160401b03811115612d54575f5ffd5b8201601f81018413612d64575f5ffd5b8051612d726121da82612b56565b8082825260208201915060208360051b850101925086831115612d93575f5ffd5b6020840193505b82841015612dbe578351612dad81612340565b825260209384019390910190612d9a565b9695505050505050565b600181811c90821680612ddc57607f821691505b602082108103612dfa57634e487b7160e01b5f52602260045260245ffd5b50919050565b838152604060208201525f612e19604083018486612a8c565b95945050505050565b5f60208284031215612e32575f5ffd5b5051919050565b5f60208284031215612e49575f5ffd5b81516001600160401b03811115612e5e575f5ffd5b8201601f81018413612e6e575f5ffd5b8051612e7c6121da82612b56565b8082825260208201915060208360051b850101925086831115612e9d575f5ffd5b602084015b83811015612cfa5780516001600160401b03811115612ebf575f5ffd5b612ece89602083890101612b78565b84525060209283019201612ea2565b838152604060208201525f612e19604083018486612a23565b818382375f9101908152919050565b5f81518060208401855e5f93019283525090919050565b5f6108ea8284612f05565b5f60208284031215612f37575f5ffd5b81516108ea816128b4565b858152606060208201525f612f5b606083018688612a8c565b8281036040840152612f6e818587612a23565b98975050505050505050565b5f5f60408385031215612f8b575f5ffd5b8251612f96816122b1565b60208401519092506001600160401b03811115612fb1575f5ffd5b612fbd85828601612b78565b9150509250929050565b83815260406020820181905281018290525f6060600584901b830181019083018583605e1936839003015b878210156130c157868503605f190184528235818112613010575f5ffd5b8901803536829003601e19018112613026575f5ffd5b81016020810190356001600160401b03811115613041575f5ffd5b8060051b3603821315613052575f5ffd5b60608852613064606089018284612a8c565b9150506020820135613075816122b1565b61308260208901826123d9565b506130906040830183612a4b565b925087820360408901526130a5828483612a23565b9750505050602083019250602084019350600182019150612ff2565b509298975050505050505050565b5f602082840312156130df575f5ffd5b81516001600160401b038111156130f4575f5ffd5b8201601f81018413613104575f5ffd5b80516131126121da82612b56565b8082825260208201915060208360051b850101925086831115613133575f5ffd5b602084015b83811015612cfa5780516001600160401b03811115613155575f5ffd5b85016040818a03601f1901121561316a575f5ffd5b613172612155565b602082015161318081612340565b815260408201516001600160401b0381111561319a575f5ffd5b60208184010192505089601f8301126131b1575f5ffd5b81516131bf6121da82612b56565b8082825260208201915060208360051b86010192508c8311156131e0575f5ffd5b6020850194505b8285101561320b5784516131fa81612354565b8252602094850194909101906131e7565b8060208501525050508085525050602083019250602081019050613138565b848152606060208201525f613243606083018587612a8c565b9050821515604083015295945050505050565b634e487b7160e01b5f52603260045260245ffd5b5f8235605e1983360301811261327e575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261329d575f5ffd5b8301803591506001600160401b038211156132b6575f5ffd5b602001915036819003821315610fdf575f5ffd5b5f5f8335601e198436030181126132df575f5ffd5b8301803591506001600160401b038211156132f8575f5ffd5b6020019150600581901b3603821315610fdf575f5ffd5b5f6020828403121561331f575f5ffd5b81356108ea816128b4565b634e487b7160e01b5f52601160045260245ffd5b5f60ff821660ff81036133535761335361332a565b60010192915050565b601f821115611c6457805f5260205f20601f840160051c810160208510156133815750805b601f840160051c820191505b818110156133a0575f815560010161338d565b5050505050565b6001600160401b038311156133be576133be612119565b6133d2836133cc8354612dc8565b8361335c565b5f601f841160018114613403575f85156133ec5750838201355b5f19600387901b1c1916600186901b1783556133a0565b5f83815260208120601f198716915b828110156134325786850135825560209485019460019092019101613412565b508682101561344e575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f813561346c816128b4565b92915050565b61347c8283613288565b6001600160401b0381111561349357613493612119565b6134a7816134a18554612dc8565b8561335c565b5f601f8211600181146134d8575f83156134c15750838201355b5f19600385901b1c1916600184901b17855561352f565b5f85815260208120601f198516915b8281101561350757868501358255602094850194600190920191016134e7565b5084821015613523575f1960f88660031b161c19848701351681555b505060018360011b0185555b505050506135406020830183613288565b61354e8183600186016133a7565b505061357861355f60408401613460565b6002830160ff1981541660ff8315151681178255505050565b5050565b602081525f611da5602083018486612a23565b5f6020828403121561359f575f5ffd5b81516001600160401b038111156135b4575f5ffd5b611da584828501612b78565b8481526001600160a01b038416602082015260606040820181905281018290525f8360808301825b858110156136195782356135fb81612354565b61360583826127fb565b5060209283019291909101906001016135e8565b50979650505050505050565b5f611da56136338386612f05565b84612f05565b5f6136448285612f05565b6001600160f81b03199390931683525050600101919050565b5f6001820161366e5761366e61332a565b506001019056fea2646970667358221220f8b15555cb7cdff5a8b4a6abc5fd5c705beaed31fd5f7dedb88877476213dfb764736f6c634300081c00333078646466323532616431626532633839623639633262303638666333373864616139353262613766313633633461313136323866353561346466353233623365663a333078646466323532616431626532633839623639633262303638666333373864616139353262613766313633633461313136323866353561346466353233623365663a32a2646970667358221220b83d5759f27fc4944e2c656f065793a9a7f158f7246bcb9941bcd833bbc450fc64736f6c634300081c0033",
}

// EventEmitterFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use EventEmitterFactoryMetaData.ABI instead.
var EventEmitterFactoryABI = EventEmitterFactoryMetaData.ABI

// EventEmitterFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventEmitterFactoryMetaData.Bin instead.
var EventEmitterFactoryBin = EventEmitterFactoryMetaData.Bin

// DeployEventEmitterFactory deploys a new Ethereum contract, binding an instance of EventEmitterFactory to it.
func DeployEventEmitterFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _validator common.Address) (common.Address, *types.Transaction, *EventEmitterFactory, error) {
	parsed, err := EventEmitterFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	roleControlAddr, _, _, _ := DeployRoleControl(auth, backend)
	EventEmitterFactoryBin = strings.ReplaceAll(EventEmitterFactoryBin, "__$5a92de94e3c11b56e4ee4f365f96da5afc$__", roleControlAddr.String()[2:])

	etherDatabaseLibAddr, _, _, _ := DeployEtherDatabaseLib(auth, backend)
	EventEmitterFactoryBin = strings.ReplaceAll(EventEmitterFactoryBin, "__$d451bf733f00e36b69819fe95d7d8eab5b$__", etherDatabaseLibAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventEmitterFactoryBin), backend, _validator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventEmitterFactory{EventEmitterFactoryCaller: EventEmitterFactoryCaller{contract: contract}, EventEmitterFactoryTransactor: EventEmitterFactoryTransactor{contract: contract}, EventEmitterFactoryFilterer: EventEmitterFactoryFilterer{contract: contract}}, nil
}

// EventEmitterFactory is an auto generated Go binding around an Ethereum contract.
type EventEmitterFactory struct {
	EventEmitterFactoryCaller     // Read-only binding to the contract
	EventEmitterFactoryTransactor // Write-only binding to the contract
	EventEmitterFactoryFilterer   // Log filterer for contract events
}

// EventEmitterFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventEmitterFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventEmitterFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventEmitterFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventEmitterFactorySession struct {
	Contract     *EventEmitterFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EventEmitterFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventEmitterFactoryCallerSession struct {
	Contract *EventEmitterFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EventEmitterFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventEmitterFactoryTransactorSession struct {
	Contract     *EventEmitterFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EventEmitterFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventEmitterFactoryRaw struct {
	Contract *EventEmitterFactory // Generic contract binding to access the raw methods on
}

// EventEmitterFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventEmitterFactoryCallerRaw struct {
	Contract *EventEmitterFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// EventEmitterFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventEmitterFactoryTransactorRaw struct {
	Contract *EventEmitterFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventEmitterFactory creates a new instance of EventEmitterFactory, bound to a specific deployed contract.
func NewEventEmitterFactory(address common.Address, backend bind.ContractBackend) (*EventEmitterFactory, error) {
	contract, err := bindEventEmitterFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactory{EventEmitterFactoryCaller: EventEmitterFactoryCaller{contract: contract}, EventEmitterFactoryTransactor: EventEmitterFactoryTransactor{contract: contract}, EventEmitterFactoryFilterer: EventEmitterFactoryFilterer{contract: contract}}, nil
}

// NewEventEmitterFactoryCaller creates a new read-only instance of EventEmitterFactory, bound to a specific deployed contract.
func NewEventEmitterFactoryCaller(address common.Address, caller bind.ContractCaller) (*EventEmitterFactoryCaller, error) {
	contract, err := bindEventEmitterFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactoryCaller{contract: contract}, nil
}

// NewEventEmitterFactoryTransactor creates a new write-only instance of EventEmitterFactory, bound to a specific deployed contract.
func NewEventEmitterFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*EventEmitterFactoryTransactor, error) {
	contract, err := bindEventEmitterFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactoryTransactor{contract: contract}, nil
}

// NewEventEmitterFactoryFilterer creates a new log filterer instance of EventEmitterFactory, bound to a specific deployed contract.
func NewEventEmitterFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*EventEmitterFactoryFilterer, error) {
	contract, err := bindEventEmitterFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactoryFilterer{contract: contract}, nil
}

// bindEventEmitterFactory binds a generic wrapper to an already deployed contract.
func bindEventEmitterFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventEmitterFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventEmitterFactory *EventEmitterFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventEmitterFactory.Contract.EventEmitterFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventEmitterFactory *EventEmitterFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.EventEmitterFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventEmitterFactory *EventEmitterFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.EventEmitterFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventEmitterFactory *EventEmitterFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventEmitterFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventEmitterFactory *EventEmitterFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventEmitterFactory *EventEmitterFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.contract.Transact(opts, method, params...)
}

// CustomContractAddresses is a free data retrieval call binding the contract method 0xabd60974.
//
// Solidity: function customContractAddresses(uint256 ) view returns(address)
func (_EventEmitterFactory *EventEmitterFactoryCaller) CustomContractAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "customContractAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomContractAddresses is a free data retrieval call binding the contract method 0xabd60974.
//
// Solidity: function customContractAddresses(uint256 ) view returns(address)
func (_EventEmitterFactory *EventEmitterFactorySession) CustomContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _EventEmitterFactory.Contract.CustomContractAddresses(&_EventEmitterFactory.CallOpts, arg0)
}

// CustomContractAddresses is a free data retrieval call binding the contract method 0xabd60974.
//
// Solidity: function customContractAddresses(uint256 ) view returns(address)
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) CustomContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _EventEmitterFactory.Contract.CustomContractAddresses(&_EventEmitterFactory.CallOpts, arg0)
}

// CustomContracts is a free data retrieval call binding the contract method 0x001f75d9.
//
// Solidity: function customContracts(address ) view returns(address contractAddress, string name, string contractABI)
func (_EventEmitterFactory *EventEmitterFactoryCaller) CustomContracts(opts *bind.CallOpts, arg0 common.Address) (struct {
	ContractAddress common.Address
	Name            string
	ContractABI     string
}, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "customContracts", arg0)

	outstruct := new(struct {
		ContractAddress common.Address
		Name            string
		ContractABI     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ContractABI = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// CustomContracts is a free data retrieval call binding the contract method 0x001f75d9.
//
// Solidity: function customContracts(address ) view returns(address contractAddress, string name, string contractABI)
func (_EventEmitterFactory *EventEmitterFactorySession) CustomContracts(arg0 common.Address) (struct {
	ContractAddress common.Address
	Name            string
	ContractABI     string
}, error) {
	return _EventEmitterFactory.Contract.CustomContracts(&_EventEmitterFactory.CallOpts, arg0)
}

// CustomContracts is a free data retrieval call binding the contract method 0x001f75d9.
//
// Solidity: function customContracts(address ) view returns(address contractAddress, string name, string contractABI)
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) CustomContracts(arg0 common.Address) (struct {
	ContractAddress common.Address
	Name            string
	ContractABI     string
}, error) {
	return _EventEmitterFactory.Contract.CustomContracts(&_EventEmitterFactory.CallOpts, arg0)
}

// Emitters is a free data retrieval call binding the contract method 0x5a0911c4.
//
// Solidity: function emitters(uint256 ) view returns(address emitterAddress, address owner)
func (_EventEmitterFactory *EventEmitterFactoryCaller) Emitters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EmitterAddress common.Address
	Owner          common.Address
}, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "emitters", arg0)

	outstruct := new(struct {
		EmitterAddress common.Address
		Owner          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmitterAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Emitters is a free data retrieval call binding the contract method 0x5a0911c4.
//
// Solidity: function emitters(uint256 ) view returns(address emitterAddress, address owner)
func (_EventEmitterFactory *EventEmitterFactorySession) Emitters(arg0 *big.Int) (struct {
	EmitterAddress common.Address
	Owner          common.Address
}, error) {
	return _EventEmitterFactory.Contract.Emitters(&_EventEmitterFactory.CallOpts, arg0)
}

// Emitters is a free data retrieval call binding the contract method 0x5a0911c4.
//
// Solidity: function emitters(uint256 ) view returns(address emitterAddress, address owner)
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) Emitters(arg0 *big.Int) (struct {
	EmitterAddress common.Address
	Owner          common.Address
}, error) {
	return _EventEmitterFactory.Contract.Emitters(&_EventEmitterFactory.CallOpts, arg0)
}

// GetCustomContractABI is a free data retrieval call binding the contract method 0xbd68dd10.
//
// Solidity: function getCustomContractABI(address contractAddress) view returns(string)
func (_EventEmitterFactory *EventEmitterFactoryCaller) GetCustomContractABI(opts *bind.CallOpts, contractAddress common.Address) (string, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "getCustomContractABI", contractAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCustomContractABI is a free data retrieval call binding the contract method 0xbd68dd10.
//
// Solidity: function getCustomContractABI(address contractAddress) view returns(string)
func (_EventEmitterFactory *EventEmitterFactorySession) GetCustomContractABI(contractAddress common.Address) (string, error) {
	return _EventEmitterFactory.Contract.GetCustomContractABI(&_EventEmitterFactory.CallOpts, contractAddress)
}

// GetCustomContractABI is a free data retrieval call binding the contract method 0xbd68dd10.
//
// Solidity: function getCustomContractABI(address contractAddress) view returns(string)
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) GetCustomContractABI(contractAddress common.Address) (string, error) {
	return _EventEmitterFactory.Contract.GetCustomContractABI(&_EventEmitterFactory.CallOpts, contractAddress)
}

// GetCustomContractEventNames is a free data retrieval call binding the contract method 0x1c32f662.
//
// Solidity: function getCustomContractEventNames(address contractAddress) view returns(string[])
func (_EventEmitterFactory *EventEmitterFactoryCaller) GetCustomContractEventNames(opts *bind.CallOpts, contractAddress common.Address) ([]string, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "getCustomContractEventNames", contractAddress)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCustomContractEventNames is a free data retrieval call binding the contract method 0x1c32f662.
//
// Solidity: function getCustomContractEventNames(address contractAddress) view returns(string[])
func (_EventEmitterFactory *EventEmitterFactorySession) GetCustomContractEventNames(contractAddress common.Address) ([]string, error) {
	return _EventEmitterFactory.Contract.GetCustomContractEventNames(&_EventEmitterFactory.CallOpts, contractAddress)
}

// GetCustomContractEventNames is a free data retrieval call binding the contract method 0x1c32f662.
//
// Solidity: function getCustomContractEventNames(address contractAddress) view returns(string[])
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) GetCustomContractEventNames(contractAddress common.Address) ([]string, error) {
	return _EventEmitterFactory.Contract.GetCustomContractEventNames(&_EventEmitterFactory.CallOpts, contractAddress)
}

// GetCustomContractSchema is a free data retrieval call binding the contract method 0xbfa33281.
//
// Solidity: function getCustomContractSchema(address contractAddress, string eventName) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitterFactory *EventEmitterFactoryCaller) GetCustomContractSchema(opts *bind.CallOpts, contractAddress common.Address, eventName string) (EventSchema, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "getCustomContractSchema", contractAddress, eventName)

	if err != nil {
		return *new(EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new(EventSchema)).(*EventSchema)

	return out0, err

}

// GetCustomContractSchema is a free data retrieval call binding the contract method 0xbfa33281.
//
// Solidity: function getCustomContractSchema(address contractAddress, string eventName) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitterFactory *EventEmitterFactorySession) GetCustomContractSchema(contractAddress common.Address, eventName string) (EventSchema, error) {
	return _EventEmitterFactory.Contract.GetCustomContractSchema(&_EventEmitterFactory.CallOpts, contractAddress, eventName)
}

// GetCustomContractSchema is a free data retrieval call binding the contract method 0xbfa33281.
//
// Solidity: function getCustomContractSchema(address contractAddress, string eventName) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) GetCustomContractSchema(contractAddress common.Address, eventName string) (EventSchema, error) {
	return _EventEmitterFactory.Contract.GetCustomContractSchema(&_EventEmitterFactory.CallOpts, contractAddress, eventName)
}

// GetCustomContracts is a free data retrieval call binding the contract method 0xbef863a6.
//
// Solidity: function getCustomContracts() view returns(address[])
func (_EventEmitterFactory *EventEmitterFactoryCaller) GetCustomContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "getCustomContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCustomContracts is a free data retrieval call binding the contract method 0xbef863a6.
//
// Solidity: function getCustomContracts() view returns(address[])
func (_EventEmitterFactory *EventEmitterFactorySession) GetCustomContracts() ([]common.Address, error) {
	return _EventEmitterFactory.Contract.GetCustomContracts(&_EventEmitterFactory.CallOpts)
}

// GetCustomContracts is a free data retrieval call binding the contract method 0xbef863a6.
//
// Solidity: function getCustomContracts() view returns(address[])
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) GetCustomContracts() ([]common.Address, error) {
	return _EventEmitterFactory.Contract.GetCustomContracts(&_EventEmitterFactory.CallOpts)
}

// GetEmitters is a free data retrieval call binding the contract method 0xe22c27ae.
//
// Solidity: function getEmitters() view returns((address,address)[])
func (_EventEmitterFactory *EventEmitterFactoryCaller) GetEmitters(opts *bind.CallOpts) ([]EventEmitterFactoryEmitterInfo, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "getEmitters")

	if err != nil {
		return *new([]EventEmitterFactoryEmitterInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventEmitterFactoryEmitterInfo)).(*[]EventEmitterFactoryEmitterInfo)

	return out0, err

}

// GetEmitters is a free data retrieval call binding the contract method 0xe22c27ae.
//
// Solidity: function getEmitters() view returns((address,address)[])
func (_EventEmitterFactory *EventEmitterFactorySession) GetEmitters() ([]EventEmitterFactoryEmitterInfo, error) {
	return _EventEmitterFactory.Contract.GetEmitters(&_EventEmitterFactory.CallOpts)
}

// GetEmitters is a free data retrieval call binding the contract method 0xe22c27ae.
//
// Solidity: function getEmitters() view returns((address,address)[])
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) GetEmitters() ([]EventEmitterFactoryEmitterInfo, error) {
	return _EventEmitterFactory.Contract.GetEmitters(&_EventEmitterFactory.CallOpts)
}

// GetGlobalSchemas is a free data retrieval call binding the contract method 0xa66ec868.
//
// Solidity: function getGlobalSchemas() pure returns((string,string,(string,string,bool)[],bytes32,uint8)[])
func (_EventEmitterFactory *EventEmitterFactoryCaller) GetGlobalSchemas(opts *bind.CallOpts) ([]EventSchema, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "getGlobalSchemas")

	if err != nil {
		return *new([]EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventSchema)).(*[]EventSchema)

	return out0, err

}

// GetGlobalSchemas is a free data retrieval call binding the contract method 0xa66ec868.
//
// Solidity: function getGlobalSchemas() pure returns((string,string,(string,string,bool)[],bytes32,uint8)[])
func (_EventEmitterFactory *EventEmitterFactorySession) GetGlobalSchemas() ([]EventSchema, error) {
	return _EventEmitterFactory.Contract.GetGlobalSchemas(&_EventEmitterFactory.CallOpts)
}

// GetGlobalSchemas is a free data retrieval call binding the contract method 0xa66ec868.
//
// Solidity: function getGlobalSchemas() pure returns((string,string,(string,string,bool)[],bytes32,uint8)[])
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) GetGlobalSchemas() ([]EventSchema, error) {
	return _EventEmitterFactory.Contract.GetGlobalSchemas(&_EventEmitterFactory.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EventEmitterFactory *EventEmitterFactoryCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventEmitterFactory.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EventEmitterFactory *EventEmitterFactorySession) Validator() (common.Address, error) {
	return _EventEmitterFactory.Contract.Validator(&_EventEmitterFactory.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EventEmitterFactory *EventEmitterFactoryCallerSession) Validator() (common.Address, error) {
	return _EventEmitterFactory.Contract.Validator(&_EventEmitterFactory.CallOpts)
}

// AddCustomContract is a paid mutator transaction binding the contract method 0x80bd201d.
//
// Solidity: function addCustomContract(address contractAddress, string name, string contractABI) returns()
func (_EventEmitterFactory *EventEmitterFactoryTransactor) AddCustomContract(opts *bind.TransactOpts, contractAddress common.Address, name string, contractABI string) (*types.Transaction, error) {
	return _EventEmitterFactory.contract.Transact(opts, "addCustomContract", contractAddress, name, contractABI)
}

// AddCustomContract is a paid mutator transaction binding the contract method 0x80bd201d.
//
// Solidity: function addCustomContract(address contractAddress, string name, string contractABI) returns()
func (_EventEmitterFactory *EventEmitterFactorySession) AddCustomContract(contractAddress common.Address, name string, contractABI string) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.AddCustomContract(&_EventEmitterFactory.TransactOpts, contractAddress, name, contractABI)
}

// AddCustomContract is a paid mutator transaction binding the contract method 0x80bd201d.
//
// Solidity: function addCustomContract(address contractAddress, string name, string contractABI) returns()
func (_EventEmitterFactory *EventEmitterFactoryTransactorSession) AddCustomContract(contractAddress common.Address, name string, contractABI string) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.AddCustomContract(&_EventEmitterFactory.TransactOpts, contractAddress, name, contractABI)
}

// AddCustomContractSchema is a paid mutator transaction binding the contract method 0xccfaf63c.
//
// Solidity: function addCustomContractSchema(address contractAddress, string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EventEmitterFactory *EventEmitterFactoryTransactor) AddCustomContractSchema(opts *bind.TransactOpts, contractAddress common.Address, name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EventEmitterFactory.contract.Transact(opts, "addCustomContractSchema", contractAddress, name, id, eventTopic, args)
}

// AddCustomContractSchema is a paid mutator transaction binding the contract method 0xccfaf63c.
//
// Solidity: function addCustomContractSchema(address contractAddress, string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EventEmitterFactory *EventEmitterFactorySession) AddCustomContractSchema(contractAddress common.Address, name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.AddCustomContractSchema(&_EventEmitterFactory.TransactOpts, contractAddress, name, id, eventTopic, args)
}

// AddCustomContractSchema is a paid mutator transaction binding the contract method 0xccfaf63c.
//
// Solidity: function addCustomContractSchema(address contractAddress, string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EventEmitterFactory *EventEmitterFactoryTransactorSession) AddCustomContractSchema(contractAddress common.Address, name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.AddCustomContractSchema(&_EventEmitterFactory.TransactOpts, contractAddress, name, id, eventTopic, args)
}

// CreateEmitter is a paid mutator transaction binding the contract method 0x176fa407.
//
// Solidity: function createEmitter() returns(address)
func (_EventEmitterFactory *EventEmitterFactoryTransactor) CreateEmitter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitterFactory.contract.Transact(opts, "createEmitter")
}

// CreateEmitter is a paid mutator transaction binding the contract method 0x176fa407.
//
// Solidity: function createEmitter() returns(address)
func (_EventEmitterFactory *EventEmitterFactorySession) CreateEmitter() (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.CreateEmitter(&_EventEmitterFactory.TransactOpts)
}

// CreateEmitter is a paid mutator transaction binding the contract method 0x176fa407.
//
// Solidity: function createEmitter() returns(address)
func (_EventEmitterFactory *EventEmitterFactoryTransactorSession) CreateEmitter() (*types.Transaction, error) {
	return _EventEmitterFactory.Contract.CreateEmitter(&_EventEmitterFactory.TransactOpts)
}

// EventEmitterFactoryCustomContractAddedIterator is returned from FilterCustomContractAdded and is used to iterate over the raw logs and unpacked data for CustomContractAdded events raised by the EventEmitterFactory contract.
type EventEmitterFactoryCustomContractAddedIterator struct {
	Event *EventEmitterFactoryCustomContractAdded // Event containing the contract specifics and raw log

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
func (it *EventEmitterFactoryCustomContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterFactoryCustomContractAdded)
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
		it.Event = new(EventEmitterFactoryCustomContractAdded)
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
func (it *EventEmitterFactoryCustomContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterFactoryCustomContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterFactoryCustomContractAdded represents a CustomContractAdded event raised by the EventEmitterFactory contract.
type EventEmitterFactoryCustomContractAdded struct {
	ContractAddress common.Address
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCustomContractAdded is a free log retrieval operation binding the contract event 0x994774dc478a0a24cbad071ebe255a31e5868d3bce654611faa96b645a940f60.
//
// Solidity: event CustomContractAdded(address indexed contractAddress, string name)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) FilterCustomContractAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*EventEmitterFactoryCustomContractAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _EventEmitterFactory.contract.FilterLogs(opts, "CustomContractAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactoryCustomContractAddedIterator{contract: _EventEmitterFactory.contract, event: "CustomContractAdded", logs: logs, sub: sub}, nil
}

// WatchCustomContractAdded is a free log subscription operation binding the contract event 0x994774dc478a0a24cbad071ebe255a31e5868d3bce654611faa96b645a940f60.
//
// Solidity: event CustomContractAdded(address indexed contractAddress, string name)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) WatchCustomContractAdded(opts *bind.WatchOpts, sink chan<- *EventEmitterFactoryCustomContractAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _EventEmitterFactory.contract.WatchLogs(opts, "CustomContractAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterFactoryCustomContractAdded)
				if err := _EventEmitterFactory.contract.UnpackLog(event, "CustomContractAdded", log); err != nil {
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

// ParseCustomContractAdded is a log parse operation binding the contract event 0x994774dc478a0a24cbad071ebe255a31e5868d3bce654611faa96b645a940f60.
//
// Solidity: event CustomContractAdded(address indexed contractAddress, string name)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) ParseCustomContractAdded(log types.Log) (*EventEmitterFactoryCustomContractAdded, error) {
	event := new(EventEmitterFactoryCustomContractAdded)
	if err := _EventEmitterFactory.contract.UnpackLog(event, "CustomContractAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterFactoryCustomContractSchemaAddedIterator is returned from FilterCustomContractSchemaAdded and is used to iterate over the raw logs and unpacked data for CustomContractSchemaAdded events raised by the EventEmitterFactory contract.
type EventEmitterFactoryCustomContractSchemaAddedIterator struct {
	Event *EventEmitterFactoryCustomContractSchemaAdded // Event containing the contract specifics and raw log

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
func (it *EventEmitterFactoryCustomContractSchemaAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterFactoryCustomContractSchemaAdded)
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
		it.Event = new(EventEmitterFactoryCustomContractSchemaAdded)
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
func (it *EventEmitterFactoryCustomContractSchemaAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterFactoryCustomContractSchemaAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterFactoryCustomContractSchemaAdded represents a CustomContractSchemaAdded event raised by the EventEmitterFactory contract.
type EventEmitterFactoryCustomContractSchemaAdded struct {
	ContractAddress common.Address
	EventName       string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCustomContractSchemaAdded is a free log retrieval operation binding the contract event 0x8c7d8ad02950ea5bd57f54af9c47ca84706c82e280cada0c04a2da9f2bb4818c.
//
// Solidity: event CustomContractSchemaAdded(address indexed contractAddress, string eventName)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) FilterCustomContractSchemaAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*EventEmitterFactoryCustomContractSchemaAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _EventEmitterFactory.contract.FilterLogs(opts, "CustomContractSchemaAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactoryCustomContractSchemaAddedIterator{contract: _EventEmitterFactory.contract, event: "CustomContractSchemaAdded", logs: logs, sub: sub}, nil
}

// WatchCustomContractSchemaAdded is a free log subscription operation binding the contract event 0x8c7d8ad02950ea5bd57f54af9c47ca84706c82e280cada0c04a2da9f2bb4818c.
//
// Solidity: event CustomContractSchemaAdded(address indexed contractAddress, string eventName)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) WatchCustomContractSchemaAdded(opts *bind.WatchOpts, sink chan<- *EventEmitterFactoryCustomContractSchemaAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _EventEmitterFactory.contract.WatchLogs(opts, "CustomContractSchemaAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterFactoryCustomContractSchemaAdded)
				if err := _EventEmitterFactory.contract.UnpackLog(event, "CustomContractSchemaAdded", log); err != nil {
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

// ParseCustomContractSchemaAdded is a log parse operation binding the contract event 0x8c7d8ad02950ea5bd57f54af9c47ca84706c82e280cada0c04a2da9f2bb4818c.
//
// Solidity: event CustomContractSchemaAdded(address indexed contractAddress, string eventName)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) ParseCustomContractSchemaAdded(log types.Log) (*EventEmitterFactoryCustomContractSchemaAdded, error) {
	event := new(EventEmitterFactoryCustomContractSchemaAdded)
	if err := _EventEmitterFactory.contract.UnpackLog(event, "CustomContractSchemaAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterFactoryEmitterCreatedIterator is returned from FilterEmitterCreated and is used to iterate over the raw logs and unpacked data for EmitterCreated events raised by the EventEmitterFactory contract.
type EventEmitterFactoryEmitterCreatedIterator struct {
	Event *EventEmitterFactoryEmitterCreated // Event containing the contract specifics and raw log

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
func (it *EventEmitterFactoryEmitterCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterFactoryEmitterCreated)
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
		it.Event = new(EventEmitterFactoryEmitterCreated)
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
func (it *EventEmitterFactoryEmitterCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterFactoryEmitterCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterFactoryEmitterCreated represents a EmitterCreated event raised by the EventEmitterFactory contract.
type EventEmitterFactoryEmitterCreated struct {
	EmitterAddress common.Address
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEmitterCreated is a free log retrieval operation binding the contract event 0xd69887947906502e8670345f1af75cd1bc3be116978564188dda62fad0aed1ba.
//
// Solidity: event EmitterCreated(address indexed emitterAddress, address indexed owner)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) FilterEmitterCreated(opts *bind.FilterOpts, emitterAddress []common.Address, owner []common.Address) (*EventEmitterFactoryEmitterCreatedIterator, error) {

	var emitterAddressRule []interface{}
	for _, emitterAddressItem := range emitterAddress {
		emitterAddressRule = append(emitterAddressRule, emitterAddressItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EventEmitterFactory.contract.FilterLogs(opts, "EmitterCreated", emitterAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFactoryEmitterCreatedIterator{contract: _EventEmitterFactory.contract, event: "EmitterCreated", logs: logs, sub: sub}, nil
}

// WatchEmitterCreated is a free log subscription operation binding the contract event 0xd69887947906502e8670345f1af75cd1bc3be116978564188dda62fad0aed1ba.
//
// Solidity: event EmitterCreated(address indexed emitterAddress, address indexed owner)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) WatchEmitterCreated(opts *bind.WatchOpts, sink chan<- *EventEmitterFactoryEmitterCreated, emitterAddress []common.Address, owner []common.Address) (event.Subscription, error) {

	var emitterAddressRule []interface{}
	for _, emitterAddressItem := range emitterAddress {
		emitterAddressRule = append(emitterAddressRule, emitterAddressItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EventEmitterFactory.contract.WatchLogs(opts, "EmitterCreated", emitterAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterFactoryEmitterCreated)
				if err := _EventEmitterFactory.contract.UnpackLog(event, "EmitterCreated", log); err != nil {
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

// ParseEmitterCreated is a log parse operation binding the contract event 0xd69887947906502e8670345f1af75cd1bc3be116978564188dda62fad0aed1ba.
//
// Solidity: event EmitterCreated(address indexed emitterAddress, address indexed owner)
func (_EventEmitterFactory *EventEmitterFactoryFilterer) ParseEmitterCreated(log types.Log) (*EventEmitterFactoryEmitterCreated, error) {
	event := new(EventEmitterFactoryEmitterCreated)
	if err := _EventEmitterFactory.contract.UnpackLog(event, "EmitterCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthDBUpdaterMetaData contains all meta data concerning the IEthDBUpdater contract.
var IEthDBUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"}]",
}

// IEthDBUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthDBUpdaterMetaData.ABI instead.
var IEthDBUpdaterABI = IEthDBUpdaterMetaData.ABI

// IEthDBUpdater is an auto generated Go binding around an Ethereum contract.
type IEthDBUpdater struct {
	IEthDBUpdaterCaller     // Read-only binding to the contract
	IEthDBUpdaterTransactor // Write-only binding to the contract
	IEthDBUpdaterFilterer   // Log filterer for contract events
}

// IEthDBUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthDBUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthDBUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthDBUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthDBUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthDBUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthDBUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthDBUpdaterSession struct {
	Contract     *IEthDBUpdater    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEthDBUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthDBUpdaterCallerSession struct {
	Contract *IEthDBUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IEthDBUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthDBUpdaterTransactorSession struct {
	Contract     *IEthDBUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IEthDBUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthDBUpdaterRaw struct {
	Contract *IEthDBUpdater // Generic contract binding to access the raw methods on
}

// IEthDBUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthDBUpdaterCallerRaw struct {
	Contract *IEthDBUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// IEthDBUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthDBUpdaterTransactorRaw struct {
	Contract *IEthDBUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthDBUpdater creates a new instance of IEthDBUpdater, bound to a specific deployed contract.
func NewIEthDBUpdater(address common.Address, backend bind.ContractBackend) (*IEthDBUpdater, error) {
	contract, err := bindIEthDBUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthDBUpdater{IEthDBUpdaterCaller: IEthDBUpdaterCaller{contract: contract}, IEthDBUpdaterTransactor: IEthDBUpdaterTransactor{contract: contract}, IEthDBUpdaterFilterer: IEthDBUpdaterFilterer{contract: contract}}, nil
}

// NewIEthDBUpdaterCaller creates a new read-only instance of IEthDBUpdater, bound to a specific deployed contract.
func NewIEthDBUpdaterCaller(address common.Address, caller bind.ContractCaller) (*IEthDBUpdaterCaller, error) {
	contract, err := bindIEthDBUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthDBUpdaterCaller{contract: contract}, nil
}

// NewIEthDBUpdaterTransactor creates a new write-only instance of IEthDBUpdater, bound to a specific deployed contract.
func NewIEthDBUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthDBUpdaterTransactor, error) {
	contract, err := bindIEthDBUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthDBUpdaterTransactor{contract: contract}, nil
}

// NewIEthDBUpdaterFilterer creates a new log filterer instance of IEthDBUpdater, bound to a specific deployed contract.
func NewIEthDBUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthDBUpdaterFilterer, error) {
	contract, err := bindIEthDBUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthDBUpdaterFilterer{contract: contract}, nil
}

// bindIEthDBUpdater binds a generic wrapper to an already deployed contract.
func bindIEthDBUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEthDBUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthDBUpdater *IEthDBUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthDBUpdater.Contract.IEthDBUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthDBUpdater *IEthDBUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthDBUpdater.Contract.IEthDBUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthDBUpdater *IEthDBUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthDBUpdater.Contract.IEthDBUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthDBUpdater *IEthDBUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthDBUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthDBUpdater *IEthDBUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthDBUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthDBUpdater *IEthDBUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthDBUpdater.Contract.contract.Transact(opts, method, params...)
}

// IEthDBUpdaterEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the IEthDBUpdater contract.
type IEthDBUpdaterEthDBPathUpdateIterator struct {
	Event *IEthDBUpdaterEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *IEthDBUpdaterEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthDBUpdaterEthDBPathUpdate)
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
		it.Event = new(IEthDBUpdaterEthDBPathUpdate)
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
func (it *IEthDBUpdaterEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthDBUpdaterEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthDBUpdaterEthDBPathUpdate represents a EthDBPathUpdate event raised by the IEthDBUpdater contract.
type IEthDBUpdaterEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_IEthDBUpdater *IEthDBUpdaterFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*IEthDBUpdaterEthDBPathUpdateIterator, error) {

	logs, sub, err := _IEthDBUpdater.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &IEthDBUpdaterEthDBPathUpdateIterator{contract: _IEthDBUpdater.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_IEthDBUpdater *IEthDBUpdaterFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *IEthDBUpdaterEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _IEthDBUpdater.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthDBUpdaterEthDBPathUpdate)
				if err := _IEthDBUpdater.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_IEthDBUpdater *IEthDBUpdaterFilterer) ParseEthDBPathUpdate(log types.Log) (*IEthDBUpdaterEthDBPathUpdate, error) {
	event := new(IEthDBUpdaterEthDBPathUpdate)
	if err := _IEthDBUpdater.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthDBUpdaterEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the IEthDBUpdater contract.
type IEthDBUpdaterEthDBUpdateIterator struct {
	Event *IEthDBUpdaterEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *IEthDBUpdaterEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthDBUpdaterEthDBUpdate)
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
		it.Event = new(IEthDBUpdaterEthDBUpdate)
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
func (it *IEthDBUpdaterEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthDBUpdaterEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthDBUpdaterEthDBUpdate represents a EthDBUpdate event raised by the IEthDBUpdater contract.
type IEthDBUpdaterEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_IEthDBUpdater *IEthDBUpdaterFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*IEthDBUpdaterEthDBUpdateIterator, error) {

	logs, sub, err := _IEthDBUpdater.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &IEthDBUpdaterEthDBUpdateIterator{contract: _IEthDBUpdater.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_IEthDBUpdater *IEthDBUpdaterFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *IEthDBUpdaterEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _IEthDBUpdater.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthDBUpdaterEthDBUpdate)
				if err := _IEthDBUpdater.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_IEthDBUpdater *IEthDBUpdaterFilterer) ParseEthDBUpdate(log types.Log) (*IEthDBUpdaterEthDBUpdate, error) {
	event := new(IEthDBUpdaterEthDBUpdate)
	if err := _IEthDBUpdater.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Multicall3MetaData contains all meta data concerning the Multicall3 contract.
var Multicall3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610c6a8061001c5f395ff3fe6080604052600436106100ef575f3560e01c80634d2301cc11610087578063a8b0574e11610057578063a8b0574e14610221578063bce38bd71461023b578063c3077fa91461024e578063ee82ac5e14610261575f5ffd5b80634d2301cc146101c357806372425d9d146101ea57806382ad56cb146101fc57806386d516e81461020f575f5ffd5b80633408e470116100c25780633408e4701461016b578063399542e91461017d5780633e64a6961461019f57806342cbb15c146101b1575f5ffd5b80630f28c97d146100f3578063174dea7114610114578063252dba421461013457806327e86d6e14610155575b5f5ffd5b3480156100fe575f5ffd5b50425b6040519081526020015b60405180910390f35b610127610122366004610958565b61027f565b60405161010b9190610a38565b610147610142366004610958565b610464565b60405161010b929190610a51565b348015610160575f5ffd5b50435f190140610101565b348015610176575f5ffd5b5046610101565b61019061018b366004610abb565b6105d2565b60405161010b93929190610b10565b3480156101aa575f5ffd5b5048610101565b3480156101bc575f5ffd5b5043610101565b3480156101ce575f5ffd5b506101016101dd366004610b37565b6001600160a01b03163190565b3480156101f5575f5ffd5b5044610101565b61012761020a366004610958565b6105ed565b34801561021a575f5ffd5b5045610101565b34801561022c575f5ffd5b5060405141815260200161010b565b610127610249366004610abb565b610766565b61019061025c366004610958565b6108f2565b34801561026c575f5ffd5b5061010161027b366004610b5d565b4090565b60605f828067ffffffffffffffff81111561029c5761029c610b74565b6040519080825280602002602001820160405280156102e157816020015b604080518082019091525f8152606060208201528152602001906001900390816102ba5790505b509250365f5b82811015610406575f85828151811061030257610302610b88565b6020026020010151905087878381811061031e5761031e610b88565b90506020028101906103309190610b9c565b6040810135958601959093506103496020850185610b37565b6001600160a01b0316816103606060870187610bba565b60405161036e929190610bfd565b5f6040518083038185875af1925050503d805f81146103a8576040519150601f19603f3d011682016040523d82523d5f602084013e6103ad565b606091505b5060208085019190915290151580845290850135176103fc5762461bcd60e51b5f526020600452601760245276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b60445260845ffd5b50506001016102e7565b5082341461045b5760405162461bcd60e51b815260206004820152601a60248201527f4d756c746963616c6c333a2076616c7565206d69736d6174636800000000000060448201526064015b60405180910390fd5b50505092915050565b436060828067ffffffffffffffff81111561048157610481610b74565b6040519080825280602002602001820160405280156104b457816020015b606081526020019060019003908161049f5790505b509150365f5b828110156105c8575f8787838181106104d5576104d5610b88565b90506020028101906104e79190610c0c565b92506104f66020840184610b37565b6001600160a01b031661050c6020850185610bba565b60405161051a929190610bfd565b5f604051808303815f865af19150503d805f8114610553576040519150601f19603f3d011682016040523d82523d5f602084013e610558565b606091505b5086848151811061056b5761056b610b88565b60209081029190910101529050806105bf5760405162461bcd60e51b8152602060048201526017602482015276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b6044820152606401610452565b506001016104ba565b5050509250929050565b43804060606105e2868686610766565b905093509350939050565b6060818067ffffffffffffffff81111561060957610609610b74565b60405190808252806020026020018201604052801561064e57816020015b604080518082019091525f8152606060208201528152602001906001900390816106275790505b509150365f5b8281101561045b575f84828151811061066f5761066f610b88565b6020026020010151905086868381811061068b5761068b610b88565b905060200281019061069d9190610c20565b92506106ac6020840184610b37565b6001600160a01b03166106c26040850185610bba565b6040516106d0929190610bfd565b5f604051808303815f865af19150503d805f8114610709576040519150601f19603f3d011682016040523d82523d5f602084013e61070e565b606091505b50602080840191909152901515808352908401351761075d5762461bcd60e51b5f526020600452601760245276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b60445260645ffd5b50600101610654565b6060818067ffffffffffffffff81111561078257610782610b74565b6040519080825280602002602001820160405280156107c757816020015b604080518082019091525f8152606060208201528152602001906001900390816107a05790505b509150365f5b828110156108e8575f8482815181106107e8576107e8610b88565b6020026020010151905086868381811061080457610804610b88565b90506020028101906108169190610c0c565b92506108256020840184610b37565b6001600160a01b031661083b6020850185610bba565b604051610849929190610bfd565b5f604051808303815f865af19150503d805f8114610882576040519150601f19603f3d011682016040523d82523d5f602084013e610887565b606091505b5060208301521515815287156108df5780516108df5760405162461bcd60e51b8152602060048201526017602482015276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b6044820152606401610452565b506001016107cd565b5050509392505050565b5f5f6060610902600186866105d2565b919790965090945092505050565b5f5f83601f840112610920575f5ffd5b50813567ffffffffffffffff811115610937575f5ffd5b6020830191508360208260051b8501011115610951575f5ffd5b9250929050565b5f5f60208385031215610969575f5ffd5b823567ffffffffffffffff81111561097f575f5ffd5b61098b85828601610910565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f82825180855260208501945060208160051b830101602085015f5b83811015610a2c57601f1985840301885281518051151584526020810151905060406020850152610a156040850182610997565b6020998a01999094509290920191506001016109e1565b50909695505050505050565b602081525f610a4a60208301846109c5565b9392505050565b5f604082018483526040602084015280845180835260608501915060608160051b8601019250602086015f5b82811015610aae57605f19878603018452610a99858351610997565b94506020938401939190910190600101610a7d565b5092979650505050505050565b5f5f5f60408486031215610acd575f5ffd5b83358015158114610adc575f5ffd5b9250602084013567ffffffffffffffff811115610af7575f5ffd5b610b0386828701610910565b9497909650939450505050565b838152826020820152606060408201525f610b2e60608301846109c5565b95945050505050565b5f60208284031215610b47575f5ffd5b81356001600160a01b0381168114610a4a575f5ffd5b5f60208284031215610b6d575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235607e19833603018112610bb0575f5ffd5b9190910192915050565b5f5f8335601e19843603018112610bcf575f5ffd5b83018035915067ffffffffffffffff821115610be9575f5ffd5b602001915036819003821315610951575f5ffd5b818382375f9101908152919050565b5f8235603e19833603018112610bb0575f5ffd5b5f8235605e19833603018112610bb0575f5ffdfea264697066735822122091a702ca43a807ae162c5c9ca466e3b922a32e2af7a5eec8b350da05c729681764736f6c634300081c0033",
}

// Multicall3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Multicall3MetaData.ABI instead.
var Multicall3ABI = Multicall3MetaData.ABI

// Multicall3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Multicall3MetaData.Bin instead.
var Multicall3Bin = Multicall3MetaData.Bin

// DeployMulticall3 deploys a new Ethereum contract, binding an instance of Multicall3 to it.
func DeployMulticall3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multicall3, error) {
	parsed, err := Multicall3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Multicall3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multicall3{Multicall3Caller: Multicall3Caller{contract: contract}, Multicall3Transactor: Multicall3Transactor{contract: contract}, Multicall3Filterer: Multicall3Filterer{contract: contract}}, nil
}

// Multicall3 is an auto generated Go binding around an Ethereum contract.
type Multicall3 struct {
	Multicall3Caller     // Read-only binding to the contract
	Multicall3Transactor // Write-only binding to the contract
	Multicall3Filterer   // Log filterer for contract events
}

// Multicall3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Multicall3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Multicall3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Multicall3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Multicall3Session struct {
	Contract     *Multicall3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Multicall3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Multicall3CallerSession struct {
	Contract *Multicall3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Multicall3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Multicall3TransactorSession struct {
	Contract     *Multicall3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Multicall3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Multicall3Raw struct {
	Contract *Multicall3 // Generic contract binding to access the raw methods on
}

// Multicall3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Multicall3CallerRaw struct {
	Contract *Multicall3Caller // Generic read-only contract binding to access the raw methods on
}

// Multicall3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Multicall3TransactorRaw struct {
	Contract *Multicall3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall3 creates a new instance of Multicall3, bound to a specific deployed contract.
func NewMulticall3(address common.Address, backend bind.ContractBackend) (*Multicall3, error) {
	contract, err := bindMulticall3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall3{Multicall3Caller: Multicall3Caller{contract: contract}, Multicall3Transactor: Multicall3Transactor{contract: contract}, Multicall3Filterer: Multicall3Filterer{contract: contract}}, nil
}

// NewMulticall3Caller creates a new read-only instance of Multicall3, bound to a specific deployed contract.
func NewMulticall3Caller(address common.Address, caller bind.ContractCaller) (*Multicall3Caller, error) {
	contract, err := bindMulticall3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Multicall3Caller{contract: contract}, nil
}

// NewMulticall3Transactor creates a new write-only instance of Multicall3, bound to a specific deployed contract.
func NewMulticall3Transactor(address common.Address, transactor bind.ContractTransactor) (*Multicall3Transactor, error) {
	contract, err := bindMulticall3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Multicall3Transactor{contract: contract}, nil
}

// NewMulticall3Filterer creates a new log filterer instance of Multicall3, bound to a specific deployed contract.
func NewMulticall3Filterer(address common.Address, filterer bind.ContractFilterer) (*Multicall3Filterer, error) {
	contract, err := bindMulticall3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Multicall3Filterer{contract: contract}, nil
}

// bindMulticall3 binds a generic wrapper to an already deployed contract.
func bindMulticall3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Multicall3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall3 *Multicall3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall3.Contract.Multicall3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall3 *Multicall3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall3.Contract.Multicall3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall3 *Multicall3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall3.Contract.Multicall3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall3 *Multicall3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall3 *Multicall3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall3 *Multicall3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall3.Contract.contract.Transact(opts, method, params...)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_Multicall3 *Multicall3Caller) GetBasefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getBasefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_Multicall3 *Multicall3Session) GetBasefee() (*big.Int, error) {
	return _Multicall3.Contract.GetBasefee(&_Multicall3.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_Multicall3 *Multicall3CallerSession) GetBasefee() (*big.Int, error) {
	return _Multicall3.Contract.GetBasefee(&_Multicall3.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall3 *Multicall3Caller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall3 *Multicall3Session) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicall3.Contract.GetBlockHash(&_Multicall3.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall3 *Multicall3CallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicall3.Contract.GetBlockHash(&_Multicall3.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall3 *Multicall3Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall3 *Multicall3Session) GetBlockNumber() (*big.Int, error) {
	return _Multicall3.Contract.GetBlockNumber(&_Multicall3.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall3 *Multicall3CallerSession) GetBlockNumber() (*big.Int, error) {
	return _Multicall3.Contract.GetBlockNumber(&_Multicall3.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_Multicall3 *Multicall3Caller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_Multicall3 *Multicall3Session) GetChainId() (*big.Int, error) {
	return _Multicall3.Contract.GetChainId(&_Multicall3.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_Multicall3 *Multicall3CallerSession) GetChainId() (*big.Int, error) {
	return _Multicall3.Contract.GetChainId(&_Multicall3.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall3 *Multicall3Caller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall3 *Multicall3Session) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicall3.Contract.GetCurrentBlockCoinbase(&_Multicall3.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall3 *Multicall3CallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicall3.Contract.GetCurrentBlockCoinbase(&_Multicall3.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall3 *Multicall3Caller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall3 *Multicall3Session) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicall3.Contract.GetCurrentBlockDifficulty(&_Multicall3.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall3 *Multicall3CallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicall3.Contract.GetCurrentBlockDifficulty(&_Multicall3.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall3 *Multicall3Caller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall3 *Multicall3Session) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicall3.Contract.GetCurrentBlockGasLimit(&_Multicall3.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall3 *Multicall3CallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicall3.Contract.GetCurrentBlockGasLimit(&_Multicall3.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall3 *Multicall3Caller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall3 *Multicall3Session) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicall3.Contract.GetCurrentBlockTimestamp(&_Multicall3.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall3 *Multicall3CallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicall3.Contract.GetCurrentBlockTimestamp(&_Multicall3.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall3 *Multicall3Caller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall3 *Multicall3Session) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicall3.Contract.GetEthBalance(&_Multicall3.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall3 *Multicall3CallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicall3.Contract.GetEthBalance(&_Multicall3.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall3 *Multicall3Caller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multicall3.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall3 *Multicall3Session) GetLastBlockHash() ([32]byte, error) {
	return _Multicall3.Contract.GetLastBlockHash(&_Multicall3.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall3 *Multicall3CallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicall3.Contract.GetLastBlockHash(&_Multicall3.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall3 *Multicall3Transactor) Aggregate(opts *bind.TransactOpts, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall3 *Multicall3Session) Aggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.Aggregate(&_Multicall3.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall3 *Multicall3TransactorSession) Aggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.Aggregate(&_Multicall3.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Transactor) Aggregate3(opts *bind.TransactOpts, calls []Multicall3Call3) (*types.Transaction, error) {
	return _Multicall3.contract.Transact(opts, "aggregate3", calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Session) Aggregate3(calls []Multicall3Call3) (*types.Transaction, error) {
	return _Multicall3.Contract.Aggregate3(&_Multicall3.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3TransactorSession) Aggregate3(calls []Multicall3Call3) (*types.Transaction, error) {
	return _Multicall3.Contract.Aggregate3(&_Multicall3.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Transactor) Aggregate3Value(opts *bind.TransactOpts, calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _Multicall3.contract.Transact(opts, "aggregate3Value", calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Session) Aggregate3Value(calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _Multicall3.Contract.Aggregate3Value(&_Multicall3.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3TransactorSession) Aggregate3Value(calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _Multicall3.Contract.Aggregate3Value(&_Multicall3.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Transactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Session) BlockAndAggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.BlockAndAggregate(&_Multicall3.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall3 *Multicall3TransactorSession) BlockAndAggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.BlockAndAggregate(&_Multicall3.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Transactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Session) TryAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.TryAggregate(&_Multicall3.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Multicall3 *Multicall3TransactorSession) TryAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.TryAggregate(&_Multicall3.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Transactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall3 *Multicall3Session) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.TryBlockAndAggregate(&_Multicall3.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall3 *Multicall3TransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _Multicall3.Contract.TryBlockAndAggregate(&_Multicall3.TransactOpts, requireSuccess, calls)
}

// PrivateVariableCustomStateEventMetaData contains all meta data concerning the PrivateVariableCustomStateEvent contract.
var PrivateVariableCustomStateEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060f58061001b5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c806320965255146034578063d09de08a146048575b5f5ffd5b5f5460405190815260200160405180910390f35b604e6050565b005b60015f5f828254605f9190609b565b90915550505f546040519081527f93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c599060200160405180910390a1565b8082018082111560b957634e487b7160e01b5f52601160045260245ffd5b9291505056fea264697066735822122037cc8191c9a76e135a48ce87c58fe30ca76609131b90c30c09485e3ae0d99c8264736f6c634300081c0033",
}

// PrivateVariableCustomStateEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PrivateVariableCustomStateEventMetaData.ABI instead.
var PrivateVariableCustomStateEventABI = PrivateVariableCustomStateEventMetaData.ABI

// PrivateVariableCustomStateEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrivateVariableCustomStateEventMetaData.Bin instead.
var PrivateVariableCustomStateEventBin = PrivateVariableCustomStateEventMetaData.Bin

// DeployPrivateVariableCustomStateEvent deploys a new Ethereum contract, binding an instance of PrivateVariableCustomStateEvent to it.
func DeployPrivateVariableCustomStateEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrivateVariableCustomStateEvent, error) {
	parsed, err := PrivateVariableCustomStateEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrivateVariableCustomStateEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivateVariableCustomStateEvent{PrivateVariableCustomStateEventCaller: PrivateVariableCustomStateEventCaller{contract: contract}, PrivateVariableCustomStateEventTransactor: PrivateVariableCustomStateEventTransactor{contract: contract}, PrivateVariableCustomStateEventFilterer: PrivateVariableCustomStateEventFilterer{contract: contract}}, nil
}

// PrivateVariableCustomStateEvent is an auto generated Go binding around an Ethereum contract.
type PrivateVariableCustomStateEvent struct {
	PrivateVariableCustomStateEventCaller     // Read-only binding to the contract
	PrivateVariableCustomStateEventTransactor // Write-only binding to the contract
	PrivateVariableCustomStateEventFilterer   // Log filterer for contract events
}

// PrivateVariableCustomStateEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivateVariableCustomStateEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateVariableCustomStateEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivateVariableCustomStateEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateVariableCustomStateEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivateVariableCustomStateEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateVariableCustomStateEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivateVariableCustomStateEventSession struct {
	Contract     *PrivateVariableCustomStateEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PrivateVariableCustomStateEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivateVariableCustomStateEventCallerSession struct {
	Contract *PrivateVariableCustomStateEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PrivateVariableCustomStateEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivateVariableCustomStateEventTransactorSession struct {
	Contract     *PrivateVariableCustomStateEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PrivateVariableCustomStateEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivateVariableCustomStateEventRaw struct {
	Contract *PrivateVariableCustomStateEvent // Generic contract binding to access the raw methods on
}

// PrivateVariableCustomStateEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivateVariableCustomStateEventCallerRaw struct {
	Contract *PrivateVariableCustomStateEventCaller // Generic read-only contract binding to access the raw methods on
}

// PrivateVariableCustomStateEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivateVariableCustomStateEventTransactorRaw struct {
	Contract *PrivateVariableCustomStateEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivateVariableCustomStateEvent creates a new instance of PrivateVariableCustomStateEvent, bound to a specific deployed contract.
func NewPrivateVariableCustomStateEvent(address common.Address, backend bind.ContractBackend) (*PrivateVariableCustomStateEvent, error) {
	contract, err := bindPrivateVariableCustomStateEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivateVariableCustomStateEvent{PrivateVariableCustomStateEventCaller: PrivateVariableCustomStateEventCaller{contract: contract}, PrivateVariableCustomStateEventTransactor: PrivateVariableCustomStateEventTransactor{contract: contract}, PrivateVariableCustomStateEventFilterer: PrivateVariableCustomStateEventFilterer{contract: contract}}, nil
}

// NewPrivateVariableCustomStateEventCaller creates a new read-only instance of PrivateVariableCustomStateEvent, bound to a specific deployed contract.
func NewPrivateVariableCustomStateEventCaller(address common.Address, caller bind.ContractCaller) (*PrivateVariableCustomStateEventCaller, error) {
	contract, err := bindPrivateVariableCustomStateEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateVariableCustomStateEventCaller{contract: contract}, nil
}

// NewPrivateVariableCustomStateEventTransactor creates a new write-only instance of PrivateVariableCustomStateEvent, bound to a specific deployed contract.
func NewPrivateVariableCustomStateEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivateVariableCustomStateEventTransactor, error) {
	contract, err := bindPrivateVariableCustomStateEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateVariableCustomStateEventTransactor{contract: contract}, nil
}

// NewPrivateVariableCustomStateEventFilterer creates a new log filterer instance of PrivateVariableCustomStateEvent, bound to a specific deployed contract.
func NewPrivateVariableCustomStateEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivateVariableCustomStateEventFilterer, error) {
	contract, err := bindPrivateVariableCustomStateEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivateVariableCustomStateEventFilterer{contract: contract}, nil
}

// bindPrivateVariableCustomStateEvent binds a generic wrapper to an already deployed contract.
func bindPrivateVariableCustomStateEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrivateVariableCustomStateEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrivateVariableCustomStateEvent.Contract.PrivateVariableCustomStateEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.Contract.PrivateVariableCustomStateEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.Contract.PrivateVariableCustomStateEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrivateVariableCustomStateEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventCaller) GetValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateVariableCustomStateEvent.contract.Call(opts, &out, "getValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventSession) GetValue() (*big.Int, error) {
	return _PrivateVariableCustomStateEvent.Contract.GetValue(&_PrivateVariableCustomStateEvent.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventCallerSession) GetValue() (*big.Int, error) {
	return _PrivateVariableCustomStateEvent.Contract.GetValue(&_PrivateVariableCustomStateEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventSession) Increment() (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.Contract.Increment(&_PrivateVariableCustomStateEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PrivateVariableCustomStateEvent.Contract.Increment(&_PrivateVariableCustomStateEvent.TransactOpts)
}

// PrivateVariableCustomStateEventValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the PrivateVariableCustomStateEvent contract.
type PrivateVariableCustomStateEventValueChangedIterator struct {
	Event *PrivateVariableCustomStateEventValueChanged // Event containing the contract specifics and raw log

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
func (it *PrivateVariableCustomStateEventValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateVariableCustomStateEventValueChanged)
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
		it.Event = new(PrivateVariableCustomStateEventValueChanged)
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
func (it *PrivateVariableCustomStateEventValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateVariableCustomStateEventValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateVariableCustomStateEventValueChanged represents a ValueChanged event raised by the PrivateVariableCustomStateEvent contract.
type PrivateVariableCustomStateEventValueChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventFilterer) FilterValueChanged(opts *bind.FilterOpts) (*PrivateVariableCustomStateEventValueChangedIterator, error) {

	logs, sub, err := _PrivateVariableCustomStateEvent.contract.FilterLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return &PrivateVariableCustomStateEventValueChangedIterator{contract: _PrivateVariableCustomStateEvent.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *PrivateVariableCustomStateEventValueChanged) (event.Subscription, error) {

	logs, sub, err := _PrivateVariableCustomStateEvent.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateVariableCustomStateEventValueChanged)
				if err := _PrivateVariableCustomStateEvent.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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

// ParseValueChanged is a log parse operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_PrivateVariableCustomStateEvent *PrivateVariableCustomStateEventFilterer) ParseValueChanged(log types.Log) (*PrivateVariableCustomStateEventValueChanged, error) {
	event := new(PrivateVariableCustomStateEventValueChanged)
	if err := _PrivateVariableCustomStateEvent.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicMappingEthDBUpdateEventMetaData contains all meta data concerning the PublicMappingEthDBUpdateEvent contract.
var PublicMappingEthDBUpdateEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mappingA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mappingB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105498061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063d09de08a14610043578063e36b791d1461004d578063f49925c21461007e575b5f5ffd5b61004b61009d565b005b61006c61005b366004610400565b60016020525f908152604090205481565b60405190815260200160405180910390f35b61006c61008c366004610400565b5f6020819052908152604090205481565b60015f8181526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546100d191610417565b5f60208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d919091556002908190527fabbb5caa7dda850e60932de0934eb1f9d0f59695050f761dc64e443e5030a5695461012d91610417565b7fabbb5caa7dda850e60932de0934eb1f9d0f59695050f761dc64e443e5030a5695560015f8190526020527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f54610185906003610417565b60016020527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f5560025f527fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f546101dd906004610417565b7fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f5560408051808201909152600b81526a6d617070696e67415b315d60a81b60208281019190915260015f90815290527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546102599190610332565b6102976040518060400160405280600b81526020016a6d617070696e67415b325d60a81b8152505f5f600281526020019081526020015f2054610332565b6102d66040518060400160405280600b81526020016a6d617070696e67425b315d60a81b81525060015f600181526020019081526020015f2054610332565b60408051808201909152600b81526a6d617070696e67425b325d60a81b60208083019190915260025f52600190527fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f546103309190610332565b565b61035f828260405160200161034991815260200190565b6040516020818303038152906040526003610363565b5050565b6040805160018082528183019092525f91816020015b606081526020019060019003908161037957905050905083815f815181106103a3576103a361043c565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a81848460058111156103e3576103e3610450565b6040516103f293929190610492565b60405180910390a150505050565b5f60208284031215610410575f5ffd5b5035919050565b8082018082111561043657634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b828110156104e957607f198786030184526104d4858351610464565b945060209384019391909101906001016104b8565b5050505082810360208401526104ff8186610464565b91505060ff8316604083015294935050505056fea2646970667358221220f98d8e7e91fa2e6c581f0ca29f1c3485d590cc63f9b58ac3495f75fd0e9a60a164736f6c634300081c0033",
}

// PublicMappingEthDBUpdateEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicMappingEthDBUpdateEventMetaData.ABI instead.
var PublicMappingEthDBUpdateEventABI = PublicMappingEthDBUpdateEventMetaData.ABI

// PublicMappingEthDBUpdateEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicMappingEthDBUpdateEventMetaData.Bin instead.
var PublicMappingEthDBUpdateEventBin = PublicMappingEthDBUpdateEventMetaData.Bin

// DeployPublicMappingEthDBUpdateEvent deploys a new Ethereum contract, binding an instance of PublicMappingEthDBUpdateEvent to it.
func DeployPublicMappingEthDBUpdateEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicMappingEthDBUpdateEvent, error) {
	parsed, err := PublicMappingEthDBUpdateEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicMappingEthDBUpdateEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicMappingEthDBUpdateEvent{PublicMappingEthDBUpdateEventCaller: PublicMappingEthDBUpdateEventCaller{contract: contract}, PublicMappingEthDBUpdateEventTransactor: PublicMappingEthDBUpdateEventTransactor{contract: contract}, PublicMappingEthDBUpdateEventFilterer: PublicMappingEthDBUpdateEventFilterer{contract: contract}}, nil
}

// PublicMappingEthDBUpdateEvent is an auto generated Go binding around an Ethereum contract.
type PublicMappingEthDBUpdateEvent struct {
	PublicMappingEthDBUpdateEventCaller     // Read-only binding to the contract
	PublicMappingEthDBUpdateEventTransactor // Write-only binding to the contract
	PublicMappingEthDBUpdateEventFilterer   // Log filterer for contract events
}

// PublicMappingEthDBUpdateEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicMappingEthDBUpdateEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicMappingEthDBUpdateEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicMappingEthDBUpdateEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicMappingEthDBUpdateEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicMappingEthDBUpdateEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicMappingEthDBUpdateEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicMappingEthDBUpdateEventSession struct {
	Contract     *PublicMappingEthDBUpdateEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PublicMappingEthDBUpdateEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicMappingEthDBUpdateEventCallerSession struct {
	Contract *PublicMappingEthDBUpdateEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// PublicMappingEthDBUpdateEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicMappingEthDBUpdateEventTransactorSession struct {
	Contract     *PublicMappingEthDBUpdateEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// PublicMappingEthDBUpdateEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicMappingEthDBUpdateEventRaw struct {
	Contract *PublicMappingEthDBUpdateEvent // Generic contract binding to access the raw methods on
}

// PublicMappingEthDBUpdateEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicMappingEthDBUpdateEventCallerRaw struct {
	Contract *PublicMappingEthDBUpdateEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicMappingEthDBUpdateEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicMappingEthDBUpdateEventTransactorRaw struct {
	Contract *PublicMappingEthDBUpdateEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicMappingEthDBUpdateEvent creates a new instance of PublicMappingEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicMappingEthDBUpdateEvent(address common.Address, backend bind.ContractBackend) (*PublicMappingEthDBUpdateEvent, error) {
	contract, err := bindPublicMappingEthDBUpdateEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicMappingEthDBUpdateEvent{PublicMappingEthDBUpdateEventCaller: PublicMappingEthDBUpdateEventCaller{contract: contract}, PublicMappingEthDBUpdateEventTransactor: PublicMappingEthDBUpdateEventTransactor{contract: contract}, PublicMappingEthDBUpdateEventFilterer: PublicMappingEthDBUpdateEventFilterer{contract: contract}}, nil
}

// NewPublicMappingEthDBUpdateEventCaller creates a new read-only instance of PublicMappingEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicMappingEthDBUpdateEventCaller(address common.Address, caller bind.ContractCaller) (*PublicMappingEthDBUpdateEventCaller, error) {
	contract, err := bindPublicMappingEthDBUpdateEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicMappingEthDBUpdateEventCaller{contract: contract}, nil
}

// NewPublicMappingEthDBUpdateEventTransactor creates a new write-only instance of PublicMappingEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicMappingEthDBUpdateEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicMappingEthDBUpdateEventTransactor, error) {
	contract, err := bindPublicMappingEthDBUpdateEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicMappingEthDBUpdateEventTransactor{contract: contract}, nil
}

// NewPublicMappingEthDBUpdateEventFilterer creates a new log filterer instance of PublicMappingEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicMappingEthDBUpdateEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicMappingEthDBUpdateEventFilterer, error) {
	contract, err := bindPublicMappingEthDBUpdateEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicMappingEthDBUpdateEventFilterer{contract: contract}, nil
}

// bindPublicMappingEthDBUpdateEvent binds a generic wrapper to an already deployed contract.
func bindPublicMappingEthDBUpdateEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicMappingEthDBUpdateEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicMappingEthDBUpdateEvent.Contract.PublicMappingEthDBUpdateEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.PublicMappingEthDBUpdateEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.PublicMappingEthDBUpdateEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicMappingEthDBUpdateEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.contract.Transact(opts, method, params...)
}

// MappingA is a free data retrieval call binding the contract method 0xf49925c2.
//
// Solidity: function mappingA(uint256 ) view returns(uint256)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventCaller) MappingA(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicMappingEthDBUpdateEvent.contract.Call(opts, &out, "mappingA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappingA is a free data retrieval call binding the contract method 0xf49925c2.
//
// Solidity: function mappingA(uint256 ) view returns(uint256)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventSession) MappingA(arg0 *big.Int) (*big.Int, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.MappingA(&_PublicMappingEthDBUpdateEvent.CallOpts, arg0)
}

// MappingA is a free data retrieval call binding the contract method 0xf49925c2.
//
// Solidity: function mappingA(uint256 ) view returns(uint256)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventCallerSession) MappingA(arg0 *big.Int) (*big.Int, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.MappingA(&_PublicMappingEthDBUpdateEvent.CallOpts, arg0)
}

// MappingB is a free data retrieval call binding the contract method 0xe36b791d.
//
// Solidity: function mappingB(uint256 ) view returns(uint256)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventCaller) MappingB(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicMappingEthDBUpdateEvent.contract.Call(opts, &out, "mappingB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappingB is a free data retrieval call binding the contract method 0xe36b791d.
//
// Solidity: function mappingB(uint256 ) view returns(uint256)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventSession) MappingB(arg0 *big.Int) (*big.Int, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.MappingB(&_PublicMappingEthDBUpdateEvent.CallOpts, arg0)
}

// MappingB is a free data retrieval call binding the contract method 0xe36b791d.
//
// Solidity: function mappingB(uint256 ) view returns(uint256)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventCallerSession) MappingB(arg0 *big.Int) (*big.Int, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.MappingB(&_PublicMappingEthDBUpdateEvent.CallOpts, arg0)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventSession) Increment() (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.Increment(&_PublicMappingEthDBUpdateEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicMappingEthDBUpdateEvent.Contract.Increment(&_PublicMappingEthDBUpdateEvent.TransactOpts)
}

// PublicMappingEthDBUpdateEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicMappingEthDBUpdateEvent contract.
type PublicMappingEthDBUpdateEventEthDBPathUpdateIterator struct {
	Event *PublicMappingEthDBUpdateEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicMappingEthDBUpdateEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicMappingEthDBUpdateEventEthDBPathUpdate)
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
		it.Event = new(PublicMappingEthDBUpdateEventEthDBPathUpdate)
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
func (it *PublicMappingEthDBUpdateEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicMappingEthDBUpdateEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicMappingEthDBUpdateEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicMappingEthDBUpdateEvent contract.
type PublicMappingEthDBUpdateEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicMappingEthDBUpdateEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicMappingEthDBUpdateEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicMappingEthDBUpdateEventEthDBPathUpdateIterator{contract: _PublicMappingEthDBUpdateEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicMappingEthDBUpdateEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicMappingEthDBUpdateEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicMappingEthDBUpdateEventEthDBPathUpdate)
				if err := _PublicMappingEthDBUpdateEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicMappingEthDBUpdateEventEthDBPathUpdate, error) {
	event := new(PublicMappingEthDBUpdateEventEthDBPathUpdate)
	if err := _PublicMappingEthDBUpdateEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicMappingEthDBUpdateEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicMappingEthDBUpdateEvent contract.
type PublicMappingEthDBUpdateEventEthDBUpdateIterator struct {
	Event *PublicMappingEthDBUpdateEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicMappingEthDBUpdateEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicMappingEthDBUpdateEventEthDBUpdate)
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
		it.Event = new(PublicMappingEthDBUpdateEventEthDBUpdate)
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
func (it *PublicMappingEthDBUpdateEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicMappingEthDBUpdateEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicMappingEthDBUpdateEventEthDBUpdate represents a EthDBUpdate event raised by the PublicMappingEthDBUpdateEvent contract.
type PublicMappingEthDBUpdateEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicMappingEthDBUpdateEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicMappingEthDBUpdateEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicMappingEthDBUpdateEventEthDBUpdateIterator{contract: _PublicMappingEthDBUpdateEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicMappingEthDBUpdateEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicMappingEthDBUpdateEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicMappingEthDBUpdateEventEthDBUpdate)
				if err := _PublicMappingEthDBUpdateEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicMappingEthDBUpdateEvent *PublicMappingEthDBUpdateEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicMappingEthDBUpdateEventEthDBUpdate, error) {
	event := new(PublicMappingEthDBUpdateEventEthDBUpdate)
	if err := _PublicMappingEthDBUpdateEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicStructEthDBUpdateEventMetaData contains all meta data concerning the PublicStructEthDBUpdateEvent contract.
var PublicStructEthDBUpdateEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"structA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"structB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103b78061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806311ba731f14610043578063d09de08a14610069578063d86387a514610073575b5f5ffd5b5f54600154610050919082565b6040805192835260208301919091520160405180910390f35b610071610081565b005b600254600354610050919082565b60015f5f015f8282546100949190610285565b909155505060018054600291905f906100ae908490610285565b909155505060028054600391905f906100c8908490610285565b909155505060038054600491905f906100e2908490610285565b909155505060408051808201909152600e81526d737472756374412e76616c75654160901b60208201525f5461011891906101b7565b61014c6040518060400160405280600e81526020016d39ba393ab1ba20973b30b63ab2a160911b8152505f600101546101b7565b6101806040518060400160405280600e81526020016d737472756374422e76616c75654160901b81525060025f01546101b7565b6101b56040518060400160405280600e81526020016d39ba393ab1ba21173b30b63ab2a160911b8152506002600101546101b7565b565b6101e482826040516020016101ce91815260200190565b60405160208183030381529060405260036101e8565b5050565b6040805160018082528183019092525f91816020015b60608152602001906001900390816101fe57905050905083815f81518110610228576102286102aa565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a8184846005811115610268576102686102be565b60405161027793929190610300565b60405180910390a150505050565b808201808211156102a457634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b8281101561035757607f198786030184526103428583516102d2565b94506020938401939190910190600101610326565b50505050828103602084015261036d81866102d2565b91505060ff8316604083015294935050505056fea26469706673582212203511370b54cb375d2e0e374d6e486c5f886f41db1a3f96b1fdbc7b05a58709a564736f6c634300081c0033",
}

// PublicStructEthDBUpdateEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicStructEthDBUpdateEventMetaData.ABI instead.
var PublicStructEthDBUpdateEventABI = PublicStructEthDBUpdateEventMetaData.ABI

// PublicStructEthDBUpdateEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicStructEthDBUpdateEventMetaData.Bin instead.
var PublicStructEthDBUpdateEventBin = PublicStructEthDBUpdateEventMetaData.Bin

// DeployPublicStructEthDBUpdateEvent deploys a new Ethereum contract, binding an instance of PublicStructEthDBUpdateEvent to it.
func DeployPublicStructEthDBUpdateEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicStructEthDBUpdateEvent, error) {
	parsed, err := PublicStructEthDBUpdateEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicStructEthDBUpdateEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicStructEthDBUpdateEvent{PublicStructEthDBUpdateEventCaller: PublicStructEthDBUpdateEventCaller{contract: contract}, PublicStructEthDBUpdateEventTransactor: PublicStructEthDBUpdateEventTransactor{contract: contract}, PublicStructEthDBUpdateEventFilterer: PublicStructEthDBUpdateEventFilterer{contract: contract}}, nil
}

// PublicStructEthDBUpdateEvent is an auto generated Go binding around an Ethereum contract.
type PublicStructEthDBUpdateEvent struct {
	PublicStructEthDBUpdateEventCaller     // Read-only binding to the contract
	PublicStructEthDBUpdateEventTransactor // Write-only binding to the contract
	PublicStructEthDBUpdateEventFilterer   // Log filterer for contract events
}

// PublicStructEthDBUpdateEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicStructEthDBUpdateEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicStructEthDBUpdateEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicStructEthDBUpdateEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicStructEthDBUpdateEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicStructEthDBUpdateEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicStructEthDBUpdateEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicStructEthDBUpdateEventSession struct {
	Contract     *PublicStructEthDBUpdateEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PublicStructEthDBUpdateEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicStructEthDBUpdateEventCallerSession struct {
	Contract *PublicStructEthDBUpdateEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// PublicStructEthDBUpdateEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicStructEthDBUpdateEventTransactorSession struct {
	Contract     *PublicStructEthDBUpdateEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PublicStructEthDBUpdateEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicStructEthDBUpdateEventRaw struct {
	Contract *PublicStructEthDBUpdateEvent // Generic contract binding to access the raw methods on
}

// PublicStructEthDBUpdateEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicStructEthDBUpdateEventCallerRaw struct {
	Contract *PublicStructEthDBUpdateEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicStructEthDBUpdateEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicStructEthDBUpdateEventTransactorRaw struct {
	Contract *PublicStructEthDBUpdateEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicStructEthDBUpdateEvent creates a new instance of PublicStructEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicStructEthDBUpdateEvent(address common.Address, backend bind.ContractBackend) (*PublicStructEthDBUpdateEvent, error) {
	contract, err := bindPublicStructEthDBUpdateEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicStructEthDBUpdateEvent{PublicStructEthDBUpdateEventCaller: PublicStructEthDBUpdateEventCaller{contract: contract}, PublicStructEthDBUpdateEventTransactor: PublicStructEthDBUpdateEventTransactor{contract: contract}, PublicStructEthDBUpdateEventFilterer: PublicStructEthDBUpdateEventFilterer{contract: contract}}, nil
}

// NewPublicStructEthDBUpdateEventCaller creates a new read-only instance of PublicStructEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicStructEthDBUpdateEventCaller(address common.Address, caller bind.ContractCaller) (*PublicStructEthDBUpdateEventCaller, error) {
	contract, err := bindPublicStructEthDBUpdateEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicStructEthDBUpdateEventCaller{contract: contract}, nil
}

// NewPublicStructEthDBUpdateEventTransactor creates a new write-only instance of PublicStructEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicStructEthDBUpdateEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicStructEthDBUpdateEventTransactor, error) {
	contract, err := bindPublicStructEthDBUpdateEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicStructEthDBUpdateEventTransactor{contract: contract}, nil
}

// NewPublicStructEthDBUpdateEventFilterer creates a new log filterer instance of PublicStructEthDBUpdateEvent, bound to a specific deployed contract.
func NewPublicStructEthDBUpdateEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicStructEthDBUpdateEventFilterer, error) {
	contract, err := bindPublicStructEthDBUpdateEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicStructEthDBUpdateEventFilterer{contract: contract}, nil
}

// bindPublicStructEthDBUpdateEvent binds a generic wrapper to an already deployed contract.
func bindPublicStructEthDBUpdateEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicStructEthDBUpdateEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicStructEthDBUpdateEvent.Contract.PublicStructEthDBUpdateEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.Contract.PublicStructEthDBUpdateEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.Contract.PublicStructEthDBUpdateEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicStructEthDBUpdateEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.Contract.contract.Transact(opts, method, params...)
}

// StructA is a free data retrieval call binding the contract method 0x11ba731f.
//
// Solidity: function structA() view returns(uint256 valueA, uint256 valueB)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventCaller) StructA(opts *bind.CallOpts) (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	var out []interface{}
	err := _PublicStructEthDBUpdateEvent.contract.Call(opts, &out, "structA")

	outstruct := new(struct {
		ValueA *big.Int
		ValueB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValueA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StructA is a free data retrieval call binding the contract method 0x11ba731f.
//
// Solidity: function structA() view returns(uint256 valueA, uint256 valueB)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventSession) StructA() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicStructEthDBUpdateEvent.Contract.StructA(&_PublicStructEthDBUpdateEvent.CallOpts)
}

// StructA is a free data retrieval call binding the contract method 0x11ba731f.
//
// Solidity: function structA() view returns(uint256 valueA, uint256 valueB)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventCallerSession) StructA() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicStructEthDBUpdateEvent.Contract.StructA(&_PublicStructEthDBUpdateEvent.CallOpts)
}

// StructB is a free data retrieval call binding the contract method 0xd86387a5.
//
// Solidity: function structB() view returns(uint256 valueA, uint256 valueB)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventCaller) StructB(opts *bind.CallOpts) (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	var out []interface{}
	err := _PublicStructEthDBUpdateEvent.contract.Call(opts, &out, "structB")

	outstruct := new(struct {
		ValueA *big.Int
		ValueB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValueA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StructB is a free data retrieval call binding the contract method 0xd86387a5.
//
// Solidity: function structB() view returns(uint256 valueA, uint256 valueB)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventSession) StructB() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicStructEthDBUpdateEvent.Contract.StructB(&_PublicStructEthDBUpdateEvent.CallOpts)
}

// StructB is a free data retrieval call binding the contract method 0xd86387a5.
//
// Solidity: function structB() view returns(uint256 valueA, uint256 valueB)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventCallerSession) StructB() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicStructEthDBUpdateEvent.Contract.StructB(&_PublicStructEthDBUpdateEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventSession) Increment() (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.Contract.Increment(&_PublicStructEthDBUpdateEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicStructEthDBUpdateEvent.Contract.Increment(&_PublicStructEthDBUpdateEvent.TransactOpts)
}

// PublicStructEthDBUpdateEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicStructEthDBUpdateEvent contract.
type PublicStructEthDBUpdateEventEthDBPathUpdateIterator struct {
	Event *PublicStructEthDBUpdateEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicStructEthDBUpdateEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicStructEthDBUpdateEventEthDBPathUpdate)
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
		it.Event = new(PublicStructEthDBUpdateEventEthDBPathUpdate)
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
func (it *PublicStructEthDBUpdateEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicStructEthDBUpdateEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicStructEthDBUpdateEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicStructEthDBUpdateEvent contract.
type PublicStructEthDBUpdateEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicStructEthDBUpdateEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicStructEthDBUpdateEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicStructEthDBUpdateEventEthDBPathUpdateIterator{contract: _PublicStructEthDBUpdateEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicStructEthDBUpdateEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicStructEthDBUpdateEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicStructEthDBUpdateEventEthDBPathUpdate)
				if err := _PublicStructEthDBUpdateEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicStructEthDBUpdateEventEthDBPathUpdate, error) {
	event := new(PublicStructEthDBUpdateEventEthDBPathUpdate)
	if err := _PublicStructEthDBUpdateEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicStructEthDBUpdateEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicStructEthDBUpdateEvent contract.
type PublicStructEthDBUpdateEventEthDBUpdateIterator struct {
	Event *PublicStructEthDBUpdateEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicStructEthDBUpdateEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicStructEthDBUpdateEventEthDBUpdate)
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
		it.Event = new(PublicStructEthDBUpdateEventEthDBUpdate)
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
func (it *PublicStructEthDBUpdateEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicStructEthDBUpdateEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicStructEthDBUpdateEventEthDBUpdate represents a EthDBUpdate event raised by the PublicStructEthDBUpdateEvent contract.
type PublicStructEthDBUpdateEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicStructEthDBUpdateEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicStructEthDBUpdateEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicStructEthDBUpdateEventEthDBUpdateIterator{contract: _PublicStructEthDBUpdateEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicStructEthDBUpdateEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicStructEthDBUpdateEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicStructEthDBUpdateEventEthDBUpdate)
				if err := _PublicStructEthDBUpdateEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicStructEthDBUpdateEvent *PublicStructEthDBUpdateEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicStructEthDBUpdateEventEthDBUpdate, error) {
	event := new(PublicStructEthDBUpdateEventEthDBUpdate)
	if err := _PublicStructEthDBUpdateEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableCustomStateEventMetaData contains all meta data concerning the PublicVariableCustomStateEvent contract.
var PublicVariableCustomStateEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060fa8061001b5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80633fa4f245146034578063d09de08a14604d575b5f5ffd5b603b5f5481565b60405190815260200160405180910390f35b60536055565b005b60015f5f8282546064919060a0565b90915550505f546040519081527f93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c599060200160405180910390a1565b8082018082111560be57634e487b7160e01b5f52601160045260245ffd5b9291505056fea2646970667358221220793f96afa39d15d1222cd765d738890b6e2188ed92c2daa2f5758721d83ab93a64736f6c634300081c0033",
}

// PublicVariableCustomStateEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableCustomStateEventMetaData.ABI instead.
var PublicVariableCustomStateEventABI = PublicVariableCustomStateEventMetaData.ABI

// PublicVariableCustomStateEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableCustomStateEventMetaData.Bin instead.
var PublicVariableCustomStateEventBin = PublicVariableCustomStateEventMetaData.Bin

// DeployPublicVariableCustomStateEvent deploys a new Ethereum contract, binding an instance of PublicVariableCustomStateEvent to it.
func DeployPublicVariableCustomStateEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableCustomStateEvent, error) {
	parsed, err := PublicVariableCustomStateEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableCustomStateEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableCustomStateEvent{PublicVariableCustomStateEventCaller: PublicVariableCustomStateEventCaller{contract: contract}, PublicVariableCustomStateEventTransactor: PublicVariableCustomStateEventTransactor{contract: contract}, PublicVariableCustomStateEventFilterer: PublicVariableCustomStateEventFilterer{contract: contract}}, nil
}

// PublicVariableCustomStateEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableCustomStateEvent struct {
	PublicVariableCustomStateEventCaller     // Read-only binding to the contract
	PublicVariableCustomStateEventTransactor // Write-only binding to the contract
	PublicVariableCustomStateEventFilterer   // Log filterer for contract events
}

// PublicVariableCustomStateEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableCustomStateEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableCustomStateEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableCustomStateEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableCustomStateEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableCustomStateEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableCustomStateEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableCustomStateEventSession struct {
	Contract     *PublicVariableCustomStateEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PublicVariableCustomStateEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableCustomStateEventCallerSession struct {
	Contract *PublicVariableCustomStateEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// PublicVariableCustomStateEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableCustomStateEventTransactorSession struct {
	Contract     *PublicVariableCustomStateEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// PublicVariableCustomStateEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableCustomStateEventRaw struct {
	Contract *PublicVariableCustomStateEvent // Generic contract binding to access the raw methods on
}

// PublicVariableCustomStateEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableCustomStateEventCallerRaw struct {
	Contract *PublicVariableCustomStateEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableCustomStateEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableCustomStateEventTransactorRaw struct {
	Contract *PublicVariableCustomStateEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableCustomStateEvent creates a new instance of PublicVariableCustomStateEvent, bound to a specific deployed contract.
func NewPublicVariableCustomStateEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableCustomStateEvent, error) {
	contract, err := bindPublicVariableCustomStateEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableCustomStateEvent{PublicVariableCustomStateEventCaller: PublicVariableCustomStateEventCaller{contract: contract}, PublicVariableCustomStateEventTransactor: PublicVariableCustomStateEventTransactor{contract: contract}, PublicVariableCustomStateEventFilterer: PublicVariableCustomStateEventFilterer{contract: contract}}, nil
}

// NewPublicVariableCustomStateEventCaller creates a new read-only instance of PublicVariableCustomStateEvent, bound to a specific deployed contract.
func NewPublicVariableCustomStateEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableCustomStateEventCaller, error) {
	contract, err := bindPublicVariableCustomStateEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableCustomStateEventCaller{contract: contract}, nil
}

// NewPublicVariableCustomStateEventTransactor creates a new write-only instance of PublicVariableCustomStateEvent, bound to a specific deployed contract.
func NewPublicVariableCustomStateEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableCustomStateEventTransactor, error) {
	contract, err := bindPublicVariableCustomStateEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableCustomStateEventTransactor{contract: contract}, nil
}

// NewPublicVariableCustomStateEventFilterer creates a new log filterer instance of PublicVariableCustomStateEvent, bound to a specific deployed contract.
func NewPublicVariableCustomStateEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableCustomStateEventFilterer, error) {
	contract, err := bindPublicVariableCustomStateEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableCustomStateEventFilterer{contract: contract}, nil
}

// bindPublicVariableCustomStateEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableCustomStateEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableCustomStateEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableCustomStateEvent.Contract.PublicVariableCustomStateEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.Contract.PublicVariableCustomStateEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.Contract.PublicVariableCustomStateEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableCustomStateEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.Contract.contract.Transact(opts, method, params...)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableCustomStateEvent.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventSession) Value() (*big.Int, error) {
	return _PublicVariableCustomStateEvent.Contract.Value(&_PublicVariableCustomStateEvent.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventCallerSession) Value() (*big.Int, error) {
	return _PublicVariableCustomStateEvent.Contract.Value(&_PublicVariableCustomStateEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.Contract.Increment(&_PublicVariableCustomStateEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableCustomStateEvent.Contract.Increment(&_PublicVariableCustomStateEvent.TransactOpts)
}

// PublicVariableCustomStateEventValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the PublicVariableCustomStateEvent contract.
type PublicVariableCustomStateEventValueChangedIterator struct {
	Event *PublicVariableCustomStateEventValueChanged // Event containing the contract specifics and raw log

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
func (it *PublicVariableCustomStateEventValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableCustomStateEventValueChanged)
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
		it.Event = new(PublicVariableCustomStateEventValueChanged)
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
func (it *PublicVariableCustomStateEventValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableCustomStateEventValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableCustomStateEventValueChanged represents a ValueChanged event raised by the PublicVariableCustomStateEvent contract.
type PublicVariableCustomStateEventValueChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventFilterer) FilterValueChanged(opts *bind.FilterOpts) (*PublicVariableCustomStateEventValueChangedIterator, error) {

	logs, sub, err := _PublicVariableCustomStateEvent.contract.FilterLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return &PublicVariableCustomStateEventValueChangedIterator{contract: _PublicVariableCustomStateEvent.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *PublicVariableCustomStateEventValueChanged) (event.Subscription, error) {

	logs, sub, err := _PublicVariableCustomStateEvent.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableCustomStateEventValueChanged)
				if err := _PublicVariableCustomStateEvent.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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

// ParseValueChanged is a log parse operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_PublicVariableCustomStateEvent *PublicVariableCustomStateEventFilterer) ParseValueChanged(log types.Log) (*PublicVariableCustomStateEventValueChanged, error) {
	event := new(PublicVariableCustomStateEventValueChanged)
	if err := _PublicVariableCustomStateEvent.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBPathUpdateMultipleEventMetaData contains all meta data concerning the PublicVariableEthDBPathUpdateMultipleEvent contract.
var PublicVariableEthDBPathUpdateMultipleEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f55600a6001553480156017575f5ffd5b50610379806100255f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df5161146100425780636af155051461005d578063d09de08a14610065575b5f5ffd5b61004b60015481565b60405190815260200160405180910390f35b61004b5f5481565b61006d61006f565b005b60015f5f8282546100809190610247565b92505081905550600260015f8282546100999190610247565b925050819055506100c96040518060400160405280600681526020016576616c75654160d01b8152505f54610179565b6100f3604051806040016040528060068152602001653b30b63ab2a160d11b815250600154610179565b60015f5f8282546101049190610247565b92505081905550600260015f82825461011d9190610247565b9250508190555061014d6040518060400160405280600681526020016576616c75654160d01b8152505f54610179565b610177604051806040016040528060068152602001653b30b63ab2a160d11b815250600154610179565b565b6101a6828260405160200161019091815260200190565b60405160208183030381529060405260036101aa565b5050565b6040805160018082528183019092525f91816020015b60608152602001906001900390816101c057905050905083815f815181106101ea576101ea61026c565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a818484600581111561022a5761022a610280565b604051610239939291906102c2565b60405180910390a150505050565b8082018082111561026657634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b8281101561031957607f19878603018452610304858351610294565b945060209384019391909101906001016102e8565b50505050828103602084015261032f8186610294565b91505060ff8316604083015294935050505056fea26469706673582212203b91edeb9c9e7527bcca893f45e6b304e0bd548eb061a770f72a3ad3dc362a7a64736f6c634300081c0033",
}

// PublicVariableEthDBPathUpdateMultipleEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBPathUpdateMultipleEventMetaData.ABI instead.
var PublicVariableEthDBPathUpdateMultipleEventABI = PublicVariableEthDBPathUpdateMultipleEventMetaData.ABI

// PublicVariableEthDBPathUpdateMultipleEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBPathUpdateMultipleEventMetaData.Bin instead.
var PublicVariableEthDBPathUpdateMultipleEventBin = PublicVariableEthDBPathUpdateMultipleEventMetaData.Bin

// DeployPublicVariableEthDBPathUpdateMultipleEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBPathUpdateMultipleEvent to it.
func DeployPublicVariableEthDBPathUpdateMultipleEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBPathUpdateMultipleEvent, error) {
	parsed, err := PublicVariableEthDBPathUpdateMultipleEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBPathUpdateMultipleEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBPathUpdateMultipleEvent{PublicVariableEthDBPathUpdateMultipleEventCaller: PublicVariableEthDBPathUpdateMultipleEventCaller{contract: contract}, PublicVariableEthDBPathUpdateMultipleEventTransactor: PublicVariableEthDBPathUpdateMultipleEventTransactor{contract: contract}, PublicVariableEthDBPathUpdateMultipleEventFilterer: PublicVariableEthDBPathUpdateMultipleEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBPathUpdateMultipleEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateMultipleEvent struct {
	PublicVariableEthDBPathUpdateMultipleEventCaller     // Read-only binding to the contract
	PublicVariableEthDBPathUpdateMultipleEventTransactor // Write-only binding to the contract
	PublicVariableEthDBPathUpdateMultipleEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBPathUpdateMultipleEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateMultipleEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBPathUpdateMultipleEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateMultipleEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBPathUpdateMultipleEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBPathUpdateMultipleEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBPathUpdateMultipleEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBPathUpdateMultipleEventSession struct {
	Contract     *PublicVariableEthDBPathUpdateMultipleEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                               // Call options to use throughout this session
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// PublicVariableEthDBPathUpdateMultipleEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBPathUpdateMultipleEventCallerSession struct {
	Contract *PublicVariableEthDBPathUpdateMultipleEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                     // Call options to use throughout this session
}

// PublicVariableEthDBPathUpdateMultipleEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBPathUpdateMultipleEventTransactorSession struct {
	Contract     *PublicVariableEthDBPathUpdateMultipleEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                     // Transaction auth options to use throughout this session
}

// PublicVariableEthDBPathUpdateMultipleEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateMultipleEventRaw struct {
	Contract *PublicVariableEthDBPathUpdateMultipleEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBPathUpdateMultipleEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateMultipleEventCallerRaw struct {
	Contract *PublicVariableEthDBPathUpdateMultipleEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBPathUpdateMultipleEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateMultipleEventTransactorRaw struct {
	Contract *PublicVariableEthDBPathUpdateMultipleEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBPathUpdateMultipleEvent creates a new instance of PublicVariableEthDBPathUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateMultipleEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBPathUpdateMultipleEvent, error) {
	contract, err := bindPublicVariableEthDBPathUpdateMultipleEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateMultipleEvent{PublicVariableEthDBPathUpdateMultipleEventCaller: PublicVariableEthDBPathUpdateMultipleEventCaller{contract: contract}, PublicVariableEthDBPathUpdateMultipleEventTransactor: PublicVariableEthDBPathUpdateMultipleEventTransactor{contract: contract}, PublicVariableEthDBPathUpdateMultipleEventFilterer: PublicVariableEthDBPathUpdateMultipleEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBPathUpdateMultipleEventCaller creates a new read-only instance of PublicVariableEthDBPathUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateMultipleEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBPathUpdateMultipleEventCaller, error) {
	contract, err := bindPublicVariableEthDBPathUpdateMultipleEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateMultipleEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBPathUpdateMultipleEventTransactor creates a new write-only instance of PublicVariableEthDBPathUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateMultipleEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBPathUpdateMultipleEventTransactor, error) {
	contract, err := bindPublicVariableEthDBPathUpdateMultipleEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateMultipleEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBPathUpdateMultipleEventFilterer creates a new log filterer instance of PublicVariableEthDBPathUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateMultipleEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBPathUpdateMultipleEventFilterer, error) {
	contract, err := bindPublicVariableEthDBPathUpdateMultipleEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateMultipleEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBPathUpdateMultipleEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBPathUpdateMultipleEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBPathUpdateMultipleEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.PublicVariableEthDBPathUpdateMultipleEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.PublicVariableEthDBPathUpdateMultipleEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.PublicVariableEthDBPathUpdateMultipleEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventCaller) ValueA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.ValueA(&_PublicVariableEthDBPathUpdateMultipleEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventCallerSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.ValueA(&_PublicVariableEthDBPathUpdateMultipleEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventCaller) ValueB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.ValueB(&_PublicVariableEthDBPathUpdateMultipleEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventCallerSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.ValueB(&_PublicVariableEthDBPathUpdateMultipleEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.Increment(&_PublicVariableEthDBPathUpdateMultipleEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateMultipleEvent.Contract.Increment(&_PublicVariableEthDBPathUpdateMultipleEvent.TransactOpts)
}

// PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBPathUpdateMultipleEvent contract.
type PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBPathUpdateMultipleEvent contract.
type PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBPathUpdateMultipleEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate)
				if err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBPathUpdateMultipleEventEthDBPathUpdate)
	if err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBPathUpdateMultipleEvent contract.
type PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate)
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
func (it *PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBPathUpdateMultipleEvent contract.
type PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateMultipleEventEthDBUpdateIterator{contract: _PublicVariableEthDBPathUpdateMultipleEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate)
				if err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBPathUpdateMultipleEvent *PublicVariableEthDBPathUpdateMultipleEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBPathUpdateMultipleEventEthDBUpdate)
	if err := _PublicVariableEthDBPathUpdateMultipleEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBPathUpdateUintEventMetaData contains all meta data concerning the PublicVariableEthDBPathUpdateUintEvent contract.
var PublicVariableEthDBPathUpdateUintEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f55600a6001553480156017575f5ffd5b506102f5806100255f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df5161146100425780636af155051461005d578063d09de08a14610065575b5f5ffd5b61004b60015481565b60405190815260200160405180910390f35b61004b5f5481565b61006d61006f565b005b60015f5f82825461008091906101c3565b92505081905550600260015f82825461009991906101c3565b925050819055506100c96040518060400160405280600681526020016576616c75654160d01b8152505f546100f5565b6100f3604051806040016040528060068152602001653b30b63ab2a160d11b8152506001546100f5565b565b610122828260405160200161010c91815260200190565b6040516020818303038152906040526003610126565b5050565b6040805160018082528183019092525f91816020015b606081526020019060019003908161013c57905050905083815f81518110610166576101666101e8565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a81848460058111156101a6576101a66101fc565b6040516101b59392919061023e565b60405180910390a150505050565b808201808211156101e257634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b8281101561029557607f19878603018452610280858351610210565b94506020938401939190910190600101610264565b5050505082810360208401526102ab8186610210565b91505060ff8316604083015294935050505056fea2646970667358221220da0c40b158e3a1b407f78890ef4d1d1caab1f7d70bbab0da7aa13a472dab9cab64736f6c634300081c0033",
}

// PublicVariableEthDBPathUpdateUintEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBPathUpdateUintEventMetaData.ABI instead.
var PublicVariableEthDBPathUpdateUintEventABI = PublicVariableEthDBPathUpdateUintEventMetaData.ABI

// PublicVariableEthDBPathUpdateUintEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBPathUpdateUintEventMetaData.Bin instead.
var PublicVariableEthDBPathUpdateUintEventBin = PublicVariableEthDBPathUpdateUintEventMetaData.Bin

// DeployPublicVariableEthDBPathUpdateUintEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBPathUpdateUintEvent to it.
func DeployPublicVariableEthDBPathUpdateUintEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBPathUpdateUintEvent, error) {
	parsed, err := PublicVariableEthDBPathUpdateUintEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBPathUpdateUintEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBPathUpdateUintEvent{PublicVariableEthDBPathUpdateUintEventCaller: PublicVariableEthDBPathUpdateUintEventCaller{contract: contract}, PublicVariableEthDBPathUpdateUintEventTransactor: PublicVariableEthDBPathUpdateUintEventTransactor{contract: contract}, PublicVariableEthDBPathUpdateUintEventFilterer: PublicVariableEthDBPathUpdateUintEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBPathUpdateUintEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateUintEvent struct {
	PublicVariableEthDBPathUpdateUintEventCaller     // Read-only binding to the contract
	PublicVariableEthDBPathUpdateUintEventTransactor // Write-only binding to the contract
	PublicVariableEthDBPathUpdateUintEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBPathUpdateUintEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateUintEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBPathUpdateUintEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateUintEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBPathUpdateUintEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBPathUpdateUintEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBPathUpdateUintEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBPathUpdateUintEventSession struct {
	Contract     *PublicVariableEthDBPathUpdateUintEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                           // Call options to use throughout this session
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PublicVariableEthDBPathUpdateUintEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBPathUpdateUintEventCallerSession struct {
	Contract *PublicVariableEthDBPathUpdateUintEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                 // Call options to use throughout this session
}

// PublicVariableEthDBPathUpdateUintEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBPathUpdateUintEventTransactorSession struct {
	Contract     *PublicVariableEthDBPathUpdateUintEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                 // Transaction auth options to use throughout this session
}

// PublicVariableEthDBPathUpdateUintEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateUintEventRaw struct {
	Contract *PublicVariableEthDBPathUpdateUintEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBPathUpdateUintEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateUintEventCallerRaw struct {
	Contract *PublicVariableEthDBPathUpdateUintEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBPathUpdateUintEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBPathUpdateUintEventTransactorRaw struct {
	Contract *PublicVariableEthDBPathUpdateUintEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBPathUpdateUintEvent creates a new instance of PublicVariableEthDBPathUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateUintEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBPathUpdateUintEvent, error) {
	contract, err := bindPublicVariableEthDBPathUpdateUintEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateUintEvent{PublicVariableEthDBPathUpdateUintEventCaller: PublicVariableEthDBPathUpdateUintEventCaller{contract: contract}, PublicVariableEthDBPathUpdateUintEventTransactor: PublicVariableEthDBPathUpdateUintEventTransactor{contract: contract}, PublicVariableEthDBPathUpdateUintEventFilterer: PublicVariableEthDBPathUpdateUintEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBPathUpdateUintEventCaller creates a new read-only instance of PublicVariableEthDBPathUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateUintEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBPathUpdateUintEventCaller, error) {
	contract, err := bindPublicVariableEthDBPathUpdateUintEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateUintEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBPathUpdateUintEventTransactor creates a new write-only instance of PublicVariableEthDBPathUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateUintEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBPathUpdateUintEventTransactor, error) {
	contract, err := bindPublicVariableEthDBPathUpdateUintEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateUintEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBPathUpdateUintEventFilterer creates a new log filterer instance of PublicVariableEthDBPathUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBPathUpdateUintEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBPathUpdateUintEventFilterer, error) {
	contract, err := bindPublicVariableEthDBPathUpdateUintEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateUintEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBPathUpdateUintEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBPathUpdateUintEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBPathUpdateUintEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.PublicVariableEthDBPathUpdateUintEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.PublicVariableEthDBPathUpdateUintEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.PublicVariableEthDBPathUpdateUintEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventCaller) ValueA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBPathUpdateUintEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.ValueA(&_PublicVariableEthDBPathUpdateUintEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventCallerSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.ValueA(&_PublicVariableEthDBPathUpdateUintEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventCaller) ValueB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBPathUpdateUintEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.ValueB(&_PublicVariableEthDBPathUpdateUintEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventCallerSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.ValueB(&_PublicVariableEthDBPathUpdateUintEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.Increment(&_PublicVariableEthDBPathUpdateUintEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBPathUpdateUintEvent.Contract.Increment(&_PublicVariableEthDBPathUpdateUintEvent.TransactOpts)
}

// PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBPathUpdateUintEvent contract.
type PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBPathUpdateUintEvent contract.
type PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateUintEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateUintEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBPathUpdateUintEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateUintEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate)
				if err := _PublicVariableEthDBPathUpdateUintEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBPathUpdateUintEventEthDBPathUpdate)
	if err := _PublicVariableEthDBPathUpdateUintEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBPathUpdateUintEvent contract.
type PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBPathUpdateUintEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBPathUpdateUintEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBPathUpdateUintEventEthDBUpdate)
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
func (it *PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBPathUpdateUintEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBPathUpdateUintEvent contract.
type PublicVariableEthDBPathUpdateUintEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateUintEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBPathUpdateUintEventEthDBUpdateIterator{contract: _PublicVariableEthDBPathUpdateUintEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBPathUpdateUintEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBPathUpdateUintEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBPathUpdateUintEventEthDBUpdate)
				if err := _PublicVariableEthDBPathUpdateUintEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBPathUpdateUintEvent *PublicVariableEthDBPathUpdateUintEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBPathUpdateUintEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBPathUpdateUintEventEthDBUpdate)
	if err := _PublicVariableEthDBPathUpdateUintEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateArrayEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateArrayEvent contract.
var PublicVariableEthDBUpdateArrayEventMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addToArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"arrayA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"arrayB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b505f8054600181810183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101819055805480820182558183527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69081018390558154808301909255015561019c806100885f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80639b88967614610043578063ad3149a414610068578063d34e768d1461007b575b5f5ffd5b61005661005136600461012a565b610085565b60405190815260200160405180910390f35b61005661007636600461012a565b6100a4565b6100836100b2565b005b60018181548110610094575f80fd5b5f91825260209091200154905081565b5f8181548110610094575f80fd5b5f80546100c0906001610141565b8154600181810184555f9384526020909320015580546100e1906002610141565b81546001810183555f9283526020909220909101556100fe610100565b565b6040517fc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c905f90a1565b5f6020828403121561013a575f5ffd5b5035919050565b8082018082111561016057634e487b7160e01b5f52601160045260245ffd5b9291505056fea26469706673582212200d0a40f17249675beceffb9e9fd972443b1698145bd98f8892cb9a14077a9d7364736f6c634300081c0033",
}

// PublicVariableEthDBUpdateArrayEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateArrayEventMetaData.ABI instead.
var PublicVariableEthDBUpdateArrayEventABI = PublicVariableEthDBUpdateArrayEventMetaData.ABI

// PublicVariableEthDBUpdateArrayEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateArrayEventMetaData.Bin instead.
var PublicVariableEthDBUpdateArrayEventBin = PublicVariableEthDBUpdateArrayEventMetaData.Bin

// DeployPublicVariableEthDBUpdateArrayEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateArrayEvent to it.
func DeployPublicVariableEthDBUpdateArrayEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateArrayEvent, error) {
	parsed, err := PublicVariableEthDBUpdateArrayEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateArrayEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateArrayEvent{PublicVariableEthDBUpdateArrayEventCaller: PublicVariableEthDBUpdateArrayEventCaller{contract: contract}, PublicVariableEthDBUpdateArrayEventTransactor: PublicVariableEthDBUpdateArrayEventTransactor{contract: contract}, PublicVariableEthDBUpdateArrayEventFilterer: PublicVariableEthDBUpdateArrayEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateArrayEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateArrayEvent struct {
	PublicVariableEthDBUpdateArrayEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateArrayEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateArrayEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateArrayEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateArrayEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateArrayEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateArrayEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateArrayEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateArrayEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateArrayEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateArrayEventSession struct {
	Contract     *PublicVariableEthDBUpdateArrayEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateArrayEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateArrayEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateArrayEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// PublicVariableEthDBUpdateArrayEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateArrayEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateArrayEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateArrayEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateArrayEventRaw struct {
	Contract *PublicVariableEthDBUpdateArrayEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateArrayEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateArrayEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateArrayEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateArrayEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateArrayEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateArrayEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateArrayEvent creates a new instance of PublicVariableEthDBUpdateArrayEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateArrayEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateArrayEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateArrayEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateArrayEvent{PublicVariableEthDBUpdateArrayEventCaller: PublicVariableEthDBUpdateArrayEventCaller{contract: contract}, PublicVariableEthDBUpdateArrayEventTransactor: PublicVariableEthDBUpdateArrayEventTransactor{contract: contract}, PublicVariableEthDBUpdateArrayEventFilterer: PublicVariableEthDBUpdateArrayEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateArrayEventCaller creates a new read-only instance of PublicVariableEthDBUpdateArrayEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateArrayEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateArrayEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateArrayEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateArrayEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateArrayEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateArrayEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateArrayEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateArrayEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateArrayEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateArrayEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateArrayEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateArrayEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateArrayEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateArrayEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateArrayEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateArrayEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateArrayEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateArrayEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateArrayEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.PublicVariableEthDBUpdateArrayEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.PublicVariableEthDBUpdateArrayEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.PublicVariableEthDBUpdateArrayEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.contract.Transact(opts, method, params...)
}

// ArrayA is a free data retrieval call binding the contract method 0xad3149a4.
//
// Solidity: function arrayA(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventCaller) ArrayA(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateArrayEvent.contract.Call(opts, &out, "arrayA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArrayA is a free data retrieval call binding the contract method 0xad3149a4.
//
// Solidity: function arrayA(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventSession) ArrayA(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.ArrayA(&_PublicVariableEthDBUpdateArrayEvent.CallOpts, arg0)
}

// ArrayA is a free data retrieval call binding the contract method 0xad3149a4.
//
// Solidity: function arrayA(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventCallerSession) ArrayA(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.ArrayA(&_PublicVariableEthDBUpdateArrayEvent.CallOpts, arg0)
}

// ArrayB is a free data retrieval call binding the contract method 0x9b889676.
//
// Solidity: function arrayB(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventCaller) ArrayB(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateArrayEvent.contract.Call(opts, &out, "arrayB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArrayB is a free data retrieval call binding the contract method 0x9b889676.
//
// Solidity: function arrayB(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventSession) ArrayB(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.ArrayB(&_PublicVariableEthDBUpdateArrayEvent.CallOpts, arg0)
}

// ArrayB is a free data retrieval call binding the contract method 0x9b889676.
//
// Solidity: function arrayB(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventCallerSession) ArrayB(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.ArrayB(&_PublicVariableEthDBUpdateArrayEvent.CallOpts, arg0)
}

// AddToArray is a paid mutator transaction binding the contract method 0xd34e768d.
//
// Solidity: function addToArray() returns()
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventTransactor) AddToArray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.contract.Transact(opts, "addToArray")
}

// AddToArray is a paid mutator transaction binding the contract method 0xd34e768d.
//
// Solidity: function addToArray() returns()
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventSession) AddToArray() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.AddToArray(&_PublicVariableEthDBUpdateArrayEvent.TransactOpts)
}

// AddToArray is a paid mutator transaction binding the contract method 0xd34e768d.
//
// Solidity: function addToArray() returns()
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventTransactorSession) AddToArray() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateArrayEvent.Contract.AddToArray(&_PublicVariableEthDBUpdateArrayEvent.TransactOpts)
}

// PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateArrayEvent contract.
type PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateArrayEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateArrayEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateArrayEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateArrayEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateArrayEvent contract.
type PublicVariableEthDBUpdateArrayEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateArrayEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateArrayEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateArrayEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateArrayEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateArrayEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateArrayEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateArrayEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateArrayEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateArrayEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateArrayEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateArrayEvent contract.
type PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateArrayEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateArrayEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateArrayEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateArrayEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateArrayEvent contract.
type PublicVariableEthDBUpdateArrayEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateArrayEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateArrayEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateArrayEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateArrayEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateArrayEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateArrayEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateArrayEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateArrayEvent *PublicVariableEthDBUpdateArrayEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateArrayEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateArrayEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateArrayEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateBoolEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateBoolEvent contract.
var PublicVariableEthDBUpdateBoolEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f805461ffff19166001179055348015601b575f5ffd5b506102ed806100295f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df51611461004257806340a3d246146100675780636af1550514610071575b5f5ffd5b5f5461005390610100900460ff1681565b604051901515815260200160405180910390f35b61006f61007d565b005b5f546100539060ff1681565b5f805461010060ff19821660ff8084161591821783900481161590920261ff001990911661ffff1990931692909217919091179182905560408051808201909152600681526576616c75654160d01b60208201526100dd92909116610110565b6040805180820190915260068152653b30b63ab2a160d11b60208201525f5461010e9190610100900460ff16610110565b565b61013f8282604051602001610129911515815260200190565b6040516020818303038152906040526002610143565b5050565b6040805160018082528183019092525f91816020015b606081526020019060019003908161015957905050905083815f81518110610183576101836101e0565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a81848460058111156101c3576101c36101f4565b6040516101d293929190610236565b60405180910390a150505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b8281101561028d57607f19878603018452610278858351610208565b9450602093840193919091019060010161025c565b5050505082810360208401526102a38186610208565b91505060ff8316604083015294935050505056fea2646970667358221220c2793f819d41d1614e7eaebefa18a6162b7c192f3d6435fafd6e604be5071e6564736f6c634300081c0033",
}

// PublicVariableEthDBUpdateBoolEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateBoolEventMetaData.ABI instead.
var PublicVariableEthDBUpdateBoolEventABI = PublicVariableEthDBUpdateBoolEventMetaData.ABI

// PublicVariableEthDBUpdateBoolEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateBoolEventMetaData.Bin instead.
var PublicVariableEthDBUpdateBoolEventBin = PublicVariableEthDBUpdateBoolEventMetaData.Bin

// DeployPublicVariableEthDBUpdateBoolEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateBoolEvent to it.
func DeployPublicVariableEthDBUpdateBoolEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateBoolEvent, error) {
	parsed, err := PublicVariableEthDBUpdateBoolEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateBoolEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateBoolEvent{PublicVariableEthDBUpdateBoolEventCaller: PublicVariableEthDBUpdateBoolEventCaller{contract: contract}, PublicVariableEthDBUpdateBoolEventTransactor: PublicVariableEthDBUpdateBoolEventTransactor{contract: contract}, PublicVariableEthDBUpdateBoolEventFilterer: PublicVariableEthDBUpdateBoolEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateBoolEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBoolEvent struct {
	PublicVariableEthDBUpdateBoolEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateBoolEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateBoolEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateBoolEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBoolEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateBoolEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBoolEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateBoolEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateBoolEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateBoolEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateBoolEventSession struct {
	Contract     *PublicVariableEthDBUpdateBoolEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateBoolEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateBoolEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateBoolEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// PublicVariableEthDBUpdateBoolEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateBoolEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateBoolEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateBoolEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBoolEventRaw struct {
	Contract *PublicVariableEthDBUpdateBoolEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateBoolEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBoolEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateBoolEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateBoolEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBoolEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateBoolEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateBoolEvent creates a new instance of PublicVariableEthDBUpdateBoolEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBoolEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateBoolEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateBoolEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBoolEvent{PublicVariableEthDBUpdateBoolEventCaller: PublicVariableEthDBUpdateBoolEventCaller{contract: contract}, PublicVariableEthDBUpdateBoolEventTransactor: PublicVariableEthDBUpdateBoolEventTransactor{contract: contract}, PublicVariableEthDBUpdateBoolEventFilterer: PublicVariableEthDBUpdateBoolEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateBoolEventCaller creates a new read-only instance of PublicVariableEthDBUpdateBoolEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBoolEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateBoolEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateBoolEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBoolEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateBoolEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateBoolEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBoolEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateBoolEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateBoolEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBoolEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateBoolEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateBoolEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBoolEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateBoolEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateBoolEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBoolEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateBoolEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateBoolEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateBoolEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.PublicVariableEthDBUpdateBoolEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.PublicVariableEthDBUpdateBoolEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.PublicVariableEthDBUpdateBoolEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(bool)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventCaller) ValueA(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateBoolEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(bool)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventSession) ValueA() (bool, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.ValueA(&_PublicVariableEthDBUpdateBoolEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(bool)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventCallerSession) ValueA() (bool, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.ValueA(&_PublicVariableEthDBUpdateBoolEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(bool)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventCaller) ValueB(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateBoolEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(bool)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventSession) ValueB() (bool, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.ValueB(&_PublicVariableEthDBUpdateBoolEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(bool)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventCallerSession) ValueB() (bool, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.ValueB(&_PublicVariableEthDBUpdateBoolEvent.CallOpts)
}

// Toggle is a paid mutator transaction binding the contract method 0x40a3d246.
//
// Solidity: function toggle() returns()
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventTransactor) Toggle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.contract.Transact(opts, "toggle")
}

// Toggle is a paid mutator transaction binding the contract method 0x40a3d246.
//
// Solidity: function toggle() returns()
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventSession) Toggle() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.Toggle(&_PublicVariableEthDBUpdateBoolEvent.TransactOpts)
}

// Toggle is a paid mutator transaction binding the contract method 0x40a3d246.
//
// Solidity: function toggle() returns()
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventTransactorSession) Toggle() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBoolEvent.Contract.Toggle(&_PublicVariableEthDBUpdateBoolEvent.TransactOpts)
}

// PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateBoolEvent contract.
type PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateBoolEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateBoolEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateBoolEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateBoolEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateBoolEvent contract.
type PublicVariableEthDBUpdateBoolEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBoolEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBoolEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateBoolEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateBoolEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBoolEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateBoolEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateBoolEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateBoolEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateBoolEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateBoolEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateBoolEvent contract.
type PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateBoolEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateBoolEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateBoolEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateBoolEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateBoolEvent contract.
type PublicVariableEthDBUpdateBoolEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBoolEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBoolEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateBoolEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateBoolEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBoolEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateBoolEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateBoolEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateBoolEvent *PublicVariableEthDBUpdateBoolEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateBoolEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateBoolEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateBoolEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateBytesEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateBytesEvent contract.
var PublicVariableEthDBUpdateBytesEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600560809081526448656c6c6f60d81b60a0525f9061002390826100f7565b5060408051808201909152600581526415dbdc9b1960da1b602082015260019061004d90826100f7565b50348015610059575f5ffd5b506101b1565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061008757607f821691505b6020821081036100a557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156100f257805f5260205f20601f840160051c810160208510156100d05750805b601f840160051c820191505b818110156100ef575f81556001016100dc565b50505b505050565b81516001600160401b038111156101105761011061005f565b6101248161011e8454610073565b846100ab565b6020601f821160018114610156575f831561013f5750848201515b5f19600385901b1c1916600184901b1784556100ef565b5f84815260208120601f198516915b828110156101855787850151825560209485019460019092019101610165565b50848210156101a257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6105e8806101be5f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df5161146100425780636af1550514610060578063a2e6204514610068575b5f5ffd5b61004a610072565b6040516100579190610323565b60405180910390f35b61004a6100fe565b61007061010a565b005b6001805461007f9061033c565b80601f01602080910402602001604051908101604052809291908181526020018280546100ab9061033c565b80156100f65780601f106100cd576101008083540402835291602001916100f6565b820191905f5260205f20905b8154815290600101906020018083116100d957829003601f168201915b505050505081565b5f805461007f9061033c565b5f5f60405160200161011c9190610374565b60405160208183030381529060405290505f60016040516020016101409190610374565b60408051601f1981840301815291905290505f61015d838261044e565b50600161016a828261044e565b5061021a6040518060400160405280600681526020016576616c75654160d01b8152505f80546101999061033c565b80601f01602080910402602001604051908101604052809291908181526020018280546101c59061033c565b80156102105780601f106101e757610100808354040283529160200191610210565b820191905f5260205f20905b8154815290600101906020018083116101f357829003601f168201915b505050505061024d565b610249604051806040016040528060068152602001653b30b63ab2a160d11b815250600180546101999061033c565b5050565b60408051600180825281830190925261024991849184916005915f91816020015b606081526020019060019003908161026e57905050905083815f8151811061029857610298610509565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a81848460058111156102d8576102d861051d565b6040516102e793929190610531565b60405180910390a150505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61033560208301846102f5565b9392505050565b600181811c9082168061035057607f821691505b60208210810361036e57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f83546103818161033c565b60018216801561039857600181146103ad576103da565b60ff19831686528115158202860193506103da565b865f5260205f205f5b838110156103d2578154888201526001909101906020016103b6565b505081860193505b5050602160f81b8252506001019392505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561044957805f5260205f20601f840160051c810160208510156104275750805b601f840160051c820191505b81811015610446575f8155600101610433565b50505b505050565b815167ffffffffffffffff811115610468576104686103ee565b61047c81610476845461033c565b84610402565b6020601f8211600181146104ae575f83156104975750848201515b5f19600385901b1c1916600184901b178455610446565b5f84815260208120601f198516915b828110156104dd57878501518255602094850194600190920191016104bd565b50848210156104fa57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b8281101561058857607f198786030184526105738583516102f5565b94506020938401939190910190600101610557565b50505050828103602084015261059e81866102f5565b91505060ff8316604083015294935050505056fea264697066735822122032b5419b7ee314f5d80039f15cbd3038c2be5ed357a4afd2b052a0261773889e64736f6c634300081c0033",
}

// PublicVariableEthDBUpdateBytesEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateBytesEventMetaData.ABI instead.
var PublicVariableEthDBUpdateBytesEventABI = PublicVariableEthDBUpdateBytesEventMetaData.ABI

// PublicVariableEthDBUpdateBytesEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateBytesEventMetaData.Bin instead.
var PublicVariableEthDBUpdateBytesEventBin = PublicVariableEthDBUpdateBytesEventMetaData.Bin

// DeployPublicVariableEthDBUpdateBytesEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateBytesEvent to it.
func DeployPublicVariableEthDBUpdateBytesEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateBytesEvent, error) {
	parsed, err := PublicVariableEthDBUpdateBytesEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateBytesEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateBytesEvent{PublicVariableEthDBUpdateBytesEventCaller: PublicVariableEthDBUpdateBytesEventCaller{contract: contract}, PublicVariableEthDBUpdateBytesEventTransactor: PublicVariableEthDBUpdateBytesEventTransactor{contract: contract}, PublicVariableEthDBUpdateBytesEventFilterer: PublicVariableEthDBUpdateBytesEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateBytesEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBytesEvent struct {
	PublicVariableEthDBUpdateBytesEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateBytesEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateBytesEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateBytesEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBytesEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateBytesEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBytesEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateBytesEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateBytesEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateBytesEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateBytesEventSession struct {
	Contract     *PublicVariableEthDBUpdateBytesEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateBytesEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateBytesEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateBytesEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// PublicVariableEthDBUpdateBytesEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateBytesEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateBytesEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateBytesEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBytesEventRaw struct {
	Contract *PublicVariableEthDBUpdateBytesEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateBytesEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBytesEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateBytesEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateBytesEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateBytesEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateBytesEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateBytesEvent creates a new instance of PublicVariableEthDBUpdateBytesEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBytesEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateBytesEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateBytesEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBytesEvent{PublicVariableEthDBUpdateBytesEventCaller: PublicVariableEthDBUpdateBytesEventCaller{contract: contract}, PublicVariableEthDBUpdateBytesEventTransactor: PublicVariableEthDBUpdateBytesEventTransactor{contract: contract}, PublicVariableEthDBUpdateBytesEventFilterer: PublicVariableEthDBUpdateBytesEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateBytesEventCaller creates a new read-only instance of PublicVariableEthDBUpdateBytesEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBytesEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateBytesEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateBytesEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBytesEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateBytesEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateBytesEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBytesEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateBytesEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateBytesEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBytesEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateBytesEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateBytesEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateBytesEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateBytesEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateBytesEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBytesEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateBytesEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateBytesEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateBytesEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.PublicVariableEthDBUpdateBytesEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.PublicVariableEthDBUpdateBytesEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.PublicVariableEthDBUpdateBytesEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(bytes)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventCaller) ValueA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateBytesEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(bytes)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventSession) ValueA() ([]byte, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.ValueA(&_PublicVariableEthDBUpdateBytesEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(bytes)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventCallerSession) ValueA() ([]byte, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.ValueA(&_PublicVariableEthDBUpdateBytesEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(bytes)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventCaller) ValueB(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateBytesEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(bytes)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventSession) ValueB() ([]byte, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.ValueB(&_PublicVariableEthDBUpdateBytesEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(bytes)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventCallerSession) ValueB() ([]byte, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.ValueB(&_PublicVariableEthDBUpdateBytesEvent.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventSession) Update() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.Update(&_PublicVariableEthDBUpdateBytesEvent.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventTransactorSession) Update() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateBytesEvent.Contract.Update(&_PublicVariableEthDBUpdateBytesEvent.TransactOpts)
}

// PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateBytesEvent contract.
type PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateBytesEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateBytesEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateBytesEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateBytesEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateBytesEvent contract.
type PublicVariableEthDBUpdateBytesEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBytesEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBytesEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateBytesEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateBytesEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBytesEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateBytesEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateBytesEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateBytesEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateBytesEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateBytesEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateBytesEvent contract.
type PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateBytesEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateBytesEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateBytesEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateBytesEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateBytesEvent contract.
type PublicVariableEthDBUpdateBytesEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBytesEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateBytesEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateBytesEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateBytesEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateBytesEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateBytesEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateBytesEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateBytesEvent *PublicVariableEthDBUpdateBytesEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateBytesEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateBytesEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateBytesEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateIntEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateIntEvent contract.
var PublicVariableEthDBUpdateIntEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526063195f55600a6001553480156018575f5ffd5b50610303806100265f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df5161146100425780636af155051461005d578063d09de08a14610065575b5f5ffd5b61004b60015481565b60405190815260200160405180910390f35b61004b5f5481565b61006d61006f565b005b60015f5f82825461008091906101c3565b92505081905550600260015f82825461009991906101c3565b925050819055506100c96040518060400160405280600681526020016576616c75654160d01b8152505f546100f5565b6100f3604051806040016040528060068152602001653b30b63ab2a160d11b8152506001546100f5565b565b610122828260405160200161010c91815260200190565b6040516020818303038152906040526004610126565b5050565b6040805160018082528183019092525f91816020015b606081526020019060019003908161013c57905050905083815f81518110610166576101666101f6565b60200260200101819052507fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a81848460058111156101a6576101a661020a565b6040516101b59392919061024c565b60405180910390a150505050565b8082018281125f8312801582168215821617156101ee57634e487b7160e01b5f52601160045260245ffd5b505092915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b828110156102a357607f1987860301845261028e85835161021e565b94506020938401939190910190600101610272565b5050505082810360208401526102b9818661021e565b91505060ff8316604083015294935050505056fea264697066735822122006968073b02db32b3679ee7eb40b17d3c3c81f35e114e0a4327677609325c4cb64736f6c634300081c0033",
}

// PublicVariableEthDBUpdateIntEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateIntEventMetaData.ABI instead.
var PublicVariableEthDBUpdateIntEventABI = PublicVariableEthDBUpdateIntEventMetaData.ABI

// PublicVariableEthDBUpdateIntEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateIntEventMetaData.Bin instead.
var PublicVariableEthDBUpdateIntEventBin = PublicVariableEthDBUpdateIntEventMetaData.Bin

// DeployPublicVariableEthDBUpdateIntEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateIntEvent to it.
func DeployPublicVariableEthDBUpdateIntEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateIntEvent, error) {
	parsed, err := PublicVariableEthDBUpdateIntEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateIntEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateIntEvent{PublicVariableEthDBUpdateIntEventCaller: PublicVariableEthDBUpdateIntEventCaller{contract: contract}, PublicVariableEthDBUpdateIntEventTransactor: PublicVariableEthDBUpdateIntEventTransactor{contract: contract}, PublicVariableEthDBUpdateIntEventFilterer: PublicVariableEthDBUpdateIntEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateIntEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateIntEvent struct {
	PublicVariableEthDBUpdateIntEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateIntEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateIntEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateIntEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateIntEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateIntEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateIntEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateIntEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateIntEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateIntEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateIntEventSession struct {
	Contract     *PublicVariableEthDBUpdateIntEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateIntEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateIntEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateIntEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// PublicVariableEthDBUpdateIntEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateIntEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateIntEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateIntEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateIntEventRaw struct {
	Contract *PublicVariableEthDBUpdateIntEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateIntEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateIntEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateIntEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateIntEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateIntEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateIntEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateIntEvent creates a new instance of PublicVariableEthDBUpdateIntEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateIntEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateIntEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateIntEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateIntEvent{PublicVariableEthDBUpdateIntEventCaller: PublicVariableEthDBUpdateIntEventCaller{contract: contract}, PublicVariableEthDBUpdateIntEventTransactor: PublicVariableEthDBUpdateIntEventTransactor{contract: contract}, PublicVariableEthDBUpdateIntEventFilterer: PublicVariableEthDBUpdateIntEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateIntEventCaller creates a new read-only instance of PublicVariableEthDBUpdateIntEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateIntEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateIntEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateIntEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateIntEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateIntEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateIntEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateIntEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateIntEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateIntEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateIntEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateIntEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateIntEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateIntEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateIntEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateIntEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateIntEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateIntEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateIntEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateIntEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateIntEvent.Contract.PublicVariableEthDBUpdateIntEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.PublicVariableEthDBUpdateIntEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.PublicVariableEthDBUpdateIntEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateIntEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(int256)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventCaller) ValueA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateIntEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(int256)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.ValueA(&_PublicVariableEthDBUpdateIntEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(int256)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventCallerSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.ValueA(&_PublicVariableEthDBUpdateIntEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(int256)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventCaller) ValueB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateIntEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(int256)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.ValueB(&_PublicVariableEthDBUpdateIntEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(int256)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventCallerSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.ValueB(&_PublicVariableEthDBUpdateIntEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.Increment(&_PublicVariableEthDBUpdateIntEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateIntEvent.Contract.Increment(&_PublicVariableEthDBUpdateIntEvent.TransactOpts)
}

// PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateIntEvent contract.
type PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateIntEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateIntEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateIntEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateIntEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateIntEvent contract.
type PublicVariableEthDBUpdateIntEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateIntEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateIntEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateIntEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateIntEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateIntEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateIntEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateIntEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateIntEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateIntEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateIntEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateIntEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateIntEvent contract.
type PublicVariableEthDBUpdateIntEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateIntEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateIntEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateIntEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateIntEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateIntEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateIntEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateIntEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateIntEvent contract.
type PublicVariableEthDBUpdateIntEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateIntEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateIntEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateIntEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateIntEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateIntEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateIntEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateIntEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateIntEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateIntEvent *PublicVariableEthDBUpdateIntEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateIntEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateIntEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateIntEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateMappingEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateMappingEvent contract.
var PublicVariableEthDBUpdateMappingEventMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mappingA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mappingB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060017fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d81905560027fabbb5caa7dda850e60932de0934eb1f9d0f59695050f761dc64e443e5030a56981905560209190915260037fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f555f5260047fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f556102ae806100b85f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063d09de08a14610043578063e36b791d1461004d578063f49925c21461007e575b5f5ffd5b61004b61009d565b005b61006c61005b36600461023c565b60016020525f908152604090205481565b60405190815260200160405180910390f35b61006c61008c36600461023c565b5f6020819052908152604090205481565b60015f8181526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546100d191610253565b5f60208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d919091556002908190527fabbb5caa7dda850e60932de0934eb1f9d0f59695050f761dc64e443e5030a5695461012d91610253565b7fabbb5caa7dda850e60932de0934eb1f9d0f59695050f761dc64e443e5030a5695560015f8190526020527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f54610185906003610253565b60016020527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f5560025f527fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f546101dd906004610253565b60025f5260016020527fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f55610210610212565b565b6040517fc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c905f90a1565b5f6020828403121561024c575f5ffd5b5035919050565b8082018082111561027257634e487b7160e01b5f52601160045260245ffd5b9291505056fea26469706673582212204f3a9c2b9e46598b4b20c27469279ec535df237028ac8e72473ddbb0ccfe39f164736f6c634300081c0033",
}

// PublicVariableEthDBUpdateMappingEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateMappingEventMetaData.ABI instead.
var PublicVariableEthDBUpdateMappingEventABI = PublicVariableEthDBUpdateMappingEventMetaData.ABI

// PublicVariableEthDBUpdateMappingEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateMappingEventMetaData.Bin instead.
var PublicVariableEthDBUpdateMappingEventBin = PublicVariableEthDBUpdateMappingEventMetaData.Bin

// DeployPublicVariableEthDBUpdateMappingEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateMappingEvent to it.
func DeployPublicVariableEthDBUpdateMappingEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateMappingEvent, error) {
	parsed, err := PublicVariableEthDBUpdateMappingEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateMappingEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateMappingEvent{PublicVariableEthDBUpdateMappingEventCaller: PublicVariableEthDBUpdateMappingEventCaller{contract: contract}, PublicVariableEthDBUpdateMappingEventTransactor: PublicVariableEthDBUpdateMappingEventTransactor{contract: contract}, PublicVariableEthDBUpdateMappingEventFilterer: PublicVariableEthDBUpdateMappingEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateMappingEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMappingEvent struct {
	PublicVariableEthDBUpdateMappingEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateMappingEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateMappingEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateMappingEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMappingEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateMappingEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMappingEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateMappingEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateMappingEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateMappingEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateMappingEventSession struct {
	Contract     *PublicVariableEthDBUpdateMappingEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateMappingEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateMappingEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateMappingEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// PublicVariableEthDBUpdateMappingEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateMappingEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateMappingEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateMappingEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMappingEventRaw struct {
	Contract *PublicVariableEthDBUpdateMappingEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateMappingEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMappingEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateMappingEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateMappingEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMappingEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateMappingEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateMappingEvent creates a new instance of PublicVariableEthDBUpdateMappingEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMappingEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateMappingEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateMappingEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMappingEvent{PublicVariableEthDBUpdateMappingEventCaller: PublicVariableEthDBUpdateMappingEventCaller{contract: contract}, PublicVariableEthDBUpdateMappingEventTransactor: PublicVariableEthDBUpdateMappingEventTransactor{contract: contract}, PublicVariableEthDBUpdateMappingEventFilterer: PublicVariableEthDBUpdateMappingEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateMappingEventCaller creates a new read-only instance of PublicVariableEthDBUpdateMappingEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMappingEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateMappingEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateMappingEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMappingEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateMappingEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateMappingEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMappingEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateMappingEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateMappingEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMappingEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateMappingEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateMappingEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMappingEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateMappingEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateMappingEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMappingEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateMappingEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateMappingEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateMappingEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.PublicVariableEthDBUpdateMappingEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.PublicVariableEthDBUpdateMappingEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.PublicVariableEthDBUpdateMappingEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.contract.Transact(opts, method, params...)
}

// MappingA is a free data retrieval call binding the contract method 0xf49925c2.
//
// Solidity: function mappingA(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventCaller) MappingA(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateMappingEvent.contract.Call(opts, &out, "mappingA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappingA is a free data retrieval call binding the contract method 0xf49925c2.
//
// Solidity: function mappingA(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventSession) MappingA(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.MappingA(&_PublicVariableEthDBUpdateMappingEvent.CallOpts, arg0)
}

// MappingA is a free data retrieval call binding the contract method 0xf49925c2.
//
// Solidity: function mappingA(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventCallerSession) MappingA(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.MappingA(&_PublicVariableEthDBUpdateMappingEvent.CallOpts, arg0)
}

// MappingB is a free data retrieval call binding the contract method 0xe36b791d.
//
// Solidity: function mappingB(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventCaller) MappingB(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateMappingEvent.contract.Call(opts, &out, "mappingB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappingB is a free data retrieval call binding the contract method 0xe36b791d.
//
// Solidity: function mappingB(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventSession) MappingB(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.MappingB(&_PublicVariableEthDBUpdateMappingEvent.CallOpts, arg0)
}

// MappingB is a free data retrieval call binding the contract method 0xe36b791d.
//
// Solidity: function mappingB(uint256 ) view returns(uint256)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventCallerSession) MappingB(arg0 *big.Int) (*big.Int, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.MappingB(&_PublicVariableEthDBUpdateMappingEvent.CallOpts, arg0)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.Increment(&_PublicVariableEthDBUpdateMappingEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMappingEvent.Contract.Increment(&_PublicVariableEthDBUpdateMappingEvent.TransactOpts)
}

// PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateMappingEvent contract.
type PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateMappingEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateMappingEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateMappingEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateMappingEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateMappingEvent contract.
type PublicVariableEthDBUpdateMappingEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMappingEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMappingEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateMappingEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateMappingEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMappingEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateMappingEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateMappingEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateMappingEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateMappingEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateMappingEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateMappingEvent contract.
type PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateMappingEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateMappingEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateMappingEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateMappingEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateMappingEvent contract.
type PublicVariableEthDBUpdateMappingEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMappingEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMappingEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateMappingEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateMappingEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMappingEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateMappingEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateMappingEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateMappingEvent *PublicVariableEthDBUpdateMappingEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateMappingEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateMappingEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateMappingEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateMultipleEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateMultipleEvent contract.
var PublicVariableEthDBUpdateMultipleEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f55600a6001553480156017575f5ffd5b50610166806100255f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df5161146100425780636af155051461005d578063d09de08a14610065575b5f5ffd5b61004b60015481565b60405190815260200160405180910390f35b61004b5f5481565b61006d61006f565b005b60015f5f828254610080919061010b565b92505081905550600260015f828254610099919061010b565b909155506100a790506100e1565b60015f5f8282546100b8919061010b565b92505081905550600260015f8282546100d1919061010b565b909155506100df90506100e1565b565b6040517fc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c905f90a1565b8082018082111561012a57634e487b7160e01b5f52601160045260245ffd5b9291505056fea26469706673582212206b1ea075361770b54f1993a413b61dbfef29e938b231a04fba2257f3187058d964736f6c634300081c0033",
}

// PublicVariableEthDBUpdateMultipleEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateMultipleEventMetaData.ABI instead.
var PublicVariableEthDBUpdateMultipleEventABI = PublicVariableEthDBUpdateMultipleEventMetaData.ABI

// PublicVariableEthDBUpdateMultipleEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateMultipleEventMetaData.Bin instead.
var PublicVariableEthDBUpdateMultipleEventBin = PublicVariableEthDBUpdateMultipleEventMetaData.Bin

// DeployPublicVariableEthDBUpdateMultipleEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateMultipleEvent to it.
func DeployPublicVariableEthDBUpdateMultipleEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateMultipleEvent, error) {
	parsed, err := PublicVariableEthDBUpdateMultipleEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateMultipleEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateMultipleEvent{PublicVariableEthDBUpdateMultipleEventCaller: PublicVariableEthDBUpdateMultipleEventCaller{contract: contract}, PublicVariableEthDBUpdateMultipleEventTransactor: PublicVariableEthDBUpdateMultipleEventTransactor{contract: contract}, PublicVariableEthDBUpdateMultipleEventFilterer: PublicVariableEthDBUpdateMultipleEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateMultipleEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMultipleEvent struct {
	PublicVariableEthDBUpdateMultipleEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateMultipleEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateMultipleEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateMultipleEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMultipleEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateMultipleEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMultipleEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateMultipleEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateMultipleEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateMultipleEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateMultipleEventSession struct {
	Contract     *PublicVariableEthDBUpdateMultipleEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                           // Call options to use throughout this session
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateMultipleEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateMultipleEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateMultipleEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                 // Call options to use throughout this session
}

// PublicVariableEthDBUpdateMultipleEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateMultipleEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateMultipleEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                 // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateMultipleEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMultipleEventRaw struct {
	Contract *PublicVariableEthDBUpdateMultipleEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateMultipleEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMultipleEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateMultipleEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateMultipleEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateMultipleEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateMultipleEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateMultipleEvent creates a new instance of PublicVariableEthDBUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMultipleEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateMultipleEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateMultipleEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMultipleEvent{PublicVariableEthDBUpdateMultipleEventCaller: PublicVariableEthDBUpdateMultipleEventCaller{contract: contract}, PublicVariableEthDBUpdateMultipleEventTransactor: PublicVariableEthDBUpdateMultipleEventTransactor{contract: contract}, PublicVariableEthDBUpdateMultipleEventFilterer: PublicVariableEthDBUpdateMultipleEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateMultipleEventCaller creates a new read-only instance of PublicVariableEthDBUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMultipleEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateMultipleEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateMultipleEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMultipleEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateMultipleEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMultipleEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateMultipleEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateMultipleEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMultipleEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateMultipleEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateMultipleEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateMultipleEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateMultipleEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateMultipleEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMultipleEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateMultipleEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateMultipleEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateMultipleEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.PublicVariableEthDBUpdateMultipleEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.PublicVariableEthDBUpdateMultipleEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.PublicVariableEthDBUpdateMultipleEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventCaller) ValueA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateMultipleEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.ValueA(&_PublicVariableEthDBUpdateMultipleEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventCallerSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.ValueA(&_PublicVariableEthDBUpdateMultipleEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventCaller) ValueB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateMultipleEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.ValueB(&_PublicVariableEthDBUpdateMultipleEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventCallerSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.ValueB(&_PublicVariableEthDBUpdateMultipleEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.Increment(&_PublicVariableEthDBUpdateMultipleEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateMultipleEvent.Contract.Increment(&_PublicVariableEthDBUpdateMultipleEvent.TransactOpts)
}

// PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateMultipleEvent contract.
type PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateMultipleEvent contract.
type PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMultipleEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMultipleEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateMultipleEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMultipleEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateMultipleEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateMultipleEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateMultipleEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateMultipleEvent contract.
type PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateMultipleEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateMultipleEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateMultipleEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateMultipleEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateMultipleEvent contract.
type PublicVariableEthDBUpdateMultipleEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMultipleEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateMultipleEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateMultipleEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateMultipleEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateMultipleEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateMultipleEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateMultipleEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateMultipleEvent *PublicVariableEthDBUpdateMultipleEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateMultipleEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateMultipleEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateMultipleEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateStringEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateStringEvent contract.
var PublicVariableEthDBUpdateStringEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600560809081526448656c6c6f60d81b60a0525f9061002390826100f7565b5060408051808201909152600581526415dbdc9b1960da1b602082015260019061004d90826100f7565b50348015610059575f5ffd5b506101b1565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061008757607f821691505b6020821081036100a557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156100f257805f5260205f20601f840160051c810160208510156100d05750805b601f840160051c820191505b818110156100ef575f81556001016100dc565b50505b505050565b81516001600160401b038111156101105761011061005f565b6101248161011e8454610073565b846100ab565b6020601f821160018114610156575f831561013f5750848201515b5f19600385901b1c1916600184901b1784556100ef565b5f84815260208120601f198516915b828110156101855787850151825560209485019460019092019101610165565b50848210156101a257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6103d9806101be5f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062df5161146100425780636af1550514610060578063a2e6204514610068575b5f5ffd5b61004a610072565b60405161005791906101a1565b60405180910390f35b61004a6100fe565b61007061010a565b005b6001805461007f906101d6565b80601f01602080910402602001604051908101604052809291908181526020018280546100ab906101d6565b80156100f65780601f106100cd576101008083540402835291602001916100f6565b820191905f5260205f20905b8154815290600101906020018083116100d957829003601f168201915b505050505081565b5f805461007f906101d6565b5f5f60405160200161011c919061020e565b60405160208183030381529060405290505f6001604051602001610140919061020e565b60408051601f1981840301815291905290505f61015d83826102e8565b50600161016a82826102e8565b50610173610177565b5050565b6040517fc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c905f90a1565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b600181811c908216806101ea57607f821691505b60208210810361020857634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f835461021b816101d6565b600182168015610232576001811461024757610274565b60ff1983168652811515820286019350610274565b865f5260205f205f5b8381101561026c57815488820152600190910190602001610250565b505081860193505b5050602160f81b8252506001019392505050565b634e487b7160e01b5f52604160045260245ffd5b601f8211156102e357805f5260205f20601f840160051c810160208510156102c15750805b601f840160051c820191505b818110156102e0575f81556001016102cd565b50505b505050565b815167ffffffffffffffff81111561030257610302610288565b6103168161031084546101d6565b8461029c565b6020601f821160018114610348575f83156103315750848201515b5f19600385901b1c1916600184901b1784556102e0565b5f84815260208120601f198516915b828110156103775787850151825560209485019460019092019101610357565b508482101561039457868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea2646970667358221220c3e3bca2dce77fc677d08f6c0bf859682da53162dfdbe0674373780f195b08f764736f6c634300081c0033",
}

// PublicVariableEthDBUpdateStringEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateStringEventMetaData.ABI instead.
var PublicVariableEthDBUpdateStringEventABI = PublicVariableEthDBUpdateStringEventMetaData.ABI

// PublicVariableEthDBUpdateStringEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateStringEventMetaData.Bin instead.
var PublicVariableEthDBUpdateStringEventBin = PublicVariableEthDBUpdateStringEventMetaData.Bin

// DeployPublicVariableEthDBUpdateStringEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateStringEvent to it.
func DeployPublicVariableEthDBUpdateStringEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateStringEvent, error) {
	parsed, err := PublicVariableEthDBUpdateStringEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateStringEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateStringEvent{PublicVariableEthDBUpdateStringEventCaller: PublicVariableEthDBUpdateStringEventCaller{contract: contract}, PublicVariableEthDBUpdateStringEventTransactor: PublicVariableEthDBUpdateStringEventTransactor{contract: contract}, PublicVariableEthDBUpdateStringEventFilterer: PublicVariableEthDBUpdateStringEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateStringEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStringEvent struct {
	PublicVariableEthDBUpdateStringEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateStringEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateStringEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateStringEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStringEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateStringEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStringEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateStringEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateStringEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateStringEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateStringEventSession struct {
	Contract     *PublicVariableEthDBUpdateStringEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateStringEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateStringEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateStringEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// PublicVariableEthDBUpdateStringEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateStringEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateStringEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateStringEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStringEventRaw struct {
	Contract *PublicVariableEthDBUpdateStringEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateStringEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStringEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateStringEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateStringEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStringEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateStringEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateStringEvent creates a new instance of PublicVariableEthDBUpdateStringEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStringEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateStringEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateStringEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStringEvent{PublicVariableEthDBUpdateStringEventCaller: PublicVariableEthDBUpdateStringEventCaller{contract: contract}, PublicVariableEthDBUpdateStringEventTransactor: PublicVariableEthDBUpdateStringEventTransactor{contract: contract}, PublicVariableEthDBUpdateStringEventFilterer: PublicVariableEthDBUpdateStringEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateStringEventCaller creates a new read-only instance of PublicVariableEthDBUpdateStringEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStringEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateStringEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateStringEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStringEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateStringEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateStringEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStringEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateStringEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateStringEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStringEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateStringEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateStringEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStringEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateStringEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateStringEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStringEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateStringEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateStringEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateStringEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateStringEvent.Contract.PublicVariableEthDBUpdateStringEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.PublicVariableEthDBUpdateStringEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.PublicVariableEthDBUpdateStringEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateStringEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(string)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventCaller) ValueA(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateStringEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(string)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventSession) ValueA() (string, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.ValueA(&_PublicVariableEthDBUpdateStringEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(string)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventCallerSession) ValueA() (string, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.ValueA(&_PublicVariableEthDBUpdateStringEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(string)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventCaller) ValueB(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateStringEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(string)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventSession) ValueB() (string, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.ValueB(&_PublicVariableEthDBUpdateStringEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(string)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventCallerSession) ValueB() (string, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.ValueB(&_PublicVariableEthDBUpdateStringEvent.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventSession) Update() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.Update(&_PublicVariableEthDBUpdateStringEvent.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventTransactorSession) Update() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStringEvent.Contract.Update(&_PublicVariableEthDBUpdateStringEvent.TransactOpts)
}

// PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateStringEvent contract.
type PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateStringEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateStringEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateStringEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateStringEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateStringEvent contract.
type PublicVariableEthDBUpdateStringEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStringEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStringEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateStringEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateStringEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStringEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateStringEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateStringEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateStringEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateStringEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateStringEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateStringEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateStringEvent contract.
type PublicVariableEthDBUpdateStringEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateStringEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateStringEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateStringEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateStringEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateStringEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateStringEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateStringEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateStringEvent contract.
type PublicVariableEthDBUpdateStringEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateStringEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStringEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStringEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateStringEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateStringEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStringEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateStringEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateStringEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateStringEvent *PublicVariableEthDBUpdateStringEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateStringEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateStringEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateStringEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateStructEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateStructEvent contract.
var PublicVariableEthDBUpdateStructEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"structA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"structB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101778061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806311ba731f14610043578063d09de08a14610069578063d86387a514610073575b5f5ffd5b5f54600154610050919082565b6040805192835260208301919091520160405180910390f35b610071610081565b005b600254600354610050919082565b60015f5f015f828254610094919061011c565b909155505060018054600291905f906100ae90849061011c565b909155505060028054600391905f906100c890849061011c565b909155505060038054600491905f906100e290849061011c565b909155506100f090506100f2565b565b6040517fc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c905f90a1565b8082018082111561013b57634e487b7160e01b5f52601160045260245ffd5b9291505056fea26469706673582212201e80858a18edd2f329d53b1cf6012e2adaacc9af9922a7b71032fabe1fd9c05d64736f6c634300081c0033",
}

// PublicVariableEthDBUpdateStructEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateStructEventMetaData.ABI instead.
var PublicVariableEthDBUpdateStructEventABI = PublicVariableEthDBUpdateStructEventMetaData.ABI

// PublicVariableEthDBUpdateStructEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateStructEventMetaData.Bin instead.
var PublicVariableEthDBUpdateStructEventBin = PublicVariableEthDBUpdateStructEventMetaData.Bin

// DeployPublicVariableEthDBUpdateStructEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateStructEvent to it.
func DeployPublicVariableEthDBUpdateStructEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateStructEvent, error) {
	parsed, err := PublicVariableEthDBUpdateStructEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateStructEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateStructEvent{PublicVariableEthDBUpdateStructEventCaller: PublicVariableEthDBUpdateStructEventCaller{contract: contract}, PublicVariableEthDBUpdateStructEventTransactor: PublicVariableEthDBUpdateStructEventTransactor{contract: contract}, PublicVariableEthDBUpdateStructEventFilterer: PublicVariableEthDBUpdateStructEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateStructEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStructEvent struct {
	PublicVariableEthDBUpdateStructEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateStructEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateStructEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateStructEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStructEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateStructEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStructEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateStructEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateStructEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateStructEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateStructEventSession struct {
	Contract     *PublicVariableEthDBUpdateStructEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateStructEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateStructEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateStructEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// PublicVariableEthDBUpdateStructEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateStructEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateStructEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateStructEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStructEventRaw struct {
	Contract *PublicVariableEthDBUpdateStructEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateStructEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStructEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateStructEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateStructEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateStructEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateStructEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateStructEvent creates a new instance of PublicVariableEthDBUpdateStructEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStructEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateStructEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateStructEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStructEvent{PublicVariableEthDBUpdateStructEventCaller: PublicVariableEthDBUpdateStructEventCaller{contract: contract}, PublicVariableEthDBUpdateStructEventTransactor: PublicVariableEthDBUpdateStructEventTransactor{contract: contract}, PublicVariableEthDBUpdateStructEventFilterer: PublicVariableEthDBUpdateStructEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateStructEventCaller creates a new read-only instance of PublicVariableEthDBUpdateStructEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStructEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateStructEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateStructEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStructEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateStructEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateStructEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStructEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateStructEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateStructEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStructEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateStructEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateStructEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateStructEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateStructEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateStructEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStructEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateStructEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateStructEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateStructEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateStructEvent.Contract.PublicVariableEthDBUpdateStructEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.PublicVariableEthDBUpdateStructEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.PublicVariableEthDBUpdateStructEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateStructEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.contract.Transact(opts, method, params...)
}

// StructA is a free data retrieval call binding the contract method 0x11ba731f.
//
// Solidity: function structA() view returns(uint256 valueA, uint256 valueB)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventCaller) StructA(opts *bind.CallOpts) (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateStructEvent.contract.Call(opts, &out, "structA")

	outstruct := new(struct {
		ValueA *big.Int
		ValueB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValueA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StructA is a free data retrieval call binding the contract method 0x11ba731f.
//
// Solidity: function structA() view returns(uint256 valueA, uint256 valueB)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventSession) StructA() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.StructA(&_PublicVariableEthDBUpdateStructEvent.CallOpts)
}

// StructA is a free data retrieval call binding the contract method 0x11ba731f.
//
// Solidity: function structA() view returns(uint256 valueA, uint256 valueB)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventCallerSession) StructA() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.StructA(&_PublicVariableEthDBUpdateStructEvent.CallOpts)
}

// StructB is a free data retrieval call binding the contract method 0xd86387a5.
//
// Solidity: function structB() view returns(uint256 valueA, uint256 valueB)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventCaller) StructB(opts *bind.CallOpts) (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateStructEvent.contract.Call(opts, &out, "structB")

	outstruct := new(struct {
		ValueA *big.Int
		ValueB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValueA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StructB is a free data retrieval call binding the contract method 0xd86387a5.
//
// Solidity: function structB() view returns(uint256 valueA, uint256 valueB)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventSession) StructB() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.StructB(&_PublicVariableEthDBUpdateStructEvent.CallOpts)
}

// StructB is a free data retrieval call binding the contract method 0xd86387a5.
//
// Solidity: function structB() view returns(uint256 valueA, uint256 valueB)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventCallerSession) StructB() (struct {
	ValueA *big.Int
	ValueB *big.Int
}, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.StructB(&_PublicVariableEthDBUpdateStructEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.Increment(&_PublicVariableEthDBUpdateStructEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateStructEvent.Contract.Increment(&_PublicVariableEthDBUpdateStructEvent.TransactOpts)
}

// PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateStructEvent contract.
type PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateStructEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateStructEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateStructEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateStructEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateStructEvent contract.
type PublicVariableEthDBUpdateStructEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStructEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStructEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateStructEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateStructEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStructEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateStructEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateStructEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateStructEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateStructEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateStructEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateStructEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateStructEvent contract.
type PublicVariableEthDBUpdateStructEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateStructEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateStructEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateStructEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateStructEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateStructEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateStructEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateStructEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateStructEvent contract.
type PublicVariableEthDBUpdateStructEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateStructEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStructEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateStructEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateStructEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateStructEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateStructEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateStructEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateStructEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateStructEvent *PublicVariableEthDBUpdateStructEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateStructEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateStructEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateStructEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateUintEventMetaData contains all meta data concerning the PublicVariableEthDBUpdateUintEvent contract.
var PublicVariableEthDBUpdateUintEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthDBUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f55600a6001553480156017575f5ffd5b5061011e806100255f395ff3fe6080604052348015600e575f5ffd5b50600436106039575f3560e01c8062df516114603d5780636af15505146057578063d09de08a14605e575b5f5ffd5b604560015481565b60405190815260200160405180910390f35b60455f5481565b60646066565b005b60015f5f8282546075919060c4565b92505081905550600260015f828254608c919060c4565b9091555060989050609a565b565b6040517fc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c905f90a1565b8082018082111560e257634e487b7160e01b5f52601160045260245ffd5b9291505056fea2646970667358221220d050c4b5e63db2068c6ce40857881474ab6122d09881582d4f9b44d8392c3afe64736f6c634300081c0033",
}

// PublicVariableEthDBUpdateUintEventABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicVariableEthDBUpdateUintEventMetaData.ABI instead.
var PublicVariableEthDBUpdateUintEventABI = PublicVariableEthDBUpdateUintEventMetaData.ABI

// PublicVariableEthDBUpdateUintEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicVariableEthDBUpdateUintEventMetaData.Bin instead.
var PublicVariableEthDBUpdateUintEventBin = PublicVariableEthDBUpdateUintEventMetaData.Bin

// DeployPublicVariableEthDBUpdateUintEvent deploys a new Ethereum contract, binding an instance of PublicVariableEthDBUpdateUintEvent to it.
func DeployPublicVariableEthDBUpdateUintEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicVariableEthDBUpdateUintEvent, error) {
	parsed, err := PublicVariableEthDBUpdateUintEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicVariableEthDBUpdateUintEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicVariableEthDBUpdateUintEvent{PublicVariableEthDBUpdateUintEventCaller: PublicVariableEthDBUpdateUintEventCaller{contract: contract}, PublicVariableEthDBUpdateUintEventTransactor: PublicVariableEthDBUpdateUintEventTransactor{contract: contract}, PublicVariableEthDBUpdateUintEventFilterer: PublicVariableEthDBUpdateUintEventFilterer{contract: contract}}, nil
}

// PublicVariableEthDBUpdateUintEvent is an auto generated Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateUintEvent struct {
	PublicVariableEthDBUpdateUintEventCaller     // Read-only binding to the contract
	PublicVariableEthDBUpdateUintEventTransactor // Write-only binding to the contract
	PublicVariableEthDBUpdateUintEventFilterer   // Log filterer for contract events
}

// PublicVariableEthDBUpdateUintEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateUintEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateUintEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateUintEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateUintEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicVariableEthDBUpdateUintEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicVariableEthDBUpdateUintEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicVariableEthDBUpdateUintEventSession struct {
	Contract     *PublicVariableEthDBUpdateUintEvent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateUintEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicVariableEthDBUpdateUintEventCallerSession struct {
	Contract *PublicVariableEthDBUpdateUintEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// PublicVariableEthDBUpdateUintEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicVariableEthDBUpdateUintEventTransactorSession struct {
	Contract     *PublicVariableEthDBUpdateUintEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// PublicVariableEthDBUpdateUintEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateUintEventRaw struct {
	Contract *PublicVariableEthDBUpdateUintEvent // Generic contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateUintEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateUintEventCallerRaw struct {
	Contract *PublicVariableEthDBUpdateUintEventCaller // Generic read-only contract binding to access the raw methods on
}

// PublicVariableEthDBUpdateUintEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicVariableEthDBUpdateUintEventTransactorRaw struct {
	Contract *PublicVariableEthDBUpdateUintEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicVariableEthDBUpdateUintEvent creates a new instance of PublicVariableEthDBUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateUintEvent(address common.Address, backend bind.ContractBackend) (*PublicVariableEthDBUpdateUintEvent, error) {
	contract, err := bindPublicVariableEthDBUpdateUintEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateUintEvent{PublicVariableEthDBUpdateUintEventCaller: PublicVariableEthDBUpdateUintEventCaller{contract: contract}, PublicVariableEthDBUpdateUintEventTransactor: PublicVariableEthDBUpdateUintEventTransactor{contract: contract}, PublicVariableEthDBUpdateUintEventFilterer: PublicVariableEthDBUpdateUintEventFilterer{contract: contract}}, nil
}

// NewPublicVariableEthDBUpdateUintEventCaller creates a new read-only instance of PublicVariableEthDBUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateUintEventCaller(address common.Address, caller bind.ContractCaller) (*PublicVariableEthDBUpdateUintEventCaller, error) {
	contract, err := bindPublicVariableEthDBUpdateUintEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateUintEventCaller{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateUintEventTransactor creates a new write-only instance of PublicVariableEthDBUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateUintEventTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicVariableEthDBUpdateUintEventTransactor, error) {
	contract, err := bindPublicVariableEthDBUpdateUintEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateUintEventTransactor{contract: contract}, nil
}

// NewPublicVariableEthDBUpdateUintEventFilterer creates a new log filterer instance of PublicVariableEthDBUpdateUintEvent, bound to a specific deployed contract.
func NewPublicVariableEthDBUpdateUintEventFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicVariableEthDBUpdateUintEventFilterer, error) {
	contract, err := bindPublicVariableEthDBUpdateUintEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateUintEventFilterer{contract: contract}, nil
}

// bindPublicVariableEthDBUpdateUintEvent binds a generic wrapper to an already deployed contract.
func bindPublicVariableEthDBUpdateUintEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicVariableEthDBUpdateUintEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateUintEvent.Contract.PublicVariableEthDBUpdateUintEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.PublicVariableEthDBUpdateUintEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.PublicVariableEthDBUpdateUintEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicVariableEthDBUpdateUintEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.contract.Transact(opts, method, params...)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventCaller) ValueA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateUintEvent.contract.Call(opts, &out, "valueA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.ValueA(&_PublicVariableEthDBUpdateUintEvent.CallOpts)
}

// ValueA is a free data retrieval call binding the contract method 0x6af15505.
//
// Solidity: function valueA() view returns(uint256)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventCallerSession) ValueA() (*big.Int, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.ValueA(&_PublicVariableEthDBUpdateUintEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventCaller) ValueB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicVariableEthDBUpdateUintEvent.contract.Call(opts, &out, "valueB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.ValueB(&_PublicVariableEthDBUpdateUintEvent.CallOpts)
}

// ValueB is a free data retrieval call binding the contract method 0x00df5161.
//
// Solidity: function valueB() view returns(uint256)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventCallerSession) ValueB() (*big.Int, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.ValueB(&_PublicVariableEthDBUpdateUintEvent.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.Increment(&_PublicVariableEthDBUpdateUintEvent.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventTransactorSession) Increment() (*types.Transaction, error) {
	return _PublicVariableEthDBUpdateUintEvent.Contract.Increment(&_PublicVariableEthDBUpdateUintEvent.TransactOpts)
}

// PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the PublicVariableEthDBUpdateUintEvent contract.
type PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator struct {
	Event *PublicVariableEthDBUpdateUintEventEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateUintEventEthDBPathUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateUintEventEthDBPathUpdate)
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
func (it *PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateUintEventEthDBPathUpdate represents a EthDBPathUpdate event raised by the PublicVariableEthDBUpdateUintEvent contract.
type PublicVariableEthDBUpdateUintEventEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateUintEvent.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateUintEventEthDBPathUpdateIterator{contract: _PublicVariableEthDBUpdateUintEvent.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateUintEventEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateUintEvent.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateUintEventEthDBPathUpdate)
				if err := _PublicVariableEthDBUpdateUintEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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

// ParseEthDBPathUpdate is a log parse operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventFilterer) ParseEthDBPathUpdate(log types.Log) (*PublicVariableEthDBUpdateUintEventEthDBPathUpdate, error) {
	event := new(PublicVariableEthDBUpdateUintEventEthDBPathUpdate)
	if err := _PublicVariableEthDBUpdateUintEvent.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicVariableEthDBUpdateUintEventEthDBUpdateIterator is returned from FilterEthDBUpdate and is used to iterate over the raw logs and unpacked data for EthDBUpdate events raised by the PublicVariableEthDBUpdateUintEvent contract.
type PublicVariableEthDBUpdateUintEventEthDBUpdateIterator struct {
	Event *PublicVariableEthDBUpdateUintEventEthDBUpdate // Event containing the contract specifics and raw log

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
func (it *PublicVariableEthDBUpdateUintEventEthDBUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicVariableEthDBUpdateUintEventEthDBUpdate)
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
		it.Event = new(PublicVariableEthDBUpdateUintEventEthDBUpdate)
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
func (it *PublicVariableEthDBUpdateUintEventEthDBUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicVariableEthDBUpdateUintEventEthDBUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicVariableEthDBUpdateUintEventEthDBUpdate represents a EthDBUpdate event raised by the PublicVariableEthDBUpdateUintEvent contract.
type PublicVariableEthDBUpdateUintEventEthDBUpdate struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthDBUpdate is a free log retrieval operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventFilterer) FilterEthDBUpdate(opts *bind.FilterOpts) (*PublicVariableEthDBUpdateUintEventEthDBUpdateIterator, error) {

	logs, sub, err := _PublicVariableEthDBUpdateUintEvent.contract.FilterLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return &PublicVariableEthDBUpdateUintEventEthDBUpdateIterator{contract: _PublicVariableEthDBUpdateUintEvent.contract, event: "EthDBUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBUpdate is a free log subscription operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventFilterer) WatchEthDBUpdate(opts *bind.WatchOpts, sink chan<- *PublicVariableEthDBUpdateUintEventEthDBUpdate) (event.Subscription, error) {

	logs, sub, err := _PublicVariableEthDBUpdateUintEvent.contract.WatchLogs(opts, "EthDBUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicVariableEthDBUpdateUintEventEthDBUpdate)
				if err := _PublicVariableEthDBUpdateUintEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
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

// ParseEthDBUpdate is a log parse operation binding the contract event 0xc4742ae10756259b50fcb91ab47eb26886137843557fd3731b4e43dbeb0fc19c.
//
// Solidity: event EthDBUpdate()
func (_PublicVariableEthDBUpdateUintEvent *PublicVariableEthDBUpdateUintEventFilterer) ParseEthDBUpdate(log types.Log) (*PublicVariableEthDBUpdateUintEventEthDBUpdate, error) {
	event := new(PublicVariableEthDBUpdateUintEventEthDBUpdate)
	if err := _PublicVariableEthDBUpdateUintEvent.contract.UnpackLog(event, "EthDBUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlMetaData contains all meta data concerning the RoleControl contract.
var RoleControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IdentityAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIdentity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"}]",
	Bin: "0x610e16610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006b575f3560e01c806333c7c7ec1461006f57806334f4a00c146100985780636215b67a146100b95780639b2642df146100d8578063f76888e7146100f8578063fdde69b814610117575b5f5ffd5b61008261007d366004610ace565b610136565b60405161008f9190610ae5565b60405180910390f35b8180156100a3575f5ffd5b506100b76100b2366004610b59565b61019a565b005b8180156100c4575f5ffd5b506100b76100d3366004610b59565b610207565b6100eb6100e6366004610ace565b61026d565b60405161008f9190610ba6565b818015610103575f5ffd5b506100b7610112366004610c6f565b6103dd565b818015610122575f5ffd5b506100b7610131366004610cad565b610448565b60608160010180548060200260200160405190810160405280929190818152602001828054801561018e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610170575b50505050509050919050565b335f90815260208490526040902054839060029060ff166101ce5760405163c597feeb60e01b815260040160405180910390fd5b6101d98233836104ae565b6101f5576040516282b42960e81b815260040160405180910390fd5b61020085858561054b565b5050505050565b335f90815260208490526040902054839060029060ff1661023b5760405163c597feeb60e01b815260040160405180910390fd5b6102468233836104ae565b610262576040516282b42960e81b815260040160405180910390fd5b61020085858561071d565b60018101546060905f9067ffffffffffffffff81111561028f5761028f610c99565b6040519080825280602002602001820160405280156102d457816020015b604080518082019091525f8152606060208201528152602001906001900390816102ad5790505b5090505f5b60018401548110156103d6575f8460010182815481106102fb576102fb610d93565b5f9182526020808320909101546001600160a01b0316808352878252604092839020835180850185528281526001820180548651818702810187019097528087529396509194909384810193919291908301828280156103a757602002820191905f5260205f20905f905b82829054906101000a900460ff16600381111561038557610385610b92565b8152602060019283018181049485019490930390920291018084116103665790505b50505050508152508484815181106103c1576103c1610d93565b602090810291909101015250506001016102d9565b5092915050565b335f90815260208390526040902054829060029060ff166104115760405163c597feeb60e01b815260040160405180910390fd5b61041c8233836104ae565b610438576040516282b42960e81b815260040160405180910390fd5b61044284846107fd565b50505050565b335f90815260208490526040902054839060029060ff1661047c5760405163c597feeb60e01b815260040160405180910390fd5b6104878233836104ae565b6104a3576040516282b42960e81b815260040160405180910390fd5b6102008585856109a2565b6001600160a01b0382165f908152602084905260408120600101815b815481101561053e578360038111156104e5576104e5610b92565b8282815481106104f7576104f7610d93565b905f5260205f2090602091828204019190069054906101000a900460ff16600381111561052657610526610b92565b0361053657600192505050610544565b6001016104ca565b505f9150505b9392505050565b6001600160a01b0382165f9081526020849052604090205460ff166105835760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f9081526020849052604081209060018201905b8154811015610715578360038111156105bc576105bc610b92565b8282815481106105ce576105ce610d93565b905f5260205f2090602091828204019190069054906101000a900460ff1660038111156105fd576105fd610b92565b0361070d578154829061061290600190610da7565b8154811061062257610622610d93565b905f5260205f2090602091828204019190069054906101000a900460ff1682828154811061065257610652610d93565b905f5260205f2090602091828204019190066101000a81548160ff0219169083600381111561068357610683610b92565b02179055508180548061069857610698610dcc565b600190038181905f5260205f2090602091828204019190066101000a81549060ff021916905590558360038111156106d2576106d2610b92565b6040516001600160a01b038716907f34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769905f90a3505050505050565b6001016105a1565b505050505050565b6001600160a01b0382165f9081526020849052604090205460ff166107555760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208481526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a9182021916908360038111156107ad576107ad610b92565b02179055508160038111156107c4576107c4610b92565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a350505050565b6001600160a01b0381165f90815260208390526040812060018101549091036108395760405163178423dd60e01b815260040160405180910390fd5b6001600160a01b0382165f908152602084905260408120805460ff19168155906108666001830182610a91565b505f90505b600184015481101561096957826001600160a01b031684600101828154811061089657610896610d93565b5f918252602090912001546001600160a01b031603610961576001808501805490916108c191610da7565b815481106108d1576108d1610d93565b5f918252602090912001546001850180546001600160a01b0390921691839081106108fe576108fe610d93565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508360010180548061093c5761093c610dcc565b5f8281526020902081015f1990810180546001600160a01b0319169055019055610969565b60010161086b565b506040516001600160a01b038316907fc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8905f90a2505050565b6001600160a01b0382165f9081526020849052604090205460ff16156109db576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602085815260408220805460ff191660019081179091558681018054918201815583529082200180546001600160a01b0319169092179091555b8151811015610a5857610a508484848481518110610a4357610a43610d93565b602002602001015161071d565b600101610a23565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a2505050565b5080545f8255601f0160209004905f5260205f2090810190610ab39190610ab6565b50565b5b80821115610aca575f8155600101610ab7565b5090565b5f60208284031215610ade575f5ffd5b5035919050565b602080825282518282018190525f918401906040840190835b81811015610b255783516001600160a01b0316835260209384019390920191600101610afe565b509095945050505050565b80356001600160a01b0381168114610b46575f5ffd5b919050565b803560048110610b46575f5ffd5b5f5f5f60608486031215610b6b575f5ffd5b83359250610b7b60208501610b30565b9150610b8960408501610b4b565b90509250925092565b634e487b7160e01b5f52602160045260245ffd5b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015610c6357868503603f19018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101905f9060608801905b80831015610c4b57835160048110610c3357634e487b7160e01b5f52602160045260245ffd5b82526020938401936001939093019290910190610c0d565b50965050506020938401939190910190600101610bcc565b50929695505050505050565b5f5f60408385031215610c80575f5ffd5b82359150610c9060208401610b30565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610cbf575f5ffd5b83359250610ccf60208501610b30565b9150604084013567ffffffffffffffff811115610cea575f5ffd5b8401601f81018613610cfa575f5ffd5b803567ffffffffffffffff811115610d1457610d14610c99565b8060051b604051601f19603f830116810181811067ffffffffffffffff82111715610d4157610d41610c99565b604052918252602081840181019290810189841115610d5e575f5ffd5b6020850194505b83851015610d8457610d7685610b4b565b815260209485019401610d65565b50809450505050509250925092565b634e487b7160e01b5f52603260045260245ffd5b81810381811115610dc657634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220b1bf7b7a04ccc02f63761c1c632f8e11c0f4e6cb9657836495f719ab2f2ed81f64736f6c634300081c0033",
}

// RoleControlABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleControlMetaData.ABI instead.
var RoleControlABI = RoleControlMetaData.ABI

// RoleControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoleControlMetaData.Bin instead.
var RoleControlBin = RoleControlMetaData.Bin

// DeployRoleControl deploys a new Ethereum contract, binding an instance of RoleControl to it.
func DeployRoleControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoleControl, error) {
	parsed, err := RoleControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoleControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoleControl{RoleControlCaller: RoleControlCaller{contract: contract}, RoleControlTransactor: RoleControlTransactor{contract: contract}, RoleControlFilterer: RoleControlFilterer{contract: contract}}, nil
}

// RoleControl is an auto generated Go binding around an Ethereum contract.
type RoleControl struct {
	RoleControlCaller     // Read-only binding to the contract
	RoleControlTransactor // Write-only binding to the contract
	RoleControlFilterer   // Log filterer for contract events
}

// RoleControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleControlSession struct {
	Contract     *RoleControl      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleControlCallerSession struct {
	Contract *RoleControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RoleControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleControlTransactorSession struct {
	Contract     *RoleControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RoleControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleControlRaw struct {
	Contract *RoleControl // Generic contract binding to access the raw methods on
}

// RoleControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleControlCallerRaw struct {
	Contract *RoleControlCaller // Generic read-only contract binding to access the raw methods on
}

// RoleControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleControlTransactorRaw struct {
	Contract *RoleControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleControl creates a new instance of RoleControl, bound to a specific deployed contract.
func NewRoleControl(address common.Address, backend bind.ContractBackend) (*RoleControl, error) {
	contract, err := bindRoleControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleControl{RoleControlCaller: RoleControlCaller{contract: contract}, RoleControlTransactor: RoleControlTransactor{contract: contract}, RoleControlFilterer: RoleControlFilterer{contract: contract}}, nil
}

// NewRoleControlCaller creates a new read-only instance of RoleControl, bound to a specific deployed contract.
func NewRoleControlCaller(address common.Address, caller bind.ContractCaller) (*RoleControlCaller, error) {
	contract, err := bindRoleControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleControlCaller{contract: contract}, nil
}

// NewRoleControlTransactor creates a new write-only instance of RoleControl, bound to a specific deployed contract.
func NewRoleControlTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleControlTransactor, error) {
	contract, err := bindRoleControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleControlTransactor{contract: contract}, nil
}

// NewRoleControlFilterer creates a new log filterer instance of RoleControl, bound to a specific deployed contract.
func NewRoleControlFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleControlFilterer, error) {
	contract, err := bindRoleControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleControlFilterer{contract: contract}, nil
}

// bindRoleControl binds a generic wrapper to an already deployed contract.
func bindRoleControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoleControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleControl *RoleControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleControl.Contract.RoleControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleControl *RoleControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleControl.Contract.RoleControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleControl *RoleControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleControl.Contract.RoleControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleControl *RoleControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleControl *RoleControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleControl *RoleControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleControl.Contract.contract.Transact(opts, method, params...)
}

// RoleControlIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the RoleControl contract.
type RoleControlIdentityCreatedIterator struct {
	Event *RoleControlIdentityCreated // Event containing the contract specifics and raw log

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
func (it *RoleControlIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlIdentityCreated)
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
		it.Event = new(RoleControlIdentityCreated)
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
func (it *RoleControlIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlIdentityCreated represents a IdentityCreated event raised by the RoleControl contract.
type RoleControlIdentityCreated struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_RoleControl *RoleControlFilterer) FilterIdentityCreated(opts *bind.FilterOpts, wallet []common.Address) (*RoleControlIdentityCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControl.contract.FilterLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlIdentityCreatedIterator{contract: _RoleControl.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_RoleControl *RoleControlFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *RoleControlIdentityCreated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControl.contract.WatchLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlIdentityCreated)
				if err := _RoleControl.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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

// ParseIdentityCreated is a log parse operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_RoleControl *RoleControlFilterer) ParseIdentityCreated(log types.Log) (*RoleControlIdentityCreated, error) {
	event := new(RoleControlIdentityCreated)
	if err := _RoleControl.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlIdentityDeletedIterator is returned from FilterIdentityDeleted and is used to iterate over the raw logs and unpacked data for IdentityDeleted events raised by the RoleControl contract.
type RoleControlIdentityDeletedIterator struct {
	Event *RoleControlIdentityDeleted // Event containing the contract specifics and raw log

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
func (it *RoleControlIdentityDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlIdentityDeleted)
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
		it.Event = new(RoleControlIdentityDeleted)
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
func (it *RoleControlIdentityDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlIdentityDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlIdentityDeleted represents a IdentityDeleted event raised by the RoleControl contract.
type RoleControlIdentityDeleted struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityDeleted is a free log retrieval operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_RoleControl *RoleControlFilterer) FilterIdentityDeleted(opts *bind.FilterOpts, wallet []common.Address) (*RoleControlIdentityDeletedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControl.contract.FilterLogs(opts, "IdentityDeleted", walletRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlIdentityDeletedIterator{contract: _RoleControl.contract, event: "IdentityDeleted", logs: logs, sub: sub}, nil
}

// WatchIdentityDeleted is a free log subscription operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_RoleControl *RoleControlFilterer) WatchIdentityDeleted(opts *bind.WatchOpts, sink chan<- *RoleControlIdentityDeleted, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControl.contract.WatchLogs(opts, "IdentityDeleted", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlIdentityDeleted)
				if err := _RoleControl.contract.UnpackLog(event, "IdentityDeleted", log); err != nil {
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

// ParseIdentityDeleted is a log parse operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_RoleControl *RoleControlFilterer) ParseIdentityDeleted(log types.Log) (*RoleControlIdentityDeleted, error) {
	event := new(RoleControlIdentityDeleted)
	if err := _RoleControl.contract.UnpackLog(event, "IdentityDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RoleControl contract.
type RoleControlRoleGrantedIterator struct {
	Event *RoleControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *RoleControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlRoleGranted)
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
		it.Event = new(RoleControlRoleGranted)
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
func (it *RoleControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlRoleGranted represents a RoleGranted event raised by the RoleControl contract.
type RoleControlRoleGranted struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_RoleControl *RoleControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*RoleControlRoleGrantedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControl.contract.FilterLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlRoleGrantedIterator{contract: _RoleControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_RoleControl *RoleControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RoleControlRoleGranted, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControl.contract.WatchLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlRoleGranted)
				if err := _RoleControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_RoleControl *RoleControlFilterer) ParseRoleGranted(log types.Log) (*RoleControlRoleGranted, error) {
	event := new(RoleControlRoleGranted)
	if err := _RoleControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RoleControl contract.
type RoleControlRoleRevokedIterator struct {
	Event *RoleControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoleControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlRoleRevoked)
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
		it.Event = new(RoleControlRoleRevoked)
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
func (it *RoleControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlRoleRevoked represents a RoleRevoked event raised by the RoleControl contract.
type RoleControlRoleRevoked struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_RoleControl *RoleControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*RoleControlRoleRevokedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControl.contract.FilterLogs(opts, "RoleRevoked", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlRoleRevokedIterator{contract: _RoleControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_RoleControl *RoleControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoleControlRoleRevoked, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControl.contract.WatchLogs(opts, "RoleRevoked", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlRoleRevoked)
				if err := _RoleControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_RoleControl *RoleControlFilterer) ParseRoleRevoked(log types.Log) (*RoleControlRoleRevoked, error) {
	event := new(RoleControlRoleRevoked)
	if err := _RoleControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
