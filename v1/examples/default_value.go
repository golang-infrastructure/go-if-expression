package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-if-expression/v1/if_expression"
)

func main() {

	threadNum := 0
	fmt.Printf("线程数： %d", if_expression.ReturnInt(threadNum == 0, 1, threadNum))

}
