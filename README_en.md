# Golang's ternary expression implementation

[中文文档](./README.md) | [English Document](./README_en.md)

# 1. The original intention of development

The lack of ternary expressions in Golang has led to the fact that in some cases, things that can be done in one line with ternary expressions have to be written very verbosely in Golang.
This is unbearable, ~~this library uses a large number of custom if functions to achieve functions similar to ternary expressions~~, the latest version is implemented based on generics.

Before using this library:

```go
if a % 2 == 0 {
     return "even number"
} else {
     return "odd number"
}
```

After using this library:

```go
return if_expression.Return(a % 2 == 0, "even", "odd")
```

Compared:

``` diff
- if a % 2 == 0 {
- return "even number"
- } else {
- return "odd number"
- }
+ return if_expression.Return(a % 2 == 0, "even", "odd")
```

## 2. Introduce dependencies

go get install:

```text
go get -u github.com/golang-infrastructure/go-if-expression
```

If you don't want to add new dependencies, just copy the following code into your utils, the generic version is very concise:

```go
package if_expression

// Return
//
// @Description: The ternary expression implemented by if
// @param boolExpression: Boolean expression, finally returns a Boolean value
// @param trueReturnValue: The value returned when the return value of boolExpression is true
// @param falseReturnValue: The value returned when boolExpression returns false
// @return bool: The result of the ternary expression, which is one of trueReturnValue or falseReturnValue
func Return[T any](boolExpression bool, trueReturnValue, falseReturnValue T) T {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnByFunc
//
// @Description: The ternary expression implemented by if
// @param boolExpression: Boolean expression, finally returns a Boolean value
// @param trueReturnValue: When the return value of boolExpression is true, execute this function and return the value
// @param falseReturnValue: Execute this function and return value when boolExpression returns false
// @return bool: The result of the ternary expression, which is one of trueReturnValue or falseReturnValue
func ReturnByFunc[T any](boolExpression bool, trueFuncForReturnValue, falseFuncForReturnValue func() T) T {
    if boolExpression {
    	return trueFuncForReturnValue()
    } else {
    	return falseFuncForReturnValue()
    }
}
```

# 3. Example

For example, the most common default value scenario:

```go
threadNum := 0
fmt.Printf("Number of threads: %d", if_expression.Return(threadNum == 0, 1, threadNum))
```

Example of use:

```go
package main

import (
    "fmt"
    if_expression "github.com/golang-infrastructure/go-if-expression"
)

func main() {

    r := if_expression.Return(true, "Yes", "No")
    fmt.Println(r)
    //Output:
	// yes

}

```

Or use a function to return, only the function that is hit by the branch will be executed, but this method is not concise and is not recommended:

```go
package main

import (
"fmt"
if_expression "github.com/golang-infrastructure/go-if-expression"
)

func main() {

    r := if_expression. ReturnByFunc(true, func() string {
    	fmt.Println("True branch is executed")
    	return "yes"
    }, func() string {
    	fmt.Println("False branch is executed")
    	return "no"
    })
    fmt.Println(r)
    //Output:
    // True branch is executed
	// yes

}
```
