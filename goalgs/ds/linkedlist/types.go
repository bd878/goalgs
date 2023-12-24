package linkedlist

type LLNode[T interface{}] interface {
  Insert(n LLNode[T]) LLNode[T]
  DeleteNext() LLNode[T]
  Next() LLNode[T]
  Item() T
  Traverse(func(LLNode[T]))
  IsEmpty() bool
}
