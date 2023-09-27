package main

func Mergesort(nums []int) {
	if len(nums) > 1 {
		mergesort(nums, 0, len(nums)-1)
	}
}

func MergesortUp(nums []int) {
  if len(nums) > 1 {
    mergesortUp(nums, 0, len(nums)-1)
  }
}

func mergesort(nums []int, from, to int) {
	if from < to {
		half := (from + to) / 2
		mergesort(nums, from, half)
		mergesort(nums, half + 1, to)
		Merge(nums, from, half, to)
	}
}

func mergesortUp(nums []int, from, to int) {
  for m := 1; m <= to-from; m = m*2 {
    // merge to adjacent lists
    for i := 0; i <= to-from; i = i+m*2 {
      ll, mm, rr := i, min(i+m-1, to), min(i+m*2-1, to)
      Merge(nums, ll, mm, rr)
    }
  }
}