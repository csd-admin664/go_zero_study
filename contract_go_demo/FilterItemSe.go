package main

import (
	"fmt"
	"github.com/csd-admin664/go_zero_study/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, _ := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	//获取合约地址
	address := common.HexToAddress("0x7fe6Dfe5FC217649bf9f0508a46d152e20a5E9b1")
	con, _ := contracts.NewStore(address, client)
	lase := uint64(20000)
	filterOpts := bind.FilterOpts{
		Start:   0,
		End:     &lase,
		Context: nil,
	}
	iter, err := con.FilterItemSet(&filterOpts)
	if err != nil {
		log.Fatalf("查询事件失败,", err)

	}
	for iter.Next() {
		event := iter.Event
		fmt.Printf("➡️  key: %s, value: %s, block: %v\n",
			string(event.Key[:]), string(event.Value[:]), event.Raw.BlockNumber)
	}
}
