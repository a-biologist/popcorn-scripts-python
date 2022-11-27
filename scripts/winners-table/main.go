package main

import (
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/stein-f/popcorn-scripts/algorand"
	"github.com/stein-f/popcorn-scripts/popcorn"
	"sort"
)

const (
	shrimpASAID       = 360019122
	shrimpWallet      = "POPCORNWIGBQSN7KTVJVGGYIP6CSUDMWD3BROJG2HMAXH73N4OQ3QJJN5M"
	transactionsAfter = "2022-11-25T12:00:00Z"
)

var gamesWeek4 = []popcorn.Game{
	{
		Team1:  "Japan",
		Team2:  "Costa Rica",
		Type:   popcorn.TypeSpread,
		Result: "Costa Rica",
	},
	{
		Team1:  "Japan",
		Team2:  "Costa Rica",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Japan",
		Team2:  "Costa Rica",
		Type:   popcorn.TypeCorners,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Belgium",
		Team2:  "Morocco",
		Type:   popcorn.TypeSpread,
		Result: "Morocco",
	},
	{
		Team1:  "Belgium",
		Team2:  "Morocco",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Belgium",
		Team2:  "Morocco",
		Type:   popcorn.TypeCorners,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Croatia",
		Team2:  "Canada",
		Type:   popcorn.TypeSpread,
		Result: "Croatia",
	},
	{
		Team1:  "Croatia",
		Team2:  "Canada",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Croatia",
		Team2:  "Canada",
		Type:   popcorn.TypeCorners,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Spain",
		Team2:  "Germany",
		Type:   popcorn.TypeSpread,
		Result: "Germany",
	},
	{
		Team1:  "Spain",
		Team2:  "Germany",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Spain",
		Team2:  "Germany",
		Type:   popcorn.TypeCorners,
		Result: popcorn.ChoiceOver,
	},
}

func main() {
	indexerClient, err := indexer.MakeClient("https://mainnet-idx.algonode.cloud", "")
	if err != nil {
		panic(err)
	}

	txns, err := algorand.FetchTransactionsAfterTime(indexerClient, shrimpWallet, shrimpASAID, transactionsAfter)
	if err != nil {
		panic(err)
	}

	resultsByUser := map[string][]popcorn.Result{}
	for _, tx := range txns {
		bet, ok := popcorn.ParseTxNote(string(tx.Note))
		if !ok {
			continue
		}
		game, ok := popcorn.FindGame(gamesWeek4, bet)
		if !ok {
			continue
		}

		result := popcorn.GetBetResult(bet, game)

		resultsByUser[tx.Sender] = append(resultsByUser[tx.Sender], result)
	}

	type counts struct {
		Wins   int
		Push   int
		Loss   int
		Shrimp int
	}

	winsByUser := map[string]counts{}
	for user, results := range resultsByUser {
		mergedResults := popcorn.MergeResults(results)
		if len(mergedResults) != len(gamesWeek4) {
			fmt.Printf("Reject %s, bets: %d\n", user, len(mergedResults))
			continue
		}
		win, push, loss := countResult(mergedResults)
		winsByUser[algorand.ResolveNfd(user)] = counts{
			Wins:   win,
			Push:   push,
			Loss:   loss,
			Shrimp: countShrimp(mergedResults),
		}
	}

	for user, wins := range winsByUser {
		fmt.Printf("%s - %d\n", user, wins)
	}

	type kv struct {
		User   string
		Wins   int
		Push   int
		Loss   int
		Amount int
	}

	var ss []kv
	for k, v := range winsByUser {
		ss = append(ss, kv{User: k, Wins: v.Wins, Push: v.Push, Loss: v.Loss, Amount: v.Shrimp})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Wins != ss[j].Wins {
			return ss[i].Wins > ss[j].Wins
		}
		if ss[i].Amount != ss[j].Amount {
			return ss[i].Amount > ss[j].Amount
		}
		return ss[i].Push > ss[j].Push
	})

	fmt.Printf("W-L-P\n")
	for _, kv := range ss {
		fmt.Printf("%d-%d-%d, shrimp: %d\t%s\n", kv.Wins, kv.Loss, kv.Push, kv.Amount, kv.User)
	}
}

func countResult(results []popcorn.Result) (int, int, int) {
	var wins, push, loss int
	for _, res := range results {
		if res.Result == popcorn.ResultWin {
			wins++
		}
		if res.Result == popcorn.ResultLose {
			loss++
		}
		if res.Result == popcorn.ResultPush {
			push++
		}
	}
	return wins, push, loss
}

func countShrimp(results []popcorn.Result) int {
	var shrimp int
	for _, res := range results {
		shrimp += res.Amount
	}
	return shrimp
}

var gamesWeek1 = []popcorn.Game{
	{
		Team1:  "Ecuador",
		Team2:  "Qatar",
		Type:   popcorn.TypeSpread,
		Result: "Ecuador",
	},
	{
		Team1:  "Ecuador",
		Team2:  "Qatar",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "England",
		Team2:  "Iran",
		Type:   popcorn.TypeSpread,
		Result: "England",
	},
	{
		Team1:  "England",
		Team2:  "Iran",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Senegal",
		Team2:  "Netherlands",
		Type:   popcorn.TypeSpread,
		Result: "Netherlands",
	},
	{
		Team1:  "Senegal",
		Team2:  "Netherlands",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "USA",
		Team2:  "Wales",
		Type:   popcorn.TypeSpread,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "USA",
		Team2:  "Wales",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Argentina",
		Team2:  "Saudi Arabia",
		Type:   popcorn.TypeSpread,
		Result: "Saudi Arabia",
	},
	{
		Team1:  "Argentina",
		Team2:  "Saudi Arabia",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Denmark",
		Team2:  "Tunisia",
		Type:   popcorn.TypeSpread,
		Result: "Tunisia",
	},
	{
		Team1:  "Denmark",
		Team2:  "Tunisia",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Mexico",
		Team2:  "Poland",
		Type:   popcorn.TypeSpread,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Mexico",
		Team2:  "Poland",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "France",
		Team2:  "Australia",
		Type:   popcorn.TypeSpread,
		Result: "France",
	},
	{
		Team1:  "France",
		Team2:  "Australia",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
}

var gamesWeek2 = []popcorn.Game{
	{
		Team1:  "Croatia",
		Team2:  "Morocco",
		Type:   popcorn.TypeSpread,
		Result: "Morocco",
	},
	{
		Team1:  "Croatia",
		Team2:  "Morocco",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Germany",
		Team2:  "Japan",
		Type:   popcorn.TypeSpread,
		Result: "Japan",
	},
	{
		Team1:  "Germany",
		Team2:  "Japan",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Spain",
		Team2:  "Costa Rica",
		Type:   popcorn.TypeSpread,
		Result: "Spain",
	},
	{
		Team1:  "Spain",
		Team2:  "Costa Rica",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Belgium",
		Team2:  "Canada",
		Type:   popcorn.TypeSpread,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Belgium",
		Team2:  "Canada",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Switzerland",
		Team2:  "Cameroon",
		Type:   popcorn.TypeSpread,
		Result: "Switzerland",
	},
	{
		Team1:  "Switzerland",
		Team2:  "Cameroon",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Uruguay",
		Team2:  "South Korea",
		Type:   popcorn.TypeSpread,
		Result: "South Korea",
	},
	{
		Team1:  "Uruguay",
		Team2:  "South Korea",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Portugal",
		Team2:  "Ghana",
		Type:   popcorn.TypeSpread,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Portugal",
		Team2:  "Ghana",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Brazil",
		Team2:  "Serbia",
		Type:   popcorn.TypeSpread,
		Result: "Brazil",
	},
	{
		Team1:  "Brazil",
		Team2:  "Serbia",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
}

var gamesWeek3 = []popcorn.Game{
	{
		Team1:  "Japan",
		Team2:  "Costa Rica",
		Type:   popcorn.TypeSpread,
		Result: "Iran",
	},
	{
		Team1:  "Wales",
		Team2:  "Iran",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "Qatar",
		Team2:  "Senegal",
		Type:   popcorn.TypeSpread,
		Result: "Senegal",
	},
	{
		Team1:  "Qatar",
		Team2:  "Senegal",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Ecuador",
		Team2:  "Netherlands",
		Type:   popcorn.TypeSpread,
		Result: "Ecuador",
	},
	{
		Team1:  "Ecuador",
		Team2:  "Netherlands",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "USA",
		Team2:  "England",
		Type:   popcorn.TypeSpread,
		Result: "USA",
	},
	{
		Team1:  "USA",
		Team2:  "England",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Tunisia",
		Team2:  "Australia",
		Type:   popcorn.TypeSpread,
		Result: "Australia",
	},
	{
		Team1:  "Tunisia",
		Team2:  "Australia",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
	{
		Team1:  "Poland",
		Team2:  "Saudi Arabia",
		Type:   popcorn.TypeSpread,
		Result: "Poland",
	},
	{
		Team1:  "Poland",
		Team2:  "Saudi Arabia",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "France",
		Team2:  "Denmark",
		Type:   popcorn.TypeSpread,
		Result: popcorn.ResultPush,
	},
	{
		Team1:  "France",
		Team2:  "Denmark",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceOver,
	},
	{
		Team1:  "Argentina",
		Team2:  "Mexico",
		Type:   popcorn.TypeSpread,
		Result: "Argentina",
	},
	{
		Team1:  "Argentina",
		Team2:  "Mexico",
		Type:   popcorn.TypeOverUnder,
		Result: popcorn.ChoiceUnder,
	},
}
