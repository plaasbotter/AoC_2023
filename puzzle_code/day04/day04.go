package day04

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

func part_a(input *[]string) {
	var sum int = 0
	for _, line := range *input {
		newline := replaceSingleDigits(line)
		splitCard := strings.Split(newline, ": ")
		splitGame := strings.Split(splitCard[1], " | ")
		winningNumbers := strings.Split(splitGame[0], " ")
		winningChecks := make([]bool, len(winningNumbers))
		playedNumbers := strings.Split(splitGame[1], " ")
		for _, play := range playedNumbers {
		inner:
			for index, test := range winningNumbers {

				if play == test {
					winningChecks[index] = true
					break inner
				}
			}
		}
		var score int = 0
		for _, check := range winningChecks {
			if check {
				if score == 0 {
					score = 1
				} else {
					score = score << 1
				}
			}
		}
		fmt.Println(score, winningChecks, winningNumbers)
		sum += score
	}
	fmt.Println(sum)
}

func part_b(input *[]string) {
	amounts := make([]int, len(*input))
	var sum int = 0
	for gameIndex, line := range *input {
		newline := replaceSingleDigits(line)
		splitCard := strings.Split(newline, ": ")
		splitGame := strings.Split(splitCard[1], " | ")
		winningNumbers := strings.Split(splitGame[0], " ")
		winningChecks := make([]bool, len(winningNumbers))
		playedNumbers := strings.Split(splitGame[1], " ")
		for _, play := range playedNumbers {
		inner:
			for index, test := range winningNumbers {
				if play == test {
					winningChecks[index] = true
					break inner
				}
			}
		}
		var counter int = gameIndex + 1
		amounts[gameIndex] += 1
		for _, check := range winningChecks {
			if check {
				amounts[counter] += amounts[gameIndex]
				counter += 1
			}
		}
		sum += amounts[gameIndex]
	}
	fmt.Println(sum)
}

func replaceSingleDigits(input string) string {
	for i := 1; i < 10; i++ {
		input = strings.ReplaceAll(input, fmt.Sprint("  ", i), fmt.Sprint(" 0", i))
	}
	return input
}
