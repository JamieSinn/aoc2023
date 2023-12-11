package day8

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var nodes = make(map[string]*Node)
var root *Node
var flow string

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func Parse(line string) (n *Node) {
	n = &Node{}
	matcher := regexp.MustCompile("([A-Z])\\w+")
	matches := matcher.FindAllString(line, -1)
	fmt.Println(matches)
	if len(matches) != 3 {
		return nil
	}
	n.Name = matches[0]
	if n.Name == "AAA" {
		root = n
	}
	if val, ok := nodes[n.Name]; ok {
		n = val
	} else {
		nodes[n.Name] = n
	}

	left := matches[1]
	right := matches[2]

	if val, ok := nodes[left]; ok {
		n.Left = val
	} else {
		n.Left = &Node{Name: left}
		nodes[left] = n.Left
	}

	if val, ok := nodes[right]; ok {
		n.Right = val
	} else {
		n.Right = &Node{Name: right}
		nodes[right] = n.Right
	}
	return
}

func Day8(fileName string) (int, int) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if i == 0 {
			flow = line
			continue
		}
		if strings.Contains(line, "=") {
			Parse(line)
			continue
		}
	}

	part1 := evaluate(root, flow)
	fmt.Println(part1)

	steps := 0
	startNodes := make([]*Node, 0)
	for _, node := range nodes {
		if node.Name[2] == 'A' {
			startNodes = append(startNodes, node)
		}
	}

	fmt.Println(steps)

	if len(startNodes) < 3 {
		return steps, -1
	}
	steps2 := evaluate2(startNodes, flow)
	fmt.Println(steps2)
	return steps, steps2
}

func evaluate2(nodes []*Node, flow string) int {
	steps := map[string]int{}
	for _, node := range nodes {
		steps[node.Name] = evaluate(node, flow)
	}
	fmt.Println(steps)
	stepsSlice := make([]int, 0)
	for _, step := range steps {
		stepsSlice = append(stepsSlice, step)
	}
	return LCM(stepsSlice[0], stepsSlice[1], stepsSlice[2:]...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func evaluate(n *Node, flow string) int {
	steps := 0
	var next *Node
	var current = n
	for i := 0; i < len(flow); {
		steps++
		direction := string(flow[i])
		switch direction {
		case "L":
			next = current.Left
		case "R":
			next = current.Right
		}
		if next == nil {
			panic("cannot have nil node")
		}
		if next.Name[2] == 'Z' {
			return steps
		}
		current = next
		if i == len(flow)-1 {
			i = 0
		} else {
			i++
		}

	}
	return steps
}
