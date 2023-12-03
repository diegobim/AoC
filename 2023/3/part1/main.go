package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type Direction struct {
	x int
	y int
}

var (
	left        = Direction{-1, 0}
	topLeft     = Direction{-1, -1}
	up          = Direction{0, -1}
	topRight    = Direction{1, -1}
	right       = Direction{1, 0}
	bottomRight = Direction{1, 1}
	down        = Direction{0, 1}
	bottomLeft  = Direction{-1, 1}
)

var ScanDirections = []Direction{left, topLeft, up, topRight, right, bottomRight, down, bottomLeft}

var DigitRegex = regexp.MustCompile(`\d+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows := []string{}

	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	var sum int
	for i, row := range rows {
		matches := DigitRegex.FindAllStringIndex(row, -1)
		for _, match := range matches {
			if scanSurroundings(i, match, rows) {
				num := row[match[0]:match[1]]
				sum += parseInt(num)
			}
		}
	}

	fmt.Println(sum)
}

func scanSurroundings(rowIndex int, match []int, rows []string) bool {
	for i := match[0]; i < match[1]; i++ {
		for _, direction := range ScanDirections {
			y := rowIndex + direction.y
			x := i + direction.x

			if y < 0 || y >= len(rows) {
				continue
			}

			if x < 0 || x >= len(rows[y]) {
				continue
			}

			neighbor := rows[y][x]

			if isSurroundedBySymbol(neighbor) {
				return true
			}
		}
	}

	return false
}

func isSurroundedBySymbol(c byte) bool {
	return !(c == '.' || unicode.IsDigit(rune(c)))
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
