package foursum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	N := len(nums)
	out := [][]int{}

	if N < 4 {
		return out
	}

	sort.Ints(nums)

	i := 0

	for i < N-3 {
		ni := nums[i]
		j := N - 1

		for j > i {
			nj := nums[j]
			k := i + 1
			l := j - 1

			for k < l {
				nk := nums[k]
				nl := nums[l]
				sum := ni + nj + nk + nl

				if sum == target {
					out = append(out, []int{ni, nj, nk, nl})
				}

				// move k until next value
				if sum <= target {
					for nums[k] == nk && k < l {
						k++
					}
				}

				// move l until next value
				if sum >= target {
					for nums[l] == nl && l > k {
						l--
					}
				}
			}

			// move j until next value
			for nums[j] == nj && j > i {
				j--
			}
		}

		// move i until next value
		for nums[i] == ni && i < N-3 {
			i++
		}
	}

	return out
}
