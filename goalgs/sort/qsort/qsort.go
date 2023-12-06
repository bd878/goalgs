package qsort

import (
  "fmt"
  "math"
  "golang.org/x/exp/constraints"

  ds "github.com/bd878/goalgs/ds/stack"
  algs "github.com/bd878/goalgs/sort/insort"
)

const bitsbyte uint8 = 8
const bitsword uint8 = 32
const bytesword uint8 = bitsword/bitsbyte

func QSortRecursive[T constraints.Ordered](a []T, l, r int) {
  if l >= r {
    return;
  }
  i := Part[T](a, l, r)
  QSortRecursive[T](a, l, i-1)
  QSortRecursive[T](a, i+1, r)
}

func QSortInsort[T constraints.Ordered](a []T, l, r int) {
  if l >= r {
    return;
  }
  if r-l <= 15 { // or any other file size, [0,25] preferred
    algs.InsortRange[T](a, l, r)
    return
  }
  i := Part[T](a, l, r)
  QSortRecursive[T](a, l, i-1)
  QSortRecursive[T](a, i+1, r)
}

func exch[T constraints.Ordered](a []T, l, r int) {
  a[l], a[r] = a[r], a[l]
}

func compexch[T constraints.Ordered](a []T, l, r int) {
  if a[l] > a[r] {
    exch[T](a, l, r)
  }
}

func qsortMedian[T constraints.Ordered](a []T, l, r int) {
  if r-l <= 15 {
    return; // keep inversions as is
  }
  // select median between most left, middle and most right
  m := (l+r)/2 // middle
  compexch[T](a, l, m)
  compexch[T](a, l, r)
  compexch[T](a, m, r) // a[l] <= a[m] <= a[r]
  exch[T](a, m, r-1)

  // Part() select most right as separator.
  // l and r are sorted already in a way that,
  // they form an inversion which
  // will be precisely sorted by insort later
  i := Part[T](a, l+1, r-1)
  qsortMedian[T](a, l, i-1)
  qsortMedian[T](a, i+1, r)
}

func HybridQSort[T constraints.Ordered](a []T, l, r int) {
  qsortMedian[T](a, l, r)
  algs.InsortRange[T](a, l, r) // handle small inversions
}

func QSort3[T constraints.Ordered](a []T, l, r int) {
  if r <= l {
    return
  }

  v := a[r]
  i, j := l-1, r
  p, q := l, r-1

  // place elements, equal v at the beginning and at the end
  // just to collect it in the center after
  for {
    i += 1
    for a[i] < v && i < j { i += 1 }
    j -= 1
    for a[j] > v && i < j { j -= 1 }

    if i >= j { break }

    a[i], a[j] = a[j], a[i] // in case when a[i] > v && a[j] < v
    if (a[i] == v && i > p) { a[p], a[i] = a[i], a[p]; p += 1 } // put a[i] at the beginning
    if (a[j] == v && j < q) { a[q], a[j] = a[j], a[q]; q -= 1 } // put a[j] at the end
  }

  a[i], a[r] = a[r], a[i] // put v in the middle position a[i]
  // i - is the middle, where v is
  j = max(l, i-1)
  i = min(i+1, r)
  // a[j] == v downside
  for k := l; k < p && j > l; k++ { a[k], a[j] = a[j], a[k]; j -= 1 }
  // a[i] == v upside
  for k := r-1; k > q && i < r-1; k-- { a[k], a[i] = a[i], a[k]; i += 1 }
  QSort3[T](a, l, j)
  QSort3[T](a, i, r)
}

// k - k'th least element in array
func SelectionMedian[T constraints.Ordered](a []T, l, r, k int) {
  if r <= l { return }
  i := Part[T](a, l, r)
  if i > k { SelectionMedian[T](a, l, i-1, k) }
  if i < k { SelectionMedian[T](a, i+1, r, k) }
  // i == k : k is already point to k'th least element
}

func Part[T constraints.Ordered](a []T, l, r int) int {
  v := a[r] // consider most right element to separate around it
  i, j := l, r-1
  for i < j {
    for ; a[i] < v && i < j; i++ {}
    for ; a[j] > v && j > i; j-- {}
    if i < j {
      a[i], a[j] = a[j], a[i]
    }
  }
  if a[r] < a[i] {
    a[i], a[r] = a[r], a[i]
    return i
  } else {
    return r
  }
}

func push(s ds.Stack[int], l, r int) {
  s.Push(r)
  s.Push(l)
}

func QSort[T constraints.Ordered](a []T, l, r int) {
  s := ds.NewArrStack[int]()
  push(s, l, r)

  for !s.IsEmpty() {
    l, _ = s.Pop()
    r, _ = s.Pop()
    if r <= l {
      continue
    }

    i := Part[T](a, l, r)
    if i > r-i { // i is above the half
      push(s, l, i-1) // larger period goes first
      push(s, i+1, r) // take smaller distance on next iteration
    } else {
      push(s, i+1, r) // larger period goes first
      push(s, l, i-1) // take smaller then
    }
  }
}

// derives pos'th bit from num
func getbit(num int32, pos uint8) uint8 {
  bytecount := pos / bitsbyte
  var bytemask int32 = (1 << bitsbyte) - 1 // to zero all digits on the left

  byte := (num >> (bitsbyte * bytecount)) & bytemask

  bitcount := pos % bitsbyte
  bit := (byte >> bitcount) & 1 // mask 1 single bit
  return uint8(bit)
}

func printbin(a []int) {
  for i := 0; i < len(a); i++ {
    fmt.Printf("%#08b ", a[i])
  }
}

func digit(num int, pos int) int {
  return (num / int(math.Pow10(pos))) % 10
}

func qsortb(a []int, l, r, d int) {
  i, j := l, r
  if r <= l || d < 0 { return; }

  for i < j {
    // 0's left, 1's right
    for getbit(int32(a[i]), uint8(d)) == 0 && (i < j) { i += 1 }
    for getbit(int32(a[j]), uint8(d)) == 1 && (j > i) { j -= 1 }
    if i < j {
      a[i], a[j] = a[j], a[i]
    }
  }

  if getbit(int32(a[r]), uint8(d)) == 0 { j += 1; }
  qsortb(a, l, j-1, d-1)
  qsortb(a, j, r, d-1)
}

func QSortB(a []int, l, r int) {
  qsortb(a, l, r, int(bitsbyte))
}

// based on qsort3 and radixMSD
func qsort3Radix(a []int, l, r, d int) {
  if d <= 0 { return; }
  if r <= l { return; }

  v := digit(a[r], d)
  i, j := l-1, r
  p, q := l-1, r

  // fmt.Printf("%d: %d, %d, %d, %v\n", d, v, l, r, a)
  for i < j {
    i += 1
    for digit(a[i], d) < v { i += 1; }
    j -= 1 // weil v am Rechts ist
    for digit(a[j], d) > v && j > l { j -= 1; }

    // if all elements on the left are less than v,
    // i == r, then do not change r
    if i > j { break; }

    a[i], a[j] = a[j], a[i]
    if digit(a[i], d) == v { p += 1; a[p], a[i] = a[i], a[p] }
    if digit(a[j], d) == v { q -= 1; a[q], a[j] = a[j], a[q] }
  }

  // fmt.Printf("%d, %d, %d, %d, %v\n",p, q, i, j, a)

  // all M. digits are same
  if math.Abs(float64(p-q)) <= 1 {
    // len(key) < keysize, aber fing bei keysize an
    qsort3Radix(a, l, r, d-1)
    return;
  }
  // if j == i, we first exchange p on j,
  // there is digit() == v on i position left,
  // should not exchange with r, hence, move forward
  if digit(a[i], d) < v { i += 1; }

  // move elements to center
  for k := l; k <= p; k++ { a[k], a[j] = a[j], a[k]; j -= 1; }
  for k := r; k >= q; k-- { a[k], a[i] = a[i], a[k]; i += 1; }
  // sort left end with the same digit
  qsort3Radix(a, l, j, d)
  // all elements on the right were the same
  if (i == r) && (digit(a[i], d) == v) { i += 1; } // i > r
  // sort the middle
  if v != 0 { qsort3Radix(a, j+1, i-1, d-1); }
  qsort3Radix(a, i, r, d)
}

func QSort3Radix(a []int, l, r int) {
  const keysize int = 5
  qsort3Radix(a, l, r, keysize)
}