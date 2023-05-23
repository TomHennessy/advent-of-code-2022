package main

import (
	"bufio"
	"log"
	"os"
)

const scoreForLoss = 0
const scoreForTie = 3
const scoreForWin = 6

const shapeRock = 1
const shapePaper = 2
const shapeScissors = 3

func main() {

	scoreSheet := map[byte]map[byte]int{
		'A': {
			'X': scoreForTie + shapeRock,      // elf: rock, me: rock
			'Y': scoreForWin + shapePaper,     // elf: rock, me: paper
			'Z': scoreForLoss + shapeScissors, // elf: rock, me: scissors
		},
		'B': {
			'X': scoreForLoss + shapeRock,    // elf: paper, me: rock
			'Y': scoreForTie + shapePaper,    // elf: paper, me: paper
			'Z': scoreForWin + shapeScissors, // elf: paper, me: scissors
		},
		'C': {
			'X': scoreForWin + shapeRock,     // elf: scissors, me: rock
			'Y': scoreForLoss + shapePaper,   // elf: scissors, me: paper
			'Z': scoreForTie + shapeScissors, // elf: scissors, me: scissors
		},
	}

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {

		thisLine := scanner.Text()

		elfHand := thisLine[0]
		myResponse := thisLine[2]

		totalScore += scoreSheet[elfHand][myResponse]

	}

	log.Println("total score is", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
