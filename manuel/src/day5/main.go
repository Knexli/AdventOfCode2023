package main

import (
	"common"
	"strconv"
	"strings"
	"unicode"
)

var tables = make([][]RangeMap, 0)

type RangeMap struct {
	source      int
	desitnation int
	width       int
}

func newRangeMap(s string) RangeMap {

	parts := strings.Split(s, " ")

	source, err := strconv.Atoi(parts[0])
	desitnation, err := strconv.Atoi(parts[1])
	width, err := strconv.Atoi(parts[2])

	common.Check(err)

	return RangeMap{
		source:      source,
		desitnation: desitnation,
		width:       width,
	}
}

func readStartValues(s string) []int {
	numberstring := strings.TrimSpace(strings.Split(s, ":")[1])
	numbers := make([]int, 0)
	for _, v := range strings.Split(numberstring, " ") {
		num, err := strconv.Atoi(string(v))
		common.Check(err)
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {
	lines := common.ReadInputFile()
	seeds := readStartValues(lines[0])

	instructionType := -1

	for _, v := range lines[1:] {
		if strings.Contains(v, ":") {
			//TODO add unmapped numbers
			instructionType++
			tables = append(tables, make([]RangeMap,0 ))
		} else if len(v) > 0 && unicode.IsDigit(rune(v[0])) {
			r := newRangeMap(v)
			tables[instructionType] = append(tables[instructionType], r)
		}
	}
	println(seeds)
}
