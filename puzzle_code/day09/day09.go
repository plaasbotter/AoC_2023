package day09

import (
	"advent_2023/utils"
	"fmt"
	"strconv"
	"strings"
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
	sum := 0
	for _, line := range *input {
		splitted := strings.Split(line, " ")
		numberSlice := make([]int, len(splitted))
		for index, number := range splitted {
			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			numberSlice[index] = value
		}
		sum += findMissingValue(numberSlice)

	}
	fmt.Println(sum)
}

func findMissingValue(numberSlice []int) int {
	newSlice := make([]int, len(numberSlice)-1)
	isLast := true
	for index := range newSlice {
		newSlice[index] = numberSlice[index+1] - numberSlice[index]
		if newSlice[index] != 0 {
			isLast = false
		}
	}
	if isLast {
		return numberSlice[len(numberSlice)-1]
	}
	return numberSlice[len(numberSlice)-1] + findMissingValue(newSlice)
}

func part_b(input *[]string) {
	sum := 0
	for _, line := range *input {
		splitted := strings.Split(line, " ")
		numberSlice := make([]int, len(splitted))
		for index, number := range splitted {
			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			numberSlice[index] = value
		}
		sum += findMissingValueReverse(numberSlice)
	}
	fmt.Println(sum)
}

func findMissingValueReverse(numberSlice []int) int {
	newSlice := make([]int, len(numberSlice)-1)
	isLast := true
	for index := range newSlice {
		newSlice[index] = numberSlice[index+1] - numberSlice[index]
		if newSlice[index] != 0 {
			isLast = false
		}
	}
	if isLast {
		return numberSlice[0]
	}
	return numberSlice[0] - findMissingValueReverse(newSlice)
}
