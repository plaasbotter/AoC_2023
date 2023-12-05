package day05

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

type mapping struct {
	destination int
	source      int
	ranger      int
}

func part_a(input *[]string) {
	seeds := make([]int, 0)
	seed_to_soil := make([]mapping, 0)
	soil_to_fertilizer := make([]mapping, 0)
	fertilizer_to_water := make([]mapping, 0)
	water_to_light := make([]mapping, 0)
	light_to_temperature := make([]mapping, 0)
	temperature_to_humidity := make([]mapping, 0)
	humidity_to_location := make([]mapping, 0)
	busyReading := false
	var pointer *[]mapping
	for index, line := range *input {
		if index == 0 {
			loadSeeds(line, &seeds)
		} else if len(line) == 0 {
			busyReading = false
		} else if !busyReading && len(line) > 0 {
			switch line {
			case "seed-to-soil map:":
				pointer = &seed_to_soil
			case "soil-to-fertilizer map:":
				pointer = &soil_to_fertilizer
			case "fertilizer-to-water map:":
				pointer = &fertilizer_to_water
			case "water-to-light map:":
				pointer = &water_to_light
			case "light-to-temperature map:":
				pointer = &light_to_temperature
			case "temperature-to-humidity map:":
				pointer = &temperature_to_humidity
			case "humidity-to-location map:":
				pointer = &humidity_to_location
			}
			busyReading = true
		} else {
			*pointer = append(*pointer, generateMap(line))
		}
	}
	var value int = 0
	var lowest int = -1
	for _, seed := range seeds {
		value = seed
		value = findIndex(&seed_to_soil, value)
		value = findIndex(&soil_to_fertilizer, value)
		value = findIndex(&fertilizer_to_water, value)
		value = findIndex(&water_to_light, value)
		value = findIndex(&light_to_temperature, value)
		value = findIndex(&temperature_to_humidity, value)
		value = findIndex(&humidity_to_location, value)
		if lowest == -1 || value < lowest {
			lowest = value
		}
	}
	fmt.Println(lowest)
}

// Finds the location within the ranges
func findIndex(almanac *[]mapping, input int) int {
	isFound := false
	value := 0
	for _, section := range *almanac {
		if section.source <= input && input <= section.source+section.ranger {
			isFound = true
			value = section.destination + (input - section.source)
			break
		}
	}
	if !isFound {
		value = input
	}
	return value
}

// Generates the mapping object
// Could of course be part of the object itself
func generateMap(input string) mapping {
	var returnValue mapping = mapping{}
	splitted := strings.Split(input, " ")
	returnValue.destination, _ = strconv.Atoi(splitted[0])
	returnValue.source, _ = strconv.Atoi(splitted[1])
	returnValue.ranger, _ = strconv.Atoi(splitted[2])
	return returnValue
}

// Eazy enough to understand
func loadSeeds(input string, seeds *[]int) {
	splitted := strings.Split(input, ": ")
	values := strings.Split(splitted[1], " ")
	for _, value := range values {
		parsed, _ := strconv.Atoi(value)
		*seeds = append(*seeds, parsed)
	}
}

func part_b(input *[]string) {
	seeds := make([]int, 0)
	seed_to_soil := make([]mapping, 0)
	soil_to_fertilizer := make([]mapping, 0)
	fertilizer_to_water := make([]mapping, 0)
	water_to_light := make([]mapping, 0)
	light_to_temperature := make([]mapping, 0)
	temperature_to_humidity := make([]mapping, 0)
	humidity_to_location := make([]mapping, 0)
	busyReading := false
	var pointer *[]mapping
	for index, line := range *input {
		if index == 0 {
			loadSeeds(line, &seeds)
		} else if len(line) == 0 {
			busyReading = false
		} else if !busyReading && len(line) > 0 {
			switch line {
			case "seed-to-soil map:":
				pointer = &seed_to_soil
			case "soil-to-fertilizer map:":
				pointer = &soil_to_fertilizer
			case "fertilizer-to-water map:":
				pointer = &fertilizer_to_water
			case "water-to-light map:":
				pointer = &water_to_light
			case "light-to-temperature map:":
				pointer = &light_to_temperature
			case "temperature-to-humidity map:":
				pointer = &temperature_to_humidity
			case "humidity-to-location map:":
				pointer = &humidity_to_location
			}
			busyReading = true
		} else {
			*pointer = append(*pointer, generateMap(line))
		}
	}
	var value int = 0
	var lowest int = -1
	// The correct way to do this would be to break the seed ranges down to trees
	// Or at least I think so
	// But i'm lazy and the computer is fast enough.
	// So calculation might take like 10 minutes, but not the worst.
	// Won't be able to get away with this later
	for i := 0; i < len(seeds); i = i + 2 {
		fmt.Println(seeds[i])
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			value = j
			value = findIndex(&seed_to_soil, value)
			value = findIndex(&soil_to_fertilizer, value)
			value = findIndex(&fertilizer_to_water, value)
			value = findIndex(&water_to_light, value)
			value = findIndex(&light_to_temperature, value)
			value = findIndex(&temperature_to_humidity, value)
			value = findIndex(&humidity_to_location, value)
			if lowest == -1 || value < lowest {
				lowest = value
			}
		}
	}
	fmt.Println(lowest)
}
