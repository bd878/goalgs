package queue_test

import (
  "testing"

  ds "github.com/bd878/goalgs/ds/queue"
)

func TestQueue(t *testing.T) {
  q := ds.New[int](10)

  q.Enqueue(1)
  if v := q.Dequeue(); v != 1 {
    t.Errorf("failed to dequeue")
  }
}