package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	MyTonken "github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal("初始化客户端失败:", err)
	}
	privateKey, err := crypto.HexToECDSA("69a9db9723bba316f267bff4d4876c459d813b58328e04af5f9fd7e9a39e2583")
	if err != nil {
		log.Fatal("加载私钥失败", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//的到地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//设置nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//设置gas
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	pid, err := client.NetworkID(context.Background())
	auth := bind.NewKeyedTransactor(privateKey, pid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := MyTonken.DeployMyToken(auth, client, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x0FB80D07Dcc78DC6f142E4F532F1a1743c7B70e3
	fmt.Println(tx.Hash().Hex()) // 0x6f5a929ad0e088cc9e710d4ec94631ae5f028b320a14cba92831fca842f394f5

	_ = instance

}
