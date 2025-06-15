package main

import (
	"context"
	"fmt"
	"github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//使用abigen生成的 go 进行合约的部署
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal("初始化客户端失败:", err)
	}
	//加载私钥 其实合约的部署也是一个 特殊的交易
	privateKey, err := crypto.HexToECDSA("f1b7ecffe901583acd475823b3b2b6c448836774f908e83f447260695f0f4b68")
	chainID, err := client.NetworkID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	address, tx, instance, err := contracts.DeployStore(auth, client, "1.0")
	if err != nil {
		log.Fatal("部署合约失败:", err)
	}

	fmt.Println(address.Hex())   //0x7fe6Dfe5FC217649bf9f0508a46d152e20a5E9b1
	fmt.Println(tx.Hash().Hex()) //	0x3ca823d0ec9ed6b8c3cac39069bdc815dd72cffff73a89083b1d9d351fb2f471

	_ = instance

}
