package tictactoe

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func MiniMax(g GameBoard, depth int, max_player bool) int {
  if g.Win() {
    if g.PreviousTurn() == "X" {
      return 100 - depth
    }
    return depth - 100
  }
  if len(g.RemainingIndices()) == 0 {
    return 0
  }

  best_value := -1000

  if max_player {
    best_value = -1000
    for _, move := range g.RemainingIndices() {
      g.Move(move, "X")
      value := MiniMax(g, depth - 1, false)
      best_value = max(best_value, value)
      return best_value
    }  
  } else {
    best_value = 1000
    for _, move := range g.RemainingIndices() {
      g.Move(move, "O")
      value := MiniMax(g, depth -1, true)
      best_value = min(best_value, value)
      return best_value
    }
  }
  return 0
}
