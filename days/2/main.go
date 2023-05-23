package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const scoreForLoss = 0
const scoreForTie = 3
const scoreForWin = 6

const shapeRock = 1
const shapePaper = 2
const shapeScissors = 3

func getOutcome(elfHand, myResponse string) int {

	var score int

	if elfHand == "A" {
		if myResponse == "X" {
			// tie
			score = scoreForTie + shapeRock
		} else if myResponse == "Y" {
			// win
			score = scoreForWin + shapePaper
		} else if myResponse == "Z" {
			// loss
			score = scoreForLoss + shapeScissors
		}
	} else if elfHand == "B" {
		if myResponse == "X" {
			// loss
			score = scoreForLoss + shapeRock
		} else if myResponse == "Y" {
			// tie
			score = scoreForTie + shapePaper
		} else if myResponse == "Z" {
			// win
			score = scoreForWin + shapeScissors
		}
	} else if elfHand == "C" {
		if myResponse == "X" {
			// win
			score = scoreForWin + shapeRock
		} else if myResponse == "Y" {
			// loss
			score = scoreForLoss + shapePaper
		} else if myResponse == "Z" {
			// tie
			score = scoreForTie + shapeScissors
		}
	}

	return score
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {

		thisLine := scanner.Text()
		// thisLine looks like this:
		// A X

		elfHand := strings.Split(thisLine, " ")[0]
		myResponse := strings.Split(thisLine, " ")[1]

		totalScore += getOutcome(elfHand, myResponse)

		// var current int
		// fmt.Sscanf(scanner.Text(), "%d", &current)

		// col 1 is the hand of the elf
		// A = rock
		// B = paper
		// C = scissors

		// col 2 is my response
		// X = rock
		// Y = paper
		// Z = scissors

		// scores:

		// shape:
		// 1 for rock
		// 2 for paper
		// 3 for scissors

		// outcomes:
		// 0 for loss
		// 3 for tie
		// 6 for win

	}

	log.Println("total score is", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
