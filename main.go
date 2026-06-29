//go:build !scratch
// +build !scratch

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
		Help: "launch game",
		Func: launchCMD,
	})

	shell.Run()
}

func launchCMD(c *ishell.Context) {
	var chosenRules ruleset
	var ruleIndex int

	approachRules := []string{string(initiativeBidding), string(noAproach)}
	ruleIndex = c.MultiChoice(approachRules, "Select an approach ruleset")
	switch ruleIndex {
	case 0:
		chosenRules.approach = initiativeBidding
	case 1:
		chosenRules.approach = noAproach
	}

	gameRules := []string{string(initiativeFIX), string(initiativeALT)}
	ruleIndex = c.MultiChoice(gameRules, "Select a ruleset")
	switch ruleIndex {
	case 0:
		chosenRules.initiative = initiativeFIX
	case 1:
		chosenRules.initiative = initiativeALT
	}

	fmt.Println("Input B1, B2")
	b1, _ := strconv.Atoi(c.ReadLine())
	b2, _ := strconv.Atoi(c.ReadLine())

	var p1 fencer
	p1.balance = b1
	var p2 fencer
	p2.balance = b2

	var moveKind moveTypeEnum
	var currentState gameState
	currentState.p1 = p1
	currentState.p2 = p2

	if chosenRules.approach == noAproach {
		fmt.Println("input attacker")
		attackerOptions := []string{"player1", "player2"}
		question := "choose attacker"
		attacker := c.MultiChoice(attackerOptions, question)

		switch attacker {
		case 0:
			p1.role = roleATK
			p2.role = roleDEF
			moveKind = moveTypeONEATK
		case 1:
			p1.role = roleDEF
			p2.role = roleATK
			moveKind = moveTypeTWOATK
		default:
			p1.role = roleERR
			p2.role = roleERR
			moveKind = moveTypeERR
		}

		currentState = gameState{
			p1:     p1,
			p2:     p2,
			parent: nil,
			kind:   stateEXCHANGE,
			rules:  chosenRules,
		}
	} else {
		currentState.kind = stateOOM
		currentState.rules = chosenRules
	}

	for true {
		switch currentState.kind {
		case stateWIN1:
			fmt.Println("Player 1 won the exchange")
			os.Exit(0)
		case stateWIN2:
			fmt.Println("player 2 has won the exchange")
			os.Exit(0)
		case stateOOM:
			if currentState.parent != nil {
				fmt.Println("players disengage")
			}
			if currentState.rules.approach == noAproach {
				os.Exit(0)
			}
		case stateERR:
			fmt.Println("malformed resulting gameState")
			os.Exit(1)
		case stateEXCHANGE:
			if currentState.parent != nil {
				fmt.Println("No hit. The exchange continues.")
			}
		}

		fmt.Println("Current State:")
		fmt.Printf("Player 1 has %d balance and is %s\n", currentState.p1.balance, currentState.p1.role)
		fmt.Printf("Pllayer 2 has %d balance and is %s\n", currentState.p2.balance, currentState.p2.role)
		fmt.Println(currentState.kind)

		fmt.Println("Input player 1 bid")
		bid1, err1 := strconv.Atoi(c.ReadPassword())
		fmt.Println("Input player 2 bid")
		bid2, err2 := strconv.Atoi(c.ReadPassword())
		if err1 != nil || err2 != nil {
			fmt.Println("malformed input")
			os.Exit(1)
		}
		fmt.Printf("Player 1 has bid %d, as %s\n", bid1, currentState.p1.role)
		fmt.Printf("Player 2 has bid %d as %s\n", bid2, currentState.p2.role)

		var currentMove move
		switch currentState.kind {
		case stateEXCHANGE:
			switch currentState.p1.role {
			case roleATK:
				moveKind = moveTypeONEATK
			case roleDEF:
				moveKind = moveTypeTWOATK
			}

			currentMove = move{
				kind: moveKind,
				bid1: bid1,
				bid2: bid2,
			}
		case stateOOM:
			currentMove.bid1 = bid1
			currentMove.bid2 = bid2
			currentMove.kind = moveTypeAPPROACH
		}

		var err error = nil
		currentState, err = getNextState(&currentState, currentMove)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)

		}
	}
	fmt.Println(currentState)

}
