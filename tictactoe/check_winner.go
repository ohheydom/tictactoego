package tictactoe

func DiagonalWin(board []string, turn string) bool {
  indexes := [][]int{[]int{0, 4, 8}, []int{2, 4, 6}}
  return MatchSlice(board, indexes, turn)
  // sliced_board := TransposeDiagonal(SliceRows(board))
  // return All(sliced_board, turn)
}

func HorizontalWin(board []string, turn string) bool {
  indexes := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8}}
  return MatchSlice(board, indexes, turn)
  // sliced_board := SliceRows(board)
  // return All(sliced_board, turn)
}

func VerticalWin(board []string, turn string) bool {
  indexes := [][]int{[]int{0, 3, 6}, []int{1, 4, 7}, []int{2, 5, 8}}
  return MatchSlice(board, indexes, turn)
  // sliced_board := Transpose(SliceRows(board))
  // return All(sliced_board, turn)
}

func (g GameBoard) Win() bool {
  if VerticalWin(g.Board, g.PreviousTurn()) || HorizontalWin(g.Board, g.PreviousTurn()) || DiagonalWin(g.Board, g.PreviousTurn()) {
    return true
  }
  return false
}
