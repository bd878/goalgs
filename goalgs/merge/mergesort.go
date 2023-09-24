package main

func Mergesort(nums []int) {
	if len(nums) > 1 {
		mergesort(nums, 0, len(nums)-1)
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