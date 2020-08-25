package main

import (
	"fmt"
	"github.com/Kwbmm/TicTacGo/board"
)

func main() {
	fmt.Println("Welcome to TicTacGo")

	currentPlayer := 'O'
	nextPlayer := 'X'
	var b board.Board

	for {
		b.Print()

		fmt.Printf("Player '%c' turn\nChoose a cell from 1 to 9\n\n", currentPlayer)
		
		choice := 0
		fmt.Scanf("%d", &choice)
		
		if choice < 1 || choice > 9 {
			fmt.Println("Invalid choice!")
			continue
		}
		
		if !b.MakeMove(choice-1, currentPlayer) {
			fmt.Println("Invalid move!")
			continue
		}

		if b.HasWinner(choice-1, currentPlayer) {
			b.Print()

			fmt.Printf("Player '%c' wins the game!!\n", currentPlayer)
			return
		}

		if !b.HasEmptyCells() {
			fmt.Println("Game is draw")
			return
		}
		currentPlayer, nextPlayer = nextPlayer, currentPlayer
	}

}
