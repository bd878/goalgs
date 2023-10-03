package flavius_test

import (
  "testing"
  "reflect"

  "github.com/bd878/goalgs/misc/flavius"
)

func TestFlavius(t *testing.T) {
  result := flavius.Run(9, 5)
  check := []int{5,1,7,4,3,6,9,2,8}
  if !reflect.DeepEqual(result, check) {
    t.Errorf("wrong result")
  }
}