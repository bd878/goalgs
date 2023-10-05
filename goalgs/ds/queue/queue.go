package queue

type Queue[T interface{}] struct {
  q []T
  n int
  head int
  tail int
}

func New[T interface{}](n int) *Queue[T] {
  q := Queue[T]{}
  q.q = make([]T, n)
  q.n = n

  return &q
}

func (q *Queue[T]) Resize(n int) {
  nq := make([]T, 0, n)
  copy(nq, q.q)
  q.q = nq
}

func (q *Queue[T]) Enqueue(v T) {
  q.q[q.tail] = v
  q.tail += 1
  q.tail = q.tail % q.n
}

func (q *Queue[T]) Dequeue() T {
  val := q.q[q.head]
  q.head += 1
  q.head = q.head % q.n
  return val
}

func (q *Queue[T]) Empty() bool {
  return q.head % q.n == q.tail
}