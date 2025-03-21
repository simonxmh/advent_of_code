package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// get is a helper function to handle errors by panicking
func get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

type pair struct {
	x, y int
}

// processLine parses a line of input and returns the target number and slice of operands

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}

	f := get(os.Open(os.Args[1]))
	defer f.Close()

	scanner := bufio.NewScanner(f)
	graph := make([][]rune, 0)
	// sum := 0

	for scanner.Scan() {
		// construct graph from scanned text
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, c := range line {
			row[i] = c
		}
		graph = append(graph, row)
	}

	//validate graph
	m := len(graph)
	n := len(graph[0])
	occupied := make(map[pair]bool)

	charMap := make(map[rune][]pair)
	for i, row := range graph {
		for j, elem := range row {
			if elem == '.' || elem == '#' {
				continue
			}
			charMap[elem] = append(charMap[elem], pair{i, j})
			occupied[pair{i, j}] = true
		}
	}

	// p1
	antinodes := make(map[pair]bool)

	// p2
	resonant := make(map[pair]bool)

	for _, pos := range charMap {
		// for each of the positions, go to every other position and check their displacement
		for i, p := range pos {
			for j, q := range pos {
				if i == j {
					continue
				}

				// get the direction of the vector
				dx := q.x - p.x
				dy := q.y - p.y

				// idea is that the antinodes are either one units of distance behind
				// or two units ahead
				a1 := pair{p.x - dx, p.y - dy}
				a2 := pair{p.x + 2*dx, p.y + 2*dy}

				if isInBounds(a1.x, a1.y, m, n) {
					antinodes[a1] = true
				}

				if isInBounds(a2.x, a2.y, m, n) {
					antinodes[a2] = true
				}

				for backward := 1; isInBounds(p.x-backward*dx, p.y-backward*dy, m, n); backward++ {
					if graph[p.x-backward*dx][p.y-backward*dy] != '.' {
						continue
					}
					resonant[pair{p.x - backward*dx, p.y - backward*dy}] = true
					graph[p.x-backward*dx][p.y-backward*dy] = '#'
				}

				for forward := 2; isInBounds(p.x+forward*dx, p.y+forward*dy, m, n); forward++ {
					if graph[p.x+forward*dx][p.y+forward*dy] != '.' {
						continue
					}
					resonant[pair{p.x + forward*dx, p.y + forward*dy}] = true
					graph[p.x+forward*dx][p.y+forward*dy] = '#'
				}
			}
		}
	}

	println(len(resonant) + len(occupied))
	printGraph(graph)
}

func isInBounds(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func printGraph(graph [][]rune) {
	for _, row := range graph {
		for _, elem := range row {
			print(string(elem))
		}
		println()
	}
}

func printCharMap(charMap map[rune][]pair) {
	for entry, pos := range charMap {
		for _, p := range pos {
			fmt.Printf("entry: %d, pos: %d, %d\n", entry, p.x, p.y)
		}
	}
}
