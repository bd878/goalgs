package main

import "sort"

func Mergesort(nums sort.Interface) {
	if nums.Len() > 1 {
		mergesort(nums, 1, nums.Len())
	}
}

func mergesort(nums sort.Interface, from, to int) {
	if from < to {
		half := (from + to) / 2
		mergesort(nums, from, half)
		mergesort(nums, half + 1, to)
		// merge(nums, from, half, to)
	}
}