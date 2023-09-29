package linkedlist

type LLNode[T interface{}] interface {
  Insert(n LLNode[T])
  DeleteNext()
  Next() LLNode[T]
  Item() T
  Traverse(func(LLNode[T]))
}
