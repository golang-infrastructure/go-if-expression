package if_expression

import (
	"fmt"
	"testing"
)

func TestReturn(t *testing.T) {
	r := Return(true, "是", "否")
	fmt.Println(r)
}

func ExampleReturn() {
	r := Return(true, "是", "否")
	fmt.Println(r)
	// Output:
	// 是
}

//func TestMap(t *testing.T) {
//	m := map[string]interface{}{
//		"foo": "bar",
//	}
//	m["bar"] = "aaa"
//	//t.Log(m["bad"])    // nil
//	var v string
//	v = Return[string](m["bad"] == nil, m["bar"], "aaa")  //  m["bar"]的类型不对
//	t.Log(v)
//}
