// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// MetaMorphoFactoryMetaData contains all meta data concerning the MetaMorphoFactory contract.
var MetaMorphoFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"morpho\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"metaMorpho\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"CreateMetaMorpho\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MORPHO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"createMetaMorpho\",\"outputs\":[{\"internalType\":\"contractIMetaMorpho\",\"name\":\"metaMorpho\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMetaMorpho\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MetaMorphoFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaMorphoFactoryMetaData.ABI instead.
var MetaMorphoFactoryABI = MetaMorphoFactoryMetaData.ABI

// MetaMorphoFactory is an auto generated Go binding around an Ethereum contract.
type MetaMorphoFactory struct {
	MetaMorphoFactoryCaller     // Read-only binding to the contract
	MetaMorphoFactoryTransactor // Write-only binding to the contract
	MetaMorphoFactoryFilterer   // Log filterer for contract events
}

// MetaMorphoFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaMorphoFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMorphoFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaMorphoFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMorphoFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaMorphoFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMorphoFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaMorphoFactorySession struct {
	Contract     *MetaMorphoFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MetaMorphoFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaMorphoFactoryCallerSession struct {
	Contract *MetaMorphoFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MetaMorphoFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaMorphoFactoryTransactorSession struct {
	Contract     *MetaMorphoFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MetaMorphoFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaMorphoFactoryRaw struct {
	Contract *MetaMorphoFactory // Generic contract binding to access the raw methods on
}

// MetaMorphoFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaMorphoFactoryCallerRaw struct {
	Contract *MetaMorphoFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MetaMorphoFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaMorphoFactoryTransactorRaw struct {
	Contract *MetaMorphoFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaMorphoFactory creates a new instance of MetaMorphoFactory, bound to a specific deployed contract.
func NewMetaMorphoFactory(address common.Address, backend bind.ContractBackend) (*MetaMorphoFactory, error) {
	contract, err := bindMetaMorphoFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoFactory{MetaMorphoFactoryCaller: MetaMorphoFactoryCaller{contract: contract}, MetaMorphoFactoryTransactor: MetaMorphoFactoryTransactor{contract: contract}, MetaMorphoFactoryFilterer: MetaMorphoFactoryFilterer{contract: contract}}, nil
}

// NewMetaMorphoFactoryCaller creates a new read-only instance of MetaMorphoFactory, bound to a specific deployed contract.
func NewMetaMorphoFactoryCaller(address common.Address, caller bind.ContractCaller) (*MetaMorphoFactoryCaller, error) {
	contract, err := bindMetaMorphoFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoFactoryCaller{contract: contract}, nil
}

// NewMetaMorphoFactoryTransactor creates a new write-only instance of MetaMorphoFactory, bound to a specific deployed contract.
func NewMetaMorphoFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaMorphoFactoryTransactor, error) {
	contract, err := bindMetaMorphoFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoFactoryTransactor{contract: contract}, nil
}

// NewMetaMorphoFactoryFilterer creates a new log filterer instance of MetaMorphoFactory, bound to a specific deployed contract.
func NewMetaMorphoFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaMorphoFactoryFilterer, error) {
	contract, err := bindMetaMorphoFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoFactoryFilterer{contract: contract}, nil
}

// bindMetaMorphoFactory binds a generic wrapper to an already deployed contract.
func bindMetaMorphoFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetaMorphoFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaMorphoFactory *MetaMorphoFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaMorphoFactory.Contract.MetaMorphoFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaMorphoFactory *MetaMorphoFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoFactory.Contract.MetaMorphoFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaMorphoFactory *MetaMorphoFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaMorphoFactory.Contract.MetaMorphoFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaMorphoFactory *MetaMorphoFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaMorphoFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaMorphoFactory *MetaMorphoFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaMorphoFactory *MetaMorphoFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaMorphoFactory.Contract.contract.Transact(opts, method, params...)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MetaMorphoFactory *MetaMorphoFactoryCaller) MORPHO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoFactory.contract.Call(opts, &out, "MORPHO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MetaMorphoFactory *MetaMorphoFactorySession) MORPHO() (common.Address, error) {
	return _MetaMorphoFactory.Contract.MORPHO(&_MetaMorphoFactory.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MetaMorphoFactory *MetaMorphoFactoryCallerSession) MORPHO() (common.Address, error) {
	return _MetaMorphoFactory.Contract.MORPHO(&_MetaMorphoFactory.CallOpts)
}

// IsMetaMorpho is a free data retrieval call binding the contract method 0x29b5352c.
//
// Solidity: function isMetaMorpho(address ) view returns(bool)
func (_MetaMorphoFactory *MetaMorphoFactoryCaller) IsMetaMorpho(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MetaMorphoFactory.contract.Call(opts, &out, "isMetaMorpho", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMetaMorpho is a free data retrieval call binding the contract method 0x29b5352c.
//
// Solidity: function isMetaMorpho(address ) view returns(bool)
func (_MetaMorphoFactory *MetaMorphoFactorySession) IsMetaMorpho(arg0 common.Address) (bool, error) {
	return _MetaMorphoFactory.Contract.IsMetaMorpho(&_MetaMorphoFactory.CallOpts, arg0)
}

// IsMetaMorpho is a free data retrieval call binding the contract method 0x29b5352c.
//
// Solidity: function isMetaMorpho(address ) view returns(bool)
func (_MetaMorphoFactory *MetaMorphoFactoryCallerSession) IsMetaMorpho(arg0 common.Address) (bool, error) {
	return _MetaMorphoFactory.Contract.IsMetaMorpho(&_MetaMorphoFactory.CallOpts, arg0)
}

// CreateMetaMorpho is a paid mutator transaction binding the contract method 0xb5102025.
//
// Solidity: function createMetaMorpho(address initialOwner, uint256 initialTimelock, address asset, string name, string symbol, bytes32 salt) returns(address metaMorpho)
func (_MetaMorphoFactory *MetaMorphoFactoryTransactor) CreateMetaMorpho(opts *bind.TransactOpts, initialOwner common.Address, initialTimelock *big.Int, asset common.Address, name string, symbol string, salt [32]byte) (*types.Transaction, error) {
	return _MetaMorphoFactory.contract.Transact(opts, "createMetaMorpho", initialOwner, initialTimelock, asset, name, symbol, salt)
}

// CreateMetaMorpho is a paid mutator transaction binding the contract method 0xb5102025.
//
// Solidity: function createMetaMorpho(address initialOwner, uint256 initialTimelock, address asset, string name, string symbol, bytes32 salt) returns(address metaMorpho)
func (_MetaMorphoFactory *MetaMorphoFactorySession) CreateMetaMorpho(initialOwner common.Address, initialTimelock *big.Int, asset common.Address, name string, symbol string, salt [32]byte) (*types.Transaction, error) {
	return _MetaMorphoFactory.Contract.CreateMetaMorpho(&_MetaMorphoFactory.TransactOpts, initialOwner, initialTimelock, asset, name, symbol, salt)
}

// CreateMetaMorpho is a paid mutator transaction binding the contract method 0xb5102025.
//
// Solidity: function createMetaMorpho(address initialOwner, uint256 initialTimelock, address asset, string name, string symbol, bytes32 salt) returns(address metaMorpho)
func (_MetaMorphoFactory *MetaMorphoFactoryTransactorSession) CreateMetaMorpho(initialOwner common.Address, initialTimelock *big.Int, asset common.Address, name string, symbol string, salt [32]byte) (*types.Transaction, error) {
	return _MetaMorphoFactory.Contract.CreateMetaMorpho(&_MetaMorphoFactory.TransactOpts, initialOwner, initialTimelock, asset, name, symbol, salt)
}

// MetaMorphoFactoryCreateMetaMorphoIterator is returned from FilterCreateMetaMorpho and is used to iterate over the raw logs and unpacked data for CreateMetaMorpho events raised by the MetaMorphoFactory contract.
type MetaMorphoFactoryCreateMetaMorphoIterator struct {
	Event *MetaMorphoFactoryCreateMetaMorpho // Event containing the contract specifics and raw log

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
func (it *MetaMorphoFactoryCreateMetaMorphoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoFactoryCreateMetaMorpho)
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
		it.Event = new(MetaMorphoFactoryCreateMetaMorpho)
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
func (it *MetaMorphoFactoryCreateMetaMorphoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoFactoryCreateMetaMorphoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoFactoryCreateMetaMorpho represents a CreateMetaMorpho event raised by the MetaMorphoFactory contract.
type MetaMorphoFactoryCreateMetaMorpho struct {
	MetaMorpho      common.Address
	Caller          common.Address
	InitialOwner    common.Address
	InitialTimelock *big.Int
	Asset           common.Address
	Name            string
	Symbol          string
	Salt            [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateMetaMorpho is a free log retrieval operation binding the contract event 0xed8c95d05909b0f217f3e68171ef917df4b278d5addfe4dda888e90279be7d1d.
//
// Solidity: event CreateMetaMorpho(address indexed metaMorpho, address indexed caller, address initialOwner, uint256 initialTimelock, address indexed asset, string name, string symbol, bytes32 salt)
func (_MetaMorphoFactory *MetaMorphoFactoryFilterer) FilterCreateMetaMorpho(opts *bind.FilterOpts, metaMorpho []common.Address, caller []common.Address, asset []common.Address) (*MetaMorphoFactoryCreateMetaMorphoIterator, error) {

	var metaMorphoRule []interface{}
	for _, metaMorphoItem := range metaMorpho {
		metaMorphoRule = append(metaMorphoRule, metaMorphoItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _MetaMorphoFactory.contract.FilterLogs(opts, "CreateMetaMorpho", metaMorphoRule, callerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoFactoryCreateMetaMorphoIterator{contract: _MetaMorphoFactory.contract, event: "CreateMetaMorpho", logs: logs, sub: sub}, nil
}

// WatchCreateMetaMorpho is a free log subscription operation binding the contract event 0xed8c95d05909b0f217f3e68171ef917df4b278d5addfe4dda888e90279be7d1d.
//
// Solidity: event CreateMetaMorpho(address indexed metaMorpho, address indexed caller, address initialOwner, uint256 initialTimelock, address indexed asset, string name, string symbol, bytes32 salt)
func (_MetaMorphoFactory *MetaMorphoFactoryFilterer) WatchCreateMetaMorpho(opts *bind.WatchOpts, sink chan<- *MetaMorphoFactoryCreateMetaMorpho, metaMorpho []common.Address, caller []common.Address, asset []common.Address) (event.Subscription, error) {

	var metaMorphoRule []interface{}
	for _, metaMorphoItem := range metaMorpho {
		metaMorphoRule = append(metaMorphoRule, metaMorphoItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _MetaMorphoFactory.contract.WatchLogs(opts, "CreateMetaMorpho", metaMorphoRule, callerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoFactoryCreateMetaMorpho)
				if err := _MetaMorphoFactory.contract.UnpackLog(event, "CreateMetaMorpho", log); err != nil {
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

// ParseCreateMetaMorpho is a log parse operation binding the contract event 0xed8c95d05909b0f217f3e68171ef917df4b278d5addfe4dda888e90279be7d1d.
//
// Solidity: event CreateMetaMorpho(address indexed metaMorpho, address indexed caller, address initialOwner, uint256 initialTimelock, address indexed asset, string name, string symbol, bytes32 salt)
func (_MetaMorphoFactory *MetaMorphoFactoryFilterer) ParseCreateMetaMorpho(log types.Log) (*MetaMorphoFactoryCreateMetaMorpho, error) {
	event := new(MetaMorphoFactoryCreateMetaMorpho)
	if err := _MetaMorphoFactory.contract.UnpackLog(event, "CreateMetaMorpho", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
