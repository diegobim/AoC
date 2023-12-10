package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Direction Point

type PipeDirectionMap map[Direction]Direction

var (
	up    = Direction{-1, 0}
	down  = Direction{1, 0}
	left  = Direction{0, -1}
	right = Direction{0, 1}
)

var Pipes = []string{"|", "-", "L", "J", "7", "F"}

var PipeToDirection = map[string]PipeDirectionMap{
	"|": {
		up:   up,
		down: down,
	},
	"-": {
		right: right,
		left:  left,
	},
	"L": {
		down: right,
		left: up,
	},
	"J": {
		right: up,
		down:  left,
	},
	"7": {
		right: down,
		up:    left,
	},
	"F": {
		up:   right,
		left: down,
	},
}

func main() {
	maze, start := parse(os.Stdin)

	maxDistance := solve(maze, start)

	fmt.Println(maxDistance)
}

func parse(file *os.File) (*[][]string, Point) {
	maze := [][]string{}
	startRow := -1
	startCol := -1

	scanner := bufio.NewScanner(os.Stdin)

	index := 0
	for scanner.Scan() {
		row := scanner.Text()
		maze = append(maze, strings.Split(row, ""))

		start := strings.Index(row, "S")
		if start != -1 {
			startCol = start
			startRow = index
		}

		index++
	}

	return &maze, Point{startRow, startCol}
}

func solve(maze *[][]string, start Point) int {
	// bruteforcing all possible pipes to replace start position ¯\_(ツ)_/¯
	for _, pipe := range Pipes {
		(*maze)[start.X][start.Y] = pipe

		path := traverse(maze, start)

		if path != nil {
			return len(path) / 2
		}
	}

	panic("RIP")
}

func traverse(maze *[][]string, start Point) []Point {
	pipe := (*maze)[start.X][start.Y]
	currentPosition := start
	currentDirection := anyKey(PipeToDirection[pipe])

	path := []Point{}

	for {
		path = append(path, currentPosition)
		pipe = (*maze)[currentPosition.X][currentPosition.Y]
		direction, ok := PipeToDirection[pipe][currentDirection]

		if !ok {
			return nil
		}

		position := Point{currentPosition.X + direction.X, currentPosition.Y + direction.Y}

		// check if we are out of bounds
		if position.X < 0 || position.X >= len(*maze) || position.Y < 0 || position.Y >= len((*maze)[position.X]) {
			return nil
		}

		// loop found
		if position == start {
			// if _, ok := PipeToDirection[(*maze)[start.X][start.Y]][direction]; !ok {
			// 	return nil
			// }
			break
		}

		currentPosition = position
		currentDirection = direction
	}

	return path
}

func anyKey(m PipeDirectionMap) Direction {
	for k := range m {
		return k
	}

	panic("RIP")
}
