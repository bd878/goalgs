package bst

// TODO: rewrite on ds/tree/binarytree
type BTreeNode[I interface{}] struct {
  Item I
  N int // internal nodes in branch
  L *BTreeNode[I]
  R *BTreeNode[I]
}

type RBBTreeNode[I interface{}] struct {
  Item I
  N int
  L *RBBTreeNode[I]
  R *RBBTreeNode[I]

  red bool // for Red-Black trees
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

func (n *RBBTreeNode[I]) Red() bool {
  if n == nil {
    return false
  }
  return n.red
}