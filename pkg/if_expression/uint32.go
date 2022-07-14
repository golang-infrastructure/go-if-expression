package if_expression

// ReturnUInt32
//  @Description: if实现的三元表达式，返回结果是uint32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的uint32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的uint32
//  @return uint32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt32(boolExpression bool, trueReturnValue, falseReturnValue uint32) uint32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt32Slice
//  @Description: if实现的三元表达式，返回结果是[]uint32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]uint32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]uint32
//  @return []uint32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt32Slice(boolExpression bool, trueReturnValue, falseReturnValue []uint32) []uint32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt32Pointer
//  @Description: if实现的三元表达式，返回结果是*uint32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*uint32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*uint32
//  @return *uint32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt32Pointer(boolExpression bool, trueReturnValue, falseReturnValue *uint32) *uint32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt32PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*uint32
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*uint32
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*uint32
//  @return *uint32: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt32PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*uint32) []*uint32 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
