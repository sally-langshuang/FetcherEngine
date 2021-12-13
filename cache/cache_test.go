package cache

import (
	"fmt"
	"testing"
)
var Users = map[int]interface{}{}

type User struct {
	ID int
	Name string
}

var fs = map[string]fi{
	"GetUser": GetUser,
}
var cache = map[string]interface{}{
	"GetUser": Users,
}

func updateRedis(name string) fi {
	function := fs[name]
	c := cache[name]
	return func (params ... interface{}) interface{}{

		// update
		user := function(params)
		c.(map[int]interface{})[params[0].(int)] = user
		return user
	}
}
type fi func(...interface{}) interface{}

func xf(a int, b string, c ... interface{}) bool {
	return false
}
func GetUser(params ...interface{}) interface{} {
	return User{params[0].([]interface{})[0].(int), params[0].([]interface{})[1].(string)}
}
type funcType1 = func(int) bool
type funcType2 = func(...interface{}) interface{}
func decorate2(finterface interface{}) interface{} {
	var res interface{}
	switch f := finterface.(type) {
	case funcType1:
		return func(id int) bool{
			fmt.Println("decorated...")
			res = f(id)
			return res.(bool)
		}
	default:
		return nil
	}

}

func K(p int) bool {
	if p > 0 {
		return true
	}
	return false
}
//	//y:= reflect.TypeOf(xf).In(0).Name()//.Elem().Name()
func TestCache(t *testing.T)  {
	fmt.Println(decorate2(K).(funcType1)(1))
	//var id int = 1
	//fmt.Println(updateRedis("GetUser")(id, "ls"))
	//fmt.Println(updateRedis("GetUser")(id, "ls2"))
	//fmt.Println(Users)
}
