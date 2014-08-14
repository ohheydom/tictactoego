package tictactoe

import "testing"

func TestMiniMax(t *testing.T) {
  win_x := []string{"X", "-", "-", "X", "O", "O", "-", "-", "-"}
  win_o := []string{"O", "O", "-", "X", "X", "-", "O", "-", "X"}
  g_x := GameBoard{Board: win_x, Turn: "X", Count: 4}
  g_o := GameBoard{Board: win_o, Turn: "O", Count: 6}
  expected_minimax_win_x := 99
  minimax_win_x := MiniMax(g_x, 0, true)
  expected_minimax_win_o := -99
  minimax_win_o := MiniMax(g_o, 0, false)

  if expected_minimax_win_x != minimax_win_x {
    t.Error("Expected: ", expected_minimax_win_x, "Got: ", minimax_win_x)
  }

  if expected_minimax_win_o != minimax_win_o {
    t.Error("Expected: ", expected_minimax_win_o, "Got: ", minimax_win_o)
  }
}

func BenchmarkBestMoveMiniMax(b *testing.B) {
  board := []string{"-", "-", "-", "X", "-", "-", "-", "-", "-"}
  g := GameBoard{Board: board, Turn: "O"}
  for i := 0; i < b.N; i++ {
    g.MiniMaxBestMove()
  }
}
