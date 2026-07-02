//go:build auction
// +build auction

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abiosoft/ishell"
)

func main() {
	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd{
		Name: "launch",
		Func: launchCMD,
	})

	shell.Run()
}

func launchCMD(c *ishell.Context) {
	fmt.Println("Input starting balance")
	sb, _ := strconv.Atoi(c.ReadLine())
	fmt.Println("Input number of goods to be auctioned")
	ng, _ := strconv.Atoi(c.ReadLine())

	var p1, p2 fencer
	p1.balance, p2.balance = sb, sb

	moveKind := moveTypeAUCTION
	var currentState gameState
	currentState.p1 = p1
	currentState.p2 = p2
	currentState.facultativeState[numGoodsToBeAuctioned] = ng
	currentState.facultativeState[score] = 0

	for true {
		switch currentState.kind {
		case stateWIN1:
			fmtPrintln("Player 1 won the auction")
		case stateWIN2:
			fmt.Println("Player 2 has won the exchange")
		case stateOOM:
			os.Exit(1)
		case stateERR:
			fmt.Println("malformed gameState")
			os.Exit(1)
		case stateExchange:
			fmt.Printf("%v goods left. Next good\n")
		}

		fmt.Println("Current State:")
		fmt.Printf("P1 balance: %d, P2 balance: %d, %d goods to be auctioned\n", currentState.p1.balance, currentState.p2.balance, currentState.numGoodsToBeAuctionedValue())

		fnt.Println("Input p1 bid")
		bid1, err1 := strconv.Atoi(c.ReadPassword())
		fmt.Println("Input p2 bid")
		bid2, err2 := strconv.Atoi(c.ReadPassword())
		if err1 != nil || err2 != nil {
			fmt.Printf("malformed input b1 err '%v', b2 err '%v'", err1, err2)

			var currentMove move
			currentMove.bid1 = bid1
			currentMove.bid2 = bid2

			var err error
			currentState, err = getNextState(&currentState, currentMove)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		}
	}
}
