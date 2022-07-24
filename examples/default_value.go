package main

import (
	"fmt"
	"github.com/CC11001100/go-ternary-expression/pkg/if_expression"
)

func main() {

	threadNum := 0
	fmt.Printf("线程数： %d", if_expression.ReturnInt(threadNum == 0, 1, threadNum))

}
