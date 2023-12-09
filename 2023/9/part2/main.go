package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	oasis := *parse(&lines)

	var sum int64
	for _, record := range oasis {
		extrapolate(&record)
		sum += record[0]
	}

	fmt.Println(sum)
}

func parse(lines *[]string) *[][]int64 {
	oasis := [][]int64{}

	for _, line := range *lines {
		oasis = append(oasis, []int64{})
		for _, char := range strings.Split(line, " ") {
			oasis[len(oasis)-1] = append(oasis[len(oasis)-1], parseInt(char))
		}
	}

	return &oasis
}

func extrapolate(record *[]int64) {
	var allZero bool = true

	diffs := []int64{}
	for i := 1; i < len(*record); i++ {
		curr := (*record)[i]
		prev := (*record)[i-1]
		diff := curr - prev
		diffs = append(diffs, diff)

		allZero = allZero && diff == 0
	}

	if !allZero {
		extrapolate(&diffs)
	}

	*record = append([]int64{(*record)[0] - diffs[0]}, (*record)...)
}

func parseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
