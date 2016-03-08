package main

import (
  "fmt"
)

var PLAYER = []string{"X", "O"}

const (
  SIZE = 7
  WINCONDITION = 4
  EMPTYFIELD = "_"
)

const (
  RIGHT = iota
  DOWN
  DIAG_UP
  DIAG_DOWN
)

type Game struct {
  board      Board
  currPlayer int
  turn       int
  players    []Player
}

func (g *Game) init() {
  g.board.init()
  g.currPlayer = 1
  g.turn = 0
  g.players = make([]Player, 2)
  g.players[0] = &Human{0}
  g.players[1] = &Bot{1}
}

func (g *Game) setField(col int) bool {
  if h := g.board.getHeight(col); h > -1 {
    g.board.set(Point{col, h}, PLAYER[g.currPlayer])
    return true
  }
  return false
}

func (g *Game) start() {
  g.init()
  g.nextTurn()
}

func (g *Game) isTied() bool {
  for _, val := range g.board.fields {
    if val == EMPTYFIELD {
      return false
    }
  }
  return true
}

func (g *Game) nextTurn() {
  g.currPlayer ^= 1
  g.turn++

  for {
    col := g.players[g.currPlayer].readInput(g.board)
    ok := g.setField(col)
    if ok {
      break
    }
  }

  if g.checkWinCondition(g.currPlayer) {
    print(g.board)
    fmt.Println(PLAYER[g.currPlayer],"has won the game!")
    return
  }

  if g.isTied() {
    print(g.board)
    fmt.Println("Tied!")
    g.start()
    return
  }

  print(g.board)

  g.nextTurn()

}

func (g *Game) checkWinCondition(currPlayer int) bool {
  for i := range g.board.fields {
    p := Point{}
    p.fromIndex(i)
    if g.checkWinDirections(p, currPlayer) {
      return true
    }
  }

  return false
}

func (g *Game) checkWinDirections(p Point, currPlayer int) bool {
  return g.checkWinDirection(p, RIGHT, 0, currPlayer) ||
  g.checkWinDirection(p, DOWN, 0, currPlayer) ||
  g.checkWinDirection(p, DIAG_UP, 0, currPlayer) ||
  g.checkWinDirection(p, DIAG_DOWN, 0, currPlayer)
}

func (g *Game) checkWinDirection(p Point, direction, depth, currPlayer int) bool {
  player := PLAYER[currPlayer]

  if v, ok := g.board.get(p); v != player || !ok {
    return false
  }

  if depth >= WINCONDITION - 1 {
    return true
  }

  switch direction {
  case RIGHT:
    p.x++
    return g.checkWinDirection(p, RIGHT, depth + 1, currPlayer)
    break
  case DOWN:
    p.y++
    return g.checkWinDirection(p, DOWN, depth + 1, currPlayer)
    break
  case DIAG_UP:
    p.x++
    p.y++
    return g.checkWinDirection(p, DIAG_UP, depth + 1, currPlayer)
    break
  case DIAG_DOWN:
    p.x++
    p.y--
    return g.checkWinDirection(p, DIAG_DOWN, depth + 1, currPlayer)
    break
  }
  return false
}

