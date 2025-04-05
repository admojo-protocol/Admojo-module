package oracle

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
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

// PerformanceOracleMetaData contains all meta data concerning the PerformanceOracle contract.
var PerformanceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deviceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"views\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taps\",\"type\":\"uint256\"}],\"name\":\"MetricsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deviceFwHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deviceSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"getAggregatedMetrics\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalViews\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMetrics\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"views\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"metrics\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"views\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fwHash\",\"type\":\"bytes32\"}],\"name\":\"setDeviceAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_views\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taps\",\"type\":\"uint256\"}],\"name\":\"updateMetrics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_views\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taps\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_firmwareHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateMetricsWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taps\",\"type\":\"uint256\"}],\"name\":\"updateTaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deviceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_views\",\"type\":\"uint256\"}],\"name\":\"updateViews\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506114858061005c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c806371b44bb81161006f57806371b44bb81461015f578063848f34aa14610190578063aa65877e146101c1578063e7f8b5ef146101f2578063f713156d1461020e578063f851a4401461022a576100a7565b806312164256146100ab578063301e3e45146100db578063303d898a146100f757806351f0d2bb146101275780636d212fc914610143575b5f5ffd5b6100c560048036038101906100c09190610b85565b610248565b6040516100d29190610bc8565b60405180910390f35b6100f560048036038101906100f09190610be1565b61025d565b005b610111600480360381019061010c9190610b85565b6103b5565b60405161011e9190610c70565b60405180910390f35b610141600480360381019061013c9190610cdd565b6103e5565b005b61015d60048036038101906101589190610be1565b6104de565b005b61017960048036038101906101749190610be1565b610635565b604051610187929190610d3c565b60405180910390f35b6101aa60048036038101906101a59190610d63565b6106c3565b6040516101b8929190610d3c565b60405180910390f35b6101db60048036038101906101d69190610d63565b61071a565b6040516101e9929190610d3c565b60405180910390f35b61020c60048036038101906102079190610da1565b610744565b005b61022860048036038101906102239190610f41565b61085f565b005b610232610a4a565b60405161023f9190610c70565b60405180910390f35b6003602052805f5260405f205f915090505481565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e390611066565b60405180910390fd5b5f5f5f8581526020019081526020015f205f8481526020019081526020015f206040518060400160405290815f8201548152602001600182015481525050905081816020018181525050805f5f8681526020019081526020015f205f8581526020019081526020015f205f820151815f0155602082015181600101559050507f56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc88484835f015184602001516040516103a79493929190611084565b60405180910390a150505050565b6002602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046b90611066565b60405180910390fd5b8160025f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f8581526020019081526020015f2081905550505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461056d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056490611066565b60405180910390fd5b5f5f5f8581526020019081526020015f205f8481526020019081526020015f206040518060400160405290815f8201548152602001600182015481525050905081815f018181525050805f5f8681526020019081526020015f205f8581526020019081526020015f205f820151815f0155602082015181600101559050507f56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc88484835f015184602001516040516106279493929190611084565b60405180910390a150505050565b5f5f5f8490505b8381116106ba575f5f5f8881526020019081526020015f205f8381526020019081526020015f206040518060400160405290815f82015481526020016001820154815250509050805f01518461069291906110f4565b93508060200151836106a491906110f4565b92505080806106b290611127565b91505061063c565b50935093915050565b5f5f5f5f5f8681526020019081526020015f205f8581526020019081526020015f206040518060400160405290815f82015481526020016001820154815250509050805f0151816020015192509250509250929050565b5f602052815f5260405f20602052805f5260405f205f9150915050805f0154908060010154905082565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ca90611066565b60405180910390fd5b6040518060400160405280838152602001828152505f5f8681526020019081526020015f205f8581526020019081526020015f205f820151815f0155602082015181600101559050507f56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc8848484846040516108519493929190611084565b60405180910390a150505050565b60035f8781526020019081526020015f205482146108b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a9906111b8565b60405180910390fd5b5f86868686866040516020016108cc959493929190611216565b6040516020818303038152906040528051906020012090505f816040516020016108f691906112c8565b6040516020818303038152906040528051906020012090505f6109198285610a6f565b905060025f8a81526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146109b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b090611337565b60405180910390fd5b6040518060400160405280888152602001878152505f5f8b81526020019081526020015f205f8a81526020019081526020015f205f820151815f0155602082015181600101559050507f56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc889898989604051610a379493929190611084565b60405180910390a1505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f6041825114610ab4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aab9061139f565b60405180910390fd5b5f5f5f602085015192506040850151915060608501515f1a9050601b8160ff161015610aea57601b81610ae791906113c9565b90505b6001868285856040515f8152602001604052604051610b0c949392919061140c565b6020604051602081039080840390855afa158015610b2c573d5f5f3e3d5ffd5b50505060206040510351935050505092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610b6481610b52565b8114610b6e575f5ffd5b50565b5f81359050610b7f81610b5b565b92915050565b5f60208284031215610b9a57610b99610b4a565b5b5f610ba784828501610b71565b91505092915050565b5f819050919050565b610bc281610bb0565b82525050565b5f602082019050610bdb5f830184610bb9565b92915050565b5f5f5f60608486031215610bf857610bf7610b4a565b5b5f610c0586828701610b71565b9350506020610c1686828701610b71565b9250506040610c2786828701610b71565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610c5a82610c31565b9050919050565b610c6a81610c50565b82525050565b5f602082019050610c835f830184610c61565b92915050565b610c9281610c50565b8114610c9c575f5ffd5b50565b5f81359050610cad81610c89565b92915050565b610cbc81610bb0565b8114610cc6575f5ffd5b50565b5f81359050610cd781610cb3565b92915050565b5f5f5f60608486031215610cf457610cf3610b4a565b5b5f610d0186828701610b71565b9350506020610d1286828701610c9f565b9250506040610d2386828701610cc9565b9150509250925092565b610d3681610b52565b82525050565b5f604082019050610d4f5f830185610d2d565b610d5c6020830184610d2d565b9392505050565b5f5f60408385031215610d7957610d78610b4a565b5b5f610d8685828601610b71565b9250506020610d9785828601610b71565b9150509250929050565b5f5f5f5f60808587031215610db957610db8610b4a565b5b5f610dc687828801610b71565b9450506020610dd787828801610b71565b9350506040610de887828801610b71565b9250506060610df987828801610b71565b91505092959194509250565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e5382610e0d565b810181811067ffffffffffffffff82111715610e7257610e71610e1d565b5b80604052505050565b5f610e84610b41565b9050610e908282610e4a565b919050565b5f67ffffffffffffffff821115610eaf57610eae610e1d565b5b610eb882610e0d565b9050602081019050919050565b828183375f83830152505050565b5f610ee5610ee084610e95565b610e7b565b905082815260208101848484011115610f0157610f00610e09565b5b610f0c848285610ec5565b509392505050565b5f82601f830112610f2857610f27610e05565b5b8135610f38848260208601610ed3565b91505092915050565b5f5f5f5f5f5f60c08789031215610f5b57610f5a610b4a565b5b5f610f6889828a01610b71565b9650506020610f7989828a01610b71565b9550506040610f8a89828a01610b71565b9450506060610f9b89828a01610b71565b9350506080610fac89828a01610cc9565b92505060a087013567ffffffffffffffff811115610fcd57610fcc610b4e565b5b610fd989828a01610f14565b9150509295509295509295565b5f82825260208201905092915050565b7f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f611050602183610fe6565b915061105b82610ff6565b604082019050919050565b5f6020820190508181035f83015261107d81611044565b9050919050565b5f6080820190506110975f830187610d2d565b6110a46020830186610d2d565b6110b16040830185610d2d565b6110be6060830184610d2d565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110fe82610b52565b915061110983610b52565b9250828201905080821115611121576111206110c7565b5b92915050565b5f61113182610b52565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611163576111626110c7565b5b600182019050919050565b7f4669726d776172652068617368206d69736d61746368000000000000000000005f82015250565b5f6111a2601683610fe6565b91506111ad8261116e565b602082019050919050565b5f6020820190508181035f8301526111cf81611196565b9050919050565b5f819050919050565b6111f06111eb82610b52565b6111d6565b82525050565b5f819050919050565b61121061120b82610bb0565b6111f6565b82525050565b5f61122182886111df565b60208201915061123182876111df565b60208201915061124182866111df565b60208201915061125182856111df565b60208201915061126182846111ff565b6020820191508190509695505050505050565b5f81905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f6112b2601c83611274565b91506112bd8261127e565b601c82019050919050565b5f6112d2826112a6565b91506112de82846111ff565b60208201915081905092915050565b7f5369676e6174757265206e6f742066726f6d20646576696365207369676e65725f82015250565b5f611321602083610fe6565b915061132c826112ed565b602082019050919050565b5f6020820190508181035f83015261134e81611315565b9050919050565b7f496e76616c6964207369676e6174757265206c656e67746800000000000000005f82015250565b5f611389601883610fe6565b915061139482611355565b602082019050919050565b5f6020820190508181035f8301526113b68161137d565b9050919050565b5f60ff82169050919050565b5f6113d3826113bd565b91506113de836113bd565b9250828201905060ff8111156113f7576113f66110c7565b5b92915050565b611406816113bd565b82525050565b5f60808201905061141f5f830187610bb9565b61142c60208301866113fd565b6114396040830185610bb9565b6114466060830184610bb9565b9594505050505056fea264697066735822122016eda4958afdefe2a96833c80bf696ece54dfaef9dddbc423733d0ada9b3daa364736f6c634300081d0033",
}

// PerformanceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PerformanceOracleMetaData.ABI instead.
var PerformanceOracleABI = PerformanceOracleMetaData.ABI

// PerformanceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PerformanceOracleMetaData.Bin instead.
var PerformanceOracleBin = PerformanceOracleMetaData.Bin

// DeployPerformanceOracle deploys a new Ethereum contract, binding an instance of PerformanceOracle to it.
func DeployPerformanceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PerformanceOracle, error) {
	parsed, err := PerformanceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PerformanceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PerformanceOracle{PerformanceOracleCaller: PerformanceOracleCaller{contract: contract}, PerformanceOracleTransactor: PerformanceOracleTransactor{contract: contract}, PerformanceOracleFilterer: PerformanceOracleFilterer{contract: contract}}, nil
}

// PerformanceOracle is an auto generated Go binding around an Ethereum contract.
type PerformanceOracle struct {
	PerformanceOracleCaller     // Read-only binding to the contract
	PerformanceOracleTransactor // Write-only binding to the contract
	PerformanceOracleFilterer   // Log filterer for contract events
}

// PerformanceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerformanceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerformanceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerformanceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerformanceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerformanceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerformanceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerformanceOracleSession struct {
	Contract     *PerformanceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PerformanceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerformanceOracleCallerSession struct {
	Contract *PerformanceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PerformanceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerformanceOracleTransactorSession struct {
	Contract     *PerformanceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PerformanceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerformanceOracleRaw struct {
	Contract *PerformanceOracle // Generic contract binding to access the raw methods on
}

// PerformanceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerformanceOracleCallerRaw struct {
	Contract *PerformanceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PerformanceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerformanceOracleTransactorRaw struct {
	Contract *PerformanceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerformanceOracle creates a new instance of PerformanceOracle, bound to a specific deployed contract.
func NewPerformanceOracle(address common.Address, backend bind.ContractBackend) (*PerformanceOracle, error) {
	contract, err := bindPerformanceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerformanceOracle{PerformanceOracleCaller: PerformanceOracleCaller{contract: contract}, PerformanceOracleTransactor: PerformanceOracleTransactor{contract: contract}, PerformanceOracleFilterer: PerformanceOracleFilterer{contract: contract}}, nil
}

// NewPerformanceOracleCaller creates a new read-only instance of PerformanceOracle, bound to a specific deployed contract.
func NewPerformanceOracleCaller(address common.Address, caller bind.ContractCaller) (*PerformanceOracleCaller, error) {
	contract, err := bindPerformanceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerformanceOracleCaller{contract: contract}, nil
}

// NewPerformanceOracleTransactor creates a new write-only instance of PerformanceOracle, bound to a specific deployed contract.
func NewPerformanceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PerformanceOracleTransactor, error) {
	contract, err := bindPerformanceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerformanceOracleTransactor{contract: contract}, nil
}

// NewPerformanceOracleFilterer creates a new log filterer instance of PerformanceOracle, bound to a specific deployed contract.
func NewPerformanceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PerformanceOracleFilterer, error) {
	contract, err := bindPerformanceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerformanceOracleFilterer{contract: contract}, nil
}

// bindPerformanceOracle binds a generic wrapper to an already deployed contract.
func bindPerformanceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PerformanceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerformanceOracle *PerformanceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerformanceOracle.Contract.PerformanceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerformanceOracle *PerformanceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.PerformanceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerformanceOracle *PerformanceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.PerformanceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerformanceOracle *PerformanceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerformanceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerformanceOracle *PerformanceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerformanceOracle *PerformanceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PerformanceOracle *PerformanceOracleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerformanceOracle.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PerformanceOracle *PerformanceOracleSession) Admin() (common.Address, error) {
	return _PerformanceOracle.Contract.Admin(&_PerformanceOracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PerformanceOracle *PerformanceOracleCallerSession) Admin() (common.Address, error) {
	return _PerformanceOracle.Contract.Admin(&_PerformanceOracle.CallOpts)
}

// DeviceFwHash is a free data retrieval call binding the contract method 0x12164256.
//
// Solidity: function deviceFwHash(uint256 ) view returns(bytes32)
func (_PerformanceOracle *PerformanceOracleCaller) DeviceFwHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PerformanceOracle.contract.Call(opts, &out, "deviceFwHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeviceFwHash is a free data retrieval call binding the contract method 0x12164256.
//
// Solidity: function deviceFwHash(uint256 ) view returns(bytes32)
func (_PerformanceOracle *PerformanceOracleSession) DeviceFwHash(arg0 *big.Int) ([32]byte, error) {
	return _PerformanceOracle.Contract.DeviceFwHash(&_PerformanceOracle.CallOpts, arg0)
}

// DeviceFwHash is a free data retrieval call binding the contract method 0x12164256.
//
// Solidity: function deviceFwHash(uint256 ) view returns(bytes32)
func (_PerformanceOracle *PerformanceOracleCallerSession) DeviceFwHash(arg0 *big.Int) ([32]byte, error) {
	return _PerformanceOracle.Contract.DeviceFwHash(&_PerformanceOracle.CallOpts, arg0)
}

// DeviceSigner is a free data retrieval call binding the contract method 0x303d898a.
//
// Solidity: function deviceSigner(uint256 ) view returns(address)
func (_PerformanceOracle *PerformanceOracleCaller) DeviceSigner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PerformanceOracle.contract.Call(opts, &out, "deviceSigner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeviceSigner is a free data retrieval call binding the contract method 0x303d898a.
//
// Solidity: function deviceSigner(uint256 ) view returns(address)
func (_PerformanceOracle *PerformanceOracleSession) DeviceSigner(arg0 *big.Int) (common.Address, error) {
	return _PerformanceOracle.Contract.DeviceSigner(&_PerformanceOracle.CallOpts, arg0)
}

// DeviceSigner is a free data retrieval call binding the contract method 0x303d898a.
//
// Solidity: function deviceSigner(uint256 ) view returns(address)
func (_PerformanceOracle *PerformanceOracleCallerSession) DeviceSigner(arg0 *big.Int) (common.Address, error) {
	return _PerformanceOracle.Contract.DeviceSigner(&_PerformanceOracle.CallOpts, arg0)
}

// GetAggregatedMetrics is a free data retrieval call binding the contract method 0x71b44bb8.
//
// Solidity: function getAggregatedMetrics(uint256 _deviceId, uint256 _startTime, uint256 _endTime) view returns(uint256 totalViews, uint256 totalTaps)
func (_PerformanceOracle *PerformanceOracleCaller) GetAggregatedMetrics(opts *bind.CallOpts, _deviceId *big.Int, _startTime *big.Int, _endTime *big.Int) (struct {
	TotalViews *big.Int
	TotalTaps  *big.Int
}, error) {
	var out []interface{}
	err := _PerformanceOracle.contract.Call(opts, &out, "getAggregatedMetrics", _deviceId, _startTime, _endTime)

	outstruct := new(struct {
		TotalViews *big.Int
		TotalTaps  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalViews = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalTaps = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAggregatedMetrics is a free data retrieval call binding the contract method 0x71b44bb8.
//
// Solidity: function getAggregatedMetrics(uint256 _deviceId, uint256 _startTime, uint256 _endTime) view returns(uint256 totalViews, uint256 totalTaps)
func (_PerformanceOracle *PerformanceOracleSession) GetAggregatedMetrics(_deviceId *big.Int, _startTime *big.Int, _endTime *big.Int) (struct {
	TotalViews *big.Int
	TotalTaps  *big.Int
}, error) {
	return _PerformanceOracle.Contract.GetAggregatedMetrics(&_PerformanceOracle.CallOpts, _deviceId, _startTime, _endTime)
}

// GetAggregatedMetrics is a free data retrieval call binding the contract method 0x71b44bb8.
//
// Solidity: function getAggregatedMetrics(uint256 _deviceId, uint256 _startTime, uint256 _endTime) view returns(uint256 totalViews, uint256 totalTaps)
func (_PerformanceOracle *PerformanceOracleCallerSession) GetAggregatedMetrics(_deviceId *big.Int, _startTime *big.Int, _endTime *big.Int) (struct {
	TotalViews *big.Int
	TotalTaps  *big.Int
}, error) {
	return _PerformanceOracle.Contract.GetAggregatedMetrics(&_PerformanceOracle.CallOpts, _deviceId, _startTime, _endTime)
}

// GetMetrics is a free data retrieval call binding the contract method 0x848f34aa.
//
// Solidity: function getMetrics(uint256 _deviceId, uint256 _timestamp) view returns(uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleCaller) GetMetrics(opts *bind.CallOpts, _deviceId *big.Int, _timestamp *big.Int) (struct {
	Views *big.Int
	Taps  *big.Int
}, error) {
	var out []interface{}
	err := _PerformanceOracle.contract.Call(opts, &out, "getMetrics", _deviceId, _timestamp)

	outstruct := new(struct {
		Views *big.Int
		Taps  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Views = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Taps = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMetrics is a free data retrieval call binding the contract method 0x848f34aa.
//
// Solidity: function getMetrics(uint256 _deviceId, uint256 _timestamp) view returns(uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleSession) GetMetrics(_deviceId *big.Int, _timestamp *big.Int) (struct {
	Views *big.Int
	Taps  *big.Int
}, error) {
	return _PerformanceOracle.Contract.GetMetrics(&_PerformanceOracle.CallOpts, _deviceId, _timestamp)
}

// GetMetrics is a free data retrieval call binding the contract method 0x848f34aa.
//
// Solidity: function getMetrics(uint256 _deviceId, uint256 _timestamp) view returns(uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleCallerSession) GetMetrics(_deviceId *big.Int, _timestamp *big.Int) (struct {
	Views *big.Int
	Taps  *big.Int
}, error) {
	return _PerformanceOracle.Contract.GetMetrics(&_PerformanceOracle.CallOpts, _deviceId, _timestamp)
}

// Metrics is a free data retrieval call binding the contract method 0xaa65877e.
//
// Solidity: function metrics(uint256 , uint256 ) view returns(uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleCaller) Metrics(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Views *big.Int
	Taps  *big.Int
}, error) {
	var out []interface{}
	err := _PerformanceOracle.contract.Call(opts, &out, "metrics", arg0, arg1)

	outstruct := new(struct {
		Views *big.Int
		Taps  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Views = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Taps = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Metrics is a free data retrieval call binding the contract method 0xaa65877e.
//
// Solidity: function metrics(uint256 , uint256 ) view returns(uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleSession) Metrics(arg0 *big.Int, arg1 *big.Int) (struct {
	Views *big.Int
	Taps  *big.Int
}, error) {
	return _PerformanceOracle.Contract.Metrics(&_PerformanceOracle.CallOpts, arg0, arg1)
}

// Metrics is a free data retrieval call binding the contract method 0xaa65877e.
//
// Solidity: function metrics(uint256 , uint256 ) view returns(uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleCallerSession) Metrics(arg0 *big.Int, arg1 *big.Int) (struct {
	Views *big.Int
	Taps  *big.Int
}, error) {
	return _PerformanceOracle.Contract.Metrics(&_PerformanceOracle.CallOpts, arg0, arg1)
}

// SetDeviceAuth is a paid mutator transaction binding the contract method 0x51f0d2bb.
//
// Solidity: function setDeviceAuth(uint256 _deviceId, address _signer, bytes32 _fwHash) returns()
func (_PerformanceOracle *PerformanceOracleTransactor) SetDeviceAuth(opts *bind.TransactOpts, _deviceId *big.Int, _signer common.Address, _fwHash [32]byte) (*types.Transaction, error) {
	return _PerformanceOracle.contract.Transact(opts, "setDeviceAuth", _deviceId, _signer, _fwHash)
}

// SetDeviceAuth is a paid mutator transaction binding the contract method 0x51f0d2bb.
//
// Solidity: function setDeviceAuth(uint256 _deviceId, address _signer, bytes32 _fwHash) returns()
func (_PerformanceOracle *PerformanceOracleSession) SetDeviceAuth(_deviceId *big.Int, _signer common.Address, _fwHash [32]byte) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.SetDeviceAuth(&_PerformanceOracle.TransactOpts, _deviceId, _signer, _fwHash)
}

// SetDeviceAuth is a paid mutator transaction binding the contract method 0x51f0d2bb.
//
// Solidity: function setDeviceAuth(uint256 _deviceId, address _signer, bytes32 _fwHash) returns()
func (_PerformanceOracle *PerformanceOracleTransactorSession) SetDeviceAuth(_deviceId *big.Int, _signer common.Address, _fwHash [32]byte) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.SetDeviceAuth(&_PerformanceOracle.TransactOpts, _deviceId, _signer, _fwHash)
}

// UpdateMetrics is a paid mutator transaction binding the contract method 0xe7f8b5ef.
//
// Solidity: function updateMetrics(uint256 _deviceId, uint256 _timestamp, uint256 _views, uint256 _taps) returns()
func (_PerformanceOracle *PerformanceOracleTransactor) UpdateMetrics(opts *bind.TransactOpts, _deviceId *big.Int, _timestamp *big.Int, _views *big.Int, _taps *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.contract.Transact(opts, "updateMetrics", _deviceId, _timestamp, _views, _taps)
}

// UpdateMetrics is a paid mutator transaction binding the contract method 0xe7f8b5ef.
//
// Solidity: function updateMetrics(uint256 _deviceId, uint256 _timestamp, uint256 _views, uint256 _taps) returns()
func (_PerformanceOracle *PerformanceOracleSession) UpdateMetrics(_deviceId *big.Int, _timestamp *big.Int, _views *big.Int, _taps *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateMetrics(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _views, _taps)
}

// UpdateMetrics is a paid mutator transaction binding the contract method 0xe7f8b5ef.
//
// Solidity: function updateMetrics(uint256 _deviceId, uint256 _timestamp, uint256 _views, uint256 _taps) returns()
func (_PerformanceOracle *PerformanceOracleTransactorSession) UpdateMetrics(_deviceId *big.Int, _timestamp *big.Int, _views *big.Int, _taps *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateMetrics(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _views, _taps)
}

// UpdateMetricsWithSig is a paid mutator transaction binding the contract method 0xf713156d.
//
// Solidity: function updateMetricsWithSig(uint256 _deviceId, uint256 _timestamp, uint256 _views, uint256 _taps, bytes32 _firmwareHash, bytes _signature) returns()
func (_PerformanceOracle *PerformanceOracleTransactor) UpdateMetricsWithSig(opts *bind.TransactOpts, _deviceId *big.Int, _timestamp *big.Int, _views *big.Int, _taps *big.Int, _firmwareHash [32]byte, _signature []byte) (*types.Transaction, error) {
	return _PerformanceOracle.contract.Transact(opts, "updateMetricsWithSig", _deviceId, _timestamp, _views, _taps, _firmwareHash, _signature)
}

// UpdateMetricsWithSig is a paid mutator transaction binding the contract method 0xf713156d.
//
// Solidity: function updateMetricsWithSig(uint256 _deviceId, uint256 _timestamp, uint256 _views, uint256 _taps, bytes32 _firmwareHash, bytes _signature) returns()
func (_PerformanceOracle *PerformanceOracleSession) UpdateMetricsWithSig(_deviceId *big.Int, _timestamp *big.Int, _views *big.Int, _taps *big.Int, _firmwareHash [32]byte, _signature []byte) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateMetricsWithSig(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _views, _taps, _firmwareHash, _signature)
}

// UpdateMetricsWithSig is a paid mutator transaction binding the contract method 0xf713156d.
//
// Solidity: function updateMetricsWithSig(uint256 _deviceId, uint256 _timestamp, uint256 _views, uint256 _taps, bytes32 _firmwareHash, bytes _signature) returns()
func (_PerformanceOracle *PerformanceOracleTransactorSession) UpdateMetricsWithSig(_deviceId *big.Int, _timestamp *big.Int, _views *big.Int, _taps *big.Int, _firmwareHash [32]byte, _signature []byte) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateMetricsWithSig(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _views, _taps, _firmwareHash, _signature)
}

// UpdateTaps is a paid mutator transaction binding the contract method 0x301e3e45.
//
// Solidity: function updateTaps(uint256 _deviceId, uint256 _timestamp, uint256 _taps) returns()
func (_PerformanceOracle *PerformanceOracleTransactor) UpdateTaps(opts *bind.TransactOpts, _deviceId *big.Int, _timestamp *big.Int, _taps *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.contract.Transact(opts, "updateTaps", _deviceId, _timestamp, _taps)
}

// UpdateTaps is a paid mutator transaction binding the contract method 0x301e3e45.
//
// Solidity: function updateTaps(uint256 _deviceId, uint256 _timestamp, uint256 _taps) returns()
func (_PerformanceOracle *PerformanceOracleSession) UpdateTaps(_deviceId *big.Int, _timestamp *big.Int, _taps *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateTaps(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _taps)
}

// UpdateTaps is a paid mutator transaction binding the contract method 0x301e3e45.
//
// Solidity: function updateTaps(uint256 _deviceId, uint256 _timestamp, uint256 _taps) returns()
func (_PerformanceOracle *PerformanceOracleTransactorSession) UpdateTaps(_deviceId *big.Int, _timestamp *big.Int, _taps *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateTaps(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _taps)
}

// UpdateViews is a paid mutator transaction binding the contract method 0x6d212fc9.
//
// Solidity: function updateViews(uint256 _deviceId, uint256 _timestamp, uint256 _views) returns()
func (_PerformanceOracle *PerformanceOracleTransactor) UpdateViews(opts *bind.TransactOpts, _deviceId *big.Int, _timestamp *big.Int, _views *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.contract.Transact(opts, "updateViews", _deviceId, _timestamp, _views)
}

// UpdateViews is a paid mutator transaction binding the contract method 0x6d212fc9.
//
// Solidity: function updateViews(uint256 _deviceId, uint256 _timestamp, uint256 _views) returns()
func (_PerformanceOracle *PerformanceOracleSession) UpdateViews(_deviceId *big.Int, _timestamp *big.Int, _views *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateViews(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _views)
}

// UpdateViews is a paid mutator transaction binding the contract method 0x6d212fc9.
//
// Solidity: function updateViews(uint256 _deviceId, uint256 _timestamp, uint256 _views) returns()
func (_PerformanceOracle *PerformanceOracleTransactorSession) UpdateViews(_deviceId *big.Int, _timestamp *big.Int, _views *big.Int) (*types.Transaction, error) {
	return _PerformanceOracle.Contract.UpdateViews(&_PerformanceOracle.TransactOpts, _deviceId, _timestamp, _views)
}

// PerformanceOracleMetricsUpdatedIterator is returned from FilterMetricsUpdated and is used to iterate over the raw logs and unpacked data for MetricsUpdated events raised by the PerformanceOracle contract.
type PerformanceOracleMetricsUpdatedIterator struct {
	Event *PerformanceOracleMetricsUpdated // Event containing the contract specifics and raw log

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
func (it *PerformanceOracleMetricsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerformanceOracleMetricsUpdated)
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
		it.Event = new(PerformanceOracleMetricsUpdated)
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
func (it *PerformanceOracleMetricsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerformanceOracleMetricsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerformanceOracleMetricsUpdated represents a MetricsUpdated event raised by the PerformanceOracle contract.
type PerformanceOracleMetricsUpdated struct {
	DeviceId  *big.Int
	Timestamp *big.Int
	Views     *big.Int
	Taps      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMetricsUpdated is a free log retrieval operation binding the contract event 0x56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc8.
//
// Solidity: event MetricsUpdated(uint256 deviceId, uint256 timestamp, uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleFilterer) FilterMetricsUpdated(opts *bind.FilterOpts) (*PerformanceOracleMetricsUpdatedIterator, error) {

	logs, sub, err := _PerformanceOracle.contract.FilterLogs(opts, "MetricsUpdated")
	if err != nil {
		return nil, err
	}
	return &PerformanceOracleMetricsUpdatedIterator{contract: _PerformanceOracle.contract, event: "MetricsUpdated", logs: logs, sub: sub}, nil
}

// WatchMetricsUpdated is a free log subscription operation binding the contract event 0x56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc8.
//
// Solidity: event MetricsUpdated(uint256 deviceId, uint256 timestamp, uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleFilterer) WatchMetricsUpdated(opts *bind.WatchOpts, sink chan<- *PerformanceOracleMetricsUpdated) (event.Subscription, error) {

	logs, sub, err := _PerformanceOracle.contract.WatchLogs(opts, "MetricsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerformanceOracleMetricsUpdated)
				if err := _PerformanceOracle.contract.UnpackLog(event, "MetricsUpdated", log); err != nil {
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

// ParseMetricsUpdated is a log parse operation binding the contract event 0x56d8f7dc7f4e1d9960c9a876b8fe805810b535361b52ba0e6c1b6b6529f2bcc8.
//
// Solidity: event MetricsUpdated(uint256 deviceId, uint256 timestamp, uint256 views, uint256 taps)
func (_PerformanceOracle *PerformanceOracleFilterer) ParseMetricsUpdated(log types.Log) (*PerformanceOracleMetricsUpdated, error) {
	event := new(PerformanceOracleMetricsUpdated)
	if err := _PerformanceOracle.contract.UnpackLog(event, "MetricsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
