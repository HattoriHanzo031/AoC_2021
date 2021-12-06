package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := bufio.NewReader(f).ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1] // Remove /n at the end (need to find a better way to parse this :))

	timers := strings.Split(input, ",")
	intTimers := make([]int, len(timers))
	for i, timer := range timers {
		t, err := strconv.Atoi(timer)
		if err != nil {
			panic(err)
		}
		intTimers[i] = t
	}
	return intTimers
}

// How to avoid this?
var cache map[int]int

// If one fish is spawn today, how many fishes will be spawn (including this one) after "days" days
func fishAfterDays(days int) int {
	if number, ok := cache[days]; ok {
		return number
	}

	number := 1
	daysRemaining := days - 9 // next fish is spawn after 9 days

	for daysRemaining > 0 {
		number += fishAfterDays(daysRemaining)
		daysRemaining -= 7 // all other fishes are spawn after each 7 days
	}

	// Cache the solution not to calculate the same thing again
	cache[days] = number
	return number
}

func main() {
	timers := parseInput("input.txt")

	solver := func(days int) int {
		result := 0
		cache = make(map[int]int)
		for _, timer := range timers {
			// Add (9 - timer) to calculate when this fish was spawn
			result += fishAfterDays(days + 9 - timer)
		}
		return result
	}

	fmt.Println("Part 1:", solver(80))
	fmt.Println("Part 2:", solver(256))
}
