package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "Pokedex > "

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		words := cleanInput(scanner.Text())
		var command string
		if len(words) == 0 {
			command = ""
		} else {
			command = words[0]
		}

		c, ok := getCommands()[command]
		if !ok {
			fmt.Printf("Command %s is not a valid command\n", command)
			continue
		}

		err := c.callback()
		if err != nil {
			fmt.Printf("Command errored with error: %v", err)
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}
