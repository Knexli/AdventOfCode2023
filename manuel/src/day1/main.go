package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readInputFile() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]string, 0)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func main() {
	linesToParse := readInputFile()
	linesToParse = corrected(linesToParse)

	println(linesToParse)
	compile := regexp.MustCompile("\\D")
	sum := 0
	for _, s := range linesToParse {
		numbers := string(compile.ReplaceAll([]byte(s), make([]byte, 0)))
		if len(numbers) > 1 {
			number := string(numbers[0]) + string(numbers[len(numbers)-1])
			add, _ := strconv.Atoi(number)
			println(add)
			sum += add

		} else {
			number := string(numbers[0]) + string(numbers[0])
			add, _ := strconv.Atoi(number)
			println(add)
			sum += add
		}
	}

	println(sum)
}

func corrected(parse []string) []string {
	nums := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, num := range nums {
		compile := regexp.MustCompile(num)
		for i2, s := range parse {
			ranges := compile.FindAllIndex([]byte(s), -1)
			temp := s
			offset := 0
			for i3, _ := range ranges {
				index := ranges[i3][0]
				temp = temp[:index+offset+1] + strconv.Itoa(i+1) + temp[index+offset+1:]
				offset++
			}
			parse[i2] = temp
		}
	}
	return parse
}
