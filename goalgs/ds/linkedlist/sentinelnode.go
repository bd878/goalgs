package linkedlist

type SentinelLLNode[T interface{}] struct {
  value T
  next *SentinelLLNode[T]
}

// first in nil
func InitSentinelLL[T interface{}]() *SentinelLLNode[T] {
  x := &SentinelLLNode[T]{}
  x.next = x
  return x
}

func NewSentinelNode[T interface{}](value T) *SentinelLLNode[T] {
  return &SentinelLLNode[T]{value: value}
}

func (x *SentinelLLNode[T]) Insert(t *SentinelLLNode[T]) *SentinelLLNode[T] {
  t.next = x.Next()
  x.next = t
  return t
}

func (x *SentinelLLNode[T]) DeleteNext() *SentinelLLNode[T] {
  t := x.Next()
  x.next = x.Next().Next()
  return t
}

func (x *SentinelLLNode[T]) Next() *SentinelLLNode[T] {
  return x.Next()
}

func (x *SentinelLLNode[T]) Item() T {
  return x.value
}

func (x *SentinelLLNode[T]) Traverse(fn func(*SentinelLLNode[T])) {
  for t := x.Next(); t != x; t = t.Next() {
    fn(t)
  }
}

func (x *SentinelLLNode[T]) IsEmpty() bool {
  return x.Next() == x
}
