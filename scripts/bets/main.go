package main

import (
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/stein-f/popcorn-scripts/algorand"
	"github.com/stein-f/popcorn-scripts/popcorn"
	"strings"
)

const (
	shrimpASAID       = 360019122
	shrimpWallet      = "POPCORNWIGBQSN7KTVJVGGYIP6CSUDMWD3BROJG2HMAXH73N4OQ3QJJN5M"
	transactionsAfter = "2022-11-22T12:00:00Z"
)

func main() {
	indexerClient, err := indexer.MakeClient("https://mainnet-idx.algonode.cloud", "")
	if err != nil {
		panic(err)
	}

	txns, err := algorand.FetchTransactionsAfterTime(indexerClient, shrimpWallet, shrimpASAID, transactionsAfter)
	if err != nil {
		panic(err)
	}

	var germanyCount, japanCount, overCount, underCount int
	for _, tx := range txns {
		bet, ok := popcorn.ParseTxNote(string(tx.Note))
		if !ok {
			continue
		}
		if !strings.Contains(bet.Game, "Switzerland") || !strings.Contains(bet.Game, "Cameroon") {
			continue
		}

		if strings.Contains(bet.Bet, "Switzerland") {
			germanyCount++
		}
		if strings.Contains(bet.Bet, "Cameroon") {
			japanCount++
		}
		if strings.Contains(bet.Bet, "Under") {
			underCount++
		}
		if strings.Contains(bet.Bet, "Over") {
			overCount++
		}
	}

	fmt.Printf("Switzerland: %d\n", germanyCount)
	fmt.Printf("Cameroon: %d\n", japanCount)
	fmt.Printf("Over: %d\n", overCount)
	fmt.Printf("Under: %d\n", underCount)
}
