package leet

import (
	"sync"
)

func lengthOfLastWord(s string) (res int) {
	var l int
	for _, c:= range s {
		if c != ' ' {
			l += 1
		} else {
			if l != 0 {
				res, l = l, 0
			}
		}
	}
	if l != 0 {
		res = l
	}
	return res
}


func quickChildren(nums []int, lo, hi, mid int, wg *sync.WaitGroup)  {
	Quick(nums, lo, mid - 1, wg)
	Quick(nums, mid + 1, hi, wg)
}

func exchange(nums []int, h, l int)  {
	var temp int
	temp = nums[h]
	nums[h] = nums[l]
	nums[l] = temp
}

func inner(nums []int, l, h *int, lo, hi int, w *sync.WaitGroup, op func()) (end bool) {
	for {
		if *l == *h || nums[*l] > nums[*h] {
			break
		}
		op()
	}

	if *h == *l {
		quickChildren(nums, lo, hi, *l, w)
		w.Wait()
		end = true
		return
	}

	exchange(nums, *h, *l)
	return
}

func Quick(nums []int, lo, hi int, wg *sync.WaitGroup)  {
	if wg != nil {
		defer func() {
			wg.Done()
		}()
	}
	if lo >= hi {
		return
	}
	w := sync.WaitGroup{}
	w.Add(2)

	for l, h := lo, hi;; {
		end := inner(nums, &l, &h, lo, hi, &w, func() {h--})
		if end {
			return
		}
		end = inner(nums, &l, &h, lo, hi, &w, func() {l++})
		if end {
			return
		}

		//for {
		//	if l == h || nums[l] > nums[h] {
		//		break
		//	}
		//	h--
		//}
		//if h == l {
		//	quickChildren(nums, lo, hi, l, &w)
		//	w.Wait()
		//	return
		//}
		//exchange(nums, h, l)
		//for {
		//	if l == h || nums[l] > nums[h] {
		//		break
		//	}
		//	l++
		//}
		//if h == l {
		//	quickChildren(nums, lo, hi, l, &w)
		//	w.Wait()
		//	return
		//}
		//exchange(nums, h, l)
	}
}