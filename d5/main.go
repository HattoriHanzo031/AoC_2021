package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Line struct {
	from, to Coordinate
}

type FloorMap map[Coordinate]int

func (floorMap FloorMap) Print() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if num, ok := floorMap[Coordinate{y, x}]; ok {
				fmt.Print(num)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func ParseCoordinate(s string) Coordinate {
	xy := strings.Split(s, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])

	return Coordinate{x, y}
}

func ParseInput(filename string) []Line {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scaner := bufio.NewScanner(f)

	var lines []Line
	for scaner.Scan() {
		text := scaner.Text()

		fields := strings.Fields(text)
		lines = append(lines, Line{ParseCoordinate(fields[0]), ParseCoordinate(fields[2])})
	}
	return lines
}

// Returns iterator function that counts from start to end.
// The other return value is true until end is reached.
// When the end is reached the next call to iterator function will return end value and false.
func Iter(start, end int) func() (int, bool) {
	increment := 1
	if start > end {
		increment = -1
	}
	current := start - increment
	return func() (int, bool) {
		current += increment
		if current == end+increment {
			current = end
			return current, false
		}
		return current, true
	}
}

func main() {
	lines := ParseInput("input.txt")

	floorMap := make(FloorMap)

	for _, line := range lines {
		//Uncomment for Part 1
		//if line.from.x != line.to.x && line.from.y != line.to.y {
		//	continue
		//}
		itX, itY := Iter(line.from.x, line.to.x), Iter(line.from.y, line.to.y)
		x, okX := itX()
		y, okY := itY()
		for okX || okY {
			floorMap[Coordinate{x, y}]++
			x, okX = itX()
			y, okY = itY()
		}
	}

	count := 0
	for _, num := range floorMap {
		if num <= 1 {
			continue
		}
		count++
	}

	floorMap.Print()
	fmt.Println("Solution: ", count)
}
