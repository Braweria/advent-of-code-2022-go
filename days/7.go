package days

import (
	"bufio"
	"fmt"
)

type item struct {
	name     string
	size     int
	children children
	category string
}

func (i *item) addChild(name string, size int, children children, category string) {
	(*i).children[name] = item{name, size, children, category}
}

type children map[string]item

func (d *children) getSumOfBelow(below int) int {
	//

	return 0
}

func D07() {
	file := getFileByNumber(7)
	defer file.Close()

	sumOfDirs := 0

	root := item{"/", 0, children{}, "dir"}
	dirs := children{"/": root}

	current := &root
	parent := &root

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

	}

	fmt.Println("Day 07")
}
