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

	for scanner.Scan() {
		line := scanner.Text()

		compartments := splitInTwo(line)

		duplicate := getDuplicateItem(compartments[0], compartments[1])
		runeDupe := []rune(duplicate)[0]

		prio := getPriority(runeDupe)

		sum += prio
	}

	fmt.Println("Day 03", sum)
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
