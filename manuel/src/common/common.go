package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInputFile() []string {
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
func ConvertToNumArray(s string) []int {
	result := make([]int, 0)
	temp := ""
	reading := false
	for _, c := range s {
		if unicode.IsDigit(c) || c == '-' {
			if !reading {
				reading = true
			}
			temp += string(c)
		} else if reading {
			value, err := strconv.Atoi(temp)
			Check(err)
			result = append(result, value)
			reading = false
			temp = ""
		}
	}
	if reading {
		value, err := strconv.Atoi(temp)
		Check(err)
		result = append(result, value)
		reading = false
		temp = ""
	}
	return result
}
