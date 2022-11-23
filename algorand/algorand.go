package algorand

import (
    "context"
    "github.com/algorand/go-algorand-sdk/client/v2/common/models"
    "github.com/algorand/go-algorand-sdk/client/v2/indexer"
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