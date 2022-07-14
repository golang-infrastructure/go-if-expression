package if_expression

// ReturnInt8
//  @Description: if实现的三元表达式，返回结果是int8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的int8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的int8
//  @return int8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt8(boolExpression bool, trueReturnValue, falseReturnValue int8) int8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt8Slice
//  @Description: if实现的三元表达式，返回结果是[]int8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]int8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]int8
//  @return []int8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt8Slice(boolExpression bool, trueReturnValue, falseReturnValue []int8) []int8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt8Pointer
//  @Description: if实现的三元表达式，返回结果是*int8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*int8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*int8
//  @return *int8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt8Pointer(boolExpression bool, trueReturnValue, falseReturnValue *int8) *int8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt8PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*int8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*int8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*int8
//  @return []*int8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt8PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*int8) []*int8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
