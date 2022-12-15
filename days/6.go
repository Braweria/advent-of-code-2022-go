package days

import (
	"bufio"
	"fmt"
)

type packet struct {
	index   int
	chars   [4]string
	counter int
}

func (p *packet) addChar(char string) {
	if p.index > 3 {
		p.index = 0
	}
	p.chars[p.index] = char
	p.index += 1
	p.counter += 1
}

func (p *packet) check() bool {
	return hasDupes(p.chars)
}

func hasDupes(arr [4]string) bool {
	mp := map[string]int{}

	for _, v := range arr {
		if v == "" {
			continue
		}
		mp[v] += 1
	}

	if len(mp) == 4 {
		return true
	}
	return false
}

func includesElementInArray(arr [4]string, element string, skip int) bool {
	for index, item := range arr {
		if index == skip {
			continue
		}
		if item == element {
			return true
		}
	}
	return false
}

func D06() {
	file := getFileByNumber(6)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	firstPacket := packet{0, [4]string{}, 0}
	result1 := 0

	for scanner.Scan() {
		char := scanner.Text()
		firstPacket.addChar(char)
		check := firstPacket.check()

		if check {
			result1 = firstPacket.counter
			break
		}
	}

	fmt.Println("Day 06", result1)
}
