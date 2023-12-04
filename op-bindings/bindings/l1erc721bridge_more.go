// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"messenger\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_contract(CrossDomainMessenger)1005\"},{\"astId\":1003,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)48_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"49\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)48_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1005\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e mapping(uint256 =\u003e bool)))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L1ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100be5760003560e01c80637f46ddb211610076578063aa5574521161005b578063aa55745214610213578063c4d66de814610226578063c89701a21461023957600080fd5b80637f46ddb2146101c8578063927ede2d146101ef57600080fd5b806354fd4d50116100a757806354fd4d50146101285780635d93a3fc14610171578063761f4493146101b557600080fd5b80633687011a146100c35780633cb747bf146100d8575b600080fd5b6100d66100d1366004610e47565b61025f565b005b6000546100fe9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101646040518060400160405280600581526020017f312e342e3000000000000000000000000000000000000000000000000000000081525081565b60405161011f9190610f35565b6101a561017f366004610f4f565b603160209081526000938452604080852082529284528284209052825290205460ff1681565b604051901515815260200161011f565b6100d66101c3366004610f90565b61030b565b6100fe7f000000000000000000000000000000000000000000000000000000000000000081565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff166100fe565b6100d6610221366004611028565b610776565b6100d661023436600461109f565b610832565b7f00000000000000000000000000000000000000000000000000000000000000006100fe565b333b156102f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b610303868633338888888861097c565b505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314801561041357507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103fb91906110bc565b73ffffffffffffffffffffffffffffffffffffffff16145b61049f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016102ea565b3073ffffffffffffffffffffffffffffffffffffffff881603610544576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c314552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016102ea565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152603160209081526040808320938a1683529281528282208683529052205460ff161515600114610613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4c314552433732314272696467653a20546f6b656e204944206973206e6f742060448201527f657363726f77656420696e20746865204c31204272696467650000000000000060648201526084016102ea565b73ffffffffffffffffffffffffffffffffffffffff87811660008181526031602090815260408083208b8616845282528083208884529091529081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f42842e0e000000000000000000000000000000000000000000000000000000008152306004820152918616602483015260448201859052906342842e0e90606401600060405180830381600087803b1580156106d357600080fd5b505af11580156106e7573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107659493929190611122565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516610819576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016102ea565b610829878733888888888861097c565b50505050505050565b600054600390610100900460ff16158015610854575060005460ff8083169116105b6108e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102ea565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff83161761010017905561091a82610cdc565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff8716610a1f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c314552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016102ea565b600063761f449360e01b888a8989898888604051602401610a469796959493929190611162565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000959095169490941790935273ffffffffffffffffffffffffffffffffffffffff8c81166000818152603186528381208e8416825286528381208b82529095529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f23b872dd000000000000000000000000000000000000000000000000000000008152908a166004820152306024820152604481018890529092506323b872dd90606401600060405180830381600087803b158015610b8657600080fd5b505af1158015610b9a573d6000803e3d6000fd5b50506000546040517f3dbb202b0000000000000000000000000000000000000000000000000000000081526201000090910473ffffffffffffffffffffffffffffffffffffffff169250633dbb202b9150610c1d907f000000000000000000000000000000000000000000000000000000000000000090859089906004016111bf565b600060405180830381600087803b158015610c3757600080fd5b505af1158015610c4b573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a589898888604051610cc99493929190611122565b60405180910390a4505050505050505050565b600054610100900460ff16610d73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102ea565b6000805473ffffffffffffffffffffffffffffffffffffffff90921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff81168114610de257600080fd5b50565b803563ffffffff81168114610df957600080fd5b919050565b60008083601f840112610e1057600080fd5b50813567ffffffffffffffff811115610e2857600080fd5b602083019150836020828501011115610e4057600080fd5b9250929050565b60008060008060008060a08789031215610e6057600080fd5b8635610e6b81610dc0565b95506020870135610e7b81610dc0565b945060408701359350610e9060608801610de5565b9250608087013567ffffffffffffffff811115610eac57600080fd5b610eb889828a01610dfe565b979a9699509497509295939492505050565b6000815180845260005b81811015610ef057602081850181015186830182015201610ed4565b81811115610f02576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610f486020830184610eca565b9392505050565b600080600060608486031215610f6457600080fd5b8335610f6f81610dc0565b92506020840135610f7f81610dc0565b929592945050506040919091013590565b600080600080600080600060c0888a031215610fab57600080fd5b8735610fb681610dc0565b96506020880135610fc681610dc0565b95506040880135610fd681610dc0565b94506060880135610fe681610dc0565b93506080880135925060a088013567ffffffffffffffff81111561100957600080fd5b6110158a828b01610dfe565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561104357600080fd5b873561104e81610dc0565b9650602088013561105e81610dc0565b9550604088013561106e81610dc0565b94506060880135935061108360808901610de5565b925060a088013567ffffffffffffffff81111561100957600080fd5b6000602082840312156110b157600080fd5b8135610f4881610dc0565b6000602082840312156110ce57600080fd5b8151610f4881610dc0565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006111586060830184866110d9565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526111b260c0830184866110d9565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006111ee6060830185610eca565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1ERC721BridgeStorageLayoutJSON), L1ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC721Bridge"] = L1ERC721BridgeStorageLayout
	deployedBytecodes["L1ERC721Bridge"] = L1ERC721BridgeDeployedBin
}