package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

/*
*

	监听新的区块
*/
func main() {
	client, er := ethclient.Dial("wss://mainnet.infura.io/ws/v3/YOUR_INFURA_KEY")
	if er != nil {
		log.Fatal("初始化客户端失败:", er)
	}
	headers := make(chan *types.Header)
	//调用客户端的 订阅区块头的方法
	sub, _ := client.SubscribeNewHead(context.Background(), headers)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}

	}

}
