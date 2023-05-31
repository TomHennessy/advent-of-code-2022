package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	directorySizes := make(map[string]int)

	for directory := range directories {

		directorySizes[directory] = getFileSize(directories, directory)

	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range directorySizes {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		if directorySizes["/"]-kv.Value <= 40000000 {
			fmt.Printf("%s, %d\n", kv.Key, kv.Value)
			break
		}

	}

}

func getFileSize(directories map[string]map[string]int, directory string) int {

	var totalSize int

	for fileName, fileSize := range directories[directory] {

		if fileSize == 0 {
			totalSize += getFileSize(directories, fileName)
		} else {
			totalSize += fileSize
		}

	}

	return totalSize

}

func createDirectoryMap() map[string]map[string]int {
	file, err := os.Open("./input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	directories := make(map[string]map[string]int)

	currentIndex := ""

	currentIndexPointer := ""

	for scanner.Scan() {

		check(scanner.Err())

		thisLine := scanner.Text()

		if thisLine[:1] == "$" {

			if thisLine[2:4] == "cd" {

				if thisLine[5:] == ".." {

					currentIndexPointerSplit := strings.Split(currentIndexPointer, "/")

					currentIndexPointer = strings.Join(currentIndexPointerSplit[:len(currentIndexPointerSplit)-2], "/") + "/"

					continue
				}

				if thisLine[5:] != "/" {

					currentIndex = thisLine[5:]

					currentIndexPointer += currentIndex + "/"

				} else {

					currentIndexPointer = "/"

				}

				if len(directories[currentIndexPointer]) == 0 {

					directories[currentIndexPointer] = make(map[string]int)

				}

			}

			continue

		}

		if thisLine[0:3] == "dir" {

			directories[currentIndexPointer][currentIndexPointer+thisLine[4:]+"/"] = 0

			continue

		}

		splitLine := strings.Split(thisLine, " ")

		fileName := splitLine[1]

		fileSize, err := strconv.Atoi(splitLine[0])

		check(err)

		directories[currentIndexPointer][fileName] = fileSize

	}

	return directories
}
