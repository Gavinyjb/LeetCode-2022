package Problem_0088_MergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	m, n = m-1, n-1
	i := len(nums1) - 1
	for ; i >= 0 && m >= 0 && n >= 0; i-- {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
	for m >= 0 {
		nums1[i] = nums1[m]
		m--
		i--
	}
}
