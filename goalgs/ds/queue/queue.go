package queue

type Queue[T interface{}] struct {
  q []T
  n int
  count int
  head int
  tail int
}

func New[T interface{}](n int) *Queue[T] {
  q := Queue[T]{}
  q.q = make([]T, n)
  q.n = n

  return &q
}

// TODO: test it!
func (q *Queue[T]) Resize(n int) {
  if n < q.n {
    panic("can not schrink")
  }

  // remap head and tail in new slice
  nq := make([]T, n)
  var j = 0
  for i := q.head; i < q.n; i++ {
    nq[j] = q.q[i]
    j += 1
  }
  for i := 0; i < q.tail; i++ {
    nq[j] = q.q[i]
    j += 1
  }
  q.head = 0;
  q.tail = j;

  q.q = nq
  q.n = n
}

func (q *Queue[T]) Enqueue(v T) {
  if q.count == q.n {
    q.Resize(2*q.n)
  }

  q.q[q.tail] = v
  q.tail += 1
  q.tail = q.tail % q.n
  q.count += 1
}

func (q *Queue[T]) Dequeue() T {
  val := q.q[q.head]
  q.head += 1
  q.count -= 1
  q.head = q.head % q.n
  return val
}

func (q *Queue[T]) IsEmpty() bool {
  return q.count == 0
}