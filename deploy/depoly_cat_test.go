package deploy

import (
	"auto-deployer/constant"
	cat_contract "auto-deployer/contract/cat"
	minter_contract "auto-deployer/contract/cat_minter"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"testing"
)

func TestDeployCat(t *testing.T) {

	polygonRpc := "https://polygon-rpc.com/"
	pk := os.Getenv(constant.PK)
	client, auth, err := NewEthCliByChainType(polygonRpc, pk)
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
	client, auth, err = NewEthCliByChainType(polygonRpc, pk)
	if err != nil {
		t.Errorf("new eth cli error:%s", err)
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
	client, auth, err = NewEthCliByChainType(polygonRpc, pk)
	if err != nil {
		t.Errorf("new eth cli error:%s", err)
		return
	}
	_, err = catInstance.AddOperator(auth, minterAddress)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	client, auth, err = NewEthCliByChainType(polygonRpc, pk)
	if err != nil {
		t.Errorf("new eth cli error:%s", err)
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
