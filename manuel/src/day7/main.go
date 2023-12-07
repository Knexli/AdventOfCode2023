package main

import (
	"common"
	"strconv"
	"unicode"
)

type result struct {
	winType int
	cards   []int
	points  int
}

func findMaxPoints(cardNums []int) int {
	counter := make([]int, 14)
	for _, v := range cardNums {
		counter[v-1]++
	}

	maxSameCard := -1
	//maxCarType := -1

	for _, v := range counter {
		if v > maxSameCard {
			maxSameCard = v
			//	maxCarType = t
		}
	}
	return 0
}

func convertRuneToCardNumber(c rune) int {
	if unicode.IsNumber(c) {
		n, err := strconv.Atoi(string(c))
		common.Check(err)
		return n
	} else {
		switch c {
		case 'A':
			return 14
		case 'K':
			return 13
		case 'Q':
			return 12
		case 'J':
			return 11
		case 'T':
			return 10
		default:
			return 0
		}
	}

}

func main() {

}
