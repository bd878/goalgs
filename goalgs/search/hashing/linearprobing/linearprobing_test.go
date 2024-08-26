package linearprobing_test

import (
  "testing"
  "github.com/bd878/goalgs/search/hashing"
  "github.com/bd878/goalgs/search/hashing/linearprobing"
)

func TestLinearProbingHashtable(t *testing.T) {
  var err error

  buckets := 10

  exist := 6
  nonExist := 25

  st := linearprobing.NewLinearProbingHashtable(buckets)
  st.Insert(hashing.NewIntItem(exist, exist))
  if _, err = st.Search(exist); err != nil {
    t.Errorf("must be in table")
  }

  if _, err = st.Search(nonExist); err == nil {
    t.Errorf("must not exist")
  }

  st.Remove(exist)
  if _, err = st.Search(exist); err == nil {
    t.Errorf("removed, must not exist")
  }
}