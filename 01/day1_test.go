package day1

import (
	"io"
	"os"
	"strings"
	"testing"
)

var lineValues = map[int]int{0: 12, 1: 38, 2: 15, 3: 77}

func TestDay1(t *testing.T) {
	file, err := os.OpenFile("calib.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		lineVal := evaluateLine(line)
		if lineValues[i] != lineVal {
			t.Errorf("index: %d did not equal expected - %d - actual %d", i, lineValues[i], lineVal)
		}
	}

}
