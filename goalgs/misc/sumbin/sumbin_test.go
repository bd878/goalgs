package sumbin_test

import (
  "testing"

  "github.com/bd878/goalgs/misc/sumbin"
)

func TestSumbin(t *testing.T) {
  for answ, ops := range map[string][]string{
    "11": []string{"10", "01"},
    "100": []string{"10", "10"},
    "1010": []string{"1000", "10"},
    "10": []string{"1", "1"},
  } {
    res := sumbin.Sum(ops[0], ops[1])
    t.Logf("%s + %s = %s, given %s", ops[0], ops[1], answ, res)
    if res != answ {
      t.Error("res != answ", res, answ)
    }
  }
}
