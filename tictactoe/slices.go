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

func SliceRows(board []string) (slicedBoard [][]string) {
	n := len(board)
	dimension := int(math.Sqrt(float64(n)))
	slicedBoard = make([][]string, 0, dimension)
	for i := 0; i < n; {
		slicedBoard = append(slicedBoard, board[i:i+dimension])
		i += dimension
	}
	return
}

func ReverseSlice(slicedBoard [][]string) [][]string {
	tempSlice := make([][]string, len(slicedBoard))
	for i, idx := len(slicedBoard)-1, 0; i >= 0; i-- {
		tempSlice[idx] = slicedBoard[i]
		idx++
	}
	return tempSlice
}

func Transpose(slicedBoard [][]string) [][]string {
	lenA := len(slicedBoard)
	lenB := len(slicedBoard[0])
	if lenA == lenB {
		return TransposeSquare(slicedBoard)
	}
	tempSlice := make([][]string, lenB)
	for i := 0; i < lenB; i++ {
		for ib := 0; ib < lenA; ib++ {
			tempSlice[i] = append(tempSlice[i], slicedBoard[ib][i])
		}
	}
	return tempSlice
}

func TransposeSquare(slicedBoard [][]string) [][]string {
	n := len(slicedBoard)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			slicedBoard[i][j], slicedBoard[j][i] = slicedBoard[j][i], slicedBoard[i][j]
		}
	}
	return slicedBoard
}

func TransposeDiagonal(slicedBoard [][]string) [][]string {
	n := len(slicedBoard)
	t, t2 := make([]string, n), make([]string, n)
	for i := 0; i < n; i++ {
		t[i] = slicedBoard[i][i]
		t2[i] = slicedBoard[i][n-i-1]
	}
	return [][]string{t, t2}
}

func All(board [][]string, value string) bool {
	for _, row := range board {
		count := 0
		for _, mark := range row {
			if mark == value {
				count++
			} else {
				break
			}
		}
		if count == len(row) {
			return true
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
