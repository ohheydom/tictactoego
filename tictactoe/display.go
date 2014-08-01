package tictactoe

import (
  "fmt"
  "math"
  "strings"
)

func Play() {
  dim := GridSizeMessage()
  game := GameBoard{CreateBoard(dim), "X"}
  DisplayBoard(game.Board)
  DisplayTurn(&game)
  LoopThroughMoves(&game)
  DisplayWinner(game)
  DisplayPlayAgain()
}

func horizontal_bars(dimension int) {
  fmt.Println("--" + strings.Repeat("-", (dimension * 2)) + "-")
}

func DisplayWinner(g GameBoard) {
  if g.Win() == true {
    fmt.Printf("Congratulations %v!! You've won!!\n", g.PreviousTurn())
  } else {
    fmt.Println("It's a draw.")
  }
}
  
func DisplayPlayAgain() {
  var yes_or_no string
  fmt.Printf("Would you like to play again?? ")
  _, err := fmt.Scanf("%s", &yes_or_no)
  if err != nil {
  } else {
    if strings.ToUpper(yes_or_no) == "YES" || strings.ToUpper(yes_or_no) == "Y" {
      Play()
    }
  }
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

func DisplayTurn(g *GameBoard) {
  fmt.Println("Current Turn:", g.Turn)
}

func DisplayAskForMove() {
  fmt.Println("Make your move: ")
}

func DisplayRemainingMoves(g *GameBoard) {
  fmt.Println("Remaining Moves:", g.RemainingIndices())
}
