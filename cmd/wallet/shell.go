package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mattn/go-colorable"

	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/utils"
	"github.com/zetamatta/go-readline-ny"
	"github.com/zetamatta/go-readline-ny/simplehistory"
)

// ShellSetup starts up a shell.
func ShellSetup(wallet *core.Wallet) {
	// We want to be able to save the history.
	history := simplehistory.New()

	// Create the new readline editor.
	editor := readline.Editor{
		Prompt: func() (int, error) {
			return fmt.Print("> ")
		},
		Writer:  colorable.NewColorableStdout(),
		History: history,
	}

	fmt.Printf("%s Wallet. Ctrl-D to quit\n", utils.Name)

	for {
		// Get the next command.
		text, err := editor.ReadLine(context.Background())

		if err != nil {
			fmt.Printf("ERR=%s\n", err.Error())
			return
		}

		// Separate the input into multiple fields so that we can process each command.
		fields := strings.Fields(text)

		if len(fields) <= 0 {
			continue
		}

		// Just add an echo here.
		fmt.Println(text)

		history.Add(text)
	}
}
