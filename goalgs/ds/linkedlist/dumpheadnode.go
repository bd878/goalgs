package linkedlist

type DumpHeadLLNode[T interface{}] struct {
  value T
  next *DumpHeadLLNode[T]
}

func NewDumpHeadNode[T interface{}]() *DumpHeadLLNode[T] {
  return &DumpHeadLLNode[T]{}
}

func NewDumpHeadLLNode[T interface{}](value T) *DumpHeadLLNode[T] {
  return &DumpHeadLLNode[T]{value: value}
}

func (x *DumpHeadLLNode[T]) Insert(t *DumpHeadLLNode[T]) {
  if x.IsEmpty() {
    x.next = t
  } else {
    t.next = x.Next()
    x.next = t
  }
}

func (x *DumpHeadLLNode[T]) DeleteNext() {
  if !x.IsEmpty() {
    x.next = x.next.next
  }
}

func (x *DumpHeadLLNode[T]) Next() *DumpHeadLLNode[T] {
  return x.next
}

func (x *DumpHeadLLNode[T]) Item() T {
  return x.value
}

func (x *DumpHeadLLNode[T]) Traverse(fn func(*DumpHeadLLNode[T])) {
  for t := x.Next(); t != nil; t = t.Next() {
    fn(t)
  }
}

func (x *DumpHeadLLNode[T]) IsEmpty() bool {
  return x.Next() == nil
}