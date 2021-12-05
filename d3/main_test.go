package main

import "testing"

type Table struct {
	in   []string
	res1 int
	res2 int
}

func testData() []Table {
	return []Table{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			198,
			230,
		},
	}
}
func TestPart1(t *testing.T) {
	for _, table := range testData() {
		result := Part1(&table.in)
		if result != table.res1 {
			t.Fail()
		}
	}
}

func TestPart2(t *testing.T) {
	for _, table := range testData() {
		result := Part2(&table.in)
		if result != table.res2 {
			t.Fail()
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1(&testData()[0].in)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2(&testData()[0].in)
	}
}
