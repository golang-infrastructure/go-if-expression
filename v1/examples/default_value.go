package main

import (
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression/v1"
)

func main() {

	threadNum := 0
	fmt.Printf("线程数： %d", if_expression.ReturnInt(threadNum == 0, 1, threadNum))

}
