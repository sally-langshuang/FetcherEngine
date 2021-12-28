package cache

import (
	"encoding/binary"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"testing"
)

func TestBloom(t *testing.T) {
	fmt.Println(^uint(0))
	filter := bloom.NewWithEstimates(100, 0.01)
	filter.Add([]byte("Love"))
	filter.Add([]byte("banan"))
	if filter.Test([]byte("Love")) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println(filter.Test([]byte("Lov")))
	fmt.Println(filter.Test([]byte("banan")))
	fmt.Println(filter.Test([]byte("Banan")))
	i := uint32(100)
	n1 := make([]byte, 4)
	binary.BigEndian.PutUint32(n1, i)
	filter.Add(n1)
}
