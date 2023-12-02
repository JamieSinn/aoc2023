package day1

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type CalibrationFile struct {
	Lines []CalibrationLine
}

func (c *CalibrationFile) FromFile(f *os.File) {
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	c.Lines = make([]CalibrationLine, len(lines))
	for i, line := range lines {
		c.Lines[i] = CalibrationLine{RawLine: line}
	}
}

func (c *CalibrationFile) Sum() (sum int) {
	for _, l := range c.Lines {
		l.Evaluate()
		sum += l.Sum()
	}
	return
}

type CalibrationLine struct {
	RawLine string
	First   *Coordinate
	Last    *Coordinate
}

func (c *CalibrationLine) RemoveWords() (ret string) {
	repl := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")
	for ret = c.RawLine; ret != repl.Replace(ret); ret = repl.Replace(ret) {
	}
	return
}

func (c *CalibrationLine) Evaluate() {
	if c.RawLine == "" {
		panic("failed to evaluate line. Raw line value is not set.")
	}

	c.First = &Coordinate{Index: -1, Number: -1}
	c.Last = &Coordinate{Index: -1, Number: -1}
	firstFound := false

	for i, ch := range c.RemoveWords() {
		if d, parseErr := strconv.ParseInt(string(ch), 10, 64); parseErr == nil {
			if !firstFound {
				c.First.Number = int(d)
				c.First.Index = i
				c.Last.Number = int(d)
				c.Last.Index = i
				firstFound = true
			} else {
				c.Last.Number = int(d)
				c.Last.Index = i
			}
		}
	}
}

func (c *CalibrationLine) Sum() int {
	if c.First == nil || c.Last == nil {
		c.Evaluate()
	}
	return 10*c.First.Number + c.Last.Number
}

type Coordinate struct {
	Number int
	Index  int
}

func Day1(fileName string) {
	calib := &CalibrationFile{}
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	calib.FromFile(file)
	fmt.Println(calib.Sum())
}
