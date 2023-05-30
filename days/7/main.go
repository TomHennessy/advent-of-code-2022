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

func main() {

	directories := createDirectoryMap()

}

func createDirectoryMap() map[string]map[string]int {
	file, err := os.Open("./input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	directories := make(map[string]map[string]int)

	var currentIndex string

	for scanner.Scan() {

		check(scanner.Err())

		thisLine := scanner.Text()

		if thisLine[:1] == "$" {

			if thisLine[2:4] == "cd" {
				currentIndex = thisLine[5:]

				if currentIndex == ".." {
					continue
				}

				directories[currentIndex] = make(map[string]int)

			} else if thisLine[2:4] == "ls" {
				continue
			}

			continue

		}

		if thisLine[0:3] == "dir" {

			directories[currentIndex][thisLine[4:]] = 0

			continue

		}

		splitLine := strings.Split(thisLine, " ")

		fileName := splitLine[1]

		fileSize, err := strconv.Atoi(splitLine[0])

		check(err)

		directories[currentIndex][fileName] = fileSize

	}

	return directories
}

// store the directories in a flat array, with the size of the files in each directory.

// key: directory name

// value: size of files in directory

// if the key is for a directory, add the size as 0

// -----

// when a size is requested for a directory, loop through the directory recursively and add the sizes of the files within and directories within.
