package day07

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

type play struct {
	simple string
	hand   []int
	bet    int
}

func emptyPlay() *[]play {
	tempPlay := make([]play, 0)
	return &tempPlay
}

// Incredibly ugly but it works
func part_a(input *[]string) {
	cardMap := map[string]int{
		"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "T": 8, "J": 9, "Q": 10, "K": 11, "A": 12,
	}
	handMap := map[string]*[]play{
		"1": emptyPlay(), "2": emptyPlay(), "22": emptyPlay(), "3": emptyPlay(), "32": emptyPlay(), "4": emptyPlay(), "5": emptyPlay(),
	}

	for _, line := range *input {
		splitted := strings.Split(line, " ")
		hand, game := getHand(splitted[0], &cardMap)
		bet, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}
		tempPlay := play{hand: hand, bet: bet, simple: splitted[0]}
		link := ""
		if game[4] == 1 {
			link = "5"
		} else if game[3] == 1 {
			link = "4"
		} else if game[2] == 1 && game[1] == 1 {
			link = "32"
		} else if game[2] == 1 {
			link = "3"
		} else if game[1] == 2 {
			link = "22"
		} else if game[1] == 1 {
			link = "2"
		} else {
			link = "1"
		}
		handMap[link] = AddPlay(tempPlay, handMap[link])
	}
	// This part is stupid, but only because maps are not in order when creating them. Maps are weird
	counter := 1
	sum := 0

	counter, sum = calculateWin(sum, counter, handMap, "1")
	counter, sum = calculateWin(sum, counter, handMap, "2")
	counter, sum = calculateWin(sum, counter, handMap, "22")
	counter, sum = calculateWin(sum, counter, handMap, "3")
	counter, sum = calculateWin(sum, counter, handMap, "32")
	counter, sum = calculateWin(sum, counter, handMap, "4")
	counter, sum = calculateWin(sum, counter, handMap, "5")
	fmt.Println(handMap, sum, counter)
}

func calculateWin(sum int, counter int, handMap map[string]*[]play, handType string) (int, int) {
	deref := *handMap[handType]
	for _, result := range deref {
		sum += counter * result.bet
		counter++
	}
	return counter, sum
}

func AddPlay(input play, results *[]play) *[]play {
	for playIndex, other := range *results {
		if isPlayLess(other, input) {
			tempResult := *results
			tempResult = append(tempResult[:playIndex+1], tempResult[playIndex:]...)
			tempResult[playIndex] = input
			results = &tempResult
			return results
		}
	}
	*results = append(*results, input)
	return results
}

func isPlayLess(left play, right play) bool {
	for index := range right.hand {
		if left.hand[index] < right.hand[index] {
			return false
		} else if left.hand[index] > right.hand[index] {
			return true
		}
	}
	return true
}

func getHand(input string, cardMap *map[string]int) ([]int, []int) {
	cardMapDereference := *cardMap
	hand, game := make([]int, len(input)), make([]int, len(input))
	handMap := make([]int, len(cardMapDereference))
	for index := range input {
		char := string(input[index])
		cardValue := cardMapDereference[char]
		hand[index] = cardValue
		handMap[cardValue] += 1
	}
	for _, amount := range handMap {
		if amount > 0 {
			game[amount-1] += 1
		}
	}
	return hand, game
}

func part_b(input *[]string) {
	cardMap := map[string]int{
		"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 0, "Q": 10, "K": 11, "A": 12,
	}
	handMap := map[string]*[]play{
		"1": emptyPlay(), "2": emptyPlay(), "22": emptyPlay(), "3": emptyPlay(), "32": emptyPlay(), "4": emptyPlay(), "5": emptyPlay(),
	}

	for _, line := range *input {
		splitted := strings.Split(line, " ")
		hand, game := getHandWithJoker(splitted[0], &cardMap)
		bet, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}
		tempPlay := play{hand: hand, bet: bet, simple: splitted[0]}
		link := ""
		if game[4] == 1 {
			link = "5"
		} else if game[3] == 1 {
			link = "4"
		} else if game[2] == 1 && game[1] == 1 {
			link = "32"
		} else if game[2] == 1 {
			link = "3"
		} else if game[1] == 2 {
			link = "22"
		} else if game[1] == 1 {
			link = "2"
		} else {
			link = "1"
		}
		handMap[link] = AddPlay(tempPlay, handMap[link])
	}
	counter := 1
	sum := 0
	counter, sum = calculateWin(sum, counter, handMap, "1")
	counter, sum = calculateWin(sum, counter, handMap, "2")
	counter, sum = calculateWin(sum, counter, handMap, "22")
	counter, sum = calculateWin(sum, counter, handMap, "3")
	counter, sum = calculateWin(sum, counter, handMap, "32")
	counter, sum = calculateWin(sum, counter, handMap, "4")
	counter, sum = calculateWin(sum, counter, handMap, "5")
	fmt.Println(handMap, sum, counter)
}

func getHandWithJoker(input string, cardMap *map[string]int) ([]int, []int) {
	cardMapDereference := *cardMap
	hand, game := make([]int, len(input)), make([]int, len(input))
	handMap := make([]int, len(cardMapDereference))
	for index := range input {
		char := string(input[index])
		cardValue := cardMapDereference[char]
		hand[index] = cardValue
		handMap[cardValue] += 1
	}
	if handMap[0] > 0 && handMap[0] != len(input) {
		maxIndex := 0
		maxValue := 0
		for i := 1; i < len(handMap); i++ {
			if handMap[i] > maxValue {
				maxValue = handMap[i]
				maxIndex = i
			}
		}
		handMap[maxIndex] += handMap[0]
		handMap[0] = 0
	}

	for _, amount := range handMap {
		if amount > 0 {
			game[amount-1] += 1
		}
	}
	return hand, game
}
