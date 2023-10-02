package linkedlist

type SentinelLLNode[T interface{}] struct {
  value T
  next *SentinelLLNode[T]
}

func InitSentinelLL[T interface{}]() *SentinelLLNode[T] {
  x := &SentinelLLNode[T]{}
  x.next = x
  return x
}

func NewSentinelNode[T interface{}](value T) *SentinelLLNode[T] {
  return &SentinelLLNode[T]{value: value}
}

func (x *SentinelLLNode[T]) Insert(t *SentinelLLNode[T]) {
  t.next = x.next
  x.next = t
}

func (x *SentinelLLNode[T]) DeleteNext() {
  x.next = x.next.next
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
