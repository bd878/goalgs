package main

import (
  "os"
  "flag"
  "testing"
  "slices"
)

var maxn = flag.Int("maxn", 32, "how many prime numbers")

var firstPrimes = []int{2, 3, 5, 7, 11, 13}

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestErat(t *testing.T) {
  primes := Erat(*maxn)
  if slices.Compare(primes[:len(firstPrimes)], firstPrimes) != 0 {
    t.Errorf("result does not contain first %d primes\n", len(firstPrimes))
  }
}