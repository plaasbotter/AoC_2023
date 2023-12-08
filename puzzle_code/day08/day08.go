package day08

import (
	"advent_2023/utils"
	"fmt"
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

type node struct {
	key   string
	left  string
	right string
}

func part_a(input *[]string) {
	nodeMap := make(map[string]node)
	instructions := (*input)[0]
	for index, line := range *input {
		if index > 1 {
			newline := strings.ReplaceAll(line, "(", "")
			newline = strings.ReplaceAll(newline, ")", "")
			splitted := strings.Split(newline, " = ")
			leftright := strings.Split(splitted[1], ", ")
			nodeMap[splitted[0]] = node{key: splitted[0], left: leftright[0], right: leftright[1]}
		}
	}
	current := nodeMap["AAA"]
	instructionCounter := 0
	counter := 0
	for {
		switch instructions[instructionCounter] {
		case 'L':
			current = nodeMap[current.left]
		case 'R':
			current = nodeMap[current.right]
		}
		counter += 1
		if current.key == "ZZZ" {
			break
		}
		instructionCounter = (instructionCounter + 1) % len(instructions)
	}
	fmt.Println(counter)
}

func part_b(input *[]string) {
	nodeMap := make(map[string]node)
	instructions := (*input)[0]
	startingNodes := make([]string, 0)
	for index, line := range *input {
		if index > 1 {
			newline := strings.ReplaceAll(line, "(", "")
			newline = strings.ReplaceAll(newline, ")", "")
			splitted := strings.Split(newline, " = ")
			if splitted[0][2] == 'A' {
				startingNodes = append(startingNodes, splitted[0])
			}
			leftright := strings.Split(splitted[1], ", ")
			nodeMap[splitted[0]] = node{key: splitted[0], left: leftright[0], right: leftright[1]}
		}
	}
	//Find least common divisor in repeats
	result := 0
	for _, start := range startingNodes {
		value := findRepeat(&instructions, nodeMap[start], &nodeMap, 0)
		if result == 0 {
			result = value
		} else {
			result = result * value / GCD(result, value)
		}
	}
	fmt.Println(result)
}

func findRepeat(instructions *string, current node, nodeMap *map[string]node, last int) int {
	instructionCounter := 0
	var counter int = 0
	for {
		switch (*instructions)[instructionCounter] {
		case 'L':
			current = (*nodeMap)[current.left]
		case 'R':
			current = (*nodeMap)[current.right]
		}
		counter += 1
		if current.key[2] == 'Z' {
			break
		}
		instructionCounter = (instructionCounter + 1) % len(*instructions)
	}
	if counter == last {
		return counter
	}
	return findRepeat(instructions, current, nodeMap, counter)
}

// greatest common divisor ia Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
