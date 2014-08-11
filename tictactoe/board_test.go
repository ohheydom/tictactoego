package tictactoe

import (
  "testing"
  "reflect"
)

func TestRemainingIndices(t *testing.T) {
  board := []string{"-", "-", "-","-", "-", "-", "-", "-", "-"}
  g := GameBoard{Board: board, Turn: "X"}
  expected_response := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  actual_response := g.RemainingIndices()
  if !reflect.DeepEqual(expected_response, actual_response) {
    t.Error("Expected: ", expected_response, "Received: ", actual_response)
  }
}

func BenchmarkRemainingIndices(b *testing.B) {
  board := []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
  g := GameBoard{Board: board, Turn: "O"}
  for i := 0; i < b.N; i++ {
    g.RemainingIndices()
  }
}
