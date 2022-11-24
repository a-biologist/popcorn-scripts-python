package popcorn_test

import (
	"github.com/stein-f/popcorn-scripts/popcorn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTxNote(t *testing.T) {
	cases := map[string]struct {
		gotTxNote string
		wantBet   popcorn.Bet
		wantIsBet bool
	}{
		"spread": {
			gotTxNote: "game: England (-1.5) vs Iran (+1.5) bet: Iran amount: 300",
			wantBet: popcorn.Bet{
				Game:   "England (-1.5) vs Iran (+1.5)",
				Bet:    "Iran",
				Amount: 300,
			},
			wantIsBet: true,
		},
		"over/under": {
			gotTxNote: "game: England vs Iran bet: Over 2.5 Goals amount: 120",
			wantBet: popcorn.Bet{
				Game:   "England vs Iran",
				Bet:    "Over 2.5 Goals",
				Amount: 120,
			},
			wantIsBet: true,
		},
		"not a bet line": {
			gotTxNote: "f214532kjg245g345h345h45h",
			wantBet:   popcorn.Bet{},
			wantIsBet: false,
		},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			bet, isBet := popcorn.ParseTxNote(test.gotTxNote)
			assert.Equal(t, test.wantBet, bet)
			assert.Equal(t, test.wantIsBet, isBet)
		})
	}
}
