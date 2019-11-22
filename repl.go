package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Enums
type MetaCommandResult int
type PrepareResult int
type StatementType int

const (
	META_COMMAND_SUCCESS MetaCommandResult = iota
	META_COMMAND_UNRECOGNIZED_COMMAND
)
const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_UNRECOGNIZED_COMMAND
)
const (
	STATEMENT_INSERT StatementType = iota
	STATEMENT_SELECT
)

type Statement struct {
	stype StatementType
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printPrompt()
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n") // strip out the trailing \n

		if input[0] == '.' {
			switch doMetaCommand(input) {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGNIZED_COMMAND:
				fmt.Printf("Unrecognized command: '%v'\n", input)
				continue
			}
		}

		prepareResult, statement := prepareStatement(input)
		switch prepareResult {
		case PREPARE_SUCCESS:
			break
		case PREPARE_UNRECOGNIZED_COMMAND:
			fmt.Printf("Unrecognized keyword at start of '%v'\n", input)
			continue
		}

		executeStatement(statement)
		fmt.Println("Executed.")
	}
}

func printPrompt() {
	fmt.Print("db > ")
}

func doMetaCommand(input string) MetaCommandResult {
	var result MetaCommandResult
	if input == ".exit" {
		os.Exit(0)
	} else {
		result = META_COMMAND_UNRECOGNIZED_COMMAND
	}
	return result
}

func prepareStatement(input string) (PrepareResult, Statement) {
	// Both insert and select will have more items afterwards and will need to be parsed accordingly
	var statement Statement
	if input == "insert" {
		statement = Statement{stype: STATEMENT_INSERT}
		return PREPARE_SUCCESS, statement
	}
	if input == "select" {
		statement = Statement{stype: STATEMENT_SELECT}
		return PREPARE_SUCCESS, statement
	}
	return PREPARE_UNRECOGNIZED_COMMAND, statement
}

func executeStatement(statement Statement) {
	switch statement.stype {
	case STATEMENT_INSERT:
		fmt.Println("This is where we do an insert.")

	case STATEMENT_SELECT:
		fmt.Println("This is where we do a select.")
	}
}
