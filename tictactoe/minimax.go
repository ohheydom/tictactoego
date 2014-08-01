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

  if max_player == true {
    best_value := -1000
    for _, move := range g.RemainingIndices() {
      g.Move(move, "X")
      value := MiniMax(g, depth + 1, false)
      g.UndoMove(move)
      best_value = max(best_value, value)
    }
    return best_value
  } else {
    best_value := 1000
    for _, move := range g.RemainingIndices() {
      g.Move(move, "O")
      value := MiniMax(g, depth + 1, true)
      g.UndoMove(move)
      best_value = min(best_value, value)
    }
    return best_value
  }
  return 0
}

func (g GameBoard) MiniMaxMoves() [][]int {
  var score [][]int
  for _, move := range g.RemainingIndices() {
    g.Move(move, g.Turn)
    arr := []int{MiniMax(g, 0, true), move}
    score = append(score, arr)
    g.UndoMove(move)
  }
  return score
}

func max_by_score(arr [][]int) []int {
  top := arr[0]
  for _, val := range arr {
    if top[0] < val[0] {
      top = val
    }
  }
  return top
}

func min_by_score(arr [][]int) []int {
  top := arr[0]
  for _, val := range arr {
    if top[0] > val[0] {
      top = val
    }
  }
  return top
}

func (g GameBoard) BestMove() int {
  if g.Turn == "X" {
    return max_by_score(g.MiniMaxMoves())[1]
  }
  return min_by_score(g.MiniMaxMoves())[1]
}
