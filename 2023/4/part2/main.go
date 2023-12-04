package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id      int
	winning []int
	actual  []int
	hits    int
}

var DigitRegex = regexp.MustCompile(`\d+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cards := []ScratchCard{}
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, *parse(line))
	}

	totalCards := countWithCopies(&cards, 0, len(cards))

	fmt.Println(totalCards)
}

func countWithCopies(cards *[]ScratchCard, copyStart int, copyEnd int) int {
	totalCards := 0

	for _, card := range (*cards)[copyStart:copyEnd] {
		totalCards++

		if card.hits == 0 {
			continue
		}

		totalCards += countWithCopies(cards, card.id, card.id+card.hits)
	}

	return totalCards
}

func parse(line string) *ScratchCard {
	parts := strings.Split(line, ":")
	game := DigitRegex.FindString(parts[0])
	numbers := strings.Split(parts[1], "|")

	id := parseInt(game)
	winning := *parseNumbers(numbers[0])
	actual := *parseNumbers(numbers[1])
	hits := calculateHits(&winning, &actual)

	return &ScratchCard{
		id,
		winning,
		actual,
		hits,
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

func calculateHits(winning *[]int, actual *[]int) int {
	hash := *intersect(winning, actual)
	return len(hash)
}

func intersect(a *[]int, b *[]int) *[]int {
	set := make([]int, 0)
	hash := make(map[int]struct{})

	for _, v := range *a {
		hash[v] = struct{}{}
	}

	for _, v := range *b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return &set
}
