package handle_str

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum1446(t *testing.T)  {
	datas := []struct{
		s string
		ans int
	}{
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
