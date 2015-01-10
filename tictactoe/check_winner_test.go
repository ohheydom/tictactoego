package tictactoe

import "testing"

func TestHorizontalWin(t *testing.T) {
	win_x := []string{"X", "X", "X", "O", "-", "O", "-", "-", "-"}
	win_x_middle := []string{"-", "-", "O", "X", "X", "X", "O", "-", "-"}
	win_x_end := []string{"-", "-", "-", "O", "-", "O", "X", "X", "X"}
	win_o := []string{"O", "O", "O", "X", "X", "-", "-", "-", "-"}
	win_o_middle := []string{"-", "-", "X", "O", "O", "O", "X", "-", "-"}
	win_o_end := []string{"-", "-", "X", "X", "-", "-", "O", "O", "O"}
	winners_x := [][]string{win_x, win_x_middle, win_x_end}
	winners_o := [][]string{win_o, win_o_middle, win_o_end}
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
	win_x := []string{"X", "O", "-", "X", "O", "O", "X", "-", "-"}
	win_x_middle := []string{"-", "X", "-", "O", "X", "O", "-", "X", "-"}
	win_o_middle := []string{"-", "O", "X", "-", "O", "-", "X", "O", "-"}
	win_o_end := []string{"-", "-", "O", "X", "X", "O", "X", "-", "O"}
	winners_x := [][]string{win_x, win_x_middle}
	winners_o := [][]string{win_o_middle, win_o_end}
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
	win_x := []string{"X", "O", "O", "-", "X", "-", "O", "X", "X"}
	if DiagonalWin(win_x, "X") == false {
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
