package day06

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
	lines := *input
	times := loadValues(lines[0])
	distances := loadValues(lines[1])
	records, value, aboveCount := 1, 0, 0
	for index := range times {
		aboveCount = 0
		for i := 0; i <= times[index]; i++ {
			value = (times[index] * i) - (i * i)
			if value > distances[index] {
				aboveCount += 1
			}
		}
		records *= aboveCount
	}
	fmt.Println(records)
}

func loadValues(line string) []int {
	returnValue := make([]int, 0)
	splitted := strings.Split(line, " ")
	for _, split := range splitted {
		if len(split) > 0 {
			num, err := strconv.Atoi(split)
			if err == nil {
				returnValue = append(returnValue, num)
			}
		}
	}
	return returnValue
}

func part_b(input *[]string) {
	lines := *input
	time := getSingleValue(lines[0])
	distance := getSingleValue(lines[1])
	var value, i, aboveCount int64 = 0, 0, 0
	for i = 0; i <= time; i++ {
		value = (time * i) - (i * i)
		if value > distance {
			aboveCount += 1
		}
	}
	fmt.Println(aboveCount)
}

func getSingleValue(line string) int64 {
	splitted := strings.Split(line, ":")
	trimmed := strings.ReplaceAll(splitted[1], " ", "")
	num, _ := strconv.ParseInt(trimmed, 10, 64)
	return num
}
