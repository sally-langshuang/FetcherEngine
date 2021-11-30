package leet

func findMedianSortedArrays(nums1 []int, nums2 []int) (res float64) {
	h1, h2, t1, t2 := 0, 0, len(nums1), len(nums2)
	for ; h1 < t1 || h2 < t2; {
		if h1 < t1 && (h2 >= t2 || nums2[h2] > nums1[h1])  {
			res = float64(nums1[h1])
			h1++
		} else if h2 < t2{
			res = float64(nums2[h2])
			h2++
		}
		if h1 < t1 && (h2 >= t2 || nums2[t2 - 1] < nums1[t1 - 1]) {
			res =  res/2 + float64(nums1[t1 - 1])/ 2
			t1--
		} else if h2 < t2{
			res = res/2 + float64(nums2[t2 -1]) / 2
			t2--
		}
	}
	return
}
func findMedianSortedArrays2(nums1 []int, nums2 []int) (res float64) {
    totalLen := len(nums1) + len(nums2)
    k := totalLen / 2
    if totalLen % 2 == 1 {
    	return float64(getKthElement(nums1, nums2, k + 1))
	}
	return float64(getKthElement(nums1, nums2, k)) / 2 + float64(getKthElement(nums1, nums2, k + 1)) / 2
}

func getKthElement(nums1 , nums2 []int, k int) int {
	if k > len(nums1) + len(nums2) { panic("k wront")}
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[k+idx2-1]
		}
		if idx2 == len(nums2) {
			return nums1[k+idx1-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		newIdx1 := min(idx1 + k / 2, len(nums1))
		newIdx2 := min(idx2 + k / 2, len(nums2))
		if nums1[newIdx1 - 1] < nums2[newIdx2 - 1] {
			k -= newIdx1 - idx1
			idx1 = newIdx1
		} else {
			k -= newIdx2 - idx2
			idx2 = newIdx2
		}
	}
	panic("not expected")
}
