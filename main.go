package main

import (
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
)

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
