package main

import (
	"bufio"
	"fmt"
	"math"
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

	// waysToWin := solve(&race)
	waysToWin := solve_quadratic(&race)

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

func solve_quadratic(race *Race) int64 {
	time := float64(race.time)
	distance := float64(race.distance)

	// Δ = b*b - 4*a
	d := math.Pow(time, 2) - 4*distance

	// x = (b +- sqrt(Δ)) / 2
	min := (time - math.Sqrt(d)) / 2
	max := (time + math.Sqrt(d)) / 2

	minHoldTime := int64(math.Floor(min + 1))
	maxHoldTime := int64(math.Ceil(max - 1))

	return maxHoldTime - minHoldTime + 1
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
