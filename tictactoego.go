package main

import (
  . "./tictactoe/"
)

func main() {
  game := GameBoard{ CreateBoard(5), "X" }
  DisplayBoard(game.Board)
}
