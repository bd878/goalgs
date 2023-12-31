package doublelinkedlist

type List[T interface{}] struct {
  i *Node[T]
}

type Node[T interface{}] struct {
  v T
  next *Node[T]
  prev *Node[T]
}

func New[T interface{}]() *List[T] {
  l := &List[T]{}

  i := &Node[T]{}
  i.next = i
  i.prev = i

  l.i = i
  return l
}

func NewNode[T interface{}](v T) *Node[T] {
  return &Node[T]{v: v}
}

func (l *List[T]) Head() *Node[T] {
  return l.i.next
}

func (l *List[T]) Tail() *Node[T] {
  return l.i.prev
}

func (l *List[T]) Insert(x *Node[T]) *Node[T] {
  x.next = l.i.next
  l.i.next.prev = x
  l.i.next = x
  x.prev = l.i
  return x
}

func (l *List[T]) Delete(x *Node[T]) {
  x.prev.next = x.next
  x.next.prev = x.prev
}

func (x *Node[T]) Item() T {
  return x.v
}

func (x *Node[T]) Next() *Node[T] {
  return x.next
}

func (x *Node[T]) Prev() *Node[T] {
  return x.prev
}

func (x *Node[T]) SetItem(v T) {
  x.v = v
}

func (l *List[T]) Traverse(fn func(*Node[T])) {
  for n := l.Head(); n != l.i; n = n.next {
    fn(n)
  }
}

func (l *List[T]) Empty() bool {
  return l.i.next == l.i
}