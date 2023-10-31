package prefixtree

import (
  tree "github.com/bd878/goalgs/ds/tree"
)

func Parse(a []rune) *tree.BTreeNode[rune] {
  i := 0
  return parse(a, &i)
}

// TODO: error handling
func parse(a []rune, i *int) *tree.BTreeNode[rune] {
  t := a[*i]
  x := tree.NewNode[rune](t)
  *i += 1

  if (t == '+') || (t == '-') {
    x.L = parse(a, i)
    x.R = parse(a, i)
  }

  return x
}