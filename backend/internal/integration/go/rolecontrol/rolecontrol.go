// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolecontrol

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

// EtherbaseCustomContract is an auto generated low-level Go binding around an user-defined struct.
type EtherbaseCustomContract struct {
	ContractAddress common.Address
	AddedBy         common.Address
	ContractABI     string
}

// EtherbaseSourceBatchEmitEvent is an auto generated low-level Go binding around an user-defined struct.
type EtherbaseSourceBatchEmitEvent struct {
	Name           string
	ArgumentTopics [][32]byte
	Data           []byte
}

// EtherbaseSourceInfo is an auto generated low-level Go binding around an user-defined struct.
type EtherbaseSourceInfo struct {
	SourceAddress common.Address
	Owner         common.Address
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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060fa8061001b5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80633fa4f245146034578063d09de08a14604d575b5f5ffd5b603b5f5481565b60405190815260200160405180910390f35b60536055565b005b60015f5f8282546064919060a0565b90915550505f546040519081527f93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c599060200160405180910390a1565b8082018082111560be57634e487b7160e01b5f52601160045260245ffd5b9291505056fea2646970667358221220c3aabacb675a9cbf35900fcab2ed7c2da4a90794b741075f625851a300985dcd64736f6c634300081c0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Counter *CounterCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Counter *CounterSession) Value() (*big.Int, error) {
	return _Counter.Contract.Value(&_Counter.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Counter *CounterCallerSession) Value() (*big.Int, error) {
	return _Counter.Contract.Value(&_Counter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// CounterValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the Counter contract.
type CounterValueChangedIterator struct {
	Event *CounterValueChanged // Event containing the contract specifics and raw log

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
func (it *CounterValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterValueChanged)
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
		it.Event = new(CounterValueChanged)
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
func (it *CounterValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterValueChanged represents a ValueChanged event raised by the Counter contract.
type CounterValueChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_Counter *CounterFilterer) FilterValueChanged(opts *bind.FilterOpts) (*CounterValueChangedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return &CounterValueChangedIterator{contract: _Counter.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_Counter *CounterFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *CounterValueChanged) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterValueChanged)
				if err := _Counter.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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
func (_Counter *CounterFilterer) ParseValueChanged(log types.Log) (*CounterValueChanged, error) {
	event := new(CounterValueChanged)
	if err := _Counter.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataValidatorMetaData contains all meta data concerning the DataValidator contract.
var DataValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"validateDataEncoding\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102f98061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063477b6e6a1461002d575b5f5ffd5b61004061003b36600461016b565b610042565b005b6001816005811115610056576100566101fd565b03610074578180602001905181019061006f9190610211565b505050565b6003816005811115610088576100886101fd565b036100a1578180602001905181019061006f9190610286565b60048160058111156100b5576100b56101fd565b036100ce578180602001905181019061006f9190610286565b60028160058111156100e2576100e26101fd565b036100fb578180602001905181019061006f919061029d565b5050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561013c5761013c6100ff565b604052919050565b5f67ffffffffffffffff82111561015d5761015d6100ff565b50601f01601f191660200190565b5f5f6040838503121561017c575f5ffd5b823567ffffffffffffffff811115610192575f5ffd5b8301601f810185136101a2575f5ffd5b80356101b56101b082610144565b610113565b8181528660208385010111156101c9575f5ffd5b816020840160208301375f602083830101528094505050506020830135600681106101f2575f5ffd5b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215610221575f5ffd5b815167ffffffffffffffff811115610237575f5ffd5b8201601f81018413610247575f5ffd5b80516102556101b082610144565b818152856020838501011115610269575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f60208284031215610296575f5ffd5b5051919050565b5f602082840312156102ad575f5ffd5b815180151581146102bc575f5ffd5b939250505056fea2646970667358221220afbca37d7453ca7ae48dffd284b8fe9c3a38bcfb6dc35cbd4d33c94e66cc7e7c64736f6c634300081c0033",
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

// EtherDatabaseImplMetaData contains all meta data concerning the EtherDatabaseImpl contract.
var EtherDatabaseImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"InvalidDataEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDataType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PathNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NewSegment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllSegments\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structEtherDatabaseLib.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getOrCreateSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"hasEntry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"removeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEtherDatabaseLib.BatchSetValue[]\",\"name\":\"values\",\"type\":\"tuple[]\"}],\"name\":\"setValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051611afa380380611afa833981016040819052602b91604e565b600180555f80546001600160a01b0319166001600160a01b038316179055506079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b611a74806100865f395ff3fe608060405234801561000f575f5ffd5b506004361061009b575f3560e01c806360778d951161006357806360778d951461013257806364ad6c881461015357806369da55d21461017657806371934ef0146101895780638998fbe81461019c575f5ffd5b80630761b8d01461009f57806317be85c3146100b45780633a5381b5146100d25780633f54d8bf146100fc5780634151df6114610111575b5f5ffd5b6100b26100ad366004611289565b6101af565b005b6100bc6101f7565b6040516100c9919061136c565b60405180910390f35b5f546100e4906001600160a01b031681565b6040516001600160a01b0390911681526020016100c9565b6101046104ca565b6040516100c99190611416565b61012461011f36600461146d565b61059e565b6040519081526020016100c9565b6101456101403660046114ac565b6105b2565b6040516100c99291906114df565b6101666101613660046114ac565b610663565b60405190151581526020016100c9565b6100b26101843660046114ac565b61069d565b61012461019736600461146d565b610785565b6100b26101aa3660046114ac565b6107af565b6101f085858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506108eb92505050565b5050505050565b60605f805b60055481101561024a576005818154811061021957610219611506565b5f91825260209091206003600490920201015460ff1615610242578161023e8161152e565b9250505b6001016101fc565b505f8167ffffffffffffffff81111561026557610265611546565b6040519080825280602002602001820160405280156102c057816020015b6102ad604080516080810190915260608152602081015f8152606060208201525f60409091015290565b8152602001906001900390816102835790505b5090505f805b6005548110156104c157600581815481106102e3576102e3611506565b5f91825260209091206003600490920201015460ff16156104b9576005818154811061031157610311611506565b905f5260205f2090600402016040518060800160405290815f820180546103379061155a565b80601f01602080910402602001604051908101604052809291908181526020018280546103639061155a565b80156103ae5780601f10610385576101008083540402835291602001916103ae565b820191905f5260205f20905b81548152906001019060200180831161039157829003601f168201915b5050509183525050600182015460209091019060ff1660058111156103d5576103d5611338565b60058111156103e6576103e6611338565b81526020016002820180546103fa9061155a565b80601f01602080910402602001604051908101604052809291908181526020018280546104269061155a565b80156104715780601f1061044857610100808354040283529160200191610471565b820191905f5260205f20905b81548152906001019060200180831161045457829003601f168201915b50505091835250506003919091015460ff161515602090910152835184908490811061049f5761049f611506565b602002602001018190525081806104b59061152e565b9250505b6001016102c6565b50909392505050565b60606003805480602002602001604051908101604052809291908181526020015f905b82821015610595578382905f5260205f2001805461050a9061155a565b80601f01602080910402602001604051908101604052809291908181526020018280546105369061155a565b80156105815780601f1061055857610100808354040283529160200191610581565b820191905f5260205f20905b81548152906001019060200180831161056457829003601f168201915b5050505050815260200190600101906104ed565b50505050905090565b5f6105a98383610c80565b90505b92915050565b5f60605f6105c08585610d97565b90506007816040516105d291906115a9565b9081526040519081900360200190205460ff1661060157505060408051602081019091525f808252915061065c565b5f61060d86865f610e5a565b9050600560068360405161062191906115a9565b9081526020016040518091039020548154811061064057610640611506565b5f91825260209091206001600490920201015460ff1693509150505b9250929050565b5f5f61066f8484610d97565b905060078160405161068191906115a9565b9081526040519081900360200190205460ff1691505092915050565b5f5b81811015610780576107788383838181106106bc576106bc611506565b90506020028101906106ce91906115b4565b6106d890806115d2565b8585858181106106ea576106ea611506565b90506020028101906106fc91906115b4565b61070d906040810190602001611618565b86868681811061071f5761071f611506565b905060200281019061073191906115b4565b61073f906040810190611631565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506108eb92505050565b60010161069f565b505050565b5f60028383604051610798929190611674565b908152602001604051809103902054905092915050565b5f6107ba8383610d97565b90506007816040516107cc91906115a9565b9081526040519081900360200190205460ff166107fc576040516357509b2760e11b815260040160405180910390fd5b5f60068260405161080d91906115a9565b90815260200160405180910390205490505f6005828154811061083257610832611506565b905f5260205f2090600402016003015f6101000a81548160ff02191690831515021790555060068260405161086791906115a9565b90815260200160405180910390205f905560078260405161088891906115a9565b9081526040805160209281900383018120805460ff1916905591820181525f80835290517fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a926108dd928892889291906116ab565b60405180910390a150505050565b5f6108f685856110d9565b90505f83600581111561090b5761090b611338565b03610a3f5760078160405161092091906115a9565b9081526040519081900360200190205460ff16156109dd575f60068260405161094991906115a9565b90815260200160405180910390205490505f6005828154811061096e5761096e611506565b905f5260205f2090600402016003015f6101000a81548160ff0219169083151502179055506006826040516109a391906115a9565b90815260200160405180910390205f90556007826040516109c491906115a9565b908152604051908190036020019020805460ff19169055505b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a858560405180602001604052805f8152505f6005811115610a2157610a21611338565b604051610a3194939291906116ab565b60405180910390a150610c7a565b600781604051610a4f91906115a9565b9081526040519081900360200190205460ff1615610aeb575f600682604051610a7891906115a9565b90815260200160405180910390205490505f60058281548110610a9d57610a9d611506565b905f5260205f209060040201905083816002019081610abc91906117b2565b5060018082018054879260ff1990911690836005811115610adf57610adf611338565b02179055505050610c2a565b5f6040518060800160405280838152602001856005811115610b0f57610b0f611338565b815260208101859052600160409091018190526005805491820181555f528151919250829160049091027f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001908190610b6890826117b2565b50602082015160018083018054909160ff1990911690836005811115610b9057610b90611338565b021790555060408201516002820190610ba990826117b2565b50606091909101516003909101805460ff1916911515919091179055600554610bd49060019061186d565b600683604051610be491906115a9565b9081526020016040518091039020819055506001600783604051610c0891906115a9565b908152604051908190036020019020805491151560ff19909216919091179055505b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a858584866005811115610c6057610c60611338565b604051610c7094939291906116ab565b60405180910390a1505b50505050565b5f5f60028484604051610c94929190611674565b9081526040805160209281900383019020545f818152600490935291205490915060ff1615610cc45790506105ac565b600180549081905f610cd58361152e565b91905055508060028686604051610ced929190611674565b90815260405190819003602001902055600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01610d39858783611880565b505f8181526004602052604090819020805460ff191660011790555181907f4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e77890610d86908890889061193a565b60405180910390a291506105ac9050565b604080515f808252602082019092526060915b83811015610e2e575f6002868684818110610dc757610dc7611506565b9050602002810190610dd99190611631565b604051610de7929190611674565b908152602001604051809103902054905082610e0282611159565b604051602001610e1392919061194d565b60408051601f19818403018152919052925050600101610daa565b50604051610e429082905f90602001611961565b60408051808303601f19018152919052949350505050565b60605f610e678585610d97565b9050600781604051610e7991906115a9565b9081526040519081900360200190205460ff16610ea9576040516357509b2760e11b815260040160405180910390fd5b5f6005600683604051610ebc91906115a9565b90815260200160405180910390205481548110610edb57610edb611506565b5f9182526020909120600490910201600381015490915060ff16610f12576040516357509b2760e11b815260040160405180910390fd5b5f846005811115610f2557610f25611338565b14158015610f5c5750836005811115610f4057610f40611338565b600182015460ff166005811115610f5957610f59611338565b14155b15610f9057600181015460405163254c060560e21b8152610f8791869160ff90911690600401611985565b60405180910390fd5b5f846005811115610fa357610fa3611338565b14158015610fc357506005846005811115610fc057610fc0611338565b14155b15611042575f546040516323bdb73560e11b81526001600160a01b039091169063477b6e6a90610ffc90600285019088906004016119a0565b5f6040518083038186803b158015611012575f5ffd5b505afa925050508015611023575060015b6110425783604051631e9c993560e21b8152600401610f879190611a30565b8060020180546110519061155a565b80601f016020809104026020016040519081016040528092919081815260200182805461107d9061155a565b80156110c85780601f1061109f576101008083540402835291602001916110c8565b820191905f5260205f20905b8154815290600101906020018083116110ab57829003601f168201915b5050505050925050505b9392505050565b604080515f808252602082019092526060915b83811015610e2e575f61112186868481811061110a5761110a611506565b905060200281019061111c9190611631565b610c80565b90508261112d82611159565b60405160200161113e92919061194d565b60408051601f198184030181529190529250506001016110ec565b6060815f03611184576040515f60208201526021016040516020818303038152906040529050919050565b6060825b608081106111c6578181607f1660801760f81b6040516020016111ac929190611961565b60408051601f19818403018152919052915060071c611188565b8181607f1660f81b6040516020016111df929190611961565b60408051601f19818403018152919052949350505050565b5f5f83601f840112611207575f5ffd5b50813567ffffffffffffffff81111561121e575f5ffd5b6020830191508360208260051b850101111561065c575f5ffd5b803560068110611246575f5ffd5b919050565b5f5f83601f84011261125b575f5ffd5b50813567ffffffffffffffff811115611272575f5ffd5b60208301915083602082850101111561065c575f5ffd5b5f5f5f5f5f6060868803121561129d575f5ffd5b853567ffffffffffffffff8111156112b3575f5ffd5b6112bf888289016111f7565b90965094506112d2905060208701611238565b9250604086013567ffffffffffffffff8111156112ed575f5ffd5b6112f98882890161124b565b969995985093965092949392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b5f52602160045260245ffd5b6006811061136857634e487b7160e01b5f52602160045260245ffd5b9052565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561140a57603f1987860301845281518051608087526113b8608088018261130a565b905060208201516113cc602089018261134c565b50604082015187820360408901526113e4828261130a565b606093840151151598909301979097525094506020938401939190910190600101611392565b50929695505050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561140a57603f1987860301845261145885835161130a565b9450602093840193919091019060010161143c565b5f5f6020838503121561147e575f5ffd5b823567ffffffffffffffff811115611494575f5ffd5b6114a08582860161124b565b90969095509350505050565b5f5f602083850312156114bd575f5ffd5b823567ffffffffffffffff8111156114d3575f5ffd5b6114a0858286016111f7565b6114e9818461134c565b604060208201525f6114fe604083018461130a565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f6001820161153f5761153f61151a565b5060010190565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061156e57607f821691505b60208210810361158c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f81518060208401855e5f93019283525090919050565b5f6105a98284611592565b5f8235605e198336030181126115c8575f5ffd5b9190910192915050565b5f5f8335601e198436030181126115e7575f5ffd5b83018035915067ffffffffffffffff821115611601575f5ffd5b6020019150600581901b360382131561065c575f5ffd5b5f60208284031215611628575f5ffd5b6105a982611238565b5f5f8335601e19843603018112611646575f5ffd5b83018035915067ffffffffffffffff821115611660575f5ffd5b60200191503681900382131561065c575f5ffd5b818382375f9101908152919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606080825281018490525f6080600586901b830181019083018783601e1936839003015b8982101561173d57868503607f1901845282358181126116ed575f5ffd5b8b0160208101903567ffffffffffffffff811115611709575f5ffd5b803603821315611717575f5ffd5b611722878284611683565b965050506020830192506020840193506001820191506116cf565b505050508281036020840152611753818661130a565b915050611765604083018460ff169052565b95945050505050565b601f82111561078057805f5260205f20601f840160051c810160208510156117935750805b601f840160051c820191505b818110156101f0575f815560010161179f565b815167ffffffffffffffff8111156117cc576117cc611546565b6117e0816117da845461155a565b8461176e565b6020601f821160018114611812575f83156117fb5750848201515b5f19600385901b1c1916600184901b1784556101f0565b5f84815260208120601f198516915b828110156118415787850151825560209485019460019092019101611821565b508482101561185e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b818103818111156105ac576105ac61151a565b67ffffffffffffffff83111561189857611898611546565b6118ac836118a6835461155a565b8361176e565b5f601f8411600181146118dd575f85156118c65750838201355b5f19600387901b1c1916600186901b1783556101f0565b5f83815260208120601f198716915b8281101561190c57868501358255602094850194600190920191016118ec565b5086821015611928575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b602081525f6114fe602083018486611683565b5f6114fe61195b8386611592565b84611592565b5f61196c8285611592565b6001600160f81b03199390931683525050600101919050565b60408101611993828561134c565b6110d2602083018461134c565b604081525f5f84546119b18161155a565b806040860152600182165f81146119cf57600181146119eb57611a1c565b60ff1983166060870152606082151560051b8701019350611a1c565b875f5260205f205f5b83811015611a13578154888201606001526001909101906020016119f4565b87016060019450505b505050809150506110d2602083018461134c565b602081016105ac828461134c56fea2646970667358221220d61ac621a17823e34816195be547a8df8db9d34bf757e1b79c393798516409bd64736f6c634300081c0033",
}

// EtherDatabaseImplABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherDatabaseImplMetaData.ABI instead.
var EtherDatabaseImplABI = EtherDatabaseImplMetaData.ABI

// EtherDatabaseImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherDatabaseImplMetaData.Bin instead.
var EtherDatabaseImplBin = EtherDatabaseImplMetaData.Bin

// DeployEtherDatabaseImpl deploys a new Ethereum contract, binding an instance of EtherDatabaseImpl to it.
func DeployEtherDatabaseImpl(auth *bind.TransactOpts, backend bind.ContractBackend, validator common.Address) (common.Address, *types.Transaction, *EtherDatabaseImpl, error) {
	parsed, err := EtherDatabaseImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherDatabaseImplBin), backend, validator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherDatabaseImpl{EtherDatabaseImplCaller: EtherDatabaseImplCaller{contract: contract}, EtherDatabaseImplTransactor: EtherDatabaseImplTransactor{contract: contract}, EtherDatabaseImplFilterer: EtherDatabaseImplFilterer{contract: contract}}, nil
}

// EtherDatabaseImpl is an auto generated Go binding around an Ethereum contract.
type EtherDatabaseImpl struct {
	EtherDatabaseImplCaller     // Read-only binding to the contract
	EtherDatabaseImplTransactor // Write-only binding to the contract
	EtherDatabaseImplFilterer   // Log filterer for contract events
}

// EtherDatabaseImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherDatabaseImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDatabaseImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherDatabaseImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDatabaseImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherDatabaseImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDatabaseImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherDatabaseImplSession struct {
	Contract     *EtherDatabaseImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EtherDatabaseImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherDatabaseImplCallerSession struct {
	Contract *EtherDatabaseImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EtherDatabaseImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherDatabaseImplTransactorSession struct {
	Contract     *EtherDatabaseImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EtherDatabaseImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherDatabaseImplRaw struct {
	Contract *EtherDatabaseImpl // Generic contract binding to access the raw methods on
}

// EtherDatabaseImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherDatabaseImplCallerRaw struct {
	Contract *EtherDatabaseImplCaller // Generic read-only contract binding to access the raw methods on
}

// EtherDatabaseImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherDatabaseImplTransactorRaw struct {
	Contract *EtherDatabaseImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherDatabaseImpl creates a new instance of EtherDatabaseImpl, bound to a specific deployed contract.
func NewEtherDatabaseImpl(address common.Address, backend bind.ContractBackend) (*EtherDatabaseImpl, error) {
	contract, err := bindEtherDatabaseImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseImpl{EtherDatabaseImplCaller: EtherDatabaseImplCaller{contract: contract}, EtherDatabaseImplTransactor: EtherDatabaseImplTransactor{contract: contract}, EtherDatabaseImplFilterer: EtherDatabaseImplFilterer{contract: contract}}, nil
}

// NewEtherDatabaseImplCaller creates a new read-only instance of EtherDatabaseImpl, bound to a specific deployed contract.
func NewEtherDatabaseImplCaller(address common.Address, caller bind.ContractCaller) (*EtherDatabaseImplCaller, error) {
	contract, err := bindEtherDatabaseImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseImplCaller{contract: contract}, nil
}

// NewEtherDatabaseImplTransactor creates a new write-only instance of EtherDatabaseImpl, bound to a specific deployed contract.
func NewEtherDatabaseImplTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherDatabaseImplTransactor, error) {
	contract, err := bindEtherDatabaseImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseImplTransactor{contract: contract}, nil
}

// NewEtherDatabaseImplFilterer creates a new log filterer instance of EtherDatabaseImpl, bound to a specific deployed contract.
func NewEtherDatabaseImplFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherDatabaseImplFilterer, error) {
	contract, err := bindEtherDatabaseImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseImplFilterer{contract: contract}, nil
}

// bindEtherDatabaseImpl binds a generic wrapper to an already deployed contract.
func bindEtherDatabaseImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherDatabaseImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDatabaseImpl *EtherDatabaseImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherDatabaseImpl.Contract.EtherDatabaseImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDatabaseImpl *EtherDatabaseImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.EtherDatabaseImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDatabaseImpl *EtherDatabaseImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.EtherDatabaseImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDatabaseImpl *EtherDatabaseImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherDatabaseImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDatabaseImpl *EtherDatabaseImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDatabaseImpl *EtherDatabaseImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.contract.Transact(opts, method, params...)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherDatabaseImpl *EtherDatabaseImplCaller) GetAllSegments(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EtherDatabaseImpl.contract.Call(opts, &out, "getAllSegments")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherDatabaseImpl *EtherDatabaseImplSession) GetAllSegments() ([]string, error) {
	return _EtherDatabaseImpl.Contract.GetAllSegments(&_EtherDatabaseImpl.CallOpts)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherDatabaseImpl *EtherDatabaseImplCallerSession) GetAllSegments() ([]string, error) {
	return _EtherDatabaseImpl.Contract.GetAllSegments(&_EtherDatabaseImpl.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherDatabaseImpl *EtherDatabaseImplCaller) GetEntries(opts *bind.CallOpts) ([]EtherDatabaseLibNode, error) {
	var out []interface{}
	err := _EtherDatabaseImpl.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]EtherDatabaseLibNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]EtherDatabaseLibNode)).(*[]EtherDatabaseLibNode)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherDatabaseImpl *EtherDatabaseImplSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EtherDatabaseImpl.Contract.GetEntries(&_EtherDatabaseImpl.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherDatabaseImpl *EtherDatabaseImplCallerSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EtherDatabaseImpl.Contract.GetEntries(&_EtherDatabaseImpl.CallOpts)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherDatabaseImpl *EtherDatabaseImplCaller) GetSegmentId(opts *bind.CallOpts, segment string) (*big.Int, error) {
	var out []interface{}
	err := _EtherDatabaseImpl.contract.Call(opts, &out, "getSegmentId", segment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherDatabaseImpl *EtherDatabaseImplSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EtherDatabaseImpl.Contract.GetSegmentId(&_EtherDatabaseImpl.CallOpts, segment)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherDatabaseImpl *EtherDatabaseImplCallerSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EtherDatabaseImpl.Contract.GetSegmentId(&_EtherDatabaseImpl.CallOpts, segment)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EtherDatabaseImpl *EtherDatabaseImplCaller) GetValue(opts *bind.CallOpts, segments []string) (uint8, []byte, error) {
	var out []interface{}
	err := _EtherDatabaseImpl.contract.Call(opts, &out, "getValue", segments)

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
func (_EtherDatabaseImpl *EtherDatabaseImplSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EtherDatabaseImpl.Contract.GetValue(&_EtherDatabaseImpl.CallOpts, segments)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EtherDatabaseImpl *EtherDatabaseImplCallerSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EtherDatabaseImpl.Contract.GetValue(&_EtherDatabaseImpl.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherDatabaseImpl *EtherDatabaseImplCaller) HasEntry(opts *bind.CallOpts, segments []string) (bool, error) {
	var out []interface{}
	err := _EtherDatabaseImpl.contract.Call(opts, &out, "hasEntry", segments)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherDatabaseImpl *EtherDatabaseImplSession) HasEntry(segments []string) (bool, error) {
	return _EtherDatabaseImpl.Contract.HasEntry(&_EtherDatabaseImpl.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherDatabaseImpl *EtherDatabaseImplCallerSession) HasEntry(segments []string) (bool, error) {
	return _EtherDatabaseImpl.Contract.HasEntry(&_EtherDatabaseImpl.CallOpts, segments)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherDatabaseImpl *EtherDatabaseImplCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherDatabaseImpl.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherDatabaseImpl *EtherDatabaseImplSession) Validator() (common.Address, error) {
	return _EtherDatabaseImpl.Contract.Validator(&_EtherDatabaseImpl.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherDatabaseImpl *EtherDatabaseImplCallerSession) Validator() (common.Address, error) {
	return _EtherDatabaseImpl.Contract.Validator(&_EtherDatabaseImpl.CallOpts)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherDatabaseImpl *EtherDatabaseImplTransactor) GetOrCreateSegmentId(opts *bind.TransactOpts, segment string) (*types.Transaction, error) {
	return _EtherDatabaseImpl.contract.Transact(opts, "getOrCreateSegmentId", segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherDatabaseImpl *EtherDatabaseImplSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.GetOrCreateSegmentId(&_EtherDatabaseImpl.TransactOpts, segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherDatabaseImpl *EtherDatabaseImplTransactorSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.GetOrCreateSegmentId(&_EtherDatabaseImpl.TransactOpts, segment)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplTransactor) RemoveValue(opts *bind.TransactOpts, segments []string) (*types.Transaction, error) {
	return _EtherDatabaseImpl.contract.Transact(opts, "removeValue", segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.RemoveValue(&_EtherDatabaseImpl.TransactOpts, segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplTransactorSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.RemoveValue(&_EtherDatabaseImpl.TransactOpts, segments)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplTransactor) SetValue(opts *bind.TransactOpts, segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherDatabaseImpl.contract.Transact(opts, "setValue", segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.SetValue(&_EtherDatabaseImpl.TransactOpts, segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplTransactorSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.SetValue(&_EtherDatabaseImpl.TransactOpts, segments, dataType, data)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplTransactor) SetValues(opts *bind.TransactOpts, values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherDatabaseImpl.contract.Transact(opts, "setValues", values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.SetValues(&_EtherDatabaseImpl.TransactOpts, values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherDatabaseImpl *EtherDatabaseImplTransactorSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherDatabaseImpl.Contract.SetValues(&_EtherDatabaseImpl.TransactOpts, values)
}

// EtherDatabaseImplEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the EtherDatabaseImpl contract.
type EtherDatabaseImplEthDBPathUpdateIterator struct {
	Event *EtherDatabaseImplEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *EtherDatabaseImplEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDatabaseImplEthDBPathUpdate)
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
		it.Event = new(EtherDatabaseImplEthDBPathUpdate)
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
func (it *EtherDatabaseImplEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDatabaseImplEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDatabaseImplEthDBPathUpdate represents a EthDBPathUpdate event raised by the EtherDatabaseImpl contract.
type EtherDatabaseImplEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherDatabaseImpl *EtherDatabaseImplFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*EtherDatabaseImplEthDBPathUpdateIterator, error) {

	logs, sub, err := _EtherDatabaseImpl.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseImplEthDBPathUpdateIterator{contract: _EtherDatabaseImpl.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherDatabaseImpl *EtherDatabaseImplFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *EtherDatabaseImplEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _EtherDatabaseImpl.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDatabaseImplEthDBPathUpdate)
				if err := _EtherDatabaseImpl.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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
func (_EtherDatabaseImpl *EtherDatabaseImplFilterer) ParseEthDBPathUpdate(log types.Log) (*EtherDatabaseImplEthDBPathUpdate, error) {
	event := new(EtherDatabaseImplEthDBPathUpdate)
	if err := _EtherDatabaseImpl.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherDatabaseImplNewSegmentIterator is returned from FilterNewSegment and is used to iterate over the raw logs and unpacked data for NewSegment events raised by the EtherDatabaseImpl contract.
type EtherDatabaseImplNewSegmentIterator struct {
	Event *EtherDatabaseImplNewSegment // Event containing the contract specifics and raw log

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
func (it *EtherDatabaseImplNewSegmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDatabaseImplNewSegment)
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
		it.Event = new(EtherDatabaseImplNewSegment)
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
func (it *EtherDatabaseImplNewSegmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDatabaseImplNewSegmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDatabaseImplNewSegment represents a NewSegment event raised by the EtherDatabaseImpl contract.
type EtherDatabaseImplNewSegment struct {
	Segment string
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSegment is a free log retrieval operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherDatabaseImpl *EtherDatabaseImplFilterer) FilterNewSegment(opts *bind.FilterOpts, id []*big.Int) (*EtherDatabaseImplNewSegmentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherDatabaseImpl.contract.FilterLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherDatabaseImplNewSegmentIterator{contract: _EtherDatabaseImpl.contract, event: "NewSegment", logs: logs, sub: sub}, nil
}

// WatchNewSegment is a free log subscription operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherDatabaseImpl *EtherDatabaseImplFilterer) WatchNewSegment(opts *bind.WatchOpts, sink chan<- *EtherDatabaseImplNewSegment, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherDatabaseImpl.contract.WatchLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDatabaseImplNewSegment)
				if err := _EtherDatabaseImpl.contract.UnpackLog(event, "NewSegment", log); err != nil {
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
func (_EtherDatabaseImpl *EtherDatabaseImplFilterer) ParseNewSegment(log types.Log) (*EtherDatabaseImplNewSegment, error) {
	event := new(EtherDatabaseImplNewSegment)
	if err := _EtherDatabaseImpl.contract.UnpackLog(event, "NewSegment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherDatabaseLibMetaData contains all meta data concerning the EtherDatabaseLib contract.
var EtherDatabaseLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"InvalidDataEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDataType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PathNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NewSegment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllSegments\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structEtherDatabaseLib.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getOrCreateSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"hasEntry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"removeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEtherDatabaseLib.BatchSetValue[]\",\"name\":\"values\",\"type\":\"tuple[]\"}],\"name\":\"setValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EtherDatabaseLibABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherDatabaseLibMetaData.ABI instead.
var EtherDatabaseLibABI = EtherDatabaseLibMetaData.ABI

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

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherDatabaseLib *EtherDatabaseLibCaller) GetAllSegments(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EtherDatabaseLib.contract.Call(opts, &out, "getAllSegments")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherDatabaseLib *EtherDatabaseLibSession) GetAllSegments() ([]string, error) {
	return _EtherDatabaseLib.Contract.GetAllSegments(&_EtherDatabaseLib.CallOpts)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherDatabaseLib *EtherDatabaseLibCallerSession) GetAllSegments() ([]string, error) {
	return _EtherDatabaseLib.Contract.GetAllSegments(&_EtherDatabaseLib.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherDatabaseLib *EtherDatabaseLibCaller) GetEntries(opts *bind.CallOpts) ([]EtherDatabaseLibNode, error) {
	var out []interface{}
	err := _EtherDatabaseLib.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]EtherDatabaseLibNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]EtherDatabaseLibNode)).(*[]EtherDatabaseLibNode)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherDatabaseLib *EtherDatabaseLibSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EtherDatabaseLib.Contract.GetEntries(&_EtherDatabaseLib.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherDatabaseLib *EtherDatabaseLibCallerSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EtherDatabaseLib.Contract.GetEntries(&_EtherDatabaseLib.CallOpts)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherDatabaseLib *EtherDatabaseLibCaller) GetSegmentId(opts *bind.CallOpts, segment string) (*big.Int, error) {
	var out []interface{}
	err := _EtherDatabaseLib.contract.Call(opts, &out, "getSegmentId", segment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherDatabaseLib *EtherDatabaseLibSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EtherDatabaseLib.Contract.GetSegmentId(&_EtherDatabaseLib.CallOpts, segment)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherDatabaseLib *EtherDatabaseLibCallerSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EtherDatabaseLib.Contract.GetSegmentId(&_EtherDatabaseLib.CallOpts, segment)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EtherDatabaseLib *EtherDatabaseLibCaller) GetValue(opts *bind.CallOpts, segments []string) (uint8, []byte, error) {
	var out []interface{}
	err := _EtherDatabaseLib.contract.Call(opts, &out, "getValue", segments)

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
func (_EtherDatabaseLib *EtherDatabaseLibSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EtherDatabaseLib.Contract.GetValue(&_EtherDatabaseLib.CallOpts, segments)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EtherDatabaseLib *EtherDatabaseLibCallerSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EtherDatabaseLib.Contract.GetValue(&_EtherDatabaseLib.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherDatabaseLib *EtherDatabaseLibCaller) HasEntry(opts *bind.CallOpts, segments []string) (bool, error) {
	var out []interface{}
	err := _EtherDatabaseLib.contract.Call(opts, &out, "hasEntry", segments)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherDatabaseLib *EtherDatabaseLibSession) HasEntry(segments []string) (bool, error) {
	return _EtherDatabaseLib.Contract.HasEntry(&_EtherDatabaseLib.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherDatabaseLib *EtherDatabaseLibCallerSession) HasEntry(segments []string) (bool, error) {
	return _EtherDatabaseLib.Contract.HasEntry(&_EtherDatabaseLib.CallOpts, segments)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherDatabaseLib *EtherDatabaseLibCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherDatabaseLib.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherDatabaseLib *EtherDatabaseLibSession) Validator() (common.Address, error) {
	return _EtherDatabaseLib.Contract.Validator(&_EtherDatabaseLib.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherDatabaseLib *EtherDatabaseLibCallerSession) Validator() (common.Address, error) {
	return _EtherDatabaseLib.Contract.Validator(&_EtherDatabaseLib.CallOpts)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherDatabaseLib *EtherDatabaseLibTransactor) GetOrCreateSegmentId(opts *bind.TransactOpts, segment string) (*types.Transaction, error) {
	return _EtherDatabaseLib.contract.Transact(opts, "getOrCreateSegmentId", segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherDatabaseLib *EtherDatabaseLibSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.GetOrCreateSegmentId(&_EtherDatabaseLib.TransactOpts, segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherDatabaseLib *EtherDatabaseLibTransactorSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.GetOrCreateSegmentId(&_EtherDatabaseLib.TransactOpts, segment)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherDatabaseLib *EtherDatabaseLibTransactor) RemoveValue(opts *bind.TransactOpts, segments []string) (*types.Transaction, error) {
	return _EtherDatabaseLib.contract.Transact(opts, "removeValue", segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherDatabaseLib *EtherDatabaseLibSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.RemoveValue(&_EtherDatabaseLib.TransactOpts, segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherDatabaseLib *EtherDatabaseLibTransactorSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.RemoveValue(&_EtherDatabaseLib.TransactOpts, segments)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherDatabaseLib *EtherDatabaseLibTransactor) SetValue(opts *bind.TransactOpts, segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherDatabaseLib.contract.Transact(opts, "setValue", segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherDatabaseLib *EtherDatabaseLibSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.SetValue(&_EtherDatabaseLib.TransactOpts, segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherDatabaseLib *EtherDatabaseLibTransactorSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.SetValue(&_EtherDatabaseLib.TransactOpts, segments, dataType, data)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherDatabaseLib *EtherDatabaseLibTransactor) SetValues(opts *bind.TransactOpts, values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherDatabaseLib.contract.Transact(opts, "setValues", values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherDatabaseLib *EtherDatabaseLibSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.SetValues(&_EtherDatabaseLib.TransactOpts, values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherDatabaseLib *EtherDatabaseLibTransactorSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherDatabaseLib.Contract.SetValues(&_EtherDatabaseLib.TransactOpts, values)
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

// EtherbaseMetaData contains all meta data concerning the Etherbase contract.
var EtherbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ContractAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyContractABI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"CustomContractAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"CustomContractDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"}],\"name\":\"CustomContractSchemaAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SourceCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractABI\",\"type\":\"string\"}],\"name\":\"addCustomContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customContractAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addedBy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractABI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"deleteCustomContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getCustomContract\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addedBy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractABI\",\"type\":\"string\"}],\"internalType\":\"structEtherbase.CustomContract\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCustomContracts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addedBy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractABI\",\"type\":\"string\"}],\"internalType\":\"structEtherbase.CustomContract[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalSchemas\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSources\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structEtherbase.SourceInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sources\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060405161585b38038061585b833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b6157d5806100865f395ff3fe608060405234801561000f575f5ffd5b50600436106100a5575f3560e01c8063a66ec8681161006e578063a66ec86814610166578063abd609741461017b578063af520da11461018e578063bef863a6146101a3578063c133bf19146101b8578063c44fb91b146101c0575f5ffd5b80621f75d9146100a9578063336d7167146100d45780633a5381b5146100f4578063736df5591461011e578063a3663d3714610133575b5f5ffd5b6100bc6100b7366004611176565b6101d3565b6040516100cb939291906111c6565b60405180910390f35b6100e76100e2366004611176565b61028a565b6040516100cb919061123a565b5f54610106906001600160a01b031681565b6040516001600160a01b0390911681526020016100cb565b61013161012c366004611176565b6103af565b005b61014661014136600461124c565b610594565b604080516001600160a01b039384168152929091166020830152016100cb565b61016e6105cb565b6040516100cb9190611263565b61010661018936600461124c565b610ab8565b610196610ae0565b6040516100cb91906113a1565b6101ab610c1e565b6040516100cb91906113fb565b610106610e2a565b6101316101ce366004611466565b610fea565b600260208190525f91825260409091208054600182015492820180546001600160a01b039283169490921692916102099061152c565b80601f01602080910402602001604051908101604052809291908181526020018280546102359061152c565b80156102805780601f1061025757610100808354040283529160200191610280565b820191905f5260205f20905b81548152906001019060200180831161026357829003601f168201915b5050505050905083565b60408051606080820183525f8083526020830152918101919091526001600160a01b038281165f90815260026020526040902054166102dc57604051633b56498960e21b815260040160405180910390fd5b6001600160a01b038083165f908152600260208181526040928390208351606081018552815486168152600182015490951691850191909152908101805491928401916103289061152c565b80601f01602080910402602001604051908101604052809291908181526020018280546103549061152c565b801561039f5780601f106103765761010080835404028352916020019161039f565b820191905f5260205f20905b81548152906001019060200180831161038257829003601f168201915b5050505050815250509050919050565b6001600160a01b038181165f90815260026020526040902054166103e657604051633b56498960e21b815260040160405180910390fd5b6001600160a01b038181165f908152600260205260409020600101541633146104225760405163ea8e4eb560e01b815260040160405180910390fd5b6001600160a01b0381165f908152600260208190526040822080546001600160a01b031990811682556001820180549091169055919061046490830182611103565b505f90505b60035481101561055d57816001600160a01b0316600382815481106104905761049061155e565b5f918252602090912001546001600160a01b03160361055557600380546104b990600190611572565b815481106104c9576104c961155e565b5f91825260209091200154600380546001600160a01b0390921691839081106104f4576104f461155e565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550600380548061053057610530611597565b5f8281526020902081015f1990810180546001600160a01b031916905501905561055d565b600101610469565b506040516001600160a01b038216907fce2cd4a1ec306fda51c35ba3ed91f20df734a8097698e844508959656ccb5118905f90a250565b600181815481106105a3575f80fd5b5f918252602090912060029091020180546001909101546001600160a01b0391821692501682565b60408051600580825260c082019092526060915f9190816020015b6106196040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b8152602001906001900390816105e6575050604080516003808252608082019092529192505f9190602082015b604080516060808201835280825260208201525f918101919091528152602001906001900390816106465750506040805160a0810182526004606082019081526366726f6d60e01b608083015281528151808301835260078152666164647265737360c81b60208281019190915282015260019181019190915281519192509082905f906106d6576106d661155e565b6020026020010181905250604051806060016040528060405180604001604052806002815260200161746f60f01b8152508152602001604051806040016040528060078152602001666164647265737360c81b8152508152602001600115158152508160018151811061074b5761074b61155e565b602002602001018190525060405180606001604052806040518060400160405280600581526020016476616c756560d81b8152508152602001604051806040016040528060078152602001663ab4b73a191a9b60c91b81525081526020015f1515815250816002815181106107c2576107c261155e565b60200260200101819052506040518060a00160405280604051806040016040528060088152602001672a3930b739b332b960c11b815250815260200160405180608001604052806044815260200161575c6044913981526020018281526020017fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8152602001600260ff16815250825f815181106108625761086261155e565b6020908102919091010152604080516003808252608082019092525f91816020015b604080516060808201835280825260208201525f918101919091528152602001906001900390816108845750506040805160a0810182526004606082019081526366726f6d60e01b608083015281528151808301835260078152666164647265737360c81b60208281019190915282015260019181019190915281519192509082905f906109145761091461155e565b6020026020010181905250604051806060016040528060405180604001604052806002815260200161746f60f01b8152508152602001604051806040016040528060078152602001666164647265737360c81b815250815260200160011515815250816001815181106109895761098961155e565b60200260200101819052506040518060600160405280604051806040016040528060078152602001661d1bdad95b925960ca1b8152508152602001604051806040016040528060078152602001663ab4b73a191a9b60c91b81525081526020016001151581525081600281518110610a0357610a0361155e565b60200260200101819052506040518060a00160405280604051806040016040528060088152602001672a3930b739b332b960c11b81525081526020016040518060800160405280604481526020016157186044913981526020018281526020017fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8152602001600360ff1681525083600281518110610aa457610aa461155e565b602002602001018190525082935050505090565b60038181548110610ac7575f80fd5b5f918252602090912001546001600160a01b0316905081565b6001546060905f9067ffffffffffffffff811115610b0057610b00611452565b604051908082528060200260200182016040528015610b4457816020015b604080518082019091525f8082526020820152815260200190600190039081610b1e5790505b5090505f5b600154811015610c185760018181548110610b6657610b6661155e565b5f91825260209091206002909102015482516001600160a01b0390911690839083908110610b9657610b9661155e565b60209081029190910101516001600160a01b0390911690526001805482908110610bc257610bc261155e565b905f5260205f2090600202016001015f9054906101000a90046001600160a01b0316828281518110610bf657610bf661155e565b6020908102919091018101516001600160a01b03909216910152600101610b49565b50919050565b6003546060905f9067ffffffffffffffff811115610c3e57610c3e611452565b604051908082528060200260200182016040528015610c8a57816020015b60408051606080820183525f808352602083015291810191909152815260200190600190039081610c5c5790505b5090505f5b600354811015610c18575f60038281548110610cad57610cad61155e565b5f9182526020808320909101546001600160a01b03908116808452600280845260408086208151606081018352815486168152600182015490951695850195909552908401805492965092939290840191610d079061152c565b80601f0160208091040260200160405190810160405280929190818152602001828054610d339061152c565b8015610d7e5780601f10610d5557610100808354040283529160200191610d7e565b820191905f5260205f20905b815481529060010190602001808311610d6157829003601f168201915b505050505081525050905081848481518110610d9c57610d9c61155e565b60200260200101515f01906001600160a01b031690816001600160a01b0316815250508060200151848481518110610dd657610dd661155e565b6020026020010151602001906001600160a01b031690816001600160a01b0316815250508060400151848481518110610e1157610e1161155e565b6020908102919091010151604001525050600101610c8f565b5f8054604051829133916001600160a01b0390911690610e499061113d565b6001600160a01b03928316815291166020820152604001604051809103905ff080158015610e79573d5f5f3e3d5ffd5b50905060016040518060400160405280836001600160a01b03168152602001836001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ed4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef891906115ab565b6001600160a01b039081169091528254600181810185555f948552602094859020845160029093020180549284166001600160a01b03199384161781559385015193018054938316939091169290921790915560408051638da5cb5b60e01b8152905191841692638da5cb5b926004808401938290030181865afa158015610f82573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fa691906115ab565b6001600160a01b0316816001600160a01b03167f04ff85b5d2c64e18db6e3d3f6814a460cd9736cf5eed34eb285b1d0a9233370d60405160405180910390a3919050565b6001600160a01b038281165f9081526002602052604090205416156110225760405163e8dc2ba560e01b815260040160405180910390fd5b80515f036110425760405162e32e1d60e21b815260040160405180910390fd5b6001600160a01b0382165f81815260026020819052604090912080546001600160a01b03199081169093178155600181018054909316331790925581016110898382611612565b50600380546001810182555f9182527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b03861690811790915560405190917f996a8af2e36bc142de1bccafc39ce86c4cd7092058777ed39fd8e795f15e7c9a91a2505050565b50805461110f9061152c565b5f825580601f1061111e575050565b601f0160209004905f5260205f209081019061113a919061114a565b50565b61404a806116ce83390190565b5b8082111561115e575f815560010161114b565b5090565b6001600160a01b038116811461113a575f5ffd5b5f60208284031215611186575f5ffd5b813561119181611162565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b038481168252831660208201526060604082018190525f906111f190830184611198565b95945050505050565b60018060a01b03815116825260018060a01b0360208201511660208301525f6040820151606060408501526112326060850182611198565b949350505050565b602081525f61119160208301846111fa565b5f6020828403121561125c575f5ffd5b5035919050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561139557603f198786030184528151805160a087526112af60a0880182611198565b9050602082015187820360208901526112c88282611198565b9150506040820151878203604089015281815180845260208401915060208160051b8501016020840193505f5b8281101561135a57601f19868303018452845180516060845261131b6060850182611198565b9050602082015184820360208601526113348282611198565b6040938401511515959093019490945250602095860195949094019391506001016112f5565b50606086015160608c01526080860151955061137b60808c018760ff169052565b995050506020968701969490940193505050600101611289565b50929695505050505050565b602080825282518282018190525f918401906040840190835b818110156113f057835180516001600160a01b0390811685526020918201511681850152909301926040909201916001016113ba565b509095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561139557603f1987860301845261143d8583516111fa565b94506020938401939190910190600101611421565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215611477575f5ffd5b823561148281611162565b9150602083013567ffffffffffffffff81111561149d575f5ffd5b8301601f810185136114ad575f5ffd5b803567ffffffffffffffff8111156114c7576114c7611452565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156114f6576114f6611452565b60405281815282820160200187101561150d575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b600181811c9082168061154057607f821691505b602082108103610c1857634e487b7160e01b5f52602260045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b8181038181111561159157634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603160045260245ffd5b5f602082840312156115bb575f5ffd5b815161119181611162565b601f82111561160d57805f5260205f20601f840160051c810160208510156115eb5750805b601f840160051c820191505b8181101561160a575f81556001016115f7565b50505b505050565b815167ffffffffffffffff81111561162c5761162c611452565b6116408161163a845461152c565b846115c6565b6020601f821160018114611672575f831561165b5750848201515b5f19600385901b1c1916600184901b17845561160a565b5f84815260208120601f198516915b828110156116a15787850151825560209485019460019092019101611681565b50848210156116be57868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fe608060405234801561000f575f5ffd5b5060405161404a38038061404a83398101604081905261002e916103c3565b6100378261005e565b6001600455600380546001600160a01b0319166001600160a01b038316179055505061041c565b600280546001600160a01b0319166001600160a01b03831617905560408051600480825260a082019092525f91602082016080803683370190505090506002815f815181106100af576100af6103f4565b602002602001019060038111156100c8576100c8610408565b908160038111156100db576100db610408565b815250505f816001815181106100f3576100f36103f4565b6020026020010190600381111561010c5761010c610408565b9081600381111561011f5761011f610408565b81525050600181600281518110610138576101386103f4565b6020026020010190600381111561015157610151610408565b9081600381111561016457610164610408565b8152505060038160038151811061017d5761017d6103f4565b6020026020010190600381111561019657610196610408565b908160038111156101a9576101a9610408565b9052506101b682826101ba565b5050565b6001600160a01b0382165f9081526020819052604090205460ff16156101f3576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602081905260408120805460ff19166001908117909155805480820182559082527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169092179091555b81518110156102915761028983838381518110610276576102766103f4565b60200260200101516102c960201b60201c565b600101610257565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a25050565b6001600160a01b0382165f9081526020819052604090205460ff166103015760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208181526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a91820219169083600381111561035957610359610408565b021790555081600381111561037057610370610408565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a3505050565b80516001600160a01b03811681146103be575f5ffd5b919050565b5f5f604083850312156103d4575f5ffd5b6103dd836103a8565b91506103eb602084016103a8565b90509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b613c21806104295f395ff3fe608060405234801561000f575f5ffd5b506004361061016d575f3560e01c806364ad6c88116100d95780638da5cb5b11610093578063c76853b81161006e578063c76853b814610370578063d50d542e14610390578063dced4783146103a3578063f30826f0146103b6575f5ffd5b80638da5cb5b14610337578063a4dbe8a01461034a578063a8d29d1d1461035d575f5ffd5b806364ad6c88146102c357806369da55d2146102d657806371934ef0146102e95780637eeede35146102fc5780638998fbe8146103115780638c71045614610324575f5ffd5b80633a5381b51161012a5780633a5381b51461020e5780633e840236146102395780633f54d8bf1461024c5780634151df611461026157806349ab123d1461028257806360778d95146102a2575f5ffd5b80630277e023146101715780630761b8d014610199578063077d3c03146101ae57806317be85c3146101c15780631bfa8601146101d65780632f2545d2146101eb575b5f5ffd5b61018461017f366004612df0565b6103be565b60405190151581526020015b60405180910390f35b6101ac6101a7366004612eac565b61042d565b005b6101ac6101bc366004612f4f565b610475565b6101c96104dc565b6040516101909190612fd6565b6101de6107ae565b6040516101909190613080565b6101fe6101f93660046130cb565b61080e565b6040516101909493929190613117565b600354610221906001600160a01b031681565b6040516001600160a01b039091168152602001610190565b6101ac610247366004612f4f565b610951565b6102546109b3565b6040516101909190613156565b61027461026f3660046131ad565b610a87565b604051908152602001610190565b6102956102903660046131ad565b610a9b565b60405161019091906131eb565b6102b56102b03660046132d8565b610e8e565b60405161019092919061330a565b6101846102d13660046132d8565b610f3f565b6101ac6102e43660046132d8565b610f79565b6102746102f73660046131ad565b61105c565b610304611086565b6040516101909190613329565b6101ac61031f3660046132d8565b6111ee565b6101846103323660046132d8565b61132a565b600254610221906001600160a01b031681565b6101ac6103583660046133da565b611427565b6101ac61036b36600461347f565b6116cc565b61038361037e3660046130cb565b611731565b6040516101909190613498565b61029561039e3660046131ad565b6117d4565b6101ac6103b13660046134aa565b611b4a565b610254611bac565b335f90815260208190526040812054819060ff166103ef5760405163c597feeb60e01b815260040160405180910390fd5b6103f93382611c77565b610415576040516282b42960e81b815260040160405180910390fd5b6104228787878787611d11565b979650505050505050565b61046e85858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611e4392505050565b5050505050565b335f9081526020819052604090205460029060ff166104a75760405163c597feeb60e01b815260040160405180910390fd5b6104b13382611c77565b6104cd576040516282b42960e81b815260040160405180910390fd5b6104d783836121d8565b505050565b60605f805b60085481101561052f57600881815481106104fe576104fe61356a565b5f91825260209091206003600490920201015460ff1615610527578161052381613592565b9250505b6001016104e1565b505f816001600160401b0381111561054957610549612d54565b6040519080825280602002602001820160405280156105a457816020015b610591604080516080810190915260608152602081015f8152606060208201525f60409091015290565b8152602001906001900390816105675790505b5090505f805b6008548110156107a557600881815481106105c7576105c761356a565b5f91825260209091206003600490920201015460ff161561079d57600881815481106105f5576105f561356a565b905f5260205f2090600402016040518060800160405290815f8201805461061b906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610647906135aa565b80156106925780601f1061066957610100808354040283529160200191610692565b820191905f5260205f20905b81548152906001019060200180831161067557829003601f168201915b5050509183525050600182015460209091019060ff1660058111156106b9576106b9612fae565b60058111156106ca576106ca612fae565b81526020016002820180546106de906135aa565b80601f016020809104026020016040519081016040528092919081815260200182805461070a906135aa565b80156107555780601f1061072c57610100808354040283529160200191610755565b820191905f5260205f20905b81548152906001019060200180831161073857829003601f168201915b50505091835250506003919091015460ff16151560209091015283518490849081106107835761078361356a565b6020026020010181905250818061079990613592565b9250505b6001016105aa565b50909392505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561080457602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116107e6575b5050505050905090565b8051602081830181018051600b82529282019190930120915280548190610834906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610860906135aa565b80156108ab5780601f10610882576101008083540402835291602001916108ab565b820191905f5260205f20905b81548152906001019060200180831161088e57829003601f168201915b5050505050908060010180546108c0906135aa565b80601f01602080910402602001604051908101604052809291908181526020018280546108ec906135aa565b80156109375780601f1061090e57610100808354040283529160200191610937565b820191905f5260205f20905b81548152906001019060200180831161091a57829003601f168201915b50505050600383015460049093015491929160ff16905084565b335f9081526020819052604090205460029060ff166109835760405163c597feeb60e01b815260040160405180910390fd5b61098d3382611c77565b6109a9576040516282b42960e81b815260040160405180910390fd5b6104d783836123a1565b60606006805480602002602001604051908101604052809291908181526020015f905b82821015610a7e578382905f5260205f200180546109f3906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1f906135aa565b8015610a6a5780601f10610a4157610100808354040283529160200191610a6a565b820191905f5260205f20905b815481529060010190602001808311610a4d57829003601f168201915b5050505050815260200190600101906109d6565b50505050905090565b5f610a928383612480565b90505b92915050565b610ace6040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b5f600c8484604051610ae19291906135dc565b90815260200160405180910390208054610afa906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610b26906135aa565b8015610b715780601f10610b4857610100808354040283529160200191610b71565b820191905f5260205f20905b815481529060010190602001808311610b5457829003601f168201915b5050505050905080515f03610b9957604051634fde869d60e01b815260040160405180910390fd5b600b81604051610ba99190613602565b90815260200160405180910390206040518060a00160405290815f82018054610bd1906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfd906135aa565b8015610c485780601f10610c1f57610100808354040283529160200191610c48565b820191905f5260205f20905b815481529060010190602001808311610c2b57829003601f168201915b50505050508152602001600182018054610c61906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8d906135aa565b8015610cd85780601f10610caf57610100808354040283529160200191610cd8565b820191905f5260205f20905b815481529060010190602001808311610cbb57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015610e66578382905f5260205f2090600302016040518060600160405290815f82018054610d35906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610d61906135aa565b8015610dac5780601f10610d8357610100808354040283529160200191610dac565b820191905f5260205f20905b815481529060010190602001808311610d8f57829003601f168201915b50505050508152602001600182018054610dc5906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610df1906135aa565b8015610e3c5780601f10610e1357610100808354040283529160200191610e3c565b820191905f5260205f20905b815481529060010190602001808311610e1f57829003601f168201915b50505091835250506002919091015460ff1615156020918201529082526001929092019101610d05565b505050908252506003820154602082015260049091015460ff16604090910152949350505050565b5f60605f610e9c8585612597565b9050600a81604051610eae9190613602565b9081526040519081900360200190205460ff16610edd57505060408051602081019091525f8082529150610f38565b5f610ee986865f61265a565b90506008600983604051610efd9190613602565b90815260200160405180910390205481548110610f1c57610f1c61356a565b5f91825260209091206001600490920201015460ff1693509150505b9250929050565b5f5f610f4b8484612597565b9050600a81604051610f5d9190613602565b9081526040519081900360200190205460ff1691505092915050565b5f5b818110156104d757611054838383818110610f9857610f9861356a565b9050602002810190610faa919061360d565b610fb4908061362b565b858585818110610fc657610fc661356a565b9050602002810190610fd8919061360d565b610fe9906040810190602001613670565b868686818110610ffb57610ffb61356a565b905060200281019061100d919061360d565b61101b906040810190613689565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611e4392505050565b600101610f7b565b5f6005838360405161106f9291906135dc565b908152602001604051809103902054905092915050565b6001546060905f906001600160401b038111156110a5576110a5612d54565b6040519080825280602002602001820160405280156110ea57816020015b604080518082019091525f8152606060208201528152602001906001900390816110c35790505b5090505f5b6001548110156111e8575f6001828154811061110d5761110d61356a565b5f9182526020808320909101546001600160a01b0316808352828252604092839020835180850185528281526001820180548651818702810187019097528087529396509194909384810193919291908301828280156111b957602002820191905f5260205f20905f905b82829054906101000a900460ff16600381111561119757611197612fae565b8152602060019283018181049485019490930390920291018084116111785790505b50505050508152508484815181106111d3576111d361356a565b602090810291909101015250506001016110ef565b50919050565b5f6111f98383612597565b9050600a8160405161120b9190613602565b9081526040519081900360200190205460ff1661123b576040516357509b2760e11b815260040160405180910390fd5b5f60098260405161124c9190613602565b90815260200160405180910390205490505f600882815481106112715761127161356a565b905f5260205f2090600402016003015f6101000a81548160ff0219169083151502179055506009826040516112a69190613602565b90815260200160405180910390205f9055600a826040516112c79190613602565b9081526040805160209281900383018120805460ff1916905591820181525f80835290517fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a9261131c928892889291906136f3565b60405180910390a150505050565b335f90815260208190526040812054819060ff1661135b5760405163c597feeb60e01b815260040160405180910390fd5b6113653382611c77565b611381576040516282b42960e81b815260040160405180910390fd5b5f5b8381101561141c573685858381811061139e5761139e61356a565b90506020028101906113b0919061360d565b90506114126113bf8280613689565b6113cc602085018561362b565b6113d96040870187613689565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611d1192505050565b5050600101611383565b506001949350505050565b335f9081526020819052604090205460019060ff166114595760405163c597feeb60e01b815260040160405180910390fd5b6114633382611c77565b61147f576040516282b42960e81b815260040160405180910390fd5b600b88886040516114919291906135dc565b90815260405190819003602001902080546114ab906135aa565b1590506114cb57604051630268c2d760e51b815260040160405180910390fd5b5f805b83811015611526578484828181106114e8576114e861356a565b90506020028101906114fa919061360d565b61150b9060608101906040016137b9565b1561151e578161151a816137d4565b9250505b6001016114ce565b5060038160ff16111561154b57604051626f3b7d60e31b815260040160405180910390fd5b5f600b8a8a60405161155e9291906135dc565b90815260405190819003602001902090508061157b8a8c8361384a565b506001810161158b888a8361384a565b50600381018690555f5b848110156115f057816002018686838181106115b3576115b361356a565b90506020028101906115c5919061360d565b81546001810183555f92835260209092209091600302016115e6828261390a565b5050600101611595565b5060048101805460ff191660ff8416179055600d80546001810182555f919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb50161163e8a8c8361384a565b508989600c8a8a6040516116539291906135dc565b9081526020016040518091039020918261166e92919061384a565b50898960405161167f9291906135dc565b60405180910390207f4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de89896040516116b8929190613a0b565b60405180910390a250505050505050505050565b335f9081526020819052604090205460029060ff166116fe5760405163c597feeb60e01b815260040160405180910390fd5b6117083382611c77565b611724576040516282b42960e81b815260040160405180910390fd5b61172d826128da565b5050565b8051602081830181018051600c8252928201919093012091528054611755906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611781906135aa565b80156117cc5780601f106117a3576101008083540402835291602001916117cc565b820191905f5260205f20905b8154815290600101906020018083116117af57829003601f168201915b505050505081565b6118076040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b600b83836040516118199291906135dc565b9081526040519081900360200190208054611833906135aa565b90505f0361185457604051634fde869d60e01b815260040160405180910390fd5b600b83836040516118669291906135dc565b90815260200160405180910390206040518060a00160405290815f8201805461188e906135aa565b80601f01602080910402602001604051908101604052809291908181526020018280546118ba906135aa565b80156119055780601f106118dc57610100808354040283529160200191611905565b820191905f5260205f20905b8154815290600101906020018083116118e857829003601f168201915b5050505050815260200160018201805461191e906135aa565b80601f016020809104026020016040519081016040528092919081815260200182805461194a906135aa565b80156119955780601f1061196c57610100808354040283529160200191611995565b820191905f5260205f20905b81548152906001019060200180831161197857829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015611b23578382905f5260205f2090600302016040518060600160405290815f820180546119f2906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611a1e906135aa565b8015611a695780601f10611a4057610100808354040283529160200191611a69565b820191905f5260205f20905b815481529060010190602001808311611a4c57829003601f168201915b50505050508152602001600182018054611a82906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611aae906135aa565b8015611af95780601f10611ad057610100808354040283529160200191611af9565b820191905f5260205f20905b815481529060010190602001808311611adc57829003601f168201915b50505091835250506002919091015460ff16151560209182015290825260019290920191016119c2565b505050908252506003820154602082015260049091015460ff166040909101529392505050565b335f9081526020819052604090205460029060ff16611b7c5760405163c597feeb60e01b815260040160405180910390fd5b611b863382611c77565b611ba2576040516282b42960e81b815260040160405180910390fd5b6104d78383612a73565b6060600d805480602002602001604051908101604052809291908181526020015f905b82821015610a7e578382905f5260205f20018054611bec906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611c18906135aa565b8015611c635780601f10611c3a57610100808354040283529160200191611c63565b820191905f5260205f20905b815481529060010190602001808311611c4657829003601f168201915b505050505081526020019060010190611bcf565b6001600160a01b0382165f908152602081905260408120600101815b8154811015611d0757836003811115611cae57611cae612fae565b828281548110611cc057611cc061356a565b905f5260205f2090602091828204019190069054906101000a900460ff166003811115611cef57611cef612fae565b03611cff57600192505050610a95565b600101611c93565b505f949350505050565b5f600b8686604051611d249291906135dc565b9081526040519081900360200190208054611d3e906135aa565b90505f03611d5f57604051634fde869d60e01b815260040160405180910390fd5b5f600b8787604051611d729291906135dc565b908152604051908190036020019020600380820154919250851115611daa5760405163321fc7cf60e11b815260040160405180910390fd5b600482015460ff168514611dd157604051631f73a41360e01b815260040160405180910390fd5b835160208501868015611dfb5760018114611e045760028114611e0f5760038114611e1f57611e30565b838383a1611e30565b8835848484a2611e30565b60208901358935858585a3611e30565b604089013560208a01358a35868686a45b5050600193505050505b95945050505050565b5f611e4e8585612b7c565b90505f836005811115611e6357611e63612fae565b03611f9757600a81604051611e789190613602565b9081526040519081900360200190205460ff1615611f35575f600982604051611ea19190613602565b90815260200160405180910390205490505f60088281548110611ec657611ec661356a565b905f5260205f2090600402016003015f6101000a81548160ff021916908315150217905550600982604051611efb9190613602565b90815260200160405180910390205f9055600a82604051611f1c9190613602565b908152604051908190036020019020805460ff19169055505b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a858560405180602001604052805f8152505f6005811115611f7957611f79612fae565b604051611f8994939291906136f3565b60405180910390a1506121d2565b600a81604051611fa79190613602565b9081526040519081900360200190205460ff1615612043575f600982604051611fd09190613602565b90815260200160405180910390205490505f60088281548110611ff557611ff561356a565b905f5260205f2090600402019050838160020190816120149190613a1e565b5060018082018054879260ff199091169083600581111561203757612037612fae565b02179055505050612182565b5f604051806080016040528083815260200185600581111561206757612067612fae565b815260208101859052600160409091018190526008805491820181555f528151919250829160049091027ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3019081906120c09082613a1e565b50602082015160018083018054909160ff19909116908360058111156120e8576120e8612fae565b0217905550604082015160028201906121019082613a1e565b50606091909101516003909101805460ff191691151591909117905560085461212c90600190613ad3565b60098360405161213c9190613602565b9081526020016040518091039020819055506001600a836040516121609190613602565b908152604051908190036020019020805491151560ff19909216919091179055505b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a8585848660058111156121b8576121b8612fae565b6040516121c894939291906136f3565b60405180910390a1505b50505050565b6001600160a01b0382165f9081526020819052604090205460ff166122105760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f9081526020819052604081209060018201905b815481101561046e5783600381111561224957612249612fae565b82828154811061225b5761225b61356a565b905f5260205f2090602091828204019190069054906101000a900460ff16600381111561228a5761228a612fae565b03612399578154829061229f90600190613ad3565b815481106122af576122af61356a565b905f5260205f2090602091828204019190069054906101000a900460ff168282815481106122df576122df61356a565b905f5260205f2090602091828204019190066101000a81548160ff0219169083600381111561231057612310612fae565b02179055508180548061232557612325613ae6565b600190038181905f5260205f2090602091828204019190066101000a81549060ff0219169055905583600381111561235f5761235f612fae565b6040516001600160a01b038716907f34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769905f90a35050505050565b60010161222e565b6001600160a01b0382165f9081526020819052604090205460ff166123d95760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208181526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a91820219169083600381111561243157612431612fae565b021790555081600381111561244857612448612fae565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a3505050565b5f5f600584846040516124949291906135dc565b9081526040805160209281900383019020545f818152600790935291205490915060ff16156124c4579050610a95565b600480549081905f6124d583613592565b919050555080600586866040516124ed9291906135dc565b90815260405190819003602001902055600680546001810182555f919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0161253985878361384a565b505f8181526007602052604090819020805460ff191660011790555181907f4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778906125869088908890613a0b565b60405180910390a29150610a959050565b604080515f808252602082019092526060915b8381101561262e575f60058686848181106125c7576125c761356a565b90506020028101906125d99190613689565b6040516125e79291906135dc565b90815260200160405180910390205490508261260282612bfc565b604051602001612613929190613afa565b60408051601f198184030181529190529250506001016125aa565b506040516126429082905f90602001613b0e565b60408051808303601f19018152919052949350505050565b60605f6126678585612597565b9050600a816040516126799190613602565b9081526040519081900360200190205460ff166126a9576040516357509b2760e11b815260040160405180910390fd5b5f60086009836040516126bc9190613602565b908152602001604051809103902054815481106126db576126db61356a565b5f9182526020909120600490910201600381015490915060ff16612712576040516357509b2760e11b815260040160405180910390fd5b5f84600581111561272557612725612fae565b1415801561275c575083600581111561274057612740612fae565b600182015460ff16600581111561275957612759612fae565b14155b1561279057600181015460405163254c060560e21b815261278791869160ff90911690600401613b32565b60405180910390fd5b5f8460058111156127a3576127a3612fae565b141580156127c3575060058460058111156127c0576127c0612fae565b14155b15612843576003546040516323bdb73560e11b81526001600160a01b039091169063477b6e6a906127fd9060028501908890600401613b4d565b5f6040518083038186803b158015612813575f5ffd5b505afa925050508015612824575060015b6128435783604051631e9c993560e21b81526004016127879190613bdd565b806002018054612852906135aa565b80601f016020809104026020016040519081016040528092919081815260200182805461287e906135aa565b80156128c95780601f106128a0576101008083540402835291602001916128c9565b820191905f5260205f20905b8154815290600101906020018083116128ac57829003601f168201915b5050505050925050505b9392505050565b6001600160a01b0381165f90815260208190526040812060018101549091036129165760405163178423dd60e01b815260040160405180910390fd5b6001600160a01b0382165f908152602081905260408120805460ff19168155906129436001830182612c9a565b505f90505b600154811015612a3b57826001600160a01b03166001828154811061296f5761296f61356a565b5f918252602090912001546001600160a01b031603612a335760018054612997908290613ad3565b815481106129a7576129a761356a565b5f91825260209091200154600180546001600160a01b0390921691839081106129d2576129d261356a565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506001805480612a0e57612a0e613ae6565b5f8281526020902081015f1990810180546001600160a01b0319169055019055612a3b565b600101612948565b506040516001600160a01b038316907fc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8905f90a25050565b6001600160a01b0382165f9081526020819052604090205460ff1615612aac576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602081905260408120805460ff19166001908117909155805480820182559082527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169092179091555b8151811015612b4457612b3c83838381518110612b2f57612b2f61356a565b60200260200101516123a1565b600101612b10565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a25050565b604080515f808252602082019092526060915b8381101561262e575f612bc4868684818110612bad57612bad61356a565b9050602002810190612bbf9190613689565b612480565b905082612bd082612bfc565b604051602001612be1929190613afa565b60408051601f19818403018152919052925050600101612b8f565b6060815f03612c27576040515f60208201526021016040516020818303038152906040529050919050565b6060825b60808110612c69578181607f1660801760f81b604051602001612c4f929190613b0e565b60408051601f19818403018152919052915060071c612c2b565b8181607f1660f81b604051602001612c82929190613b0e565b60408051601f19818403018152919052949350505050565b5080545f8255601f0160209004905f5260205f2090810190612cbc9190612cbf565b50565b5b80821115612cd3575f8155600101612cc0565b5090565b5f5f83601f840112612ce7575f5ffd5b5081356001600160401b03811115612cfd575f5ffd5b602083019150836020828501011115610f38575f5ffd5b5f5f83601f840112612d24575f5ffd5b5081356001600160401b03811115612d3a575f5ffd5b6020830191508360208260051b8501011115610f38575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715612d9057612d90612d54565b604052919050565b5f5f6001600160401b03841115612db157612db1612d54565b50601f8301601f1916602001612dc681612d68565b915050828152838383011115612dda575f5ffd5b828260208301375f602084830101529392505050565b5f5f5f5f5f60608688031215612e04575f5ffd5b85356001600160401b03811115612e19575f5ffd5b612e2588828901612cd7565b90965094505060208601356001600160401b03811115612e43575f5ffd5b612e4f88828901612d14565b90945092505060408601356001600160401b03811115612e6d575f5ffd5b8601601f81018813612e7d575f5ffd5b612e8c88823560208401612d98565b9150509295509295909350565b803560068110612ea7575f5ffd5b919050565b5f5f5f5f5f60608688031215612ec0575f5ffd5b85356001600160401b03811115612ed5575f5ffd5b612ee188828901612d14565b9096509450612ef4905060208701612e99565b925060408601356001600160401b03811115612f0e575f5ffd5b612f1a88828901612cd7565b969995985093965092949392505050565b80356001600160a01b0381168114612ea7575f5ffd5b803560048110612ea7575f5ffd5b5f5f60408385031215612f60575f5ffd5b612f6983612f2b565b9150612f7760208401612f41565b90509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b5f52602160045260245ffd5b60068110612fd257612fd2612fae565b9052565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561307457603f1987860301845281518051608087526130226080880182612f80565b905060208201516130366020890182612fc2565b506040820151878203604089015261304e8282612f80565b606093840151151598909301979097525094506020938401939190910190600101612ffc565b50929695505050505050565b602080825282518282018190525f918401906040840190835b818110156130c05783516001600160a01b0316835260209384019390920191600101613099565b509095945050505050565b5f602082840312156130db575f5ffd5b81356001600160401b038111156130f0575f5ffd5b8201601f81018413613100575f5ffd5b61310f84823560208401612d98565b949350505050565b608081525f6131296080830187612f80565b828103602084015261313b8187612f80565b91505083604083015260ff8316606083015295945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561307457603f19878603018452613198858351612f80565b9450602093840193919091019060010161317c565b5f5f602083850312156131be575f5ffd5b82356001600160401b038111156131d3575f5ffd5b6131df85828601612cd7565b90969095509350505050565b602081525f825160a0602084015261320660c0840182612f80565b90506020840151601f198483030160408501526132238282612f80565b6040860151858203601f190160608701528051808352919350602090810192508084019190600582901b8501015f5b828110156132b757601f1986830301845284518051606084526132786060850182612f80565b9050602082015184820360208601526132918282612f80565b604093840151151595909301949094525060209586019594909401939150600101613252565b50606088015160808801526080880151945061042260a088018660ff169052565b5f5f602083850312156132e9575f5ffd5b82356001600160401b038111156132fe575f5ffd5b6131df85828601612d14565b6133148184612fc2565b604060208201525f61310f6040830184612f80565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561307457868503603f19018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101905f9060608801905b808310156133c2578351600481106133aa576133aa612fae565b82526020938401936001939093019290910190613390565b5096505050602093840193919091019060010161334f565b5f5f5f5f5f5f5f6080888a0312156133f0575f5ffd5b87356001600160401b03811115613405575f5ffd5b6134118a828b01612cd7565b90985096505060208801356001600160401b0381111561342f575f5ffd5b61343b8a828b01612cd7565b9096509450506040880135925060608801356001600160401b03811115613460575f5ffd5b61346c8a828b01612d14565b989b979a50959850939692959293505050565b5f6020828403121561348f575f5ffd5b610a9282612f2b565b602081525f610a926020830184612f80565b5f5f604083850312156134bb575f5ffd5b6134c483612f2b565b915060208301356001600160401b038111156134de575f5ffd5b8301601f810185136134ee575f5ffd5b80356001600160401b0381111561350757613507612d54565b8060051b61351760208201612d68565b91825260208184018101929081019088841115613532575f5ffd5b6020850194505b8385101561355b5761354a85612f41565b825260209485019490910190613539565b80955050505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f600182016135a3576135a361357e565b5060010190565b600181811c908216806135be57607f821691505b6020821081036111e857634e487b7160e01b5f52602260045260245ffd5b818382375f9101908152919050565b5f81518060208401855e5f93019283525090919050565b5f610a9282846135eb565b5f8235605e19833603018112613621575f5ffd5b9190910192915050565b5f5f8335601e19843603018112613640575f5ffd5b8301803591506001600160401b03821115613659575f5ffd5b6020019150600581901b3603821315610f38575f5ffd5b5f60208284031215613680575f5ffd5b610a9282612e99565b5f5f8335601e1984360301811261369e575f5ffd5b8301803591506001600160401b038211156136b7575f5ffd5b602001915036819003821315610f38575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606080825281018490525f6080600586901b830181019083018783601e1936839003015b8982101561378457868503607f190184528235818112613735575f5ffd5b8b016020810190356001600160401b03811115613750575f5ffd5b80360382131561375e575f5ffd5b6137698782846136cb565b96505050602083019250602084019350600182019150613717565b50505050828103602084015261379a8186612f80565b915050611e3a604083018460ff169052565b8015158114612cbc575f5ffd5b5f602082840312156137c9575f5ffd5b81356128d3816137ac565b5f60ff821660ff81036137e9576137e961357e565b60010192915050565b601f8211156104d757805f5260205f20601f840160051c810160208510156138175750805b601f840160051c820191505b8181101561046e575f8155600101613823565b5f19600383901b1c191660019190911b1790565b6001600160401b0383111561386157613861612d54565b6138758361386f83546135aa565b836137f2565b5f601f8411600181146138a1575f851561388f5750838201355b6138998682613836565b84555061046e565b5f83815260208120601f198716915b828110156138d057868501358255602094850194600190920191016138b0565b50868210156138ec575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f8135610a95816137ac565b6139148283613689565b6001600160401b0381111561392b5761392b612d54565b61393f8161393985546135aa565b856137f2565b5f601f82116001811461396b575f83156139595750838201355b6139638482613836565b8655506139c2565b5f85815260208120601f198516915b8281101561399a578685013582556020948501946001909201910161397a565b50848210156139b6575f1960f88660031b161c19848701351681555b505060018360011b0185555b505050506139d36020830183613689565b6139e181836001860161384a565b505061172d6139f2604084016138fe565b6002830160ff1981541660ff8315151681178255505050565b602081525f61310f6020830184866136cb565b81516001600160401b03811115613a3757613a37612d54565b613a4b81613a4584546135aa565b846137f2565b6020601f821160018114613a78575f8315613a665750848201515b613a708482613836565b85555061046e565b5f84815260208120601f198516915b82811015613aa75787850151825560209485019460019092019101613a87565b5084821015613ac457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b81810381811115610a9557610a9561357e565b634e487b7160e01b5f52603160045260245ffd5b5f61310f613b0883866135eb565b846135eb565b5f613b1982856135eb565b6001600160f81b03199390931683525050600101919050565b60408101613b408285612fc2565b6128d36020830184612fc2565b604081525f5f8454613b5e816135aa565b806040860152600182165f8114613b7c5760018114613b9857613bc9565b60ff1983166060870152606082151560051b8701019350613bc9565b875f5260205f205f5b83811015613bc057815488820160600152600190910190602001613ba1565b87016060019450505b505050809150506128d36020830184612fc2565b60208101610a958284612fc256fea26469706673582212202b5b229f00090a107f7e35504651973effbb65a775bad46745053d0513b056e564736f6c634300081c00333078646466323532616431626532633839623639633262303638666333373864616139353262613766313633633461313136323866353561346466353233623365663a333078646466323532616431626532633839623639633262303638666333373864616139353262613766313633633461313136323866353561346466353233623365663a32a2646970667358221220518e9910fe996b3eb694496a54c94f7025c6c6b0431e5eeb998df0029748ab4564736f6c634300081c0033",
}

// EtherbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherbaseMetaData.ABI instead.
var EtherbaseABI = EtherbaseMetaData.ABI

// EtherbaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherbaseMetaData.Bin instead.
var EtherbaseBin = EtherbaseMetaData.Bin

// DeployEtherbase deploys a new Ethereum contract, binding an instance of Etherbase to it.
func DeployEtherbase(auth *bind.TransactOpts, backend bind.ContractBackend, _validator common.Address) (common.Address, *types.Transaction, *Etherbase, error) {
	parsed, err := EtherbaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherbaseBin), backend, _validator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Etherbase{EtherbaseCaller: EtherbaseCaller{contract: contract}, EtherbaseTransactor: EtherbaseTransactor{contract: contract}, EtherbaseFilterer: EtherbaseFilterer{contract: contract}}, nil
}

// Etherbase is an auto generated Go binding around an Ethereum contract.
type Etherbase struct {
	EtherbaseCaller     // Read-only binding to the contract
	EtherbaseTransactor // Write-only binding to the contract
	EtherbaseFilterer   // Log filterer for contract events
}

// EtherbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherbaseSession struct {
	Contract     *Etherbase        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherbaseCallerSession struct {
	Contract *EtherbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EtherbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherbaseTransactorSession struct {
	Contract     *EtherbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EtherbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherbaseRaw struct {
	Contract *Etherbase // Generic contract binding to access the raw methods on
}

// EtherbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherbaseCallerRaw struct {
	Contract *EtherbaseCaller // Generic read-only contract binding to access the raw methods on
}

// EtherbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherbaseTransactorRaw struct {
	Contract *EtherbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherbase creates a new instance of Etherbase, bound to a specific deployed contract.
func NewEtherbase(address common.Address, backend bind.ContractBackend) (*Etherbase, error) {
	contract, err := bindEtherbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Etherbase{EtherbaseCaller: EtherbaseCaller{contract: contract}, EtherbaseTransactor: EtherbaseTransactor{contract: contract}, EtherbaseFilterer: EtherbaseFilterer{contract: contract}}, nil
}

// NewEtherbaseCaller creates a new read-only instance of Etherbase, bound to a specific deployed contract.
func NewEtherbaseCaller(address common.Address, caller bind.ContractCaller) (*EtherbaseCaller, error) {
	contract, err := bindEtherbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherbaseCaller{contract: contract}, nil
}

// NewEtherbaseTransactor creates a new write-only instance of Etherbase, bound to a specific deployed contract.
func NewEtherbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherbaseTransactor, error) {
	contract, err := bindEtherbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherbaseTransactor{contract: contract}, nil
}

// NewEtherbaseFilterer creates a new log filterer instance of Etherbase, bound to a specific deployed contract.
func NewEtherbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherbaseFilterer, error) {
	contract, err := bindEtherbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherbaseFilterer{contract: contract}, nil
}

// bindEtherbase binds a generic wrapper to an already deployed contract.
func bindEtherbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etherbase *EtherbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etherbase.Contract.EtherbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etherbase *EtherbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherbase.Contract.EtherbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etherbase *EtherbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etherbase.Contract.EtherbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etherbase *EtherbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etherbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etherbase *EtherbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etherbase *EtherbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etherbase.Contract.contract.Transact(opts, method, params...)
}

// CustomContractAddresses is a free data retrieval call binding the contract method 0xabd60974.
//
// Solidity: function customContractAddresses(uint256 ) view returns(address)
func (_Etherbase *EtherbaseCaller) CustomContractAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "customContractAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomContractAddresses is a free data retrieval call binding the contract method 0xabd60974.
//
// Solidity: function customContractAddresses(uint256 ) view returns(address)
func (_Etherbase *EtherbaseSession) CustomContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _Etherbase.Contract.CustomContractAddresses(&_Etherbase.CallOpts, arg0)
}

// CustomContractAddresses is a free data retrieval call binding the contract method 0xabd60974.
//
// Solidity: function customContractAddresses(uint256 ) view returns(address)
func (_Etherbase *EtherbaseCallerSession) CustomContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _Etherbase.Contract.CustomContractAddresses(&_Etherbase.CallOpts, arg0)
}

// CustomContracts is a free data retrieval call binding the contract method 0x001f75d9.
//
// Solidity: function customContracts(address ) view returns(address contractAddress, address addedBy, string contractABI)
func (_Etherbase *EtherbaseCaller) CustomContracts(opts *bind.CallOpts, arg0 common.Address) (struct {
	ContractAddress common.Address
	AddedBy         common.Address
	ContractABI     string
}, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "customContracts", arg0)

	outstruct := new(struct {
		ContractAddress common.Address
		AddedBy         common.Address
		ContractABI     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AddedBy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ContractABI = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// CustomContracts is a free data retrieval call binding the contract method 0x001f75d9.
//
// Solidity: function customContracts(address ) view returns(address contractAddress, address addedBy, string contractABI)
func (_Etherbase *EtherbaseSession) CustomContracts(arg0 common.Address) (struct {
	ContractAddress common.Address
	AddedBy         common.Address
	ContractABI     string
}, error) {
	return _Etherbase.Contract.CustomContracts(&_Etherbase.CallOpts, arg0)
}

// CustomContracts is a free data retrieval call binding the contract method 0x001f75d9.
//
// Solidity: function customContracts(address ) view returns(address contractAddress, address addedBy, string contractABI)
func (_Etherbase *EtherbaseCallerSession) CustomContracts(arg0 common.Address) (struct {
	ContractAddress common.Address
	AddedBy         common.Address
	ContractABI     string
}, error) {
	return _Etherbase.Contract.CustomContracts(&_Etherbase.CallOpts, arg0)
}

// GetCustomContract is a free data retrieval call binding the contract method 0x336d7167.
//
// Solidity: function getCustomContract(address contractAddress) view returns((address,address,string))
func (_Etherbase *EtherbaseCaller) GetCustomContract(opts *bind.CallOpts, contractAddress common.Address) (EtherbaseCustomContract, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "getCustomContract", contractAddress)

	if err != nil {
		return *new(EtherbaseCustomContract), err
	}

	out0 := *abi.ConvertType(out[0], new(EtherbaseCustomContract)).(*EtherbaseCustomContract)

	return out0, err

}

// GetCustomContract is a free data retrieval call binding the contract method 0x336d7167.
//
// Solidity: function getCustomContract(address contractAddress) view returns((address,address,string))
func (_Etherbase *EtherbaseSession) GetCustomContract(contractAddress common.Address) (EtherbaseCustomContract, error) {
	return _Etherbase.Contract.GetCustomContract(&_Etherbase.CallOpts, contractAddress)
}

// GetCustomContract is a free data retrieval call binding the contract method 0x336d7167.
//
// Solidity: function getCustomContract(address contractAddress) view returns((address,address,string))
func (_Etherbase *EtherbaseCallerSession) GetCustomContract(contractAddress common.Address) (EtherbaseCustomContract, error) {
	return _Etherbase.Contract.GetCustomContract(&_Etherbase.CallOpts, contractAddress)
}

// GetCustomContracts is a free data retrieval call binding the contract method 0xbef863a6.
//
// Solidity: function getCustomContracts() view returns((address,address,string)[])
func (_Etherbase *EtherbaseCaller) GetCustomContracts(opts *bind.CallOpts) ([]EtherbaseCustomContract, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "getCustomContracts")

	if err != nil {
		return *new([]EtherbaseCustomContract), err
	}

	out0 := *abi.ConvertType(out[0], new([]EtherbaseCustomContract)).(*[]EtherbaseCustomContract)

	return out0, err

}

// GetCustomContracts is a free data retrieval call binding the contract method 0xbef863a6.
//
// Solidity: function getCustomContracts() view returns((address,address,string)[])
func (_Etherbase *EtherbaseSession) GetCustomContracts() ([]EtherbaseCustomContract, error) {
	return _Etherbase.Contract.GetCustomContracts(&_Etherbase.CallOpts)
}

// GetCustomContracts is a free data retrieval call binding the contract method 0xbef863a6.
//
// Solidity: function getCustomContracts() view returns((address,address,string)[])
func (_Etherbase *EtherbaseCallerSession) GetCustomContracts() ([]EtherbaseCustomContract, error) {
	return _Etherbase.Contract.GetCustomContracts(&_Etherbase.CallOpts)
}

// GetGlobalSchemas is a free data retrieval call binding the contract method 0xa66ec868.
//
// Solidity: function getGlobalSchemas() pure returns((string,string,(string,string,bool)[],bytes32,uint8)[])
func (_Etherbase *EtherbaseCaller) GetGlobalSchemas(opts *bind.CallOpts) ([]EventSchema, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "getGlobalSchemas")

	if err != nil {
		return *new([]EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventSchema)).(*[]EventSchema)

	return out0, err

}

// GetGlobalSchemas is a free data retrieval call binding the contract method 0xa66ec868.
//
// Solidity: function getGlobalSchemas() pure returns((string,string,(string,string,bool)[],bytes32,uint8)[])
func (_Etherbase *EtherbaseSession) GetGlobalSchemas() ([]EventSchema, error) {
	return _Etherbase.Contract.GetGlobalSchemas(&_Etherbase.CallOpts)
}

// GetGlobalSchemas is a free data retrieval call binding the contract method 0xa66ec868.
//
// Solidity: function getGlobalSchemas() pure returns((string,string,(string,string,bool)[],bytes32,uint8)[])
func (_Etherbase *EtherbaseCallerSession) GetGlobalSchemas() ([]EventSchema, error) {
	return _Etherbase.Contract.GetGlobalSchemas(&_Etherbase.CallOpts)
}

// GetSources is a free data retrieval call binding the contract method 0xaf520da1.
//
// Solidity: function getSources() view returns((address,address)[])
func (_Etherbase *EtherbaseCaller) GetSources(opts *bind.CallOpts) ([]EtherbaseSourceInfo, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "getSources")

	if err != nil {
		return *new([]EtherbaseSourceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]EtherbaseSourceInfo)).(*[]EtherbaseSourceInfo)

	return out0, err

}

// GetSources is a free data retrieval call binding the contract method 0xaf520da1.
//
// Solidity: function getSources() view returns((address,address)[])
func (_Etherbase *EtherbaseSession) GetSources() ([]EtherbaseSourceInfo, error) {
	return _Etherbase.Contract.GetSources(&_Etherbase.CallOpts)
}

// GetSources is a free data retrieval call binding the contract method 0xaf520da1.
//
// Solidity: function getSources() view returns((address,address)[])
func (_Etherbase *EtherbaseCallerSession) GetSources() ([]EtherbaseSourceInfo, error) {
	return _Etherbase.Contract.GetSources(&_Etherbase.CallOpts)
}

// Sources is a free data retrieval call binding the contract method 0xa3663d37.
//
// Solidity: function sources(uint256 ) view returns(address sourceAddress, address owner)
func (_Etherbase *EtherbaseCaller) Sources(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SourceAddress common.Address
	Owner         common.Address
}, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "sources", arg0)

	outstruct := new(struct {
		SourceAddress common.Address
		Owner         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Sources is a free data retrieval call binding the contract method 0xa3663d37.
//
// Solidity: function sources(uint256 ) view returns(address sourceAddress, address owner)
func (_Etherbase *EtherbaseSession) Sources(arg0 *big.Int) (struct {
	SourceAddress common.Address
	Owner         common.Address
}, error) {
	return _Etherbase.Contract.Sources(&_Etherbase.CallOpts, arg0)
}

// Sources is a free data retrieval call binding the contract method 0xa3663d37.
//
// Solidity: function sources(uint256 ) view returns(address sourceAddress, address owner)
func (_Etherbase *EtherbaseCallerSession) Sources(arg0 *big.Int) (struct {
	SourceAddress common.Address
	Owner         common.Address
}, error) {
	return _Etherbase.Contract.Sources(&_Etherbase.CallOpts, arg0)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Etherbase *EtherbaseCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etherbase.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Etherbase *EtherbaseSession) Validator() (common.Address, error) {
	return _Etherbase.Contract.Validator(&_Etherbase.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Etherbase *EtherbaseCallerSession) Validator() (common.Address, error) {
	return _Etherbase.Contract.Validator(&_Etherbase.CallOpts)
}

// AddCustomContract is a paid mutator transaction binding the contract method 0xc44fb91b.
//
// Solidity: function addCustomContract(address contractAddress, string contractABI) returns()
func (_Etherbase *EtherbaseTransactor) AddCustomContract(opts *bind.TransactOpts, contractAddress common.Address, contractABI string) (*types.Transaction, error) {
	return _Etherbase.contract.Transact(opts, "addCustomContract", contractAddress, contractABI)
}

// AddCustomContract is a paid mutator transaction binding the contract method 0xc44fb91b.
//
// Solidity: function addCustomContract(address contractAddress, string contractABI) returns()
func (_Etherbase *EtherbaseSession) AddCustomContract(contractAddress common.Address, contractABI string) (*types.Transaction, error) {
	return _Etherbase.Contract.AddCustomContract(&_Etherbase.TransactOpts, contractAddress, contractABI)
}

// AddCustomContract is a paid mutator transaction binding the contract method 0xc44fb91b.
//
// Solidity: function addCustomContract(address contractAddress, string contractABI) returns()
func (_Etherbase *EtherbaseTransactorSession) AddCustomContract(contractAddress common.Address, contractABI string) (*types.Transaction, error) {
	return _Etherbase.Contract.AddCustomContract(&_Etherbase.TransactOpts, contractAddress, contractABI)
}

// CreateSource is a paid mutator transaction binding the contract method 0xc133bf19.
//
// Solidity: function createSource() returns(address)
func (_Etherbase *EtherbaseTransactor) CreateSource(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherbase.contract.Transact(opts, "createSource")
}

// CreateSource is a paid mutator transaction binding the contract method 0xc133bf19.
//
// Solidity: function createSource() returns(address)
func (_Etherbase *EtherbaseSession) CreateSource() (*types.Transaction, error) {
	return _Etherbase.Contract.CreateSource(&_Etherbase.TransactOpts)
}

// CreateSource is a paid mutator transaction binding the contract method 0xc133bf19.
//
// Solidity: function createSource() returns(address)
func (_Etherbase *EtherbaseTransactorSession) CreateSource() (*types.Transaction, error) {
	return _Etherbase.Contract.CreateSource(&_Etherbase.TransactOpts)
}

// DeleteCustomContract is a paid mutator transaction binding the contract method 0x736df559.
//
// Solidity: function deleteCustomContract(address contractAddress) returns()
func (_Etherbase *EtherbaseTransactor) DeleteCustomContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Etherbase.contract.Transact(opts, "deleteCustomContract", contractAddress)
}

// DeleteCustomContract is a paid mutator transaction binding the contract method 0x736df559.
//
// Solidity: function deleteCustomContract(address contractAddress) returns()
func (_Etherbase *EtherbaseSession) DeleteCustomContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Etherbase.Contract.DeleteCustomContract(&_Etherbase.TransactOpts, contractAddress)
}

// DeleteCustomContract is a paid mutator transaction binding the contract method 0x736df559.
//
// Solidity: function deleteCustomContract(address contractAddress) returns()
func (_Etherbase *EtherbaseTransactorSession) DeleteCustomContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Etherbase.Contract.DeleteCustomContract(&_Etherbase.TransactOpts, contractAddress)
}

// EtherbaseCustomContractAddedIterator is returned from FilterCustomContractAdded and is used to iterate over the raw logs and unpacked data for CustomContractAdded events raised by the Etherbase contract.
type EtherbaseCustomContractAddedIterator struct {
	Event *EtherbaseCustomContractAdded // Event containing the contract specifics and raw log

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
func (it *EtherbaseCustomContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseCustomContractAdded)
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
		it.Event = new(EtherbaseCustomContractAdded)
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
func (it *EtherbaseCustomContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseCustomContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseCustomContractAdded represents a CustomContractAdded event raised by the Etherbase contract.
type EtherbaseCustomContractAdded struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCustomContractAdded is a free log retrieval operation binding the contract event 0x996a8af2e36bc142de1bccafc39ce86c4cd7092058777ed39fd8e795f15e7c9a.
//
// Solidity: event CustomContractAdded(address indexed contractAddress)
func (_Etherbase *EtherbaseFilterer) FilterCustomContractAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*EtherbaseCustomContractAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Etherbase.contract.FilterLogs(opts, "CustomContractAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseCustomContractAddedIterator{contract: _Etherbase.contract, event: "CustomContractAdded", logs: logs, sub: sub}, nil
}

// WatchCustomContractAdded is a free log subscription operation binding the contract event 0x996a8af2e36bc142de1bccafc39ce86c4cd7092058777ed39fd8e795f15e7c9a.
//
// Solidity: event CustomContractAdded(address indexed contractAddress)
func (_Etherbase *EtherbaseFilterer) WatchCustomContractAdded(opts *bind.WatchOpts, sink chan<- *EtherbaseCustomContractAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Etherbase.contract.WatchLogs(opts, "CustomContractAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseCustomContractAdded)
				if err := _Etherbase.contract.UnpackLog(event, "CustomContractAdded", log); err != nil {
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

// ParseCustomContractAdded is a log parse operation binding the contract event 0x996a8af2e36bc142de1bccafc39ce86c4cd7092058777ed39fd8e795f15e7c9a.
//
// Solidity: event CustomContractAdded(address indexed contractAddress)
func (_Etherbase *EtherbaseFilterer) ParseCustomContractAdded(log types.Log) (*EtherbaseCustomContractAdded, error) {
	event := new(EtherbaseCustomContractAdded)
	if err := _Etherbase.contract.UnpackLog(event, "CustomContractAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseCustomContractDeletedIterator is returned from FilterCustomContractDeleted and is used to iterate over the raw logs and unpacked data for CustomContractDeleted events raised by the Etherbase contract.
type EtherbaseCustomContractDeletedIterator struct {
	Event *EtherbaseCustomContractDeleted // Event containing the contract specifics and raw log

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
func (it *EtherbaseCustomContractDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseCustomContractDeleted)
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
		it.Event = new(EtherbaseCustomContractDeleted)
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
func (it *EtherbaseCustomContractDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseCustomContractDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseCustomContractDeleted represents a CustomContractDeleted event raised by the Etherbase contract.
type EtherbaseCustomContractDeleted struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCustomContractDeleted is a free log retrieval operation binding the contract event 0xce2cd4a1ec306fda51c35ba3ed91f20df734a8097698e844508959656ccb5118.
//
// Solidity: event CustomContractDeleted(address indexed contractAddress)
func (_Etherbase *EtherbaseFilterer) FilterCustomContractDeleted(opts *bind.FilterOpts, contractAddress []common.Address) (*EtherbaseCustomContractDeletedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Etherbase.contract.FilterLogs(opts, "CustomContractDeleted", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseCustomContractDeletedIterator{contract: _Etherbase.contract, event: "CustomContractDeleted", logs: logs, sub: sub}, nil
}

// WatchCustomContractDeleted is a free log subscription operation binding the contract event 0xce2cd4a1ec306fda51c35ba3ed91f20df734a8097698e844508959656ccb5118.
//
// Solidity: event CustomContractDeleted(address indexed contractAddress)
func (_Etherbase *EtherbaseFilterer) WatchCustomContractDeleted(opts *bind.WatchOpts, sink chan<- *EtherbaseCustomContractDeleted, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Etherbase.contract.WatchLogs(opts, "CustomContractDeleted", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseCustomContractDeleted)
				if err := _Etherbase.contract.UnpackLog(event, "CustomContractDeleted", log); err != nil {
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

// ParseCustomContractDeleted is a log parse operation binding the contract event 0xce2cd4a1ec306fda51c35ba3ed91f20df734a8097698e844508959656ccb5118.
//
// Solidity: event CustomContractDeleted(address indexed contractAddress)
func (_Etherbase *EtherbaseFilterer) ParseCustomContractDeleted(log types.Log) (*EtherbaseCustomContractDeleted, error) {
	event := new(EtherbaseCustomContractDeleted)
	if err := _Etherbase.contract.UnpackLog(event, "CustomContractDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseCustomContractSchemaAddedIterator is returned from FilterCustomContractSchemaAdded and is used to iterate over the raw logs and unpacked data for CustomContractSchemaAdded events raised by the Etherbase contract.
type EtherbaseCustomContractSchemaAddedIterator struct {
	Event *EtherbaseCustomContractSchemaAdded // Event containing the contract specifics and raw log

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
func (it *EtherbaseCustomContractSchemaAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseCustomContractSchemaAdded)
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
		it.Event = new(EtherbaseCustomContractSchemaAdded)
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
func (it *EtherbaseCustomContractSchemaAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseCustomContractSchemaAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseCustomContractSchemaAdded represents a CustomContractSchemaAdded event raised by the Etherbase contract.
type EtherbaseCustomContractSchemaAdded struct {
	ContractAddress common.Address
	EventName       string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCustomContractSchemaAdded is a free log retrieval operation binding the contract event 0x8c7d8ad02950ea5bd57f54af9c47ca84706c82e280cada0c04a2da9f2bb4818c.
//
// Solidity: event CustomContractSchemaAdded(address indexed contractAddress, string eventName)
func (_Etherbase *EtherbaseFilterer) FilterCustomContractSchemaAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*EtherbaseCustomContractSchemaAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Etherbase.contract.FilterLogs(opts, "CustomContractSchemaAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseCustomContractSchemaAddedIterator{contract: _Etherbase.contract, event: "CustomContractSchemaAdded", logs: logs, sub: sub}, nil
}

// WatchCustomContractSchemaAdded is a free log subscription operation binding the contract event 0x8c7d8ad02950ea5bd57f54af9c47ca84706c82e280cada0c04a2da9f2bb4818c.
//
// Solidity: event CustomContractSchemaAdded(address indexed contractAddress, string eventName)
func (_Etherbase *EtherbaseFilterer) WatchCustomContractSchemaAdded(opts *bind.WatchOpts, sink chan<- *EtherbaseCustomContractSchemaAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Etherbase.contract.WatchLogs(opts, "CustomContractSchemaAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseCustomContractSchemaAdded)
				if err := _Etherbase.contract.UnpackLog(event, "CustomContractSchemaAdded", log); err != nil {
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
func (_Etherbase *EtherbaseFilterer) ParseCustomContractSchemaAdded(log types.Log) (*EtherbaseCustomContractSchemaAdded, error) {
	event := new(EtherbaseCustomContractSchemaAdded)
	if err := _Etherbase.contract.UnpackLog(event, "CustomContractSchemaAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceCreatedIterator is returned from FilterSourceCreated and is used to iterate over the raw logs and unpacked data for SourceCreated events raised by the Etherbase contract.
type EtherbaseSourceCreatedIterator struct {
	Event *EtherbaseSourceCreated // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceCreated)
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
		it.Event = new(EtherbaseSourceCreated)
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
func (it *EtherbaseSourceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceCreated represents a SourceCreated event raised by the Etherbase contract.
type EtherbaseSourceCreated struct {
	SourceAddress common.Address
	Owner         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSourceCreated is a free log retrieval operation binding the contract event 0x04ff85b5d2c64e18db6e3d3f6814a460cd9736cf5eed34eb285b1d0a9233370d.
//
// Solidity: event SourceCreated(address indexed sourceAddress, address indexed owner)
func (_Etherbase *EtherbaseFilterer) FilterSourceCreated(opts *bind.FilterOpts, sourceAddress []common.Address, owner []common.Address) (*EtherbaseSourceCreatedIterator, error) {

	var sourceAddressRule []interface{}
	for _, sourceAddressItem := range sourceAddress {
		sourceAddressRule = append(sourceAddressRule, sourceAddressItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Etherbase.contract.FilterLogs(opts, "SourceCreated", sourceAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceCreatedIterator{contract: _Etherbase.contract, event: "SourceCreated", logs: logs, sub: sub}, nil
}

// WatchSourceCreated is a free log subscription operation binding the contract event 0x04ff85b5d2c64e18db6e3d3f6814a460cd9736cf5eed34eb285b1d0a9233370d.
//
// Solidity: event SourceCreated(address indexed sourceAddress, address indexed owner)
func (_Etherbase *EtherbaseFilterer) WatchSourceCreated(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceCreated, sourceAddress []common.Address, owner []common.Address) (event.Subscription, error) {

	var sourceAddressRule []interface{}
	for _, sourceAddressItem := range sourceAddress {
		sourceAddressRule = append(sourceAddressRule, sourceAddressItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Etherbase.contract.WatchLogs(opts, "SourceCreated", sourceAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceCreated)
				if err := _Etherbase.contract.UnpackLog(event, "SourceCreated", log); err != nil {
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

// ParseSourceCreated is a log parse operation binding the contract event 0x04ff85b5d2c64e18db6e3d3f6814a460cd9736cf5eed34eb285b1d0a9233370d.
//
// Solidity: event SourceCreated(address indexed sourceAddress, address indexed owner)
func (_Etherbase *EtherbaseFilterer) ParseSourceCreated(log types.Log) (*EtherbaseSourceCreated, error) {
	event := new(EtherbaseSourceCreated)
	if err := _Etherbase.contract.UnpackLog(event, "SourceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceMetaData contains all meta data concerning the EtherbaseSource contract.
var EtherbaseSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EventNameAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EventNameNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfTopics\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"InvalidDataEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDataType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIdentity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PathNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyIndexedargs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyTopics\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"EthDBPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NewSegment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"SchemaRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"initialRoles\",\"type\":\"uint8[]\"}],\"name\":\"createWalletIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"deleteIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"ArgumentTopics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitEvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"argumentTopics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEtherbaseSource.BatchEmitEvent[]\",\"name\":\"events\",\"type\":\"tuple[]\"}],\"name\":\"emitEvents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventSchemas\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllIdentities\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structRoleControl.IdentityView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSegments\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structEtherDatabaseLib.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getEventSchemaFromId\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getEventSchemaFromName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"numIndexedArgs\",\"type\":\"uint8\"}],\"internalType\":\"structEventSchema\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getOrCreateSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredEventNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"segment\",\"type\":\"string\"}],\"name\":\"getSegmentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"hasEntry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"idToName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"eventTopic\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"argType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isIndexed\",\"type\":\"bool\"}],\"internalType\":\"structArgument[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"registerEventSchema\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"removeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"},{\"internalType\":\"enumEtherDatabaseLib.DataType\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEtherDatabaseLib.BatchSetValue[]\",\"name\":\"values\",\"type\":\"tuple[]\"}],\"name\":\"setValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161404a38038061404a83398101604081905261002e916103c3565b6100378261005e565b6001600455600380546001600160a01b0319166001600160a01b038316179055505061041c565b600280546001600160a01b0319166001600160a01b03831617905560408051600480825260a082019092525f91602082016080803683370190505090506002815f815181106100af576100af6103f4565b602002602001019060038111156100c8576100c8610408565b908160038111156100db576100db610408565b815250505f816001815181106100f3576100f36103f4565b6020026020010190600381111561010c5761010c610408565b9081600381111561011f5761011f610408565b81525050600181600281518110610138576101386103f4565b6020026020010190600381111561015157610151610408565b9081600381111561016457610164610408565b8152505060038160038151811061017d5761017d6103f4565b6020026020010190600381111561019657610196610408565b908160038111156101a9576101a9610408565b9052506101b682826101ba565b5050565b6001600160a01b0382165f9081526020819052604090205460ff16156101f3576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602081905260408120805460ff19166001908117909155805480820182559082527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169092179091555b81518110156102915761028983838381518110610276576102766103f4565b60200260200101516102c960201b60201c565b600101610257565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a25050565b6001600160a01b0382165f9081526020819052604090205460ff166103015760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208181526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a91820219169083600381111561035957610359610408565b021790555081600381111561037057610370610408565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a3505050565b80516001600160a01b03811681146103be575f5ffd5b919050565b5f5f604083850312156103d4575f5ffd5b6103dd836103a8565b91506103eb602084016103a8565b90509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b613c21806104295f395ff3fe608060405234801561000f575f5ffd5b506004361061016d575f3560e01c806364ad6c88116100d95780638da5cb5b11610093578063c76853b81161006e578063c76853b814610370578063d50d542e14610390578063dced4783146103a3578063f30826f0146103b6575f5ffd5b80638da5cb5b14610337578063a4dbe8a01461034a578063a8d29d1d1461035d575f5ffd5b806364ad6c88146102c357806369da55d2146102d657806371934ef0146102e95780637eeede35146102fc5780638998fbe8146103115780638c71045614610324575f5ffd5b80633a5381b51161012a5780633a5381b51461020e5780633e840236146102395780633f54d8bf1461024c5780634151df611461026157806349ab123d1461028257806360778d95146102a2575f5ffd5b80630277e023146101715780630761b8d014610199578063077d3c03146101ae57806317be85c3146101c15780631bfa8601146101d65780632f2545d2146101eb575b5f5ffd5b61018461017f366004612df0565b6103be565b60405190151581526020015b60405180910390f35b6101ac6101a7366004612eac565b61042d565b005b6101ac6101bc366004612f4f565b610475565b6101c96104dc565b6040516101909190612fd6565b6101de6107ae565b6040516101909190613080565b6101fe6101f93660046130cb565b61080e565b6040516101909493929190613117565b600354610221906001600160a01b031681565b6040516001600160a01b039091168152602001610190565b6101ac610247366004612f4f565b610951565b6102546109b3565b6040516101909190613156565b61027461026f3660046131ad565b610a87565b604051908152602001610190565b6102956102903660046131ad565b610a9b565b60405161019091906131eb565b6102b56102b03660046132d8565b610e8e565b60405161019092919061330a565b6101846102d13660046132d8565b610f3f565b6101ac6102e43660046132d8565b610f79565b6102746102f73660046131ad565b61105c565b610304611086565b6040516101909190613329565b6101ac61031f3660046132d8565b6111ee565b6101846103323660046132d8565b61132a565b600254610221906001600160a01b031681565b6101ac6103583660046133da565b611427565b6101ac61036b36600461347f565b6116cc565b61038361037e3660046130cb565b611731565b6040516101909190613498565b61029561039e3660046131ad565b6117d4565b6101ac6103b13660046134aa565b611b4a565b610254611bac565b335f90815260208190526040812054819060ff166103ef5760405163c597feeb60e01b815260040160405180910390fd5b6103f93382611c77565b610415576040516282b42960e81b815260040160405180910390fd5b6104228787878787611d11565b979650505050505050565b61046e85858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611e4392505050565b5050505050565b335f9081526020819052604090205460029060ff166104a75760405163c597feeb60e01b815260040160405180910390fd5b6104b13382611c77565b6104cd576040516282b42960e81b815260040160405180910390fd5b6104d783836121d8565b505050565b60605f805b60085481101561052f57600881815481106104fe576104fe61356a565b5f91825260209091206003600490920201015460ff1615610527578161052381613592565b9250505b6001016104e1565b505f816001600160401b0381111561054957610549612d54565b6040519080825280602002602001820160405280156105a457816020015b610591604080516080810190915260608152602081015f8152606060208201525f60409091015290565b8152602001906001900390816105675790505b5090505f805b6008548110156107a557600881815481106105c7576105c761356a565b5f91825260209091206003600490920201015460ff161561079d57600881815481106105f5576105f561356a565b905f5260205f2090600402016040518060800160405290815f8201805461061b906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610647906135aa565b80156106925780601f1061066957610100808354040283529160200191610692565b820191905f5260205f20905b81548152906001019060200180831161067557829003601f168201915b5050509183525050600182015460209091019060ff1660058111156106b9576106b9612fae565b60058111156106ca576106ca612fae565b81526020016002820180546106de906135aa565b80601f016020809104026020016040519081016040528092919081815260200182805461070a906135aa565b80156107555780601f1061072c57610100808354040283529160200191610755565b820191905f5260205f20905b81548152906001019060200180831161073857829003601f168201915b50505091835250506003919091015460ff16151560209091015283518490849081106107835761078361356a565b6020026020010181905250818061079990613592565b9250505b6001016105aa565b50909392505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561080457602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116107e6575b5050505050905090565b8051602081830181018051600b82529282019190930120915280548190610834906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610860906135aa565b80156108ab5780601f10610882576101008083540402835291602001916108ab565b820191905f5260205f20905b81548152906001019060200180831161088e57829003601f168201915b5050505050908060010180546108c0906135aa565b80601f01602080910402602001604051908101604052809291908181526020018280546108ec906135aa565b80156109375780601f1061090e57610100808354040283529160200191610937565b820191905f5260205f20905b81548152906001019060200180831161091a57829003601f168201915b50505050600383015460049093015491929160ff16905084565b335f9081526020819052604090205460029060ff166109835760405163c597feeb60e01b815260040160405180910390fd5b61098d3382611c77565b6109a9576040516282b42960e81b815260040160405180910390fd5b6104d783836123a1565b60606006805480602002602001604051908101604052809291908181526020015f905b82821015610a7e578382905f5260205f200180546109f3906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1f906135aa565b8015610a6a5780601f10610a4157610100808354040283529160200191610a6a565b820191905f5260205f20905b815481529060010190602001808311610a4d57829003601f168201915b5050505050815260200190600101906109d6565b50505050905090565b5f610a928383612480565b90505b92915050565b610ace6040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b5f600c8484604051610ae19291906135dc565b90815260200160405180910390208054610afa906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610b26906135aa565b8015610b715780601f10610b4857610100808354040283529160200191610b71565b820191905f5260205f20905b815481529060010190602001808311610b5457829003601f168201915b5050505050905080515f03610b9957604051634fde869d60e01b815260040160405180910390fd5b600b81604051610ba99190613602565b90815260200160405180910390206040518060a00160405290815f82018054610bd1906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfd906135aa565b8015610c485780601f10610c1f57610100808354040283529160200191610c48565b820191905f5260205f20905b815481529060010190602001808311610c2b57829003601f168201915b50505050508152602001600182018054610c61906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8d906135aa565b8015610cd85780601f10610caf57610100808354040283529160200191610cd8565b820191905f5260205f20905b815481529060010190602001808311610cbb57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015610e66578382905f5260205f2090600302016040518060600160405290815f82018054610d35906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610d61906135aa565b8015610dac5780601f10610d8357610100808354040283529160200191610dac565b820191905f5260205f20905b815481529060010190602001808311610d8f57829003601f168201915b50505050508152602001600182018054610dc5906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610df1906135aa565b8015610e3c5780601f10610e1357610100808354040283529160200191610e3c565b820191905f5260205f20905b815481529060010190602001808311610e1f57829003601f168201915b50505091835250506002919091015460ff1615156020918201529082526001929092019101610d05565b505050908252506003820154602082015260049091015460ff16604090910152949350505050565b5f60605f610e9c8585612597565b9050600a81604051610eae9190613602565b9081526040519081900360200190205460ff16610edd57505060408051602081019091525f8082529150610f38565b5f610ee986865f61265a565b90506008600983604051610efd9190613602565b90815260200160405180910390205481548110610f1c57610f1c61356a565b5f91825260209091206001600490920201015460ff1693509150505b9250929050565b5f5f610f4b8484612597565b9050600a81604051610f5d9190613602565b9081526040519081900360200190205460ff1691505092915050565b5f5b818110156104d757611054838383818110610f9857610f9861356a565b9050602002810190610faa919061360d565b610fb4908061362b565b858585818110610fc657610fc661356a565b9050602002810190610fd8919061360d565b610fe9906040810190602001613670565b868686818110610ffb57610ffb61356a565b905060200281019061100d919061360d565b61101b906040810190613689565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611e4392505050565b600101610f7b565b5f6005838360405161106f9291906135dc565b908152602001604051809103902054905092915050565b6001546060905f906001600160401b038111156110a5576110a5612d54565b6040519080825280602002602001820160405280156110ea57816020015b604080518082019091525f8152606060208201528152602001906001900390816110c35790505b5090505f5b6001548110156111e8575f6001828154811061110d5761110d61356a565b5f9182526020808320909101546001600160a01b0316808352828252604092839020835180850185528281526001820180548651818702810187019097528087529396509194909384810193919291908301828280156111b957602002820191905f5260205f20905f905b82829054906101000a900460ff16600381111561119757611197612fae565b8152602060019283018181049485019490930390920291018084116111785790505b50505050508152508484815181106111d3576111d361356a565b602090810291909101015250506001016110ef565b50919050565b5f6111f98383612597565b9050600a8160405161120b9190613602565b9081526040519081900360200190205460ff1661123b576040516357509b2760e11b815260040160405180910390fd5b5f60098260405161124c9190613602565b90815260200160405180910390205490505f600882815481106112715761127161356a565b905f5260205f2090600402016003015f6101000a81548160ff0219169083151502179055506009826040516112a69190613602565b90815260200160405180910390205f9055600a826040516112c79190613602565b9081526040805160209281900383018120805460ff1916905591820181525f80835290517fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a9261131c928892889291906136f3565b60405180910390a150505050565b335f90815260208190526040812054819060ff1661135b5760405163c597feeb60e01b815260040160405180910390fd5b6113653382611c77565b611381576040516282b42960e81b815260040160405180910390fd5b5f5b8381101561141c573685858381811061139e5761139e61356a565b90506020028101906113b0919061360d565b90506114126113bf8280613689565b6113cc602085018561362b565b6113d96040870187613689565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611d1192505050565b5050600101611383565b506001949350505050565b335f9081526020819052604090205460019060ff166114595760405163c597feeb60e01b815260040160405180910390fd5b6114633382611c77565b61147f576040516282b42960e81b815260040160405180910390fd5b600b88886040516114919291906135dc565b90815260405190819003602001902080546114ab906135aa565b1590506114cb57604051630268c2d760e51b815260040160405180910390fd5b5f805b83811015611526578484828181106114e8576114e861356a565b90506020028101906114fa919061360d565b61150b9060608101906040016137b9565b1561151e578161151a816137d4565b9250505b6001016114ce565b5060038160ff16111561154b57604051626f3b7d60e31b815260040160405180910390fd5b5f600b8a8a60405161155e9291906135dc565b90815260405190819003602001902090508061157b8a8c8361384a565b506001810161158b888a8361384a565b50600381018690555f5b848110156115f057816002018686838181106115b3576115b361356a565b90506020028101906115c5919061360d565b81546001810183555f92835260209092209091600302016115e6828261390a565b5050600101611595565b5060048101805460ff191660ff8416179055600d80546001810182555f919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb50161163e8a8c8361384a565b508989600c8a8a6040516116539291906135dc565b9081526020016040518091039020918261166e92919061384a565b50898960405161167f9291906135dc565b60405180910390207f4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de89896040516116b8929190613a0b565b60405180910390a250505050505050505050565b335f9081526020819052604090205460029060ff166116fe5760405163c597feeb60e01b815260040160405180910390fd5b6117083382611c77565b611724576040516282b42960e81b815260040160405180910390fd5b61172d826128da565b5050565b8051602081830181018051600c8252928201919093012091528054611755906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611781906135aa565b80156117cc5780601f106117a3576101008083540402835291602001916117cc565b820191905f5260205f20905b8154815290600101906020018083116117af57829003601f168201915b505050505081565b6118076040518060a001604052806060815260200160608152602001606081526020015f81526020015f60ff1681525090565b600b83836040516118199291906135dc565b9081526040519081900360200190208054611833906135aa565b90505f0361185457604051634fde869d60e01b815260040160405180910390fd5b600b83836040516118669291906135dc565b90815260200160405180910390206040518060a00160405290815f8201805461188e906135aa565b80601f01602080910402602001604051908101604052809291908181526020018280546118ba906135aa565b80156119055780601f106118dc57610100808354040283529160200191611905565b820191905f5260205f20905b8154815290600101906020018083116118e857829003601f168201915b5050505050815260200160018201805461191e906135aa565b80601f016020809104026020016040519081016040528092919081815260200182805461194a906135aa565b80156119955780601f1061196c57610100808354040283529160200191611995565b820191905f5260205f20905b81548152906001019060200180831161197857829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020015f905b82821015611b23578382905f5260205f2090600302016040518060600160405290815f820180546119f2906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611a1e906135aa565b8015611a695780601f10611a4057610100808354040283529160200191611a69565b820191905f5260205f20905b815481529060010190602001808311611a4c57829003601f168201915b50505050508152602001600182018054611a82906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611aae906135aa565b8015611af95780601f10611ad057610100808354040283529160200191611af9565b820191905f5260205f20905b815481529060010190602001808311611adc57829003601f168201915b50505091835250506002919091015460ff16151560209182015290825260019290920191016119c2565b505050908252506003820154602082015260049091015460ff166040909101529392505050565b335f9081526020819052604090205460029060ff16611b7c5760405163c597feeb60e01b815260040160405180910390fd5b611b863382611c77565b611ba2576040516282b42960e81b815260040160405180910390fd5b6104d78383612a73565b6060600d805480602002602001604051908101604052809291908181526020015f905b82821015610a7e578382905f5260205f20018054611bec906135aa565b80601f0160208091040260200160405190810160405280929190818152602001828054611c18906135aa565b8015611c635780601f10611c3a57610100808354040283529160200191611c63565b820191905f5260205f20905b815481529060010190602001808311611c4657829003601f168201915b505050505081526020019060010190611bcf565b6001600160a01b0382165f908152602081905260408120600101815b8154811015611d0757836003811115611cae57611cae612fae565b828281548110611cc057611cc061356a565b905f5260205f2090602091828204019190069054906101000a900460ff166003811115611cef57611cef612fae565b03611cff57600192505050610a95565b600101611c93565b505f949350505050565b5f600b8686604051611d249291906135dc565b9081526040519081900360200190208054611d3e906135aa565b90505f03611d5f57604051634fde869d60e01b815260040160405180910390fd5b5f600b8787604051611d729291906135dc565b908152604051908190036020019020600380820154919250851115611daa5760405163321fc7cf60e11b815260040160405180910390fd5b600482015460ff168514611dd157604051631f73a41360e01b815260040160405180910390fd5b835160208501868015611dfb5760018114611e045760028114611e0f5760038114611e1f57611e30565b838383a1611e30565b8835848484a2611e30565b60208901358935858585a3611e30565b604089013560208a01358a35868686a45b5050600193505050505b95945050505050565b5f611e4e8585612b7c565b90505f836005811115611e6357611e63612fae565b03611f9757600a81604051611e789190613602565b9081526040519081900360200190205460ff1615611f35575f600982604051611ea19190613602565b90815260200160405180910390205490505f60088281548110611ec657611ec661356a565b905f5260205f2090600402016003015f6101000a81548160ff021916908315150217905550600982604051611efb9190613602565b90815260200160405180910390205f9055600a82604051611f1c9190613602565b908152604051908190036020019020805460ff19169055505b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a858560405180602001604052805f8152505f6005811115611f7957611f79612fae565b604051611f8994939291906136f3565b60405180910390a1506121d2565b600a81604051611fa79190613602565b9081526040519081900360200190205460ff1615612043575f600982604051611fd09190613602565b90815260200160405180910390205490505f60088281548110611ff557611ff561356a565b905f5260205f2090600402019050838160020190816120149190613a1e565b5060018082018054879260ff199091169083600581111561203757612037612fae565b02179055505050612182565b5f604051806080016040528083815260200185600581111561206757612067612fae565b815260208101859052600160409091018190526008805491820181555f528151919250829160049091027ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3019081906120c09082613a1e565b50602082015160018083018054909160ff19909116908360058111156120e8576120e8612fae565b0217905550604082015160028201906121019082613a1e565b50606091909101516003909101805460ff191691151591909117905560085461212c90600190613ad3565b60098360405161213c9190613602565b9081526020016040518091039020819055506001600a836040516121609190613602565b908152604051908190036020019020805491151560ff19909216919091179055505b7fe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a8585848660058111156121b8576121b8612fae565b6040516121c894939291906136f3565b60405180910390a1505b50505050565b6001600160a01b0382165f9081526020819052604090205460ff166122105760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f9081526020819052604081209060018201905b815481101561046e5783600381111561224957612249612fae565b82828154811061225b5761225b61356a565b905f5260205f2090602091828204019190069054906101000a900460ff16600381111561228a5761228a612fae565b03612399578154829061229f90600190613ad3565b815481106122af576122af61356a565b905f5260205f2090602091828204019190069054906101000a900460ff168282815481106122df576122df61356a565b905f5260205f2090602091828204019190066101000a81548160ff0219169083600381111561231057612310612fae565b02179055508180548061232557612325613ae6565b600190038181905f5260205f2090602091828204019190066101000a81549060ff0219169055905583600381111561235f5761235f612fae565b6040516001600160a01b038716907f34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769905f90a35050505050565b60010161222e565b6001600160a01b0382165f9081526020819052604090205460ff166123d95760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208181526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a91820219169083600381111561243157612431612fae565b021790555081600381111561244857612448612fae565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a3505050565b5f5f600584846040516124949291906135dc565b9081526040805160209281900383019020545f818152600790935291205490915060ff16156124c4579050610a95565b600480549081905f6124d583613592565b919050555080600586866040516124ed9291906135dc565b90815260405190819003602001902055600680546001810182555f919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0161253985878361384a565b505f8181526007602052604090819020805460ff191660011790555181907f4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778906125869088908890613a0b565b60405180910390a29150610a959050565b604080515f808252602082019092526060915b8381101561262e575f60058686848181106125c7576125c761356a565b90506020028101906125d99190613689565b6040516125e79291906135dc565b90815260200160405180910390205490508261260282612bfc565b604051602001612613929190613afa565b60408051601f198184030181529190529250506001016125aa565b506040516126429082905f90602001613b0e565b60408051808303601f19018152919052949350505050565b60605f6126678585612597565b9050600a816040516126799190613602565b9081526040519081900360200190205460ff166126a9576040516357509b2760e11b815260040160405180910390fd5b5f60086009836040516126bc9190613602565b908152602001604051809103902054815481106126db576126db61356a565b5f9182526020909120600490910201600381015490915060ff16612712576040516357509b2760e11b815260040160405180910390fd5b5f84600581111561272557612725612fae565b1415801561275c575083600581111561274057612740612fae565b600182015460ff16600581111561275957612759612fae565b14155b1561279057600181015460405163254c060560e21b815261278791869160ff90911690600401613b32565b60405180910390fd5b5f8460058111156127a3576127a3612fae565b141580156127c3575060058460058111156127c0576127c0612fae565b14155b15612843576003546040516323bdb73560e11b81526001600160a01b039091169063477b6e6a906127fd9060028501908890600401613b4d565b5f6040518083038186803b158015612813575f5ffd5b505afa925050508015612824575060015b6128435783604051631e9c993560e21b81526004016127879190613bdd565b806002018054612852906135aa565b80601f016020809104026020016040519081016040528092919081815260200182805461287e906135aa565b80156128c95780601f106128a0576101008083540402835291602001916128c9565b820191905f5260205f20905b8154815290600101906020018083116128ac57829003601f168201915b5050505050925050505b9392505050565b6001600160a01b0381165f90815260208190526040812060018101549091036129165760405163178423dd60e01b815260040160405180910390fd5b6001600160a01b0382165f908152602081905260408120805460ff19168155906129436001830182612c9a565b505f90505b600154811015612a3b57826001600160a01b03166001828154811061296f5761296f61356a565b5f918252602090912001546001600160a01b031603612a335760018054612997908290613ad3565b815481106129a7576129a761356a565b5f91825260209091200154600180546001600160a01b0390921691839081106129d2576129d261356a565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506001805480612a0e57612a0e613ae6565b5f8281526020902081015f1990810180546001600160a01b0319169055019055612a3b565b600101612948565b506040516001600160a01b038316907fc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8905f90a25050565b6001600160a01b0382165f9081526020819052604090205460ff1615612aac576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602081905260408120805460ff19166001908117909155805480820182559082527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169092179091555b8151811015612b4457612b3c83838381518110612b2f57612b2f61356a565b60200260200101516123a1565b600101612b10565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a25050565b604080515f808252602082019092526060915b8381101561262e575f612bc4868684818110612bad57612bad61356a565b9050602002810190612bbf9190613689565b612480565b905082612bd082612bfc565b604051602001612be1929190613afa565b60408051601f19818403018152919052925050600101612b8f565b6060815f03612c27576040515f60208201526021016040516020818303038152906040529050919050565b6060825b60808110612c69578181607f1660801760f81b604051602001612c4f929190613b0e565b60408051601f19818403018152919052915060071c612c2b565b8181607f1660f81b604051602001612c82929190613b0e565b60408051601f19818403018152919052949350505050565b5080545f8255601f0160209004905f5260205f2090810190612cbc9190612cbf565b50565b5b80821115612cd3575f8155600101612cc0565b5090565b5f5f83601f840112612ce7575f5ffd5b5081356001600160401b03811115612cfd575f5ffd5b602083019150836020828501011115610f38575f5ffd5b5f5f83601f840112612d24575f5ffd5b5081356001600160401b03811115612d3a575f5ffd5b6020830191508360208260051b8501011115610f38575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715612d9057612d90612d54565b604052919050565b5f5f6001600160401b03841115612db157612db1612d54565b50601f8301601f1916602001612dc681612d68565b915050828152838383011115612dda575f5ffd5b828260208301375f602084830101529392505050565b5f5f5f5f5f60608688031215612e04575f5ffd5b85356001600160401b03811115612e19575f5ffd5b612e2588828901612cd7565b90965094505060208601356001600160401b03811115612e43575f5ffd5b612e4f88828901612d14565b90945092505060408601356001600160401b03811115612e6d575f5ffd5b8601601f81018813612e7d575f5ffd5b612e8c88823560208401612d98565b9150509295509295909350565b803560068110612ea7575f5ffd5b919050565b5f5f5f5f5f60608688031215612ec0575f5ffd5b85356001600160401b03811115612ed5575f5ffd5b612ee188828901612d14565b9096509450612ef4905060208701612e99565b925060408601356001600160401b03811115612f0e575f5ffd5b612f1a88828901612cd7565b969995985093965092949392505050565b80356001600160a01b0381168114612ea7575f5ffd5b803560048110612ea7575f5ffd5b5f5f60408385031215612f60575f5ffd5b612f6983612f2b565b9150612f7760208401612f41565b90509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b5f52602160045260245ffd5b60068110612fd257612fd2612fae565b9052565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561307457603f1987860301845281518051608087526130226080880182612f80565b905060208201516130366020890182612fc2565b506040820151878203604089015261304e8282612f80565b606093840151151598909301979097525094506020938401939190910190600101612ffc565b50929695505050505050565b602080825282518282018190525f918401906040840190835b818110156130c05783516001600160a01b0316835260209384019390920191600101613099565b509095945050505050565b5f602082840312156130db575f5ffd5b81356001600160401b038111156130f0575f5ffd5b8201601f81018413613100575f5ffd5b61310f84823560208401612d98565b949350505050565b608081525f6131296080830187612f80565b828103602084015261313b8187612f80565b91505083604083015260ff8316606083015295945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561307457603f19878603018452613198858351612f80565b9450602093840193919091019060010161317c565b5f5f602083850312156131be575f5ffd5b82356001600160401b038111156131d3575f5ffd5b6131df85828601612cd7565b90969095509350505050565b602081525f825160a0602084015261320660c0840182612f80565b90506020840151601f198483030160408501526132238282612f80565b6040860151858203601f190160608701528051808352919350602090810192508084019190600582901b8501015f5b828110156132b757601f1986830301845284518051606084526132786060850182612f80565b9050602082015184820360208601526132918282612f80565b604093840151151595909301949094525060209586019594909401939150600101613252565b50606088015160808801526080880151945061042260a088018660ff169052565b5f5f602083850312156132e9575f5ffd5b82356001600160401b038111156132fe575f5ffd5b6131df85828601612d14565b6133148184612fc2565b604060208201525f61310f6040830184612f80565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561307457868503603f19018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101905f9060608801905b808310156133c2578351600481106133aa576133aa612fae565b82526020938401936001939093019290910190613390565b5096505050602093840193919091019060010161334f565b5f5f5f5f5f5f5f6080888a0312156133f0575f5ffd5b87356001600160401b03811115613405575f5ffd5b6134118a828b01612cd7565b90985096505060208801356001600160401b0381111561342f575f5ffd5b61343b8a828b01612cd7565b9096509450506040880135925060608801356001600160401b03811115613460575f5ffd5b61346c8a828b01612d14565b989b979a50959850939692959293505050565b5f6020828403121561348f575f5ffd5b610a9282612f2b565b602081525f610a926020830184612f80565b5f5f604083850312156134bb575f5ffd5b6134c483612f2b565b915060208301356001600160401b038111156134de575f5ffd5b8301601f810185136134ee575f5ffd5b80356001600160401b0381111561350757613507612d54565b8060051b61351760208201612d68565b91825260208184018101929081019088841115613532575f5ffd5b6020850194505b8385101561355b5761354a85612f41565b825260209485019490910190613539565b80955050505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f600182016135a3576135a361357e565b5060010190565b600181811c908216806135be57607f821691505b6020821081036111e857634e487b7160e01b5f52602260045260245ffd5b818382375f9101908152919050565b5f81518060208401855e5f93019283525090919050565b5f610a9282846135eb565b5f8235605e19833603018112613621575f5ffd5b9190910192915050565b5f5f8335601e19843603018112613640575f5ffd5b8301803591506001600160401b03821115613659575f5ffd5b6020019150600581901b3603821315610f38575f5ffd5b5f60208284031215613680575f5ffd5b610a9282612e99565b5f5f8335601e1984360301811261369e575f5ffd5b8301803591506001600160401b038211156136b7575f5ffd5b602001915036819003821315610f38575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606080825281018490525f6080600586901b830181019083018783601e1936839003015b8982101561378457868503607f190184528235818112613735575f5ffd5b8b016020810190356001600160401b03811115613750575f5ffd5b80360382131561375e575f5ffd5b6137698782846136cb565b96505050602083019250602084019350600182019150613717565b50505050828103602084015261379a8186612f80565b915050611e3a604083018460ff169052565b8015158114612cbc575f5ffd5b5f602082840312156137c9575f5ffd5b81356128d3816137ac565b5f60ff821660ff81036137e9576137e961357e565b60010192915050565b601f8211156104d757805f5260205f20601f840160051c810160208510156138175750805b601f840160051c820191505b8181101561046e575f8155600101613823565b5f19600383901b1c191660019190911b1790565b6001600160401b0383111561386157613861612d54565b6138758361386f83546135aa565b836137f2565b5f601f8411600181146138a1575f851561388f5750838201355b6138998682613836565b84555061046e565b5f83815260208120601f198716915b828110156138d057868501358255602094850194600190920191016138b0565b50868210156138ec575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f8135610a95816137ac565b6139148283613689565b6001600160401b0381111561392b5761392b612d54565b61393f8161393985546135aa565b856137f2565b5f601f82116001811461396b575f83156139595750838201355b6139638482613836565b8655506139c2565b5f85815260208120601f198516915b8281101561399a578685013582556020948501946001909201910161397a565b50848210156139b6575f1960f88660031b161c19848701351681555b505060018360011b0185555b505050506139d36020830183613689565b6139e181836001860161384a565b505061172d6139f2604084016138fe565b6002830160ff1981541660ff8315151681178255505050565b602081525f61310f6020830184866136cb565b81516001600160401b03811115613a3757613a37612d54565b613a4b81613a4584546135aa565b846137f2565b6020601f821160018114613a78575f8315613a665750848201515b613a708482613836565b85555061046e565b5f84815260208120601f198516915b82811015613aa75787850151825560209485019460019092019101613a87565b5084821015613ac457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b81810381811115610a9557610a9561357e565b634e487b7160e01b5f52603160045260245ffd5b5f61310f613b0883866135eb565b846135eb565b5f613b1982856135eb565b6001600160f81b03199390931683525050600101919050565b60408101613b408285612fc2565b6128d36020830184612fc2565b604081525f5f8454613b5e816135aa565b806040860152600182165f8114613b7c5760018114613b9857613bc9565b60ff1983166060870152606082151560051b8701019350613bc9565b875f5260205f205f5b83811015613bc057815488820160600152600190910190602001613ba1565b87016060019450505b505050809150506128d36020830184612fc2565b60208101610a958284612fc256fea26469706673582212202b5b229f00090a107f7e35504651973effbb65a775bad46745053d0513b056e564736f6c634300081c0033",
}

// EtherbaseSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherbaseSourceMetaData.ABI instead.
var EtherbaseSourceABI = EtherbaseSourceMetaData.ABI

// EtherbaseSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherbaseSourceMetaData.Bin instead.
var EtherbaseSourceBin = EtherbaseSourceMetaData.Bin

// DeployEtherbaseSource deploys a new Ethereum contract, binding an instance of EtherbaseSource to it.
func DeployEtherbaseSource(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, validator common.Address) (common.Address, *types.Transaction, *EtherbaseSource, error) {
	parsed, err := EtherbaseSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherbaseSourceBin), backend, _owner, validator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherbaseSource{EtherbaseSourceCaller: EtherbaseSourceCaller{contract: contract}, EtherbaseSourceTransactor: EtherbaseSourceTransactor{contract: contract}, EtherbaseSourceFilterer: EtherbaseSourceFilterer{contract: contract}}, nil
}

// EtherbaseSource is an auto generated Go binding around an Ethereum contract.
type EtherbaseSource struct {
	EtherbaseSourceCaller     // Read-only binding to the contract
	EtherbaseSourceTransactor // Write-only binding to the contract
	EtherbaseSourceFilterer   // Log filterer for contract events
}

// EtherbaseSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherbaseSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherbaseSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherbaseSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherbaseSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherbaseSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherbaseSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherbaseSourceSession struct {
	Contract     *EtherbaseSource  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherbaseSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherbaseSourceCallerSession struct {
	Contract *EtherbaseSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EtherbaseSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherbaseSourceTransactorSession struct {
	Contract     *EtherbaseSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EtherbaseSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherbaseSourceRaw struct {
	Contract *EtherbaseSource // Generic contract binding to access the raw methods on
}

// EtherbaseSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherbaseSourceCallerRaw struct {
	Contract *EtherbaseSourceCaller // Generic read-only contract binding to access the raw methods on
}

// EtherbaseSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherbaseSourceTransactorRaw struct {
	Contract *EtherbaseSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherbaseSource creates a new instance of EtherbaseSource, bound to a specific deployed contract.
func NewEtherbaseSource(address common.Address, backend bind.ContractBackend) (*EtherbaseSource, error) {
	contract, err := bindEtherbaseSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSource{EtherbaseSourceCaller: EtherbaseSourceCaller{contract: contract}, EtherbaseSourceTransactor: EtherbaseSourceTransactor{contract: contract}, EtherbaseSourceFilterer: EtherbaseSourceFilterer{contract: contract}}, nil
}

// NewEtherbaseSourceCaller creates a new read-only instance of EtherbaseSource, bound to a specific deployed contract.
func NewEtherbaseSourceCaller(address common.Address, caller bind.ContractCaller) (*EtherbaseSourceCaller, error) {
	contract, err := bindEtherbaseSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceCaller{contract: contract}, nil
}

// NewEtherbaseSourceTransactor creates a new write-only instance of EtherbaseSource, bound to a specific deployed contract.
func NewEtherbaseSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherbaseSourceTransactor, error) {
	contract, err := bindEtherbaseSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceTransactor{contract: contract}, nil
}

// NewEtherbaseSourceFilterer creates a new log filterer instance of EtherbaseSource, bound to a specific deployed contract.
func NewEtherbaseSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherbaseSourceFilterer, error) {
	contract, err := bindEtherbaseSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceFilterer{contract: contract}, nil
}

// bindEtherbaseSource binds a generic wrapper to an already deployed contract.
func bindEtherbaseSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherbaseSourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherbaseSource *EtherbaseSourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherbaseSource.Contract.EtherbaseSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherbaseSource *EtherbaseSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.EtherbaseSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherbaseSource *EtherbaseSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.EtherbaseSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherbaseSource *EtherbaseSourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherbaseSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherbaseSource *EtherbaseSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherbaseSource *EtherbaseSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.contract.Transact(opts, method, params...)
}

// EventSchemas is a free data retrieval call binding the contract method 0x2f2545d2.
//
// Solidity: function eventSchemas(string ) view returns(string name, string id, bytes32 eventTopic, uint8 numIndexedArgs)
func (_EtherbaseSource *EtherbaseSourceCaller) EventSchemas(opts *bind.CallOpts, arg0 string) (struct {
	Name           string
	Id             string
	EventTopic     [32]byte
	NumIndexedArgs uint8
}, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "eventSchemas", arg0)

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
func (_EtherbaseSource *EtherbaseSourceSession) EventSchemas(arg0 string) (struct {
	Name           string
	Id             string
	EventTopic     [32]byte
	NumIndexedArgs uint8
}, error) {
	return _EtherbaseSource.Contract.EventSchemas(&_EtherbaseSource.CallOpts, arg0)
}

// EventSchemas is a free data retrieval call binding the contract method 0x2f2545d2.
//
// Solidity: function eventSchemas(string ) view returns(string name, string id, bytes32 eventTopic, uint8 numIndexedArgs)
func (_EtherbaseSource *EtherbaseSourceCallerSession) EventSchemas(arg0 string) (struct {
	Name           string
	Id             string
	EventTopic     [32]byte
	NumIndexedArgs uint8
}, error) {
	return _EtherbaseSource.Contract.EventSchemas(&_EtherbaseSource.CallOpts, arg0)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_EtherbaseSource *EtherbaseSourceCaller) GetAllIdentities(opts *bind.CallOpts) ([]RoleControlIdentityView, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getAllIdentities")

	if err != nil {
		return *new([]RoleControlIdentityView), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleControlIdentityView)).(*[]RoleControlIdentityView)

	return out0, err

}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_EtherbaseSource *EtherbaseSourceSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _EtherbaseSource.Contract.GetAllIdentities(&_EtherbaseSource.CallOpts)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _EtherbaseSource.Contract.GetAllIdentities(&_EtherbaseSource.CallOpts)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherbaseSource *EtherbaseSourceCaller) GetAllSegments(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getAllSegments")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherbaseSource *EtherbaseSourceSession) GetAllSegments() ([]string, error) {
	return _EtherbaseSource.Contract.GetAllSegments(&_EtherbaseSource.CallOpts)
}

// GetAllSegments is a free data retrieval call binding the contract method 0x3f54d8bf.
//
// Solidity: function getAllSegments() view returns(string[])
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetAllSegments() ([]string, error) {
	return _EtherbaseSource.Contract.GetAllSegments(&_EtherbaseSource.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_EtherbaseSource *EtherbaseSourceCaller) GetAllWallets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getAllWallets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_EtherbaseSource *EtherbaseSourceSession) GetAllWallets() ([]common.Address, error) {
	return _EtherbaseSource.Contract.GetAllWallets(&_EtherbaseSource.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetAllWallets() ([]common.Address, error) {
	return _EtherbaseSource.Contract.GetAllWallets(&_EtherbaseSource.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherbaseSource *EtherbaseSourceCaller) GetEntries(opts *bind.CallOpts) ([]EtherDatabaseLibNode, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]EtherDatabaseLibNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]EtherDatabaseLibNode)).(*[]EtherDatabaseLibNode)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherbaseSource *EtherbaseSourceSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EtherbaseSource.Contract.GetEntries(&_EtherbaseSource.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes,uint8,bytes,bool)[])
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetEntries() ([]EtherDatabaseLibNode, error) {
	return _EtherbaseSource.Contract.GetEntries(&_EtherbaseSource.CallOpts)
}

// GetEventSchemaFromId is a free data retrieval call binding the contract method 0x49ab123d.
//
// Solidity: function getEventSchemaFromId(string id) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EtherbaseSource *EtherbaseSourceCaller) GetEventSchemaFromId(opts *bind.CallOpts, id string) (EventSchema, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getEventSchemaFromId", id)

	if err != nil {
		return *new(EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new(EventSchema)).(*EventSchema)

	return out0, err

}

// GetEventSchemaFromId is a free data retrieval call binding the contract method 0x49ab123d.
//
// Solidity: function getEventSchemaFromId(string id) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EtherbaseSource *EtherbaseSourceSession) GetEventSchemaFromId(id string) (EventSchema, error) {
	return _EtherbaseSource.Contract.GetEventSchemaFromId(&_EtherbaseSource.CallOpts, id)
}

// GetEventSchemaFromId is a free data retrieval call binding the contract method 0x49ab123d.
//
// Solidity: function getEventSchemaFromId(string id) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetEventSchemaFromId(id string) (EventSchema, error) {
	return _EtherbaseSource.Contract.GetEventSchemaFromId(&_EtherbaseSource.CallOpts, id)
}

// GetEventSchemaFromName is a free data retrieval call binding the contract method 0xd50d542e.
//
// Solidity: function getEventSchemaFromName(string name) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EtherbaseSource *EtherbaseSourceCaller) GetEventSchemaFromName(opts *bind.CallOpts, name string) (EventSchema, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getEventSchemaFromName", name)

	if err != nil {
		return *new(EventSchema), err
	}

	out0 := *abi.ConvertType(out[0], new(EventSchema)).(*EventSchema)

	return out0, err

}

// GetEventSchemaFromName is a free data retrieval call binding the contract method 0xd50d542e.
//
// Solidity: function getEventSchemaFromName(string name) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EtherbaseSource *EtherbaseSourceSession) GetEventSchemaFromName(name string) (EventSchema, error) {
	return _EtherbaseSource.Contract.GetEventSchemaFromName(&_EtherbaseSource.CallOpts, name)
}

// GetEventSchemaFromName is a free data retrieval call binding the contract method 0xd50d542e.
//
// Solidity: function getEventSchemaFromName(string name) view returns((string,string,(string,string,bool)[],bytes32,uint8))
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetEventSchemaFromName(name string) (EventSchema, error) {
	return _EtherbaseSource.Contract.GetEventSchemaFromName(&_EtherbaseSource.CallOpts, name)
}

// GetRegisteredEventNames is a free data retrieval call binding the contract method 0xf30826f0.
//
// Solidity: function getRegisteredEventNames() view returns(string[])
func (_EtherbaseSource *EtherbaseSourceCaller) GetRegisteredEventNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getRegisteredEventNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetRegisteredEventNames is a free data retrieval call binding the contract method 0xf30826f0.
//
// Solidity: function getRegisteredEventNames() view returns(string[])
func (_EtherbaseSource *EtherbaseSourceSession) GetRegisteredEventNames() ([]string, error) {
	return _EtherbaseSource.Contract.GetRegisteredEventNames(&_EtherbaseSource.CallOpts)
}

// GetRegisteredEventNames is a free data retrieval call binding the contract method 0xf30826f0.
//
// Solidity: function getRegisteredEventNames() view returns(string[])
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetRegisteredEventNames() ([]string, error) {
	return _EtherbaseSource.Contract.GetRegisteredEventNames(&_EtherbaseSource.CallOpts)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherbaseSource *EtherbaseSourceCaller) GetSegmentId(opts *bind.CallOpts, segment string) (*big.Int, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getSegmentId", segment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherbaseSource *EtherbaseSourceSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EtherbaseSource.Contract.GetSegmentId(&_EtherbaseSource.CallOpts, segment)
}

// GetSegmentId is a free data retrieval call binding the contract method 0x71934ef0.
//
// Solidity: function getSegmentId(string segment) view returns(uint256)
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetSegmentId(segment string) (*big.Int, error) {
	return _EtherbaseSource.Contract.GetSegmentId(&_EtherbaseSource.CallOpts, segment)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EtherbaseSource *EtherbaseSourceCaller) GetValue(opts *bind.CallOpts, segments []string) (uint8, []byte, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "getValue", segments)

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
func (_EtherbaseSource *EtherbaseSourceSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EtherbaseSource.Contract.GetValue(&_EtherbaseSource.CallOpts, segments)
}

// GetValue is a free data retrieval call binding the contract method 0x60778d95.
//
// Solidity: function getValue(string[] segments) view returns(uint8, bytes)
func (_EtherbaseSource *EtherbaseSourceCallerSession) GetValue(segments []string) (uint8, []byte, error) {
	return _EtherbaseSource.Contract.GetValue(&_EtherbaseSource.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherbaseSource *EtherbaseSourceCaller) HasEntry(opts *bind.CallOpts, segments []string) (bool, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "hasEntry", segments)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherbaseSource *EtherbaseSourceSession) HasEntry(segments []string) (bool, error) {
	return _EtherbaseSource.Contract.HasEntry(&_EtherbaseSource.CallOpts, segments)
}

// HasEntry is a free data retrieval call binding the contract method 0x64ad6c88.
//
// Solidity: function hasEntry(string[] segments) view returns(bool)
func (_EtherbaseSource *EtherbaseSourceCallerSession) HasEntry(segments []string) (bool, error) {
	return _EtherbaseSource.Contract.HasEntry(&_EtherbaseSource.CallOpts, segments)
}

// IdToName is a free data retrieval call binding the contract method 0xc76853b8.
//
// Solidity: function idToName(string ) view returns(string)
func (_EtherbaseSource *EtherbaseSourceCaller) IdToName(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "idToName", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IdToName is a free data retrieval call binding the contract method 0xc76853b8.
//
// Solidity: function idToName(string ) view returns(string)
func (_EtherbaseSource *EtherbaseSourceSession) IdToName(arg0 string) (string, error) {
	return _EtherbaseSource.Contract.IdToName(&_EtherbaseSource.CallOpts, arg0)
}

// IdToName is a free data retrieval call binding the contract method 0xc76853b8.
//
// Solidity: function idToName(string ) view returns(string)
func (_EtherbaseSource *EtherbaseSourceCallerSession) IdToName(arg0 string) (string, error) {
	return _EtherbaseSource.Contract.IdToName(&_EtherbaseSource.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherbaseSource *EtherbaseSourceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherbaseSource *EtherbaseSourceSession) Owner() (common.Address, error) {
	return _EtherbaseSource.Contract.Owner(&_EtherbaseSource.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherbaseSource *EtherbaseSourceCallerSession) Owner() (common.Address, error) {
	return _EtherbaseSource.Contract.Owner(&_EtherbaseSource.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherbaseSource *EtherbaseSourceCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherbaseSource.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherbaseSource *EtherbaseSourceSession) Validator() (common.Address, error) {
	return _EtherbaseSource.Contract.Validator(&_EtherbaseSource.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_EtherbaseSource *EtherbaseSourceCallerSession) Validator() (common.Address, error) {
	return _EtherbaseSource.Contract.Validator(&_EtherbaseSource.CallOpts)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) CreateWalletIdentity(opts *bind.TransactOpts, walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "createWalletIdentity", walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_EtherbaseSource *EtherbaseSourceSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.CreateWalletIdentity(&_EtherbaseSource.TransactOpts, walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.CreateWalletIdentity(&_EtherbaseSource.TransactOpts, walletAddress, initialRoles)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) DeleteIdentity(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "deleteIdentity", wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_EtherbaseSource *EtherbaseSourceSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.DeleteIdentity(&_EtherbaseSource.TransactOpts, wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.DeleteIdentity(&_EtherbaseSource.TransactOpts, wallet)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x0277e023.
//
// Solidity: function emitEvent(string name, bytes32[] ArgumentTopics, bytes data) returns(bool success)
func (_EtherbaseSource *EtherbaseSourceTransactor) EmitEvent(opts *bind.TransactOpts, name string, ArgumentTopics [][32]byte, data []byte) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "emitEvent", name, ArgumentTopics, data)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x0277e023.
//
// Solidity: function emitEvent(string name, bytes32[] ArgumentTopics, bytes data) returns(bool success)
func (_EtherbaseSource *EtherbaseSourceSession) EmitEvent(name string, ArgumentTopics [][32]byte, data []byte) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.EmitEvent(&_EtherbaseSource.TransactOpts, name, ArgumentTopics, data)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x0277e023.
//
// Solidity: function emitEvent(string name, bytes32[] ArgumentTopics, bytes data) returns(bool success)
func (_EtherbaseSource *EtherbaseSourceTransactorSession) EmitEvent(name string, ArgumentTopics [][32]byte, data []byte) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.EmitEvent(&_EtherbaseSource.TransactOpts, name, ArgumentTopics, data)
}

// EmitEvents is a paid mutator transaction binding the contract method 0x8c710456.
//
// Solidity: function emitEvents((string,bytes32[],bytes)[] events) returns(bool success)
func (_EtherbaseSource *EtherbaseSourceTransactor) EmitEvents(opts *bind.TransactOpts, events []EtherbaseSourceBatchEmitEvent) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "emitEvents", events)
}

// EmitEvents is a paid mutator transaction binding the contract method 0x8c710456.
//
// Solidity: function emitEvents((string,bytes32[],bytes)[] events) returns(bool success)
func (_EtherbaseSource *EtherbaseSourceSession) EmitEvents(events []EtherbaseSourceBatchEmitEvent) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.EmitEvents(&_EtherbaseSource.TransactOpts, events)
}

// EmitEvents is a paid mutator transaction binding the contract method 0x8c710456.
//
// Solidity: function emitEvents((string,bytes32[],bytes)[] events) returns(bool success)
func (_EtherbaseSource *EtherbaseSourceTransactorSession) EmitEvents(events []EtherbaseSourceBatchEmitEvent) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.EmitEvents(&_EtherbaseSource.TransactOpts, events)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherbaseSource *EtherbaseSourceTransactor) GetOrCreateSegmentId(opts *bind.TransactOpts, segment string) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "getOrCreateSegmentId", segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherbaseSource *EtherbaseSourceSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.GetOrCreateSegmentId(&_EtherbaseSource.TransactOpts, segment)
}

// GetOrCreateSegmentId is a paid mutator transaction binding the contract method 0x4151df61.
//
// Solidity: function getOrCreateSegmentId(string segment) returns(uint256)
func (_EtherbaseSource *EtherbaseSourceTransactorSession) GetOrCreateSegmentId(segment string) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.GetOrCreateSegmentId(&_EtherbaseSource.TransactOpts, segment)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) GrantRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "grantRole", wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_EtherbaseSource *EtherbaseSourceSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.GrantRole(&_EtherbaseSource.TransactOpts, wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.GrantRole(&_EtherbaseSource.TransactOpts, wallet, role)
}

// RegisterEventSchema is a paid mutator transaction binding the contract method 0xa4dbe8a0.
//
// Solidity: function registerEventSchema(string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) RegisterEventSchema(opts *bind.TransactOpts, name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "registerEventSchema", name, id, eventTopic, args)
}

// RegisterEventSchema is a paid mutator transaction binding the contract method 0xa4dbe8a0.
//
// Solidity: function registerEventSchema(string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EtherbaseSource *EtherbaseSourceSession) RegisterEventSchema(name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.RegisterEventSchema(&_EtherbaseSource.TransactOpts, name, id, eventTopic, args)
}

// RegisterEventSchema is a paid mutator transaction binding the contract method 0xa4dbe8a0.
//
// Solidity: function registerEventSchema(string name, string id, bytes32 eventTopic, (string,string,bool)[] args) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) RegisterEventSchema(name string, id string, eventTopic [32]byte, args []Argument) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.RegisterEventSchema(&_EtherbaseSource.TransactOpts, name, id, eventTopic, args)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) RemoveValue(opts *bind.TransactOpts, segments []string) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "removeValue", segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherbaseSource *EtherbaseSourceSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.RemoveValue(&_EtherbaseSource.TransactOpts, segments)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x8998fbe8.
//
// Solidity: function removeValue(string[] segments) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) RemoveValue(segments []string) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.RemoveValue(&_EtherbaseSource.TransactOpts, segments)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) RevokeRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "revokeRole", wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_EtherbaseSource *EtherbaseSourceSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.RevokeRole(&_EtherbaseSource.TransactOpts, wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.RevokeRole(&_EtherbaseSource.TransactOpts, wallet, role)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) SetValue(opts *bind.TransactOpts, segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "setValue", segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherbaseSource *EtherbaseSourceSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.SetValue(&_EtherbaseSource.TransactOpts, segments, dataType, data)
}

// SetValue is a paid mutator transaction binding the contract method 0x0761b8d0.
//
// Solidity: function setValue(string[] segments, uint8 dataType, bytes data) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) SetValue(segments []string, dataType uint8, data []byte) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.SetValue(&_EtherbaseSource.TransactOpts, segments, dataType, data)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherbaseSource *EtherbaseSourceTransactor) SetValues(opts *bind.TransactOpts, values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherbaseSource.contract.Transact(opts, "setValues", values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherbaseSource *EtherbaseSourceSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.SetValues(&_EtherbaseSource.TransactOpts, values)
}

// SetValues is a paid mutator transaction binding the contract method 0x69da55d2.
//
// Solidity: function setValues((string[],uint8,bytes)[] values) returns()
func (_EtherbaseSource *EtherbaseSourceTransactorSession) SetValues(values []EtherDatabaseLibBatchSetValue) (*types.Transaction, error) {
	return _EtherbaseSource.Contract.SetValues(&_EtherbaseSource.TransactOpts, values)
}

// EtherbaseSourceEthDBPathUpdateIterator is returned from FilterEthDBPathUpdate and is used to iterate over the raw logs and unpacked data for EthDBPathUpdate events raised by the EtherbaseSource contract.
type EtherbaseSourceEthDBPathUpdateIterator struct {
	Event *EtherbaseSourceEthDBPathUpdate // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceEthDBPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceEthDBPathUpdate)
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
		it.Event = new(EtherbaseSourceEthDBPathUpdate)
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
func (it *EtherbaseSourceEthDBPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceEthDBPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceEthDBPathUpdate represents a EthDBPathUpdate event raised by the EtherbaseSource contract.
type EtherbaseSourceEthDBPathUpdate struct {
	Path     []string
	Data     []byte
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDBPathUpdate is a free log retrieval operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterEthDBPathUpdate(opts *bind.FilterOpts) (*EtherbaseSourceEthDBPathUpdateIterator, error) {

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceEthDBPathUpdateIterator{contract: _EtherbaseSource.contract, event: "EthDBPathUpdate", logs: logs, sub: sub}, nil
}

// WatchEthDBPathUpdate is a free log subscription operation binding the contract event 0xe028433ec1ff3b8c839d1a18dc250bba72715735efde55f11debeb85fa863c5a.
//
// Solidity: event EthDBPathUpdate(string[] path, bytes data, uint8 dataType)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchEthDBPathUpdate(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceEthDBPathUpdate) (event.Subscription, error) {

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "EthDBPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceEthDBPathUpdate)
				if err := _EtherbaseSource.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseEthDBPathUpdate(log types.Log) (*EtherbaseSourceEthDBPathUpdate, error) {
	event := new(EtherbaseSourceEthDBPathUpdate)
	if err := _EtherbaseSource.contract.UnpackLog(event, "EthDBPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the EtherbaseSource contract.
type EtherbaseSourceIdentityCreatedIterator struct {
	Event *EtherbaseSourceIdentityCreated // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceIdentityCreated)
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
		it.Event = new(EtherbaseSourceIdentityCreated)
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
func (it *EtherbaseSourceIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceIdentityCreated represents a IdentityCreated event raised by the EtherbaseSource contract.
type EtherbaseSourceIdentityCreated struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterIdentityCreated(opts *bind.FilterOpts, wallet []common.Address) (*EtherbaseSourceIdentityCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceIdentityCreatedIterator{contract: _EtherbaseSource.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceIdentityCreated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceIdentityCreated)
				if err := _EtherbaseSource.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseIdentityCreated(log types.Log) (*EtherbaseSourceIdentityCreated, error) {
	event := new(EtherbaseSourceIdentityCreated)
	if err := _EtherbaseSource.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceIdentityDeletedIterator is returned from FilterIdentityDeleted and is used to iterate over the raw logs and unpacked data for IdentityDeleted events raised by the EtherbaseSource contract.
type EtherbaseSourceIdentityDeletedIterator struct {
	Event *EtherbaseSourceIdentityDeleted // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceIdentityDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceIdentityDeleted)
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
		it.Event = new(EtherbaseSourceIdentityDeleted)
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
func (it *EtherbaseSourceIdentityDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceIdentityDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceIdentityDeleted represents a IdentityDeleted event raised by the EtherbaseSource contract.
type EtherbaseSourceIdentityDeleted struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityDeleted is a free log retrieval operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterIdentityDeleted(opts *bind.FilterOpts, wallet []common.Address) (*EtherbaseSourceIdentityDeletedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "IdentityDeleted", walletRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceIdentityDeletedIterator{contract: _EtherbaseSource.contract, event: "IdentityDeleted", logs: logs, sub: sub}, nil
}

// WatchIdentityDeleted is a free log subscription operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchIdentityDeleted(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceIdentityDeleted, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "IdentityDeleted", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceIdentityDeleted)
				if err := _EtherbaseSource.contract.UnpackLog(event, "IdentityDeleted", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseIdentityDeleted(log types.Log) (*EtherbaseSourceIdentityDeleted, error) {
	event := new(EtherbaseSourceIdentityDeleted)
	if err := _EtherbaseSource.contract.UnpackLog(event, "IdentityDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceNewSegmentIterator is returned from FilterNewSegment and is used to iterate over the raw logs and unpacked data for NewSegment events raised by the EtherbaseSource contract.
type EtherbaseSourceNewSegmentIterator struct {
	Event *EtherbaseSourceNewSegment // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceNewSegmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceNewSegment)
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
		it.Event = new(EtherbaseSourceNewSegment)
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
func (it *EtherbaseSourceNewSegmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceNewSegmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceNewSegment represents a NewSegment event raised by the EtherbaseSource contract.
type EtherbaseSourceNewSegment struct {
	Segment string
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSegment is a free log retrieval operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterNewSegment(opts *bind.FilterOpts, id []*big.Int) (*EtherbaseSourceNewSegmentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceNewSegmentIterator{contract: _EtherbaseSource.contract, event: "NewSegment", logs: logs, sub: sub}, nil
}

// WatchNewSegment is a free log subscription operation binding the contract event 0x4f9f48819058c88a77c38d63c60c9ed90e3557391e8fba7523e0230a2528e778.
//
// Solidity: event NewSegment(string segment, uint256 indexed id)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchNewSegment(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceNewSegment, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "NewSegment", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceNewSegment)
				if err := _EtherbaseSource.contract.UnpackLog(event, "NewSegment", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseNewSegment(log types.Log) (*EtherbaseSourceNewSegment, error) {
	event := new(EtherbaseSourceNewSegment)
	if err := _EtherbaseSource.contract.UnpackLog(event, "NewSegment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EtherbaseSource contract.
type EtherbaseSourceRoleGrantedIterator struct {
	Event *EtherbaseSourceRoleGranted // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceRoleGranted)
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
		it.Event = new(EtherbaseSourceRoleGranted)
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
func (it *EtherbaseSourceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceRoleGranted represents a RoleGranted event raised by the EtherbaseSource contract.
type EtherbaseSourceRoleGranted struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterRoleGranted(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*EtherbaseSourceRoleGrantedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceRoleGrantedIterator{contract: _EtherbaseSource.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceRoleGranted, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceRoleGranted)
				if err := _EtherbaseSource.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseRoleGranted(log types.Log) (*EtherbaseSourceRoleGranted, error) {
	event := new(EtherbaseSourceRoleGranted)
	if err := _EtherbaseSource.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EtherbaseSource contract.
type EtherbaseSourceRoleRevokedIterator struct {
	Event *EtherbaseSourceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceRoleRevoked)
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
		it.Event = new(EtherbaseSourceRoleRevoked)
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
func (it *EtherbaseSourceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceRoleRevoked represents a RoleRevoked event raised by the EtherbaseSource contract.
type EtherbaseSourceRoleRevoked struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*EtherbaseSourceRoleRevokedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "RoleRevoked", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceRoleRevokedIterator{contract: _EtherbaseSource.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceRoleRevoked, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "RoleRevoked", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceRoleRevoked)
				if err := _EtherbaseSource.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseRoleRevoked(log types.Log) (*EtherbaseSourceRoleRevoked, error) {
	event := new(EtherbaseSourceRoleRevoked)
	if err := _EtherbaseSource.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherbaseSourceSchemaRegisteredIterator is returned from FilterSchemaRegistered and is used to iterate over the raw logs and unpacked data for SchemaRegistered events raised by the EtherbaseSource contract.
type EtherbaseSourceSchemaRegisteredIterator struct {
	Event *EtherbaseSourceSchemaRegistered // Event containing the contract specifics and raw log

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
func (it *EtherbaseSourceSchemaRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherbaseSourceSchemaRegistered)
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
		it.Event = new(EtherbaseSourceSchemaRegistered)
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
func (it *EtherbaseSourceSchemaRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherbaseSourceSchemaRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherbaseSourceSchemaRegistered represents a SchemaRegistered event raised by the EtherbaseSource contract.
type EtherbaseSourceSchemaRegistered struct {
	Name common.Hash
	Id   string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSchemaRegistered is a free log retrieval operation binding the contract event 0x4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de.
//
// Solidity: event SchemaRegistered(string indexed name, string id)
func (_EtherbaseSource *EtherbaseSourceFilterer) FilterSchemaRegistered(opts *bind.FilterOpts, name []string) (*EtherbaseSourceSchemaRegisteredIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _EtherbaseSource.contract.FilterLogs(opts, "SchemaRegistered", nameRule)
	if err != nil {
		return nil, err
	}
	return &EtherbaseSourceSchemaRegisteredIterator{contract: _EtherbaseSource.contract, event: "SchemaRegistered", logs: logs, sub: sub}, nil
}

// WatchSchemaRegistered is a free log subscription operation binding the contract event 0x4574cea21b42336b945bdff0cad6c2fe8d1c2e0006ea0090a19d7459404914de.
//
// Solidity: event SchemaRegistered(string indexed name, string id)
func (_EtherbaseSource *EtherbaseSourceFilterer) WatchSchemaRegistered(opts *bind.WatchOpts, sink chan<- *EtherbaseSourceSchemaRegistered, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _EtherbaseSource.contract.WatchLogs(opts, "SchemaRegistered", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherbaseSourceSchemaRegistered)
				if err := _EtherbaseSource.contract.UnpackLog(event, "SchemaRegistered", log); err != nil {
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
func (_EtherbaseSource *EtherbaseSourceFilterer) ParseSchemaRegistered(log types.Log) (*EtherbaseSourceSchemaRegistered, error) {
	event := new(EtherbaseSourceSchemaRegistered)
	if err := _EtherbaseSource.contract.UnpackLog(event, "SchemaRegistered", log); err != nil {
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
	Bin: "0x6080604052348015600e575f5ffd5b50610c6a8061001c5f395ff3fe6080604052600436106100ef575f3560e01c80634d2301cc11610087578063a8b0574e11610057578063a8b0574e14610221578063bce38bd71461023b578063c3077fa91461024e578063ee82ac5e14610261575f5ffd5b80634d2301cc146101c357806372425d9d146101ea57806382ad56cb146101fc57806386d516e81461020f575f5ffd5b80633408e470116100c25780633408e4701461016b578063399542e91461017d5780633e64a6961461019f57806342cbb15c146101b1575f5ffd5b80630f28c97d146100f3578063174dea7114610114578063252dba421461013457806327e86d6e14610155575b5f5ffd5b3480156100fe575f5ffd5b50425b6040519081526020015b60405180910390f35b610127610122366004610958565b61027f565b60405161010b9190610a38565b610147610142366004610958565b610464565b60405161010b929190610a51565b348015610160575f5ffd5b50435f190140610101565b348015610176575f5ffd5b5046610101565b61019061018b366004610abb565b6105d2565b60405161010b93929190610b10565b3480156101aa575f5ffd5b5048610101565b3480156101bc575f5ffd5b5043610101565b3480156101ce575f5ffd5b506101016101dd366004610b37565b6001600160a01b03163190565b3480156101f5575f5ffd5b5044610101565b61012761020a366004610958565b6105ed565b34801561021a575f5ffd5b5045610101565b34801561022c575f5ffd5b5060405141815260200161010b565b610127610249366004610abb565b610766565b61019061025c366004610958565b6108f2565b34801561026c575f5ffd5b5061010161027b366004610b5d565b4090565b60605f828067ffffffffffffffff81111561029c5761029c610b74565b6040519080825280602002602001820160405280156102e157816020015b604080518082019091525f8152606060208201528152602001906001900390816102ba5790505b509250365f5b82811015610406575f85828151811061030257610302610b88565b6020026020010151905087878381811061031e5761031e610b88565b90506020028101906103309190610b9c565b6040810135958601959093506103496020850185610b37565b6001600160a01b0316816103606060870187610bba565b60405161036e929190610bfd565b5f6040518083038185875af1925050503d805f81146103a8576040519150601f19603f3d011682016040523d82523d5f602084013e6103ad565b606091505b5060208085019190915290151580845290850135176103fc5762461bcd60e51b5f526020600452601760245276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b60445260845ffd5b50506001016102e7565b5082341461045b5760405162461bcd60e51b815260206004820152601a60248201527f4d756c746963616c6c333a2076616c7565206d69736d6174636800000000000060448201526064015b60405180910390fd5b50505092915050565b436060828067ffffffffffffffff81111561048157610481610b74565b6040519080825280602002602001820160405280156104b457816020015b606081526020019060019003908161049f5790505b509150365f5b828110156105c8575f8787838181106104d5576104d5610b88565b90506020028101906104e79190610c0c565b92506104f66020840184610b37565b6001600160a01b031661050c6020850185610bba565b60405161051a929190610bfd565b5f604051808303815f865af19150503d805f8114610553576040519150601f19603f3d011682016040523d82523d5f602084013e610558565b606091505b5086848151811061056b5761056b610b88565b60209081029190910101529050806105bf5760405162461bcd60e51b8152602060048201526017602482015276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b6044820152606401610452565b506001016104ba565b5050509250929050565b43804060606105e2868686610766565b905093509350939050565b6060818067ffffffffffffffff81111561060957610609610b74565b60405190808252806020026020018201604052801561064e57816020015b604080518082019091525f8152606060208201528152602001906001900390816106275790505b509150365f5b8281101561045b575f84828151811061066f5761066f610b88565b6020026020010151905086868381811061068b5761068b610b88565b905060200281019061069d9190610c20565b92506106ac6020840184610b37565b6001600160a01b03166106c26040850185610bba565b6040516106d0929190610bfd565b5f604051808303815f865af19150503d805f8114610709576040519150601f19603f3d011682016040523d82523d5f602084013e61070e565b606091505b50602080840191909152901515808352908401351761075d5762461bcd60e51b5f526020600452601760245276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b60445260645ffd5b50600101610654565b6060818067ffffffffffffffff81111561078257610782610b74565b6040519080825280602002602001820160405280156107c757816020015b604080518082019091525f8152606060208201528152602001906001900390816107a05790505b509150365f5b828110156108e8575f8482815181106107e8576107e8610b88565b6020026020010151905086868381811061080457610804610b88565b90506020028101906108169190610c0c565b92506108256020840184610b37565b6001600160a01b031661083b6020850185610bba565b604051610849929190610bfd565b5f604051808303815f865af19150503d805f8114610882576040519150601f19603f3d011682016040523d82523d5f602084013e610887565b606091505b5060208301521515815287156108df5780516108df5760405162461bcd60e51b8152602060048201526017602482015276135d5b1d1a58d85b1b0cce8818d85b1b0819985a5b1959604a1b6044820152606401610452565b506001016107cd565b5050509392505050565b5f5f6060610902600186866105d2565b919790965090945092505050565b5f5f83601f840112610920575f5ffd5b50813567ffffffffffffffff811115610937575f5ffd5b6020830191508360208260051b8501011115610951575f5ffd5b9250929050565b5f5f60208385031215610969575f5ffd5b823567ffffffffffffffff81111561097f575f5ffd5b61098b85828601610910565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f82825180855260208501945060208160051b830101602085015f5b83811015610a2c57601f1985840301885281518051151584526020810151905060406020850152610a156040850182610997565b6020998a01999094509290920191506001016109e1565b50909695505050505050565b602081525f610a4a60208301846109c5565b9392505050565b5f604082018483526040602084015280845180835260608501915060608160051b8601019250602086015f5b82811015610aae57605f19878603018452610a99858351610997565b94506020938401939190910190600101610a7d565b5092979650505050505050565b5f5f5f60408486031215610acd575f5ffd5b83358015158114610adc575f5ffd5b9250602084013567ffffffffffffffff811115610af7575f5ffd5b610b0386828701610910565b9497909650939450505050565b838152826020820152606060408201525f610b2e60608301846109c5565b95945050505050565b5f60208284031215610b47575f5ffd5b81356001600160a01b0381168114610a4a575f5ffd5b5f60208284031215610b6d575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235607e19833603018112610bb0575f5ffd5b9190910192915050565b5f5f8335601e19843603018112610bcf575f5ffd5b83018035915067ffffffffffffffff821115610be9575f5ffd5b602001915036819003821315610951575f5ffd5b818382375f9101908152919050565b5f8235603e19833603018112610bb0575f5ffd5b5f8235605e19833603018112610bb0575f5ffdfea26469706673582212203aff4c47df40ea02be979aa87ae8576f97a70a788f2ec4c97287e6ce8eba568064736f6c634300081c0033",
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

// RoleControlMetaData contains all meta data concerning the RoleControl contract.
var RoleControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IdentityAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIdentity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"initialRoles\",\"type\":\"uint8[]\"}],\"name\":\"createWalletIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"deleteIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllIdentities\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structRoleControl.IdentityView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RoleControlABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleControlMetaData.ABI instead.
var RoleControlABI = RoleControlMetaData.ABI

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

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_RoleControl *RoleControlCaller) GetAllIdentities(opts *bind.CallOpts) ([]RoleControlIdentityView, error) {
	var out []interface{}
	err := _RoleControl.contract.Call(opts, &out, "getAllIdentities")

	if err != nil {
		return *new([]RoleControlIdentityView), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleControlIdentityView)).(*[]RoleControlIdentityView)

	return out0, err

}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_RoleControl *RoleControlSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _RoleControl.Contract.GetAllIdentities(&_RoleControl.CallOpts)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_RoleControl *RoleControlCallerSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _RoleControl.Contract.GetAllIdentities(&_RoleControl.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_RoleControl *RoleControlCaller) GetAllWallets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoleControl.contract.Call(opts, &out, "getAllWallets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_RoleControl *RoleControlSession) GetAllWallets() ([]common.Address, error) {
	return _RoleControl.Contract.GetAllWallets(&_RoleControl.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_RoleControl *RoleControlCallerSession) GetAllWallets() ([]common.Address, error) {
	return _RoleControl.Contract.GetAllWallets(&_RoleControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoleControl *RoleControlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoleControl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoleControl *RoleControlSession) Owner() (common.Address, error) {
	return _RoleControl.Contract.Owner(&_RoleControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoleControl *RoleControlCallerSession) Owner() (common.Address, error) {
	return _RoleControl.Contract.Owner(&_RoleControl.CallOpts)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_RoleControl *RoleControlTransactor) CreateWalletIdentity(opts *bind.TransactOpts, walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _RoleControl.contract.Transact(opts, "createWalletIdentity", walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_RoleControl *RoleControlSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _RoleControl.Contract.CreateWalletIdentity(&_RoleControl.TransactOpts, walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_RoleControl *RoleControlTransactorSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _RoleControl.Contract.CreateWalletIdentity(&_RoleControl.TransactOpts, walletAddress, initialRoles)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_RoleControl *RoleControlTransactor) DeleteIdentity(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _RoleControl.contract.Transact(opts, "deleteIdentity", wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_RoleControl *RoleControlSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _RoleControl.Contract.DeleteIdentity(&_RoleControl.TransactOpts, wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_RoleControl *RoleControlTransactorSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _RoleControl.Contract.DeleteIdentity(&_RoleControl.TransactOpts, wallet)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_RoleControl *RoleControlTransactor) GrantRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControl.contract.Transact(opts, "grantRole", wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_RoleControl *RoleControlSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControl.Contract.GrantRole(&_RoleControl.TransactOpts, wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_RoleControl *RoleControlTransactorSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControl.Contract.GrantRole(&_RoleControl.TransactOpts, wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_RoleControl *RoleControlTransactor) RevokeRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControl.contract.Transact(opts, "revokeRole", wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_RoleControl *RoleControlSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControl.Contract.RevokeRole(&_RoleControl.TransactOpts, wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_RoleControl *RoleControlTransactorSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControl.Contract.RevokeRole(&_RoleControl.TransactOpts, wallet, role)
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

// RoleControlImplMetaData contains all meta data concerning the RoleControlImpl contract.
var RoleControlImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IdentityAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdentityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIdentity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"IdentityDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"initialRoles\",\"type\":\"uint8[]\"}],\"name\":\"createWalletIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"deleteIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllIdentities\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structRoleControl.IdentityView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"enumRoleControl.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516111a73803806111a783398101604081905261002e91610387565b6100378161003d565b506103dc565b600280546001600160a01b0319166001600160a01b03831617905560408051600480825260a082019092525f91602082016080803683370190505090506002815f8151811061008e5761008e6103b4565b602002602001019060038111156100a7576100a76103c8565b908160038111156100ba576100ba6103c8565b815250505f816001815181106100d2576100d26103b4565b602002602001019060038111156100eb576100eb6103c8565b908160038111156100fe576100fe6103c8565b81525050600181600281518110610117576101176103b4565b60200260200101906003811115610130576101306103c8565b90816003811115610143576101436103c8565b8152505060038160038151811061015c5761015c6103b4565b60200260200101906003811115610175576101756103c8565b90816003811115610188576101886103c8565b9052506101958282610199565b5050565b6001600160a01b0382165f9081526020819052604090205460ff16156101d2576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602081905260408120805460ff19166001908117909155805480820182559082527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169092179091555b81518110156102705761026883838381518110610255576102556103b4565b60200260200101516102a860201b60201c565b600101610236565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a25050565b6001600160a01b0382165f9081526020819052604090205460ff166102e05760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208181526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a918202191690836003811115610338576103386103c8565b021790555081600381111561034f5761034f6103c8565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a3505050565b5f60208284031215610397575f5ffd5b81516001600160a01b03811681146103ad575f5ffd5b9392505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b610dbe806103e95f395ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c80637eeede35116100585780637eeede35146100c45780638da5cb5b146100d9578063a8d29d1d14610104578063dced478314610117575f5ffd5b8063077d3c031461007e5780631bfa8601146100935780633e840236146100b1575b5f5ffd5b61009161008c366004610ad6565b61012a565b005b61009b610191565b6040516100a89190610b07565b60405180910390f35b6100916100bf366004610ad6565b6101f1565b6100cc610253565b6040516100a89190610b66565b6002546100ec906001600160a01b031681565b6040516001600160a01b0390911681526020016100a8565b610091610112366004610c2f565b6103bc565b610091610125366004610c63565b610421565b335f9081526020819052604090205460029060ff1661015c5760405163c597feeb60e01b815260040160405180910390fd5b6101663382610483565b610182576040516282b42960e81b815260040160405180910390fd5b61018c838361051f565b505050565b606060018054806020026020016040519081016040528092919081815260200182805480156101e757602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116101c9575b5050505050905090565b335f9081526020819052604090205460029060ff166102235760405163c597feeb60e01b815260040160405180910390fd5b61022d3382610483565b610249576040516282b42960e81b815260040160405180910390fd5b61018c83836106ef565b6001546060905f9067ffffffffffffffff81111561027357610273610c4f565b6040519080825280602002602001820160405280156102b857816020015b604080518082019091525f8152606060208201528152602001906001900390816102915790505b5090505f5b6001548110156103b6575f600182815481106102db576102db610d41565b5f9182526020808320909101546001600160a01b03168083528282526040928390208351808501855282815260018201805486518187028101870190975280875293965091949093848101939192919083018282801561038757602002820191905f5260205f20905f905b82829054906101000a900460ff16600381111561036557610365610b52565b8152602060019283018181049485019490930390920291018084116103465790505b50505050508152508484815181106103a1576103a1610d41565b602090810291909101015250506001016102bd565b50919050565b335f9081526020819052604090205460029060ff166103ee5760405163c597feeb60e01b815260040160405180910390fd5b6103f83382610483565b610414576040516282b42960e81b815260040160405180910390fd5b61041d826107ce565b5050565b335f9081526020819052604090205460029060ff166104535760405163c597feeb60e01b815260040160405180910390fd5b61045d3382610483565b610479576040516282b42960e81b815260040160405180910390fd5b61018c8383610967565b6001600160a01b0382165f908152602081905260408120600101815b8154811015610513578360038111156104ba576104ba610b52565b8282815481106104cc576104cc610d41565b905f5260205f2090602091828204019190069054906101000a900460ff1660038111156104fb576104fb610b52565b0361050b57600192505050610519565b60010161049f565b505f9150505b92915050565b6001600160a01b0382165f9081526020819052604090205460ff166105575760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f9081526020819052604081209060018201905b81548110156106e85783600381111561059057610590610b52565b8282815481106105a2576105a2610d41565b905f5260205f2090602091828204019190069054906101000a900460ff1660038111156105d1576105d1610b52565b036106e057815482906105e690600190610d55565b815481106105f6576105f6610d41565b905f5260205f2090602091828204019190069054906101000a900460ff1682828154811061062657610626610d41565b905f5260205f2090602091828204019190066101000a81548160ff0219169083600381111561065757610657610b52565b02179055508180548061066c5761066c610d74565b600190038181905f5260205f2090602091828204019190066101000a81549060ff021916905590558360038111156106a6576106a6610b52565b6040516001600160a01b038716907f34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769905f90a35050505050565b600101610575565b5050505050565b6001600160a01b0382165f9081526020819052604090205460ff166107275760405163c597feeb60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260208181526040822060018082018054918201815584529282902091830490910180549192849260ff601f9092166101000a91820219169083600381111561077f5761077f610b52565b021790555081600381111561079657610796610b52565b6040516001600160a01b038516907faa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121905f90a3505050565b6001600160a01b0381165f908152602081905260408120600181015490910361080a5760405163178423dd60e01b815260040160405180910390fd5b6001600160a01b0382165f908152602081905260408120805460ff19168155906108376001830182610a70565b505f90505b60015481101561092f57826001600160a01b03166001828154811061086357610863610d41565b5f918252602090912001546001600160a01b031603610927576001805461088b908290610d55565b8154811061089b5761089b610d41565b5f91825260209091200154600180546001600160a01b0390921691839081106108c6576108c6610d41565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550600180548061090257610902610d74565b5f8281526020902081015f1990810180546001600160a01b031916905501905561092f565b60010161083c565b506040516001600160a01b038316907fc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8905f90a25050565b6001600160a01b0382165f9081526020819052604090205460ff16156109a0576040516335a081ab60e11b815260040160405180910390fd5b6001600160a01b0382165f818152602081905260408120805460ff19166001908117909155805480820182559082527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169092179091555b8151811015610a3857610a3083838381518110610a2357610a23610d41565b60200260200101516106ef565b600101610a04565b506040516001600160a01b038316907fac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0905f90a25050565b5080545f8255601f0160209004905f5260205f2090810190610a929190610a95565b50565b5b80821115610aa9575f8155600101610a96565b5090565b80356001600160a01b0381168114610ac3575f5ffd5b919050565b803560048110610ac3575f5ffd5b5f5f60408385031215610ae7575f5ffd5b610af083610aad565b9150610afe60208401610ac8565b90509250929050565b602080825282518282018190525f918401906040840190835b81811015610b475783516001600160a01b0316835260209384019390920191600101610b20565b509095945050505050565b634e487b7160e01b5f52602160045260245ffd5b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015610c2357868503603f19018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101905f9060608801905b80831015610c0b57835160048110610bf357634e487b7160e01b5f52602160045260245ffd5b82526020938401936001939093019290910190610bcd565b50965050506020938401939190910190600101610b8c565b50929695505050505050565b5f60208284031215610c3f575f5ffd5b610c4882610aad565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610c74575f5ffd5b610c7d83610aad565b9150602083013567ffffffffffffffff811115610c98575f5ffd5b8301601f81018513610ca8575f5ffd5b803567ffffffffffffffff811115610cc257610cc2610c4f565b8060051b604051601f19603f830116810181811067ffffffffffffffff82111715610cef57610cef610c4f565b604052918252602081840181019290810188841115610d0c575f5ffd5b6020850194505b83851015610d3257610d2485610ac8565b815260209485019401610d13565b50809450505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b8181038181111561051957634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220618656ff9ee477b47992f0b0115244a72220c1ecc9d9dd09a78541131b988fe864736f6c634300081c0033",
}

// RoleControlImplABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleControlImplMetaData.ABI instead.
var RoleControlImplABI = RoleControlImplMetaData.ABI

// RoleControlImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoleControlImplMetaData.Bin instead.
var RoleControlImplBin = RoleControlImplMetaData.Bin

// DeployRoleControlImpl deploys a new Ethereum contract, binding an instance of RoleControlImpl to it.
func DeployRoleControlImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *RoleControlImpl, error) {
	parsed, err := RoleControlImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoleControlImplBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoleControlImpl{RoleControlImplCaller: RoleControlImplCaller{contract: contract}, RoleControlImplTransactor: RoleControlImplTransactor{contract: contract}, RoleControlImplFilterer: RoleControlImplFilterer{contract: contract}}, nil
}

// RoleControlImpl is an auto generated Go binding around an Ethereum contract.
type RoleControlImpl struct {
	RoleControlImplCaller     // Read-only binding to the contract
	RoleControlImplTransactor // Write-only binding to the contract
	RoleControlImplFilterer   // Log filterer for contract events
}

// RoleControlImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleControlImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleControlImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleControlImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleControlImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleControlImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleControlImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleControlImplSession struct {
	Contract     *RoleControlImpl  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleControlImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleControlImplCallerSession struct {
	Contract *RoleControlImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RoleControlImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleControlImplTransactorSession struct {
	Contract     *RoleControlImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RoleControlImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleControlImplRaw struct {
	Contract *RoleControlImpl // Generic contract binding to access the raw methods on
}

// RoleControlImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleControlImplCallerRaw struct {
	Contract *RoleControlImplCaller // Generic read-only contract binding to access the raw methods on
}

// RoleControlImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleControlImplTransactorRaw struct {
	Contract *RoleControlImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleControlImpl creates a new instance of RoleControlImpl, bound to a specific deployed contract.
func NewRoleControlImpl(address common.Address, backend bind.ContractBackend) (*RoleControlImpl, error) {
	contract, err := bindRoleControlImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleControlImpl{RoleControlImplCaller: RoleControlImplCaller{contract: contract}, RoleControlImplTransactor: RoleControlImplTransactor{contract: contract}, RoleControlImplFilterer: RoleControlImplFilterer{contract: contract}}, nil
}

// NewRoleControlImplCaller creates a new read-only instance of RoleControlImpl, bound to a specific deployed contract.
func NewRoleControlImplCaller(address common.Address, caller bind.ContractCaller) (*RoleControlImplCaller, error) {
	contract, err := bindRoleControlImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplCaller{contract: contract}, nil
}

// NewRoleControlImplTransactor creates a new write-only instance of RoleControlImpl, bound to a specific deployed contract.
func NewRoleControlImplTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleControlImplTransactor, error) {
	contract, err := bindRoleControlImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplTransactor{contract: contract}, nil
}

// NewRoleControlImplFilterer creates a new log filterer instance of RoleControlImpl, bound to a specific deployed contract.
func NewRoleControlImplFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleControlImplFilterer, error) {
	contract, err := bindRoleControlImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplFilterer{contract: contract}, nil
}

// bindRoleControlImpl binds a generic wrapper to an already deployed contract.
func bindRoleControlImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoleControlImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleControlImpl *RoleControlImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleControlImpl.Contract.RoleControlImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleControlImpl *RoleControlImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.RoleControlImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleControlImpl *RoleControlImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.RoleControlImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleControlImpl *RoleControlImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleControlImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleControlImpl *RoleControlImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleControlImpl *RoleControlImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.contract.Transact(opts, method, params...)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_RoleControlImpl *RoleControlImplCaller) GetAllIdentities(opts *bind.CallOpts) ([]RoleControlIdentityView, error) {
	var out []interface{}
	err := _RoleControlImpl.contract.Call(opts, &out, "getAllIdentities")

	if err != nil {
		return *new([]RoleControlIdentityView), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleControlIdentityView)).(*[]RoleControlIdentityView)

	return out0, err

}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_RoleControlImpl *RoleControlImplSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _RoleControlImpl.Contract.GetAllIdentities(&_RoleControlImpl.CallOpts)
}

// GetAllIdentities is a free data retrieval call binding the contract method 0x7eeede35.
//
// Solidity: function getAllIdentities() view returns((address,uint8[])[])
func (_RoleControlImpl *RoleControlImplCallerSession) GetAllIdentities() ([]RoleControlIdentityView, error) {
	return _RoleControlImpl.Contract.GetAllIdentities(&_RoleControlImpl.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_RoleControlImpl *RoleControlImplCaller) GetAllWallets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoleControlImpl.contract.Call(opts, &out, "getAllWallets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_RoleControlImpl *RoleControlImplSession) GetAllWallets() ([]common.Address, error) {
	return _RoleControlImpl.Contract.GetAllWallets(&_RoleControlImpl.CallOpts)
}

// GetAllWallets is a free data retrieval call binding the contract method 0x1bfa8601.
//
// Solidity: function getAllWallets() view returns(address[])
func (_RoleControlImpl *RoleControlImplCallerSession) GetAllWallets() ([]common.Address, error) {
	return _RoleControlImpl.Contract.GetAllWallets(&_RoleControlImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoleControlImpl *RoleControlImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoleControlImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoleControlImpl *RoleControlImplSession) Owner() (common.Address, error) {
	return _RoleControlImpl.Contract.Owner(&_RoleControlImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoleControlImpl *RoleControlImplCallerSession) Owner() (common.Address, error) {
	return _RoleControlImpl.Contract.Owner(&_RoleControlImpl.CallOpts)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_RoleControlImpl *RoleControlImplTransactor) CreateWalletIdentity(opts *bind.TransactOpts, walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _RoleControlImpl.contract.Transact(opts, "createWalletIdentity", walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_RoleControlImpl *RoleControlImplSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.CreateWalletIdentity(&_RoleControlImpl.TransactOpts, walletAddress, initialRoles)
}

// CreateWalletIdentity is a paid mutator transaction binding the contract method 0xdced4783.
//
// Solidity: function createWalletIdentity(address walletAddress, uint8[] initialRoles) returns()
func (_RoleControlImpl *RoleControlImplTransactorSession) CreateWalletIdentity(walletAddress common.Address, initialRoles []uint8) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.CreateWalletIdentity(&_RoleControlImpl.TransactOpts, walletAddress, initialRoles)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_RoleControlImpl *RoleControlImplTransactor) DeleteIdentity(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _RoleControlImpl.contract.Transact(opts, "deleteIdentity", wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_RoleControlImpl *RoleControlImplSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.DeleteIdentity(&_RoleControlImpl.TransactOpts, wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_RoleControlImpl *RoleControlImplTransactorSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.DeleteIdentity(&_RoleControlImpl.TransactOpts, wallet)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_RoleControlImpl *RoleControlImplTransactor) GrantRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControlImpl.contract.Transact(opts, "grantRole", wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_RoleControlImpl *RoleControlImplSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.GrantRole(&_RoleControlImpl.TransactOpts, wallet, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address wallet, uint8 role) returns()
func (_RoleControlImpl *RoleControlImplTransactorSession) GrantRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.GrantRole(&_RoleControlImpl.TransactOpts, wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_RoleControlImpl *RoleControlImplTransactor) RevokeRole(opts *bind.TransactOpts, wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControlImpl.contract.Transact(opts, "revokeRole", wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_RoleControlImpl *RoleControlImplSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.RevokeRole(&_RoleControlImpl.TransactOpts, wallet, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address wallet, uint8 role) returns()
func (_RoleControlImpl *RoleControlImplTransactorSession) RevokeRole(wallet common.Address, role uint8) (*types.Transaction, error) {
	return _RoleControlImpl.Contract.RevokeRole(&_RoleControlImpl.TransactOpts, wallet, role)
}

// RoleControlImplIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the RoleControlImpl contract.
type RoleControlImplIdentityCreatedIterator struct {
	Event *RoleControlImplIdentityCreated // Event containing the contract specifics and raw log

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
func (it *RoleControlImplIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlImplIdentityCreated)
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
		it.Event = new(RoleControlImplIdentityCreated)
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
func (it *RoleControlImplIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlImplIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlImplIdentityCreated represents a IdentityCreated event raised by the RoleControlImpl contract.
type RoleControlImplIdentityCreated struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_RoleControlImpl *RoleControlImplFilterer) FilterIdentityCreated(opts *bind.FilterOpts, wallet []common.Address) (*RoleControlImplIdentityCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControlImpl.contract.FilterLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplIdentityCreatedIterator{contract: _RoleControlImpl.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address indexed wallet)
func (_RoleControlImpl *RoleControlImplFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *RoleControlImplIdentityCreated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControlImpl.contract.WatchLogs(opts, "IdentityCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlImplIdentityCreated)
				if err := _RoleControlImpl.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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
func (_RoleControlImpl *RoleControlImplFilterer) ParseIdentityCreated(log types.Log) (*RoleControlImplIdentityCreated, error) {
	event := new(RoleControlImplIdentityCreated)
	if err := _RoleControlImpl.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlImplIdentityDeletedIterator is returned from FilterIdentityDeleted and is used to iterate over the raw logs and unpacked data for IdentityDeleted events raised by the RoleControlImpl contract.
type RoleControlImplIdentityDeletedIterator struct {
	Event *RoleControlImplIdentityDeleted // Event containing the contract specifics and raw log

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
func (it *RoleControlImplIdentityDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlImplIdentityDeleted)
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
		it.Event = new(RoleControlImplIdentityDeleted)
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
func (it *RoleControlImplIdentityDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlImplIdentityDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlImplIdentityDeleted represents a IdentityDeleted event raised by the RoleControlImpl contract.
type RoleControlImplIdentityDeleted struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityDeleted is a free log retrieval operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_RoleControlImpl *RoleControlImplFilterer) FilterIdentityDeleted(opts *bind.FilterOpts, wallet []common.Address) (*RoleControlImplIdentityDeletedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControlImpl.contract.FilterLogs(opts, "IdentityDeleted", walletRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplIdentityDeletedIterator{contract: _RoleControlImpl.contract, event: "IdentityDeleted", logs: logs, sub: sub}, nil
}

// WatchIdentityDeleted is a free log subscription operation binding the contract event 0xc24bb9f4ebb7a8b8d37e375af6ac3f7e39d0218d2042bbd057d783048a904cd8.
//
// Solidity: event IdentityDeleted(address indexed wallet)
func (_RoleControlImpl *RoleControlImplFilterer) WatchIdentityDeleted(opts *bind.WatchOpts, sink chan<- *RoleControlImplIdentityDeleted, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _RoleControlImpl.contract.WatchLogs(opts, "IdentityDeleted", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlImplIdentityDeleted)
				if err := _RoleControlImpl.contract.UnpackLog(event, "IdentityDeleted", log); err != nil {
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
func (_RoleControlImpl *RoleControlImplFilterer) ParseIdentityDeleted(log types.Log) (*RoleControlImplIdentityDeleted, error) {
	event := new(RoleControlImplIdentityDeleted)
	if err := _RoleControlImpl.contract.UnpackLog(event, "IdentityDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlImplRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RoleControlImpl contract.
type RoleControlImplRoleGrantedIterator struct {
	Event *RoleControlImplRoleGranted // Event containing the contract specifics and raw log

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
func (it *RoleControlImplRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlImplRoleGranted)
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
		it.Event = new(RoleControlImplRoleGranted)
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
func (it *RoleControlImplRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlImplRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlImplRoleGranted represents a RoleGranted event raised by the RoleControlImpl contract.
type RoleControlImplRoleGranted struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_RoleControlImpl *RoleControlImplFilterer) FilterRoleGranted(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*RoleControlImplRoleGrantedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControlImpl.contract.FilterLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplRoleGrantedIterator{contract: _RoleControlImpl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed wallet, uint8 indexed role)
func (_RoleControlImpl *RoleControlImplFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RoleControlImplRoleGranted, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControlImpl.contract.WatchLogs(opts, "RoleGranted", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlImplRoleGranted)
				if err := _RoleControlImpl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RoleControlImpl *RoleControlImplFilterer) ParseRoleGranted(log types.Log) (*RoleControlImplRoleGranted, error) {
	event := new(RoleControlImplRoleGranted)
	if err := _RoleControlImpl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleControlImplRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RoleControlImpl contract.
type RoleControlImplRoleRevokedIterator struct {
	Event *RoleControlImplRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoleControlImplRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleControlImplRoleRevoked)
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
		it.Event = new(RoleControlImplRoleRevoked)
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
func (it *RoleControlImplRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleControlImplRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleControlImplRoleRevoked represents a RoleRevoked event raised by the RoleControlImpl contract.
type RoleControlImplRoleRevoked struct {
	Wallet common.Address
	Role   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_RoleControlImpl *RoleControlImplFilterer) FilterRoleRevoked(opts *bind.FilterOpts, wallet []common.Address, role []uint8) (*RoleControlImplRoleRevokedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControlImpl.contract.FilterLogs(opts, "RoleRevoked", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &RoleControlImplRoleRevokedIterator{contract: _RoleControlImpl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed wallet, uint8 indexed role)
func (_RoleControlImpl *RoleControlImplFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoleControlImplRoleRevoked, wallet []common.Address, role []uint8) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _RoleControlImpl.contract.WatchLogs(opts, "RoleRevoked", walletRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleControlImplRoleRevoked)
				if err := _RoleControlImpl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RoleControlImpl *RoleControlImplFilterer) ParseRoleRevoked(log types.Log) (*RoleControlImplRoleRevoked, error) {
	event := new(RoleControlImplRoleRevoked)
	if err := _RoleControlImpl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"segments\",\"type\":\"string[]\"}],\"name\":\"encodePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f553480156012575f5ffd5b506105a5806100205f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c80632fc377871461002d575b5f5ffd5b61004061003b36600461028f565b610056565b60405161004d9190610300565b60405180910390f35b6060610062838361006b565b90505b92915050565b604080515f808252602082019092526060915b838110156100eb575f6100b386868481811061009c5761009c610349565b90506020028101906100ae919061035d565b610117565b9050826100bf826101f1565b6040516020016100d09291906103be565b60408051601f1981840301815291905292505060010161007e565b506040516100ff9082905f906020016103da565b60408051808303601f19018152919052949350505050565b5f5f6001848460405161012b9291906103fe565b9081526040805160209281900383019020545f818152600290935291205490915060ff161561015b579050610065565b5f80549081908061016b8361040d565b919050555080600186866040516101839291906103fe565b90815260405190819003602001902055600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b016101cf8587836104b5565b505f818152600260205260409020805460ff1916600117905591506100659050565b6060815f0361021c576040515f60208201526021016040516020818303038152906040529050919050565b6060825b6080811061025e578181607f1660801760f81b6040516020016102449291906103da565b60408051601f19818403018152919052915060071c610220565b8181607f1660f81b6040516020016102779291906103da565b60408051601f19818403018152919052949350505050565b5f5f602083850312156102a0575f5ffd5b823567ffffffffffffffff8111156102b6575f5ffd5b8301601f810185136102c6575f5ffd5b803567ffffffffffffffff8111156102dc575f5ffd5b8560208260051b84010111156102f0575f5ffd5b6020919091019590945092505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112610372575f5ffd5b83018035915067ffffffffffffffff82111561038c575f5ffd5b6020019150368190038213156103a0575f5ffd5b9250929050565b5f81518060208401855e5f93019283525090919050565b5f6103d26103cc83866103a7565b846103a7565b949350505050565b5f6103e582856103a7565b6001600160f81b03199390931683525050600101919050565b818382375f9101908152919050565b5f6001820161042a57634e487b7160e01b5f52601160045260245ffd5b5060010190565b600181811c9082168061044557607f821691505b60208210810361046357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156104b057805f5260205f20601f840160051c8101602085101561048e5750805b601f840160051c820191505b818110156104ad575f815560010161049a565b50505b505050565b67ffffffffffffffff8311156104cd576104cd610335565b6104e1836104db8354610431565b83610469565b5f601f841160018114610512575f85156104fb5750838201355b5f19600387901b1c1916600186901b1783556104ad565b5f83815260208120601f198716915b828110156105415786850135825560209485019460019092019101610521565b508682101561055d575f1960f88860031b161c19848701351681555b505060018560011b018355505050505056fea26469706673582212208e30f1dabd966a2ed0405feb303a8c05b094df37217eb9731475684baeaac2d864736f6c634300081c0033",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// TestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMetaData.Bin instead.
var TestBin = TestMetaData.Bin

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// EncodePath is a paid mutator transaction binding the contract method 0x2fc37787.
//
// Solidity: function encodePath(string[] segments) returns(bytes)
func (_Test *TestTransactor) EncodePath(opts *bind.TransactOpts, segments []string) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "encodePath", segments)
}

// EncodePath is a paid mutator transaction binding the contract method 0x2fc37787.
//
// Solidity: function encodePath(string[] segments) returns(bytes)
func (_Test *TestSession) EncodePath(segments []string) (*types.Transaction, error) {
	return _Test.Contract.EncodePath(&_Test.TransactOpts, segments)
}

// EncodePath is a paid mutator transaction binding the contract method 0x2fc37787.
//
// Solidity: function encodePath(string[] segments) returns(bytes)
func (_Test *TestTransactorSession) EncodePath(segments []string) (*types.Transaction, error) {
	return _Test.Contract.EncodePath(&_Test.TransactOpts, segments)
}
