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