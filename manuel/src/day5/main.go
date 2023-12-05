package main

import (
	"common"
	"math"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var tables = make([][]RangeMap, 0)

type RangeMap struct {
	source      int
	destination int
	width       int
}

func (m RangeMap) isInRange(x int) bool {
	if x >= m.source && x < m.source+m.width {
		return true
	}
	return false
}

func (m RangeMap) findDestination(x int) int {
	if m.isInRange(x) {
		return (x - m.source) + m.destination
	}
	return -1
}

func newRangeMap(s string) RangeMap {
	parts := strings.Split(s, " ")

	source, err := strconv.Atoi(parts[1])
	destination, err := strconv.Atoi(parts[0])
	width, err := strconv.Atoi(parts[2])

	common.Check(err)

	return RangeMap{
		source:      source,
		destination: destination,
		width:       width,
	}
}

func readStartValues(s string) []int {
	numberString := strings.TrimSpace(strings.Split(s, ":")[1])
	numbers := make([]int, 0)
	for _, v := range strings.Split(numberString, " ") {
		num, err := strconv.Atoi(v)
		common.Check(err)
		numbers = append(numbers, num)
	}
	return numbers
}

func getDestination(x int, level int) int {
	for _, m := range tables[level] {
		r := m.findDestination(x)
		if r != -1 {
			return r
		}
	}
	return x
}

func main() {
	lines := common.ReadInputFile()
	seeds := readStartValues(lines[0])

	instructionType := -1

	for _, v := range lines[1:] {
		if strings.Contains(v, ":") {
			instructionType++
			tables = append(tables, make([]RangeMap, 0))
		} else if len(v) > 0 && unicode.IsDigit(rune(v[0])) {
			r := newRangeMap(v)
			tables[instructionType] = append(tables[instructionType], r)
		}
	}
	res := make([]int, 0)

	for _, seed := range seeds {
		idx := seed
		level := 0
		for level != len(tables) {
			idx = getDestination(idx, level)
			//fmt.Printf("level: %d, idx: %d\n", level, idx)
			level++
		}
		res = append(res, idx)

	}
	lowest := math.MaxInt

	for _, v := range res {
		if v < lowest {
			lowest = v
		}
	}
	//	println(lowest)

	//flattenedMap := convertToFlatMap(tables)

	lowest = math.MaxInt
	l := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		i := i
		go func() {

			for x := seeds[i]; x < seeds[i]+seeds[i+1]; x++ {
				idx := x
				//println(x)
				level := 0
				for level != len(tables) {
					//fmt.Printf("level: %d, idx: %d\n", level, idx)
					idx = getDestination(idx, level)
					level++
				}
				if idx < lowest {
					l.Lock()
					lowest = idx
					l.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	println(lowest)
}
