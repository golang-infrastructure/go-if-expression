package if_expression

// ReturnInt32
//  @Description: if实现的三元表达式，返回结果是int32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的int32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的int32
//  @return int32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt32(boolExpression bool, trueReturnValue, falseReturnValue int32) int32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt32Slice
//  @Description: if实现的三元表达式，返回结果是[]int32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]int32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]int32
//  @return []int32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt32Slice(boolExpression bool, trueReturnValue, falseReturnValue []int32) []int32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt32Pointer
//  @Description: if实现的三元表达式，返回结果是*int32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*int32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*int32
//  @return *int32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt32Pointer(boolExpression bool, trueReturnValue, falseReturnValue *int32) *int32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
// ReturnInt32PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*int32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*int32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*int32
//  @return []*int32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt32PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*int32) []*int32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}