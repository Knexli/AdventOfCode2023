package common

import (
	"bufio"
	"fmt"
	"os"
)

func Check(err error)  {
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
