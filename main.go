package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	conf := &config{}
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
	next     string
	previous string
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
	var target_url string
	if conf.next != "" {
		target_url = conf.next
	} else {
		target_url = "https://pokeapi.co/api/v2/location-area"
	}
	resp, err := http.Get(target_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var poke_maps LocationAreaResponse

	err = json.Unmarshal(content, &poke_maps)
	if err != nil {
		return err
	}

	for _, item := range poke_maps.Results {
		fmt.Println(item.Name)
	}

	conf.next = poke_maps.Next
	conf.previous = poke_maps.Previous

	return nil

}

func commandMapb(conf *config) error {
	var target_url string
	if conf.previous != "" {
		target_url = conf.previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := http.Get(target_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var poke_maps LocationAreaResponse

	err = json.Unmarshal(content, &poke_maps)
	if err != nil {
		return err
	}

	for _, item := range poke_maps.Results {
		fmt.Println(item.Name)
	}

	conf.next = poke_maps.Next
	conf.previous = poke_maps.Previous

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
