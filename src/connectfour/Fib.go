package main


func fib(n int) int {
  if n <= 1 {
    return 1
  }
  return fib(n -1) + fib(n - 2)
}

func fibGo(n int, ch chan int) {
  if n <= 1 {
    ch <- 1
    return
  }
  ch2, ch3 := make(chan int), make(chan int)
  go fibGo(n - 1, ch2)
  go fibGo(n - 2, ch3)
  ch<- <-ch2 + <-ch3
}