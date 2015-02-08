package tictactoe

import "strings"

func AlphaBeta(g GameBoard, depth int, alpha int, beta int, maxPlayer bool) int {
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
		for _, move := range g.RemainingIndices() {
			g.Move(move, "X")
			alpha = Max(alpha, AlphaBeta(g, depth+1, alpha, beta, false))
			g.UndoMove(move)
			if beta <= alpha {
				g.Result[boardString] = alpha
				return alpha
			}
		}
		return alpha
	} else {
		for _, move := range g.RemainingIndices() {
			g.Move(move, "O")
			beta = Min(beta, AlphaBeta(g, depth+1, alpha, beta, true))
			g.UndoMove(move)
			if beta <= alpha {
				return beta
			}
		}
		return beta
	}
}

func (g GameBoard) AlphaBetaBestMove() int {
	if g.Turn == "X" {
		return MaxBy(g.AlphaBetaMoves(), 0)[1]
	}
	return MinBy(g.AlphaBetaMoves(), 0)[1]
}

func (g GameBoard) AlphaBetaMoves() (score [][]int) {
	score = make([][]int, 0, len(g.RemainingIndices()))
	for _, move := range g.RemainingIndices() {
		g.Move(move, g.Turn)
		arr := []int{AlphaBeta(g, 0, -1000, 1000, true), move}
		score = append(score, arr)
		g.UndoMove(move)
	}
	return
}
