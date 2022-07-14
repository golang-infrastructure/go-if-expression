package if_expression

// ReturnInt16
//  @Description: if实现的三元表达式，返回结果是int16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的int16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的int16
//  @return int16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt16(boolExpression bool, trueReturnValue, falseReturnValue int16) int16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt16Slice
//  @Description: if实现的三元表达式，返回结果是[]int16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]int16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]int16
//  @return []int16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt16Slice(boolExpression bool, trueReturnValue, falseReturnValue []int16) []int16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt16Pointer
//  @Description: if实现的三元表达式，返回结果是*int16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*int16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*int16
//  @return *int16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt16Pointer(boolExpression bool, trueReturnValue, falseReturnValue *int16) *int16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInt16PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*int16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*int16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*int16
//  @return []*int16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnInt16PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*int16) []*int16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
