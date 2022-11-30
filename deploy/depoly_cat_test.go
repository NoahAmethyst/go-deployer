package deploy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tristan-club/kit/chain_info"
	"go-deployer/constant"
	cat_contract "go-deployer/contract/cat"
	minter_contract "go-deployer/contract/cat_minter"
	"os"
	"testing"
)

func TestDeployCat(t *testing.T) {

	chainType := uint32(chain_info.ChainTypePolygon)
	pk := os.Getenv(constant.PK)
	client, auth, err := NewEthCliByChainType(chainType, pk)
	if err != nil {
		t.Errorf("new eth cli error:%s", err)
		return
	}
	fromAddress := auth.From

	//Deploy cat contract
	catAddress, catDeployTx, catInstance, err := cat_contract.DeployStore(auth, client, "Test", "Test cat",
		common.HexToAddress("0x4a99683dd308243f4b2488683cF0C92C37789799"), "Test")
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	t.Logf("cat contract address:%s", catAddress.Hex())
	t.Logf("cat contract deploy hash:%s", catDeployTx.Hash().Hex())

	//Deploy cat contract
	if err := refreshTxCallData(client, auth); err != nil {
		t.Errorf("set new nonce error:%s", err)
		return
	}
	minterAddress, minterDeployTx, minterInstance, err := minter_contract.DeployStore(auth, client, catAddress)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	t.Logf("minter contract address:%s", minterAddress.Hex())
	t.Logf("minter contract deploy hash:%s", minterDeployTx.Hash().Hex())

	//call contract
	if err = refreshTxCallData(client, auth); err != nil {
		t.Errorf("set new nonce error:%s", err)
		return
	}
	_, err = catInstance.AddOperator(auth, minterAddress)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	if err = refreshTxCallData(client, auth); err != nil {
		t.Errorf("set new nonce error:%s", err)
		return
	}
	_, err = minterInstance.FreeMint(auth, fromAddress)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	of, err := catInstance.BalanceOf(nil, fromAddress)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	t.Logf("balance:%s", of.String())
}

func TestCallContract(t *testing.T) {
	chainType := uint32(chain_info.ChainTypePolygon)
	pk := os.Getenv(constant.PK)
	client, auth, err := NewEthCliByChainType(chainType, pk)
	if err != nil {
		t.Errorf("new eth cli error:%s", err)
		return
	}
	fromAddress := auth.From

	instance, err := cat_contract.NewStore(common.HexToAddress("0xe20EA3D034bA8c7dEdAD73FF16BCaE966e204278"), client)
	if err != nil {
		t.Errorf("get instance error:%s", err)
		return
	}

	of, err := instance.BalanceOf(nil, fromAddress)
	if err != nil {
		t.Errorf("call contract error:%s", err)
		return
	}
	t.Logf("balance:%s", of.String())

}
