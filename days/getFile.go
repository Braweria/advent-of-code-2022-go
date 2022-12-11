package days

import (
	"log"
	"os"
	"strconv"
)

func getFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func getFileByNumber(number int) *os.File {
	filename := "inputs/" + strconv.Itoa(number) + ".txt"

	return getFile(filename)
}
