package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// go from left to right, then right to left

// each pass, keep track of the current highest. At each tree, if the current tree is higher than the current highest, update an array with its index.

// do this for each row

func main() {

	file, err := os.Open("./input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// currentHighest := 0

	// var seenFromLeft []int[]int
	// seenFromLeft := make(map[int]map[int]int)

	// fileStats, err := file.Stat()

	// check(err)

	// for i := int64(0); i < int64(fileStats.Size()); i++ {

	// }

	// treeArrays := make([][]int, 4)

	// tree arrays should have 4 arrays, each containing arrats of ints. The 4 arrays are the 4 directions and have the keys left, right, top, bottom

	treeArrays := make(map[string]map[int]map[int]int, 0)

	treeArrays["left"] = make(map[int]map[int]int, 0)
	treeArrays["right"] = make(map[int]map[int]int, 0)
	treeArrays["top"] = make(map[int]map[int]int, 0)
	treeArrays["bottom"] = make(map[int]map[int]int, 0)

	rowIndex := 0

	for scanner.Scan() {

		thisLine := scanner.Text()

		check(scanner.Err())

		// fmt.Println(seenFromLeft)

		// thisLineLeft := make(map[int]int, 0)

		// left

		// right

		treeArrays["left"][rowIndex] = make(map[int]int, len(thisLine))
		treeArrays["right"][rowIndex] = make(map[int]int, len(thisLine))

		for i, char := range strings.Split(thisLine, "") {

			charAsInt, err := strconv.Atoi(char)

			check(err)

			// thisLineLeft[i] = charAsInt

			treeArrays["left"][rowIndex][i] = charAsInt
			treeArrays["right"][rowIndex][len(thisLine)-i-1] = charAsInt

		}

		rowIndex++

	}

	// now get the top down and bottom up arrays

	for i := 0; i < len(treeArrays["left"]); i++ {

		arrayLength := len(treeArrays["left"][i])

		for j := 0; j < arrayLength; j++ {

			if i == 0 {
				treeArrays["top"][j] = make(map[int]int, len(treeArrays["left"]))
				treeArrays["bottom"][j] = make(map[int]int, len(treeArrays["left"]))
			}

			height := treeArrays["left"][i][j]

			// fmt.Println(i, j, height)

			treeArrays["top"][j][i] = height

			treeArrays["bottom"][j][arrayLength-i-1] = height

		}

	}

	// we now have the left and right arrays for each row

}

func canBeSeen(array []int) []int {

	currentHighest := 0
	tallTrees := make([]int, 0)

	// for i := 0; i < len(thisLine); i++ {
	for index, height := range array {

		if height >= currentHighest {

			currentHighest = height

			tallTrees = append(tallTrees, index)

		}

	}

	return tallTrees
}
