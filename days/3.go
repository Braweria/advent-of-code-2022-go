package days

import (
	"bufio"
	"fmt"
	"regexp"
	"unicode"
)

func D03() {
	file := getFile("inputs/3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	badgeSum := 0
	group := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		compartments := splitInTwo(line)
		duplicate := getDuplicateItem(compartments[0], compartments[1])
		runeDupe := []rune(duplicate)[0]
		prio := getPriority(runeDupe)

		sum += prio

		group = append(group, line)
		if len(group) == 3 {
			badge := getBadge(group)
			badgePrio := getPriority(badge)
			badgeSum += badgePrio

			group = []string{}
		}
	}

	fmt.Println("Day 03", sum, badgeSum)
}

func splitInTwo(line string) []string {
	splits := []string{}
	splits = append(splits, line[:len(line)/2])
	splits = append(splits, line[len(line)/2:])
	return splits
}

func getDuplicateItem(left, right string) string {
	regex := regexp.MustCompile("[" + left + "]")
	matches := regex.FindStringSubmatch(right)
	return matches[0]
}

func getPriority(char rune) int {
	switch {
	case unicode.IsLower(char):
		return int(char-'a') + 1
	default:
		return int(char-'A') + 27
	}
}

func getBadge(group []string) rune {

	firstBag := group[0]
	secondBag := group[1]
	thirdBag := group[2]

	for _, item := range firstBag {
		if includesElement(secondBag, item) {
			if includesElement(thirdBag, item) {
				return item
			}
		}
	}
	panic("No badge found")
}

func includesElement(arr string, element rune) bool {
	for _, item := range arr {
		if item == element {
			return true
		}
	}
	return false
}
