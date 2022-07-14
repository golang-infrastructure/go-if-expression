package if_expression

// ReturnInt
//  @Description: if实现的三元表达式，返回结果是int
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的int
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的int
//  @return int: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt(boolExpression bool, trueReturnValue, falseReturnValue int) int {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnIntSlice
//  @Description: if实现的三元表达式，返回结果是[]int
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]int
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]int
//  @return []int: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnIntSlice(boolExpression bool, trueReturnValue, falseReturnValue []int) []int {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnIntPointer
//  @Description: if实现的三元表达式，返回结果是*int
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*int
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*int
//  @return *int: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnIntPointer(boolExpression bool, trueReturnValue, falseReturnValue *int) *int {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnIntPointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*int
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*int
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*int
//  @return []*int: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnIntPointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*int) []*int {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}