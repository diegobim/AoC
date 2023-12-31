package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var DigitRegex = regexp.MustCompile(`\d+`)

var KindToOrder = map[string]int{
	"HC":   1,
	"1P":   2,
	"2P":   3,
	"3oaK": 4,
	"FH":   5,
	"4oaK": 6,
	"5oaK": 7,
}

var LabelToOrder = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Hand struct {
	cards string
	bid   int
	kind  string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	hands := *parse(&lines)

	sortHands(&hands)

	winnings := solve(&hands)

	fmt.Println(winnings)
}

func parse(lines *[]string) *[]Hand {
	hands := []Hand{}

	for _, line := range *lines {
		hands = append(hands, *parseHand(line))
	}

	return &hands
}

func parseHand(line string) *Hand {
	blocks := strings.Split(line, " ")

	cards := blocks[0]
	bid := parseInt(blocks[1])
	kind := findKind(&cards)

	return &Hand{cards, bid, kind}
}

func findKind(cards *string) string {
	occurrences := map[string]int{}

	for _, c := range *cards {
		occurrences[string(c)]++
	}

	if len(occurrences) == 1 {
		return "5oaK"
	}

	if len(occurrences) == 2 {
		var has3eq, hasPair bool

		for _, v := range occurrences {
			if v == 4 {
				return "4oaK"
			}
			if v == 3 {
				has3eq = true
			}
			if v == 2 {
				hasPair = true
			}
		}

		if has3eq && hasPair {
			return "FH"
		}

		return "3oaK"
	}

	if len(occurrences) == 3 {
		for _, v := range occurrences {
			if v == 3 {
				return "3oaK"
			}
			if v == 2 {
				return "2P"
			}
		}
	}

	if len(occurrences) == 4 {
		return "1P"
	}

	return "HC"
}

func sortHands(hands *[]Hand) {
	sort.Slice(*hands, func(i, j int) bool {
		iType, jType := (*hands)[i].kind, (*hands)[j].kind
		iTypeOrder, jTypeOrder := KindToOrder[iType], KindToOrder[jType]

		if iTypeOrder != jTypeOrder {
			return iTypeOrder < jTypeOrder
		}

		iCards, jCards := (*hands)[i].cards, (*hands)[j].cards

		for i := 0; i < len(iCards); i++ {
			iLabel, jLabel := string(iCards[i]), string(jCards[i])
			iLabelOrder, jLabelOrder := LabelToOrder[iLabel], LabelToOrder[jLabel]

			if iLabelOrder != jLabelOrder {
				return iLabelOrder < jLabelOrder
			}
		}

		return false
	})
}

func solve(hands *[]Hand) int {
	winnings := 0

	for rank, hand := range *hands {
		winnings += (rank + 1) * hand.bid
	}

	return winnings
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
