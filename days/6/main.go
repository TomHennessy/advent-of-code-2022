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

	// fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(bytesToRead))

	getStartOfPacket(file)

	// fmt.Println("o2 = ", o2)

}

func getStartOfPacket(file *os.File) string {

	fileStats, err := file.Stat()

	check(err)

	for i := int64(0); i < int64(fileStats.Size()); i++ {

		// move along the file 6 bytes
		o2, err := file.Seek(i, 0)

		check(err)

		// read 2 bytes from the current position
		bytesToRead := make([]byte, 4)

		n2, err := file.Read(bytesToRead)

		check(err)

		// return string(bytesToRead)

		// charList := strings.Split(string(bytesToRead), "")

		if unique(string(bytesToRead)) {
			fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(bytesToRead))
			fmt.Println("Characters processed: ", i+4)
			return string(bytesToRead)
		}

	}

	return ""

	// return o2
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
