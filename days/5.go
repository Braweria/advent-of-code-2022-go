package days

import (
	"fmt"
	"strconv"
	"strings"
)

type crateStack []string

func (crates *crateStack) grabCrate() string {
	crate := (*crates)[len(*crates)-1]
	*crates = (*crates)[:len(*crates)-1]
	return crate
}

func (crates *crateStack) addCrate(crate string) {
	*crates = append(*crates, crate)
}

func (crates *crateStack) getTopCrate() string {
	return (*crates)[len(*crates)-1]
}

type instruction struct {
	amount, startPosition, endPosition int
}

type instructions []instruction

func (inst *instructions) add(amount, start, end int) {
	instruction := instruction{amount, start, end}
	*inst = append(*inst, instruction)
}

func getNumbers(line string) (int, int, int) {
	split := strings.Split(line, " ")
	nums := []int{}

	for _, char := range split {
		num, err := strconv.Atoi(char)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}

	return nums[0], nums[1], nums[2]
}

func D05() {
	// file := getFileByNumber(5)
	// defer file.Close()

	guides := instructions{}

	fmt.Println(guides)
	n1, n2, n3 := getNumbers("move 1 from 2 to 1")
	guides.add(n1, n2, n3)
	fmt.Println(guides)

	fmt.Println("Day 05")
}

func panicCheck(err error) {
	if err != nil {
		panic(err)
	}
}
