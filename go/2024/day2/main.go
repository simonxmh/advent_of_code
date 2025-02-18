package main

import (
	"bufio"
	"fmt"
	"os"
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

	s := bufio.NewScanner(f)

	safe := 0
	ds := 0
	for s.Scan() {
		report := strings.Fields(s.Text())

		pr := parse(report)

		if isSafe(pr) {
			safe += 1
		} else if isDampenedSafe(pr) {
			ds += 1
		}
	}

	fmt.Printf("safe: %d\n", safe)
	fmt.Printf("total after dampened: %d\n", ds+safe)
}

func parse(xs []string) []int {
	r := make([]int, 0, len(xs))
	for _, si := range xs {
		r = append(r, get(strconv.Atoi(si)))
	}
	return r
}

func isSafe(report []int) bool {
	inc := false
	dec := false
	pre := report[0]
	for _, cur := range report[1:] {
		if cur > pre {
			if dec {
				return false
			}
			if cur-pre > 3 {
				return false
			}
			pre = cur
			inc = true
		} else if cur < pre {
			if inc {
				return false
			}
			if cur-pre < -3 {
				return false
			}
			pre = cur
			dec = true
		} else {
			return false
		}
	}

	return true
}

func isDampenedSafe(report []int) bool {
	for i := range report {
		removed := make([]int, 0, len(report)-1)
		removed = append(removed, report[:i]...)
		removed = append(removed, report[i+1:]...)
		if isSafe(removed) {
			return true
		}
	}
	return false
}
