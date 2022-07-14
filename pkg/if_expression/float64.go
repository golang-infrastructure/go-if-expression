package if_expression

// ReturnFloat64
//  @Description: if实现的三元表达式，返回结果是float64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的float64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的float64
//  @return float64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat64(boolExpression bool, trueReturnValue, falseReturnValue float64) float64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnFloat64Slice
//  @Description: if实现的三元表达式，返回结果是[]float64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]float64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]float64
//  @return []float64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat64Slice(boolExpression bool, trueReturnValue, falseReturnValue []float64) []float64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnFloat64Pointer
//  @Description: if实现的三元表达式，返回结果是*float64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*float64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*float64
//  @return *float64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat64Pointer(boolExpression bool, trueReturnValue, falseReturnValue *float64) *float64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}


// ReturnFloat64PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*float64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*float64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*float64
//  @return []*float64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat64PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*float64) []*float64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
