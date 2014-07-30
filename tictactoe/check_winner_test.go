package tictactoe

import "testing"
import "reflect"

func TestTranspose(t *testing.T) {
  initial_board := [][]string{ []string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"} }
  transposed_board := [][]string{ []string{"X", "-", "-"}, []string{"X", "O", "X"}, []string{"-", "O", "O"} }
  tran := transpose(initial_board)
  if !reflect.DeepEqual(tran, transposed_board) {
    t.Error("Expected", transposed_board, "got", tran)
  } 
}

func TestTransposeDiagonal(t *testing.T) {
  initial_board := [][]string{ []string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"} }
  transposed_board := [][]string{ []string{"X", "O", "O"}, []string{"-", "O", "-"} }
  tran := transpose_diagonal(initial_board)
  if !reflect.DeepEqual(tran, transposed_board) {
    t.Error("Expected", transposed_board, "got", tran)
  } 
}

func TestHorizontalWin(t *testing.T) {
  win_x := []string{ "X", "X", "X", "O", "-", "O", "-", "-", "-" }
  win_x_middle := []string{ "-", "-", "O", "X", "X", "X", "O", "-", "-" }
  win_x_end := []string{ "-", "-", "-", "O", "-", "O", "X", "X", "X" }
  win_o := []string{ "O", "O", "O", "X", "X", "-", "-", "-", "-" }
  win_o_middle := []string{ "-", "-", "X", "O", "O", "O", "X", "-", "-" }
  win_o_end := []string{ "-", "-", "X", "X", "-", "-", "O", "O", "O" }
  winners_x := [][]string{ win_x, win_x_middle, win_x_end }
  winners_o := [][]string{ win_o, win_o_middle, win_o_end }
  for _, result := range winners_x {
    if HorizontalWin(result, "X") == false {
      t.Error("Expected true")
    }
  }

  for _, result := range winners_o {
    if HorizontalWin(result, "O") == false {
      t.Error("Expected true")
    }
  }
}

func TestVerticalWin(t *testing.T) {
  win_x := []string{ "X", "O", "-", "X", "O", "O", "X", "-", "-" }
  win_x_middle := []string{ "-", "X", "-", "O", "X", "O", "-", "X", "-" }
  win_o_middle := []string{ "-", "O", "X", "-", "O", "-", "X", "O", "-" }
  win_o_end := []string{ "-", "-", "O", "X", "X", "O", "X", "-", "O" }
  winners_x := [][]string{ win_x, win_x_middle }
  winners_o := [][]string{ win_o_middle, win_o_end }
  for _, result := range winners_x {
    if VerticalWin(result, "X") == false {
      t.Error("Expected true")
    }
  }

  for _, result := range winners_o {
    if VerticalWin(result, "O") == false {
      t.Error("Expected true")
    }
  }
}

func TestDiagonalWin(t *testing.T) {
  win_x := []string{ "X", "O", "O", "-", "X", "-", "O", "X", "X" }
  if DiagonalWin(win_x, "X") == false {
    t.Error("Expected True")
  }
}
