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

// ERC20TokenOwnableMetaData contains all meta data concerning the ERC20TokenOwnable contract.
var ERC20TokenOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161203f38038061203f83398181016040528101906100319190610748565b808585816003908161004391906109fe565b50806004908161005391906109fe565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100c6575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100bd9190610adc565b60405180910390fd5b6100d5816101d760201b60201c565b505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610144576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013b90610b75565b60405180910390fd5b5f8360ff1611801561015a575060128360ff1611155b610199576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019090610c03565b60405180910390fd5b82600560146101000a81548160ff021916908360ff1602179055505f8211156101cd576101cc818361029a60201b60201c565b5b5050505050610cde565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361030a575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016103019190610adc565b60405180910390fd5b61031b5f838361031f60201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361036f578060025f8282546103639190610c4e565b9250508190555061043d565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156103f8578381836040517fe450d38c0000000000000000000000000000000000000000000000000000000081526004016103ef93929190610c90565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610484578060025f82825403925050819055506104ce565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161052b9190610cc5565b60405180910390a3505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61059782610551565b810181811067ffffffffffffffff821117156105b6576105b5610561565b5b80604052505050565b5f6105c8610538565b90506105d4828261058e565b919050565b5f67ffffffffffffffff8211156105f3576105f2610561565b5b6105fc82610551565b9050602081019050919050565b8281835e5f83830152505050565b5f610629610624846105d9565b6105bf565b9050828152602081018484840111156106455761064461054d565b5b610650848285610609565b509392505050565b5f82601f83011261066c5761066b610549565b5b815161067c848260208601610617565b91505092915050565b5f60ff82169050919050565b61069a81610685565b81146106a4575f5ffd5b50565b5f815190506106b581610691565b92915050565b5f819050919050565b6106cd816106bb565b81146106d7575f5ffd5b50565b5f815190506106e8816106c4565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610717826106ee565b9050919050565b6107278161070d565b8114610731575f5ffd5b50565b5f815190506107428161071e565b92915050565b5f5f5f5f5f60a0868803121561076157610760610541565b5b5f86015167ffffffffffffffff81111561077e5761077d610545565b5b61078a88828901610658565b955050602086015167ffffffffffffffff8111156107ab576107aa610545565b5b6107b788828901610658565b94505060406107c8888289016106a7565b93505060606107d9888289016106da565b92505060806107ea88828901610734565b9150509295509295909350565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061084557607f821691505b60208210810361085857610857610801565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108ba7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261087f565b6108c4868361087f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6108ff6108fa6108f5846106bb565b6108dc565b6106bb565b9050919050565b5f819050919050565b610918836108e5565b61092c61092482610906565b84845461088b565b825550505050565b5f5f905090565b610943610934565b61094e81848461090f565b505050565b5b81811015610971576109665f8261093b565b600181019050610954565b5050565b601f8211156109b6576109878161085e565b61099084610870565b8101602085101561099f578190505b6109b36109ab85610870565b830182610953565b50505b505050565b5f82821c905092915050565b5f6109d65f19846008026109bb565b1980831691505092915050565b5f6109ee83836109c7565b9150826002028217905092915050565b610a07826107f7565b67ffffffffffffffff811115610a2057610a1f610561565b5b610a2a825461082e565b610a35828285610975565b5f60209050601f831160018114610a66575f8415610a54578287015190505b610a5e85826109e3565b865550610ac5565b601f198416610a748661085e565b5f5b82811015610a9b57848901518255600182019150602085019450602081019050610a76565b86831015610ab85784890151610ab4601f8916826109c7565b8355505b6001600288020188555050505b505050505050565b610ad68161070d565b82525050565b5f602082019050610aef5f830184610acd565b92915050565b5f82825260208201905092915050565b7f4552433230546f6b656e4f776e61626c653a206f776e65722063616e6e6f74205f8201527f6265207a65726f20616464726573730000000000000000000000000000000000602082015250565b5f610b5f602f83610af5565b9150610b6a82610b05565b604082019050919050565b5f6020820190508181035f830152610b8c81610b53565b9050919050565b7f4552433230546f6b656e4f776e61626c653a20646563696d616c73206d7573745f8201527f206265206265747765656e203120616e64203138000000000000000000000000602082015250565b5f610bed603483610af5565b9150610bf882610b93565b604082019050919050565b5f6020820190508181035f830152610c1a81610be1565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c58826106bb565b9150610c63836106bb565b9250828201905080821115610c7b57610c7a610c21565b5b92915050565b610c8a816106bb565b82525050565b5f606082019050610ca35f830186610acd565b610cb06020830185610c81565b610cbd6040830184610c81565b949350505050565b5f602082019050610cd85f830184610c81565b92915050565b61135480610ceb5f395ff3fe608060405234801561000f575f5ffd5b50600436106100f3575f3560e01c806370a082311161009557806395d89b411161006457806395d89b411461025d578063a9059cbb1461027b578063dd62ed3e146102ab578063f2fde38b146102db576100f3565b806370a08231146101e9578063715018a61461021957806379cc6790146102235780638da5cb5b1461023f576100f3565b806323b872dd116100d157806323b872dd14610163578063313ce5671461019357806340c10f19146101b157806342966c68146101cd576100f3565b806306fdde03146100f7578063095ea7b31461011557806318160ddd14610145575b5f5ffd5b6100ff6102f7565b60405161010c9190610f14565b60405180910390f35b61012f600480360381019061012a9190610fc5565b610387565b60405161013c919061101d565b60405180910390f35b61014d6103a9565b60405161015a9190611045565b60405180910390f35b61017d6004803603810190610178919061105e565b6103b2565b60405161018a919061101d565b60405180910390f35b61019b6103e0565b6040516101a891906110c9565b60405180910390f35b6101cb60048036038101906101c69190610fc5565b6103f6565b005b6101e760048036038101906101e291906110e2565b61047a565b005b61020360048036038101906101fe919061110d565b610487565b6040516102109190611045565b60405180910390f35b6102216104cc565b005b61023d60048036038101906102389190610fc5565b6104df565b005b6102476104f8565b6040516102549190611147565b60405180910390f35b610265610520565b6040516102729190610f14565b60405180910390f35b61029560048036038101906102909190610fc5565b6105b0565b6040516102a2919061101d565b60405180910390f35b6102c560048036038101906102c09190611160565b6105d2565b6040516102d29190611045565b60405180910390f35b6102f560048036038101906102f0919061110d565b610654565b005b606060038054610306906111cb565b80601f0160208091040260200160405190810160405280929190818152602001828054610332906111cb565b801561037d5780601f106103545761010080835404028352916020019161037d565b820191905f5260205f20905b81548152906001019060200180831161036057829003601f168201915b5050505050905090565b5f5f6103916106d8565b905061039e8185856106df565b600191505092915050565b5f600254905090565b5f5f6103bc6106d8565b90506103c98582856106f1565b6103d4858585610784565b60019150509392505050565b5f600560149054906101000a900460ff16905090565b6103fe610874565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361046c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104639061126b565b60405180910390fd5b61047682826108fb565b5050565b610484338261097a565b50565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104d4610874565b6104dd5f6109f9565b565b6104ea8233836106f1565b6104f4828261097a565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461052f906111cb565b80601f016020809104026020016040519081016040528092919081815260200182805461055b906111cb565b80156105a65780601f1061057d576101008083540402835291602001916105a6565b820191905f5260205f20905b81548152906001019060200180831161058957829003601f168201915b5050505050905090565b5f5f6105ba6106d8565b90506105c7818585610784565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61065c610874565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106cc575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106c39190611147565b60405180910390fd5b6106d5816109f9565b50565b5f33905090565b6106ec8383836001610abc565b505050565b5f6106fc84846105d2565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561077e578181101561076f578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161076693929190611289565b60405180910390fd5b61077d84848484035f610abc565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107f4575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107eb9190611147565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610864575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161085b9190611147565b60405180910390fd5b61086f838383610c8b565b505050565b61087c6106d8565b73ffffffffffffffffffffffffffffffffffffffff1661089a6104f8565b73ffffffffffffffffffffffffffffffffffffffff16146108f9576108bd6106d8565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016108f09190611147565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361096b575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016109629190611147565b60405180910390fd5b6109765f8383610c8b565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109ea575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016109e19190611147565b60405180910390fd5b6109f5825f83610c8b565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b2c575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610b239190611147565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b9c575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610b939190611147565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610c85578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610c7c9190611045565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610cdb578060025f828254610ccf91906112eb565b92505081905550610da9565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610d64578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610d5b93929190611289565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610df0578060025f8282540392505081905550610e3a565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e979190611045565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610ee682610ea4565b610ef08185610eae565b9350610f00818560208601610ebe565b610f0981610ecc565b840191505092915050565b5f6020820190508181035f830152610f2c8184610edc565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f6182610f38565b9050919050565b610f7181610f57565b8114610f7b575f5ffd5b50565b5f81359050610f8c81610f68565b92915050565b5f819050919050565b610fa481610f92565b8114610fae575f5ffd5b50565b5f81359050610fbf81610f9b565b92915050565b5f5f60408385031215610fdb57610fda610f34565b5b5f610fe885828601610f7e565b9250506020610ff985828601610fb1565b9150509250929050565b5f8115159050919050565b61101781611003565b82525050565b5f6020820190506110305f83018461100e565b92915050565b61103f81610f92565b82525050565b5f6020820190506110585f830184611036565b92915050565b5f5f5f6060848603121561107557611074610f34565b5b5f61108286828701610f7e565b935050602061109386828701610f7e565b92505060406110a486828701610fb1565b9150509250925092565b5f60ff82169050919050565b6110c3816110ae565b82525050565b5f6020820190506110dc5f8301846110ba565b92915050565b5f602082840312156110f7576110f6610f34565b5b5f61110484828501610fb1565b91505092915050565b5f6020828403121561112257611121610f34565b5b5f61112f84828501610f7e565b91505092915050565b61114181610f57565b82525050565b5f60208201905061115a5f830184611138565b92915050565b5f5f6040838503121561117657611175610f34565b5b5f61118385828601610f7e565b925050602061119485828601610f7e565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111e257607f821691505b6020821081036111f5576111f461119e565b5b50919050565b7f4552433230546f6b656e4f776e61626c653a2063616e6e6f74206d696e7420745f8201527f6f207a65726f2061646472657373000000000000000000000000000000000000602082015250565b5f611255602e83610eae565b9150611260826111fb565b604082019050919050565b5f6020820190508181035f83015261128281611249565b9050919050565b5f60608201905061129c5f830186611138565b6112a96020830185611036565b6112b66040830184611036565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6112f582610f92565b915061130083610f92565b9250828201905080821115611318576113176112be565b5b9291505056fea264697066735822122091b8c21ca4ea3902e842d31019814c3ea3ba0a5457ed21ad4604152857ae70b164736f6c634300081f0033",
}

// ERC20TokenOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenOwnableMetaData.ABI instead.
var ERC20TokenOwnableABI = ERC20TokenOwnableMetaData.ABI

// ERC20TokenOwnableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenOwnableMetaData.Bin instead.
var ERC20TokenOwnableBin = ERC20TokenOwnableMetaData.Bin

// DeployERC20TokenOwnable deploys a new Ethereum contract, binding an instance of ERC20TokenOwnable to it.
func DeployERC20TokenOwnable(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, initialSupply_ *big.Int, owner_ common.Address) (common.Address, *types.Transaction, *ERC20TokenOwnable, error) {
	parsed, err := ERC20TokenOwnableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenOwnableBin), backend, name_, symbol_, decimals_, initialSupply_, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenOwnable{ERC20TokenOwnableCaller: ERC20TokenOwnableCaller{contract: contract}, ERC20TokenOwnableTransactor: ERC20TokenOwnableTransactor{contract: contract}, ERC20TokenOwnableFilterer: ERC20TokenOwnableFilterer{contract: contract}}, nil
}

// ERC20TokenOwnable is an auto generated Go binding around an Ethereum contract.
type ERC20TokenOwnable struct {
	ERC20TokenOwnableCaller     // Read-only binding to the contract
	ERC20TokenOwnableTransactor // Write-only binding to the contract
	ERC20TokenOwnableFilterer   // Log filterer for contract events
}

// ERC20TokenOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenOwnableSession struct {
	Contract     *ERC20TokenOwnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC20TokenOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenOwnableCallerSession struct {
	Contract *ERC20TokenOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC20TokenOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenOwnableTransactorSession struct {
	Contract     *ERC20TokenOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20TokenOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenOwnableRaw struct {
	Contract *ERC20TokenOwnable // Generic contract binding to access the raw methods on
}

// ERC20TokenOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenOwnableCallerRaw struct {
	Contract *ERC20TokenOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenOwnableTransactorRaw struct {
	Contract *ERC20TokenOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenOwnable creates a new instance of ERC20TokenOwnable, bound to a specific deployed contract.
func NewERC20TokenOwnable(address common.Address, backend bind.ContractBackend) (*ERC20TokenOwnable, error) {
	contract, err := bindERC20TokenOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnable{ERC20TokenOwnableCaller: ERC20TokenOwnableCaller{contract: contract}, ERC20TokenOwnableTransactor: ERC20TokenOwnableTransactor{contract: contract}, ERC20TokenOwnableFilterer: ERC20TokenOwnableFilterer{contract: contract}}, nil
}

// NewERC20TokenOwnableCaller creates a new read-only instance of ERC20TokenOwnable, bound to a specific deployed contract.
func NewERC20TokenOwnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenOwnableCaller, error) {
	contract, err := bindERC20TokenOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnableCaller{contract: contract}, nil
}

// NewERC20TokenOwnableTransactor creates a new write-only instance of ERC20TokenOwnable, bound to a specific deployed contract.
func NewERC20TokenOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenOwnableTransactor, error) {
	contract, err := bindERC20TokenOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnableTransactor{contract: contract}, nil
}

// NewERC20TokenOwnableFilterer creates a new log filterer instance of ERC20TokenOwnable, bound to a specific deployed contract.
func NewERC20TokenOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenOwnableFilterer, error) {
	contract, err := bindERC20TokenOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnableFilterer{contract: contract}, nil
}

// bindERC20TokenOwnable binds a generic wrapper to an already deployed contract.
func bindERC20TokenOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenOwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenOwnable *ERC20TokenOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenOwnable.Contract.ERC20TokenOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenOwnable *ERC20TokenOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.ERC20TokenOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenOwnable *ERC20TokenOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.ERC20TokenOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenOwnable.Contract.Allowance(&_ERC20TokenOwnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenOwnable.Contract.Allowance(&_ERC20TokenOwnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenOwnable.Contract.BalanceOf(&_ERC20TokenOwnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenOwnable.Contract.BalanceOf(&_ERC20TokenOwnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Decimals() (uint8, error) {
	return _ERC20TokenOwnable.Contract.Decimals(&_ERC20TokenOwnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) Decimals() (uint8, error) {
	return _ERC20TokenOwnable.Contract.Decimals(&_ERC20TokenOwnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Name() (string, error) {
	return _ERC20TokenOwnable.Contract.Name(&_ERC20TokenOwnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) Name() (string, error) {
	return _ERC20TokenOwnable.Contract.Name(&_ERC20TokenOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Owner() (common.Address, error) {
	return _ERC20TokenOwnable.Contract.Owner(&_ERC20TokenOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) Owner() (common.Address, error) {
	return _ERC20TokenOwnable.Contract.Owner(&_ERC20TokenOwnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Symbol() (string, error) {
	return _ERC20TokenOwnable.Contract.Symbol(&_ERC20TokenOwnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) Symbol() (string, error) {
	return _ERC20TokenOwnable.Contract.Symbol(&_ERC20TokenOwnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenOwnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenOwnable.Contract.TotalSupply(&_ERC20TokenOwnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenOwnable *ERC20TokenOwnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenOwnable.Contract.TotalSupply(&_ERC20TokenOwnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Approve(&_ERC20TokenOwnable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Approve(&_ERC20TokenOwnable.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Burn(&_ERC20TokenOwnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Burn(&_ERC20TokenOwnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.BurnFrom(&_ERC20TokenOwnable.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.BurnFrom(&_ERC20TokenOwnable.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Mint(&_ERC20TokenOwnable.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Mint(&_ERC20TokenOwnable.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.RenounceOwnership(&_ERC20TokenOwnable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.RenounceOwnership(&_ERC20TokenOwnable.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Transfer(&_ERC20TokenOwnable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.Transfer(&_ERC20TokenOwnable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.TransferFrom(&_ERC20TokenOwnable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.TransferFrom(&_ERC20TokenOwnable.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenOwnable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.TransferOwnership(&_ERC20TokenOwnable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenOwnable *ERC20TokenOwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenOwnable.Contract.TransferOwnership(&_ERC20TokenOwnable.TransactOpts, newOwner)
}

// ERC20TokenOwnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20TokenOwnable contract.
type ERC20TokenOwnableApprovalIterator struct {
	Event *ERC20TokenOwnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenOwnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenOwnableApproval)
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
		it.Event = new(ERC20TokenOwnableApproval)
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
func (it *ERC20TokenOwnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenOwnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenOwnableApproval represents a Approval event raised by the ERC20TokenOwnable contract.
type ERC20TokenOwnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenOwnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenOwnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnableApprovalIterator{contract: _ERC20TokenOwnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenOwnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenOwnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenOwnableApproval)
				if err := _ERC20TokenOwnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) ParseApproval(log types.Log) (*ERC20TokenOwnableApproval, error) {
	event := new(ERC20TokenOwnableApproval)
	if err := _ERC20TokenOwnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenOwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TokenOwnable contract.
type ERC20TokenOwnableOwnershipTransferredIterator struct {
	Event *ERC20TokenOwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20TokenOwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenOwnableOwnershipTransferred)
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
		it.Event = new(ERC20TokenOwnableOwnershipTransferred)
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
func (it *ERC20TokenOwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenOwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenOwnableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TokenOwnable contract.
type ERC20TokenOwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TokenOwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenOwnable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnableOwnershipTransferredIterator{contract: _ERC20TokenOwnable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TokenOwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenOwnable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenOwnableOwnershipTransferred)
				if err := _ERC20TokenOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TokenOwnableOwnershipTransferred, error) {
	event := new(ERC20TokenOwnableOwnershipTransferred)
	if err := _ERC20TokenOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenOwnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20TokenOwnable contract.
type ERC20TokenOwnableTransferIterator struct {
	Event *ERC20TokenOwnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenOwnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenOwnableTransfer)
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
		it.Event = new(ERC20TokenOwnableTransfer)
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
func (it *ERC20TokenOwnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenOwnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenOwnableTransfer represents a Transfer event raised by the ERC20TokenOwnable contract.
type ERC20TokenOwnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenOwnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenOwnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenOwnableTransferIterator{contract: _ERC20TokenOwnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenOwnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenOwnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenOwnableTransfer)
				if err := _ERC20TokenOwnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20TokenOwnable *ERC20TokenOwnableFilterer) ParseTransfer(log types.Log) (*ERC20TokenOwnableTransfer, error) {
	event := new(ERC20TokenOwnableTransfer)
	if err := _ERC20TokenOwnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
