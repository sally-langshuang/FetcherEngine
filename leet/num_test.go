package leet

import (
	"fmt"
	"testing"
)

func TestNum10(t *testing.T)  {
	fmt.Println(isMatch("aaa", "a*a"))

}
func TestQuick(t *testing.T) {
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//Quick([]int{-2,1,3,4,7,8,100,22}, 2, 2, &wg)
	tt:= []struct {
		Nums []int
	}{{[]int{1, 4, 8, 22, 7, 3 , 100, -2}}}
	for _, x:= range tt {
		Quick(x.Nums, 0, len(x.Nums) - 1, nil)
		fmt.Println(x.Nums)
	}
}

func TestNum58(t *testing.T)  {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon"))
}
var i = 0
var j = 2
func TestNum8(t *testing.T)  {
	arr := []int{1, 2, 3}
	fmt.Println(arr[3:])
}

func TestNum5(t *testing.T)  {
	tests := []struct{s , exp string} {
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"aacabdkacaa", "aca"},
		{"a", "a"},
		{"ac", "a"},
	}
	for _, row := range tests {
		if res:= longestPalindrome(row.s); res != row.exp {
			t.Errorf("%v != %v", res, row.exp)
		} else {
			t.Logf("right!!!")
		}
	}
}
func TestNum4(t *testing.T)  {
	tests:= []struct{
		nums1, nums2 []int
		exp float64
	} {
		{[]int{1, 4},  []int{2, 3, 5, 6}, 3.5},
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
	}
	for _, row:= range tests {
		if res:= findMedianSortedArrays2(row.nums1, row.nums2); res != row.exp {
			t.Errorf("%v != %v", res, row.exp)
		} else {
			t.Logf("right!!!")
		}
	}
}
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
