package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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
		fmt.Printf("Your command was: %s\n", parts[0])
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
		if parts[0] == "exit" {
			break
		}
	}

}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}
