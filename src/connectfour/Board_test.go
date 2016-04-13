package main

import (
  "time"
  "log"
  "testing"
)
/*

import "testing"

func BenchmarkWinCondition(b *testing.B) {
  g := Game{}
  g.init()
  g.board.mockUp()

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    g.checkWinCondition(0)
  }
}

func BenchmarkWinConditionGo(b *testing.B) {
  g := Game{}
  g.init()
  g.board.mockUp()

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    g.checkWinConditionGo(0)
  }
}
func BenchmarkWinConditionGo2(b *testing.B) {
  g := Game{}
  g.init()
  g.board.mockUp()

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    g.checkWinConditionGo2(0)
  }
}*/

const (
  TIMES = 10
)

func TestPerformance(test *testing.T) {
  g := Game{}
  g.init()
  //g.board.mockUp()


  /*tests := []func(int, chan<- int) int {
    fib,
    fibGo,
  }*/
  s, t := trace("Call Func")
  for k := 0; k < TIMES; k++ {
    fib(10)
  }
  un(s, t)

  ch := make(chan int)
  s1, t1 := trace("Call Func")
  for k := 0; k < TIMES; k++ {
    go fibGo(10, ch)
  }
  <-ch
  un(s1, t1)

}

func trace(s string) (string, time.Time) {
  log.Println("START:", s)
  return s, time.Now()
}

func un(s string, startTime time.Time) {
  endTime := time.Now()
  log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime) / TIMES)
}

