package stack_test

import (
  "testing"
  ds "github.com/bd878/goalgs/ds/stack"
)

func TestStack(t *testing.T) {
  for scenario, fn := range map[string]func(*testing.T){
    "stack push and pop values": stackPushValues,
  } {
    t.Run(scenario, fn)
  }
}

func stackPushValues(t *testing.T) {
  st := &ds.Stack[int]{}

  st.Push(1)
  if i, _ := st.Pop(); i != 1 {
    t.Errorf("popped wrong value")
  }
}