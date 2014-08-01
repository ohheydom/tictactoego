package tictactoe

func DiagonalWin(board []string, turn string) bool {
  sliced_board := TransposeDiagonal(SliceRows(board))
  are_equal := false
  for _, val := range sliced_board {
    count := 0
    for _, value := range val {
      if value == turn { 
        count += 1
      }
      if count == len(val) {
        are_equal = true
        break
      }
    }
  }
  return are_equal
}

func HorizontalWin(board []string, turn string) bool {
  sliced_board := SliceRows(board)
  are_equal := false
  for _, val := range sliced_board {
    count := 0
    for _, value := range val {
      if value == turn { 
        count += 1
      }
      if count == len(val) {
        are_equal = true
        break
      }
    }
  }
  return are_equal
}

func VerticalWin(board []string, turn string) bool {
  sliced_board := Transpose(SliceRows(board))
  are_equal := false
  for _, val := range sliced_board {
    count := 0
    for _, value := range val {
      if value == turn { 
        count += 1
      }
      if count == len(val) {
        are_equal = true
        break
      }
    }
  }
  return are_equal
}

func (g GameBoard) Win() bool {
  if VerticalWin(g.Board, g.PreviousTurn()) == true || HorizontalWin(g.Board, g.PreviousTurn()) == true || DiagonalWin(g.Board, g.PreviousTurn()) == true {
    return true
  }
  return false
}
