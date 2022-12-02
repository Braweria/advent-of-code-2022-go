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

	scoreSystem := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3, "Lost": 0, "Draw": 3, "Won": 6}
	score := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		handShapes := strings.Fields(line)
		score += scoreSystem[shapes[handShapes[1]]]
		outcome := checkOutcome(shapes[handShapes[0]], shapes[handShapes[1]])
		score += scoreSystem[outcome]
	}

	fmt.Println(score)
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
