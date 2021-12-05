package main

import "fmt"

type Field struct {
	column, row int
}

type Board struct {
	Fields     map[int]Field
	RowHits    [5]int
	ColumnHits [5]int
	bingo      bool
}

// Returns true if board has bingo
func (b *Board) Mark(num int) bool {
	field, ok := b.Fields[num]
	if !ok {
		return false
	}

	// Deleting hit numbers so they don't add up in board score
	// Other option would be to extend the Field struct with "hit" field and skip them during score calculation
	delete(b.Fields, num)

	b.ColumnHits[field.column]++
	b.RowHits[field.row]++

	b.bingo = b.ColumnHits[field.column] == 5 || b.RowHits[field.row] == 5 // TODO: determin programmatically
	return b.bingo
}

func (b *Board) Score() int {
	sum := 0
	for key := range b.Fields {
		sum += key
	}
	return sum
}

// Helper for debugging
func (b *Board) Print() {
	var fields [5][5]int
	for v, f := range b.Fields {
		fields[f.row][f.column] = v
	}
	for _, row := range fields {
		fmt.Println(row)
	}
	fmt.Println()
}
