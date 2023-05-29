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

			printStacks("arrangedStacks", arrangedStacks)

			continue
		}

		if isStartingStacks {

			startingStacks = append(startingStacks, thisLine)

			continue

		}

		// from here, we're dealing with the operations

		processedStacks = performOperation(thisLine, processedStacks)

		// -----------------

		// then perform the operations on the stacks

	}

	// fmt.Println(startingStacks)

	// fmt.Println("startingStacks", startingStacks)

	printStacks("processedStacks", processedStacks)

	fmt.Println("answer::")
	fmt.Println("------------------")

	for _, stack := range processedStacks {
		fmt.Print(stack[0])
	}

	// fmt.Println("------------------")

	// fmt.Println("arrangedStacks")
	// fmt.Println("------------------")

	// for _, stack := range arrangedStacks {
	// 	fmt.Println(stack)
	// }

	// fmt.Println("------------------")

	// fmt.Println("startingStacks")
	// fmt.Println("------------------")

	// for _, stack := range startingStacks {
	// 	fmt.Println(stack)
	// }

	// fmt.Println("------------------")

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

	// move 3 from 9 to 7

	splitOperation := strings.Split(operation, " ")

	operationNumber, _ := strconv.Atoi(splitOperation[1])

	operationFrom, _ := strconv.Atoi(splitOperation[3])

	operationTo, _ := strconv.Atoi(splitOperation[5])

	fmt.Println("------------------")

	// fmt.Println("stacks", stacks)

	fmt.Println("operationNumber", operationNumber)

	fmt.Println("operationFrom", operationFrom)

	fmt.Println("operationTo", operationTo)

	fmt.Println("------------------")

	// then, perform the operation

	itemsMoved := 0

	// remove the items from the from stack and load them into the itemsToMove array
	// for i, item := range stacks[operationFrom-1] {
	for i := 0; i < len(stacks[operationFrom-1]); i++ {

		item := stacks[operationFrom-1][i]

		itemsMoved++

		// add the item to the to stack
		stacks[operationTo-1] = prependStr(stacks[operationTo-1], item)

		// remove the item from the stack
		stacks[operationFrom-1] = append(stacks[operationFrom-1][:i], stacks[operationFrom-1][i+1:]...)

		i--

		if itemsMoved == operationNumber {

			break

		}

	}

	fmt.Println("------------------")

	fmt.Println("stacks", stacks)

	// os.Exit(1)

	return stacks
}

func prependStr(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}

func printStacks(label string, stacks [][]string) {

	fmt.Println("------------------")

	fmt.Println(label)

	for _, stack := range stacks {

		fmt.Println(stack)

	}

}
