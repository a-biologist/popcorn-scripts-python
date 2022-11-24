package algorand

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"net/http"
)

func FetchTransactionsAfterTime(indexerClient *indexer.Client, accountID string, assetID uint64, afterTime string) ([]models.Transaction, error) {
	next := ""
	var txns []models.Transaction
	for {
		tx, err := indexerClient.LookupAccountTransactions(accountID).
			AfterTimeString(afterTime).
			Limit(1000).
			NextToken(next).
			AssetID(assetID).
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

func TruncateAddress(s string) string {
	first5 := s[0:5]
	last5 := s[len(s)-5:]
	return fmt.Sprintf("%s...%s", first5, last5)
}

func ResolveNfd(wallet string) string {
	res, err := http.DefaultClient.Get("https://api.nf.domains/nfd/address?address=" + wallet)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		return TruncateAddress(wallet)
	}
	var recs []nfdRes
	if err := json.NewDecoder(res.Body).Decode(&recs); err != nil {
		panic(err)
	}
	if len(recs) == 0 {
		return wallet
	}
	return recs[0].Name
}

type nfdRes struct {
	Name string `json:"name"`
}
