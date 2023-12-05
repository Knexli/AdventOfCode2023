package main

import (
	"common"
	"fmt"
	"strconv"
	"strings"
)

type scratch struct {
	winningNumbers []int
	numbers        []int
	points         int
	amount         int
	id             int
	mentioned      int
}

func (s *scratch) print() {
	fmt.Printf("%d: mentioned: %d, amount: %d \n", s.id, s.mentioned, s.amount)
}

func newScratch(id int, line string) *scratch {
	s := scratch{
		winningNumbers: make([]int, 0),
		numbers:        make([]int, 0),
	}
	s.readNumbers(line)
	s.calculateScore()
	s.id = id
	return &s
}

func (s *scratch) calculateScore() {
	match := make([]int, 0)
	for _, winNum := range s.winningNumbers {
		for _, number := range s.numbers {
			if winNum == number {
				match = append(match, winNum)
				break
			}
		}
	}
	score := 0
	if len(match) > 1 {
		score = 1
		for i := 1; i < len(match); i++ {
			score *= 2
		}
	} else if len(match) == 1 {
		score = 1
	}
	s.points = score
	s.amount = len(match)
}

func (s *scratch) readNumbers(line string) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, "|")
	winNums := strings.Split(strings.TrimSpace(parts[0]), " ")
	for _, v := range winNums {
		if v == "" {
			continue
		}
		number, err := strconv.Atoi(v)
		check(err)
		s.winningNumbers = append(s.winningNumbers, number)
	}
	nums := strings.Split(strings.TrimSpace(parts[1]), " ")
	for _, v := range nums {
		if v == "" {
			continue
		}
		number, err := strconv.Atoi(v)
		check(err)
		s.numbers = append(s.numbers, number)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines := common.ReadInputFile()
	sum := 0
	scratches := make([]*scratch, 0)
	for i, line := range lines {
		parts := strings.Split(line, ":")
		s := newScratch(i+1, parts[1])
		scratches = append(scratches, s)
		sum += s.points
	}
	println(sum)

	sum = 0
	for i, s := range scratches {
		s.mentioned++
		sum += s.mentioned
		s.print()
		for offset := 1; offset <= s.amount; offset++ {
			if i+offset < len(scratches) {
				scratches[i+offset].mentioned += s.mentioned
			}
		}
	}
	println(sum)
}
