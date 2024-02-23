package indexing

import (
  "io"
  "log"
  "errors"
  "github.com/bd878/goalgs/search/bst"
  "github.com/bd878/goalgs/search/types"
)

const MAX_LEN = 10e2

// text index is global, that each StringItem
// could get a Key through idx
var textIndex []byte = make([]byte, MAX_LEN)

type StringItem[K interface{ string }] struct {
  idx int
}

func (s *StringItem[K]) Key() string {
  // TODO: bad bad bad, return []byte or store index as single string
  return string(textIndex[s.idx:]) // till end, since slice is only a reference
}

func (s *StringItem[K]) Null() bool {
  return string(textIndex[s.idx]) == ""
}

func (s *StringItem[K]) Rand() {
  panic("not implemented")
}

type TextIndex struct {
  Corpus types.SearchTable[string, types.Item[string]]
}

func NewIndex(fio io.Reader) *TextIndex {
  corpus := bst.NewBinaryST[string, types.Item[string]]()
  n, err := fio.Read(textIndex)
  if err != nil {
    panic(err)
  }
  log.Printf("read %d bytes\n", n)
  textIndex = textIndex[:n]
  // TODO: improve: index by whole words, not by symbols
  for i := 0; i < len(textIndex); i++ {
    corpus.Insert(&StringItem[string]{idx:i})
  }

  return &TextIndex{Corpus:corpus}
}

func (i *TextIndex) Search(v string) (int, error) {
  item := i.Corpus.Search(v)
  if item == nil {
    return -1, errors.New("not found")
  }
  // TODO: can we go better, than type conversion???
  sItem, ok := item.(*StringItem[string])
  if !ok {
    panic("not a StringItem")
  }
  log.Printf("found string %s at position %d in index\n", 
    v, sItem.idx)
  return sItem.idx, nil
}

func (i *TextIndex) Print() {
  c, ok := i.Corpus.(*bst.BinaryST[string, types.Item[string]])
  if !ok {
    panic("not binary st")
  }
  c.Print()
}