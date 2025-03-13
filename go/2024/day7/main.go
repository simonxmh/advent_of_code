package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// get is a helper function to handle errors by panicking
func get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// processLine parses a line of input and returns the target number and slice of operands
func processLine(line string) (int, []int) {
	columns := strings.Fields(line)

	// First number is the target, with a colon at the end that we need to remove
	target := get(strconv.Atoi(columns[0][:len(columns[0])-1]))

	// Convert remaining strings to integers
	operands := make([]int, len(columns)-1)
	for i, num := range columns[1:] {
		operands[i] = get(strconv.Atoi(num))
	}

	return target, operands
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}

	f := get(os.Open(os.Args[1]))
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		target, operands := processLine(scanner.Text())
		if canMake(target, operands, 1, operands[0], "+") {
			sum += target
		}
	}

	println("Final sum:", sum)
}

// canMake determines if it's possible to create the target number using the given operands
// with addition and multiplication operations, respecting order of operations
func canMake(target int, operands []int, currentIndex, currentTotal int, path string) bool {
	// Base case: reached the end of operands
	if currentIndex >= len(operands) {
		return currentTotal == target
	}

	// Optimization: if we've exceeded the target, no need to continue
	if currentTotal > target {
		return false
	}

	// Try three possible operations:
	// 1. Simple addition with current number
	if canMake(target, operands, currentIndex+1, currentTotal+operands[currentIndex], path+"+") {
		return true
	}

	// 2. Simple multiplication with current number
	if canMake(target, operands, currentIndex+1, currentTotal*operands[currentIndex], path+"*") {
		return true
	}

	numString := fmt.Sprintf("%d%d", currentTotal, operands[currentIndex])
	if canMake(target, operands, currentIndex+1, get(strconv.Atoi(numString)), path+"|") {
		return true
	}

	return false
}
