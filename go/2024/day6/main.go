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

var directions = [][2]int{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1}, // Up -> right -> down -> left
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

	vis := make(map[[2]int]bool)

	for r := range graph {
		for c := range graph[r] {
			if graph[r][c] == '^' {
				vis[[2]int{r, c}] = true
				walk(graph, r, c, -1, 0, vis) //upwards
			}
		}
	}
	fmt.Printf("%v", len(vis))
}

func walk(graph [][]rune, r, c, dr, dc int, vis map[[2]int]bool) {
	if r+dr < 0 || c+dc < 0 || r+dr >= len(graph) || c+dc >= len(graph[r]) {
		return
	}
	fmt.Printf("%v %v %v %v\n", r, c, dr, dc)

	if graph[r+dr][c+dc] == '#' {
		//change dir
		var newDr, newDc int
		for i, dirs := range directions {
			if dr == dirs[0] && dc == dirs[1] {
				newDr = directions[(i+1)%len(directions)][0] // circle back
				newDc = directions[(i+1)%len(directions)][1] // circle back
				fmt.Printf("newdr: %v %v\n", newDr, newDc)
			}
		}
		vis[[2]int{r + newDr, c + newDc}] = true
		walk(graph, r+newDr, c+newDc, newDr, newDc, vis)
	} else {
		vis[[2]int{r + dr, c + dc}] = true
		walk(graph, r+dr, c+dc, dr, dc, vis)
	}
	return
}
