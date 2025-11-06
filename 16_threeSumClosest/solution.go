package threesumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	N := len(nums)
	closestDiff := math.MaxInt

	sort.Ints(nums)

	for i := 0; i < N-2; i++ {
		j := i + 1
		k := N - 1

		for j < k {
			diff := nums[i] + nums[j] + nums[k] - target

			if diff == 0 {
				return target + diff
			}

			if diff < 0 {
				j++
			} else {
				k--
			}

			if abs(diff) < abs(closestDiff) {
				closestDiff = diff
			}
		}
	}

	return target + closestDiff
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
