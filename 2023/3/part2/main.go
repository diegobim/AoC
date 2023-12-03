package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"

	"golang.org/x/exp/maps"
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
		for j, char := range row {
			if isGearCandidate(char) {
				partNumbers := scanSurroundings(i, j, rows)
				if len(partNumbers) == 2 {
					sum += parseInt(partNumbers[0]) * parseInt(partNumbers[1])
				}
			}
		}
	}

	fmt.Println(sum)
}

func scanSurroundings(rowIndex int, charIndex int, rows []string) []string {
	rowMatches := map[int][][]int{}
	adjacentMatches := map[string]string{}

	for _, direction := range ScanDirections {
		y := rowIndex + direction.y
		x := charIndex + direction.x

		if y < 0 || y >= len(rows) {
			continue
		}

		if x < 0 || x >= len(rows[y]) {
			continue
		}

		row := rows[y]
		neighbor := row[x]

		if isDigit(neighbor) {
			if _, ok := rowMatches[y]; !ok {
				rowMatches[y] = DigitRegex.FindAllStringIndex(row, -1)
			}

			for _, match := range rowMatches[y] {
				key := fmt.Sprintf("%v:%v:%v", y, match[0], match[1])

				if _, ok := adjacentMatches[key]; !ok {
					if match[0] <= x && match[1] > x {
						adjacentMatches[key] = row[match[0]:match[1]]
					}
				}
			}
		}
	}

	return maps.Values(adjacentMatches)
}

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func isGearCandidate(c rune) bool {
	return c == '*'
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
