package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var startingStacks []string
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
			startingStacks = arrangeStacks(startingStacks)

			continue
		}

		if isStartingStacks {

			startingStacks = append(startingStacks, thisLine)

		}

		// -----------------

		// then perform the operations on the stacks

	}

	// fmt.Println(startingStacks)

	// fmt.Println("startingStacks", startingStacks)

	fmt.Println("startingStacks")
	fmt.Println("------------------")

	for _, stack := range startingStacks {
		fmt.Println(stack)
	}

	fmt.Println("------------------")

}

func arrangeStacks(stacks []string) []string {

	spacerModifier := 0.25

	arrangedStacks := make([][]string, len(stacks))

	for i, stack := range stacks {

		// var thisStack []string

		thisStack := make([]string, (int(float64(len(stack))+spacerModifier) / 4))

		fmt.Println("stack", thisStack)
		fmt.Println("thisStack len", len(thisStack))

		// for every n + 1 % 4, it's a stack.

		for j, _ := range stack {
			// if i is a multiple of 3, minus 1 for each stack we've already had, it's a stack
			if (j+1)%4 == 0 {
				// arrangedStacks[i] = stack[1:len(stack)-1]
				// thisStack = thisStack + string(char)
				// thisStack = append(thisStack, string(char-1))

				// arrangedStacks[(j+1)/4] = strings.Join(thisStack, "")

				previousChar := stack[j-2 : j-1]

				fmt.Println("char", previousChar)
				thisStack[((j+1)/4)-1] = previousChar

			}

			// previousChar = string(char)

		}

		fmt.Println("char", string(stack[len(stack)-2]))

		thisStack[len(arrangedStacks)-1] = string(stack[len(stack)-2])

		// arrangedStacks = append(arrangedStacks, thisStack)
		arrangedStacks[i] = thisStack

	}

	fmt.Println("arrangedStacks")
	fmt.Println("------------------")

	for _, stack := range arrangedStacks {
		fmt.Println(stack)
	}

	fmt.Println("------------------")

	return stacks
}
