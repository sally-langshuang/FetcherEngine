package leet

func twoSum(nums []int, target int) []int {
	peers :=  make(map[int]int)
	for idx, n := range nums {
		if peerIdx, ok := peers[n]; ok {
			return []int {peerIdx, idx}
		}
		peers[target - n]  = idx
	}
	return []int{}
}