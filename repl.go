package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printPrompt()
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n") // strip out the trailing \n

		if input == ".exit" {
			return
		} else {
			fmt.Printf("Unrecognized command: %v\n", input)
		}
	}
}

func printPrompt() {
	fmt.Print("db > ")
}
