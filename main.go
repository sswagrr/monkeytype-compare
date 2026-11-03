package main

import (
	"errors"
	"fmt"
	"os"
)

func parseArgs() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", fmt.Errorf("usage: %s <player1> <player2>", os.Args[0])
	}
	return os.Args[1], os.Args[2], nil
}

type result struct {
	rec Records
	err error
}

func fetchAsync(name string) <-chan result {
	ch := make(chan result, 1)
	go func() {
		rec, err := FetchRecords(name)
		ch <- result{rec, err}
	}()
	return ch
}

func main() {
	player1, player2, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	ch1, ch2 := fetchAsync(player1), fetchAsync(player2)
	res1, res2 := <-ch1, <-ch2

	for _, r := range []struct {
		name string
		err  error
	}{{player1, res1.err}, {player2, res2.err}} {
		if errors.Is(r.err, ErrUserNotFound) {
			fmt.Printf("Username '%s' does not exist.\n", r.name)
			return
		} else if r.err != nil {
			fmt.Println("Error:", r.err)
			return
		}
	}

	DisplayRecords(res1.rec, res2.rec)
}
