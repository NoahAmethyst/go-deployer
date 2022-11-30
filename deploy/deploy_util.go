package deploy

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tristan-club/kit/chain_info"
	"go-deployer/utils/log"
	"math/big"
)

func NewEthCliByChainType(chainType uint32, pk string) (*ethclient.Client, *bind.TransactOpts, error) {
	net := chain_info.GetNetByChainType(chainType)
	if len(net.RpcUrl) == 0 {
		return nil, nil, errors.New("unsupported chain")
	}
	return NewEthCli(net.RpcUrl, pk)
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

// Todo 维护nonce
// 方案：1.从nodebalance获取nonce 2.nonce存入tstore,该账户不可人工使用
func refreshTxCallData(client *ethclient.Client, auth *bind.TransactOpts) error {
	//refresh nonce
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Error().Msgf("refresh nonce error:%s%s", err)
		return err
	}
	if nonce == auth.Nonce.Uint64() {
		auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	} else {
		auth.Nonce = new(big.Int).SetUint64(nonce)
	}

	//refresh gasprice
	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
		log.Error().Msgf("refresh gasprice error:%s", err)
		return err
	} else {
		auth.GasPrice = gasPrice
	}

	//auth.Nonce = newNonce.SetUint64(nonce)
	//auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	return nil
}
