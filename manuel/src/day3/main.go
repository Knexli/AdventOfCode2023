package main

import (
	"common"
	"strconv"
	"unicode"
)

var (
	numbers = make([]number, 0)
	parts   = make([]part, 0)
)

func main() {
	lines := common.ReadInputFile()
	readNumbers(lines)
	adjacent := getAdjacentNumbers()
	sum := 0
	for _, n := range adjacent {
		sum += n.value
		println(n.value)
	}
	println(sum)
}

func getAdjacentNumbers() []number {
	nums := make([]number, 0)
	for _, n := range numbers {
		found := false
		for _, p := range parts {
			if !found && n.isAdjacent(p.x, p.y) {
				nums = append(nums, n)
				found = true
			}
		}
	}
	return nums
}

func readNumbers(lines []string) {
	num := make([]rune, 0)
	for y, line := range lines {
		for x, c := range line {
			if unicode.IsNumber(c) {
				num = append(num, c)
				nextX := x + 1
				if nextX > len(line)-1 || !unicode.IsNumber(rune(line[nextX])) {
					n := number{width: len(num)}
					n.setValue(string(num))
					n.setCoordinate(x-n.width+1, y)
					numbers = append(numbers, n)
					num = make([]rune, 0)
				}
			} else {
				if c != '.' {
					p := part{}
					p.setCoordinate(x, y)
					parts = append(parts, p)
				}
			}

		}

	}
}

type field interface {
	setValue(string)
	setCoordinate(int, int)
}
type part struct {
	x int
	y int
}

func (p *part) setCoordinate(x int, y int) {
	p.y = y
	p.x = x
}

func (p *part) setValue(string) {

}

type number struct {
	value int
	x     int
	y     int
	width int
}

func (n *number) setCoordinate(x int, y int) {
	n.y = y
	n.x = x
}

func (n *number) setWidth(s string) {
	n.width = len(s)
}

func (n *number) setValue(s string) {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	n.value = v
}

func (n *number) isAdjacent(x int, y int) bool {
	if n.x-1 > x || n.y-1 > y || n.x+n.width < x || n.y+1 < y {
		return false
	}
	return true

}
