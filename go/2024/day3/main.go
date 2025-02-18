package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
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

	cont := string(get(os.ReadFile(os.Args[1])))

	re := regexp.MustCompile(`^mul\((\d+),(\d+)\)`)
	enableRe := regexp.MustCompile(`^do\(\)`)
	disableRe := regexp.MustCompile(`^don't\(\)`)

	total := 0

	enabled := true
	i := 0
	for i < len(cont) {
		sub := cont[i:]

		if enabled && re.MatchString(sub) {
			matches := re.FindStringSubmatch(sub)
			a, b := parseMatch(string(matches[0]))
			total += a * b
		}
		if enableRe.MatchString(sub) {
			enabled = true
			i += len("do()")
			continue
		}
		if disableRe.MatchString(sub) {
			enabled = false
			i += len("don't()")
			continue
		}
		i += 1
	}
	println(total)
}

func parseMatch(s string) (int, int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	match := re.FindStringSubmatch(s)
	if match == nil {
		panic("invalid match format: " + s)
	}
	a := get(strconv.Atoi(string(match[1])))
	b := get(strconv.Atoi(string(match[2])))
	println(a, b)

	return a, b
}
