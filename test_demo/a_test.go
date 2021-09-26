package test_demo

import (
	//"bou.ke/monkey"
	"fmt"
	"github.com/agiledragon/gomonkey"
	//"github.com/sally-langshuang/FetcherEngine/test_demo/infra/os-encap"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)


//func TestA(t *testing.T)  {
//	Convey("test has digit", t, func() {
//		Convey("for succ", func() {
//			guard := monkey.Patch(osencap.Exec, func(cmd string, args ...string) (string, error) {
//				return "outputExpect", nil
//			})
//			defer guard.Unpatch()
//			output, err := osencap.F("any", "any")
//			So(output, ShouldEqual, "outputExpect")
//			So(err, ShouldBeNil)
//		})
//	})
//}

func logicFunc(a,b int) (int, error) {
	sum, err := netWorkFunc(a, b)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func netWorkFunc(a,b int) (int,error){
	if a < 0 && b < 0 {
		errmsg := "a<0 && b<0" //gomonkey有bug，函数一定要有栈分配变量，不然mock不住
		return 0, fmt.Errorf("%v",errmsg)
	}

	return a+b, nil
}

func TestMockFunc(t *testing.T) {
	convey.Convey("TestMockFunc1", t, func() {
		var p1 = gomonkey.ApplyFunc(netWorkFunc, func(a, b int) (int, error) {
			fmt.Println("in mock function")
			return a+b, nil
		})
		defer p1.Reset()

		sum, err := logicFunc(10, 20)
		convey.So(sum, convey.ShouldEqual, 30)
		convey.So(err, convey.ShouldBeNil)
	})

}