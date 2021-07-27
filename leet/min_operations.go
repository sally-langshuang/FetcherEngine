package leet

import (
	"fmt"
)

type node struct {
	TargetIdx, ArrIdx int
	before    *node
}

//Sort.Search
func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

//https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/solution/de-dao-zi-xu-lie-de-zui-shao-cao-zuo-ci-hefgl/
func minOperations(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	d := []*node{}  //  d[i]表示 存在长度为i+1的idx子串，末尾值是d[i]
	for i, val := range arr {  // arr中不在target的元素可以不用考虑
		if idx, has := pos[val]; has {  //前面含有比idx小的子串长度， 更新tails
			p := Search(len(d), func(i int) bool {
				return d[i].TargetIdx >= idx
			})
			n := node{TargetIdx: idx, ArrIdx: i}
			// 遍历的新值idx不能放到p+1以后，因为p+1的最小元素大于idx
			if p < len(d) { // 最长子串要保证结尾元素尽量小，才能保证后遍历的元素能加入子串，后遍历的下标值可以插入到d[p]，因为d[p-1]是最大的小于idx的下标值，而p+1是结尾为idx的子串最大长度
				d[p] = &n
			} else {
				d = append(d, &n)  //此事最长子串的长度初出现更优解
			}
			if p-1 >= 0 {
				n.before = d[p-1]
			}
		}
	}
	subNodes := []node{}[:]
	for x := d[len(d)-1]; x !=nil ;x= x.before {
		subNodes = append([]node{*x}, subNodes... )
	}
	fmt.Println(target, arr)
	for _, n:= range subNodes {
		fmt.Printf("target[%d], arr[%d]: %d ==> ", n.TargetIdx,n.ArrIdx, target[n.TargetIdx])
	}
	return n - len(d)
}
