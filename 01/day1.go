package day1

import (
	"io"
	"os"
	"regexp"
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
	regexPattern := "(one|two|three|four|five|six|seven|eight|nine)"
	numberPositions := map[int]int{0: -1, 1: -1}
	firstNumber := 0
	lastNumber := 0
	firstFound := false
	sum := 0
	wordMatcher, err := regexp.Compile(regexPattern)
	if err != nil {
		panic(err)
	}

	for i, c := range line {
		if d, err := strconv.ParseInt(string(c), 10, 64); err == nil {
			if !firstFound {
				firstNumber = int(d)
				lastNumber = firstNumber
				firstFound = true
				numberPositions[0] = i
				numberPositions[1] = i
			} else {
				lastNumber = int(d)
				numberPositions[1] = i
			}
		}
	}

	wordMatches := wordMatcher.FindAllString(line, -1)
	workingLine := line
	for _, match := range wordMatches {
		n := -1
		switch match {
		case "one":
			n = 1
			break
		case "two":
			n = 2
			break
		case "three":
			n = 3
			break
		case "four":
			n = 4
			break
		case "five":
			n = 5
			break
		case "six":
			n = 6
			break
		case "seven":
			n = 7
			break
		case "eight":
			n = 8
			break
		case "nine":
			n = 9
			break
		default:
			panic("unknown word")
		}
		index := strings.Index(workingLine, match)
		if numberPositions[0] < 0 || index < numberPositions[0] {
			firstNumber = n
			numberPositions[0] = index
		} else if numberPositions[1] < 0 || index > numberPositions[1] {
			lastNumber = n
			numberPositions[1] = index
		}
		workingLine = strings.Replace(workingLine, match, strings.Repeat("_", len(match)), 1)
	}
	sum += firstNumber*10 + lastNumber
	return sum
}
