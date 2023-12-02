package day2

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day2(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}

	games := strings.Split(string(data), "\n")
	matcher := regexp.MustCompile("([0-9]+) (red|blue|green)")
	possibleGames := 0
	powerSum := 0
	for i, game := range games {
		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		roundsPossible := true
		maxColors := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, round := range rounds {
			m := matcher.FindAllStringSubmatch(round, -1)
			for _, cMatch := range m {
				count, _ := strconv.ParseInt(cMatch[1], 10, 64)
				colour := cMatch[2]
				if count > int64(bag[colour]) {
					roundsPossible = false
				}
				maxColors[colour] = int(math.Max(float64(maxColors[colour]), float64(int(count))))
			}
			if len(m) == 0 {
				panic("failed to parse round")
			}
		}
		roundSum := 1
		for _, max := range maxColors {
			roundSum *= max
		}
		powerSum += roundSum
		if roundsPossible {
			possibleGames += i + 1
		}
	}
	fmt.Println("Existing Possible Games", possibleGames)
	fmt.Println("Power Sum", powerSum)
}
