package tictactoe

func CreateBoard(size int) []string {
  dim := size * size
  var board []string
  for i := 0; i < dim; i++ {
    board = append(board, "-")
  }
  return board
}
