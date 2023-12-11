package index

import (
  "math/rand"
  "fmt"
  "io"
  "time"
)

const MAX_KEY int = 10e2

type STItem struct {
  key uint
  value float32
}

func NewSTItem() *STItem {
  return &STItem{key: uint(MAX_KEY)}
}

func (i *STItem) Key() uint {
  return i.key
}

func (i *STItem) IsNull() bool {
  return int(i.key) == MAX_KEY
}

func (i *STItem) Rand() {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  i.key = uint(r.Uint32()) % uint(MAX_KEY)
  i.value = r.Float32()
}

func (i *STItem) Print(w io.Writer) {
  w.Write([]byte(fmt.Sprintf("%d: %g\n", i.key, i.value)))
}

func (i *STItem) Value() float32 {
  return i.value
}

func (i *STItem) SetValue(v float32) {
  i.value = v
}

func (i *STItem) SetKey(k uint) {
  i.key = k
}

type SearchTable struct {
  st [](*STItem)
}

func New() *SearchTable {
  st := make([](*STItem), MAX_KEY)
  return &SearchTable{st}
}

func (s *SearchTable) Search(key uint) *STItem {
  return s.st[key]
}

func (s *SearchTable) Insert(i *STItem) {
  s.st[i.Key()] = i
}

func (s *SearchTable) Remove(i *STItem) {
  s.st[i.Key()] = nil
}

func (s *SearchTable) Select(k uint) *STItem {
  for i := 0; i < len(s.st); i++ {
    if s.st[i] != nil && !s.st[i].IsNull() {
      if k == 0 { return s.st[i]; }
      k -= 1
    }
  }
  return NewSTItem()
}

func (s *SearchTable) Count() int {
  var result int
  for i := 0; i < len(s.st); i++ {
    if s.st[i] != nil && !s.st[i].IsNull() {
      result += 1
    }
  }
  return result
}

func (s *SearchTable) Print(w io.Writer) {
  for i := 0; i < len(s.st); i++ {
    if s.st[i] != nil && !s.st[i].IsNull() {
      s.st[i].Print(w)
    }
  }
}
