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

// returns next inserted note
func (x *PtrNode[T]) Insert(n LLNode[T]) LLNode[T] {
  var t, ok = n.(*PtrNode[T])
  if !ok {
    panic("not *PtrNode[T]")
  }

  if x.IsEmpty() {
    // make this node new node with value
    *x = *t
    x.next = nil
    return x
  } else {
    t.next = x.next
    x.next = t
    return t
  }
}

// returns next note
func (x *PtrNode[T]) SetNext(n LLNode[T]) LLNode[T] {
  var t, ok = n.(*PtrNode[T])
  if !ok {
    panic("not *PtrNode[T")
  }

  if x.IsEmpty() {
    *x = *t
    return x
  } else {
    x.next = t
    return t
  }
}

func (x *PtrNode[T]) DeleteNext() LLNode[T] {
  if x.next != nil {
    t := x.next
    x.next = t.next
    return t
  } else {
    t := *x
    *x = *(InitPtrLL[T]())
    return &t
  }
}

func (x *PtrNode[T]) Next() LLNode[T] {
  return x.next
}

func (x *PtrNode[T]) Item() T {
  return x.value
}

func (x *PtrNode[T]) Traverse(fn func(LLNode[T])) {
  if x.IsEmpty() {
    return
  }

  for t := x; t != nil; t = t.next {
    fn(t)
  }
}

// we introduce init flag in case
// ll would contain empty nodes only
func (x *PtrNode[T]) IsEmpty() bool {
  if x == nil { return true }
  return x.next == nil && x.init
}