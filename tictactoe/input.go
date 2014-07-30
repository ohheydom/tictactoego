package tictactoe

import "fmt"

func (g GameBoard) Move(ind int, turn string) {
  if g.Board[ind] == "-" {
    g.Board[ind] = turn
  }
}

func GridSizeMessage()(dim int) {
  fmt.Printf("What size grid would you like to play on?? Please enter only one dimension (ie 3, 4, or 5): ")
  _, err := fmt.Scanf("%d", &dim)
  if err != nil {
  }
  return
}
