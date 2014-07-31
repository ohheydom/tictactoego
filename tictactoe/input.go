package tictactoe

import "fmt"

func (g GameBoard) Move(ind int, turn string) {
  if g.Board[ind] == "-" {
    g.Board[ind] = turn
  }
}

func ValidMove(remaining_indices []int, move int) bool {
  for _, valid_move := range remaining_indices {
    if valid_move == move {
      return true
    }
  }
  return false
}

func GridSizeMessage()(dim int) {
  fmt.Printf("What size grid would you like to play on?? Please enter only one dimension (ie 3, 4, or 5): ")
  _, err := fmt.Scanf("%d", &dim)
  if err != nil {
  }
  return
}

func LoopThroughMoves(g GameBoard) {
  for g.Win() == false {
    var move int
    DisplayAskForMove()
    _, err := fmt.Scanf("%d", &move)
    if err != nil || ValidMove(g.RemainingIndices(), move) == false {
      fmt.Println("Please Enter A Valid Move")
    } else {
      g.Move(move, g.Turn)
      g.SwitchTurn()
    } 
  DisplayBoard(g.Board)
  DisplayRemainingMoves(g)
    if len(g.RemainingIndices()) == 0 {
      break
    }
  }
}
