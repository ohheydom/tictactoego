package tictactoe

import "testing"

func TestHorizontalWin(t *testing.T) {
	winX := []string{"X", "X", "X", "O", "-", "O", "-", "-", "-"}
	winXMiddle := []string{"-", "-", "O", "X", "X", "X", "O", "-", "-"}
	winXEnd := []string{"-", "-", "-", "O", "-", "O", "X", "X", "X"}
	winO := []string{"O", "O", "O", "X", "X", "-", "-", "-", "-"}
	winOMiddle := []string{"-", "-", "X", "O", "O", "O", "X", "-", "-"}
	winOEnd := []string{"-", "-", "X", "X", "-", "-", "O", "O", "O"}
	winnersX := [][]string{winX, winXMiddle, winXEnd}
	winnersO := [][]string{winO, winOMiddle, winOEnd}
	for _, result := range winnersX {
		if HorizontalWin(SliceRows(result), "X") == false {
			t.Error("Expected true")
		}
	}

	for _, result := range winnersO {
		if HorizontalWin(SliceRows(result), "O") == false {
			t.Error("Expected true")
		}
	}
}

func TestVerticalWin(t *testing.T) {
	winX := []string{"X", "O", "-", "X", "O", "O", "X", "-", "-"}
	winXMiddle := []string{"-", "X", "-", "O", "X", "O", "-", "X", "-"}
	winOMiddle := []string{"-", "O", "X", "-", "O", "-", "X", "O", "-"}
	winOEnd := []string{"-", "-", "O", "X", "X", "O", "X", "-", "O"}
	winnersX := [][]string{winX, winXMiddle}
	winnersO := [][]string{winOMiddle, winOEnd}
	for _, result := range winnersX {
		if VerticalWin(SliceRows(result), "X") == false {
			t.Error("Expected true")
		}
	}

	for _, result := range winnersO {
		if VerticalWin(SliceRows(result), "O") == false {
			t.Error("Expected true")
		}
	}
}

func TestDiagonalWin(t *testing.T) {
	winX := []string{"X", "O", "O", "-", "X", "-", "O", "X", "X"}
	if DiagonalWin(SliceRows(winX), "X") == false {
		t.Error("Expected True")
	}
}

func BenchmarkWin(b *testing.B) {
	board := []string{"-", "-", "-", "X", "-", "-", "-", "-", "-"}
	g := GameBoard{Board: board, Turn: "O"}
	for i := 0; i < b.N; i++ {
		g.Win()
	}
}
