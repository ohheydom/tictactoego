package tictactoe

import "testing"

func TestMiniMax(t *testing.T) {
	winX := []string{"X", "-", "-", "X", "O", "O", "-", "-", "-"}
	winO := []string{"O", "O", "-", "X", "X", "-", "O", "-", "X"}
	results := make(map[string]int)
	gX := GameBoard{Board: winX, Turn: "X", Count: 4, Result: results}
	gO := GameBoard{Board: winO, Turn: "O", Count: 6, Result: results}
	expectedMinimaxWinX := 99
	minimaxWinX := MiniMax(gX, 0, true)
	expectedMinimaxWinO := -99
	minimaxWinO := MiniMax(gO, 0, false)

	if expectedMinimaxWinX != minimaxWinX {
		t.Error("Expected: ", expectedMinimaxWinX, "Got: ", minimaxWinX)
	}

	if expectedMinimaxWinO != minimaxWinO {
		t.Error("Expected: ", expectedMinimaxWinO, "Got: ", minimaxWinO)
	}
}

func BenchmarkBestMoveMiniMax(b *testing.B) {
	board := []string{"-", "-", "-", "X", "-", "-", "-", "-", "-"}
	results := make(map[string]int)
	g := GameBoard{Board: board, Turn: "O", Count: 1, Result: results}
	for i := 0; i < b.N; i++ {
		g.MiniMaxBestMove()
	}
}
