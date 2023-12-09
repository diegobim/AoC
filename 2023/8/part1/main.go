package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	value string
	edges map[string]string
}

type Network struct {
	directions string
	nodes      map[string]Node
}

var NodeMapRegex = regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

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
	moves := 0
	current := "AAA"

	for current != "ZZZ" {
		for i := 0; i < len(network.directions); i++ {
			direction := string(network.directions[i])
			current = network.nodes[current].edges[direction]
			moves++

			if current == "ZZZ" {
				break
			}
		}
	}

	return moves
}
