package hashing

import "errors"

var ErrNoItem = errors.New("no item")

type Item interface {
  Key() int
  Value() int
}

type IntItem struct {
  value int
  key int
}

func NewIntItem(value, key int) *IntItem {
  return &IntItem{value, key}
}

func (v *IntItem) Key() int {
  return v.key
}

func (v *IntItem) Value() int {
  return v.value
}