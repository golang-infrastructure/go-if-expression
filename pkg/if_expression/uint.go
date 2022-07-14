package if_expression

// ReturnUInt
//  @Description: if实现的三元表达式，返回结果是uint
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的uint
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的uint
//  @return uint: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUInt(boolExpression bool, trueReturnValue, falseReturnValue uint) uint {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUIntSlice
//  @Description: if实现的三元表达式，返回结果是[]uint
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]uint
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]uint
//  @return []uint: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUIntSlice(boolExpression bool, trueReturnValue, falseReturnValue []uint) []uint {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUIntPointer
//  @Description: if实现的三元表达式，返回结果是*uint
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*uint
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*uint
//  @return *uint: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUIntPointer(boolExpression bool, trueReturnValue, falseReturnValue *uint) *uint {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnUIntPointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*uint
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*uint
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*uint
//  @return *uint: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnUIntPointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*uint) []*uint {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
