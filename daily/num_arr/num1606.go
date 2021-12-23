package num_arr

type PQueue struct {
	arr  []interface{}
	Less func(i, j interface{}) bool
}

func (q *PQueue) Len() int {
	return len((*q).arr)
}
func (q *PQueue) Peek() interface{} {
	return (*q).arr[0]
}
func (q *PQueue) Pop() interface{} {
	v := q.Peek()
	(*q).arr[0] = (*q).arr[q.Len()-1]
	(*q).arr = (*q).arr[:q.Len()-1]
	q.Down(0)
	return v
}
func (q *PQueue) Push(v interface{}) {
	(*q).arr = append((*q).arr, v)
	q.Up(len(q.arr) - 1)
}
func (q *PQueue) Down(i int) {
	for {
		j := 2*i + 1
		if j >= q.Len() {
			break
		}
		if j2 := 2*i + 2; j2 < q.Len() && q.Less((*q).arr[j2], (*q).arr[j]) {
			j = j2
		}
		if q.Less((*q).arr[i], (*q).arr[j]) {
			break
		}
		(*q).arr[i], (*q).arr[j] = (*q).arr[j], (*q).arr[i]
		i = j
	}
}
func (q *PQueue) Up(i int) {
	for i != 0 {
		j := (i - 1) / 2
		if q.Less((*q).arr[j], (*q).arr[i]) {
			break
		}
		(*q).arr[i], (*q).arr[j] = (*q).arr[j], (*q).arr[i]
		i = j
	}
}

func busiestServers(k int, arrival []int, load []int) []int {
	count := make([]int, k)
	free := PQueue{arr: []interface{}{}, Less: func(i, j interface{}) bool {
		return i.(int) <= j.(int)
	}} // id
	busy := PQueue{arr: []interface{}{}, Less: func(i, j interface{}) bool {
		return i.([]int)[0] <= j.([]int)[0]
	}} //0 endTIme
	for i := 0; i < k; i++ {
		free.Push(i)
	}
	for i := range arrival {
		for busy.Len() > 0 && busy.Peek().([]int)[0] <= arrival[i] {
			n := busy.Pop().([]int)
			id := n[1]
			free.Push(i + ((id-i)%k+k)%k)
		}
		if free.Len() == 0 {
			continue
		}
		id := free.Pop().(int) % k
		busy.Push([]int{arrival[i] + load[i], id})
		count[id]++
	}
	maxCount, ans := 0, []int{}
	for i := range count {
		if count[i] > maxCount {
			maxCount = count[i]
			ans = []int{i}
		} else if count[i] == maxCount {
			ans = append(ans, i)
		}
	}
	return ans
}
