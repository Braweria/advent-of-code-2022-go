package days

import (
	"bufio"
	"fmt"
)

type packet struct {
	index          int
	chars          [4]string
	counter        int
	messageIndex   int
	message        [14]string
	messageCounter int
}

func (p *packet) addChar(char string) {
	if p.index > 3 {
		p.index = 0
	}
	p.chars[p.index] = char
	p.index += 1
	p.counter += 1
}

func (p *packet) addMessage(char string) {
	if p.messageIndex > 13 {
		p.messageIndex = 0
	}
	p.message[p.messageIndex] = char
	p.messageIndex += 1
	p.messageCounter += 1
}

func (p *packet) check() (bool, bool) {
	return hasDupesInFour(p.chars), hasDupesInFourteen(p.message)
}

func hasDupesInFour(arr [4]string) bool {
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

func hasDupesInFourteen(arr [14]string) bool {
	mp := map[string]int{}

	for _, v := range arr {
		if v == "" {
			continue
		}
		mp[v] += 1
	}

	if len(mp) == 14 {
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

	firstPacket := packet{0, [4]string{}, 0, 0, [14]string{}, 0}
	result1 := 0
	result2 := 0

	for scanner.Scan() {
		char := scanner.Text()
		firstPacket.addChar(char)
		firstPacket.addMessage(char)

		check4, check14 := firstPacket.check()

		if check4 && result1 == 0 {
			result1 = firstPacket.counter
		}

		if check14 && result2 == 0 {
			result2 = firstPacket.messageCounter
		}
	}

	fmt.Println("Day 06", result1, result2)
}
