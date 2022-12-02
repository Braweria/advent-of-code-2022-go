package days

import (
	"bufio"
	"fmt"
	"strings"
)

func D02() {
	file := getFile("inputs/2.txt")
	defer file.Close()

	shapes := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors", "X": "Rock", "Y": "Paper", "Z": "Scissors"}
	updatedShapes := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors", "X": "Lost", "Y": "Draw", "Z": "Won"}

	scoreSystem := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3, "Lost": 0, "Draw": 3, "Won": 6}
	firstScore := 0
	secondScore := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		handShapes := strings.Fields(line)

		firstScore += scoreSystem[shapes[handShapes[1]]]
		outcome := checkOutcome(shapes[handShapes[0]], shapes[handShapes[1]])
		firstScore += scoreSystem[outcome]

		secondScore += scoreSystem[updatedShapes[handShapes[1]]]
		reverseOutcome := checkReverseOutcome(updatedShapes[handShapes[0]], updatedShapes[handShapes[1]])
		secondScore += scoreSystem[reverseOutcome]
	}

	fmt.Println(firstScore, secondScore)
}

func checkReverseOutcome(opponentShape string, outcome string) string {
	myShape := opponentShape
	if outcome == "Draw" {
		myShape = opponentShape
	}

	if outcome == "Lost" {
		switch opponentShape {
		case "Rock":
			myShape = "Scissors"
		case "Scissors":
			myShape = "Paper"
		case "Paper":
			myShape = "Rock"
		}
	}

	if outcome == "Won" {
		switch opponentShape {
		case "Rock":
			myShape = "Paper"
		case "Scissors":
			myShape = "Rock"
		case "Paper":
			myShape = "Scissors"
		}
	}

	return myShape
}

func checkOutcome(opponentShape string, myShape string) string {
	if myShape == opponentShape {
		return "Draw"
	}
	beatSystem := map[string]string{"Rock": "Scissors", "Scissors": "Paper", "Paper": "Rock"}

	wouldBeat := beatSystem[myShape]

	if wouldBeat == opponentShape {
		return "Won"
	}

	return "Lost"

	// switch {
	// case wouldBeat == opponentShape:
	// 	return "Won"
	// case myShape == opponentShape:
	// 	return "Draw"
	// case wouldBeat != opponentShape:
	// 	return "Lost"
	// }
}
