package algorand

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/algorand/go-algorand-sdk/client/v2/indexer"
    "sort"
)

func GetAssetConfigTxJSON(ctx context.Context, indexerClient *indexer.Client, index uint64) (map[string]interface{}, error) {
    transactionsResponse, err := indexerClient.SearchForTransactions().
        AssetID(index).
        TxType("acfg").
        Do(ctx)
    if err != nil {
        panic(err)
    }

    sort.Slice(transactionsResponse.Transactions, func(i, j int) bool {
        return transactionsResponse.Transactions[i].RoundTime > transactionsResponse.Transactions[j].RoundTime
    })

    for _, tx := range transactionsResponse.Transactions {
        var arc69Metadata map[string]interface{}
        if err = json.Unmarshal(tx.Note, &arc69Metadata); err != nil {
            return nil, fmt.Errorf("failed to parse arc69 metadata. %w", err)
        }
        if arc69Metadata["standard"] == "arc69" {
            return arc69Metadata, nil
        }
    }

    return nil, fmt.Errorf("failed to find arc69 metadata for asset %d", index)
}