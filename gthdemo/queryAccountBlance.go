package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")

	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x59CB4FF9Fa196fE78E0A515B88F21A15A526d5A0")
	addressStr := account.Hex()
	fmt.Println("地址字符串：", addressStr)
	//获取账户余额 nil 代表最新余额
	blance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("报错了", err)
	}

	log.Println("余额为:", blance)
	fbalance := new(big.Float)
	fbalance.SetString(blance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	log.Println("转换为--ETH余额为转换:", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("正在处理的余额：", pendingBalance)

}
