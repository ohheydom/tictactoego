package tictactoe

func DiagonalWin(board []string, turn string) bool {
  sliced_board := TransposeDiagonal(SliceRows(board))
  return All(sliced_board, turn)
}

func HorizontalWin(board []string, turn string) bool {
  sliced_board := SliceRows(board)
  return All(sliced_board, turn)
}

func VerticalWin(board []string, turn string) bool {
  sliced_board := Transpose(SliceRows(board))
  return All(sliced_board, turn)
}

func (g GameBoard) Win() bool {
  if VerticalWin(g.Board, g.PreviousTurn()) || HorizontalWin(g.Board, g.PreviousTurn()) || DiagonalWin(g.Board, g.PreviousTurn()) {
    return true
  }
  return false
}
