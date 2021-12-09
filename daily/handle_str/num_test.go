package handle_str

import (
	"fmt"
	"reflect"
	"testing"
)


func TestNum44(t *testing.T)  {
	datas := []struct{
		s, p string
		ans bool
	}{
		{"abcabczzzde","*abc???de*", true},
		{"zacabz","*a?b*", false},
		{"acdcb", "a*c?b", false},
		{"adceb", "*a*b", true},
		{"cb", "?a", false},
		{"aa", "*", true},
		{"aa", "a", false},
		{"", "*", true},
		{"", "", true},
	}
	for _, x := range datas{
		fmt.Println(isMatch(x.s, x.p) == x.ans)
	}
}
func TestNum43(t *testing.T)  {
	datas := []struct{
		num1, num2, ans string
	}{
		{"12", "3", "36"},
		{"9", "9", "81"},
		{"123", "456", "56088"},
		{"123", "0", "0"},
	}
	for _, x := range datas {
		fmt.Println(multiply(x.num1, x.num2) == x.ans)
	}
}
func TestNum1446(t *testing.T)  {
	datas := []struct{
		s string
		ans int
	}{
		{"j", 1},
		{"cc", 2},
		{"tourist", 1},
		{"hooraaaaaaaaaaay", 11},
		{"triplepillooooow", 5},
		{"abbcccddddeeeeedcba", 5},
		{"leetcode", 2},
	}
	for _, x := range datas {
		fmt.Println(reflect.DeepEqual(maxPower(x.s), x.ans))
	}
}
