package tictactoe

func DiagonalWin(board []string, turn string) bool {
	slicedBoard := TransposeDiagonal(SliceRows(board))
	return All(slicedBoard, turn)
}

func HorizontalWin(board []string, turn string) bool {
	slicedBoard := SliceRows(board)
	return All(slicedBoard, turn)
}

func VerticalWin(board []string, turn string) bool {
	slicedBoard := Transpose(SliceRows(board))
	x := All(slicedBoard, turn)
	Transpose(slicedBoard)
	return x
}

func (g GameBoard) Win() bool {
	if g.Count < 5 {
		return false
	}
	if VerticalWin(g.Board, g.PreviousTurn()) || HorizontalWin(g.Board, g.PreviousTurn()) || DiagonalWin(g.Board, g.PreviousTurn()) {
		return true
	}
	return false
}
