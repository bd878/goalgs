package skiplist

type Item interface {
  Key() int
  Value() int
}

type IntItem struct {
  key int
  value int
}

func NewIntItem(key, value int) *IntItem {
  return &IntItem{key, value}
}

func (i *IntItem) Key() int {
  return i.key
}

func (i *IntItem) Value() int {
  return i.value
}