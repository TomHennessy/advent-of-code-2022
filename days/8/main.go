package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("./input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	treeArrays := make(map[string]map[int]map[int]int, 0)

	treeArrays["left"] = make(map[int]map[int]int, 0)
	treeArrays["right"] = make(map[int]map[int]int, 0)
	treeArrays["top"] = make(map[int]map[int]int, 0)
	treeArrays["bottom"] = make(map[int]map[int]int, 0)

	rowIndex := 0

	var arraysLength int

	for scanner.Scan() {

		thisLine := scanner.Text()

		check(scanner.Err())

		arraysLength = len(thisLine)

		treeArrays["left"][rowIndex] = make(map[int]int, arraysLength)
		treeArrays["right"][rowIndex] = make(map[int]int, arraysLength)

		for i, char := range strings.Split(thisLine, "") {

			if rowIndex == 0 {
				treeArrays["top"][i] = make(map[int]int, arraysLength)
				treeArrays["bottom"][i] = make(map[int]int, arraysLength)
			}

			charAsInt, err := strconv.Atoi(char)

			check(err)

			treeArrays["left"][rowIndex][i] = charAsInt
			treeArrays["right"][rowIndex][arraysLength-i-1] = charAsInt

			treeArrays["top"][i][rowIndex] = charAsInt
			treeArrays["bottom"][i][arraysLength-rowIndex-1] = charAsInt

		}

		rowIndex++

	}

	// we now have the left and right, top and bottom arrays for each row
	// I then check each index against each of the other lists to see if it can be seen from any direction

	treesThatCanBeSeen := make(map[string]map[int]map[int]int, 0)

	treesThatCanBeSeen["left"] = make(map[int]map[int]int, 0)
	treesThatCanBeSeen["right"] = make(map[int]map[int]int, 0)
	treesThatCanBeSeen["top"] = make(map[int]map[int]int, 0)
	treesThatCanBeSeen["bottom"] = make(map[int]map[int]int, 0)

	for i := 0; i < arraysLength; i++ {

		treesThatCanBeSeen["left"][i] = canBeSeen(treeArrays["left"][i])
		treesThatCanBeSeen["right"][i] = canBeSeen(treeArrays["right"][i])
		treesThatCanBeSeen["top"][i] = canBeSeen(treeArrays["top"][i])
		treesThatCanBeSeen["bottom"][i] = canBeSeen(treeArrays["bottom"][i])

	}

	// then check the lists against each other
	totalCanBeSeen := 0

	for y := 0; y < arraysLength; y++ {

		for x := 0; x < arraysLength; x++ {

			reverseX := arraysLength - x - 1
			reverseY := arraysLength - y - 1

			if treesThatCanBeSeen["left"][y][x] == 1 ||
				treesThatCanBeSeen["right"][y][reverseX] == 1 ||

				treesThatCanBeSeen["top"][x][y] == 1 ||
				treesThatCanBeSeen["bottom"][x][reverseY] == 1 {

				totalCanBeSeen++

			}

		}

	}

	fmt.Println(totalCanBeSeen)

}

func canBeSeen(array map[int]int) map[int]int {

	// this cuts down any trees that can't be seen, marking them as 1

	currentHighest := -1
	canBeSeenList := make(map[int]int, len(array))

	for index := 0; index < len(array); index++ {

		height := array[index]

		if height > currentHighest {
			currentHighest = height

			canBeSeenList[index] = 1

		} else {
			canBeSeenList[index] = 0
		}

	}

	return canBeSeenList
}
