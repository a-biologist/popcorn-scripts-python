package main

import (
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/stein-f/popcorn-scripts/algorand"
	"github.com/stein-f/popcorn-scripts/popcorn"
)

const (
	shrimpASAID       = 360019122
	wallet            = "EE7CRFODWJDOXUVPVBIDYHKSRJHNZ4R3B6AKHBU5VRB5YBDXGVSZULEYIQ"
	transactionsAfter = "2022-11-21T12:00:00Z"
)

func main() {
	indexerClient, err := indexer.MakeClient("https://mainnet-idx.algonode.cloud", "")
	if err != nil {
		panic(err)
	}

	txns, err := algorand.FetchTransactionsAfterTime(indexerClient, wallet, shrimpASAID, transactionsAfter)
	if err != nil {
		panic(err)
	}

	for _, tx := range txns {
		bet, ok := popcorn.ParseTxNote(string(tx.Note))
		if ok {
			fmt.Println(bet)
		}
	}
}
