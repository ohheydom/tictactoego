package main

import (
  . "./tictactoe/"
)

func main() {
  dim := GridSizeMessage()
  game := GameBoard{CreateBoard(dim), "X"}
  DisplayBoard(game.Board)
  game.Move(4, "X")
  DisplayBoard(game.Board)
}
