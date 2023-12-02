package day1

import (
	"io"
	"os"
	"strings"
	"testing"
)

var calibration1 = map[int]int{0: 12, 1: 38, 2: 15, 3: 77, 4: 22}
var calibration2 = map[int]int{0: 29, 1: 83, 2: 13, 3: 24, 4: 42, 5: 14, 6: 76, 7: 88, 8: 29, 9: 19, 10: 11, 11: 18, 12: 83, 13: 79, 14: 88, 15: 56}

func TestDay1(t *testing.T) {
	file, err := os.OpenFile("calib1.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		calibLine := &CalibrationLine{RawLine: line}
		calibLine.Evaluate()
		lineVal := calibLine.Sum()
		if calibration1[i] != lineVal {
			t.Errorf("index: %d (%s) did not equal expected - %d - actual %d", i, line, calibration1[i], lineVal)
		}
	}
}

func TestDay1_Words(t *testing.T) {
	file, err := os.OpenFile("calib2.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		calibLine := &CalibrationLine{RawLine: line}
		calibLine.Evaluate()
		lineVal := calibLine.Sum()
		if calibration2[i] != lineVal {
			t.Errorf("index: %d (%s) did not equal expected - %d - actual %d", i, line, calibration2[i], lineVal)
		}
	}
}
