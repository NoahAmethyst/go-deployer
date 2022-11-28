// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cat\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"}],\"name\":\"batchFreeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_basePrice\",\"type\":\"uint256\"}],\"name\":\"changeBasePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"changeFreeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cat\",\"type\":\"address\"}],\"name\":\"changeIcat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"changePublicThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"freeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"icat\",\"outputs\":[{\"internalType\":\"contractICAT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whiteListMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x60806040526706f05b59d3b200006003556101f46005556107d06006553480156200002957600080fd5b506040516200252a3803806200252a83398181016040528101906200004f919062000198565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050620001ca565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001608262000133565b9050919050565b620001728162000153565b81146200017e57600080fd5b50565b600081519050620001928162000167565b92915050565b600060208284031215620001b157620001b06200012e565b5b6000620001c18482850162000181565b91505092915050565b61235080620001da6000396000f3fe60806040526004361061011f5760003560e01c80639870d7fe116100a0578063c6cfe19211610064578063c6cfe1921461038a578063c7876ea4146103b3578063db5e5c3d146103de578063dc53fd9214610407578063fa07ce1d146104325761011f565b80639870d7fe146102bb57806398eef4b7146102e4578063ac8a584a1461030d578063ae47b27c14610336578063bdf884c3146103615761011f565b80633ed50bb3116100e75780633ed50bb3146101e85780634fb2e45d146102255780635dc89b341461024e5780636196c279146102655780638da5cb5b146102905761011f565b806313e7c9d81461012457806325f051bd1461016157806326092b831461018c57806330df47af14610196578063370fc62d146101bf575b600080fd5b34801561013057600080fd5b5061014b6004803603810190610146919061181f565b61045b565b6040516101589190611867565b60405180910390f35b34801561016d57600080fd5b5061017661047b565b604051610183919061189b565b60405180910390f35b610194610481565b005b3480156101a257600080fd5b506101bd60048036038101906101b891906118f4565b6106bc565b005b3480156101cb57600080fd5b506101e660048036038101906101e1919061181f565b61079e565b005b3480156101f457600080fd5b5061020f600480360381019061020a919061181f565b610870565b60405161021c9190611867565b60405180910390f35b34801561023157600080fd5b5061024c6004803603810190610247919061181f565b610890565b005b34801561025a57600080fd5b50610263610961565b005b34801561027157600080fd5b5061027a610c38565b6040516102879190611980565b60405180910390f35b34801561029c57600080fd5b506102a5610c5e565b6040516102b291906119aa565b60405180910390f35b3480156102c757600080fd5b506102e260048036038101906102dd919061181f565b610c82565b005b3480156102f057600080fd5b5061030b60048036038101906103069190611b1e565b610d6a565b005b34801561031957600080fd5b50610334600480360381019061032f919061181f565b610fb9565b005b34801561034257600080fd5b5061034b6110a2565b604051610358919061189b565b60405180910390f35b34801561036d57600080fd5b5061038860048036038101906103839190611b93565b6110a8565b005b34801561039657600080fd5b506103b160048036038101906103ac9190611b93565b611140565b005b3480156103bf57600080fd5b506103c86111d8565b6040516103d5919061189b565b60405180910390f35b3480156103ea57600080fd5b5061040560048036038101906104009190611b93565b6111de565b005b34801561041357600080fd5b5061041c611276565b604051610429919061189b565b60405180910390f35b34801561043e57600080fd5b506104596004803603810190610454919061181f565b611378565b005b60016020528060005260406000206000915054906101000a900460ff1681565b60055481565b600654600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105159190611bd5565b1015610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054d90611c5f565b60405180910390fd5b6000610560611276565b9050803410156105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059c90611ccb565b60405180910390fd5b60006105ba823461157a90919063ffffffff16565b111561061a573373ffffffffffffffffffffffffffffffffffffffff166108fc6105ed833461157a90919063ffffffff16565b9081150290604051600060405180830381858888f19350505050158015610618573d6000803e3d6000fd5b505b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a627842336040518263ffffffff1660e01b815260040161067591906119aa565b6020604051808303816000875af1158015610694573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b89190611bd5565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461074a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074190611d37565b60405180910390fd5b600047111561079b578073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610799573d6000803e3d6000fd5b505b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461082c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082390611d37565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60046020528060005260406000206000915054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461091e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091590611d37565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600554600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f59190611bd5565b1015610a36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2d90611dc9565b60405180910390fd5b600654600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aca9190611bd5565b10610b0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0190611e35565b60405180910390fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610b97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8e90611ea1565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a627842336040518263ffffffff1660e01b8152600401610bf291906119aa565b6020604051808303816000875af1158015610c11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c359190611bd5565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0790611d37565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610df6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ded90611f0d565b60405180910390fd5b600554610e9d8251600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8f9190611bd5565b6115c490919063ffffffff16565b10610edd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed490611f79565b60405180910390fd5b60005b8151811015610fb557600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a627842838381518110610f3a57610f39611f99565b5b60200260200101516040518263ffffffff1660e01b8152600401610f5e91906119aa565b6020604051808303816000875af1158015610f7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa19190611bd5565b508080610fad90611ff7565b915050610ee0565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611047576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103e90611d37565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60065481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611136576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112d90611d37565b60405180910390fd5b8060068190555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c590611d37565b60405180910390fd5b8060038190555050565b60035481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461126c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161126390611d37565b60405180910390fd5b8060058190555050565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130a9190611bd5565b905060065481101561132157600354915050611375565b61137161136060016113526103e86113446006548761157a90919063ffffffff16565b61162290919063ffffffff16565b6115c490919063ffffffff16565b60035461166c90919063ffffffff16565b9150505b90565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611404576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fb90611f0d565b60405180910390fd5b600554600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611474573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114989190611bd5565b106114d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114cf90611f79565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a627842826040518263ffffffff1660e01b815260040161153391906119aa565b6020604051808303816000875af1158015611552573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115769190611bd5565b5050565b60006115bc83836040518060400160405280601481526020017f7375627472616374696f6e206f766572666c6f770000000000000000000000008152506116e6565b905092915050565b60008082846115d3919061203f565b905083811015611618576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160f906120e1565b60405180910390fd5b8091505092915050565b600061166483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061174a565b905092915050565b600080830361167e57600090506116e0565b6000828461168c9190612101565b905082848261169b919061218a565b146116db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d29061222d565b60405180910390fd5b809150505b92915050565b600083831115829061172e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161172591906122c4565b60405180910390fd5b506000838561173d91906122e6565b9050809150509392505050565b60008083118290611791576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161178891906122c4565b60405180910390fd5b50600083856117a0919061218a565b9050809150509392505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006117ec826117c1565b9050919050565b6117fc816117e1565b811461180757600080fd5b50565b600081359050611819816117f3565b92915050565b600060208284031215611835576118346117b7565b5b60006118438482850161180a565b91505092915050565b60008115159050919050565b6118618161184c565b82525050565b600060208201905061187c6000830184611858565b92915050565b6000819050919050565b61189581611882565b82525050565b60006020820190506118b0600083018461188c565b92915050565b60006118c1826117c1565b9050919050565b6118d1816118b6565b81146118dc57600080fd5b50565b6000813590506118ee816118c8565b92915050565b60006020828403121561190a576119096117b7565b5b6000611918848285016118df565b91505092915050565b6000819050919050565b600061194661194161193c846117c1565b611921565b6117c1565b9050919050565b60006119588261192b565b9050919050565b600061196a8261194d565b9050919050565b61197a8161195f565b82525050565b60006020820190506119956000830184611971565b92915050565b6119a4816117e1565b82525050565b60006020820190506119bf600083018461199b565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611a13826119ca565b810181811067ffffffffffffffff82111715611a3257611a316119db565b5b80604052505050565b6000611a456117ad565b9050611a518282611a0a565b919050565b600067ffffffffffffffff821115611a7157611a706119db565b5b602082029050602081019050919050565b600080fd5b6000611a9a611a9584611a56565b611a3b565b90508083825260208201905060208402830185811115611abd57611abc611a82565b5b835b81811015611ae65780611ad2888261180a565b845260208401935050602081019050611abf565b5050509392505050565b600082601f830112611b0557611b046119c5565b5b8135611b15848260208601611a87565b91505092915050565b600060208284031215611b3457611b336117b7565b5b600082013567ffffffffffffffff811115611b5257611b516117bc565b5b611b5e84828501611af0565b91505092915050565b611b7081611882565b8114611b7b57600080fd5b50565b600081359050611b8d81611b67565b92915050565b600060208284031215611ba957611ba86117b7565b5b6000611bb784828501611b7e565b91505092915050565b600081519050611bcf81611b67565b92915050565b600060208284031215611beb57611bea6117b7565b5b6000611bf984828501611bc0565b91505092915050565b600082825260208201905092915050565b7f5075626c6963206d696e74206e6f74206f70656e207965740000000000000000600082015250565b6000611c49601883611c02565b9150611c5482611c13565b602082019050919050565b60006020820190508181036000830152611c7881611c3c565b9050919050565b7f496e73756666696369656e740000000000000000000000000000000000000000600082015250565b6000611cb5600c83611c02565b9150611cc082611c7f565b602082019050919050565b60006020820190508181036000830152611ce481611ca8565b9050919050565b7f4f6e6c79204f776e657200000000000000000000000000000000000000000000600082015250565b6000611d21600a83611c02565b9150611d2c82611ceb565b602082019050919050565b60006020820190508181036000830152611d5081611d14565b9050919050565b7f4e6f74207265616368207768697465206c697374206d696e742072657175697260008201527f656d656e74000000000000000000000000000000000000000000000000000000602082015250565b6000611db3602583611c02565b9150611dbe82611d57565b604082019050919050565b60006020820190508181036000830152611de281611da6565b9050919050565b7f5768697465206c697374206d696e7420636f756e742065786365656465640000600082015250565b6000611e1f601e83611c02565b9150611e2a82611de9565b602082019050919050565b60006020820190508181036000830152611e4e81611e12565b9050919050565b7f4f6e65207768697465206c6973742061646472657373206f6e6c79206f6e6365600082015250565b6000611e8b602083611c02565b9150611e9682611e55565b602082019050919050565b60006020820190508181036000830152611eba81611e7e565b9050919050565b7f4f6e6c79204f70657261746f7200000000000000000000000000000000000000600082015250565b6000611ef7600d83611c02565b9150611f0282611ec1565b602082019050919050565b60006020820190508181036000830152611f2681611eea565b9050919050565b7f46726565206d696e7420636f756e742065786365656465640000000000000000600082015250565b6000611f63601883611c02565b9150611f6e82611f2d565b602082019050919050565b60006020820190508181036000830152611f9281611f56565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061200282611882565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361203457612033611fc8565b5b600182019050919050565b600061204a82611882565b915061205583611882565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561208a57612089611fc8565b5b828201905092915050565b7f6164646974696f6e206f766572666c6f77000000000000000000000000000000600082015250565b60006120cb601183611c02565b91506120d682612095565b602082019050919050565b600060208201905081810360008301526120fa816120be565b9050919050565b600061210c82611882565b915061211783611882565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156121505761214f611fc8565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061219582611882565b91506121a083611882565b9250826121b0576121af61215b565b5b828204905092915050565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008201527f7700000000000000000000000000000000000000000000000000000000000000602082015250565b6000612217602183611c02565b9150612222826121bb565b604082019050919050565b600060208201905081810360008301526122468161220a565b9050919050565b600081519050919050565b60005b8381101561227657808201518184015260208101905061225b565b83811115612285576000848401525b50505050565b60006122968261224d565b6122a08185611c02565b93506122b0818560208601612258565b6122b9816119ca565b840191505092915050565b600060208201905081810360008301526122de818461228b565b905092915050565b60006122f182611882565b91506122fc83611882565b92508282101561230f5761230e611fc8565b5b82820390509291505056fea2646970667358221220a8dbd179d6d395c0647f460c82d0ea53c1a8c4c21b11d9a20f87a0dda379703164736f6c634300080d0033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _cat common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _cat)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// BasePrice is a free data retrieval call binding the contract method 0xc7876ea4.
//
// Solidity: function basePrice() view returns(uint256)
func (_Store *StoreCaller) BasePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "basePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePrice is a free data retrieval call binding the contract method 0xc7876ea4.
//
// Solidity: function basePrice() view returns(uint256)
func (_Store *StoreSession) BasePrice() (*big.Int, error) {
	return _Store.Contract.BasePrice(&_Store.CallOpts)
}

// BasePrice is a free data retrieval call binding the contract method 0xc7876ea4.
//
// Solidity: function basePrice() view returns(uint256)
func (_Store *StoreCallerSession) BasePrice() (*big.Int, error) {
	return _Store.Contract.BasePrice(&_Store.CallOpts)
}

// FreeThreshold is a free data retrieval call binding the contract method 0x25f051bd.
//
// Solidity: function freeThreshold() view returns(uint256)
func (_Store *StoreCaller) FreeThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "freeThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeThreshold is a free data retrieval call binding the contract method 0x25f051bd.
//
// Solidity: function freeThreshold() view returns(uint256)
func (_Store *StoreSession) FreeThreshold() (*big.Int, error) {
	return _Store.Contract.FreeThreshold(&_Store.CallOpts)
}

// FreeThreshold is a free data retrieval call binding the contract method 0x25f051bd.
//
// Solidity: function freeThreshold() view returns(uint256)
func (_Store *StoreCallerSession) FreeThreshold() (*big.Int, error) {
	return _Store.Contract.FreeThreshold(&_Store.CallOpts)
}

// Icat is a free data retrieval call binding the contract method 0x6196c279.
//
// Solidity: function icat() view returns(address)
func (_Store *StoreCaller) Icat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "icat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Icat is a free data retrieval call binding the contract method 0x6196c279.
//
// Solidity: function icat() view returns(address)
func (_Store *StoreSession) Icat() (common.Address, error) {
	return _Store.Contract.Icat(&_Store.CallOpts)
}

// Icat is a free data retrieval call binding the contract method 0x6196c279.
//
// Solidity: function icat() view returns(address)
func (_Store *StoreCallerSession) Icat() (common.Address, error) {
	return _Store.Contract.Icat(&_Store.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Store *StoreCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Store *StoreSession) Operators(arg0 common.Address) (bool, error) {
	return _Store.Contract.Operators(&_Store.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Store *StoreCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Store.Contract.Operators(&_Store.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_Store *StoreCaller) PublicMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "publicMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_Store *StoreSession) PublicMintPrice() (*big.Int, error) {
	return _Store.Contract.PublicMintPrice(&_Store.CallOpts)
}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_Store *StoreCallerSession) PublicMintPrice() (*big.Int, error) {
	return _Store.Contract.PublicMintPrice(&_Store.CallOpts)
}

// PublicThreshold is a free data retrieval call binding the contract method 0xae47b27c.
//
// Solidity: function publicThreshold() view returns(uint256)
func (_Store *StoreCaller) PublicThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "publicThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicThreshold is a free data retrieval call binding the contract method 0xae47b27c.
//
// Solidity: function publicThreshold() view returns(uint256)
func (_Store *StoreSession) PublicThreshold() (*big.Int, error) {
	return _Store.Contract.PublicThreshold(&_Store.CallOpts)
}

// PublicThreshold is a free data retrieval call binding the contract method 0xae47b27c.
//
// Solidity: function publicThreshold() view returns(uint256)
func (_Store *StoreCallerSession) PublicThreshold() (*big.Int, error) {
	return _Store.Contract.PublicThreshold(&_Store.CallOpts)
}

// WhiteListMinted is a free data retrieval call binding the contract method 0x3ed50bb3.
//
// Solidity: function whiteListMinted(address ) view returns(bool)
func (_Store *StoreCaller) WhiteListMinted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "whiteListMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteListMinted is a free data retrieval call binding the contract method 0x3ed50bb3.
//
// Solidity: function whiteListMinted(address ) view returns(bool)
func (_Store *StoreSession) WhiteListMinted(arg0 common.Address) (bool, error) {
	return _Store.Contract.WhiteListMinted(&_Store.CallOpts, arg0)
}

// WhiteListMinted is a free data retrieval call binding the contract method 0x3ed50bb3.
//
// Solidity: function whiteListMinted(address ) view returns(bool)
func (_Store *StoreCallerSession) WhiteListMinted(arg0 common.Address) (bool, error) {
	return _Store.Contract.WhiteListMinted(&_Store.CallOpts, arg0)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_Store *StoreTransactor) AddOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addOperator", _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_Store *StoreSession) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddOperator(&_Store.TransactOpts, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_Store *StoreTransactorSession) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddOperator(&_Store.TransactOpts, _operator)
}

// BatchFreeMint is a paid mutator transaction binding the contract method 0x98eef4b7.
//
// Solidity: function batchFreeMint(address[] _receivers) returns()
func (_Store *StoreTransactor) BatchFreeMint(opts *bind.TransactOpts, _receivers []common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "batchFreeMint", _receivers)
}

// BatchFreeMint is a paid mutator transaction binding the contract method 0x98eef4b7.
//
// Solidity: function batchFreeMint(address[] _receivers) returns()
func (_Store *StoreSession) BatchFreeMint(_receivers []common.Address) (*types.Transaction, error) {
	return _Store.Contract.BatchFreeMint(&_Store.TransactOpts, _receivers)
}

// BatchFreeMint is a paid mutator transaction binding the contract method 0x98eef4b7.
//
// Solidity: function batchFreeMint(address[] _receivers) returns()
func (_Store *StoreTransactorSession) BatchFreeMint(_receivers []common.Address) (*types.Transaction, error) {
	return _Store.Contract.BatchFreeMint(&_Store.TransactOpts, _receivers)
}

// ChangeBasePrice is a paid mutator transaction binding the contract method 0xc6cfe192.
//
// Solidity: function changeBasePrice(uint256 _basePrice) returns()
func (_Store *StoreTransactor) ChangeBasePrice(opts *bind.TransactOpts, _basePrice *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changeBasePrice", _basePrice)
}

// ChangeBasePrice is a paid mutator transaction binding the contract method 0xc6cfe192.
//
// Solidity: function changeBasePrice(uint256 _basePrice) returns()
func (_Store *StoreSession) ChangeBasePrice(_basePrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangeBasePrice(&_Store.TransactOpts, _basePrice)
}

// ChangeBasePrice is a paid mutator transaction binding the contract method 0xc6cfe192.
//
// Solidity: function changeBasePrice(uint256 _basePrice) returns()
func (_Store *StoreTransactorSession) ChangeBasePrice(_basePrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangeBasePrice(&_Store.TransactOpts, _basePrice)
}

// ChangeFreeThreshold is a paid mutator transaction binding the contract method 0xdb5e5c3d.
//
// Solidity: function changeFreeThreshold(uint256 _value) returns()
func (_Store *StoreTransactor) ChangeFreeThreshold(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changeFreeThreshold", _value)
}

// ChangeFreeThreshold is a paid mutator transaction binding the contract method 0xdb5e5c3d.
//
// Solidity: function changeFreeThreshold(uint256 _value) returns()
func (_Store *StoreSession) ChangeFreeThreshold(_value *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangeFreeThreshold(&_Store.TransactOpts, _value)
}

// ChangeFreeThreshold is a paid mutator transaction binding the contract method 0xdb5e5c3d.
//
// Solidity: function changeFreeThreshold(uint256 _value) returns()
func (_Store *StoreTransactorSession) ChangeFreeThreshold(_value *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangeFreeThreshold(&_Store.TransactOpts, _value)
}

// ChangeIcat is a paid mutator transaction binding the contract method 0x370fc62d.
//
// Solidity: function changeIcat(address _cat) returns()
func (_Store *StoreTransactor) ChangeIcat(opts *bind.TransactOpts, _cat common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changeIcat", _cat)
}

// ChangeIcat is a paid mutator transaction binding the contract method 0x370fc62d.
//
// Solidity: function changeIcat(address _cat) returns()
func (_Store *StoreSession) ChangeIcat(_cat common.Address) (*types.Transaction, error) {
	return _Store.Contract.ChangeIcat(&_Store.TransactOpts, _cat)
}

// ChangeIcat is a paid mutator transaction binding the contract method 0x370fc62d.
//
// Solidity: function changeIcat(address _cat) returns()
func (_Store *StoreTransactorSession) ChangeIcat(_cat common.Address) (*types.Transaction, error) {
	return _Store.Contract.ChangeIcat(&_Store.TransactOpts, _cat)
}

// ChangePublicThreshold is a paid mutator transaction binding the contract method 0xbdf884c3.
//
// Solidity: function changePublicThreshold(uint256 _value) returns()
func (_Store *StoreTransactor) ChangePublicThreshold(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changePublicThreshold", _value)
}

// ChangePublicThreshold is a paid mutator transaction binding the contract method 0xbdf884c3.
//
// Solidity: function changePublicThreshold(uint256 _value) returns()
func (_Store *StoreSession) ChangePublicThreshold(_value *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangePublicThreshold(&_Store.TransactOpts, _value)
}

// ChangePublicThreshold is a paid mutator transaction binding the contract method 0xbdf884c3.
//
// Solidity: function changePublicThreshold(uint256 _value) returns()
func (_Store *StoreTransactorSession) ChangePublicThreshold(_value *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangePublicThreshold(&_Store.TransactOpts, _value)
}

// FreeMint is a paid mutator transaction binding the contract method 0xfa07ce1d.
//
// Solidity: function freeMint(address _to) returns()
func (_Store *StoreTransactor) FreeMint(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "freeMint", _to)
}

// FreeMint is a paid mutator transaction binding the contract method 0xfa07ce1d.
//
// Solidity: function freeMint(address _to) returns()
func (_Store *StoreSession) FreeMint(_to common.Address) (*types.Transaction, error) {
	return _Store.Contract.FreeMint(&_Store.TransactOpts, _to)
}

// FreeMint is a paid mutator transaction binding the contract method 0xfa07ce1d.
//
// Solidity: function freeMint(address _to) returns()
func (_Store *StoreTransactorSession) FreeMint(_to common.Address) (*types.Transaction, error) {
	return _Store.Contract.FreeMint(&_Store.TransactOpts, _to)
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Store *StoreTransactor) PublicMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "publicMint")
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Store *StoreSession) PublicMint() (*types.Transaction, error) {
	return _Store.Contract.PublicMint(&_Store.TransactOpts)
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Store *StoreTransactorSession) PublicMint() (*types.Transaction, error) {
	return _Store.Contract.PublicMint(&_Store.TransactOpts)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_Store *StoreTransactor) RemoveOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "removeOperator", _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_Store *StoreSession) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _Store.Contract.RemoveOperator(&_Store.TransactOpts, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_Store *StoreTransactorSession) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _Store.Contract.RemoveOperator(&_Store.TransactOpts, _operator)
}

// TransferEth is a paid mutator transaction binding the contract method 0x30df47af.
//
// Solidity: function transferEth(address _to) returns()
func (_Store *StoreTransactor) TransferEth(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferEth", _to)
}

// TransferEth is a paid mutator transaction binding the contract method 0x30df47af.
//
// Solidity: function transferEth(address _to) returns()
func (_Store *StoreSession) TransferEth(_to common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferEth(&_Store.TransactOpts, _to)
}

// TransferEth is a paid mutator transaction binding the contract method 0x30df47af.
//
// Solidity: function transferEth(address _to) returns()
func (_Store *StoreTransactorSession) TransferEth(_to common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferEth(&_Store.TransactOpts, _to)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address _newOwner) returns()
func (_Store *StoreTransactor) TransferOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwner", _newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address _newOwner) returns()
func (_Store *StoreSession) TransferOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwner(&_Store.TransactOpts, _newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address _newOwner) returns()
func (_Store *StoreTransactorSession) TransferOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwner(&_Store.TransactOpts, _newOwner)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x5dc89b34.
//
// Solidity: function whiteListMint() returns()
func (_Store *StoreTransactor) WhiteListMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "whiteListMint")
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x5dc89b34.
//
// Solidity: function whiteListMint() returns()
func (_Store *StoreSession) WhiteListMint() (*types.Transaction, error) {
	return _Store.Contract.WhiteListMint(&_Store.TransactOpts)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x5dc89b34.
//
// Solidity: function whiteListMint() returns()
func (_Store *StoreTransactorSession) WhiteListMint() (*types.Transaction, error) {
	return _Store.Contract.WhiteListMint(&_Store.TransactOpts)
}
