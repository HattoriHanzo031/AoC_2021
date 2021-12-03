package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1(values *[]string) int {
	numLines := len(*values)
	numBits := len((*values)[0])

	// number of ones for each bit possition
	onesCounts := make([]int, numBits)
	for _, val := range *values {
		for bit, c := range val {
			onesCounts[bit] += int(c - '0')
		}
	}

	var gamaRate, epsilonRate int

	// For each bit position, if number of ones is bigger than number of zeros
	// epsilonRate gets 1 at that position, othervise, gamaRate gets 1 at that position
	for _, oneBits := range onesCounts {
		gamaRate <<= 1
		epsilonRate <<= 1
		if oneBits > numLines-oneBits {
			epsilonRate += 1
		} else {
			gamaRate += 1
		}
	}

	fmt.Println(onesCounts, gamaRate, epsilonRate, gamaRate*epsilonRate)
	return gamaRate * epsilonRate
}

// Splits the input slice and returns two slices, numbers with majority bit at 'bit' position
// and numbers with minority bit at bit position, in that order
func split(values *[]string, bit int) ([]string, []string) {
	var ones, zeros []string
	for _, val := range *values {
		if val[bit] == '1' {
			ones = append(ones, val)
		} else {
			zeros = append(zeros, val)
		}
	}

	if len(zeros) > len(ones) {
		return zeros, ones
	} else {
		return ones, zeros
	}
}

func Part2(values *[]string) int {
	numBits := len((*values)[0])

	// For every bit position set o2 to majority bit slice and co2 to minority bit slice
	// until there is only one number in both o2 and co2
	o2, co2 := split(values, 0)
	for bit := 1; bit < numBits; bit++ {
		if len(o2) > 1 {
			o2, _ = split(&o2, bit)
		}
		if len(co2) > 1 {
			_, co2 = split(&co2, bit)
		}
	}

	o2Val, _ := strconv.ParseUint(o2[0], 2, 64)
	co2Val, _ := strconv.ParseUint(co2[0], 2, 64)

	fmt.Println(o2, o2Val, co2, co2Val, o2Val*co2Val)
	return int(o2Val * co2Val)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scaner := bufio.NewScanner(f)
	var values []string
	for scaner.Scan() {
		values = append(values, scaner.Text())
	}

	fmt.Println("Part 1:", Part1(&values))
	fmt.Println("Part 2:", Part2(&values))
}
