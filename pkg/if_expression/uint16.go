package if_expression

// ReturnUInt16
//  @Description: if实现的三元表达式，返回结果是uint16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的uint16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的uint16
//  @return uint16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt16(boolExpression bool, trueReturnValue, falseReturnValue uint16) uint16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt16Slice
//  @Description: if实现的三元表达式，返回结果是[]uint16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]uint16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]uint16
//  @return []uint16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt16Slice(boolExpression bool, trueReturnValue, falseReturnValue []uint16) []uint16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt16Pointer
//  @Description: if实现的三元表达式，返回结果是*uint16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*uint16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*uint16
//  @return *uint16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt16Pointer(boolExpression bool, trueReturnValue, falseReturnValue *uint16) *uint16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt16PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*uint16
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*uint16
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*uint16
//  @return *uint16: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt16PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*uint16) []*uint16 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
