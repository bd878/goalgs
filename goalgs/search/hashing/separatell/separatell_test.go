package separatell_test

import (
  "testing"
  "github.com/bd878/goalgs/search/hashing/separatell"
)

type IntItem struct {
  value int
  key int
}

func (v *IntItem) Key() int {
  return v.key
}

func (v *IntItem) Value() int {
  return v.value
}

func TestSeparateLL(t *testing.T) {
  var err error

  amount := 30
  bucketsCount := 5
  exist := 6
  notExist := 25

  table := separatell.NewHashtableLL(amount, bucketsCount)
  values := []int{
    1, exist, /* 1 */
    2, 7, 12, 17, /* 2 */
    19, 24, /* 4*/
  }
  for _, v := range values {
    table.Insert(&IntItem{value: v, key: v})
  }

  _, err = table.Search(exist)
  if err != nil {
    t.Errorf("must be in table")
  }

  _, err = table.Search(notExist)
  if err == nil {
    t.Error("must not exist")
  }
}