package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SensorReading struct {
	Entries []int
	Child   *SensorReading
}

func (s *SensorReading) String() string {
	if s.Child == nil {
		return fmt.Sprintf("%v\n", s.Entries)
	}
	return fmt.Sprintf("%v\n%v\n", s.Entries, s.Child.String())
}

func (s *SensorReading) allZero() bool {
	allZero := true
	for _, entry := range s.Entries {
		if entry == 0 {
			continue
		} else {
			allZero = false
			break
		}
	}
	return allZero
}

func (s *SensorReading) MakeChild() {
	if s.allZero() {
		s.Child = nil
		return
	}

	s.Child = &SensorReading{Entries: []int{}}
	for i := 0; i < len(s.Entries)-1; i++ {
		root := s.Entries[i]
		next := s.Entries[i+1]
		s.Child.Entries = append(s.Child.Entries, next-root)
	}
	s.Child.MakeChild()
}

func (s *SensorReading) Predict() int {
	if s.allZero() {
		return 0
	}

	s.Entries = append(s.Entries, s.Entries[len(s.Entries)-1]+s.Child.Predict())
	return s.Entries[len(s.Entries)-1]
}

func (s *SensorReading) History() int {
	if s.allZero() {
		return 0
	}
	s.Entries = append([]int{s.Entries[0] - s.Child.History()}, s.Entries...)
	return s.Entries[0]
}

func Day9(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	sum := 0
	historySum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		sensor := SensorReading{}
		for _, entry := range strings.Split(line, " ") {
			value, err := strconv.Atoi(entry)
			if err != nil {
				panic(err)
			}
			sensor.Entries = append(sensor.Entries, value)
		}

		sensor.MakeChild()
		prediction := sensor.Predict()
		history := sensor.History()
		sum += prediction
		historySum += history
		fmt.Printf("Prediction: %d\n", prediction)
		fmt.Printf("History: %d\n")
		fmt.Println(sensor.String())
	}

	fmt.Println(sum)
	fmt.Println(historySum)
}
