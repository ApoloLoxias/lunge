package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/manifoldco/promptui"
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
		Name: "pp",
		Help: "prompts user to type in something and returns it to stdout",
		Func: ppCmd,
	})

	shell.Run()
}

func ppCmd(c *ishell.Context) {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    ">>>",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Malformed prompt: %v", err)
		fmt.Println()
		return
	}

	fmt.Println(result)
}
