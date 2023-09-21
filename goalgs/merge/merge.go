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