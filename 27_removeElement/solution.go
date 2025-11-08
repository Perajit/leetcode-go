package removeelement

import (
	"sort"
)

func removeElement(nums []int, val int) int {
	N := len(nums)

	if N == 0 {
		return 0
	}

	if N == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	sort.Ints(nums)

	i := 0
	j := N - 1

	for i < j {
		for i < j && nums[i] != val {
			i++
		}

		for i < j && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp

		i++
		j--
	}

	for i >= 0 && nums[i] == val {
		i--
	}

	return i + 1
}

/*
[2,2,3,3], 3
i=0->1->2, j=3->2			=>	[2,2,3,3], i=2, j=2

[2,2,3,3,4], 2
i=0, j=4							=>	[4,2,3,3,2], i=1, j=3
i=1, j=3							=>	[4,3,3,2,2], i=2, j=2
i=2. j=2							=>	[4,3,3,2,2], i=2, j=2

[0,1,2,3], 4
i=0->1->2->3, j=3			=>	[0,1,2,3], i=3, j=3

[1,2,2,3,3,4,5], 3
i=0->1->2->3, j=6			=>	[1,2,2,5,3,4,3], i=4, j=5
i=4, j=5							=>	[1,2,2,5,4,3,3], i=5, j=4

[3,3,3], 3
i=0->1->2->3, j=2			=>	[3,3,3], i=3, j=2

[2,2,2], 0
i=0->1->2, j=2				=>	[2,2,2], i=2, j=2
*/
