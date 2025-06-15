package main

import (
	"fmt"
	"github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal("初始化客户端失败:", err)
	}
	address := common.HexToAddress("0x0FB80D07Dcc78DC6f142E4F532F1a1743c7B70e3")

	instance, err := contracts.NewMyToken(address, client)
	if err != nil {
		log.Fatal("加载合约手失败	:", err)
	}
	fmt.Println("contract is loaded")
	_ = instance
}
