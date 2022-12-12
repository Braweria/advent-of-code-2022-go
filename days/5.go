package days

import (
	"bufio"
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
	file := getFileByNumber(5)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := [10]crateStack{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		col := 1
		for i := 1; i < len(line); i += 4 {
			char := string(line[i])

			_, err := strconv.Atoi(char)
			if err == nil {
				break
			}

			if char == string(' ') {
				col++
				continue
			}
			stacks[col] = append(crateStack{char}, stacks[col]...)
			col++
		}
	}

	guide := instructions{}

	for scanner.Scan() {
		line := scanner.Text()
		amount, from, to := getNumbers(line)
		guide.add(amount, from, to)
	}

	for _, item := range guide {
		for i := 0; i < item.amount; i++ {
			crate := stacks[item.startPosition].grabCrate()
			stacks[item.endPosition].addCrate(crate)
		}
	}

	result := ""

	for _, char := range stacks {
		if len(char) == 0 {
			continue
		}
		top := char.getTopCrate()
		result += top
	}

	fmt.Println("Day 05", result)
}

func panicCheck(err error) {
	if err != nil {
		panic(err)
	}
}
