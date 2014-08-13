package tictactoe

import (
  "testing"
  "reflect"
)

func TestAll(t *testing.T) {
  board := [][]string{[]string{"X", "X", "X"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  expected_value := true
  received_value := All(board, "X")
  if expected_value != received_value {
    t.Error("Expected: ", expected_value, "Received: ", received_value)
  }
}

func TestReverseSlice(t *testing.T) {
  initial_board := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  expected_board := [][]string{[]string{"-", "X", "O"}, []string{"-", "O", "O"}, []string{"X", "X", "-"}}
  received_board := ReverseSlice(initial_board)
  if !reflect.DeepEqual(expected_board, received_board) {
    t.Error("Expected", expected_board, "got", received_board)
  } 
}

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

func TestMatchSlice(t *testing.T) {
  win_x := []string{"X", "X", "X", "O", "-", "O", "-", "-", "-"}
  indexes := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8}}
  expected_value := true
  received_value := MatchSlice(win_x, indexes, "X")
  if expected_value != received_value {
    t.Error("Expected True")
  }
}

func BenchmarkAll(b *testing.B) {
  board := [][]string{[]string{"X", "X", "X"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
  for i := 0; i < b.N; i++ {
    All(board, "X")
  }
}

func BenchmarkMaxBy(b *testing.B) {
  arr := [][]int{[]int{0, 99}, []int{1, 100}, []int{2, 98}, []int{3, 93}, []int{4,101}, []int{5, -101}, []int{6, 45}, []int{7, 88}, []int{8, -22}}
  for i := 0; i < b.N; i++ {
    MaxBy(arr, 0)
  }
} 

func BenchmarkSliceRows(b *testing.B) {
  board := []string{"X", "-", "-", "X", "O", "O", "-", "-", "-"}
  for i := 0; i < b.N; i++ {
    SliceRows(board)
  }
}

func BenchmarkMatchSlice(b *testing.B) {
  win_x := []string{"X", "X", "X", "O", "-", "O", "-", "-", "-"}
  indexes := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8}}
  for i:= 0; i < b.N; i++ {
    MatchSlice(win_x, indexes, "X")
  }
}

func BenchmarkReverseSlice(b *testing.B) {
  board := [][]string{[]string{"X", "O", "X"}, []string{"X", "O", "O"}, []string{"X", "O", "X"}}
  for i:= 0; i < b.N; i++ {
    ReverseSlice(board)
  }
}

func BenchmarkAddOneToSliceValues(b *testing.B) {
  board := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  for i:= 0; i < b.N; i++ {
    AddOneToSliceValues(board)
  }
}
