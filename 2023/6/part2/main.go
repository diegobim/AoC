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
	time     int64
	distance int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := ""
	for scanner.Scan() {
		line := scanner.Text()
		lines = lines + line + "\n"
	}

	race := *parse(&lines)
	fmt.Println(race)

	waysToWin := solve(&race)

	fmt.Println(waysToWin)
}

func parse(lines *string) *Race {
	blocks := strings.Split(*lines, "\n")
	times := DigitRegex.FindAllString(blocks[0], -1)
	distances := DigitRegex.FindAllString(blocks[1], -1)

	return &Race{
		time:     parseInt(strings.Join(times, "")),
		distance: parseInt(strings.Join(distances, "")),
	}
}

func solve(race *Race) int64 {
	return calculateWaysToWin(race)
}

func calculateWaysToWin(race *Race) int64 {
	wins := int64(0)
	for i := int64(0); i < race.time; i++ {
		speed := i
		remainingTime := race.time - i
		distanceTraveled := remainingTime * speed

		if distanceTraveled > race.distance {
			wins++
		}
	}

	return wins
}

func parseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
