package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/GabrielPereira187/pokedexcli/initializers"
	"github.com/GabrielPereira187/pokedexcli/types"
)

func printCli(commands map[string]types.CliCommand) {
	for {
		fmt.Printf("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		if command , ok := commands[strings.Split(line, " ")[0]]; ok == true {
			words := strings.Fields(line)

			if line == "exit" {
				break
			}
			if err := command.Callback(words...); err != nil {
				fmt.Println(err.Error())
			}	
		}
	}
}

func main() {
	mapCommands := initializers.LoadCommands()
	printCli(mapCommands)


	fmt.Println("\n\n\nObrigado por usar essa CLI")
}