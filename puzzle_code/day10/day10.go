package day10

import (
	"advent_2023/utils"
	"fmt"
)

const (
	up    = "|F7"
	down  = "|JL"
	left  = "-LF"
	right = "-J7"
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
	x     int
	y     int
	steps int
	pipe  rune
}

func part_a(input *[]string) {
	max := 0
	nodeMap := LoadNodeMap(input)
	for _, value := range nodeMap {
		if value.steps > max {
			max = value.steps
		}
	}
	fmt.Println(max)
}

func LoadNodeMap(grid *[]string) map[string]*node {
	// Find start
	nodeMap := make(map[string]*node)
	nodes := make([]node, 0)
	for yIndex, line := range *grid {
		for xIndex, char := range line {
			if char == 'S' {
				nodes = append(nodes, node{key: fmt.Sprint(xIndex, ",", yIndex), x: xIndex, y: yIndex, steps: 0, pipe: char})
				break
			}
		}
	}

	// Traverse nodes
	for len(nodes) > 0 {
		tempNode := nodes[0]
		nodes = nodes[1:]
		nodeMap[tempNode.key] = &tempNode
		switch tempNode.pipe {
		case 'S':
			checkUp(tempNode, grid, &nodeMap, &nodes)
			checkDown(tempNode, grid, &nodeMap, &nodes)
			checkLeft(tempNode, grid, &nodeMap, &nodes)
			checkRight(tempNode, grid, &nodeMap, &nodes)
		case '|':
			checkUp(tempNode, grid, &nodeMap, &nodes)
			checkDown(tempNode, grid, &nodeMap, &nodes)
		case '-':
			checkLeft(tempNode, grid, &nodeMap, &nodes)
			checkRight(tempNode, grid, &nodeMap, &nodes)
		case 'F':
			checkDown(tempNode, grid, &nodeMap, &nodes)
			checkRight(tempNode, grid, &nodeMap, &nodes)
		case 'L':
			checkUp(tempNode, grid, &nodeMap, &nodes)
			checkRight(tempNode, grid, &nodeMap, &nodes)
		case 'J':
			checkLeft(tempNode, grid, &nodeMap, &nodes)
			checkUp(tempNode, grid, &nodeMap, &nodes)
		case '7':
			checkDown(tempNode, grid, &nodeMap, &nodes)
			checkLeft(tempNode, grid, &nodeMap, &nodes)
		}
	}
	return nodeMap
}

func alreadyInMap(inputNode node, xOffset int, yOffset int, nodeMap *map[string]*node) bool {
	newX := inputNode.x + xOffset
	newY := inputNode.y + yOffset
	existingNode, exists := (*nodeMap)[fmt.Sprint(newX, ",", newY)]
	if exists && existingNode.steps > inputNode.steps {
		existingNode.steps = inputNode.steps
	}
	return exists
}

func checkUp(inputNode node, grid *[]string, nodeMap *map[string]*node, nodes *[]node) {
	xOffset := 0
	yOffset := -1
	moveSet := up
	if inputNode.y > 0 {
		if canMove(inputNode, xOffset, yOffset, grid, moveSet) && !alreadyInMap(inputNode, xOffset, yOffset, nodeMap) {
			*nodes = append(*nodes, generateNode(inputNode, xOffset, yOffset, grid))
		}
	}
}

func checkDown(inputNode node, grid *[]string, nodeMap *map[string]*node, nodes *[]node) {
	xOffset := 0
	yOffset := +1
	moveSet := down
	if inputNode.y < len(*grid) {
		if canMove(inputNode, xOffset, yOffset, grid, moveSet) && !alreadyInMap(inputNode, xOffset, yOffset, nodeMap) {
			*nodes = append(*nodes, generateNode(inputNode, xOffset, yOffset, grid))
		}
	}
}

func checkLeft(inputNode node, grid *[]string, nodeMap *map[string]*node, nodes *[]node) {
	xOffset := -1
	yOffset := 0
	moveSet := left
	if inputNode.x > 0 {
		if canMove(inputNode, xOffset, yOffset, grid, moveSet) && !alreadyInMap(inputNode, xOffset, yOffset, nodeMap) {
			*nodes = append(*nodes, generateNode(inputNode, xOffset, yOffset, grid))
		}
	}
}

func checkRight(inputNode node, grid *[]string, nodeMap *map[string]*node, nodes *[]node) {
	xOffset := +1
	yOffset := 0
	moveSet := right
	if inputNode.x < len((*grid)[inputNode.y]) {
		if canMove(inputNode, xOffset, yOffset, grid, moveSet) && !alreadyInMap(inputNode, xOffset, yOffset, nodeMap) {
			*nodes = append(*nodes, generateNode(inputNode, xOffset, yOffset, grid))
		}
	}
}

func canMove(inputNode node, xOffset int, yOffset int, grid *[]string, moveSet string) bool {
	newPipe := rune((*grid)[inputNode.y+yOffset][inputNode.x+xOffset])
	for _, pipe := range moveSet {
		if newPipe == pipe {
			return true
		}
	}
	return false
}

func generateNode(inputNode node, xOffset int, yOffset int, grid *[]string) node {
	newX := inputNode.x + xOffset
	newY := inputNode.y + yOffset
	char := (*grid)[newY][newX]
	return node{key: fmt.Sprint(newX, ",", newY), x: newX, y: newY, steps: inputNode.steps + 1, pipe: rune(char)}
}

func findAndReplaceStart(grid *[]string, x int, y int, ylen int, xlen int) rune {
	top := false
	if y > 0 {
		char := (*grid)[y-1][x]
		top = char == '|' || char == '7' || char == 'F'
	}

	bottom := false
	if y < ylen-1 {
		char := (*grid)[y+1][x]
		bottom = char == '|' || char == 'J' || char == 'L'
	}

	left := false
	if x > 0 {
		char := (*grid)[y][x-1]
		left = char == '-' || char == 'L' || char == 'F'
	}

	right := false
	if x < xlen-1 {
		char := (*grid)[y][x+1]
		right = char == '-' || char == 'J' || char == '7'
	}

	if top && bottom {
		return rune('|')
	}
	if top && left {
		return rune('J')
	}
	if top && right {
		return rune('L')
	}
	if left && right {
		return rune('-')
	}
	if bottom && left {
		return rune('7')
	}
	return rune('F')
}

func part_b(input *[]string) {
	// Make Grid
	ylen := len(*input) * 3
	xlen := len((*input)[0]) * 3
	grid := make([][]byte, ylen)
	for index := range grid {
		grid[index] = make([]byte, xlen)
	}
	for yIndex, line := range *input {
		for xIndex, char := range line {
			if char != '.' {
				newX := xIndex*3 + 1
				newY := yIndex*3 + 1
				if char == 'S' {
					char = findAndReplaceStart(input, xIndex, yIndex, ylen, xlen)
				}
				switch char {
				case '|':
					grid[newY-1][newX] = '#'
					grid[newY][newX] = '#'
					grid[newY+1][newX] = '#'
				case '-':
					grid[newY][newX-1] = '#'
					grid[newY][newX] = '#'
					grid[newY][newX+1] = '#'
				case 'F':
					grid[newY+1][newX] = '#'
					grid[newY][newX] = '#'
					grid[newY][newX+1] = '#'
				case 'L':
					grid[newY-1][newX] = '#'
					grid[newY][newX] = '#'
					grid[newY][newX+1] = '#'
				case 'J':
					grid[newY-1][newX] = '#'
					grid[newY][newX] = '#'
					grid[newY][newX-1] = '#'
				case '7':
					grid[newY+1][newX] = '#'
					grid[newY][newX] = '#'
					grid[newY][newX-1] = '#'
				}
			}
		}
	}
	// Breath fill
	seen := make(map[string]bool, 0)
	floodFill(&grid, &seen, 0, 0)
	// Check for amount not seen
	sum := 0
	for yIndex, line := range *input {
		for xIndex := range line {
			inside := true
			newX := xIndex * 3
			newY := yIndex * 3
		out:
			for x := 0; x <= 2; x++ {
				for y := 0; y <= 2; y++ {
					key := fmt.Sprint(newX+x, ",", newY+y)
					_, hasSeen := seen[key]
					if hasSeen {
						inside = false
						break out
					}
				}
			}
			if inside {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}

func floodFill(grid *[][]byte, seen *map[string]bool, x int, y int) {
	if y < 0 || y >= len(*grid) || x < 0 || x >= len((*grid)[y]) {
		return
	}
	key := fmt.Sprint(x, ",", y)
	_, hasSeen := (*seen)[key]
	if hasSeen {
		return
	}
	if (*grid)[y][x] != 0 {
		return
	}
	(*seen)[key] = true
	floodFill(grid, seen, x, y-1)
	floodFill(grid, seen, x, y+1)
	floodFill(grid, seen, x-1, y)
	floodFill(grid, seen, x+1, y)
}
