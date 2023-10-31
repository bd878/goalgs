package binarytree

import (
  "math"
  "fmt"

  stack "github.com/bd878/goalgs/ds/stack"
  queue "github.com/bd878/goalgs/ds/queue"
)

type BTreeNode[T interface{}] struct {
  V T
  L *BTreeNode[T]
  R *BTreeNode[T]
}

func Init[T interface{}]() *BTreeNode[T] {
  return nil
}

func NewNode[T interface{}](v T) *BTreeNode[T] {
  return &BTreeNode[T]{V: v}
}

func (n *BTreeNode[T]) IsEmpty() bool {
  return n == nil
}

func (n *BTreeNode[T]) SetItem(v T) *BTreeNode[T] {
  if n == nil {
    n = &BTreeNode[T]{V: v}
  } else {
    n.V = v
  }
  return n
}

func (n *BTreeNode[T]) Serialize() map[int][]*BTreeNode[T] {
  if n == nil {
    panic("n is nil, SetItem first")
  }

  var e *BTreeNode[T]
  res := make(map[int][]*BTreeNode[T])
  arr := make([]*BTreeNode[T], 1)
  var v *BTreeNode[T]
  var level, pos int

  q := queue.New[*BTreeNode[T]](2)
  q.Enqueue(n)
  q.Enqueue(e)

  for ; !q.IsEmpty(); {
    v = q.Dequeue()

    if v == e {
      if q.IsEmpty() {
        break;
      }

      res[level] = arr
      level += 1
      arr = make([]*BTreeNode[T], int(math.Pow(2, float64(level))))
      pos = 0
      q.Enqueue(e)
      continue
    }

    arr[pos] = v
    pos += 1

    if v.L != nil {
      q.Enqueue(v.L)
    }
    if v.R != nil {
      q.Enqueue(v.R)
    }
  }

  return res
}

func (n *BTreeNode[T]) Height() int {
  if n == nil {
    return -1
  }

  res := n.Serialize()
  return len(res)
}

func (n *BTreeNode[T]) CountTotal() int {
  if n == nil {
    return 0
  }

  var total int

  s := stack.New[*BTreeNode[T]]()
  s.Push(n)
  total += 1
  var top *BTreeNode[T]
  var err error

  for ; !s.IsEmpty(); {
    top, err = s.Pop() // remove root from stack
    if err != nil {
      panic(err)
    }

    if top.R != nil {
      s.Push(top.R)
      total += 1
    }
    if top.L != nil {
      s.Push(top.L)
      total += 1
    }
  }

  return total
}

func (n *BTreeNode[T]) CountTotalRecursive() int {
  if n == nil {
    return 0
  }

  return n.L.CountTotalRecursive() + n.R.CountTotalRecursive() + 1;
}

func (n *BTreeNode[T]) HeightRecursive() int {
  if n == nil {
    return 0;
  }

  lHeight := n.L.HeightRecursive()
  rHeight := n.R.HeightRecursive()

  return max(lHeight, rHeight) + 1;
}

func (n *BTreeNode[T]) TraverseUpDown(visit func(*BTreeNode[T])) {
  if n == nil {
    return;
  }

  s := stack.New[*BTreeNode[T]]()
  var v *BTreeNode[T]

  s.Push(n)
  for ; !s.IsEmpty(); {
    v, _ = s.Pop()
    // order is important
    if v.R != nil {
      s.Push(v.R)
    }
    if v.L != nil {
      s.Push(v.L)
    }
    visit(v)
  }
}

// recursive
func (n *BTreeNode[T]) TraverseLeftRight(visit func(*BTreeNode[T])) {
  if n == nil {
    return;
  }

  if n.L == nil && n.R == nil {
    visit(n)
    return
  }

  if n.L != nil {
    visit(n.L)
  }

  visit(n)

  if n.R != nil {
    visit(n.R)
  }
}

// recursive
func (n *BTreeNode[T]) TraverseDownUp(visit func(*BTreeNode[T])) {
  if n == nil {
    return;
  }

  if n.L == nil && n.R == nil {
    visit(n)
    return
  }

  if n.L != nil {
    visit(n.L)
  }
  if n.R != nil {
    visit(n.R)
  }
  visit(n)
}

func (n *BTreeNode[T]) TraverseDeep(visit func(*BTreeNode[T])) {
  if n == nil {
    return;
  }

  s := stack.New[*BTreeNode[T]]()
  s.Push(n)
  var v *BTreeNode[T]

  for ; !s.IsEmpty(); {
    v, _ = s.Pop()
    visit(v)

    if v.L != nil {
      s.Push(v.L)
    }
    if v.R != nil {
      s.Push(v.R)
    }
  }
}

func (n *BTreeNode[T]) TraverseLevel(visit func(*BTreeNode[T])) {
  if n == nil {
    return;
  }

  q := queue.New[*BTreeNode[T]](1)
  q.Enqueue(n)
  var v *BTreeNode[T]

  for ; !q.IsEmpty(); {
    v = q.Dequeue()
    visit(v)

    if v.L != nil {
      q.Enqueue(v.L)
    }
    if v.R != nil {
      q.Enqueue(v.R)
    }
  }
}

func (n *BTreeNode[T]) TraverseRecursive(visit func(*BTreeNode[T])) {
  if n == nil {
    return;
  }

  visit(n)

  n.L.TraverseRecursive(visit)
  n.R.TraverseRecursive(visit)
}

// return root because insert is valid to root only
func (n *BTreeNode[T]) Insert(t *BTreeNode[T]) *BTreeNode[T] {
  if n == nil {
    // init tree with value
    return t;
  }

  q := queue.New[*BTreeNode[T]](1)
  var v *BTreeNode[T]

  q.Enqueue(n)
  for ; !q.IsEmpty(); {
    v = q.Dequeue()
    if v.L == nil {
      v.L = t
      break
    } else {
      q.Enqueue(v.L)
    }

    if v.R == nil {
      v.R = t
      break
    } else {
      q.Enqueue(v.R)
    }
  }

  return n
}

func (n *BTreeNode[T]) print(printer func(T, int), h int) {
  if n == nil {
    var t T
    printer(t, h)
    return;
  }

  printer(n.V, h)
  n.R.print(printer, h+1)
  n.L.print(printer, h+1)
}

// recursive
func (n *BTreeNode[T]) Print(printer func(T, int)) {
  n.print(printer, 0)
}

func PrintRune(r rune, h int) {
  if r == 0 {
    fmt.Printf("%" + fmt.Sprint(h+3) + "v\n", "*")
  } else {
    fmt.Printf("%" + fmt.Sprint(h+3) + "q\n", r)
  }
}

func PrintInt(r int, h int) {
  fmt.Printf("%" + fmt.Sprint(h+3) + "d\n", r)
}