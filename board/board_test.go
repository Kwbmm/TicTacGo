package board

import (
	"strconv"
	"strings"
	"testing"
)

var emptyBoard = Board{
	State: [9]Cell{}}

var partialBoard = Board{
	State: [9]Cell{' ', 'O', ' ', 'X', 'X', ' ', 'O', ' ', ' '}}

var fullBoardNoWinner = Board{
	State: [9]Cell{'O', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'X'}}

var fullBoardWinnerIsO = Board{
	State: [9]Cell{'X', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'X'}}

var fullBoardWinnerIsX = Board{
	State: [9]Cell{'O', 'X', 'X', 'O', 'X', 'O', 'X', 'X', 'O'}}

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
		npCell := Cell(nextPlayer)
		pCell := Cell(player)
		boards["123"] = make(map[rune]Board)
		boards["123"][player] = Board{
			State: [9]Cell{pCell, pCell, pCell, npCell, pCell, npCell, pCell, npCell, npCell}}
		boards["147"] = make(map[rune]Board)
		boards["147"][player] = Board{
			State: [9]Cell{pCell, npCell, npCell, pCell, pCell, npCell, pCell, npCell, pCell}}
		boards["159"] = make(map[rune]Board)
		boards["159"][player] = Board{
			State: [9]Cell{pCell, npCell, npCell, npCell, pCell, npCell, npCell, npCell, pCell}}
		boards["258"] = make(map[rune]Board)
		boards["258"][player] = Board{
			State: [9]Cell{npCell, pCell, npCell, pCell, pCell, npCell, npCell, pCell, pCell}}
		boards["357"] = make(map[rune]Board)
		boards["357"][player] = Board{
			State: [9]Cell{npCell, npCell, pCell, npCell, pCell, npCell, pCell, npCell, npCell}}
		boards["369"] = make(map[rune]Board)
		boards["369"][player] = Board{
			State: [9]Cell{npCell, npCell, pCell, pCell, npCell, pCell, npCell, pCell, pCell}}
		boards["456"] = make(map[rune]Board)
		boards["456"][player] = Board{
			State: [9]Cell{npCell, npCell, pCell, pCell, pCell, pCell, npCell, pCell, npCell}}
		boards["789"] = make(map[rune]Board)
		boards["789"][player] = Board{
			State: [9]Cell{npCell, npCell, pCell, pCell, npCell, npCell, pCell, pCell, pCell}}
	}

	//Test + Verify
	for key, playerBoards := range boards {
		for player, board := range playerBoards {
			result, _ := board.HasWinner()
			if result != expectedResult {
				t.Errorf("Expecting winner '%c' with sequence %s", player, key)
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
	result := board.State[0].setValue(expectedPlayerSet)
	playerSet := board.State[0]

	//Verify
	if result != expectedResult || playerSet != Cell(expectedPlayerSet) {
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
	result := board.State[1].setValue(expectedPlayerSet)
	playerSet := board.State[1]

	//Verify
	if result != expectedResult || playerSet != Cell(expectedPlayerSet) {
		t.Errorf("Expecting result = %t and player = %c\nGot result = %t and player = %c",
			expectedResult, expectedPlayerSet,
			result, playerSet)
	}
}
