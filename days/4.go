package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type assignment struct {
	start, end int
}

type elfGroup struct {
	first, second assignment
}

func (group *elfGroup) isEngulfing() bool {
	first := group.first
	second := group.second

	if first.start <= second.start && first.end >= second.end {
		return true
	}

	if second.start <= first.start && second.end >= first.end {
		return true
	}

	return false
}

func (group *elfGroup) isOverlapping() bool {
	first := group.first
	second := group.second

	if first.end < second.start {
		return false
	}

	if second.end < first.start {
		return false
	}

	return true
}

func D04() {
	file := getFile("inputs/4.txt")

	scanner := bufio.NewScanner(file)

	engulfing := 0
	overlaps := 0

	for scanner.Scan() {
		group := elfGroup{}

		line := scanner.Text()
		elves := strings.Split(line, ",")

		for i, value := range elves {
			ranges := strings.Split(value, "-")
			start, startErr := strconv.Atoi(ranges[0])
			end, endErr := strconv.Atoi(ranges[1])

			if startErr != nil {
				panic(startErr)
			}

			if endErr != nil {
				panic(endErr)
			}

			if i == 0 {
				group.first.start = start
				group.first.end = end
				continue
			}
			group.second.start = start
			group.second.end = end
		}

		if group.isEngulfing() {
			engulfing++
		}

		if group.isOverlapping() {
			overlaps++
		}

	}

	fmt.Println("Day 04", engulfing, overlaps)
}
