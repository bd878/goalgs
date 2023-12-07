package radix

import (
  "math"
)

const keysize int = 8
const bitsbyte int = 8
const baseR int = 10

func digit(num int, pos int) int {
  return (num / int(math.Pow10(pos))) % 10
}

// based on distcount alg
func radixMSD(a []int, l, r, d int) {
  // define local static variable
  // temp buffer for keys
  aux := make([]int, 10e3) // MaxN

  var radix func([]int, int, int, int)
  radix = func(a []int, l, r, d int) {
    if d <= 0 { return; }
    if r <= l { return; }
    // crop small-size files
    // if r-l < 15 { algs.InsortRange[int](a, l, r); return; }

    // partial sums
    count := make([]int, baseR+1);
     // count nums for each digit
    for i := l; i <= r; i++ { count[digit(a[i], d) + 1] += 1 }
    // count elements below
    for j := 1; j < baseR; j++ { count[j] += count[j-1] }
    // count elements below and on this position
    for i := l; i <= r; i++ {
      aux[l+count[digit(a[i], d)]] = a[i]
      // count grows up (keeps space for values below),
      // hence += 1
      count[digit(a[i], d)] += 1
    }
    // restore an order
    for i := l; i <= r; i++ { a[i] = aux[i] }
    // to handle trailing zeroes, all counts are the same = r-l
    radix(a, l, l+count[0]-1, d-1)
    // baseR-1, because take gap from count[j] to count[j+1] 
    for j := 0; j < baseR-1; j++ {
      radix(a, l+count[j], l+count[j+1]-1, d-1)
    }
  }

  radix(a, l, r, d)
}

func RadixMSD(a []int, l, r int) {
  radixMSD(a, l, r, keysize)
}

// we can't use qsort.Part, because it is unstable.
// Instead, we use stable index count sort
func RadixLSD(a []int, l, r int) {
  const bytesword int = 10

  aux := make([]int, r-l+1)

  for d := 0; d <= bytesword; d++ {
    count := make([]int, baseR+1)

    for i := l; i <= r; i++ { count[digit(a[i], d)+1] += 1; }
    for j := 1; j < baseR; j++ { count[j] += count[j-1]; }
    for i := l; i <= r; i++ {
      aux[count[digit(a[i], d)]] = a[i]
      count[digit(a[i], d)] += 1
    }
    for i := l; i <= r; i++ { a[i] = aux[i]; }
  }
}