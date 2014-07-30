package tictactoe

import (
  "fmt"
  "math"
  "strings"
)

func DisplayBoard(board []string) {
  dimension := int(math.Sqrt(float64(len(board))))
  fmt.Println("---" + strings.Repeat("-", (dimension * 2)) + "--")
  for i := 0; i < dimension; i++ {
    top := i + dimension
    fmt.Println("| ", strings.Join(board[i:top], " "), " |")
  }
  fmt.Println("---" + strings.Repeat("-", (dimension * 2)) + "--")
}
