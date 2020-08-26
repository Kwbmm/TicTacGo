package minimax

import (
    "math"
    "github.com/Kwbmm/TicTacGo/board"
)

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func Minimax(b board.Board, isMaximizing bool) int {
    hasWinner, who := b.HasWinner()
    if !b.HasEmptyCells() || hasWinner {
        aiCell, playerCell := board.Cell('X'), board.Cell('O')
        if who == aiCell {
            return 1
        } else if who == playerCell {
            return -1
        }
        return 0
    }
    
    if isMaximizing {
        maxEval := math.MinInt32
        for _, emptyCell := range b.GetEmptyCells() {
            b.State[emptyCell] = board.Cell('X')
            score := Minimax(b, false)
            maxEval = Max(maxEval, score)
        }
        return maxEval
    }
    
    minEval := math.MaxInt32
    for _, emptyCell := range b.GetEmptyCells() {
        b.State[emptyCell] = board.Cell('O')
        score := Minimax(b, true)
        minEval = Min(minEval, score)
    }
    return minEval
}