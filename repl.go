package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mike-jacks/pokedexcli/command"
	"github.com/mike-jacks/pokedexcli/config"
)

func startRepl(config *config.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := command.CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		commands := command.GetCommandsMap()
		command, exists := commands[commandName]
		if exists {
			err := command.Callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
