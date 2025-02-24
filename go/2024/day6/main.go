package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

func get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

var directions = [][2]int{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1}, // Up -> right -> down -> left
}

func p2() {
	if len(os.Args) != 2 {
		log.Fatal("error with filepath")
	}

	f := get(os.Open(os.Args[1]))
	defer f.Close()

	r := bufio.NewScanner(f)

	var graph [][]rune
	for r.Scan() {
		row := []rune(r.Text())
		graph = append(graph, row)
	}

	var total int32
	var wg sync.WaitGroup
	sm := make(chan struct{}, 8) // 8 goroutines

	for r := range graph {
		for c := range graph[r] {
			if graph[r][c] == '.' {
				// do these in parallel
				wg.Add(1)
				sm <- struct{}{} // get the slot

				go func(r, c int) {
					defer wg.Done()
					defer func() { <-sm }() //relesae slot
					graphCopy := deepCopyGraph(graph)
					graphCopy[r][c] = '#'
					vis := make(map[[4]int]bool)
					atomic.AddInt32(&total, int32(simulate(graphCopy, vis)))
				}(r, c)
			}
		}
	}
	wg.Wait()
	fmt.Printf("%v", total)
}

func deepCopyGraph(graph [][]rune) [][]rune {
	copyGraph := make([][]rune, len(graph))
	for i := range graph {
		copyGraph[i] = append([]rune(nil), graph[i]...)
	}
	return copyGraph
}

func simulate(graph [][]rune, vis map[[4]int]bool) int {
	for r := range graph {
		for c := range graph[r] {
			if graph[r][c] == '^' {
				vis[[4]int{r, c, -1, 0}] = true
				return walk(graph, r, c, -1, 0, vis, 0) //upwards
			}
		}
	}
	return 0
}

func walk(graph [][]rune, r, c, dr, dc int, vis map[[4]int]bool, turns int) int {
	if r+dr < 0 || c+dc < 0 || r+dr >= len(graph) || c+dc >= len(graph[r]) {
		return 0
	}
	if turns > 10000 {
		return 1
	}
	// fmt.Printf("%v %v %v %v\n", r, c, dr, dc)
	// hack to see if they are stuck in a loop
	// if vis[[4]int{r + dr, c + dc, dr, dc}] {
	// 	return 1
	// }

	if graph[r+dr][c+dc] == '#' {
		//change dir
		var newDr, newDc int
		for i, dirs := range directions {
			if dr == dirs[0] && dc == dirs[1] {
				newDr = directions[(i+1)%len(directions)][0] // circle back
				newDc = directions[(i+1)%len(directions)][1] // circle back
			}
		}
		vis[[4]int{r + newDr, c + newDc, newDr, newDc}] = true
		return walk(graph, r+newDr, c+newDc, newDr, newDc, vis, turns+1)
	}
	vis[[4]int{r + dr, c + dc, dr, dc}] = true
	return walk(graph, r+dr, c+dc, dr, dc, vis, turns+1)
}
