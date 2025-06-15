package main

import (
	"fmt"
	"github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, _ := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	//获取合约地址
	address := common.HexToAddress("0x7fe6Dfe5FC217649bf9f0508a46d152e20a5E9b1")
	//加载合约
	contract, _ := contracts.NewStore(address, client)
	version, err := contract.Version(nil)
	if err != nil {
		fmt.Println("加载合约失败:", err)
	}
	fmt.Println("contract is loaded Version", version)

}
