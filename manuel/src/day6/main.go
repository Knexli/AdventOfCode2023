package main

import (
	"common"
	"math"
	"strconv"
	"unicode"
)

type race struct {
	timeLimit      int
	recordDistance int
}

func (r race) getWinningPossibilityCount() int {
	floor := mitternacht(float64(r.timeLimit), float64(r.recordDistance), false)
	ceil := mitternacht(float64(r.timeLimit), float64(r.recordDistance), true)
	floor = math.Floor(floor) + 1
	ceil = math.Ceil(ceil) - 1
	return int(math.Abs(ceil-floor)) + 1
}

func mitternacht(t float64, s float64, top bool) float64 {
	if top {
		return (-t - math.Sqrt(t*t-(4*s))) / -2
	} else {
		return (-t + math.Sqrt(t*t-4*s)) / -2
	}
}

func main() {
	lines := common.ReadInputFile()
	top := convertToNumArray(lines[0])
	bottom := convertToNumArray(lines[1])
	races := make([]race, 0)
	for i := 0; i < len(top); i++ {
		races = append(races, race{timeLimit: top[i], recordDistance: bottom[i]})
	}

	sum := 1
	for _, r := range races {
		sum *= r.getWinningPossibilityCount()
	}
	println("*********** part one")
	println(sum)

	t := justReadNums(lines[0])
	s := justReadNums(lines[1])
	race := race{timeLimit: t, recordDistance: s}
	println("*********** part two")
	println(race.getWinningPossibilityCount())
}

func justReadNums(s string) int {
	num := ""
	for _, c := range s {
		if unicode.IsNumber(c) {
			num += string(c)
		}
	}
	res, err := strconv.Atoi(num)
	common.Check(err)
	return res
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
