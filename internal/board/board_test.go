package board

import (
	"strconv"
	"strings"
	"testing"
)

var emptyBoard = Board{
	BoardState: [9]Cell{
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '}}}

var partialBoard = Board{
	BoardState: [9]Cell{
		Cell{Value: ' '},
		Cell{Value: 'O'},
		Cell{Value: ' '},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: ' '},
		Cell{Value: 'O'},
		Cell{Value: ' '},
		Cell{Value: ' '}}}

var fullBoardNoWinner = Board{
	BoardState: [9]Cell{
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'}}}

var fullBoardWinnerIsO = Board{
	BoardState: [9]Cell{
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'O'},
		Cell{Value: 'X'}}}

var fullBoardWinnerIsX = Board{
	BoardState: [9]Cell{
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: 'O'}}}

func TestWhenBoardHasEmptyCellsShouldReturnTrue(t *testing.T) {
	//Given
	board := emptyBoard
	expectedResult := true

	//Test
	result := board.HasEmptyCells()

	//Verify
	if result != expectedResult {
		t.Errorf("Board is full! Expected %t but got %t", expectedResult, result)
	}

	//Given
	board = partialBoard
	expectedResult = true

	//Test
	result = board.HasEmptyCells()

	//Verify
	if result != expectedResult {
		t.Errorf("Board is full! Expected %t but got %t", expectedResult, result)
	}
}

func TestWhenBoardHasNoEmptyCellsShouldReturnFalse(t *testing.T) {
	//Given
	board := fullBoardNoWinner
	expectedResult := false

	//Test
	result := board.HasEmptyCells()

	//Verify
	if result != expectedResult {
		t.Errorf("Board is full! Expected %t but got %t", expectedResult, result)
	}
}

func TestWhenThereIsWinnerShouldReturnTrue(t *testing.T) {
	//Given
	players := [...]rune{'O', 'X'}
	var boards = make(map[string]map[rune]Board)
	expectedResult := true

	for index, player := range players {
		var nextPlayer rune
		if index == 0 {
			nextPlayer = players[1]
		} else {
			nextPlayer = players[0]
		}
		boards["123"] = make(map[rune]Board)
		boards["123"][player] = Board{
			BoardState: [9]Cell{
				{Value: player},
				{Value: player},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer}}}
		boards["147"] = make(map[rune]Board)
		boards["147"][player] = Board{
			BoardState: [9]Cell{
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: player}}}
		boards["159"] = make(map[rune]Board)
		boards["159"][player] = Board{
			BoardState: [9]Cell{
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player}}}
		boards["258"] = make(map[rune]Board)
		boards["258"][player] = Board{
			BoardState: [9]Cell{
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: player}}}
		boards["357"] = make(map[rune]Board)
		boards["357"][player] = Board{
			BoardState: [9]Cell{
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer}}}
		boards["369"] = make(map[rune]Board)
		boards["369"][player] = Board{
			BoardState: [9]Cell{
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: player}}}
		boards["456"] = make(map[rune]Board)
		boards["456"][player] = Board{
			BoardState: [9]Cell{
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: player},
				{Value: player},
				{Value: player},
				{Value: nextPlayer},
				{Value: player},
				{Value: nextPlayer}}}
		boards["789"] = make(map[rune]Board)
		boards["789"][player] = Board{
			BoardState: [9]Cell{
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: player},
				{Value: nextPlayer},
				{Value: nextPlayer},
				{Value: player},
				{Value: player},
				{Value: player}}}
		
	}

	//Test + Verify
	for key, playerBoards := range boards {
		indexes := strings.Split(key, "")
		for player, board := range playerBoards {
			for _, index := range indexes {
				i, _ := strconv.Atoi(index)
				result := board.HasWinner(i-1, player)
				if result != expectedResult {
					t.Errorf("Expecting winner '%c' with sequence %s", player, key)
				}
			}
		}
	}
}

func TestWhenSettingEmptyCellShouldReturnTrue(t *testing.T) {
	//Given
	expectedPlayerSet := 'O'
	expectedResult := true
	var board = partialBoard

	//Test
	result := board.SetCellValue(0, expectedPlayerSet)
	playerSet := board.BoardState[0].getValue()

	//Verify
	if result != expectedResult || playerSet != expectedPlayerSet {
		t.Errorf("Expecting result = %t and player = %c\nGot result = %t and player = %c",
			expectedResult, expectedPlayerSet,
			result, playerSet)
	}
}

func TestWhenSettingTakenCellShouldReturnFalse(t *testing.T) {
	//Given
	expectedPlayerSet := 'O'
	expectedResult := false
	var board = partialBoard

	//Test
	result := board.SetCellValue(1, expectedPlayerSet)
	playerSet := board.BoardState[1].getValue()

	//Verify
	if result != expectedResult || playerSet != expectedPlayerSet {
		t.Errorf("Expecting result = %t and player = %c\nGot result = %t and player = %c",
			expectedResult, expectedPlayerSet,
			result, playerSet)
	}
}
