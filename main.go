package main

import (
	"fmt"
	"os"
)

func parseArgs() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", fmt.Errorf("usage: %s <player1> <player2>", os.Args[0])
	}
	return os.Args[1], os.Args[2], nil
}

func main() {
	var exist bool
	var err error
	var player1, player2 string

	player1, player2, err = parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	exist, err = UsernameExists(player1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if !exist {
		fmt.Printf("Username '%s' does not exist.\n", player1)
		return
	}

	exist, err = UsernameExists(player2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if !exist {
		fmt.Printf("Username '%s' does not exist.\n", player2)
		return
	}

	rec1, err := FetchRecords(player1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rec2, err := FetchRecords(player2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	DisplayRecords(rec1, rec2)
}
