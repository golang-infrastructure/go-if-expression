package if_expression

// ReturnComplex128
//  @Description: if实现的三元表达式，返回结果是complex128
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的complex128
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的complex128
//  @return complex128: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex128(boolExpression bool, trueReturnValue, falseReturnValue complex128) complex128 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnComplex128Slice
//  @Description: if实现的三元表达式，返回结果是[]complex128
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]complex128
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]complex128
//  @return []complex128: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex128Slice(boolExpression bool, trueReturnValue, falseReturnValue []complex128) []complex128 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnComplex128Pointer
//  @Description: if实现的三元表达式，返回结果是*complex128
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*complex128
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*complex128
//  @return *complex128: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex128Pointer(boolExpression bool, trueReturnValue, falseReturnValue *complex128) *complex128 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnComplex128PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*complex128
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*complex128
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*complex128
//  @return []*complex128: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnComplex128PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*complex128) []*complex128 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
