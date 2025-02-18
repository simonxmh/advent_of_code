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
	// order := make(map[int]int)
	// visited := make(map[int]bool)

	scanRelation := true
	for r.Scan() {
		// relations
		a := strings.Split(r.Text(), "|")

		// fmt.Println(a, len(a))
		if len(a) != 2 {
			//convert to scanning the ordering
			scanRelation = false
			// construct with map walk
			// for node := range relation {
			// 	if !visited[node] {
			// 		start := Node{val: node, lvl: 0}
			// 		walk(relation, start, order, visited)
			// 	}
			// }
			// for k, v := range relation {
			// 	fmt.Printf("%v, %v", k, v)
			// }
		}

		if scanRelation {
			k := get(strconv.Atoi(a[0]))
			v := get(strconv.Atoi(a[1]))

			relation[k] = append(relation[k], v)
		} else {
			b := strings.Split(r.Text(), ",")
			if len(b) == 1 && b[0] == "" {
				continue
			}
			sorted := true
			for i, v := range b {
				vv := get(strconv.Atoi(v))

				// if previoous instances contain future
				for _, pre := range b[:i] {
					p := get(strconv.Atoi(pre))
					if slices.Contains(relation[vv], p) {
						sorted = false
					}
				}
			}
			if sorted {
				// println(b[len(b)/2])
				sum += get(strconv.Atoi(b[len(b)/2]))
			}

			// custom comparator
			// bb := b
			// sort.SliceStable(bb, func(i, j int) bool {
			// 	return order[i] < order[j]
			// })
			// if reflect.DeepEqual(b, bb) {
			// 	fmt.Printf("%v", b)
			// }
		}

	}
	fmt.Println(sum)
}

type Node struct {
	val, lvl int
}

func walk(relation map[int][]int, start Node, order map[int]int, visited map[int]bool) {
	q := []Node{start}

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		order[n.val] = n.lvl

		fmt.Printf("%v\n", relation[n.val])
		for _, nei := range relation[n.val] {
			if !visited[nei] {
				visited[nei] = true
				q = append(q, Node{val: nei, lvl: n.lvl + 1})
			}
		}

	}
}
