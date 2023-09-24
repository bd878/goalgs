package main

func MergeAB(a []int, n int, b []int, m int) []int {
  c := make([]int, n + m)

  for i, j, k := 0, 0, 0; k < n + m; k++ {
    if i == n { c[k] = b[j]; j++; continue; }
    if j == m { c[k] = a[i]; i++; continue; }

    if a[i] < b[j] {
      c[k] = a[i]
      i++
    } else {
      c[k] = b[j]
      j++
    }
  } 

  return c
}

func Merge(a []int, l, m, r int) {
  var i, j int
  var total, ll, rl, rr = r-l+1, 0, m+1, r-l
  aux := make([]int, total)

  for i = l; i <= m; i++ {
    aux[i-l] = a[i]
  }
  i = ll

  // reverse order
  for j = rl; j <= r; j++ {
    aux[(r-l)-(j-rl)] = a[j]
  }
  j = rr

  for k := l; k <= r; k++ {
    // most right less than most left
    if aux[j] < aux[i] {
      a[k] = aux[j]
      j = j-1
    } else {
      a[k] = aux[i]
      i = i+1
    }
  }
}