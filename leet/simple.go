package leet

//https://leetcode-cn.com/problems/excel-sheet-column-number/
func titleToNumber(columnTitle string) int {
	var res int
	for i :=0 ; i < len(columnTitle) ; i++{
		c := columnTitle[i]
		res = res * 26 +  int(c) - int('A') + 1
	}
	return res
}

//https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/
func pathInZigZagTree(label int) []int {
	arr := []int {label}
	for n:= label; n > 1;{
		n/=2
		arr = append(arr, n)
	}
	f := func (i int ) int{
		x := arr[i]
		return x
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = f(i)
	}
	return arr
}
