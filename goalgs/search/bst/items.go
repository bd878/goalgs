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
