package main

import (
	"fmt"
	"os"
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

	startOfPacket := getFirstNonRepeating(file, 0, 4)

	startOfMessage := getFirstNonRepeating(file, startOfPacket, 14)

	fmt.Println("startOfMessage =", startOfMessage)

}

func getFirstNonRepeating(file *os.File, start int64, length int64) int64 {

	fileStats, err := file.Stat()

	check(err)

	for i := start; i < int64(fileStats.Size()); i++ {

		file.Seek(i, 0)

		check(err)

		bytesToRead := make([]byte, length)

		file.Read(bytesToRead)

		check(err)

		if unique(string(bytesToRead)) {
			return i + length
		}

	}

	return -1 // not found, allows for error checking

}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
