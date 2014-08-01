package tictactoe

import "testing"
import "reflect"

func TestTranspose(t *testing.T) {
  initial_board := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  transposed_board := [][]string{[]string{"X", "-", "-"}, []string{"X", "O", "X"}, []string{"-", "O", "O"}}
  tran := Transpose(initial_board)
  if !reflect.DeepEqual(tran, transposed_board) {
    t.Error("Expected", transposed_board, "got", tran)
  } 
}

func TestTransposeDiagonal(t *testing.T) {
  initial_board := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  transposed_board := [][]string{[]string{"X", "O", "O"}, []string{"-", "O", "-"}}
  tran := TransposeDiagonal(initial_board)
  if !reflect.DeepEqual(tran, transposed_board) {
    t.Error("Expected", transposed_board, "got", tran)
  } 
}
