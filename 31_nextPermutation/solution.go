package nextpermutation

import (
	"math"
)

func nextPermutation(nums []int) {
	N := len(nums)

	if N < 2 {
		return
	}

	// find first index of asc order from last
	n := math.MinInt
	i := N - 1

	for ; i >= 0; i-- {
		if nums[i] < n {
			break
		}

		n = nums[i]
	}

	// for last permutation: next = first permutation
	if i < 0 {
		reverse(&nums, i+1, N)
		return
	}

	// find smallest value greater than nums[i] from last
	j := N - 1
	for ; j > i; j-- {
		if nums[j] > nums[i] {
			break
		}
	}

	// swap ni <-> nj and sort rhs asc (reverse)
	n = nums[i]
	nums[i] = nums[j]
	nums[j] = n
	reverse(&nums, i+1, N)
}

func reverse(nums *[]int, start int, end int) {
	bound := (end - start) / 2
	offset := 0

	for ; offset < bound; offset++ {
		temp := (*nums)[start+offset]
		(*nums)[start+offset] = (*nums)[end-1-offset]
		(*nums)[end-1-offset] = temp
	}
}
