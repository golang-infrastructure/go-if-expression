package if_expression

// ReturnUInt8
//  @Description: if实现的三元表达式，返回结果是uint8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的uint8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的uint8
//  @return uint8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt8(boolExpression bool, trueReturnValue, falseReturnValue uint8) uint8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt8Slice
//  @Description: if实现的三元表达式，返回结果是[]uint8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]uint8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]uint8
//  @return []uint8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt8Slice(boolExpression bool, trueReturnValue, falseReturnValue []uint8) []uint8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt8Pointer
//  @Description: if实现的三元表达式，返回结果是*uint8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*uint8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*uint8
//  @return *uint8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt8Pointer(boolExpression bool, trueReturnValue, falseReturnValue *uint8) *uint8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt8PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*uint8
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*uint8
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*uint8
//  @return []*uint8: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt8PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*uint8) []*uint8 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
