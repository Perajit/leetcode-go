package removeduplicates

func removeDuplicates(nums []int) int {
	N := len(nums)

	if N < 2 {
		return N
	}

	i := 0
	j := 1

	for i < N {
		// find index of next number
		for j < N && nums[i] == nums[j] {
			j++
		}

		// no next number
		if j > N-1 {
			break
		}

		// swap
		temp := nums[i+1]
		nums[i+1] = nums[j]
		nums[j] = temp

		i++
		j++
	}

	return i + 1
}

/*
[1,1,2]
i=0, j=1->2				=>	[1,2,1], i=1, j=3
i=1, j=3					=>	[1,2,1], i=1, j=3

[0,0,1,1,1,2,2,3,3,4]
i=0, j=1->2				=>	[0,1,0,1,1,2,2,3,3,4], i=1, j=3
i=1, j=3->4->5		=>	[0,1,2,1,1,0,2,3,3,4], i=2, j=6
i=2, j=6->7				=>	[0,1,2,3,1,0,2,1,3,4], i=3, j=8
i=3, j=8->9				=>	[0,1,2,3,4,0,2,1,3,1], i=4, j=10
i=4, j=10					=>	[0,1,2,3,4,0,2,1,3,1], i=4, j=10

[1,2,3]
i=0, j=1					=>	[1,2,3], i=1, j=2
i=1, j=2					=>	[1,2,3], i=2, j=3
i=2, j=3					=>	[1,2,3], i=2, j=3

[1,1]
i=0, j=1->2				=>	[1,1], i=0, j=2

[1,2,2]
i=0, j=1					=>	[1,2,2], i=1, j=2
i=1, j=2->3				=>	[1,2,2], i=1, j=2
*/
