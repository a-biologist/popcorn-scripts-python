# popcorn-scripts

Hacky scripts to parse popcorn lounge transactions and output stats

# Prerequisites

The script depends on Golang being installed on your machine. You can install from here: https://go.dev/dl/

# Scripts

## Count bets on teams

`go run scripts/bets/main.go`

Output

```text
Germany: 35
Japan: 17
Over: 22
Under: 20
```

# Run tests

```bash
make test
```
