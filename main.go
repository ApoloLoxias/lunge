package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
)

/* main;go */
func main() {
	shell := ishell.New() //includees default exit, clear, help cmds

	shell.Println(`
      _                                 _   _ 
     | |                               | | | |
     | |    _   _  _ __    __ _   ___  | | | |
     | |   | | | ||  _ \  / _  | /   \ | | | |
     | |___| |_| || | | || |_| ||  |_/ |_| |_|
     |_____\___|_||_| |_| \___ | \___| (_) (_)
                           __/ |              
                          |___/
	`)

	shell.AddCmd(&ishell.Cmd{
		Name: "fence",
		Help: "fence",
		Func: fenceCMD,
	})

	shell.Run()
}

func fenceCMD(c *ishell.Context) {
	fmt.Println("Input B1, B2")
	b1, _ := strconv.Atoi(c.ReadLine())
	b2, _ := strconv.Atoi(c.ReadLine())

	fence(b1, b2, c)
}

func fence(b1, b2 int, c *ishell.Context) {
	p1 := Fencer{
		Balance: b1,
		RoW:     true,
	}
	p2 := Fencer{
		Balance: b2,
		RoW:     false,
	}

	hit := hitFAL
	for hit == hitFAL {
		a, d := func(c *ishell.Context) (int, int) {
			fmt.Println("input a")
			a, _ := strconv.Atoi(c.ReadLine())

			fmt.Println("input b")
			d, _ := strconv.Atoi(c.ReadLine())

			return a, d
		}(c)
		fmt.Printf("a,d = %v,%v", a, d)
		fmt.Println()

		currenthit, err := Strike(&p1, &p2, a, d)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		hit = currenthit
		fmt.Println(hit)
	}
}

/* fencer.go */

type Fencer struct {
	Balance int
	RoW     bool
}

type hitEnum string

const (
	hitONE hitEnum = "p1 hits p2"
	hitTWO hitEnum = "p2 hits p1"
	hitFAL hitEnum = "no hit"
	hitDIS hitEnum = "disengage"
)

func Strike(o, d *Fencer, atk, def int) (hitEnum, error) {
	if [2]bool{o.RoW, d.RoW} != [2]bool{true, false} {
		return hitFAL, errors.New("invalid row pairing")
	}
	if atk > o.Balance || def > d.Balance {
		return hitFAL, errors.New("invalid balance spenditure")
	}

	if atk == 0 && def == 0 {
		return hitDIS, nil
	}
	if atk > def {
		return hitONE, nil
	}

	o.Balance -= atk
	d.Balance -= def

	if o.Balance == 0 && d.Balance == 0 {
		return hitDIS, nil
	}
	if o.Balance == 0 && d.Balance == 0 {
		return hitDIS, nil
	}
	if o.Balance == 0 { //d.Balance != 0 is implied
		return hitTWO, nil
	}
	if d.Balance == 0 { //o.Balance !=0 is implied
		return hitONE, nil
	}
	return hitFAL, nil
}
