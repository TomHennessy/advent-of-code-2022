package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	intersect "github.com/juliangruber/go-intersect/v2"
)

func main() {

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalPriority int

	var itemLists [][]string

	for scanner.Scan() {

		thisLine := scanner.Text()

		itemLists = append(itemLists, strings.Split(thisLine, ""))

		if len(itemLists) == 3 {

			sharedItem := getShared(itemLists) //intersections[0]

			totalPriority += getPriority(sharedItem[0])

			itemLists = [][]string{}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("totalPriority")
	fmt.Println(totalPriority)

}

func getShared(arr [][]string) string {
	// arr is an array of strings.

	intersections := arr[0] // initialize with first line

	for i := 1; i < len(arr); i++ {

		intersections = intersect.SimpleGeneric(intersections, arr[i])

	}

	return intersections[0]

}

func getPriority(char byte) int {

	if char >= 91 {

		return int(char - 96)

	} else {

		return int(char - 38)

	}

}

// every 3 lines is a new group
// the only value shared between all 3 lines is the badge number
