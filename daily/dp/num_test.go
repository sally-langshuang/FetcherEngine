package dp

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestA(t *testing.T)  {
	// Declaring u
	var u uint32

	// For loop
	for i:= 0; i < 315; i++ {

		// Function with AddUint32 method
		go func() {
			atomic.AddUint32(&u, 2)
		}()
	}

	// Prints loaded values address
	fmt.Println(atomic.LoadUint32(&u))
}
func TestNum630(t *testing.T)  {
	datas := []struct{
		counts [][]int
		ans int
	} {
		//{[][]int{{3, 2}, {4, 3}}, 0},
		{[][]int{{2,5},{2,19},{1,8},{1,3}}, 4},
	}
	for _, x := range datas {
		fmt.Println(scheduleCourse(x.counts))
	}
}
