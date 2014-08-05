package tictactoe

import "math"

func MaxBy(arr [][]int, index int) []int {
  top := arr[0]
  for _, val := range arr {
    if top[index] < val[index] {
      top = val
    }
  }
  return top
}

func MinBy(arr [][]int, index int) []int {
  top := arr[0]
  for _, val := range arr {
    if top[index] > val[index] {
      top = val
    }
  }
  return top
}

func SliceRows(board []string) [][]string {
  dimension := int(math.Sqrt(float64(len(board))))
  var sliced_board [][]string
  for i := 0; i < len(board); {
    row_end := i + dimension
    sliced_board = append(sliced_board, board[i:row_end])
    i += dimension
  }
  return sliced_board
}

func ReverseSlice(sliced_board [][]string) [][]string {
  var temp_slice [][]string
  for i := len(sliced_board) - 1; i >= 0; i-- {
    temp_slice = append(temp_slice, sliced_board[i])
  }
  return temp_slice
}

func Transpose(sliced_board [][]string) [][]string {
  len_a := len(sliced_board) //2
  len_b := len(sliced_board[0]) //3
  temp_slice := make([][]string, len_b)
  for i := 0; i < len_b; i++ {
    for ib := 0; ib < len_a; ib++ {
      temp_slice[i] = append(temp_slice[i], sliced_board[ib][i])
    }  
  }
  return temp_slice
}

func TransposeDiagonal(sliced_board [][]string) [][]string {
  var temp_slice []string
  dimension := len(sliced_board)
  reverse_sliced_board := ReverseSlice(sliced_board)
  for row_ind, row := range sliced_board {
    for val_ind := range row {
      if val_ind == row_ind {
        temp_slice = append(temp_slice, sliced_board[val_ind][row_ind])
      }
    }
  }

  for row_ind, row := range reverse_sliced_board {
    for val_ind := range row {
      if val_ind == row_ind {
        temp_slice = append(temp_slice, reverse_sliced_board[val_ind][row_ind])
      }
    }
  }
  return [][]string{temp_slice[0:dimension], temp_slice[dimension:dimension * 2]}
}

func All(board [][]string, value string) bool {
  for _, row := range board {
    count := 0
    for _, mark := range row {
      if mark == value { 
        count += 1
      } else { break }
      if count == len(row) {
        return true
      }
    }
  }
  return false
}

func AddOneToSliceValues(slice []int) []int {
  var new_slice []int
  for _, val := range slice {
    new_slice = append(new_slice, val + 1)
  }
  return new_slice
}
