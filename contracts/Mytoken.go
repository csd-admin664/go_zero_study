// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// MyTokenMetaData contains all meta data concerning the MyToken contract.
var MyTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"staabigenteMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611a23380380611a2383398181016040528101906100319190610267565b336040518060400160405280600c81526020017f52434344656d6f546f6b656e00000000000000000000000000000000000000008152506040518060400160405280600381526020017f524454000000000000000000000000000000000000000000000000000000000081525081600390816100ad91906104cf565b5080600490816100bd91906104cf565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610130575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161012791906105ad565b60405180910390fd5b61013f8161014660201b60201c565b50506105c6565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102368261020d565b9050919050565b6102468161022c565b8114610250575f5ffd5b50565b5f815190506102618161023d565b92915050565b5f6020828403121561027c5761027b610209565b5b5f61028984828501610253565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061030d57607f821691505b6020821081036103205761031f6102c9565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610347565b61038c8683610347565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103d06103cb6103c6846103a4565b6103ad565b6103a4565b9050919050565b5f819050919050565b6103e9836103b6565b6103fd6103f5826103d7565b848454610353565b825550505050565b5f5f905090565b610414610405565b61041f8184846103e0565b505050565b5b81811015610442576104375f8261040c565b600181019050610425565b5050565b601f8211156104875761045881610326565b61046184610338565b81016020851015610470578190505b61048461047c85610338565b830182610424565b50505b505050565b5f82821c905092915050565b5f6104a75f198460080261048c565b1980831691505092915050565b5f6104bf8383610498565b9150826002028217905092915050565b6104d882610292565b67ffffffffffffffff8111156104f1576104f061029c565b5b6104fb82546102f6565b610506828285610446565b5f60209050601f831160018114610537575f8415610525578287015190505b61052f85826104b4565b865550610596565b601f19841661054586610326565b5f5b8281101561056c57848901518255600182019150602085019450602081019050610547565b868310156105895784890151610585601f891682610498565b8355505b6001600288020188555050505b505050505050565b6105a78161022c565b82525050565b5f6020820190506105c05f83018461059e565b92915050565b611450806105d35f395ff3fe6080604052600436106100f6575f3560e01c8063715018a611610089578063d586b74c11610058578063d586b74c14610315578063dd62ed3e1461033f578063e086e5ec1461037b578063f2fde38b1461039157610105565b8063715018a61461026f5780638da5cb5b1461028557806395d89b41146102af578063a9059cbb146102d957610105565b806323b872dd116100c557806323b872dd146101a3578063313ce567146101df578063664e97041461020957806370a082311461023357610105565b806306fdde0314610109578063095ea7b3146101335780631249c58b1461016f57806318160ddd1461017957610105565b36610105576101036103b9565b005b5f5ffd5b348015610114575f5ffd5b5061011d610423565b60405161012a9190610fb8565b60405180910390f35b34801561013e575f5ffd5b5061015960048036038101906101549190611069565b6104b3565b60405161016691906110c1565b60405180910390f35b6101776103b9565b005b348015610184575f5ffd5b5061018d6104d5565b60405161019a91906110e9565b60405180910390f35b3480156101ae575f5ffd5b506101c960048036038101906101c49190611102565b6104de565b6040516101d691906110c1565b60405180910390f35b3480156101ea575f5ffd5b506101f361050c565b604051610200919061116d565b60405180910390f35b348015610214575f5ffd5b5061021d610514565b60405161022a91906110e9565b60405180910390f35b34801561023e575f5ffd5b5061025960048036038101906102549190611186565b61051c565b60405161026691906110e9565b60405180910390f35b34801561027a575f5ffd5b50610283610561565b005b348015610290575f5ffd5b50610299610574565b6040516102a691906111c0565b60405180910390f35b3480156102ba575f5ffd5b506102c361059c565b6040516102d09190610fb8565b60405180910390f35b3480156102e4575f5ffd5b506102ff60048036038101906102fa9190611069565b61062c565b60405161030c91906110c1565b60405180910390f35b348015610320575f5ffd5b5061032961064e565b60405161033691906110e9565b60405180910390f35b34801561034a575f5ffd5b50610365600480360381019061036091906111d9565b610659565b60405161037291906110e9565b60405180910390f35b348015610386575f5ffd5b5061038f6106db565b005b34801561039c575f5ffd5b506103b760048036038101906103b29190611186565b610777565b005b66038d7ea4c68000341015610403576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fa90611261565b60405180910390fd5b5f6305f5e1003461041491906112ac565b905061042033826107fb565b50565b6060600380546104329061131a565b80601f016020809104026020016040519081016040528092919081815260200182805461045e9061131a565b80156104a95780601f10610480576101008083540402835291602001916104a9565b820191905f5260205f20905b81548152906001019060200180831161048c57829003601f168201915b5050505050905090565b5f5f6104bd61087a565b90506104ca818585610881565b600191505092915050565b5f600254905090565b5f5f6104e861087a565b90506104f5858285610893565b610500858585610926565b60019150509392505050565b5f6012905090565b6305f5e10081565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610569610a16565b6105725f610a9d565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546105ab9061131a565b80601f01602080910402602001604051908101604052809291908181526020018280546105d79061131a565b80156106225780601f106105f957610100808354040283529160200191610622565b820191905f5260205f20905b81548152906001019060200180831161060557829003601f168201915b5050505050905090565b5f5f61063661087a565b9050610643818585610926565b600191505092915050565b66038d7ea4c6800081565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6106e3610a16565b5f4790505f8111610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072090611394565b60405180910390fd5b610731610574565b73ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610773573d5f5f3e3d5ffd5b5050565b61077f610a16565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107ef575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016107e691906111c0565b60405180910390fd5b6107f881610a9d565b50565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361086b575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161086291906111c0565b60405180910390fd5b6108765f8383610b60565b5050565b5f33905090565b61088e8383836001610d79565b505050565b5f61089e8484610659565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156109205781811015610911578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610908939291906113b2565b60405180910390fd5b61091f84848484035f610d79565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610996575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161098d91906111c0565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a06575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016109fd91906111c0565b60405180910390fd5b610a11838383610b60565b505050565b610a1e61087a565b73ffffffffffffffffffffffffffffffffffffffff16610a3c610574565b73ffffffffffffffffffffffffffffffffffffffff1614610a9b57610a5f61087a565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610a9291906111c0565b60405180910390fd5b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bb0578060025f828254610ba491906113e7565b92505081905550610c7e565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610c39578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610c30939291906113b2565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610cc5578060025f8282540392505081905550610d0f565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d6c91906110e9565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610de9575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610de091906111c0565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610e59575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610e5091906111c0565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610f42578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610f3991906110e9565b60405180910390a35b50505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610f8a82610f48565b610f948185610f52565b9350610fa4818560208601610f62565b610fad81610f70565b840191505092915050565b5f6020820190508181035f830152610fd08184610f80565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61100582610fdc565b9050919050565b61101581610ffb565b811461101f575f5ffd5b50565b5f813590506110308161100c565b92915050565b5f819050919050565b61104881611036565b8114611052575f5ffd5b50565b5f813590506110638161103f565b92915050565b5f5f6040838503121561107f5761107e610fd8565b5b5f61108c85828601611022565b925050602061109d85828601611055565b9150509250929050565b5f8115159050919050565b6110bb816110a7565b82525050565b5f6020820190506110d45f8301846110b2565b92915050565b6110e381611036565b82525050565b5f6020820190506110fc5f8301846110da565b92915050565b5f5f5f6060848603121561111957611118610fd8565b5b5f61112686828701611022565b935050602061113786828701611022565b925050604061114886828701611055565b9150509250925092565b5f60ff82169050919050565b61116781611152565b82525050565b5f6020820190506111805f83018461115e565b92915050565b5f6020828403121561119b5761119a610fd8565b5b5f6111a884828501611022565b91505092915050565b6111ba81610ffb565b82525050565b5f6020820190506111d35f8301846111b1565b92915050565b5f5f604083850312156111ef576111ee610fd8565b5b5f6111fc85828601611022565b925050602061120d85828601611022565b9150509250929050565b7f4e6f7420656e6f756768204554482073656e74000000000000000000000000005f82015250565b5f61124b601383610f52565b915061125682611217565b602082019050919050565b5f6020820190508181035f8301526112788161123f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6112b682611036565b91506112c183611036565b92508282026112cf81611036565b915082820484148315176112e6576112e561127f565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061133157607f821691505b602082108103611344576113436112ed565b5b50919050565b7f4e6f2045544820746f20776974686472617700000000000000000000000000005f82015250565b5f61137e601283610f52565b91506113898261134a565b602082019050919050565b5f6020820190508181035f8301526113ab81611372565b9050919050565b5f6060820190506113c55f8301866111b1565b6113d260208301856110da565b6113df60408301846110da565b949350505050565b5f6113f182611036565b91506113fc83611036565b92508282019050808211156114145761141361127f565b5b9291505056fea264697066735822122068335fa0795a6465d9f5d88ff9647d13e94ca6d6b4326d7afeb90da0bf52817364736f6c634300081e0033",
}

// MyTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MyTokenMetaData.ABI instead.
var MyTokenABI = MyTokenMetaData.ABI

// MyTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyTokenMetaData.Bin instead.
var MyTokenBin = MyTokenMetaData.Bin

// DeployMyToken deploys a new Ethereum contract, binding an instance of MyToken to it.
func DeployMyToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *MyToken, error) {
	parsed, err := MyTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyTokenBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// MyToken is an auto generated Go binding around an Ethereum contract.
type MyToken struct {
	MyTokenCaller     // Read-only binding to the contract
	MyTokenTransactor // Write-only binding to the contract
	MyTokenFilterer   // Log filterer for contract events
}

// MyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyTokenSession struct {
	Contract     *MyToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyTokenCallerSession struct {
	Contract *MyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyTokenTransactorSession struct {
	Contract     *MyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyTokenRaw struct {
	Contract *MyToken // Generic contract binding to access the raw methods on
}

// MyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyTokenCallerRaw struct {
	Contract *MyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyTokenTransactorRaw struct {
	Contract *MyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyToken creates a new instance of MyToken, bound to a specific deployed contract.
func NewMyToken(address common.Address, backend bind.ContractBackend) (*MyToken, error) {
	contract, err := bindMyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenCaller(address common.Address, caller bind.ContractCaller) (*MyTokenCaller, error) {
	contract, err := bindMyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenCaller{contract: contract}, nil
}

// NewMyTokenTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyTokenTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransactor{contract: contract}, nil
}

// NewMyTokenFilterer creates a new log filterer instance of MyToken, bound to a specific deployed contract.
func NewMyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyTokenFilterer, error) {
	contract, err := bindMyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyTokenFilterer{contract: contract}, nil
}

// bindMyToken binds a generic wrapper to an already deployed contract.
func bindMyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.MyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transact(opts, method, params...)
}

// MINETH is a free data retrieval call binding the contract method 0xd586b74c.
//
// Solidity: function MIN_ETH() view returns(uint256)
func (_MyToken *MyTokenCaller) MINETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "MIN_ETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINETH is a free data retrieval call binding the contract method 0xd586b74c.
//
// Solidity: function MIN_ETH() view returns(uint256)
func (_MyToken *MyTokenSession) MINETH() (*big.Int, error) {
	return _MyToken.Contract.MINETH(&_MyToken.CallOpts)
}

// MINETH is a free data retrieval call binding the contract method 0xd586b74c.
//
// Solidity: function MIN_ETH() view returns(uint256)
func (_MyToken *MyTokenCallerSession) MINETH() (*big.Int, error) {
	return _MyToken.Contract.MINETH(&_MyToken.CallOpts)
}

// RATE is a free data retrieval call binding the contract method 0x664e9704.
//
// Solidity: function RATE() view returns(uint256)
func (_MyToken *MyTokenCaller) RATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RATE is a free data retrieval call binding the contract method 0x664e9704.
//
// Solidity: function RATE() view returns(uint256)
func (_MyToken *MyTokenSession) RATE() (*big.Int, error) {
	return _MyToken.Contract.RATE(&_MyToken.CallOpts)
}

// RATE is a free data retrieval call binding the contract method 0x664e9704.
//
// Solidity: function RATE() view returns(uint256)
func (_MyToken *MyTokenCallerSession) RATE() (*big.Int, error) {
	return _MyToken.Contract.RATE(&_MyToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyToken *MyTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyToken *MyTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyToken.Contract.Allowance(&_MyToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyToken *MyTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyToken.Contract.Allowance(&_MyToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyToken *MyTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyToken *MyTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyToken *MyTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenCallerSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenCallerSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyToken *MyTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyToken *MyTokenSession) Owner() (common.Address, error) {
	return _MyToken.Contract.Owner(&_MyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyToken *MyTokenCallerSession) Owner() (common.Address, error) {
	return _MyToken.Contract.Owner(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenCallerSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenSession) TotalSupply() (*big.Int, error) {
	return _MyToken.Contract.TotalSupply(&_MyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MyToken.Contract.TotalSupply(&_MyToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyToken *MyTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyToken *MyTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.Approve(&_MyToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyToken *MyTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.Approve(&_MyToken.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_MyToken *MyTokenTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_MyToken *MyTokenSession) Mint() (*types.Transaction, error) {
	return _MyToken.Contract.Mint(&_MyToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_MyToken *MyTokenTransactorSession) Mint() (*types.Transaction, error) {
	return _MyToken.Contract.Mint(&_MyToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyToken *MyTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyToken *MyTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyToken.Contract.RenounceOwnership(&_MyToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyToken *MyTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyToken.Contract.RenounceOwnership(&_MyToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyToken *MyTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyToken *MyTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyToken *MyTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyToken *MyTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyToken *MyTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyToken *MyTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyToken *MyTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyToken *MyTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyToken.Contract.TransferOwnership(&_MyToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyToken *MyTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyToken.Contract.TransferOwnership(&_MyToken.TransactOpts, newOwner)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_MyToken *MyTokenTransactor) WithdrawETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "withdrawETH")
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_MyToken *MyTokenSession) WithdrawETH() (*types.Transaction, error) {
	return _MyToken.Contract.WithdrawETH(&_MyToken.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_MyToken *MyTokenTransactorSession) WithdrawETH() (*types.Transaction, error) {
	return _MyToken.Contract.WithdrawETH(&_MyToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyToken *MyTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyToken *MyTokenSession) Receive() (*types.Transaction, error) {
	return _MyToken.Contract.Receive(&_MyToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyToken *MyTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _MyToken.Contract.Receive(&_MyToken.TransactOpts)
}

// MyTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyToken contract.
type MyTokenApprovalIterator struct {
	Event *MyTokenApproval // Event containing the contract specifics and raw log

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
func (it *MyTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenApproval)
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
		it.Event = new(MyTokenApproval)
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
func (it *MyTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenApproval represents a Approval event raised by the MyToken contract.
type MyTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MyToken *MyTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenApprovalIterator{contract: _MyToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MyToken *MyTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenApproval)
				if err := _MyToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MyToken *MyTokenFilterer) ParseApproval(log types.Log) (*MyTokenApproval, error) {
	event := new(MyTokenApproval)
	if err := _MyToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyToken contract.
type MyTokenOwnershipTransferredIterator struct {
	Event *MyTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MyTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenOwnershipTransferred)
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
		it.Event = new(MyTokenOwnershipTransferred)
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
func (it *MyTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MyToken contract.
type MyTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyToken *MyTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenOwnershipTransferredIterator{contract: _MyToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyToken *MyTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenOwnershipTransferred)
				if err := _MyToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyToken *MyTokenFilterer) ParseOwnershipTransferred(log types.Log) (*MyTokenOwnershipTransferred, error) {
	event := new(MyTokenOwnershipTransferred)
	if err := _MyToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyToken contract.
type MyTokenTransferIterator struct {
	Event *MyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenTransfer)
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
		it.Event = new(MyTokenTransfer)
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
func (it *MyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenTransfer represents a Transfer event raised by the MyToken contract.
type MyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyToken *MyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransferIterator{contract: _MyToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyToken *MyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenTransfer)
				if err := _MyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MyToken *MyTokenFilterer) ParseTransfer(log types.Log) (*MyTokenTransfer, error) {
	event := new(MyTokenTransfer)
	if err := _MyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
