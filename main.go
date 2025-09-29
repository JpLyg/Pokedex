package main

import (
	"Pokedex/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	conf := &config{
		API: pokeapi.NewClient("https://pokeapi.co", 5*time.Minute),
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		line = strings.ToLower(strings.TrimSpace(line))
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd, ok := commands[parts[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(conf); err != nil {
			fmt.Println("Error:", err)
		}

		//fmt.Printf("Your command was: %s\n", parts[0])

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	Next     *string
	Previous *string
	API      *pokeapi.Client
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the map",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the map",
		callback:    commandMapb,
	},
}

func commandMap(conf *config) error {
	var page pokeapi.LocationAreaList
	var err error
	if conf.Next == nil {
		page, err = conf.API.GetLocationAreasFirstPage()
	} else {
		page, err = conf.API.GetLocationAreasByURL(*conf.Next)
	}
	if err != nil {
		return err
	}

	for _, r := range page.Results {
		fmt.Println(r.Name)
	}

	conf.Next = page.Next
	conf.Previous = page.Previous
	return nil
}

func commandMapb(conf *config) error {
	if conf.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	page, err := conf.API.GetLocationAreasByURL(*conf.Previous)
	if err != nil {
		return err
	}

	for _, r := range page.Results {
		fmt.Println(r.Name)
	}

	conf.Next = page.Next
	conf.Previous = page.Previous
	return nil
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}
