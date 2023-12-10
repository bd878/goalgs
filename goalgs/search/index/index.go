package index

import (
  "math/rand"
  "fmt"
  "io"
  "time"
)

const maxKey uint = 10e2

type IndexSTItem struct {
  key uint
  value float32
}

func NewItem() *IndexSTItem {
  return &IndexSTItem{key: maxKey}
}

func (i *IndexSTItem) Key() uint {
  return i.key
}

func (i *IndexSTItem) IsNull() bool {
  return i.key == maxKey
}

func (i *IndexSTItem) Rand() {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  i.key = uint(r.Uint32()) % maxKey
  i.value = r.Float32()
}

func (i *IndexSTItem) Print(w io.Writer) {
  w.Write([]byte(fmt.Sprintf("%d: %g\n", i.key, i.value)))
}

func (i *IndexSTItem) Value() float32 {
  return i.value
}

func (i *IndexSTItem) SetValue(v float32) {
  i.value = v
}

type IndexST struct {
  st [](*IndexSTItem)
}

func New() *IndexST {
  st := make([](*IndexSTItem), maxKey)
  return &IndexST{st}
}

func (s *IndexST) Search(key uint) *IndexSTItem {
  return s.st[key]
}

func (s *IndexST) Insert(i *IndexSTItem) {
  s.st[i.Key()] = i
}

func (s *IndexST) Remove(i *IndexSTItem) {
  s.st[i.Key()] = nil
}

func (s *IndexST) Select(k uint) *IndexSTItem {
  for i := 0; i < len(s.st); i++ {
    if s.st[i] != nil && !s.st[i].IsNull() {
      if k == 0 { return s.st[i]; }
      k -= 1
    }
  }
  return NewItem()
}

func (s *IndexST) Count() uint {
  var result uint
  for i := 0; i < len(s.st); i++ {
    if s.st[i] != nil && !s.st[i].IsNull() {
      result += 1
    }
  }
  return result
}

func (s *IndexST) Print(w io.Writer) {
  for i := 0; i < len(s.st); i++ {
    if s.st[i] != nil && !s.st[i].IsNull() {
      s.st[i].Print(w)
    }
  }
}
