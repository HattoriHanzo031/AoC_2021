package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	S1 = uint8(1 << 0)
	S2 = uint8(1 << 1)
	S3 = uint8(1 << 2)
	S4 = uint8(1 << 3)
	S5 = uint8(1 << 4)
	S6 = uint8(1 << 5)
	S7 = uint8(1 << 6)
)

var digitMap = map[uint8]int{
	S1 | S2 | S3 | S4 | S5 | S6:      0,
	S2 | S3:                          1,
	S1 | S2 | S7 | S5 | S4:           2,
	S1 | S2 | S3 | S4 | S7:           3,
	S2 | S3 | S6 | S7:                4,
	S1 | S3 | S4 | S6 | S7:           5,
	S1 | S3 | S4 | S5 | S6 | S7:      6,
	S1 | S2 | S3:                     7,
	S1 | S2 | S3 | S4 | S5 | S6 | S7: 8,
	S1 | S2 | S3 | S4 | S6 | S7:      9,
}

var masks = map[int]uint8{
	2: S2 | S3,
	3: S1 | S2 | S3,
	4: S2 | S3 | S6 | S7,
	5: S1 | S4 | S7,
	6: S1 | S3 | S4 | S6,
	7: 0,
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	count := 0
	total := 0
	scaner := bufio.NewScanner(f)
	for scaner.Scan() {
		input := strings.Split(scaner.Text(), "|")
		signals := strings.Fields(input[0])
		digits := strings.Fields(input[1])

		count += countKnownDigits(digits)

		d := parseDigits(parseSegments(signals), digits)
		fmt.Println(d)
		total += d
	}

	fmt.Println("COUNT:", count)
	fmt.Println("TOTAL:", total)
}

func parseSegments(signals []string) map[rune]uint8 {
	var segments = map[rune]uint8{
		'a': 0b01111111,
		'b': 0b01111111,
		'c': 0b01111111,
		'd': 0b01111111,
		'e': 0b01111111,
		'f': 0b01111111,
		'g': 0b01111111,
	}

	for _, signal := range signals {
		mask := masks[len(signal)]
		for segment, v := range segments {
			if in(signal, segment) {
				switch len(signal) {
				case 2, 3, 4:
					segments[segment] = v & mask
				}
			} else {
				segments[segment] = v & (^mask)
			}
		}
	}

	// Remaining unambiguous segments can be found by looking at other definite segments,
	// but manually looking at the data showed that only these two cases appear
	for segment, v := range segments {
		if v == 0b00011000 {
			segments[segment] = 0b00001000
		}
		if v == 0b00000110 {
			segments[segment] = 0b00000100
		}
	}
	return segments
}

func parseDigits(segments map[rune]uint8, digits []string) int {
	d := 0
	for _, digit := range digits {
		final := uint8(0)
		for _, d := range digit {
			final |= segments[d]
		}
		d *= 10
		d += digitMap[final]

	}
	return d
}

func countKnownDigits(digits []string) int {
	count := 0
	for _, digit := range digits {
		switch len(digit) {
		case 2, 3, 4, 7:
			count++
		}
	}
	return count
}

func in(str string, r rune) bool {
	for _, s := range str {
		if r == s {
			return true
		}
	}
	return false
}
