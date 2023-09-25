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

func MergeStop(a []int, l, m, r int) {
  var maxn int

  n1 := m-l+1
  n2 := r-m

  ln := make([]int, n1+1)
  rn := make([]int, n2+1)
  for i := 0; i < n1; i++ {
    ln[i] = a[l+i]
    maxn = max(maxn, ln[i])
  }
  for j := 0; j < n2; j++ {
    rn[j] = a[m+j+1]
    maxn = max(maxn, rn[j])
  }

  // signal flags
  maxn++
  ln[n1] = maxn
  rn[n2] = maxn

  for k, i, j := l, 0, 0; k <= r; k++ {
    if ln[i] <= rn[j] {
      a[k] = ln[i]
      i = i+1
    } else {
      a[k] = rn[j]
      j = j+1
    }
  }
}

func MergeNonStop(a []int, l, m, r int) {
  ln := make([]int, m-l+1)
  rn := make([]int, r-m)
  for i := 0; i < len(ln); i++ {
    ln[i] = a[l+i]
  }
  for j := 0; j < len(rn); j++ {
    rn[j] = a[m+j+1]
  }

  n1 := len(ln)
  n2 := len(rn)
  for k, i, j := l, 0, 0; k <= r; k++ {
    if n2 == 0 || (n1 > 0 && ln[i] <= rn[j]) {
      a[k] = ln[i]
      i = i+1
      n1 = n1-1
    } else {
      a[k] = rn[j]
      j = j+1
      n2 = n2-1
    }
  }
}