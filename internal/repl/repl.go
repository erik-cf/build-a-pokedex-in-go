package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

const prompt = "Pokedex > "

type cliCommandCallback func(*state.PokedexConfig, ...string) error

type cliCommand struct {
	name        string
	description string
	callback    cliCommandCallback
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func StartRepl(config *state.PokedexConfig) {
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

		var args []string
		if len(words) == 1 {
			args = make([]string, 0)
		} else {
			args = words[1:]
		}
		err := c.callback(config, args...)
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
		"map": {
			name:        "map",
			description: "Get Next Location Area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous location area",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemons in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon from the Pokedex",
			callback:    commandInspect,
		},
	}
}
