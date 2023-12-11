package day11

import (
	"advent_2023/utils"
	"fmt"
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

type galaxy struct {
	x int
	y int
}

func makeKey(x int, y int) string {
	return fmt.Sprint(x, ",", y)
}

func part_a(input *[]string) {
	additional := 2
	emptyLines := make(map[int]bool)
	for yIndex, line := range *input {
		shouldAdd := true
	innerLines:
		for _, char := range line {
			if char == '#' {
				shouldAdd = false
				break innerLines
			}
		}
		if shouldAdd {
			emptyLines[yIndex] = true
		}
	}
	// Check Emtpy Columns
	emptyColumns := make(map[int]bool)
	for xIndex := range (*input)[0] {
		shouldAdd := true
	innerColumns:
		for yIndex := range *input {
			if (*input)[yIndex][xIndex] == '#' {
				shouldAdd = false
				break innerColumns
			}
		}
		if shouldAdd {
			emptyColumns[xIndex] = true
		}
	}
	galaxies := make(map[string]galaxy)
	for yIndex, line := range *input {
		for xIndex, char := range line {
			if char == '#' {
				key := makeKey(xIndex, yIndex)
				galaxies[key] = galaxy{x: xIndex, y: yIndex}
			}
		}
	}
	var distances uint64 = 0
	// Can make this log time, but slightly more effort
	// Improve if performance is low
	for keyLeft, valueLeft := range galaxies {
		for keyRight, valueRight := range galaxies {
			if keyLeft != keyRight {
				// Start with X
				distance := 0
				direction := 1
				if valueLeft.x > valueRight.x {
					direction = -1
				}
				for x := valueLeft.x; x != valueRight.x; x = x + direction {
					_, exists := emptyColumns[x]
					if exists {
						distance += (additional * direction)
					} else {
						distance += direction
					}
				}
				distances += (uint64)(AbsInt(distance))
				// Start with Y
				distance = 0
				direction = 1
				if valueLeft.y > valueRight.y {
					direction = -1
				}
				for y := valueLeft.y; y != valueRight.y; y = y + direction {
					_, exists := emptyLines[y]
					if exists {
						distance += (additional * direction)
					} else {
						distance += direction
					}
				}
				distances += (uint64)(AbsInt(distance))
			}
		}
	}
	fmt.Println(distances / 2)
}

func AbsInt(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

// Cool, so rules have changed drastically
func part_b(input *[]string) {
	additional := 1000000
	emptyLines := make(map[int]bool)
	for yIndex, line := range *input {
		shouldAdd := true
	innerLines:
		for _, char := range line {
			if char == '#' {
				shouldAdd = false
				break innerLines
			}
		}
		if shouldAdd {
			emptyLines[yIndex] = true
		}
	}
	// Check Emtpy Columns
	emptyColumns := make(map[int]bool)
	for xIndex := range (*input)[0] {
		shouldAdd := true
	innerColumns:
		for yIndex := range *input {
			if (*input)[yIndex][xIndex] == '#' {
				shouldAdd = false
				break innerColumns
			}
		}
		if shouldAdd {
			emptyColumns[xIndex] = true
		}
	}
	galaxies := make(map[string]galaxy)
	for yIndex, line := range *input {
		for xIndex, char := range line {
			if char == '#' {
				key := makeKey(xIndex, yIndex)
				galaxies[key] = galaxy{x: xIndex, y: yIndex}
			}
		}
	}
	var distances uint64 = 0
	// Can make this log time, but slightly more effort
	// Improve if performance is low
	for keyLeft, valueLeft := range galaxies {
		for keyRight, valueRight := range galaxies {
			if keyLeft != keyRight {
				// Start with X
				distance := 0
				direction := 1
				if valueLeft.x > valueRight.x {
					direction = -1
				}
				for x := valueLeft.x; x != valueRight.x; x = x + direction {
					_, exists := emptyColumns[x]
					if exists {
						distance += (additional * direction)
					} else {
						distance += direction
					}
				}
				distances += (uint64)(AbsInt(distance))
				// Start with Y
				distance = 0
				direction = 1
				if valueLeft.y > valueRight.y {
					direction = -1
				}
				for y := valueLeft.y; y != valueRight.y; y = y + direction {
					_, exists := emptyLines[y]
					if exists {
						distance += (additional * direction)
					} else {
						distance += direction
					}
				}
				distances += (uint64)(AbsInt(distance))
			}
		}
	}
	fmt.Println(distances / 2)
}
