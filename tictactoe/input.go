package tictactoe

func (g GameBoard) Move(ind int, turn string) {
  if g.Board[ind] == "-" {
    g.Board[ind] = turn
  }
}
