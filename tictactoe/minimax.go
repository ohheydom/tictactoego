package tictactoe

import "strings"

func MiniMax(g GameBoard, depth int, maxPlayer bool) int {
	boardString := strings.Join(g.Board, "")
	val, ok := g.Result[boardString]
	if ok {
		return val
	}
	if g.Win() {
		if g.PreviousTurn() == "X" {
			return 100 - depth
		}
		return depth - 100
	}
	if len(g.RemainingIndices()) == 0 {
		return 0
	}

	if maxPlayer {
		bestValue := -1000
		for _, move := range g.RemainingIndices() {
			g.Move(move, "X")
			value := MiniMax(g, depth+1, false)
			g.UndoMove(move)
			bestValue = Max(bestValue, value)
		}
		g.Result[boardString] = bestValue
		return bestValue
	} else {
		bestValue := 1000
		for _, move := range g.RemainingIndices() {
			g.Move(move, "O")
			value := MiniMax(g, depth+1, true)
			g.UndoMove(move)
			bestValue = Min(bestValue, value)
		}
		g.Result[boardString] = bestValue
		return bestValue
	}
}

func (g GameBoard) MiniMaxMoves() [][]int {
	var score [][]int
	for _, move := range g.RemainingIndices() {
		g.Move(move, g.Turn)
		arr := []int{MiniMax(g, 0, true), move}
		score = append(score, arr)
		g.UndoMove(move)
	}
	return score
}

func (g GameBoard) MiniMaxBestMove() int {
	if g.Turn == "X" {
		return MaxBy(g.MiniMaxMoves(), 0)[1]
	}
	return MinBy(g.MiniMaxMoves(), 0)[1]
}
