package main

import (
	"math"
	"math/rand"
	"fmt"
	"unicode"

	"github.com/Kwbmm/TicTacGo/board"
	"github.com/Kwbmm/TicTacGo/minimax"
)

func main() {
	fmt.Println("Welcome to TicTacGo")

	var playType rune

	for {
		fmt.Println("Do you want to play with another [p]layer or agains the [A]I?")

		fmt.Scanf("%c\n", &playType)
		playType = unicode.ToLower(playType)

		if playType != 'p' && playType != 'a' {
			fmt.Printf("Invalid choice, please choose p (for player) or a (for AI)\n\n")
			continue
		}

		break
	}

	if playType == 'p' {
		playMultiplayer()
		return
	}
	playAi()
	return
}

func playAi() {
	players := [...]rune{'O', 'X', 'X', 'O'}
	i := rand.Intn(len(players[:2]))
	var b board.Board
	currentPlayer := players[i]
	nextPlayer := players[i+2]

	for {
		b.Print()

		fmt.Printf("Player '%c' turn\nChoose a cell from 1 to 9\n\n", currentPlayer)

		if currentPlayer == 'X' {
			fmt.Println("AI is choosing...")
			bestScore := math.MinInt32
			tempBoard := b
			var chosenCellIndex int
			for _, emptyCell := range b.GetEmptyCells() {
				score := minimax.Minimax(tempBoard, false)
				if bestScore < score{
					bestScore = score
					chosenCellIndex = emptyCell
				}
			}
			if !b.MakeMove(chosenCellIndex, currentPlayer) {
				fmt.Println("AI invalid move!")
				continue
			}
		} else {
			choice := 0
			fmt.Scanf("%d\n", &choice)

			if choice < 1 || choice > 9 {
				fmt.Println("Invalid choice!")
				continue
			}

			if !b.MakeMove(choice-1, currentPlayer) {
				fmt.Println("Invalid move!")
				continue
			}
		}
		if res, _ := b.HasWinner(); res {
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

func playMultiplayer() {
	var b board.Board
	currentPlayer := 'O'
	nextPlayer := 'X'
	for {
		b.Print()

		fmt.Printf("Player '%c' turn\nChoose a cell from 1 to 9\n\n", currentPlayer)

		choice := 0
		fmt.Scanf("%d\n", &choice)

		if choice < 1 || choice > 9 {
			fmt.Println("Invalid choice!")
			continue
		}

		if !b.MakeMove(choice-1, currentPlayer) {
			fmt.Println("Invalid move!")
			continue
		}

		if res, _ := b.HasWinner(); res {
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
