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
  g.players[0] = &Bot{0}
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
    fmt.Println(PLAYER[g.currPlayer], "has won the game!")
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
  //return g.checkWinConditionGo(currPlayer)
  //---------------------------------
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

//-------------------

func (g *Game) checkWinConditionGo(currPlayer int) bool {
  ch := make(chan bool, len(g.board.fields))
  for i := range g.board.fields {
    p := Point{}
    p.fromIndex(i)
    go g.checkWinDirectionsGo(p, currPlayer, ch)
  }

  for i := 0; i < len(g.board.fields); i++ {
    if <-ch {
      return true
    }
  }

  return false
}

func (g *Game) checkWinDirectionsGo(p Point, currPlayer int, ch chan bool) {
  ch2 := make(chan bool, 4)

  for i := 0; i < 4; i++ {
    go g.checkWinDirectionGo(p, i, 0, currPlayer, ch2)
  }

  for i := 0; i < 4; i++ {
    if <-ch2 {
      ch <- true
      return
    }
  }
  ch <- false

}

func (g *Game) checkWinDirectionGo(p Point, direction, depth, currPlayer int, ch chan bool) {
  player := PLAYER[currPlayer]

  if v, ok := g.board.get(p); v != player || !ok {
    ch <- false
    return
  }

  if depth >= WINCONDITION - 1 {
    ch <- true
    return
  }

  switch direction {
  case RIGHT:
    p.x++
    g.checkWinDirectionGo(p, RIGHT, depth + 1, currPlayer, ch)
    return
    break
  case DOWN:
    p.y++
    g.checkWinDirectionGo(p, DOWN, depth + 1, currPlayer, ch)
    return
    break
  case DIAG_UP:
    p.x++
    p.y++
    g.checkWinDirectionGo(p, DIAG_UP, depth + 1, currPlayer, ch)
    return
    break
  case DIAG_DOWN:
    p.x++
    p.y--
    g.checkWinDirectionGo(p, DIAG_DOWN, depth + 1, currPlayer, ch)
    return
    break
  }
  ch <- false
  return
}

func (g *Game) checkWinConditionGo2(currPlayer int) bool {
  ch := make(chan bool, len(g.board.fields))
  for i := range g.board.fields {
    p := Point{}
    p.fromIndex(i)
    go g.checkWinDirectionsGo2(p, currPlayer, ch)
  }

  for i := 0; i < len(g.board.fields); i++ {
    if <-ch {
      return true
    }
  }

  return false
}

func (g *Game) checkWinDirectionsGo2(p Point, currPlayer int, ch chan bool) {

  ch2 := make(chan bool, 4)

  for i := 0; i < 4; i++ {
    go g.checkWinDirectionGo2(p, i, 0, currPlayer, ch2)
  }

  for i := 0; i < 4; i++ {
    if <-ch2 {
      ch <- true
      return
    }
  }
  ch <- false

}

func (g *Game) checkWinDirectionGo2(p Point, direction, depth, currPlayer int, ch chan bool) {
  player := PLAYER[currPlayer]

  if v, ok := g.board.get(p); v != player || !ok {
    ch <- false
    return
  }

  if depth >= WINCONDITION - 1 {
    ch <- true
    return
  }

  switch direction {
  case RIGHT:
    p.x++
    go g.checkWinDirectionGo2(p, RIGHT, depth + 1, currPlayer, ch)
    return
    break
  case DOWN:
    p.y++
    go g.checkWinDirectionGo2(p, DOWN, depth + 1, currPlayer, ch)
    return
    break
  case DIAG_UP:
    p.x++
    p.y++
    go g.checkWinDirectionGo2(p, DIAG_UP, depth + 1, currPlayer, ch)
    return
    break
  case DIAG_DOWN:
    p.x++
    p.y--
    go g.checkWinDirectionGo2(p, DIAG_DOWN, depth + 1, currPlayer, ch)
    return
    break
  }
  ch <- false
  return
}
