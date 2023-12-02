package day03

import (
	"advent_2023/utils"
	"fmt"
	"strconv"
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
				foundSymbol := false

				startY := yIndex - 1
				if startY < 0 {
					startY = 0
				}
				endY := yIndex + 1
				if endY >= len(lines)-1 {
					endY = len(lines) - 1
				}
				startX := numberStart - 1
				if startX < 0 {
					startX = 0
				}
				endX := xIndex + 1
				if endX >= len(line)-1 {
					endX = len(line) - 1
				}
				//out:
				for y := startY; y < endY+1; y++ {
					for x := startX; x < endX; x++ {
						fmt.Print(string(lines[y][x]))
						if !unicode.IsDigit(rune(lines[y][x])) && lines[y][x] != '.' {

							/*
								if value == 547 {
									fmt.Println(string(lines[y][x]), value)
									for yPrint := startY; yPrint < endY+1; yPrint++ {
										fmt.Println(lines[yPrint][startX : endX+1])
									}
									fmt.Println()
								}
							*/
							foundSymbol = true
							//break out
						}

					}
					fmt.Println()
				}
				fmt.Println()
				if foundSymbol {
					sum += value

				}
			}
		}
	}
	fmt.Println(sum)
}

// 528616 < x < 531806

func part_b(input *[]string) {
}
