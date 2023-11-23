package distcountsort

const MAX_M = 10e2

func Distcount(a []int, l, r int) {
  cnt := make([]int, MAX_M, MAX_M)
  b := make([]int, len(a))

  for i := l; i <= r; i++ {
    cnt[a[i]+1] += 1 // keeps amount of a[i] keys
  }

  for j := l+1; j < MAX_M; j++ {
    cnt[j] += cnt[j-1] // keeps particle sums of keys < this key
  }

  // move keys to b
  for i := l; i <= r; i++ {
    b[cnt[a[i]]] = a[i] // element a[i] must be on position cnt[a[i]] in slice b
    cnt[a[i]] += 1
  }

  // move back sorted keys from b to a
  for i := l; i <= r; i++ {
    a[i] = b[i-l]
  }
}