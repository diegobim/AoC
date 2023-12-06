package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var DigitRegex = regexp.MustCompile(`\d+`)

type Race struct {
	time           int
	recordDistance int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := ""
	for scanner.Scan() {
		line := scanner.Text()
		lines = lines + line + "\n"
	}

	races := *parse(&lines)

	waysToWin := solve(&races)

	fmt.Println(waysToWin)
}

func parse(lines *string) *[]Race {
	blocks := strings.Split(*lines, "\n")
	times := *parseNumbers(blocks[0])
	distances := *parseNumbers(blocks[1])

	races := []Race{}

	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], distances[i]})
	}

	return &races
}

func solve(races *[]Race) int {
	total := 1

	for _, race := range *races {
		total *= calculateWaysToWin(&race)
	}

	return total
}

func calculateWaysToWin(race *Race) int {
	wins := 0
	for i := 0; i < race.time; i++ {
		speed := i
		remainingTime := race.time - i
		distanceTraveled := remainingTime * speed

		if distanceTraveled > race.recordDistance {
			wins++
		}
	}

	return wins
}

func parseNumbers(s string) *[]int {
	numbers := []int{}
	for _, n := range DigitRegex.FindAllString(s, -1) {
		numbers = append(numbers, parseInt(n))
	}
	return &numbers
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
