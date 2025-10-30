package twosum

func twoSum(nums []int, target int) []int {
	indexLookup := make(map[int]int)

	for i, n := range nums {
		compliment := target - n
		j, ok := indexLookup[compliment]

		if ok {
			return []int{i, j}
		}

		indexLookup[n] = i
	}

	return []int{}
}
