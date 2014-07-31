package tictactoe

import (
  "fmt"
  "math"
  "strings"
)

func horizontal_bars(dimension int) {
  fmt.Println("--" + strings.Repeat("-", (dimension * 2)) + "-")
}

func DisplayBoard(board []string) {
  dimension := int(math.Sqrt(float64(len(board))))
  horizontal_bars(dimension)
  for i := 0; i < len(board); {
    row_end := i + dimension
    fmt.Println("|", strings.Join(board[i:row_end], " "), "|")
    i += dimension
  }
  horizontal_bars(dimension)
}

func DisplayTurn(g GameBoard) {
  fmt.Println("Current Turn:", g.Turn)
}

func DisplayAskForMove() {
  fmt.Println("Make your move: ")
}

func DisplayRemainingMoves(g GameBoard) {
  fmt.Println("Remaining Moves:", g.RemainingIndices())
}
