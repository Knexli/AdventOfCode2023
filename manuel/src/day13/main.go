package main

import (
	"common"
	"strconv"
	"strings"
)

type ash struct {
	field  []int64
	fieldR []int64
}

func newAsh(lines []string) *ash {
	field := make([]int64, len(lines))
	fieldR := make([]int64, len(lines[0]))
	for y := 0; y < len(lines); y++ {
		lines[y] = strings.Replace(lines[y], "#", "1", -1)
		lines[y] = strings.Replace(lines[y], ".", "0", -1)
		num, err := strconv.ParseInt(lines[y], 2, 64)
		if err != nil {
			panic(err)
		}
		field[y] = num
	}
	for x := 0; x < len(lines[0]); x++ {
		currentNum := ""
		for y := 0; y < len(lines); y++ {
			currentNum += string(lines[y][x])
		}
		num, err := strconv.ParseInt(currentNum, 2, 64)
		if err != nil {
			panic(err)
		}
		fieldR[x] = num
	}

	return &ash{
		field:  field,
		fieldR: fieldR,
	}
}

func (a ash) findHorizontalMirror() (above int, found bool) {
	return getMirrorIndex(a.field)
}

func getMirrorIndex(field []int64) (int, bool) {
	for middleLine := 1; middleLine < len(field); middleLine++ {
		dStart := middleLine
		dEnd := len(field) - middleLine
		dCheck := min(dStart, dEnd)
		allTheSame := true
		//println("############## Starting from: ", middleLine, " #####################")
		for distance := 0; distance < dCheck; distance++ {
			same := field[middleLine-distance-1] == field[middleLine+distance]
			//	println(field[middleLine-distance-1], " == ", field[middleLine+distance], " := ", same)
			//	println(middleLine-distance-1, " <--> ", middleLine+distance, "    :Max=", len(field))
			if !same {
				allTheSame = false
			}
		}
		if allTheSame {
			return middleLine, true
		}
	}
	return -1, false
}

func (a ash) findVerticalMirror() (left int, found bool) {
	return getMirrorIndex(a.fieldR)
}

func main() {
	lines := common.ReadInputFile()
	fields := make([]ash, 0)
	currentField := make([]string, 0)

	for _, line := range lines {
		if len(line) == 0 {
			fields = append(
				fields,
				*newAsh(currentField),
			)
			currentField = make([]string, 0)
		} else {
			currentField = append(currentField, line)
		}
	}
	fields = append(
		fields,
		*newAsh(currentField),
	)
	currentField = make([]string, 0)
	sum := 0
	for _, field := range fields {
		num1, v := field.findVerticalMirror()
		num2, h := field.findHorizontalMirror()
		if v {
			sum += num1
		} else if h {
			sum += num2 * 100
		} else {
			panic("ahhhhhhhhhhhhhhhhhhhh")
		}
	}
	println(sum)
}
