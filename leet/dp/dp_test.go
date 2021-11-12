package dp

import (
	"fmt"
	"testing"
)

func assertEqual(res, exp, msg interface{})  {
	switch v:= res.(type) {
	case int:
		if v != exp.(int) {
			fmt.Printf("res(%d) != exp(%d) params: %v\n", v, exp.(int), msg)
		}
	}
}
func TestNum11(t *testing.T)  {
	datas := []struct {
		arr []int
		res int
	}{
		{arr: []int{1,8,6,2,5,4,8,3,7}, res: 49},
	}
	for _, x:= range datas {
		assertEqual(maxArea(x.arr), x.res, x.arr)
	}
}
func TestNum375(t *testing.T)  {
	fmt.Println(getMoneyAmount(10)) //16
}
func TestNum629(t *testing.T)  {
 	fmt.Println(kInversePairs(5, 5))
}
