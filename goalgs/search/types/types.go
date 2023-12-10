package types

import "golang.org/x/exp/constraints"

type Item[K interface{}] interface {
  Key() K
  IsNull() bool
  Rand()
}

type IndexedItem[K constraints.Integer] interface {
  Item[K]
}

type SearchTable[K interface{}, I Item[K]] interface {
  Search(K) I
  Insert(I)
  Remove(I)
  Select(i int) I // select i'th grossest Item
  Count() int
}