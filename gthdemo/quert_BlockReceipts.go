package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

func main() {
	client, errmsg := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if nil != errmsg {
		log.Fatalln("初始化客户端失败:", errmsg)

	}
	blockNumber := big.NewInt(8545272)
	//查询区块收据
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatalln("获取区块失败:", err)
	}

	for _, receipt := range receiptsByNum {
		log.Println("交易hash:", receipt.TxHash.Hex())
		log.Println("交易状态:", receipt.Status)
		log.Println("gasUsed:", receipt.GasUsed)
		log.Println("gasPrice:", receipt.EffectiveGasPrice)
		log.Println("blobGasUsed:", receipt.BlobGasUsed)
		log.Println("blobGasPrice:", receipt.BlobGasPrice)
		log.Println("cumulativeGasUsed:", receipt.CumulativeGasUsed)
		log.Println("logs:", receipt.Logs)
		log.Println("contractAddress:", receipt.ContractAddress)
		log.Println("blockHash:", receipt.BlockHash.Hex())
		log.Println("blockNumber:", receipt.BlockNumber)
		log.Println("transactionIndex:", receipt.TransactionIndex)
		log.Println("type:", receipt.Type)
		log.Println("root:", receipt.PostState)
		log.Println("bloom:", receipt.Bloom)

	}

	res := common.HexToHash("0xae10e79423d2aa15a58b5cb9519dc35fe6d02a2f754dd55a213d51e79fc860af")
	tracRep, err := client.TransactionReceipt(context.Background(), res)
	if err != nil {
		log.Fatalln("获取交易收据失败:", err)
	}
	log.Println("交易hash:", tracRep.TxHash.Hex())
	log.Println("交易状态:", tracRep.Status)
	log.Println("gasUsed:", tracRep.GasUsed)
	log.Println("gasPrice:", tracRep.EffectiveGasPrice)
	log.Println("blobGasUsed:", tracRep.BlobGasUsed)
	log.Println("blobGasPrice:", tracRep.BlobGasPrice)
	log.Println("cumulativeGasUsed:", tracRep.CumulativeGasUsed)
	log.Println("logs:", tracRep.Logs)
	log.Println("contractAddress:", tracRep.ContractAddress)
	log.Println("blockHash:", tracRep.BlockHash.Hex())
	log.Println("blockNumber:", tracRep.BlockNumber)
	log.Println("transactionIndex:", tracRep.TransactionIndex)
	log.Println("type:", tracRep.Type)
	log.Println("root:", tracRep.PostState)
	log.Println("bloom:", tracRep.Bloom)

}
