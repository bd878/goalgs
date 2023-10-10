package doublelinkedlist

type DoubleLL[T interface{}] struct {
  i *Node[T]
}

type Node[T interface{}] struct {
  v T
  next *Node[T]
  prev *Node[T]
}

func New[T interface{}]() *DoubleLL[T] {
  l := &DoubleLL[T]{}

  i := &Node[T]{}
  i.next = i
  i.prev = i

  l.i = i
  return l
}

func NewNode[T interface{}](v T) *Node[T] {
  return &Node[T]{v: v}
}

func (l *DoubleLL[T]) Head() *Node[T] {
  return l.i.next
}

func (l *DoubleLL[T]) Tail() *Node[T] {
  return l.i.prev
}

func (l *DoubleLL[T]) Insert(x *Node[T]) *Node[T] {
  x.next = l.i.next
  l.i.next.prev = x
  l.i.next = x
  x.prev = l.i
  return x
}

func (l *DoubleLL[T]) Delete(x *Node[T]) {
  x.prev.next = x.next
  x.next.prev = x.prev
}

func (x *Node[T]) Item() T {
  return x.v
}

func (l *DoubleLL[T]) Traverse(fn func(*Node[T])) {
  for n := l.i.next; n != l.i; n = n.next {
    fn(n)
  }
}
