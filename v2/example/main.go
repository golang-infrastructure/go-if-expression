package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-if-expression/v2/if_expression"
)

func main() {

	r := if_expression.Return(true, "是", "否")
	fmt.Println(r)
	// Output:
	// 是

}
