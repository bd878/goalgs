package linkedlist

type LLNode[T interface{}] interface {
  // returns next inserted node
  Insert(n LLNode[T]) LLNode[T]
  DeleteNext() LLNode[T]
  Next() LLNode[T]
  // returns next node
  SetNext(n LLNode[T]) LLNode[T]
  Item() T
  Traverse(func(LLNode[T]))
  IsEmpty() bool
}
