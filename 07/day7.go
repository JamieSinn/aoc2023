package day7

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

func (h Hand) CardsValue() int {
	value := 0
	matcher := regexp.MustCompile("(A)|(K)|(Q)|(J)|(T)|(9)|(8)|(7)|(6)|(5)|(4)|(3)|(2)")
	matches := matcher.FindAllStringSubmatch(h.Cards, -1)
	fmt.Println(matches)
	return value
}

func Day7(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var ranked []Hand
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		bid, _ := strconv.Atoi(split[1])
		ranked = append(ranked, Hand{Cards: split[0], Bid: bid})
	}
	for i, hand := range ranked {
		fmt.Printf("File Index: %d, Cards: %s, Value: %d, Bid: %d\n", i, hand.Cards, hand.CardsValue(), hand.Bid)
	}
}
