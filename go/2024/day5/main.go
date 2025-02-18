package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	if len(os.Args) != 2 {
		log.Fatal("error with filepath")
	}

	f := get(os.Open(os.Args[1]))
	r := bufio.NewScanner(f)

	relation := make(map[int][]int)
	sum := 0
	p2sum := 0

	scanRelation := true
	for r.Scan() {
		// relations
		a := strings.Split(r.Text(), "|")

		if len(a) != 2 {
			//convert to scanning the ordering
			scanRelation = false
		}

		if scanRelation {
			k := get(strconv.Atoi(a[0]))
			v := get(strconv.Atoi(a[1]))

			relation[k] = append(relation[k], v)
			continue
		}
		b := strings.Split(r.Text(), ",")
		if len(b) == 1 && b[0] == "" {
			continue
		}
		sorted := true
		for i, v := range b {
			vv := get(strconv.Atoi(v))

			// if previoous instances contain future
			for j, pre := range b[:i] {
				p := get(strconv.Atoi(pre))
				if slices.Contains(relation[vv], p) {
					sorted = false
					// just swap for correct ordering, effectively bubble sort
					b[i], b[j] = b[j], b[i]
				}
			}
		}
		if sorted {
			// println(b[len(b)/2])
			sum += get(strconv.Atoi(b[len(b)/2]))
		} else {
			// part2
			p2sum += get(strconv.Atoi(b[len(b)/2]))
			fmt.Printf("%v", b)
		}
	}
	fmt.Println(sum)
	fmt.Println(p2sum)
}

type Node struct {
	val, lvl int
}
