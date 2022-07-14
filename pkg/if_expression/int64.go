package if_expression

// ReturnInt64
//  @Description: if实现的三元表达式，返回结果是int64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的int64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的int64
//  @return int64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt64(boolExpression bool, trueReturnValue, falseReturnValue int64) int64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt64Slice
//  @Description: if实现的三元表达式，返回结果是[]int64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]int64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]int64
//  @return []int64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt64Slice(boolExpression bool, trueReturnValue, falseReturnValue []int64) []int64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt64Pointer
//  @Description: if实现的三元表达式，返回结果是*int64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*int64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*int64
//  @return *int64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt64Pointer(boolExpression bool, trueReturnValue, falseReturnValue *int64) *int64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt64PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*int64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*int64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*int64
//  @return []*int64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt64PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*int64) []*int64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
