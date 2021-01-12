package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m+n)

	i, j, k := 0, 0, 0
	for m > j && n > k && i < m+n {
		if nums1[j] <= nums2[k] {
			nums[i] = nums1[j]
			j++
		} else {
			nums[i] = nums2[k]
			k++
		}
		i++
	}

	for m > j {
		nums[i] = nums1[j]
		i++
		j++
	}
	for n > k {
		nums[i] = nums2[k]
		i++
		k++
	}
	copy(nums1, nums)
	return
}
