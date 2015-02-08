package tictactoe

import (
	"reflect"
	"testing"
)

func TestRemainingIndices(t *testing.T) {
	board := []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
	g := GameBoard{Board: board, Turn: "X"}
	expectedResponse := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	actualResponse := g.RemainingIndices()
	if !reflect.DeepEqual(expectedResponse, actualResponse) {
		t.Error("Expected: ", expectedResponse, "Received: ", actualResponse)
	}
}

func BenchmarkRemainingIndices(b *testing.B) {
	board := []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
	g := GameBoard{Board: board, Turn: "O"}
	for i := 0; i < b.N; i++ {
		g.RemainingIndices()
	}
}

func BenchmarkValidMove(b *testing.B) {
	remainingIndices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		ValidMove(remainingIndices, 7)
	}
}
