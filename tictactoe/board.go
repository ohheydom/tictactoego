package tictactoe

type GameBoard struct {
  Board []string
  Turn string
}

func CreateBoard(size int) []string {
  dim := size * size
  var board []string
  for i := 0; i < dim; i++ {
    board = append(board, "-")
  }
  return board
}

func (g *GameBoard) Move(ind int, turn string) {
  if g.Board[ind] == "-" {
    g.Board[ind] = turn
    g.SwitchTurn()
  }
}

func (g GameBoard) PreviousTurn() string {
  if g.Turn == "X" {
    return "O"
  }
  return "X"
}

func (g GameBoard) RemainingIndices() (remaining_indices []int) {
  for i, value := range g.Board {
    if value == "-" {
      remaining_indices = append(remaining_indices, i)
    }
  }
  return
}

func (g *GameBoard) SwitchTurn() {
  if g.Turn == "X" {
    g.Turn = "O"
  } else {
    g.Turn = "X" 
  }
}

func (g *GameBoard) UndoMove(ind int) {
  g.Board[ind] = "-"
  g.SwitchTurn()
}

func ValidMove(remaining_indices []int, move int) bool {
  for _, valid_move := range remaining_indices {
    if valid_move == move {
      return true
    }
  }
  return false
}
