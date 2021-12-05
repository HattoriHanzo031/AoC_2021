package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ParseInput(filename string) ([]string, []Board) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scaner := bufio.NewScanner(f)

	// Get the drawn numbers
	scaner.Scan()
	numbers := strings.Split(scaner.Text(), ",")
	scaner.Scan() // Skip empty line

	// Get the boards from the input
	var boards []Board
	var board Board
	board.Fields = make(map[int]Field)
	var row int
	for scaner.Scan() {
		text := scaner.Text()
		// Empty line between boards in the input
		if text == "" {
			boards = append(boards, board)
			board = Board{}
			board.Fields = make(map[int]Field)
			row = 0
			continue
		}

		for column, val := range strings.Fields(text) {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			board.Fields[num] = Field{column, row}
		}
		row++
	}
	boards = append(boards, board) // append the last one
	return numbers, boards
}

func main() {
	numbers, boards := ParseInput("input.txt")

	var doOnce sync.Once
	var bingoCount int

	for _, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(boards); i++ {
			// Skip the boards with bingo
			if boards[i].bingo {
				continue
			}
			// Mark the board and check if it got bingo
			if boards[i].Mark(n) {
				// For the first bingo, calculate Part 1 solution
				doOnce.Do(func() {
					fmt.Println("Part 1:", boards[i].Score()*n)
				})

				// Count the boards with bingo and check if all boards have bingo
				bingoCount++
				if bingoCount == len(boards) {
					fmt.Println("Part 2:", boards[i].Score()*n)
					return
				}
			}
		}
	}

	panic("No solution")
}
