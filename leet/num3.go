package leet

func lengthOfLongestSubstring(s string) int {
	head, tail, max := 0, 0, 0
	elementLastIdx := make(map[byte]int)
	for ; tail < len(s); tail++ {
		if i, ok := elementLastIdx[s[tail]]; ok && i + 1 >head {
			head = i + 1
		}
		if tail - head + 1 > max {
			max = tail - head + 1
		}
		elementLastIdx[s[tail]] = tail
		//fmt.Println(head, tail , max, elementLastIdx)
	}
	return max
}