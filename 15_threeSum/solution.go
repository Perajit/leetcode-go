package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	N := len(nums)
	out := [][]int{}

	if N < 3 {
		return out
	}

	sort.Ints(nums)

	i := 0

	for i < N-2 {
		ni := nums[i]
		j := i + 1
		k := N - 1

		for j < k {
			nj := nums[j]
			nk := nums[k]
			sum := nj + nk

			if sum == -ni {
				// found match, collect result
				out = append(out, []int{ni, nums[j], nums[k]})
			}

			if sum <= -ni {
				// found match or need higher value, move b forward until new number is found
				for nums[j] == nj && j < k {
					j++
				}
			}

			if sum >= -ni {
				// found match or need lower value, move c backward until new number is found
				for nums[k] == nk && k > j {
					k--
				}
			}
		}

		// avoid duplication, move a forward until new number is found
		for nums[i] == ni && i < N-2 {
			i++
		}
	}

	return out
}
