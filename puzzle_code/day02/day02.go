package day02

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

//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

func part_a(input *[]string) {
	cubeMap := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}
	var sum int = 0
	for index, line := range *input {
		isValid := true
		splitGame := strings.Split(line, ": ")
		splitDraw := strings.Split(splitGame[1], "; ")
	out:
		for _, draw := range splitDraw {
			splitCubes := strings.Split(draw, ", ")
			for _, cube := range splitCubes {
				splitRecord := strings.Split(cube, " ")
				num, err := strconv.Atoi(splitRecord[0])
				if err != nil {
					panic("Can't convert")
				}
				if num > cubeMap[splitRecord[1]] {
					isValid = false
					break out
				}
			}
		}
		if isValid {
			sum += index + 1
		}
	}
	fmt.Println(sum)
}

func part_b(input *[]string) {
	var sum int = 0
	cubeMap := map[string]int{
		"red": 0, "green": 0, "blue": 0,
	}
	for _, line := range *input {
		cubeMap["red"] = -1
		cubeMap["green"] = -1
		cubeMap["blue"] = -1
		splitGame := strings.Split(line, ": ")
		splitDraw := strings.Split(splitGame[1], "; ")
		for _, draw := range splitDraw {
			splitCubes := strings.Split(draw, ", ")
			for _, cube := range splitCubes {
				splitRecord := strings.Split(cube, " ")
				num, err := strconv.Atoi(splitRecord[0])
				if err != nil {
					panic("Can't convert")
				}
				if cubeMap[splitRecord[1]] == -1 || num > cubeMap[splitRecord[1]] {
					cubeMap[splitRecord[1]] = num
				}
			}
		}
		sum += (cubeMap["red"] * cubeMap["green"] * cubeMap["blue"])
	}
	fmt.Println(sum)
}
