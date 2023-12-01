package day01

import (
	"advent_2023/utils"
	"fmt"
	"strings"
	"unicode"
)

// Entry point for the day
func Run(filename string, part int) {
	input, err := utils.LoadFile(filename)
	if err != nil {
		fmt.Println("Error Loading Data:", err)
		return
	}
	switch part {
	case 1:
		part_a(&input)
	case 2:
		part_b(&input)
	default:
		panic("No part selected")
	}
}

func part_a(input *[]string) {
	var sum int = 0
	for _, value := range *input {
		sum += (findFirstAndLastInt(value))
	}
	fmt.Println(sum)
}

func part_b(input *[]string) {
	numMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	var sum int = 0
	for _, value := range *input {
		sum += (findFirstAndLastIntWithWords(value, numMap))
	}
	fmt.Println(sum)
}

func findFirstAndLastInt(input string) int {
	foundFirst := false
	var returnValue int = 0
	var digit int = 0
	for _, char := range input {
		if unicode.IsNumber(char) {
			digit = int(char - '0')
			if !foundFirst {
				returnValue += 10 * digit
				foundFirst = true
			}
		}
	}
	returnValue += digit
	return returnValue
}

func findFirstAndLastIntWithWords(input string, wordmap map[string]int) int {
	var firstValue int = 0
	var firstIndex int = len(input) + 1
	var lastValue int = 0
	var lastIndex int = -1
	for index, char := range input {
		if unicode.IsNumber(char) {
			digit := int(char - '0')
			if index < firstIndex {
				firstValue = digit
				firstIndex = index
			}
			if index > lastIndex {
				lastValue = digit
				lastIndex = index
			}
		}
	}
	for key, value := range wordmap {
		indexFirst := strings.Index(input, key)
		if indexFirst >= 0 && indexFirst < firstIndex {
			firstValue = value
			firstIndex = indexFirst
		}
		indexLast := strings.LastIndex(input, key)
		if indexLast >= 0 && indexLast > lastIndex {
			lastValue = value
			lastIndex = indexLast
		}
	}
	return firstValue*10 + lastValue
}
