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

type ScratchCard struct {
	id      int
	winning []int
	actual  []int
	points  int
}

var DigitRegex = regexp.MustCompile(`\d+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cards := []ScratchCard{}
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, *parse(line))
	}

	totalPoints := 0
	for _, card := range cards {
		totalPoints += card.points
	}

	fmt.Println(totalPoints)
}

func parse(line string) *ScratchCard {
	parts := strings.Split(line, ":")
	game := DigitRegex.FindString(parts[0])
	numbers := strings.Split(parts[1], "|")

	id := parseInt(game)
	winning := *parseNumbers(numbers[0])
	actual := *parseNumbers(numbers[1])
	points := calculatePoints(&winning, &actual)

	return &ScratchCard{
		id,
		winning,
		actual,
		points,
	}
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

func calculatePoints(winning *[]int, actual *[]int) int {
	hash := intersect(*winning, *actual)
	return int(math.Pow(2, float64(len(hash)-1)))
}

func intersect(a []int, b []int) []int {
	set := make([]int, 0)
	hash := make(map[int]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
