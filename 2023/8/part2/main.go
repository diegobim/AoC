package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	value string
	edges map[string]string
}

type Network struct {
	directions string
	nodes      map[string]Node
}

var NodeMapRegex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	network := *parse(&lines)

	moves := solve(&network)

	fmt.Println(moves)
}

func parse(lines *[]string) *Network {
	directions := (*lines)[0]

	nodes := map[string]Node{}
	for _, line := range (*lines)[2:] {
		matches := NodeMapRegex.FindStringSubmatch(line)
		node, left, right := matches[1], matches[2], matches[3]

		if _, ok := nodes[node]; !ok {
			nodes[node] = Node{node, map[string]string{"L": left, "R": right}}
		}
	}

	return &Network{directions, nodes}
}

func solve(network *Network) int {
	paths := map[string]string{}
	for node := range network.nodes {
		if strings.HasSuffix(node, "A") {
			paths[node] = node
		}
	}

	movesByPath := []int{}
	for path := range paths {
		moves := 0
		current := ""

		for !strings.HasSuffix(current, "Z") {
			for i := 0; i < len(network.directions); i++ {
				direction := string(network.directions[i])
				current = paths[path]
				paths[path] = network.nodes[current].edges[direction]

				if strings.HasSuffix(current, "Z") {
					movesByPath = append(movesByPath, moves)
					break
				}

				moves++
			}
		}
	}

	return lcm(movesByPath[0], movesByPath[1], movesByPath...)
}

// Greatest Common Divisor (Euclidean algorithm)
func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// Least Common Multiple w/ GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
