package linkedlist

type PtrNode[T interface{}] struct {
  value T
  next *PtrNode[T]
  init bool
}

func InitPtrLL[T interface{}]() *PtrNode[T] {
  return &PtrNode[T]{init: true}
}

func NewPtrNode[T interface{}](v T) *PtrNode[T] {
  return &PtrNode[T]{value: v}
}

func (x *PtrNode[T]) Insert(t *PtrNode[T]) *PtrNode[T] {
  if x.IsEmpty() {
    // make this node new node with value
    *x = *t
    x.next = nil
    return x
  } else {
    t.next = x.Next()
    x.next = t
    return t
  }
}

func (x *PtrNode[T]) DeleteNext() *PtrNode[T] {
  if x.next != nil {
    t := x.Next()
    x.next = t.Next()
    return t
  } else {
    t := *x
    *x = *(InitPtrLL[T]())
    return &t
  }
}

func (x *PtrNode[T]) Next() *PtrNode[T] {
  return x.next
}

func (x *PtrNode[T]) Item() T {
  return x.value
}

func (x *PtrNode[T]) Traverse(fn func(*PtrNode[T])) {
  if x.IsEmpty() {
    return
  }

  for t := x; t != nil; t = t.Next() {
    fn(t)
  }
}

func (x *PtrNode[T]) IsEmpty() bool {
  return x.next == nil && x.init
}