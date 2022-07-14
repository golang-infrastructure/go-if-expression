package example

import (
	"fmt"
	"github.com/CC11001100/go-ternary-expression/pkg/if_expression"
	"math/rand"
	"reflect"
)

func BoolExample() {

	// 对于布尔的确实有点多余，就是为了看起来整整齐齐...
	b := if_expression.ReturnBool(rand.Int()%2 == 0, false, true)
	fmt.Println(b) // output: true or false

	trueValue := true
	falseValue := false
	bp := if_expression.ReturnBoolPointer(rand.Int()%2 == 1, &trueValue, &falseValue)
	fmt.Println(reflect.TypeOf(bp)) // output: *bool

}
