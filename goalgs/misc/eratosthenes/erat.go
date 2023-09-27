package main

func Erat(n int) []int {
  if n > 1 {
    return erat(n)
  } else {
    return []int{}
  }
}

func erat(n int) []int {
  nums := make([]int, n)
  for i := 2; i < n; i++ {
    if nums[i] == 0 {
      for j := i; j*i < n; j++ {
        nums[i*j] = 1
      }
    }
  }

  primes := make([]int, 0)
  for i := 2; i < n; i++ {
    if nums[i] == 0 {
      primes = append(primes, i)
    }
  }

  return primes
}