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
			'X': scoreForLoss + shapeScissors, // elf: rock, me: loss
			'Y': scoreForTie + shapeRock,      // elf: rock, me: tie
			'Z': scoreForWin + shapePaper,     // elf: rock, me: win
		},
		'B': {
			'X': scoreForLoss + shapeRock,    // elf: paper, me: loss
			'Y': scoreForTie + shapePaper,    // elf: paper, me: tie
			'Z': scoreForWin + shapeScissors, // elf: paper, me: win
		},
		'C': {
			'X': scoreForLoss + shapePaper,   // elf: scissors, me: loss
			'Y': scoreForTie + shapeScissors, // elf: scissors, me: tie
			'Z': scoreForWin + shapeRock,     // elf: scissors, me: win
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
