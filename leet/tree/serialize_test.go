package tree

import (
	"fmt"
	"testing"
)

func TestPreorder(t *testing.T)  {
	datas := []struct{
		svc TreeSerialize
		str string
	}{
		{pre, "1,2,#,4,#,#,3,#,#"},
		{post, "#,#,#,4,2,#,#,3,1"},
		{level, "1,2,3,#,4,#,#,#,#"},
	}
	for _, x := range datas {
		svc := x.svc
		tree := svc.Deserialize(x.str)
		tree.Print()
		s := svc.Serialize(tree)
		if  s != x.str {
			t.Errorf("%s != %s", s, x.str)
		}

		fmt.Println(svc.Serialize(tree))
	}
}
