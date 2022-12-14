package deploy

import (
	"auto-deployer/utils/log"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewEthCliByChainType(rpc, pk string) (*ethclient.Client, *bind.TransactOpts, error) {

	return NewEthCli(rpc, pk)
}

func NewEthCli(rpc, pk string) (*ethclient.Client, *bind.TransactOpts, error) {
	c := context.Background()

	var client *ethclient.Client
	var err error
	var auth *bind.TransactOpts

	client, err = ethclient.Dial(rpc)
	if err != nil {
		log.Error().Msgf("init eth cli error:%s", err)
		return client, auth, err
	}
	id, err := client.ChainID(c)
	if err != nil {
		log.Error().Msgf("get chain id error:%s", err)
	} else {
		log.Info().Msgf("chainId:%s", id.String())
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Error().Msgf("%s", err)
		return client, auth, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Msgf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return client, auth, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error().Msgf("%s", err)
		return client, auth, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error().Msgf("%s", err)
		return client, auth, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, id)
	if err != nil {
		log.Error().Msgf("%s", err)
		return client, auth, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return client, auth, err
}
