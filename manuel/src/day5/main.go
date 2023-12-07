package main

import (
	"common"
	"strconv"
	"unicode"
)

type race struct {
	timeLimit      int
	recordDistance int
}

func (receiver race) getWinningPossibilityCount() int {
}

func main() {
	lines := common.ReadInputFile()
	top := convertToNumArray(lines[0])
	bottom := convertToNumArray(lines[1])
	races := make([]race, 0)
	for i := 0; i < len(top); i++ {
		races = append(races, race{timeLimit: top[i], recordDistance: bottom[i]})
		println(races[i].timeLimit)
	}

}

func convertToNumArray(s string) []int {
	result := make([]int, 0)
	temp := ""
	reading := false
	for _, c := range s {
		if unicode.IsDigit(c) {
			if !reading {
				reading = true
			}
			temp += string(c)
		} else if reading {
			value, err := strconv.Atoi(temp)
			common.Check(err)
			result = append(result, value)
			reading = false
			temp = ""
		}
	}
	if reading {
		value, err := strconv.Atoi(temp)
		common.Check(err)
		result = append(result, value)
		reading = false
		temp = ""
	}
	return result
}
