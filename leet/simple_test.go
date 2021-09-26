package leet

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T)  {
	data := []struct {
		S string; R int;
	}{
		//{"A", 1},
		{"AB", 28}, //1*26 + 2
	}
	for _, d:= range data {
		r := titleToNumber(d.S)
		if r != d.R {
			fmt.Printf("%s => %d shoud be %d\n", d.S, r, d.R)
		}
	}

}


func Test2(t *testing.T) {
	data := []struct{
		Label int; R []int
	}{
		{14, []int{1,3,4,14}},
		{26, []int{1,2,6,10,26}},
	}
	for _, d := range data{
		r := pathInZigZagTree(d.Label)
		if ! func(a , b []int) (y bool) {
			if len(a) == len(b) {
				y = true
				for i, ae := range a {
					if b[i] != ae {
						y = false
						return
					}
				}
			}
			return
		}( r , d.R) {
			t.Errorf("label: %dreal: %d expected: %d", d.Label, r, d.R)
		}
	}
}