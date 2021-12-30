package num_arr

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}
	sort.Ints(hand)
	cnt := map[int]int{}
	for _, x := range hand {
		cnt[x]++
	}
	for i := 0; i < len(hand); i++ {
		if cnt[hand[i]] <= 0 {
			continue
		}
		for j := 0; j < groupSize; j++ {
			if val := hand[i] + j; cnt[val] <= 0 {
				return false
			} else {
				cnt[val]--
			}
		}
	}
	return true
}
func isNStraightHand0(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}
	sort.Ints(hand)
	heads := map[int]map[int]int{}
	for _, x := range hand {
		follow, ok := heads[x-1]
		if !ok {
			if heads[x] == nil {
				heads[x] = map[int]int{groupSize - 1: 1}
			} else {
				heads[x][groupSize-1]++
			}
			continue
		}
		var tailsNum int
		for k := range follow {
			tailsNum = k
			break
		}
		follow[tailsNum]--
		if follow[tailsNum] == 0 {
			delete(follow, tailsNum)
			if len(follow) == 0 {
				delete(heads, x-1)
			}
		}
		if tailsNum > 1 {
			if heads[x] == nil {
				heads[x] = map[int]int{}
			}
			heads[x][tailsNum-1]++
		}
	}
	return len(heads) == 0
}
