package bst

import "math/rand"

const MAX_KEY int = 10e2

const MAX_LEN int = 10

var chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
  "abcdefghijklmnopqurstuvwxyz" + "123456789"

type StringItem struct {
  key int
  value string
}

func (si *StringItem) SetKey(k int) {
  si.key = k
}

func (si *StringItem) Key() int {
  return si.key
}

func (si *StringItem) Null() bool {
  return si.key == 0
}

func (si *StringItem) Rand() {
  si.key = rand.Intn(MAX_KEY)+1
  for i := 0; i < MAX_LEN; i++ {
    si.value += string(chars[rand.Intn(len(chars))])
  }
}

type IntItem struct {
  key int
  value int
}

func (ii *IntItem) SetKey(k int) {
  ii.key = k
}

func (ii *IntItem) Key() int {
  return ii.key
}

func (ii *IntItem) Null() bool {
  return ii.key == 0
}

func (ii *IntItem) Rand() {
  ii.key = rand.Intn(MAX_KEY)+1
  ii.value = rand.Intn(MAX_KEY)
}
