package prefixtree_test

import (
  "testing"

  tree "github.com/bd878/goalgs/ds/tree"
  "github.com/bd878/goalgs/misc/prefixtree"
)

func TestPrefixtree(t *testing.T) {
  arr := []rune{'+', '-', '3', '2', '1'}
  root := prefixtree.Parse(arr)

  if root.CountTotal() != len(arr) {
    t.Error("CountTotal != len", root.CountTotal(), len(arr))
  }

  root.Print(tree.PrintRune)
}
