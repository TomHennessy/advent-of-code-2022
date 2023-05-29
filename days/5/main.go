package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var startingStacks []string
	var arrangedStacks [][]string
	var processedStacks [][]string
	isStartingStacks := true

	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		thisLine := scanner.Text()

		// first, loop through the file until there's a blank line

		// the last line of this can be ignored

		// -----------------

		if thisLine == "" {
			isStartingStacks = false
			// fmt.Println("blank line")
			startingStacks = startingStacks[:len(startingStacks)-1]

			// then format the stacks into appropriate data structures
			arrangedStacks = arrangeStacks(startingStacks)

			processedStacks = arrangedStacks

			continue
		}

		if isStartingStacks {

			startingStacks = append(startingStacks, thisLine)

			continue

		}

		// from here, we're dealing with the operations

		processedStacks = performOperation(thisLine, processedStacks)

	}

	fmt.Println("answer::")
	fmt.Println("------------------")

	for _, stack := range processedStacks {
		fmt.Print(stack[0])
	}

}

func arrangeStacks(stacks []string) [][]string {

	// the length of each stack is 4 (including a space at the end of each)
	arrangedStacks := make([][]string, (len(stacks[0])/4)+1)

	for _, stack := range stacks {

		// add the space so they're all uniform length and layout
		stackWithChar := stack + " "

		for j := range stackWithChar {

			if (j+1)%4 == 0 {

				previousChar := string(stackWithChar[j-2])

				if previousChar == " " {

					continue

				}

				insertIndex := ((j + 1) / 4) - 1

				arrangedStacks[insertIndex] = append(arrangedStacks[insertIndex], previousChar)

			}

		}

	}

	return arrangedStacks
}

func performOperation(operation string, stacks [][]string) [][]string {

	// first, determine which stack to take from and which to put on

	// example: move 3 from 9 to 7

	splitOperation := strings.Split(operation, " ")

	operationNumber, _ := strconv.Atoi(splitOperation[1])

	operationFrom, _ := strconv.Atoi(splitOperation[3])

	operationTo, _ := strconv.Atoi(splitOperation[5])

	// then, perform the operation

	itemsToMove := make([]string, 0)

	// remove the items from the from stack and add them to the to stack
	for i := 0; i < len(stacks[operationFrom-1]); i++ {

		item := stacks[operationFrom-1][i]

		itemsToMove = append(itemsToMove, item)

		if len(itemsToMove) == operationNumber {

			for j := operationNumber - 1; j >= 0; j-- {

				// add the item to the to stack
				stacks[operationTo-1] = prependStr(stacks[operationTo-1], itemsToMove[j])

				// remove the item from the stack
				stacks[operationFrom-1] = append(stacks[operationFrom-1][:j], stacks[operationFrom-1][j+1:]...)

			}

			break

		}

	}

	return stacks

}

func prependStr(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}
