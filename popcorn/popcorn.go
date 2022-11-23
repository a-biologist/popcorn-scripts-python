package popcorn

import (
	"regexp"
	"strings"
)

var txNoteRegexp = regexp.MustCompile(`game:(?P<game>.+)bet:(?P<bet>.+)amount:(?P<amount>.+)`)

type Bet struct {
	Game   string
	Bet    string
	Amount string
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

	return Bet{Game: game, Bet: bet, Amount: amount}, true
}
