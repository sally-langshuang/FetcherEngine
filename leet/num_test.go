package leet

import (
	"testing"
)
func TestNum3(t *testing.T) {
	tests := []struct{
		s string
		exp int
	} {
		//{"abcabcbb", 3},
		//{"bbbbb", 1},
		//{"pwwkew", 3},
		//{"", 0},
		{"abba", 2},
	}
	for _, row := range tests {
		res := lengthOfLongestSubstring(row.s)
		if res != row.exp {
			t.Errorf("%v != %v", res, row.exp)
		} else {
			t.Logf("right")
		}

	}
}

func TestNum2(t *testing.T)  {
	tests := []struct {
		a, b, exp []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7,0,8}},
		{[]int{9,9,9,9,9,9,9},[]int{9,9,9,9},[]int{8,9,9,9,0,0,0,1}},
	}
	for _, row:= range tests {
		res := addTwoNumbers(initListNode(row.a), initListNode(row.b))
		exp := initListNode(row.exp)
		if ok, _:= exp.Equal(res); !ok {
			t.Errorf("%v != %v", res, exp)
		} else {
			t.Logf("%v == %v, right !!!", res, exp)
		}

	}

}