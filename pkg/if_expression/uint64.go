package if_expression

// ReturnUInt64
//  @Description: if实现的三元表达式，返回结果是uint64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的uint64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的uint64
//  @return uint64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt64(boolExpression bool, trueReturnValue, falseReturnValue uint64) uint64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt64Slice
//  @Description: if实现的三元表达式，返回结果是[]uint64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]uint64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]uint64
//  @return []uint64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt64Slice(boolExpression bool, trueReturnValue, falseReturnValue []uint64) []uint64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt64Pointer
//  @Description: if实现的三元表达式，返回结果是*uint64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*uint64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*uint64
//  @return *uint64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt64Pointer(boolExpression bool, trueReturnValue, falseReturnValue *uint64) *uint64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUInt64PointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*uint64
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*uint64
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*uint64
//  @return []*uint64: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt64PointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*uint64) []*uint64 {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
