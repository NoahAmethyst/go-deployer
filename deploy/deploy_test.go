package deploy

import (
	"auto-deployer/constant"
	token_contract "auto-deployer/contract/token"
	"auto-deployer/utils/log"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Test_Deploy(t *testing.T) {
	deploy()
}

func deploy() {

	c := context.Background()

	client, err := ethclient.Dial("https://rpc-mainnet.maticvigil.com/")
	if err != nil {
		log.Error().Msgf("init eth cli error:%s", err)
	}
	id, err := client.ChainID(c)
	if err != nil {
		log.Error().Msgf("get chain id error:%s", err)
	} else {
		log.Info().Msgf("chainId:%s", id.String())
	}

	pk := os.Getenv(constant.PK)

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Error().Msgf("%s", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Msgf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error().Msgf("%s", err)

	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error().Msgf("%s", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
	if err != nil {
		log.Error().Msgf("%s", err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//Deploy contract
	address, tx, instance, err := token_contract.DeployStore(auth, client, "T", "T")
	if err != nil {
		log.Error().Msgf("%s", err)
		return
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	//call contract
	symbol, err := instance.Symbol(nil)
	if err != nil {
		log.Error().Msgf("%s", err)
		return
	}

	log.Info().Msgf("%s", symbol)

	_, err = instance.Mint(auth, fromAddress, big.NewInt(100000))
	if err != nil {
		log.Error().Msgf("%s", err)
		return
	}

	of, err := instance.BalanceOf(nil, fromAddress)
	if err != nil {
		log.Error().Msgf("%s", err)
		return
	}

	log.Info().Msgf("get balance:%s", of.String())
}
