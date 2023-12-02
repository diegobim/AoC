package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum := 0
	re := regexp.MustCompile(`\d`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		digit1 := matches[0]
		digit2 := matches[len(matches)-1]
		parsed, _ := strconv.Atoi(digit1 + digit2)

		sum += parsed
	}

	fmt.Printf("Sum: %d\n", sum)
}
