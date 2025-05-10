package services

import (
    "context"
    "crypto/ecdsa"
    _ "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
    "traceability/config"
)

func WriteToBlockchain(ipfsHash string) (string, error) {
    client, err := ethclient.Dial(config.BNBEndpoint)
    if err != nil {
        return "", err
    }

    privateKey, err := crypto.HexToECDSA(config.PrivateKey)
    if err != nil {
        return "", err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
    gasPrice, _ := client.SuggestGasPrice(context.Background())

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(300000)
    auth.GasPrice = gasPrice

    tx := types.NewTransaction(nonce, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(0), 300000, gasPrice, []byte(ipfsHash))
    signedTx, _ := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(97)), privateKey)

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return "", err
    }

    return signedTx.Hash().Hex(), nil
}
