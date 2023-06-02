package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	defer timer("main")()

	file, err := os.Open("../input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	treeArrays := make(map[string]map[int]map[int]int, 0)

	treeArrays["left"] = make(map[int]map[int]int, 0)
	treeArrays["top"] = make(map[int]map[int]int, 0)

	rowIndex := 0

	var arraysLength int

	for scanner.Scan() {

		thisLine := scanner.Text()

		check(scanner.Err())

		arraysLength = len(thisLine)

		treeArrays["left"][rowIndex] = make(map[int]int, arraysLength)

		for i, char := range strings.Split(thisLine, "") {

			if rowIndex == 0 {
				treeArrays["top"][i] = make(map[int]int, arraysLength)
			}

			charAsInt, err := strconv.Atoi(char)

			check(err)

			treeArrays["left"][rowIndex][i] = charAsInt

			treeArrays["top"][i][rowIndex] = charAsInt

		}

		rowIndex++

	}

	// now we want to check each index (x, y) and see what the four distances are that it can see

	highestDistance := 0

	for y := 0; y < arraysLength; y++ {

		for x := 0; x < arraysLength; x++ {

			distances := make(map[string]int, 0)

			distances["left"] = 0
			distances["right"] = 0
			distances["top"] = 0
			distances["bottom"] = 0

			// left
			for i := x - 1; i >= 0; i-- {
				distances["left"]++
				if treeArrays["left"][y][i] >= treeArrays["left"][y][x] {
					break
				}
			}

			// right
			for i := x + 1; i < arraysLength; i++ {
				distances["right"]++
				if treeArrays["left"][y][i] >= treeArrays["left"][y][x] {
					break
				}
			}

			// top
			for i := y - 1; i >= 0; i-- {
				distances["top"]++
				if treeArrays["top"][x][i] >= treeArrays["top"][x][y] {
					break
				}
			}

			// bottom
			for i := y + 1; i < arraysLength; i++ {
				distances["bottom"]++
				if treeArrays["top"][x][i] >= treeArrays["top"][x][y] {
					break
				}
			}

			thisTotalDistance := distances["left"] * distances["right"] * distances["top"] * distances["bottom"]

			if thisTotalDistance > highestDistance {
				highestDistance = thisTotalDistance

			}

		}

	}

	fmt.Println(highestDistance)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
