package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

var visitedPositions map[string]int

// var headPositions map[string]int
// var tailPositions map[string]int

func main() {
	defer timer("main")()

	file, err := os.Open("../input.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	visitedPositions = make(map[string]int, 0) // this should just be appended to with the key being the position and the value being the number of times it's been visited

	headPositions := map[string]int{"x": 0, "y": 0}
	tailPositions := map[string]int{"x": 0, "y": 0}

	count := 0

	printPositions(headPositions, tailPositions)

	for scanner.Scan() {

		thisLine := scanner.Text()

		check(scanner.Err())

		moveRope(thisLine, headPositions, tailPositions)

		fmt.Println("line: ", count)

		// if len(visitedPositions) == 12 {
		// 	fmt.Println(len(visitedPositions))
		// 	fmt.Println(visitedPositions)
		// 	os.Exit(0)
		// }

		count++

	}

	fmt.Println(len(visitedPositions))

}

func moveRope(instruction string, headPositions map[string]int, tailPositions map[string]int) {

	direction, distance := string(instruction[0]), string(instruction[2:])

	distanceAsInt, err := strconv.Atoi(distance)

	check(err)

	// this should be in a loop so the tail can move when it needs to

	for i := 0; i < distanceAsInt; i++ {

		headPositions["x"], headPositions["y"] = stepPosition(direction, headPositions)

		tailPositions["x"], tailPositions["y"] = stepTail(direction, headPositions, tailPositions)
		visitedPositions[fmt.Sprintf("%d,%d", tailPositions["x"], tailPositions["y"])]++

		// fmt.Println("head", headPositions, "tail", tailPositions)

		fmt.Println("direction", direction)
		printPositions(headPositions, tailPositions)

	}

	// return headPositions["x"], headPositions["y"]

}

func printPositions(headPositions map[string]int, tailPositions map[string]int) {
	fmt.Println("head", headPositions, "tail", tailPositions)
	fmt.Println("--------")
}

func stepPosition(direction string, positions map[string]int) (int, int) {

	switch direction {
	case "U":
		positions["y"] += 1
	case "D":
		positions["y"] -= 1
	case "R":
		positions["x"] += 1
	case "L":
		positions["x"] -= 1
	}

	return positions["x"], positions["y"]

}

func stepTail(direction string, headPositions map[string]int, tailPositions map[string]int) (int, int) {

	// this is the clever bit which will need to figure out how the tail moves

	// if the tail is gapped behind the head, in the direction of travel, move it in the same direction as the head

	// ----

	// the tail should go behind the head. Unless:

	// the head is diagonal to the tail but headX - tailX == 1 and headY - tailY == 1 (or -1)

	// switch direction {
	// case "U":
	// 	tailPositions["y"] += 1
	// case "D":
	// 	tailPositions["y"] -= 1
	// case "R":
	// 	tailPositions["x"] += 1
	// case "L":
	// 	tailPositions["x"] -= 1
	// }

	distanceX := headPositions["x"] - tailPositions["x"]
	distanceY := headPositions["y"] - tailPositions["y"]

	if (distanceX == 1 &&
		distanceY == 1) ||
		(distanceX == -1 &&
			distanceY == -1) {
		// tailPositions

		// they're diagonal to each other

		fmt.Println("1")

		return tailPositions["x"], tailPositions["y"]

		// switch direction {
		// case "U":
		// 	return stepPosition("D", headPositions)
		// case "D":
		// 	return stepPosition("U", headPositions)
		// case "R":
		// 	return stepPosition("L", headPositions)
		// case "L":
		// 	return stepPosition("R", headPositions)
		// }
	}

	// if headPositions["x"] == tailPositions["x"] && headPositions["y"] == tailPositions["y"] {
	if (distanceX == 1 ||
		distanceY == 1) ||
		(distanceX == -1 ||
			distanceY == -1) {

		if distanceX > 1 || distanceX < -1 || distanceY > 1 || distanceY < -1 {

			tailPositions["x"], tailPositions["y"] = headPositions["x"], headPositions["y"]

			// they've moved too far apart, so move the tail to behind the head in the direction of travel

			fmt.Println("2")

			switch direction {
			case "U":
				return stepPosition("D", tailPositions)
			case "D":
				return stepPosition("U", tailPositions)
			case "R":
				return stepPosition("L", tailPositions)
			case "L":
				return stepPosition("R", tailPositions)
			}
		}

		// they're in the same position, so don't move the tail

		fmt.Println("3")

		return tailPositions["x"], tailPositions["y"]
	}

	if distanceX == 0 && distanceY == 0 {
		// they're in the same position, so don't move the tail

		fmt.Println("4")

		return tailPositions["x"], tailPositions["y"]

	}

	fmt.Println("5")

	return stepPosition(direction, tailPositions)

}

// determine how far apart the two positions are
// if they're more than one move apart, move the tail

// # # # # #
// # # # # #
// # T # # #
// H # # # #
// # # # # #
