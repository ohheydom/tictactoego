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
	n := len(board)
	dimension := int(math.Sqrt(float64(n)))
	sliced_board := make([][]string, 0, dimension)
	for i := 0; i < n; {
		sliced_board = append(sliced_board, board[i:i+dimension])
		i += dimension
	}
	return sliced_board
}

func ReverseSlice(sliced_board [][]string) [][]string {
	temp_slice := make([][]string, len(sliced_board))
	for i, idx := len(sliced_board)-1, 0; i >= 0; i-- {
		temp_slice[idx] = sliced_board[i]
		idx += 1
	}
	return temp_slice
}

func Transpose(sliced_board [][]string) [][]string {
	len_a := len(sliced_board)
	len_b := len(sliced_board[0])
	if len_a == len_b {
		return TransposeSquare(sliced_board)
	}
	temp_slice := make([][]string, len_b)
	for i := 0; i < len_b; i++ {
		for ib := 0; ib < len_a; ib++ {
			temp_slice[i] = append(temp_slice[i], sliced_board[ib][i])
		}
	}
	return temp_slice
}

func TransposeSquare(sliced_board [][]string) [][]string {
	len_a := len(sliced_board)
	for n := 0; n < len_a-1; n++ {
		for m := n + 1; m < len_a; m++ {
			sliced_board[n][m], sliced_board[m][n] = sliced_board[m][n], sliced_board[n][m]
		}
	}
	return sliced_board
}

func TransposeDiagonal(sliced_board [][]string) [][]string {
	dimension := len(sliced_board)
	temp_slice := make([]string, dimension)
	temp_slice_2 := make([]string, dimension)
	for i := 0; i < dimension; i++ {
		temp_slice[i] = sliced_board[i][i]
		temp_slice_2[i] = sliced_board[i][dimension-i-1]
	}
	return [][]string{temp_slice, temp_slice_2}
}

func All(board [][]string, value string) bool {
	for _, row := range board {
		count := 0
		for _, mark := range row {
			if mark == value {
				count += 1
			} else {
				break
			}
			if count == len(row) {
				return true
			}
		}
	}
	return false
}

func AddOneToSliceValues(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] + 1
	}
	return slice
}

func MatchSlice(arr []string, slice [][]int, turn string) bool {
	n := len(slice)
	nn := len(slice[0])
	for i := 0; i < n; i++ {
		for ib, count := 1, 0; ib < nn; ib++ {
			if arr[slice[i][ib]] != turn {
				break
			}
			if arr[slice[i][ib]] == arr[slice[i][ib-1]] {
				count++
			}
			if count == nn-1 {
				return true
			}
		}
	}
	return false
}
