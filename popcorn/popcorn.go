package popcorn

import (
	"fmt"
	"github.com/stein-f/popcorn-scripts/lang/slice"
	"regexp"
	"strconv"
	"strings"
)

var txNoteRegexp = regexp.MustCompile(`game:(?P<game>.+)bet:(?P<bet>.+)amount:(?P<amount>.+)`)

const (
	ResultWin  = "WIN"
	ResultLose = "LOSE"
	ResultPush = "PUSH"

	TypeSpread    = "SPREAD"
	TypeOverUnder = "OU"

	ChoiceOver  = "OVER"
	ChoiceUnder = "UNDER"
)

type Bet struct {
	Game   string
	Bet    string
	Amount int
}

func (bet Bet) String() string {
	return fmt.Sprintf("Game: %s, Bet: %s, Amount: %d", bet.Game, bet.Bet, bet.Amount)
}

func (bet Bet) IsSpreadBet() bool {
	return !strings.Contains(bet.Bet, "Over") && !strings.Contains(bet.Bet, "Under")
}

func (bet Bet) IsOver() bool {
	return !bet.IsSpreadBet() && strings.Contains(bet.Bet, "Over")
}

func (bet Bet) IsUnder() bool {
	return !bet.IsSpreadBet() && strings.Contains(bet.Bet, "Under")
}

type Game struct {
	Type   string
	Team1  string
	Team2  string
	Result string
}

type Result struct {
	Result string
	Game   Game
	Amount int
}

func (r Result) IsEqual(res Result) bool {
	return r.Game.Result == res.Game.Result &&
		r.Game.Team1 == res.Game.Team1 &&
		r.Game.Team2 == res.Game.Team2 &&
		r.Game.Type == res.Game.Type
}

// ParseTxNote parses the transaction note field. Covers these basic examples only
//
//	game: England (-1.5) vs Iran (+1.5) bet: Iran amount: 300
//	game: England vs Iran bet: Over 2.5 Goals amount: 300
func ParseTxNote(txNote string) (Bet, bool) {
	matches := txNoteRegexp.FindStringSubmatch(txNote)
	if matches == nil {
		return Bet{}, false
	}

	game := strings.TrimSpace(matches[txNoteRegexp.SubexpIndex("game")])
	bet := strings.TrimSpace(matches[txNoteRegexp.SubexpIndex("bet")])
	amount := strings.TrimSpace(matches[txNoteRegexp.SubexpIndex("amount")])

	amt, err := strconv.Atoi(amount)
	if err != nil {
		panic(err)
	}
	return Bet{Game: game, Bet: bet, Amount: amt}, true
}

func FindGame(games []Game, bet Bet) (Game, bool) {
	game := slice.Filter(games, func(g Game, _ int) bool {
		correctTeams := strings.Contains(bet.Game, g.Team1) && strings.Contains(bet.Game, g.Team2)
		if !correctTeams {
			return false
		}
		if bet.IsSpreadBet() {
			return g.Type == TypeSpread
		} else {
			return g.Type == TypeOverUnder
		}
	})
	if len(game) == 0 {
		fmt.Printf("skipping bet %s\n", bet.Game)
		return Game{}, false
	}
	if len(game) != 1 {
		panic("more than 1 matching game")
	}
	return game[0], true
}

func GetBetResult(bet Bet, game Game) Result {
	// handle pushes
	if game.Result == ResultPush {
		return Result{
			Result: ResultPush,
			Game:   game,
			Amount: bet.Amount,
		}
	}

	// handle spreads
	if bet.IsSpreadBet() {
		result := ResultLose
		if bet.Bet == game.Result {
			result = ResultWin
		}
		return Result{
			Result: result,
			Game:   game,
			Amount: bet.Amount,
		}
	}

	// handle over/under
	if game.Result == ChoiceUnder && bet.IsUnder() {
		return Result{
			Result: ResultWin,
			Game:   game,
			Amount: bet.Amount,
		}
	}
	if game.Result == ChoiceOver && bet.IsOver() {
		return Result{
			Result: ResultWin,
			Game:   game,
			Amount: bet.Amount,
		}
	}
	return Result{
		Result: ResultLose,
		Game:   game,
		Amount: bet.Amount,
	}
}

func MergeResults(results []Result) []Result {
	var out []Result
	for _, res := range results {
		var isEqual bool
		for i := range out {
			if res.IsEqual(out[i]) {
				isEqual = true
				out[i].Amount += res.Amount
			}
		}
		if !isEqual {
			out = append(out, res)
		}
	}
	return out
}
