package tictactoe

import "testing"
import "reflect"

func TestTranspose(t *testing.T) {
  initial_board := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  expected_transposed_board := [][]string{[]string{"X", "-", "-"}, []string{"X", "O", "X"}, []string{"-", "O", "O"}}
  calculated_transposed_board := Transpose(initial_board)
  if !reflect.DeepEqual(expected_transposed_board, calculated_transposed_board) {
    t.Error("Expected", expected_transposed_board, "got", calculated_transposed_board)
  } 

  initial_uneven_board := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}}
  expected_uneven_tran := [][]string{[]string{"X", "-"}, []string{"X", "O"}, []string{"-", "O"}}
  calculated_uneven_tran := Transpose(initial_uneven_board)
  if !reflect.DeepEqual(expected_uneven_tran, calculated_uneven_tran) {
    t.Error("Expected", expected_uneven_tran, "got", calculated_uneven_tran)
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

func TestAll(t *testing.T) {
  board := [][]string{[]string{"X", "X", "X"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  expected_value := true
  received_value := All(board, "X")
  if expected_value != received_value {
    t.Error("Expected: ", expected_value, "Received: ", received_value)
  }
}

func BenchmarkAll(b *testing.B) {
  board := [][]string{[]string{"X", "X", "X"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  for i := 0; i < b.N; i++ {
    All(board, "X")
  }
}

