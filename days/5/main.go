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

	// spacerModifier := 0.25

	arrangedStacks := make([][]string, (len(stacks[0])/4)+1)

	for _, stack := range stacks {

		// var thisStack []string

		// thisStack := make([]string, (int(float64(len(stack))+spacerModifier) / 4))

		// fmt.Println("stack", thisStack)
		// fmt.Println("thisStack len", len(thisStack))

		// for every n + 1 % 4, it's a stack.

		stackWithChar := stack + " "

		for j, _ := range stackWithChar {
			// if i is a multiple of 3, minus 1 for each stack we've already had, it's a stack

			// refIndex := j

			// if j == len(stack)-1 {
			// 	refIndex = j + 1
			// }

			fmt.Println("---")
			fmt.Print("----", stackWithChar, "----")
			fmt.Println("---")

			if (j+1)%4 == 0 { // || j == len(stack)-1
				// arrangedStacks[i] = stack[1:len(stack)-1]
				// thisStack = thisStack + string(char)
				// thisStack = append(thisStack, string(char-1))

				// arrangedStacks[(j+1)/4] = strings.Join(thisStack, "")

				// previousChar := stack[refIndex-2 : refIndex-1]

				// if j == 1 {
				// 	continue
				// }

				previousChar := string(stackWithChar[j-2])

				fmt.Println("char", previousChar)
				fmt.Println("j", j)
				fmt.Println("len(stackWithChar)", len(stackWithChar))
				// arrangedStacks[i][((j+1)/4)-1] = previousChar

				insertIndex := ((j + 1) / 4) - 1

				arrangedStacks[insertIndex] = append(arrangedStacks[insertIndex], previousChar)

				fmt.Println("arrangedStacks", arrangedStacks)

			}

			// previousChar = string(char)
			fmt.Println("------------------")

		}

		fmt.Println("char", string(stackWithChar[len(stackWithChar)-2]))

		// arrangedStacks[i][len(arrangedStacks)-1] = string(stack[len(stack)-2])

		// arrangedStacks[len(arrangedStacks)-1] = append(arrangedStacks[len(arrangedStacks)-1], string(stack[len(stack)-2]))

		// arrangedStacks = append(arrangedStacks, thisStack)
		// arrangedStacks[i] = thisStack

	}

	fmt.Println("arrangedStacks")
	fmt.Println("------------------")

	for _, stack := range arrangedStacks {
		fmt.Println(stack)
	}

	fmt.Println("------------------")

	return stacks
}
