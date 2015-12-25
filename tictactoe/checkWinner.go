package tictactoe

func DiagonalWin(slicedBoard [][]string, turn string) bool {
	diagBoard := TransposeDiagonal(slicedBoard)
	return All(diagBoard, turn)
}

func HorizontalWin(slicedBoard [][]string, turn string) bool {
	return All(slicedBoard, turn)
}

func VerticalWin(slicedBoard [][]string, turn string) bool {
	for i, n := 0, len(slicedBoard); i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if turn == slicedBoard[j][i] {
				count++
			} else {
				break
			}
		}
		if count == n {
			return true
		}
	}
	return false
}

func (g GameBoard) Win() bool {
	if g.Count < 5 {
		return false
	}
	slicedBoard := SliceRows(g.Board)
	if VerticalWin(slicedBoard, g.PreviousTurn()) || HorizontalWin(slicedBoard, g.PreviousTurn()) || DiagonalWin(slicedBoard, g.PreviousTurn()) {
		return true
	}
	return false
}
