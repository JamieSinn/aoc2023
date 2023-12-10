package day6

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day6(fileName string) {
	file, err := os.ReadFile(fileName)
	matcher := regexp.MustCompile("\\d+")
	if err != nil {
		panic(err)
	}

	for _, file := range []string{strings.ReplaceAll(string(file), " ", ""), string(file)} {

		lines := strings.Split(string(file), "\n")

		times, records := matcher.FindAllString(lines[0], -1), matcher.FindAllString(lines[1], -1)

		marginProduct := 1
		for i, t := range times {
			raceTime, _ := strconv.Atoi(t)
			record, _ := strconv.Atoi(records[i])
			margin := len(optimizeRace(raceTime, record))
			fmt.Println(margin)
			marginProduct *= margin
		}
		fmt.Println(marginProduct)
	}
}

func optimizeRace(time int, record int) []int {
	timeDistanceMap := map[int]int{}
	var recordBreakers []int
	for i := 0; i < time; i++ {
		timeDistanceMap[i] = i * (time - i)
		if timeDistanceMap[i] > record {
			recordBreakers = append(recordBreakers, i)
		}
	}
	return recordBreakers
}
