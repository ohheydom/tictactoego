package tictactoe

func (g GameBoard) Move(ind int, turn string) {
  g.Board[ind] = turn
}
