package hashing

type Hasher interface {
  Hash() int
}

func HashInt(v, m int) int {
  return v & (m - 1)
}
