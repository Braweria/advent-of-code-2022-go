package days

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func D01() {
	file := getFile("inputs/1.txt")
	defer file.Close()

	maxCarrying := 0
	sum := 0
	var carrying []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		char := scanner.Text()
		if char == "" {
			if sum > maxCarrying {
				maxCarrying = sum
			}
			carrying = append(carrying, sum)

			sum = 0
			continue
		}
		calories, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		sum += calories
	}
	carrying = append(carrying, sum)
	sort.Sort(sort.Reverse(sort.IntSlice(carrying)))
	fmt.Println(maxCarrying, carrying[0]+carrying[1]+carrying[2])
}
