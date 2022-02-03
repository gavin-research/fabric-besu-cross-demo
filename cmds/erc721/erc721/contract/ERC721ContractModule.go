// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccountData is an auto generated low-level Go binding around an user-defined struct.
type AccountData struct {
	Id       []byte
	AuthType AuthTypeData
}

// AuthTypeData is an auto generated low-level Go binding around an user-defined struct.
type AuthTypeData struct {
	Mode   uint8
	Option GoogleProtobufAnyData
}

// CrossContext is an auto generated low-level Go binding around an user-defined struct.
type CrossContext struct {
	TxId    []byte
	TxIndex uint8
	Signers []AccountData
}

// GoogleProtobufAnyData is an auto generated low-level Go binding around an user-defined struct.
type GoogleProtobufAnyData struct {
	TypeUrl string
	Value   []byte
}

// ERC721ContractModuleABI is the input ABI used to generate the binding from.
const ERC721ContractModuleABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CALLABLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE_URL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"setCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"tx_id\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"tx_index\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumAuthType.AuthMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"type_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structGoogleProtobufAny.Data\",\"name\":\"option\",\"type\":\"tuple\"}],\"internalType\":\"structAuthType.Data\",\"name\":\"auth_type\",\"type\":\"tuple\"}],\"internalType\":\"structAccount.Data[]\",\"name\":\"signers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCrossContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"callInfo\",\"type\":\"bytes\"}],\"name\":\"onContractCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721ContractModuleBin is the compiled bytecode used for deploying new contracts.
var ERC721ContractModuleBin = "0x60806040523480156200001157600080fd5b50620000286000336001600160e01b036200002e16565b62000166565b6200004382826001600160e01b036200004716565b5050565b6000828152602081815260409091206200006c9183906200078b620000c9821b17901c565b156200004357620000856001600160e01b03620000f216565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000620000e9836001600160a01b0384166001600160e01b03620000f616565b90505b92915050565b3390565b60006200010d83836001600160e01b036200014e16565b6200014557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620000ec565b506000620000ec565b60009081526001919091016020526040902054151590565b61167980620001766000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063a217fddf11610081578063c25054ab1161005b578063c25054ab146101ad578063ca15c873146101b5578063d547741f146101c8576100d4565b8063a217fddf1461018a578063a23f0b6a14610192578063beb92f551461019a576100d4565b80637376ec87116100b25780637376ec871461012a5780639010d07c1461014a57806391d148541461016a576100d4565b8063248a9ca3146100d95780632f2ff15d1461010257806336568abe14610117575b600080fd5b6100ec6100e7366004611022565b6101db565b6040516100f99190611214565b60405180910390f35b61011561011036600461103a565b6101f3565b005b61011561012536600461103a565b610244565b61013d6101383660046110b4565b610286565b6040516100f9919061121d565b61015d610158366004611074565b610691565b6040516100f991906111d1565b61017d61017836600461103a565b6106b8565b6040516100f99190611209565b6100ec6106d6565b6100ec6106db565b6101156101a8366004610ffb565b6106f2565b6100ec610712565b6100ec6101c3366004611022565b61073a565b6101156101d636600461103a565b610751565b6000818152602081905260409020600201545b919050565b600082815260208190526040902060020154610211906101786107a0565b6102365760405162461bcd60e51b815260040161022d906112cd565b60405180910390fd5b61024082826107a4565b5050565b61024c6107a0565b6001600160a01b0316816001600160a01b03161461027c5760405162461bcd60e51b815260040161022d9061152e565b6102408282610813565b60606102a5604051610297906111a8565b6040518091039020336106b8565b6102c15760405162461bcd60e51b815260040161022d906114c0565b6102ce604085018561158b565b90506001146102ef5760405162461bcd60e51b815260040161022d906114f7565b60036102fe604086018661158b565b600081811061030957fe5b905060200281019061031b9190611624565b610329906020810190611624565b610337906020810190611095565b600381111561034257fe5b1461035f5760405162461bcd60e51b815260040161022d906113be565b60405160200161036e90611159565b60405160208183030381529060405280519060200120848060400190610394919061158b565b600081811061039f57fe5b90506020028101906103b19190611624565b6103bf906020810190611624565b6103cd906020810190611624565b6103d790806115dd565b6040516020016103e8929190611149565b604051602081830303815290604052805190602001201461041b5760405162461bcd60e51b815260040161022d906113f5565b610423610f9d565b61046284848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061088292505050565b60408051808201909152600c81527f7472616e7366657246726f6d00000000000000000000000000000000000000006020918201528082015180519101209091507f0c41b033f4a31df8067c24d1e9b550a2ce75fd4a29e1147af9752174f0e6cb2014156106725760606104d98260400151610934565b905080516002146104fc5760405162461bcd60e51b815260040161022d9061132a565b600061056f61050e604089018961158b565b600081811061051957fe5b905060200281019061052b9190611624565b61053590806115dd565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a1c92505050565b905060006105908360008151811061058357fe5b6020026020010151610a23565b905060006105b1846001815181106105a457fe5b6020026020010151610a39565b85516040517f23b872dd0000000000000000000000000000000000000000000000000000000081529192506001600160a01b0316906323b872dd906105fe908690869086906004016111e5565b600060405180830381600087803b15801561061857600080fd5b505af115801561062c573d6000803e3d6000fd5b505050506040518060400160405280600c81526020017f63616c6c207375636365656400000000000000000000000000000000000000008152509550505050505061068a565b60405162461bcd60e51b815260040161022d90611452565b9392505050565b60008281526020819052604081206106af908363ffffffff610a8716565b90505b92915050565b60008281526020819052604081206106af908363ffffffff610a9316565b600081565b6040516106e7906111a8565b604051809103902081565b61070f604051610701906111a8565b6040518091039020826101f3565b50565b60405160200161072190611159565b6040516020818303038152906040528051906020012081565b60008181526020819052604081206106b290610aa8565b60008281526020819052604090206002015461076f906101786107a0565b61027c5760405162461bcd60e51b815260040161022d90611361565b60006106af836001600160a01b038416610ab3565b3390565b60008281526020819052604090206107c2908263ffffffff61078b16565b15610240576107cf6107a0565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610831908263ffffffff610afd16565b156102405761083e6107a0565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b61088a610f9d565b81606061089e61089983610b12565b610934565b90506002815110156108c25760405162461bcd60e51b815260040161022d90611489565b6108ca610fc1565b6108db6108d686610b12565b610b37565b90506108ee6108e982610b81565b610a23565b6001600160a01b0316845261090a61090582610b81565b610bcf565b602085015261091881610c4c565b1561092c5761092681610b81565b60408501525b505050919050565b606061093f82610c6d565b61094857600080fd5b600061095383610ca7565b905060608167ffffffffffffffff8111801561096e57600080fd5b506040519080825280602002602001820160405280156109a857816020015b610995610fe1565b81526020019060019003908161098d5790505b50905060006109ba8560200151610d03565b60208601510190506000805b84811015610a11576109d783610d66565b91506040518060400160405280838152602001848152508482815181106109fa57fe5b6020908102919091010152918101916001016109c6565b509195945050505050565b6014015190565b8051600090601514610a3457600080fd5b6106b2825b805160009015801590610a4e57508151602110155b610a5757600080fd5b600080610a6384610dff565b815191935091506020821015610a7f5760208290036101000a90045b949350505050565b60006106af8383610e25565b60006106af836001600160a01b038416610e6a565b60006106b282610e82565b6000610abf8383610e6a565b610af5575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556106b2565b5060006106b2565b60006106af836001600160a01b038416610e86565b610b1a610fe1565b506040805180820190915281518152602082810190820152919050565b610b3f610fc1565b610b4882610c6d565b610b5157600080fd5b6000610b608360200151610d03565b60208085015160408051808201909152868152920190820152915050919050565b610b89610fe1565b610b9282610c4c565b610b9b57600080fd5b60208201516000610bab82610d66565b80830160209586015260408051808201909152908152938401919091525090919050565b8051606090610bdd57600080fd5b600080610be984610dff565b9150915060608167ffffffffffffffff81118015610c0657600080fd5b506040519080825280601f01601f191660200182016040528015610c31576020820181803683370190505b50905060208101610c43848285610f4c565b50949350505050565b6000610c56610fe1565b505080518051602091820151919092015191011190565b8051600090610c7e575060006101ee565b6020820151805160001a9060c0821015610c9d576000925050506101ee565b5060019392505050565b8051600090610cb8575060006101ee565b60008090506000610ccc8460200151610d03565b602085015185519181019250015b80821015610cfa57610ceb82610d66565b60019093019290910190610cda565b50909392505050565b8051600090811a6080811015610d1d5760009150506101ee565b60b8811080610d38575060c08110801590610d38575060f881105b15610d475760019150506101ee565b60c0811015610d5b5760b5190190506101ee565b60f5190190506101ee565b80516000908190811a6080811015610d815760019150610df8565b60b8811015610d9657607e1981019150610df8565b60c0811015610dc35760b78103600185019450806020036101000a85510460018201810193505050610df8565b60f8811015610dd85760be1981019150610df8565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b6000806000610e118460200151610d03565b602085015194519481019594039392505050565b81546000908210610e485760405162461bcd60e51b815260040161022d90611270565b826000018281548110610e5757fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60008181526001830160205260408120548015610f425783546000198083019190810190600090879083908110610eb957fe5b9060005260206000200154905080876000018481548110610ed657fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080610f0657fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506106b2565b60009150506106b2565b80610f5657610f98565b5b60208110610f76578251825260209283019290910190601f1901610f57565b8015610f98578251825160208390036101000a60001901801990921691161782525b505050565b6040805160608082018352600082526020820152908101610fbc610fe1565b905290565b6040518060400160405280610fd4610fe1565b8152602001600081525090565b604051806040016040528060008152602001600081525090565b60006020828403121561100c578081fd5b81356001600160a01b038116811461068a578182fd5b600060208284031215611033578081fd5b5035919050565b6000806040838503121561104c578081fd5b8235915060208301356001600160a01b0381168114611069578182fd5b809150509250929050565b60008060408385031215611086578182fd5b50508035926020909101359150565b6000602082840312156110a6578081fd5b81356004811061068a578182fd5b6000806000604084860312156110c8578081fd5b833567ffffffffffffffff808211156110df578283fd5b818601606081890312156110f1578384fd5b94506020860135915080821115611106578283fd5b81860187601f820112611117578384fd5b8035925081831115611127578384fd5b876020848301011115611138578384fd5b949760209095019650909450505050565b6000828483379101908152919050565b7f2f657874656e73696f6e2e74797065732e4265737541757468457874656e736981527f6f6e000000000000000000000000000000000000000000000000000000000000602082015260220190565b7f43414c4c41424c455f524f4c45000000000000000000000000000000000000008152600d0190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b901515815260200190565b90815260200190565b6000602080835283518082850152825b818110156112495785810183015185820160400152820161122d565b8181111561125a5783604083870101525b50601f01601f1916929092016040019392505050565b60208082526022908201527f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60408201527f6473000000000000000000000000000000000000000000000000000000000000606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201527f2061646d696e20746f206772616e740000000000000000000000000000000000606082015260800190565b60208082526016908201527f696e76616c6964206c656e677468206f66206172677300000000000000000000604082015260600190565b60208082526030908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201527f2061646d696e20746f207265766f6b6500000000000000000000000000000000606082015260800190565b6020808252601b908201527f61757468206d6f6465206d75737420626520455854454e53494f4e0000000000604082015260600190565b60208082526023908201527f657874656e73696f6e206d757374206265204265737541757468457874656e7360408201527f696f6e0000000000000000000000000000000000000000000000000000000000606082015260800190565b60208082526010908201527f696e76616c69642063616c6c496e666f00000000000000000000000000000000604082015260600190565b6020808252601b908201527f6974656d73206d757374206265206d6f7265207468616e2074776f0000000000604082015260600190565b6020808252601d908201527f6d73672e73656e646572206d757374206265207468652063616c6c6572000000604082015260600190565b60208082526018908201527f7369676e657273206c656e677468206d75737420626520310000000000000000604082015260600190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201527f20726f6c657320666f722073656c660000000000000000000000000000000000606082015260800190565b6000808335601e198436030181126115a1578283fd5b8084018035925067ffffffffffffffff8311156115bc578384fd5b60208101935050506020810236038213156115d657600080fd5b9250929050565b6000808335601e198436030181126115f3578283fd5b8084018035925067ffffffffffffffff83111561160e578384fd5b602001925050368190038213156115d657600080fd5b60008235603e19833603018112611639578182fd5b919091019291505056fea26469706673582212205a27f6659bbe19979feebee45296d7eb963d3b67c7ea169e8d2f5e25f708cfad64736f6c63430006080033"

// DeployERC721ContractModule deploys a new Ethereum contract, binding an instance of ERC721ContractModule to it.
func DeployERC721ContractModule(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721ContractModule, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ContractModuleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721ContractModuleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721ContractModule{ERC721ContractModuleCaller: ERC721ContractModuleCaller{contract: contract}, ERC721ContractModuleTransactor: ERC721ContractModuleTransactor{contract: contract}, ERC721ContractModuleFilterer: ERC721ContractModuleFilterer{contract: contract}}, nil
}

// ERC721ContractModule is an auto generated Go binding around an Ethereum contract.
type ERC721ContractModule struct {
	ERC721ContractModuleCaller     // Read-only binding to the contract
	ERC721ContractModuleTransactor // Write-only binding to the contract
	ERC721ContractModuleFilterer   // Log filterer for contract events
}

// ERC721ContractModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721ContractModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ContractModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721ContractModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ContractModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721ContractModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ContractModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721ContractModuleSession struct {
	Contract     *ERC721ContractModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC721ContractModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721ContractModuleCallerSession struct {
	Contract *ERC721ContractModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ERC721ContractModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721ContractModuleTransactorSession struct {
	Contract     *ERC721ContractModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ERC721ContractModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721ContractModuleRaw struct {
	Contract *ERC721ContractModule // Generic contract binding to access the raw methods on
}

// ERC721ContractModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721ContractModuleCallerRaw struct {
	Contract *ERC721ContractModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721ContractModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721ContractModuleTransactorRaw struct {
	Contract *ERC721ContractModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721ContractModule creates a new instance of ERC721ContractModule, bound to a specific deployed contract.
func NewERC721ContractModule(address common.Address, backend bind.ContractBackend) (*ERC721ContractModule, error) {
	contract, err := bindERC721ContractModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModule{ERC721ContractModuleCaller: ERC721ContractModuleCaller{contract: contract}, ERC721ContractModuleTransactor: ERC721ContractModuleTransactor{contract: contract}, ERC721ContractModuleFilterer: ERC721ContractModuleFilterer{contract: contract}}, nil
}

// NewERC721ContractModuleCaller creates a new read-only instance of ERC721ContractModule, bound to a specific deployed contract.
func NewERC721ContractModuleCaller(address common.Address, caller bind.ContractCaller) (*ERC721ContractModuleCaller, error) {
	contract, err := bindERC721ContractModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModuleCaller{contract: contract}, nil
}

// NewERC721ContractModuleTransactor creates a new write-only instance of ERC721ContractModule, bound to a specific deployed contract.
func NewERC721ContractModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721ContractModuleTransactor, error) {
	contract, err := bindERC721ContractModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModuleTransactor{contract: contract}, nil
}

// NewERC721ContractModuleFilterer creates a new log filterer instance of ERC721ContractModule, bound to a specific deployed contract.
func NewERC721ContractModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721ContractModuleFilterer, error) {
	contract, err := bindERC721ContractModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModuleFilterer{contract: contract}, nil
}

// bindERC721ContractModule binds a generic wrapper to an already deployed contract.
func bindERC721ContractModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ContractModuleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721ContractModule *ERC721ContractModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721ContractModule.Contract.ERC721ContractModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721ContractModule *ERC721ContractModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.ERC721ContractModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721ContractModule *ERC721ContractModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.ERC721ContractModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721ContractModule *ERC721ContractModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721ContractModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721ContractModule *ERC721ContractModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721ContractModule *ERC721ContractModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.contract.Transact(opts, method, params...)
}

// CALLABLEROLE is a free data retrieval call binding the contract method 0xa23f0b6a.
//
// Solidity: function CALLABLE_ROLE() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCaller) CALLABLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "CALLABLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CALLABLEROLE is a free data retrieval call binding the contract method 0xa23f0b6a.
//
// Solidity: function CALLABLE_ROLE() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleSession) CALLABLEROLE() ([32]byte, error) {
	return _ERC721ContractModule.Contract.CALLABLEROLE(&_ERC721ContractModule.CallOpts)
}

// CALLABLEROLE is a free data retrieval call binding the contract method 0xa23f0b6a.
//
// Solidity: function CALLABLE_ROLE() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) CALLABLEROLE() ([32]byte, error) {
	return _ERC721ContractModule.Contract.CALLABLEROLE(&_ERC721ContractModule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC721ContractModule.Contract.DEFAULTADMINROLE(&_ERC721ContractModule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC721ContractModule.Contract.DEFAULTADMINROLE(&_ERC721ContractModule.CallOpts)
}

// TYPEURL is a free data retrieval call binding the contract method 0xc25054ab.
//
// Solidity: function TYPE_URL() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCaller) TYPEURL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "TYPE_URL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TYPEURL is a free data retrieval call binding the contract method 0xc25054ab.
//
// Solidity: function TYPE_URL() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleSession) TYPEURL() ([32]byte, error) {
	return _ERC721ContractModule.Contract.TYPEURL(&_ERC721ContractModule.CallOpts)
}

// TYPEURL is a free data retrieval call binding the contract method 0xc25054ab.
//
// Solidity: function TYPE_URL() view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) TYPEURL() ([32]byte, error) {
	return _ERC721ContractModule.Contract.TYPEURL(&_ERC721ContractModule.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC721ContractModule.Contract.GetRoleAdmin(&_ERC721ContractModule.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC721ContractModule.Contract.GetRoleAdmin(&_ERC721ContractModule.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC721ContractModule *ERC721ContractModuleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC721ContractModule *ERC721ContractModuleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC721ContractModule.Contract.GetRoleMember(&_ERC721ContractModule.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC721ContractModule.Contract.GetRoleMember(&_ERC721ContractModule.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC721ContractModule *ERC721ContractModuleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC721ContractModule *ERC721ContractModuleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC721ContractModule.Contract.GetRoleMemberCount(&_ERC721ContractModule.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC721ContractModule.Contract.GetRoleMemberCount(&_ERC721ContractModule.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721ContractModule *ERC721ContractModuleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721ContractModule.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721ContractModule *ERC721ContractModuleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC721ContractModule.Contract.HasRole(&_ERC721ContractModule.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721ContractModule *ERC721ContractModuleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC721ContractModule.Contract.HasRole(&_ERC721ContractModule.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.GrantRole(&_ERC721ContractModule.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.GrantRole(&_ERC721ContractModule.TransactOpts, role, account)
}

// OnContractCall is a paid mutator transaction binding the contract method 0x7376ec87.
//
// Solidity: function onContractCall((bytes,uint8,(bytes,(uint8,(string,bytes)))[]) context, bytes callInfo) returns(bytes)
func (_ERC721ContractModule *ERC721ContractModuleTransactor) OnContractCall(opts *bind.TransactOpts, context CrossContext, callInfo []byte) (*types.Transaction, error) {
	return _ERC721ContractModule.contract.Transact(opts, "onContractCall", context, callInfo)
}

// OnContractCall is a paid mutator transaction binding the contract method 0x7376ec87.
//
// Solidity: function onContractCall((bytes,uint8,(bytes,(uint8,(string,bytes)))[]) context, bytes callInfo) returns(bytes)
func (_ERC721ContractModule *ERC721ContractModuleSession) OnContractCall(context CrossContext, callInfo []byte) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.OnContractCall(&_ERC721ContractModule.TransactOpts, context, callInfo)
}

// OnContractCall is a paid mutator transaction binding the contract method 0x7376ec87.
//
// Solidity: function onContractCall((bytes,uint8,(bytes,(uint8,(string,bytes)))[]) context, bytes callInfo) returns(bytes)
func (_ERC721ContractModule *ERC721ContractModuleTransactorSession) OnContractCall(context CrossContext, callInfo []byte) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.OnContractCall(&_ERC721ContractModule.TransactOpts, context, callInfo)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.RenounceRole(&_ERC721ContractModule.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.RenounceRole(&_ERC721ContractModule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.RevokeRole(&_ERC721ContractModule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.RevokeRole(&_ERC721ContractModule.TransactOpts, role, account)
}

// SetCaller is a paid mutator transaction binding the contract method 0xbeb92f55.
//
// Solidity: function setCaller(address caller) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactor) SetCaller(opts *bind.TransactOpts, caller common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.contract.Transact(opts, "setCaller", caller)
}

// SetCaller is a paid mutator transaction binding the contract method 0xbeb92f55.
//
// Solidity: function setCaller(address caller) returns()
func (_ERC721ContractModule *ERC721ContractModuleSession) SetCaller(caller common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.SetCaller(&_ERC721ContractModule.TransactOpts, caller)
}

// SetCaller is a paid mutator transaction binding the contract method 0xbeb92f55.
//
// Solidity: function setCaller(address caller) returns()
func (_ERC721ContractModule *ERC721ContractModuleTransactorSession) SetCaller(caller common.Address) (*types.Transaction, error) {
	return _ERC721ContractModule.Contract.SetCaller(&_ERC721ContractModule.TransactOpts, caller)
}

// ERC721ContractModuleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC721ContractModule contract.
type ERC721ContractModuleRoleAdminChangedIterator struct {
	Event *ERC721ContractModuleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC721ContractModuleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractModuleRoleAdminChanged)
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
		it.Event = new(ERC721ContractModuleRoleAdminChanged)
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
func (it *ERC721ContractModuleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractModuleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractModuleRoleAdminChanged represents a RoleAdminChanged event raised by the ERC721ContractModule contract.
type ERC721ContractModuleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721ContractModule *ERC721ContractModuleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC721ContractModuleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ERC721ContractModule.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModuleRoleAdminChangedIterator{contract: _ERC721ContractModule.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721ContractModule *ERC721ContractModuleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC721ContractModuleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC721ContractModule.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractModuleRoleAdminChanged)
				if err := _ERC721ContractModule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ERC721ContractModule *ERC721ContractModuleFilterer) ParseRoleAdminChanged(log types.Log) (*ERC721ContractModuleRoleAdminChanged, error) {
	event := new(ERC721ContractModuleRoleAdminChanged)
	if err := _ERC721ContractModule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ContractModuleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC721ContractModule contract.
type ERC721ContractModuleRoleGrantedIterator struct {
	Event *ERC721ContractModuleRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC721ContractModuleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractModuleRoleGranted)
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
		it.Event = new(ERC721ContractModuleRoleGranted)
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
func (it *ERC721ContractModuleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractModuleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractModuleRoleGranted represents a RoleGranted event raised by the ERC721ContractModule contract.
type ERC721ContractModuleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721ContractModule *ERC721ContractModuleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC721ContractModuleRoleGrantedIterator, error) {

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

	logs, sub, err := _ERC721ContractModule.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModuleRoleGrantedIterator{contract: _ERC721ContractModule.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721ContractModule *ERC721ContractModuleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC721ContractModuleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC721ContractModule.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractModuleRoleGranted)
				if err := _ERC721ContractModule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ERC721ContractModule *ERC721ContractModuleFilterer) ParseRoleGranted(log types.Log) (*ERC721ContractModuleRoleGranted, error) {
	event := new(ERC721ContractModuleRoleGranted)
	if err := _ERC721ContractModule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ContractModuleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC721ContractModule contract.
type ERC721ContractModuleRoleRevokedIterator struct {
	Event *ERC721ContractModuleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC721ContractModuleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractModuleRoleRevoked)
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
		it.Event = new(ERC721ContractModuleRoleRevoked)
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
func (it *ERC721ContractModuleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractModuleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractModuleRoleRevoked represents a RoleRevoked event raised by the ERC721ContractModule contract.
type ERC721ContractModuleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721ContractModule *ERC721ContractModuleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC721ContractModuleRoleRevokedIterator, error) {

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

	logs, sub, err := _ERC721ContractModule.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractModuleRoleRevokedIterator{contract: _ERC721ContractModule.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721ContractModule *ERC721ContractModuleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC721ContractModuleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC721ContractModule.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractModuleRoleRevoked)
				if err := _ERC721ContractModule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ERC721ContractModule *ERC721ContractModuleFilterer) ParseRoleRevoked(log types.Log) (*ERC721ContractModuleRoleRevoked, error) {
	event := new(ERC721ContractModuleRoleRevoked)
	if err := _ERC721ContractModule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
