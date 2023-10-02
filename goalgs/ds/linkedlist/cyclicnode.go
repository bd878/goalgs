package linkedlist

type CyclicNode[T interface{}] struct {
  value T
  next *CyclicNode[T]
}

func NewCyclicNode[T interface{}](value T) *CyclicNode[T] {
  x := &CyclicNode[T]{value: value}
  x.next = x
  return x
}

func (x *CyclicNode[T]) Insert(t *CyclicNode[T]) {
  t.next = x.next
  x.next = t
}

func (x *CyclicNode[T]) DeleteNext() {
  x.next = x.next.next
}

func (x *CyclicNode[T]) Next() *CyclicNode[T] {
  return x.next
}

func (x *CyclicNode[T]) Item() T {
  return x.value
}

func (x *CyclicNode[T]) IsOnlyOne() bool {
  return x.next == x
}

func (x *CyclicNode[T]) Traverse(fn func(*CyclicNode[T])) {
  fn(x)
  if !x.IsOnlyOne() {
    for t := x.Next(); t != x; t = t.Next() {
      fn(t)
    } 
  }
}