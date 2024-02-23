package indexing_test

import (
  "log"
  "io"
  "testing"
  "strings"
  search "github.com/bd878/goalgs/search/indexing"
)

func TestStringsIndexing(t *testing.T) {
  for scenario, fn := range map[string]func(*testing.T) {
    "text found": testStringFound,
  } {
    t.Run(scenario, fn)
  }
}

func testStringFound(t *testing.T) {
  log.SetOutput(io.Discard)

  var index int = 10

  idx := search.NewIndex(strings.NewReader("find this text"))

  n, err := idx.Search("text")
  if err != nil {
    t.Error(err)
  }
  if n != index {
    t.Errorf("wrong index found: %d != %d\n", n, index)
  }
}