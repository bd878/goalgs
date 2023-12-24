package linkedlist

type CyclicNode[T interface{}] struct {
  value T
  next *CyclicNode[T]
  init bool
}

// cyclic && first node is NOT empty
func NewCyclicNode[T interface{}](value T) *CyclicNode[T] {
  x := &CyclicNode[T]{value: value}
  x.next = x
  return x
}

func initCyclicNode[T interface{}]() *CyclicNode[T] {
  x := &CyclicNode[T]{init:true}
  x.next = x
  return x
}

func (x *CyclicNode[T]) Insert(t *CyclicNode[T]) *CyclicNode[T] {
  t.next = x.next
  x.next = t
  return t
}

func (x *CyclicNode[T]) DeleteNext() *CyclicNode[T] {
  if x.Next().IsEmpty() {
    result := *x
    *x = *(initCyclicNode[T]())
    return &result
  }

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

func (x *CyclicNode[T]) IsEmpty() bool {
  return x.next == x && x.init
}