package bst

type BTreeNode[I interface{}] struct {
  Item I
  Red bool // red is, when it's upper node is red
  N int // internal nodes in branch
  L *BTreeNode[I]
  R *BTreeNode[I]
}

func (n *BTreeNode[I]) print(printer func(*BTreeNode[I], int), h int) {
  if n == nil {
    printer(nil, h)
    return;
  }

  printer(n, h)
  n.R.print(printer, h+1)
  n.L.print(printer, h+1)
}
