package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := m - 1
	ptr2 := n - 1
	for ptr1 >= 0 || ptr2 >= 0 {
		if ptr1 < 0 {
			nums1[ptr1+ptr2+1] = nums2[ptr2]
			ptr2--
		} else if ptr2 < 0 {
			nums1[ptr1+ptr2+1] = nums1[ptr1]
			ptr1--
		} else if nums1[ptr1] > nums2[ptr2] {
			nums1[ptr1+ptr2+1] = nums1[ptr1]
			ptr1--
		} else {
			nums1[ptr1+ptr2+1] = nums2[ptr2]
			ptr2--
		}
	}
}
