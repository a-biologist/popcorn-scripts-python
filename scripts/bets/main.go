package main

import (
	"context"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"regexp"
	"strings"
)

const (
	shrimpASAID  = 360019122
	shrimpWallet = "POPCORNWIGBQSN7KTVJVGGYIP6CSUDMWD3BROJG2HMAXH73N4OQ3QJJN5M"
)

type Bet struct {
	Game   string
	Bet    string
	Amount string
}

var txNoteRegexp = regexp.MustCompile(`game:(?P<game>.+)bet:(?P<bet>.+)amount:(?P<amount>.+)`)

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
		bet, ok := parseTxNote(string(tx.Note))
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

// Parses the transaction note field. Covers these basic examples only
// game: England (-1.5) vs Iran (+1.5) bet: Iran amount: 300
// game: England vs Iran bet: Over 2.5 Goals amount: 300
func parseTxNote(txNote string) (Bet, bool) {
	matches := txNoteRegexp.FindStringSubmatch(txNote)
	if matches == nil {
		return Bet{}, false
	}

	game := strings.TrimSpace(matches[txNoteRegexp.SubexpIndex("game")])
	bet := strings.TrimSpace(matches[txNoteRegexp.SubexpIndex("bet")])
	amount := strings.TrimSpace(matches[txNoteRegexp.SubexpIndex("amount")])

	return Bet{Game: game, Bet: bet, Amount: amount}, true
}
