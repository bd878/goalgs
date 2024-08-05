package skiplist_test

import (
  "testing"

  "github.com/bd878/goalgs/search/skiplist"
)

func TestInsert(t *testing.T) {
  sl := skiplist.NewSkiplist()

  nonExistentNode := skiplist.NewIntItem(1234, 4321)

  slNode1 := skiplist.NewIntItem(1, 101)
  slNode2 := skiplist.NewIntItem(2, 102)
  slNode5 := skiplist.NewIntItem(5, 105)
  sl.Insert(slNode1)
  sl.Insert(slNode2)
  sl.Insert(slNode5)

  if _, err := sl.Search(slNode2.Key()); err != nil {
    t.Errorf("node not found for key: %d\n", slNode2.Key())
  }

  if _, err := sl.Search(slNode5.Key()); err != nil {
    t.Errorf("node not found for key: %d\n", slNode5.Key())
  }

  sl.Remove(slNode5.Key())

  if _, err := sl.Search(slNode5.Key()); err == nil {
    t.Errorf("node is found for non-existent key: %d\n", slNode5.Key())
  }

  if _, err := sl.Search(nonExistentNode.Key()); err == nil {
    t.Errorf("node is found for non-existent key: %d\n", nonExistentNode.Key())
  }
}