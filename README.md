# Golang的三元表达式实现

# 一、开发初衷

Golang中缺少三元表达式，就导致某些情况三元表达式一行就能搞定的事情到Golang里面就得写得很啰嗦，
这是无法忍受的，~~这个库就是借助大量自定义的if函数来实现类似三元表达式的功能~~，最新版本是基于泛型实现的。

使用此库之前：

```go
if a % 2 == 0 {
    return "偶数"
} else {
    return "奇数"
}
```

使用此库之后：

```go
return if_expression.Return(a % 2 == 0, "偶数", "奇数")
```

对比：

```diff
- if a % 2 == 0 {
- 	return "偶数"
- } else {
- 	return "奇数"
- }
+ return if_expression.Return(a % 2 == 0, "偶数", "奇数")
```

## 二、引入依赖

go get安装：

```text
go get -u github.com/golang-infrastructure/go-if-expression
```

如果你不想增加新的依赖，直接拷贝下面这段代码到你的utils中，泛型的实现版本非常简洁：

```go
package if_expression

// Return
//
//	@Description: if实现的三元表达式
//	@param boolExpression: 布尔表达式，最终返回一个布尔值
//	@param trueReturnValue: 当boolExpression返回值为true的时候返回的值
//	@param falseReturnValue: 当boolExpression返回值为false的时候返回的值
//	@return bool: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func Return[T any](boolExpression bool, trueReturnValue, falseReturnValue T) T {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
```

# 三、 Example

比如最常见的默认值场景：

```go
threadNum := 0
fmt.Printf("线程数： %d", if_expression.Return(threadNum == 0, 1, threadNum))
```

使用的例子：

```go
package main

import (
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
)

func main() {

	r := if_expression.Return(true, "是", "否")
	fmt.Println(r)
	// Output:
	// 是

}

```

