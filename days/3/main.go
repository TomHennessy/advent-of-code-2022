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

	for scanner.Scan() {

		thisLine := scanner.Text()

		firstHalfItems := strings.Split(thisLine[:len(thisLine)/2], "")
		secondHalf := strings.Split(thisLine[len(thisLine)/2:], "")

		intersections := intersect.SimpleGeneric(firstHalfItems, secondHalf)

		sharedItem := intersections[0]

		priority := sharedItem[0]

		if priority >= 91 {

			priority = priority - 96

		} else {

			priority = priority - 38

		}

		totalPriority += int(priority)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("totalPriority")
	fmt.Println(totalPriority)

}

// func itemPriority(item string) int {

// 	var priority rune

// 	priority = rune(item[0])
// 	if priority >= 91 {
// 		priority = priority - 96
// 	} else {
// 		priority = priority - 38
// 	}

// 	return int(priority)
// }
