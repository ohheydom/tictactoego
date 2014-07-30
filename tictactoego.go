package main

import (
  "./tictactoe/"
)

func main() {
  board := tictactoe.CreateBoard(4)
  tictactoe.DisplayBoard(board)
}
