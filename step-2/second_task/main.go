package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {

	if lineNum <= 0 {
		return ""
	}

	f, err := os.Open(inputFilename)

	if err != nil {
		return ""
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentLine := 0

	for scanner.Scan() {

		if currentLine == lineNum {
			return scanner.Text()
		}

		currentLine++

	}

	return ""

}
