package linkedlist

type CyclicNode[T interface{}] struct {
  value T
  next *CyclicNode[T]
}

// cyclic && first node is NOT empty
func NewCyclicNode[T interface{}](value T) *CyclicNode[T] {
  x := &CyclicNode[T]{value: value}
  x.next = x
  return x
}

func (x *CyclicNode[T]) Insert(t *CyclicNode[T]) *CyclicNode[T] {
  t.next = x.next
  x.next = t
  return t
}

func (x *CyclicNode[T]) DeleteNext() *CyclicNode[T] {
  result := x.Next()
  x.next = x.Next().Next()
  return result
}

func (x *CyclicNode[T]) Next() *CyclicNode[T] {
  return x.next
}

func (x *CyclicNode[T]) Item() T {
  return x.value
}

func (x *CyclicNode[T]) Traverse(fn func(*CyclicNode[T])) {
  fn(x)
  for t := x.Next(); t != x; t = t.Next() {
    fn(t)
  } 
}