package days

import (
	"bufio"
	"fmt"
	"strconv"
)

func D01() {
	file := getFile("inputs/1.txt")
	defer file.Close()

	maxCarrying := 0
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		char := scanner.Text()
		if char == "" {
			if sum > maxCarrying {
				maxCarrying = sum
			}
			sum = 0
			continue
		}
		calories, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		sum += calories
	}
	fmt.Println(maxCarrying)
}
