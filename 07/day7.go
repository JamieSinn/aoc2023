package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Hand struct {
	Cards []string
	Bet   int
	Type  int
}

func (hand *Hand) HandType(part2 bool) int {
	//Return 7: 5 of a kind
	//Return 6: 4 of a kind
	//Return 5: Full House
	//Return 4: 3 of a kind
	//Return 3: 2 pair
	//Return 2: 1 pair
	//return 1: High card

	dict := make(map[string]int)
	for _, num := range hand.Cards {
		dict[num] = dict[num] + 1
	}
	numJs := dict["J"]

	if part2 {
		if numJs == 5 {
			return 7
		}
		delete(dict, "J")
	}

	keys := maps.Keys(dict)
	maxMapKey := keys[0]

	for _, key := range keys {
		if dict[key] > dict[maxMapKey] {
			maxMapKey = key
		}
	}

	if part2 {
		dict[maxMapKey] += numJs
	}

	if len(keys) == 1 {
		hand.Type = 7
		return 7
	} else if len(keys) == 2 {
		if dict[keys[0]] == 4 || dict[keys[1]] == 4 {
			hand.Type = 6
			return 6
		} else if (dict[keys[0]] == 3 && dict[keys[1]] == 2) || (dict[keys[0]] == 2 && dict[keys[1]] == 3) {
			hand.Type = 5
			return 5
		}
	} else if len(keys) == 3 {
		if dict[keys[0]] == 3 || dict[keys[1]] == 3 || dict[keys[2]] == 3 {
			hand.Type = 4
			return 4
		} else if (dict[keys[0]] == 2 && dict[keys[1]] == 2) || (dict[keys[0]] == 2 && dict[keys[2]] == 2) || (dict[keys[1]] == 2 && dict[keys[2]] == 2) {
			hand.Type = 3
			return 3
		}
	} else if len(keys) == 4 {
		hand.Type = 2
		return 2
	}
	hand.Type = 1
	return 1
}
func evaluate(hands []Hand, part2 bool) {

	sort.Slice(hands, func(i, j int) bool {
		handIType := hands[i].HandType(part2)
		handJType := hands[j].HandType(part2)
		//A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
		cardMap := map[string]string{}
		if !part2 {
			cardMap = map[string]string{
				"A": "14",
				"K": "13",
				"Q": "12",
				"J": "11",
				"T": "10",
				"9": "9",
				"8": "8",
				"7": "7",
				"6": "6",
				"5": "5",
				"4": "4",
				"3": "3",
				"2": "2",
			}
		} else {
			cardMap = map[string]string{
				"A": "14",
				"K": "13",
				"Q": "12",
				"T": "10",
				"9": "9",
				"8": "8",
				"7": "7",
				"6": "6",
				"5": "5",
				"4": "4",
				"3": "3",
				"2": "2",
				"J": "1",
			}
		}

		if handIType == handJType {
			for index := 0; index < 5; index++ {
				currentCardIString, success := cardMap[hands[i].Cards[index]]
				if !success {
					currentCardIString = hands[i].Cards[index]
				}
				currentCardI, _ := strconv.Atoi(currentCardIString)

				currentCardJString, success := cardMap[hands[j].Cards[index]]
				if !success {
					currentCardJString = hands[j].Cards[index]
				}
				currentCardJ, _ := strconv.Atoi(currentCardJString)

				if currentCardI == currentCardJ {
					continue
				} else if currentCardI < currentCardJ {
					return true
				}
				return false
			}
		}
		return handIType < handJType
	})

	totalScore := 0
	for i, hand := range hands {
		totalScore += hand.Bet * (i + 1)
	}

	fmt.Println(totalScore)
}
func Day7(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make([]Hand, 0)
	for scanner.Scan() {
		currentHand := parseLine(scanner.Text())
		hands = append(hands, currentHand)
	}

	evaluate(hands, false)
	evaluate(hands, true)
}

func parseLine(line string) Hand {
	lineContentsSlice := strings.Split(line, " ")
	currentHand := Hand{
		Cards: make([]string, 0),
		Bet:   0,
	}

	currentBet, _ := strconv.Atoi(lineContentsSlice[1])
	currentHand.Bet = currentBet

	currentCards := strings.Split(lineContentsSlice[0], "")
	currentHand.Cards = append(currentHand.Cards, currentCards...)

	return currentHand
}
