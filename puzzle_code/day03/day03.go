package day03

import (
	"advent_2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

type coord struct {
	x int
	y int
}

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

// Too many ifs, but universal solution
func part_a(input *[]string) {
	var sum int = 0
	lines := *input

	for yIndex, line := range lines {
		numberStart := 0
		countStarted := false
		for xIndex, char := range line {
			if unicode.IsDigit(char) && !countStarted {
				if !countStarted {
					numberStart = xIndex
					countStarted = true
				}
			} else if countStarted && (!unicode.IsDigit(char) || xIndex >= len(line)-1) {
				countStarted = false
				value, _ := strconv.Atoi(line[numberStart:xIndex])
				if xIndex == len(line)-1 && unicode.IsDigit(char) {
					value, _ = strconv.Atoi(line[numberStart : xIndex+1])
				}

				foundSymbol := false
				coords := make([]coord, 0)
				startX := numberStart - 1
				for x := startX; x < xIndex+1; x++ {
					coords = append(coords, coord{x: x, y: yIndex - 1})
					coords = append(coords, coord{x: x, y: yIndex + 1})
				}
				coords = append(coords, coord{x: startX, y: yIndex})
				if xIndex != len(line)-1 {
					coords = append(coords, coord{x: xIndex, y: yIndex})
				}
				for _, pos := range coords {
					if pos.x >= 0 && pos.x < len(line) && pos.y >= 0 && pos.y < len(lines) {
						if lines[pos.y][pos.x] != '.' {
							foundSymbol = true
							break
						}
					}
				}
				if foundSymbol {
					fmt.Println((value))
					sum += value
				}
			}
		}
	}
	fmt.Println(sum)
}

// There doesn't appear to be a star at the edge
func part_b(input *[]string) {
	var sum int = 0
	lines := *input
	countStarted := false
	var count int = 0
	var value int = 1
	for yIndex, line := range lines {
		for xIndex, char := range line {
			if char == '*' {
				count = 0
				value = 1
				for y := yIndex - 1; y < yIndex+2; y++ {
					countStarted = false
					for x := xIndex - 1; x < xIndex+2; x++ {
						if y == yIndex && x == xIndex {
							countStarted = false
							continue
						}
						if !countStarted && unicode.IsDigit(rune(lines[y][x])) {
							countStarted = true
							value *= findNumber(lines[y], x)
							count++
						} else if countStarted && lines[y][x] == '.' {
							countStarted = false
						}
					}
				}
				if count > 1 {
					fmt.Println(value)
					sum += value
				}
			}
		}
	}
	fmt.Println(sum)
}

func findNumber(input string, pos int) int {
	start := pos
	end := pos
	for i := pos; i >= 0; i-- {
		if !unicode.IsDigit(rune(input[i])) {
			break
		}
		start = i
	}
	for i := pos; i < len(input); i++ {
		if !unicode.IsDigit(rune(input[i])) {
			break
		}
		end = i
	}
	value, _ := strconv.Atoi(input[start : end+1])
	if value == 0 {
		println("mario")
	}
	return value
}
