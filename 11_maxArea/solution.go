package maxarea

func maxArea(height []int) int {
	N := len(height)
	i := 0
	j := N - 1
	maxAmount := 0

	for i < j {
		hi := height[i]
		hj := height[j]
		maxHeight := min(hi, hj)
		amount := maxHeight * (j - i)

		if amount > maxAmount {
			maxAmount = amount
		}

		if hi < hj {
			i++
		} else {
			j--
		}
	}

	return maxAmount
}
