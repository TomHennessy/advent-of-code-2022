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

	totalMatches := 0

	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		thisLine := scanner.Text()

		lineSegments := strings.Split(thisLine, ",")

		var previousSegment []int

		for i := 0; i < len(lineSegments); i++ {

			thisLineSegment := getSections(lineSegments[i])

			if len(previousSegment) == 0 {
				previousSegment = thisLineSegment
				continue
			}

			segmentsAreTheSame := aFullyContainsB(thisLineSegment, previousSegment)

			if segmentsAreTheSame {
				fmt.Println("thisLineSegment")
				fmt.Println(thisLineSegment)
				fmt.Println("previousSegment")
				fmt.Println(previousSegment)
				fmt.Println("-----")

				totalMatches++
			}

			previousSegment = thisLineSegment
		}

	}

	fmt.Println("totalMatches")
	fmt.Println(totalMatches)
}

func getSections(rangeString string) []int {
	// this will get the sections from the min and max

	var sections []int

	splitRange := strings.Split(rangeString, "-")

	rangeMin, err := strconv.Atoi(splitRange[0])

	if err != nil {
		log.Fatal(err)
	}

	rangeMax, err := strconv.Atoi(splitRange[1])

	if err != nil {
		log.Fatal(err)
	}
	// splitRange := []int{strings.Split(rangeString, "-")[0], strings.Split(rangeString, "-")[0]}

	for i := rangeMin; i <= rangeMax; i++ {
		sections = append(sections, i)
	}

	return sections
}

func aFullyContainsB(a, b []int) bool {
	// this will check if a fully contains b

	if a[0] <= b[0] && a[len(a)-1] >= b[len(b)-1] {
		return true
	}

	if a[0] >= b[0] && a[len(a)-1] <= b[len(b)-1] {
		return true
	}

	return false
}

func slicesAreEqual(a, b []int) bool {
	// this will check if two slices are equal

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			// fmt.Println("not equal")
			return false
		}
	}

	return true
}
