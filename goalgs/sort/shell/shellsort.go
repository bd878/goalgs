package shellsort

import "golang.org/x/exp/constraints"

func Shellsort[T constraints.Ordered](a []T, l, r int) {
  var h int
  for h = 1; h <= (r-l)/9; h = 3*h+1 {}
  for ; h > 0; h /= 3 {
    for i := l+h; i < r; i++ {
      j := i;
      v := a[i]

      for j >= l+h && v < a[j-h] {
        a[j] = a[j-h]
        j -= h
      }
      a[j] = v
    }
  }
}