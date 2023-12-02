package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var DigitParser = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

var DigitRegex = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		digit1, digit2 := findMatches(line)
		sum += parseCoord(digit1, digit2)
	}

	fmt.Printf("Sum: %d\n", sum)
}

func findMatches(line string) (string, string) {
	first := ""
	last := ""

	tokens := ""
	for i := 0; i < len(line); i++ {
		tokens = tokens + line[i:i+1]

		first = DigitRegex.FindString(tokens)

		if first != "" {
			break
		}
	}

	tokens = ""
	for i := len(line) - 1; i >= 0; i-- {
		tokens = line[i:i+1] + tokens

		last = DigitRegex.FindString(tokens)

		if last != "" {
			break
		}
	}

	return first, last
}

func parseDigit(digit string) string {
	parsed, ok := DigitParser[digit]

	if !ok {
		panic(fmt.Sprintf("Invalid digit: %s", digit))
	}

	return parsed
}

func parseCoord(digit1 string, digit2 string) int {
	parsed, err := strconv.Atoi(parseDigit(digit1) + parseDigit(digit2))

	if err != nil {
		panic(err)
	}

	return parsed
}
