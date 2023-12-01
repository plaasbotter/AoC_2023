package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ConvertInputToNumbers(input *[]string) ([]int, error) {
	var numbers []int
	for _, value := range *input {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func LoadFile(filename string) ([]string, error) {
	file, err := os.Open("./puzzle_data/" + filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	// Executes at the end
	defer file.Close()

	var returnValue []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		returnValue = append(returnValue, scanner.Text())
	}

	return returnValue, nil
}
