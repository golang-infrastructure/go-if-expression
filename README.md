# Golang的三元表达式实现

# 开发初衷

Golang中缺少三元表达式，就导致某些情况三元表达式一行就能搞定的事情到Golang里面就得写得很啰嗦， 这是无法忍受的，这个库就是借助大量自定义的if函数来实现类似三元表达式的功能。

## 引入依赖

非泛型版本：

```
go get -u github.com/golang-infrastructure/go-if-expression/v1
```

泛型版本： 

```text
go get -u github.com/golang-infrastructure/go-if-expression/v2
```

# V1 Example 

所有的API都在if_expression包下，根据三元表达式返回的不同类型有不同的名字：
![](.README_images/26e47025.png)

比如最常见的默认值场景：

```go
threadNum := 0
fmt.Printf("线程数： %d", if_expression.ReturnInt(threadNum == 0, 1, threadNum))
```



# V2 Example

```go
package main

import (
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression/v2"
)

func main() {

	r := if_expression.Return(true, "是", "否")
	fmt.Println(r)
	// Output:
	// 是

}
```

