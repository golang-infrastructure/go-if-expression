package if_expression

// ReturnComplex64
//  @Description: if实现的三元表达式，返回结果是complex64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的complex64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的complex64
//  @return complex64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex64(boolExpression bool, trueReturnValue, falseReturnValue complex64) complex64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnComplex64Slice
//  @Description: if实现的三元表达式，返回结果是[]complex64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]complex64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]complex64
//  @return []complex64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex64Slice(boolExpression bool, trueReturnValue, falseReturnValue []complex64) []complex64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnComplex64Pointer
//  @Description: if实现的三元表达式，返回结果是*complex64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*complex64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*complex64
//  @return *complex64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex64Pointer(boolExpression bool, trueReturnValue, falseReturnValue *complex64) *complex64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnComplex64PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*complex64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*complex64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*complex64
//  @return []*complex64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex64PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*complex64) []*complex64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
