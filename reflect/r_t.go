package testdemo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAA(t *testing.T) {
	for _, v := range []interface{}{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println("a", v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}
}

func TestB(t *testing.T) {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}

type child struct {
	Name float64
}

type S struct {
	F    float32 `species:"gopher" color:"blue"`
	Name string
	Food []int
	child
	s struct{ age int }
}

func (s S) Me(args ...interface{}) {
	fmt.Println("call==>", args)
}

func Me(a int, args ...interface{}) int {
	return a + 1
}

type ChanDir int

const (
	RecvDir  ChanDir             = 1 << iota // <-chan
	SendDir                                  // chan<-
	BothDir  = RecvDir | SendDir             // chan
	ForthDir = iota
	Fifth
)

func TestC(t *testing.T) {
	//	Int*, Uint*, Float*, Complex*: Bits
	//	Array: Elem, Len
	//	Chan: ChanDir, Elem
	//	Func: In, NumIn, Out, NumOut, IsVariadic.
	//	Map: Key, Elem
	//	Ptr: Elem
	//	Slice: Elem
	//	Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField
	var ty reflect.Type
	ty = reflect.TypeOf(2)
	fmt.Println(ty.Bits())

	ch := make(chan<- int)
	ty = reflect.TypeOf(ch)

	if ty.Kind() == reflect.Chan {
		fmt.Println(reflect.TypeOf(ch).ChanDir())
	}

	l := []int{1, 2}
	ty = reflect.TypeOf(l)
	if ty.Kind() == reflect.Array {
		fmt.Println(reflect.ValueOf(l).Len())
		fmt.Println(ty.Elem())
	}

	ty = reflect.TypeOf(Me)
	if ty.Kind() == reflect.Func {
		fmt.Println(ty.IsVariadic())
		for i := 0; i < ty.NumIn(); i++ {
			fmt.Println(ty.In(i))
		}
		for i := 0; i < ty.NumOut(); i++ {
			fmt.Println(ty.Out(i))
		}
	}
	v := S{}
	ty = reflect.TypeOf(v)
	if ty.Kind() == reflect.Struct {
		for i := 0; i < ty.NumField(); i++ {
			field := ty.Field(i)
			fmt.Printf("field = %[1]v, Name= %s, tag=%v, offset=%v, index=%v ,Anonymous=%v PkgPath=%s, type=%v\n", field, field.Name, field.Tag, field.Offset, field.Index, field.Anonymous, field.PkgPath, field.Type)
		}
		idx := []int{3, 0}
		fmt.Println(ty.FieldByIndex(idx))
		fmt.Println(ty.FieldByNameFunc(func(s string) bool { return s == "Name" }))
		fmt.Println(ty.NumMethod())
		for j := 0; j < ty.NumMethod(); j++ {
			method := ty.Method(j)
			fmt.Printf("func=%v, index=%d, name=%s, pkg=%v, type=%v\n", method.Func, method.Index, method.Name, method.PkgPath, method.Type)
			fmt.Println("------s-----")
			fmt.Println(method.Type.In(1).Kind(), method.Type.In(1).Elem())
			fmt.Println("------s-----")
			ps := reflect.New(method.Type.In(1))
			method.Func.Call([]reflect.Value{reflect.ValueOf(v), ps})
		}
	}

}
