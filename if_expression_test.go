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
