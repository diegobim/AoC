package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

var DigitRegex = regexp.MustCompile(`\d+`)

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id    int
	sets  []Set
	valid bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseGame(line))
	}

	setPowerSum := 0
	for _, game := range games {
		var maxRed, maxGreen, maxBlue int

		for _, set := range game.sets {
			if set.red > maxRed {
				maxRed = set.red
			}

			if set.green > maxGreen {
				maxGreen = set.green
			}

			if set.blue > maxBlue {
				maxBlue = set.blue
			}
		}

		setPowerSum += maxRed * maxGreen * maxBlue
	}

	fmt.Println(setPowerSum)
}

func parseGame(line string) Game {
	parts := strings.Split(line, ":")
	game := parseInt(DigitRegex.FindString(parts[0]))
	sets := parseSets(strings.Split(parts[1], ";"))

	return Game{
		id:    game,
		sets:  sets,
		valid: validateSets(sets),
	}
}

func parseSets(sets []string) []Set {
	result := []Set{}

	for _, set := range sets {

		var red, green, blue int
		for _, color := range strings.Split(set, ",") {
			if strings.Contains(color, "red") {
				red = parseInt(DigitRegex.FindString(color))
			} else if strings.Contains(color, "green") {
				green = parseInt(DigitRegex.FindString(color))
			} else if strings.Contains(color, "blue") {
				blue = parseInt(DigitRegex.FindString(color))
			}
		}

		result = append(result, Set{red, green, blue})
	}

	return result
}

func validateSets(sets []Set) bool {
	for _, set := range sets {
		if set.blue > MAX_BLUE || set.green > MAX_GREEN || set.red > MAX_RED {
			return false
		}
	}

	return true
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
