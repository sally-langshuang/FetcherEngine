package leet

//https://leetcode-cn.com/problems/network-delay-time/
//	data := []struct {
//		Times   [][]int;
//		N, K, R int
//	}{
//		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
//		{[][]int{{1, 2, 1}}, 2, 1, 1},
//		{[][]int{{1, 2, 1}}, 2, 2, -1},
//	}

func makeArr(m, n int) [][]int{
	res := make([][]int, m)
	for i, _ := range res {
		res[i] = make([]int, n)
	}
	return res
}

func F743(times [][]int, n, k int ) int {
	arr := makeArr(n, n)
	for _, t := range times {
		arr[t[0] - 1][t[1] - 1] = t[2]
	}
	return -1
}