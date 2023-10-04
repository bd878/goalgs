package linkedlist

// TODO: PtrCyclicNode

type PtrNode[T interface{}] struct {
  value T
  next *PtrNode[T]
}

func NewPtrNode[T interface{}](value T) *PtrNode[T] {
  return &PtrNode[T]{value: value}
}

// returns next node
func (x *PtrNode[T]) Insert(t *PtrNode[T]) *PtrNode[T] {
  if x == nil {
    x = t
    x.next = nil
    return x
  } else {
    t.next = x.Next()
    x.next = t
    return t
  }
}

func (x *PtrNode[T]) DeleteNext() *PtrNode[T] {
  if !x.IsEmpty() {
    t := x.Next()
    if !t.IsEmpty() {
      x.next = t.Next()
    }
    return t
  }
  return nil
}

func (x *PtrNode[T]) Next() *PtrNode[T] {
  if !x.IsEmpty() {
    return x.next
  } else {
    return nil
  }
}

func (x *PtrNode[T]) Item() T {
  if !x.IsEmpty() {
    return x.value
  } else {
    var result T
    return result
  }
}

func (x *PtrNode[T]) Traverse(fn func(*PtrNode[T])) {
  if x.IsEmpty() {
    return
  }
  for t := x; t != nil; t = t.Next() {
    fn(x)
  }
}

func (x *PtrNode[T]) IsEmpty() bool {
  return x == nil
}