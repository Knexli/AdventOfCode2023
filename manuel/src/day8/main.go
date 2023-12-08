package main

import (
	"common"
	"math"
	"strings"
)

type node struct {
	left  string
	right string
}

func newNode(s string) (node, string) {
	parts := strings.Split(s, "=")
	key := parts[0][0:3]
	left := parts[1][2:5]
	right := parts[1][7:10]
	return node{left: left, right: right}, key
}

var nodes = make(map[string]node)

func main() {
	lines := common.ReadInputFile()
	instruction := lines[0]

	for i := 2; i < len(lines); i++ {
		node, key := newNode(lines[i])
		nodes[key] = node
	}

	stepCount := getStepCount(instruction, nodes, "AAA", "ZZZ")
	println(stepCount)

	stepCount = getParallelStepCount(instruction)

	println(stepCount)
}

func getParallelStepCount(instruction string) int {
	starts := make([]string, 0)
	for v := range nodes {
		if v[2] == 'A' {
			starts = append(starts, v)
		}
	}
	count := 1
	for _, start := range starts {
		res := getStepCount(instruction, nodes, start, "Z")
		count = int(lcm(float64(count), float64(res)))
	}

	return count
}

func lcm(num1 float64, num2 float64) float64 {
	if num1 == 0 || num2 == 0 {
		return 0
	}
	abs1 := math.Abs(num1)
	abs2 := math.Abs(num2)
	highest := math.Max(abs2, abs1)
	lowest := math.Min(abs1, abs2)
	lcm := highest
	for int(lcm)%int(lowest) != 0 {
		lcm += highest
	}
	return lcm
}

func getStepCount(instruction string, nodes map[string]node, start string, end string) int {
	count := 0
	current := start
	q := common.Queue[string]{}

	for _, c := range instruction {
		q.Enqueue(string(c))
	}
	for !q.Empty() {
		move, err := q.Dequeue()
		common.Check(err)

		count++
		switch move {
		case "R":
			current = nodes[current].right
		case "L":
			current = nodes[current].left
		}

		check := ""
		if current != "" {
			check = current[len(current)-len(end):]

		}

		if check == end {
			return count
		}

		if q.Size() == 1 {
			for _, c := range instruction {
				q.Enqueue(string(c))
			}
		}
	}

	return count
}
