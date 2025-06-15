package main

import (
	"context"
	"fmt"
	"github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

//查询事件

func main() {

	client, errMsg := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")

	if nil != errMsg {
		log.Fatalln("初始化客户端失败:", errMsg)

	}

	//合约地址
	contractAddress := common.HexToAddress("0x7fe6Dfe5FC217649bf9f0508a46d152e20a5E9b1")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(8545236),
		ToBlock:   big.NewInt(8545272),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalln("查询日志失败:", err)
	}

	for _, val := range logs {
		fmt.Println(val.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(val.BlockNumber)     // 2394201
		fmt.Println(val.TxHash.Hex())
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		errm := contractAbi.UnpackIntoInterface(&event, "ItemSet", val.Data)
		if errm != nil {
			log.Fatal(err)
		}
		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range val.Topics {
			topics[i] = val.Topics[i].Hex()
		}

		fmt.Println(topics[0])

	}
}
