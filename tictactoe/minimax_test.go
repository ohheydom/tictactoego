package tictactoe

import "testing"

func TestMiniMax(t *testing.T) {
  win_x := []string{"X", "-", "-", "X", "O", "O", "-", "-", "-"}
  win_o := []string{"O", "O", "-", "X", "X", "-", "O", "-", "X"}
  g_x := GameBoard{Board: win_x, Turn: "X"}
  g_o := GameBoard{Board: win_o, Turn: "O"}
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
