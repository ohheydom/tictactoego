package tictactoe

type GameBoard struct {
  Board []string
}

func CreateBoard(size int) []string {
  dim := size * size
  var board []string
  for i := 0; i < dim; i++ {
    board = append(board, "-")
  }
  return board
}

func SwitchTurn(turn *string) {
  if *turn == "X" {
    *turn = "O" } else {
    *turn = "X" }
}

func (g GameBoard) RemainingIndices() (remaining_indices []int) {
  for i, value := range g.Board {
    if value == "-" {
      remaining_indices = append(remaining_indices, i)
    }
  }
  return
}
