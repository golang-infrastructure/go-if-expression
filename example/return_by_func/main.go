package main

import (
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
)

func main() {

	r := if_expression.ReturnByFunc(true, func() string {
		fmt.Println("True分支被执行了")
		return "是"
	}, func() string {
		fmt.Println("False分支被执行了")
		return "否"
	})
	fmt.Println(r)
	// Output:
	// True分支被执行了
	// 是

}
