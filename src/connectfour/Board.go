package main

type Board struct {
  fields []string
}

func (b *Board) set(p Point, val string) {
  b.fields[p.toIndex()] = val
}

func (b *Board) get(p Point) (string, bool) {
  if (p.x >= SIZE || p.y >= SIZE) {
    return ".", false
  }
  if (p.x < 0 || p.y < 0) {
    return ".", false
  }
  return b.fields[p.toIndex()], true
}

func (b *Board) mockUp() []string {
  return []string{
    "_", "X", "_",
    "_", "X", "O",
    "X", "O", "O",
  }
}

func (b *Board) init() {
  b.fields = make([]string, SIZE * SIZE)
  for i := 0; i < SIZE * SIZE; i++ {
    b.fields[i] = EMPTYFIELD
  }
  //b.fields = b.mockUp()
}

func (b *Board) getHeight(col int) int {
  for i := SIZE - 1; i >= 0; i-- {
    if v, ok := b.get(Point{col, i}); v == EMPTYFIELD && ok {
      return i
    }
  }
  return -1
}

func (b *Board) getField(col int) (Point, bool) {
  h := b.getHeight(col)
  if h < 0 {
    return Point{}, false
  }
  return Point{col, h}, true
}

func (b *Board) isValidMove(i int) bool {
  if i < 0 || i >= SIZE {
    return false
  }
  p, ok := b.getField(i)
  if !ok {
    return false
  }
  return b.fields[p.toIndex()] == EMPTYFIELD
}