package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scaner := bufio.NewScanner(f)
	var depths []int
	for scaner.Scan() {
		depth, err := strconv.Atoi(scaner.Text())
		if err != nil {
			panic(err)
		}
		depths = append(depths, depth)
	}

	fmt.Println(countIncreased(depths, 1))
	fmt.Println(countIncreased(depths, 3))
}

func countIncreased(depths []int, windowSize int) int {
	count := 0
	for i := windowSize; i < len(depths); i++ {
		if depths[i] > depths[i-windowSize] {
			count++
		}
	}
	return count
}
