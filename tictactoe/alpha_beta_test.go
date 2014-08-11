package tictactoe

import "testing"

func TestAlphaBeta(t *testing.T) {
  win_x := []string{"X", "-", "-", "X", "O", "O", "-", "-", "-"}
  win_o := []string{"O", "O", "-", "X", "X", "-", "O", "-", "X"}
  g_x := GameBoard{Board: win_x, Turn: "X"}
  g_o := GameBoard{Board: win_o, Turn: "O"}
  expected_minimax_win_x := 99
  minimax_win_x := AlphaBeta(g_x, 0, -1000, 1000, true)
  expected_minimax_win_o := -99
  minimax_win_o := AlphaBeta(g_o, 0, -1000, 1000, false)

  if expected_minimax_win_x != minimax_win_x {
    t.Error("Expected: ", expected_minimax_win_x, "Got: ", minimax_win_x)
  }

  if expected_minimax_win_o != minimax_win_o {
    t.Error("Expected: ", expected_minimax_win_o, "Got: ", minimax_win_o)
  }
}

func BenchmarkBestMoveAlphaBeta(b *testing.B) {
  board := []string{"-", "-", "-", "X", "-", "-", "-", "-", "-"}
  g := GameBoard{Board: board, Turn: "O"}
  for i := 0; i < b.N; i++ {
    g.AlphaBetaBestMove()
  }
}
