package dp

import (
	"container/heap"
	"fmt"
	"sort"
)
func isAsc(a, b[]int) bool {
	if a[1] < b[1] {
		return true
	}
	if a[1] == b[1] && a[0] <= b[0] {
		return true
	}
	return false
}

func sort1(courses [][]int)  {
	if len(courses) <= 1 {
		return
	}
	i, j:= 0, len(courses) - 1
	for i < j {
		for isAsc(courses[i], courses[j]) {
			j--
			if j == i {
				goto out
			}
		}
		courses[i], courses[j] = courses[j], courses[i]
		for isAsc(courses[i], courses[j]){
			i++
			if j == i {
				goto out
			}
		}
		courses[i], courses[j] = courses[j], courses[i]
	}
	out:
	sort1(courses[:i])
	sort1(courses[i + 1:])
	return
}
type arr []int

type MyHeap struct {
	arr
}

func (h *MyHeap) up(i int)  {
	for x := i; x > 0; {
		up := (x - 1) / 2
		if h.arr[up] < h.arr[x] {
			h.arr[up], h.arr[x] = h.arr[x], h.arr[up]
		} else {
			break
		}
		x = up
	}
}

func (h *MyHeap) Push(v int)  {
	h.arr = append(h.arr, v)
	h.up(len(h.arr) - 1)
}

func (h *MyHeap) down(i0 int) bool {
	i := i0
	for {
		j := i * 2 + 1
		if j >= len(h.arr) {
			break
		}
		if j2 := j + 1; j2 < len(h.arr) && h.arr[j2] > h.arr[j] {
			j = j2
		}
		if h.arr[j] <= h.arr[i] {
			break
		}
		h.arr[j], h.arr[i] = h.arr[i], h.arr[j]
		i = j
	}
	return i > i0

}
func (h *MyHeap) Fix(i int)  {
	if !h.down(i) {
		h.up(i)
	}
}

func scheduleCourse0(courses [][]int) int {
	sort1(courses)
	var offset, maxCount int
	h := MyHeap{arr{}}
	for i := range courses {
		if courses[i][0] + offset <= courses[i][1] {
			offset += courses[i][0]
			h.Push(courses[i][0])
			maxCount++
		} else if len(h.arr) > 0 {
			if h.arr[0] > courses[i][0] {
				offset += - h.arr[0] + courses[i][0]
				h.arr[0] = courses[i][0]
				h.Fix(0)
			}
		}
	}
	return len(h.arr)
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &Heap{}
	total := 0 // 优先队列中所有课程的总时间
	for _, course := range courses {
		if h.Len() == 191 {
			fmt.Println("debug")
		}
		if t := course[0]; total+t <= course[1] {
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}

