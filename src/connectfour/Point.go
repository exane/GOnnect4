package main


type Point struct {
  x, y int
}

func (p *Point) toIndex() int {
  return p.x + SIZE * p.y
}

func (p *Point) fromIndex(index int) *Point {
  p.x = index % SIZE
  p.y = (index - p.x) / SIZE
  return p
}

