package main

import (
	"context"
	"fmt"
	"github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")

	if err != nil {
		log.Fatalf("init client fial  ")
	}
	//加载合约
	address := common.HexToAddress("0x7fe6Dfe5FC217649bf9f0508a46d152e20a5E9b1")
	contract, err := contracts.NewStore(address, client)
	if err != nil {
		log.Fatal("load contract fail")

	}

	//加载私钥 其实合约的部署也是一个 特殊的交易
	privateKey, err := crypto.HexToECDSA("f1b7ecffe901583acd475823b3b2b6c448836774f908e83f447260695f0f4b68")
	chainID, err := client.NetworkID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foaaaao"))
	copy(value[:], []byte("0202123020203"))
	tx, err := contract.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("等待交易确认失败:", err)
	}
	if receipt.Status != 1 {
		log.Fatal("交易执行失败，Status=0")
	}
	fmt.Println("交易已成功上链！")

	// 读取状态
	result, err := contract.Items(nil, key)
	if err != nil {
		log.Fatal("读取状态失败:", err)
	}
	fmt.Println("✅ 最新查询结果是:", string(result[:]))

}
