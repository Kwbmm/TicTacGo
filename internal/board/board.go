package board

import (
	"fmt"
)

type Board struct {
	State [9]Cell
}

func (b *Board) MakeMove(i int, player rune) bool {
	if i < 0 || i >= len(b.State) || b.State[i] != 0 {
		return false
	}

	b.State[i] = Cell(player)

	return true
}

func (b *Board) HasWinner(index int, player rune) bool {
	for i := 0; i < 3; i++ {
		if b.checkRow(i) {
			return true
		}

		if b.checkCol(i) {
			return true
		}
	}

	cells := b.State

	if cells[0] != 0 && cells[0] == cells[4] && cells[0] == cells[8]  {
		return true
	}

	if cells[2] != 0 && cells[2] == cells[6] && cells[2] == cells[6] {
		return true
	}

	return false
}

func (b *Board) checkRow(i int) bool {
	i *= 3
	return b.State[i] != 0 && b.State[i] == b.State[i+1] && b.State[i] == b.State[i+2]
}

func (b *Board) checkCol(i int) bool {
	return b.State[i] != 0 && b.State[i] == b.State[i+3] && b.State[i] == b.State[i+6]
}

func (b *Board) HasEmptyCells() bool {
	for _, cell := range b.State {
		if cell != 'O' && cell != 'X' {
			return true
		}
	}
	return false
}

func (board *Board) printHeaderFooter() {
	fmt.Println("+-------|-------|-------+")
}

func (board *Board) printFiller() {
	fmt.Println("|       |       |       |")
}

func (b *Board) Print() {
	for i, c := range b.State {
		indexInRow := i % 3
		if c == 0 {
			c = Cell('0' + 1 + i)
		}
		switch indexInRow {
		case 0:
			b.printHeaderFooter()
			b.printFiller()
			fmt.Printf("|   %c   ", c)
		case 1:
			fmt.Printf("|   %c   |", c)
		case 2:
			fmt.Printf("   %c   |\n", c)
			b.printFiller()
		}
	}
	b.printHeaderFooter()}
