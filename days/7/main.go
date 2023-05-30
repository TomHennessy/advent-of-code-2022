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

	// directories := make(map[string]map[string]int)

	// fmt.Println("directories =", directories)

	// for dir, size := range directories {
	// 	// println(directory, size)

	// 	fmt.Println("dir", dir)
	// 	fmt.Println("size", size)
	// 	fmt.Println("----------------------")

	// }

	// b, err := json.Marshal(directories)

	// check(err)

	// fmt.Println("-----------")
	// fmt.Println("-----------")
	// fmt.Println("-----------")
	// fmt.Println(string(b))

	// os.Exit(0)

	// directories = [map[a: e:0 f:29116 g:2557 h.lst:62596]]
	// directories["a"] = make(map[string]int)
	// directories["a"]["e"] = 0
	// directories["a"]["f"] = 29116
	// directories["a"]["g"] = 2557
	// directories["a"]["h.lst"] = 62596

	// directories["e"] = make(map[string]int)
	// directories["e"]["i"] = 584

	// fmt.Println("directories =", directories)

	directorySizes := make(map[string]int)

	for directory := range directories {

		directorySizes[directory] = getFileSize(directories, directory)
		// fmt.Println("directorySizes", directory, directorySizes[directory])

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

	// previousKey := ""
	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		fmt.Println("not this--")
		if directorySizes["/"]-kv.Value <= 40000000 {
			fmt.Printf("%s, %d\n", kv.Key, kv.Value)
			fmt.Println("--")
			break
		}

		// previousKey = kv.Key
	}

	// fmt.Println("previousKey =", previousKey)
	// fmt.Println("directorySizes[previousKey] =", directorySizes[previousKey])
	// fmt.Println(directorySizes)
	fmt.Println("Current used space:", directorySizes["/"])

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

	// fmt.Println("totalSize =", totalSize)

	return totalSize

}

func createDirectoryMap() map[string]map[string]int {
	file, err := os.Open("./input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	directories := make(map[string]map[string]int)

	currentIndex := ""
	// currentIndexHash := 0

	currentIndexPointer := ""

	// var dirNameMap = make(map[string]string)

	for scanner.Scan() {

		check(scanner.Err())

		thisLine := scanner.Text()

		if thisLine[:1] == "$" {

			if thisLine[2:4] == "cd" {

				if thisLine[5:] == ".." {

					// fmt.Println("-----------------------")
					// fmt.Println("-----------------------")
					// fmt.Println("currentIndexPointer =", currentIndexPointer)

					// currentIndexPointer = currentIndexPointer[:strings.LastIndex(currentIndexPointer, "/")] + "/"

					currentIndexPointerSplit := strings.Split(currentIndexPointer, "/")

					currentIndexPointer = strings.Join(currentIndexPointerSplit[:len(currentIndexPointerSplit)-2], "/") + "/"

					// fmt.Println("currentIndexPointerNew =", currentIndexPointer)

					// fmt.Println("-----------------------")
					// fmt.Println("-----------------------")
					continue
				}

				if thisLine[5:] != "/" {

					// fmt.Println("--------------")
					// fmt.Println("--------------")
					// fmt.Println("--------------")
					// fmt.Println("--------------")
					// fmt.Println("--------------")
					// fmt.Println("directories[thisLine[5:]] == nil", directories[thisLine[5:]])

					// if len(directories[thisLine[5:]]) != 0 {

					currentIndex = thisLine[5:]

					// } else {

					// 	// currentIndex = currentIndexHash + "/" + thisLine[5:]

					// 	// current index should have a prefix of the currentIndexHash

					// 	currentIndex = strconv.Itoa(currentIndexHash) + "_" + thisLine[5:]

					// 	dirNameMap[thisLine[5:]] = ""
					// 	dirNameMap[thisLine[5:]] = currentIndex

					// 	currentIndexHash++

					// 	fmt.Println("--------------")
					// 	fmt.Println("--------------")
					// 	fmt.Println("--------------")
					// 	fmt.Println("--------------")
					// 	fmt.Println("--------------")
					// 	fmt.Println("currentIndex =", currentIndex)

					// }

					currentIndexPointer += currentIndex + "/"

					// if currentIndex == "gftgshl" {

					// 	fmt.Println("-----------")
					// 	fmt.Println("-----------")
					// 	fmt.Println("-----------")
					// 	fmt.Println("-----------")
					// 	fmt.Println("currentIndex =", currentIndex)
					// 	fmt.Println("directories =", directories)

					// 	// os.Exit(1)

					// }
				} else {

					currentIndexPointer = "/"

				}

				// fmt.Println("currentIndexPointer lalala::", currentIndexPointer)

				if len(directories[currentIndexPointer]) == 0 {

					directories[currentIndexPointer] = make(map[string]int)

				}

			}
			// else if thisLine[2:4] == "ls" {
			// 	continue
			// }

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

		// if currentIndex == "ddgtnw" {

		// 	fmt.Println("-----------")

		// 	fmt.Println("currentIndex =", currentIndex)
		// 	fmt.Println("fileName =", fileName)
		// 	fmt.Println("fileSize =", fileSize)
		// 	fmt.Println("directories =", directories)

		// 	fmt.Println("-----------")

		// 	// os.Exit(1)

		// }
	}

	return directories
}

// store the directories in a flat array, with the size of the files in each directory.

// key: directory name

// value: size of files in directory

// if the key is for a directory, add the size as 0

// -----

// when a size is requested for a directory, loop through the directory recursively and add the sizes of the files within and directories within.
