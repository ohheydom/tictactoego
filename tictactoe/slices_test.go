package tictactoe

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	board := [][]string{[]string{"X", "X", "X"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
	expectedValue := true
	receivedValue := All(board, "X")
	if expectedValue != receivedValue {
		t.Error("Expected: ", expectedValue, "Received: ", receivedValue)
	}
}

func TestReverseSlice(t *testing.T) {
	initialBoard := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
	expectedBoard := [][]string{[]string{"-", "X", "O"}, []string{"-", "O", "O"}, []string{"X", "X", "-"}}
	receivedBoard := ReverseSlice(initialBoard)
	if !reflect.DeepEqual(expectedBoard, receivedBoard) {
		t.Error("Expected", expectedBoard, "got", receivedBoard)
	}
}

func TestTranspose(t *testing.T) {
	initialBoard := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
	expectedTransposedBoard := [][]string{[]string{"X", "-", "-"}, []string{"X", "O", "X"}, []string{"-", "O", "O"}}
	calculatedTransposedBoard := Transpose(initialBoard)
	if !reflect.DeepEqual(expectedTransposedBoard, calculatedTransposedBoard) {
		t.Error("Expected", expectedTransposedBoard, "got", calculatedTransposedBoard)
	}

	initialUnevenBoard := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}}
	expectedUnevenTran := [][]string{[]string{"X", "-"}, []string{"X", "O"}, []string{"-", "O"}}
	calculatedUnevenTran := Transpose(initialUnevenBoard)
	if !reflect.DeepEqual(expectedUnevenTran, calculatedUnevenTran) {
		t.Error("Expected", expectedUnevenTran, "got", calculatedUnevenTran)
	}
}

func TestTransposeSquare(t *testing.T) {
	initialBoard := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
	expectedTransposedBoard := [][]string{[]string{"X", "-", "-"}, []string{"X", "O", "X"}, []string{"-", "O", "O"}}
	calculatedTransposedBoard := TransposeSquare(initialBoard)
	if !reflect.DeepEqual(expectedTransposedBoard, calculatedTransposedBoard) {
		t.Error("Expected", expectedTransposedBoard, "got", calculatedTransposedBoard)
	}
}

func TestTransposeDiagonal(t *testing.T) {
	initialBoard := [][]string{[]string{"X", "X", "-"}, []string{"-", "O", "O"}, []string{"-", "X", "O"}}
	transposedBoard := [][]string{[]string{"X", "O", "O"}, []string{"-", "O", "-"}}
	tran := TransposeDiagonal(initialBoard)
	if !reflect.DeepEqual(tran, transposedBoard) {
		t.Error("Expected", transposedBoard, "got", tran)
	}
}

func TestMatchSlice(t *testing.T) {
	winX := []string{"X", "X", "X", "O", "-", "O", "-", "-", "-"}
	indexes := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8}}
	expectedValue := true
	receivedValue := MatchSlice(winX, indexes, "X")
	if expectedValue != receivedValue {
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
	arr := [][]int{[]int{0, 99}, []int{1, 100}, []int{2, 98}, []int{3, 93}, []int{4, 101}, []int{5, -101}, []int{6, 45}, []int{7, 88}, []int{8, -22}}
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
	winX := []string{"X", "X", "X", "O", "-", "O", "-", "-", "-"}
	indexes := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8}}
	for i := 0; i < b.N; i++ {
		MatchSlice(winX, indexes, "X")
	}
}

func BenchmarkReverseSlice(b *testing.B) {
	board := [][]string{[]string{"X", "O", "X"}, []string{"X", "O", "O"}, []string{"X", "O", "X"}}
	for i := 0; i < b.N; i++ {
		ReverseSlice(board)
	}
}

func BenchmarkTranspose(b *testing.B) {
	board := [][]string{[]string{"X", "O", "X"}, []string{"X", "O", "O"}, []string{"X", "O", "X"}}
	for i := 0; i < b.N; i++ {
		Transpose(board)
	}
}

func BenchmarkTransposeDiagonal(b *testing.B) {
	board := [][]string{[]string{"X", "O", "X"}, []string{"X", "O", "O"}, []string{"X", "O", "X"}}
	for i := 0; i < b.N; i++ {
		TransposeDiagonal(board)
	}
}
func BenchmarkAddOneToSliceValues(b *testing.B) {
	board := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		AddOneToSliceValues(board)
	}
}
