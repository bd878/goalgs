package stack

type Stack[T interface{}] interface {
  Push(T)
  Pop() (T, error)
  IsEmpty() bool
}