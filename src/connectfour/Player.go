package main

import (
  "bufio"
  "os"
  "fmt"
  "strconv"
  "strings"
)

type Player interface {
  readInput(board Board) int
}

type Bot struct {
  id int
}

func (b *Bot) readInput(board Board) int {
  res := b.doTurn(board)
  fmt.Printf("Player %s set: %d\n", PLAYER[b.id], res + 1)
  return res
}

func (b *Bot) minimaxAlgorithm(field []string, depth int, maximizing bool, alpha, beta int) int {
  board := Board{}
  game := Game{}
  board.fields = make([]string, len(field))
  copy(board.fields, field)

  game.board = board

  if game.checkWinCondition(b.id) {
    return 100 - depth
  }
  if game.checkWinCondition(1 ^ b.id) {
    return -100 + depth
  }
  if game.isTied() {
    return 0
  }
  if depth >= 10 {
    return 0
  }

  score := make(map[int]int)
  var v int

  if maximizing {
    v = -9000
    game.currPlayer = b.id
  } else {
    v = 9000
    game.currPlayer = 1 ^ b.id
  }

  var bestMove int
  for i := 0; i < SIZE; i++ {
    board.fields = make([]string, len(field))
    copy(board.fields, field)
    game.board = board

    if !board.isValidMove(i) {
      continue
    }

    game.setField(i)

    score[i] = b.minimaxAlgorithm(board.fields, depth + 1, !maximizing, alpha, beta)

    if maximizing {
      if score[i] > v {
        v = score[i]
        bestMove = i
      }
      alpha = max(alpha, v)
      if beta <= alpha {
        break
      }
    } else {
      if score[i] < v {
        v = score[i]
        bestMove = i
      }
      beta = min(beta, v)
      if beta <= alpha {
        break
      }
    }

  }

  if (depth == 0) {
    return bestMove
  }
  return v
}

func (b *Bot) doTurn(board Board) int {
  return b.minimaxAlgorithm(board.fields, 0, true, -9000, 9000)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

type Human struct {
  id int
}

func (h *Human) readInput(board Board) int {
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("Player %s set: ", PLAYER[h.id])
  input, _ := reader.ReadString('\n')
  i, ok := strconv.ParseInt(strings.Trim(input, "\n\r"), 10, 16)

  if ok != nil {
    fmt.Println(ok)
    return h.readInput(board)
  }

  i--
  if i < 0 || i >= SIZE {
    fmt.Println("range error")
    return h.readInput(board)
  }
  return int(i)
}