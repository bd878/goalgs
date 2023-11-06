package stack_test

import (
  "testing"
  ds "github.com/bd878/goalgs/ds/stack"
)

func TestStack(t *testing.T) {
  for scenario, fn := range map[string]func(*testing.T){
    "stack push and pop values": stackPushValues,
    "stack is empty": stackIsEmpty,
  } {
    t.Run(scenario, fn)
  }
}

func stackPushValues(t *testing.T) {
  for _, st := range []ds.Stack[int]{
    ds.NewArrStack[int](),
    ds.NewLLStack[int](),
  } {
    st.Push(1)
    v, err := st.Pop()
    if err != nil {
      t.Error("stack returned error:", err)
    }

    if v != 1 {
      t.Errorf("popped wrong value")
    }
  }
}

func stackIsEmpty(t *testing.T) {
  for _, st := range []ds.Stack[int]{
    ds.NewArrStack[int](),
    ds.NewLLStack[int](),
  } {
    st.Push(1)

    if st.IsEmpty() {
      t.Errorf("stack empty")
    }
    st.Pop()
    if !st.IsEmpty() {
      t.Errorf("stack not empty")
    }
  }
}