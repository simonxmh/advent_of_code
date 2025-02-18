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

func main() {
	if len(os.Args) != 2 {
		log.Fatal("error with filepath")
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
			if graph[r][c] == 'A' {
				total += match(graph, r, c)
			}
		}
	}
	fmt.Println(total)
	return
}

func match(graph [][]rune, r, c int) int {
	if r < 1 || c < 1 || r >= len(graph)-1 || c >= len(graph[r])-1 {
		return 0
	}
	tl, tr, bl, br := graph[r-1][c-1], graph[r-1][c+1], graph[r+1][c-1], graph[r+1][c+1]
	tlbr, trbl := false, false
	if tl == 'M' && br == 'S' || tl == 'S' && br == 'M' {
		tlbr = true
	}
	if tr == 'M' && bl == 'S' || tr == 'S' && bl == 'M' {
		trbl = true
	}
	if tlbr && trbl {
		return 1
	}
	return 0
}
