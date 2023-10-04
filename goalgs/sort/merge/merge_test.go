package merge

import (
  "testing"
  "math/rand"
  "os"
  "flag"
  "sort"

  ds "github.com/bd878/goalgs/ds/linkedlist"
)

const MAX_N = 10e1

var (
  size = flag.Int("size", 10e4, "sort size length")
  alen = flag.Int("alen", 10e3, "a length")
  blen = flag.Int("blen", 10e2, "b length")
)

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestMergeSort(t *testing.T) {
  funcs := [](func([]int)){
    Mergesort, MergesortUp,
  }

  for _, fn := range funcs {
    ns := rand.Perm(*size)

    fn(ns)
    if !sort.IsSorted(sort.IntSlice(ns)) {
      t.Error("perm is not sorted")
    }
  }
}

func TestMergeStop(t *testing.T) {
  funcs := [](func([]int, int, int, int)){
    MergeStop, MergeNonStop,
  }

  for _, fn := range funcs {
    a, b := getSorted()

    c := make([]int, 0, a.Len()+b.Len())
    c = append(c, a...)
    c = append(c, b...)

    fn(c, 0, a.Len()-1, a.Len()+b.Len()-1)
    if !sort.IsSorted(sort.IntSlice(c)) {
      t.Error("c is not sorted")
    }

    c = []int{3, 2}
    fn(c, 0, 0, 1)
    if !sort.IsSorted(sort.IntSlice(c)) {
      t.Error("small c is not sorted")
    }
  }
}

func TestMergeAB(t *testing.T) {
  a, b := getSorted()

  c := MergeAB(a, a.Len(), b, b.Len())
  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("c is not sorted")
  }

  a = []int{3}
  b = []int{2}
  c = MergeAB(a, a.Len(), b, b.Len())
  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("small c is not sorted")
  }
}

func TestMerge(t *testing.T) {
  a, b := getSorted()

  c := make([]int, 0, a.Len()+b.Len())
  c = append(c, a...)
  c = append(c, b...)

  Merge(c, 0, a.Len()-1, a.Len()+b.Len()-1)

  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("c is not sorted")
  }

  c = []int{3, 2}
  Merge(c, 0, 0, 1)
  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("small c is not sorted")
  }

  c = []int{1}
  Merge(c, 0, 0, 0)
  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("one-elem c is not sorted")
  }
}

func TestMergeLL(t *testing.T) {
  a := &ds.DumpHeadNode[int]{}
  b := &ds.DumpHeadNode[int]{}
  heada, headb := a, b

  perm1, perm2 := getSorted()

  perms := make([]int, len(perm1)+len(perm2))
  copy(perms, perm1)
  copy(perms[len(perm1):], perm2)
  sort.Sort(sort.IntSlice(perms))

  for _, v := range perm1 {
    a = a.Insert(ds.NewDumpHeadNode[int](v))
  }

  for _, v := range perm2 {
    b = b.Insert(ds.NewDumpHeadNode[int](v))
  }

  c := MergeLL(heada.Next(), headb.Next())
  for i, v := range perms {
    if v != c.Item() {
      t.Errorf("=== %dth values not equal: %d != %d\n", i, v, c.Item())
    }
    c = c.Next()
  }
}

func getSorted() (a, b sort.IntSlice) {
  a = make([]int, *alen)
  b = make([]int, *blen)
  for i := 0; i < *alen; i++ {
    a[i] = rand.Intn(MAX_N)
  }
  for i := 0; i < *blen; i++ {
    b[i] = rand.Intn(MAX_N)
  }

  sort.Sort(a)
  sort.Sort(b)

  return a, b
}