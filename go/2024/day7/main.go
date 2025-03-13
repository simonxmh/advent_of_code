package main

import (
	"bufio"
	"log"
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
	if len(os.Args) != 2 {
		log.Fatal("error with filepath")
	}

	f := get(os.Open(os.Args[1]))
	r := bufio.NewScanner(f)

	sum := 0

	for r.Scan() {
		line := r.Text()

		columns := strings.Fields(line)

		wantedTotal := get(strconv.Atoi(columns[0][:len(columns[0])-1]))
		rest := columns[1:]
		restInt := make([]int, len(rest))
		for i := range rest {
			restInt[i] = get(strconv.Atoi(rest[i]))
		}

		if canMake(wantedTotal, restInt, 0, 0) {
			sum += wantedTotal
		}

	}
	println(sum)

}

func canMake(wantedTotal int, rest []int, currentIndex, currentTotal int) bool {
	if currentIndex >= len(rest) {
		return currentTotal == wantedTotal
	}

	// Base case: if we've exceeded the target, no need to continue
	if currentTotal > wantedTotal {
		return false
	}

	if canMake(wantedTotal, rest, currentIndex+1, currentTotal+rest[currentIndex]) ||
		canMake(wantedTotal, rest, currentIndex+1, currentTotal*rest[currentIndex]) ||
		(1 < currentIndex && currentIndex+2 < len(rest) && canMake(wantedTotal, rest, currentIndex+2, currentTotal-rest[currentIndex-1]+rest[currentIndex-1]*rest[currentIndex])) {
		return true
	}
	return false
}
