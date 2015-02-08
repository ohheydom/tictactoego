package tictactoe

import "testing"

func TestAlphaBeta(t *testing.T) {
	winX := []string{"X", "-", "-", "X", "O", "O", "-", "-", "-"}
	winO := []string{"O", "O", "-", "X", "X", "-", "O", "-", "X"}
	results := make(map[string]int)
	gX := GameBoard{Board: winX, Turn: "X", Count: 4, Result: results}
	gO := GameBoard{Board: winO, Turn: "O", Count: 6, Result: results}
	expectedMinimaxWinX := 99
	minimaxWinX := AlphaBeta(gX, 0, -1000, 1000, true)
	expectedMinimaxWinO := -99
	minimaxWinO := AlphaBeta(gO, 0, -1000, 1000, false)

	if expectedMinimaxWinX != minimaxWinX {
		t.Error("Expected: ", expectedMinimaxWinX, "Got: ", minimaxWinX)
	}

	if expectedMinimaxWinO != minimaxWinO {
		t.Error("Expected: ", expectedMinimaxWinO, "Got: ", minimaxWinO)
	}
}

func BenchmarkBestMoveAlphaBeta(b *testing.B) {
	board := []string{"-", "-", "-", "X", "-", "-", "-", "-", "-"}
	results := make(map[string]int)
	g := GameBoard{Board: board, Turn: "O", Result: results}
	for i := 0; i < b.N; i++ {
		g.AlphaBetaBestMove()
	}
}
