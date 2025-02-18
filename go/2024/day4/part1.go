package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

var NEWPATH = []rune{'M', 'A', 'S'}

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Up, Down, Left, Right
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonals: Top-left, Top-right, Bottom-left, Bottom-right
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: d3 [filepath]")
	}

	f := get(os.Open(os.Args[1]))
	r := bufio.NewScanner(f)

	var graph [][]rune
	for r.Scan() {
		row := []rune(r.Text())
		graph = append(graph, row)
	}

	total := 0
	for r := range graph {
		for c := range graph[r] {
			if graph[r][c] == 'M' {
				visited := make([][]bool, len(graph))
				for i := range visited {
					visited[i] = make([]bool, len(graph[i]))
				}
				total += walk(graph, visited, r, c, 0, 0, 0)
			}
		}
	}
	fmt.Println(total)
	return
}

func walk(graph [][]rune, visited [][]bool, r, c, dr, dc, ind int) int {
	if r < 0 || c < 0 || r >= len(graph) || c >= len(graph[r]) {
		return 0
	}
	if visited[r][c] || graph[r][c] != PATH[ind] {
		return 0
	}
	if ind == 3 {
		return 1
	}
	total := 0
	for _, d := range directions {
		newDr, newDc := d[0], d[1]
		if dr == 0 && dc == 0 || (dr == newDr && dc == newDc) {
			total += walk(graph, visited, r+newDr, c+newDc, newDr, newDc, ind+1)
		}
	}
	return total
}
