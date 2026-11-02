# monkeytype-compare

A small CLI that compares the personal bests of two [Monkeytype](https://monkeytype.com) players across the 8 standard modes.

## Usage

```bash
go build *.go -o monkeytype-compare
./monkeytype-compare <player1> <player2>
```

Example output:

```
                   15s    30s    60s   120s     10     25     50    100
alice           142.00 135.40 128.10 120.55 160.00 150.20 140.30 132.00
bob             138.50 136.00 125.00 118.20 155.40 148.00 141.10 129.90
diff              3.50  -0.60   3.10   2.35   4.60   2.20  -0.80   2.10
'alice' wins by 16.45 wpm.
```

## What is compared

For each player, the best WPM is taken for:

- **Time:** 15, 30, 60, 120 seconds
- **Words:** 10, 25, 50, 100

A missing PB counts as 0.
The winner is the player with the higher sum of PBs across all 8 modes.

## Data source

Profiles are fetched from the public [Monkeytype API](https://api.monkeytype.com). Both usernames must exist and have public profiles.
