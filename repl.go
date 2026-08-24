package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "Pokedex > "

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
		fmt.Printf("Your command was: %s\n", command)
	}

}
