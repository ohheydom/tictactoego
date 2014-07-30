package tictactoe

import "math"

func slice_rows(board []string) [][]string {
  dimension := int(math.Sqrt(float64(len(board))))
  var sliced_board [][]string
  for i := 0; i < len(board); {
    row_end := i + dimension
    sliced_board = append(sliced_board, board[i:row_end])
    i += dimension
  }
  return sliced_board
}

func transpose(sliced_board [][]string) [][]string{
  var temp_slice []string
  for row_ind, row := range sliced_board {
    for val_ind := range row {
      temp_slice = append(temp_slice, sliced_board[val_ind][row_ind])
    }
  }
  return slice_rows(temp_slice)
}

func HorizontalWin(board []string, turn string) bool {
  sliced_board := slice_rows(board)
  are_equal := false
  for _, val := range sliced_board {
    test := 0
    for _, value := range val {
      if value == turn { 
        test +=1
      }
      if test == len(val) {
        are_equal = true
        break
      }
    }
  }
  return are_equal
}

func VerticalWin(board []string, turn string) bool {
  sliced_board := transpose(slice_rows(board))
  are_equal := false
  for _, val := range sliced_board {
    if val[0] == val[1] && val[1] == val[2] {
      are_equal = true
      break
    }
  }
  return are_equal
}
