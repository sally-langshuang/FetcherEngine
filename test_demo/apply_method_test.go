package test_demo

import (
"fmt"
"reflect"
"testing"

"github.com/agiledragon/gomonkey"
. "github.com/smartystreets/goconvey/convey"
)

type myType struct {
}

func (m *myType) logicFunc(a, b int) (int, error) {
	sum, err := m.NetWorkFunc(a, b)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (m *myType) NetWorkFunc(a, b int) (int, error) {
	if a < 0 && b < 0 {
		errmsg := "a<0 && b<0"
		return 0, fmt.Errorf("%v", errmsg)
	}

	return a + b, nil
}

func TestMockMethod(t *testing.T) {
	Convey("TestMockMethod", t, func() {
		var p *myType
		fmt.Printf("method num:%d\n", reflect.TypeOf(p).NumMethod())
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(p), "NetWorkFunc", func(_ *myType, a, b int) (int, error) {
			if a < 0 && b < 0 {
				errmsg := "a<0 && b<0"
				return 0, fmt.Errorf("%v", errmsg)
			}
			return a + b + 1, nil
		})
		defer p1.Reset()

		var m myType
		sum, err := m.logicFunc(10, 20)
		So(sum, ShouldEqual, 31)
		So(err, ShouldBeNil)
	})
}
