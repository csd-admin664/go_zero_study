package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 查询区块头 HeaderByNumber  2 查询完整区块
// 调用区块头查询参数涉及到两个接口 1
func main() {
	client, err := ethclient.Dial("https://ethereum.publicnode.com/?sepolia")
	if err != nil {
		panic(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	fmt.Println(header.Number.String()) //22675215
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number().Uint64())     // 22675215
	fmt.Println(block.Time())                // 1749571199
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 208

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144

}
