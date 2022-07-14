package if_expression

// ReturnFloat32
//  @Description: if实现的三元表达式，返回结果是float32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的float32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的float32
//  @return float32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat32(boolExpression bool, trueReturnValue, falseReturnValue float32) float32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnFloat32Slice
//  @Description: if实现的三元表达式，返回结果是[]float32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]float32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]float32
//  @return []float32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat32Slice(boolExpression bool, trueReturnValue, falseReturnValue []float32) []float32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnFloat32Pointer
//  @Description: if实现的三元表达式，返回结果是*float32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*float32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*float32
//  @return *float32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat32Pointer(boolExpression bool, trueReturnValue, falseReturnValue *float32) *float32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnFloat32PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*float32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*float32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*float32
//  @return []*float32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnFloat32PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*float32) []*float32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
