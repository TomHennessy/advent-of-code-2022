package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	currentElfTotal := 0

	var elfTotals []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if scanner.Text() == "" {
			elfTotals = append(elfTotals, currentElfTotal)
			currentElfTotal = 0
			continue
		}

		var current int
		fmt.Sscanf(scanner.Text(), "%d", &current)

		currentElfTotal += current

	}

	sort.Slice(elfTotals, func(i, j int) bool {
		return elfTotals[i] > elfTotals[j]
	})

	fmt.Println(elfTotals[0] + elfTotals[1] + elfTotals[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
