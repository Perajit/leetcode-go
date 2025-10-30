package findmediansortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	L := l1 + l2

	if l1 == 0 {
		return findMedianFromArray(nums2)
	}

	if l2 == 0 {
		return findMedianFromArray(nums1)
	}

	end := L/2 + 1
	i := 0
	j := 0
	count := 0

	if L%2 > 0 {
		// odd
		latest := 0

		for count < end {
			if i < l1 && (j >= l2 || nums1[i] <= nums2[j]) {
				latest = nums1[i]
				i++
			} else {
				latest = nums2[j]
				j++
			}

			count++
		}

		return float64(latest)
	} else {
		// even
		latest1 := 0
		latest2 := 0

		for count < end {
			latest2 = latest1

			if i < l1 && (j >= l2 || nums1[i] <= nums2[j]) {
				latest1 = nums1[i]
				i++
			} else {
				latest1 = nums2[j]
				j++
			}

			count++
		}

		return float64(latest1+latest2) / 2
	}
}

func findMedianFromArray(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		return 0
	} else if l == 1 {
		return float64(nums[0])
	} else if l%2 > 0 {
		// odd
		return float64(nums[l/2])
	} else {
		// even
		mi := l / 2
		return float64(nums[mi]+nums[mi-1]) / 2
	}
}
