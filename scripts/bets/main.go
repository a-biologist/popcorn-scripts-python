package main

import (
	"context"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/stein-f/popcorn-scripts/popcorn"
	"strings"
)

const (
	shrimpASAID  = 360019122
	shrimpWallet = "POPCORNWIGBQSN7KTVJVGGYIP6CSUDMWD3BROJG2HMAXH73N4OQ3QJJN5M"
)

func main() {
	indexerClient, err := indexer.MakeClient("https://mainnet-idx.algonode.cloud", "")
	if err != nil {
		panic(err)
	}

	txns, err := fetchTransactionsAfterTime(indexerClient, "2022-11-21T12:00:00Z")
	if err != nil {
		panic(err)
	}

	var germanyCount, japanCount, overCount, underCount int
	for _, tx := range txns {
		bet, ok := popcorn.ParseTxNote(string(tx.Note))
		if !ok {
			continue
		}

		if strings.Contains(bet.Game, "Germany") && strings.Contains(bet.Game, "Japan") {
			if strings.Contains(bet.Bet, "Germany") {
				germanyCount++
			}
			if strings.Contains(bet.Bet, "Japan") {
				japanCount++
			}
			if strings.Contains(bet.Bet, "Under") {
				underCount++
			}
			if strings.Contains(bet.Bet, "Over") {
				overCount++
			}
		}
	}

	fmt.Printf("Germany: %d\n", germanyCount)
	fmt.Printf("Japan: %d\n", japanCount)
	fmt.Printf("Over: %d\n", overCount)
	fmt.Printf("Under: %d\n", underCount)
}

func fetchTransactionsAfterTime(indexerClient *indexer.Client, afterTime string) ([]models.Transaction, error) {
	next := ""
	var txns []models.Transaction
	for {
		tx, err := indexerClient.LookupAccountTransactions(shrimpWallet).
			AfterTimeString(afterTime).
			Limit(1000).
			NextToken(next).
			AssetID(shrimpASAID).
			Do(context.TODO())
		if err != nil {
			return nil, err
		}
		next = tx.NextToken
		txns = append(txns, tx.Transactions...)
		if tx.NextToken == "" {
			break
		}
	}
	return txns, nil
}
