package day1

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func Day1(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	sum := 0
	for _, line := range lines {
		sum += evaluateLine(line)
	}
	println(sum)
}

func evaluateLine(line string) int {
	firstNumber := 0
	lastNumber := 0
	firstFound := false
	sum := 0

	for _, c := range line {
		if d, err := strconv.ParseInt(string(c), 10, 64); err == nil {
			if !firstFound {
				firstNumber = int(d)
				lastNumber = firstNumber
				firstFound = true
			} else {
				lastNumber = int(d)
			}
		}
	}
	sum += firstNumber*10 + lastNumber
	return sum
}
