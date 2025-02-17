package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	f := get(os.Open("./example.txt"))

	var xs, ys []int

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()

		columns := strings.Fields(line)

		xs = append(xs, get(strconv.Atoi(columns[0])))
		ys = append(ys, get(strconv.Atoi(columns[1])))
	}

	sort.Ints(xs)
	sort.Ints(ys)

	// get distance
	dist := 0
	for i, x := range xs {
		d := x - ys[i]
		if d < 0 {
			d = -1 * d
		}
		dist += d
	}
	fmt.Println(dist)

	// get similarity
	sim := 0
	freqs := make(map[int]int)
	for _, y := range ys {
		freqs[y] += 1
	}
	for _, x := range xs {
		sim += x * freqs[x]
	}
	fmt.Println(sim)
}
