package if_expression

// ReturnByte
//  @Description: if实现的三元表达式，返回结果是字节byte
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的字节byte
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的字节byte
//  @return byte: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnByte(boolExpression bool, trueReturnValue, falseReturnValue byte) byte {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnByteArray
//  @Description: if实现的三元表达式，返回结果是字节数组[]byte
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的值[]byte
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]byte
//  @return []byte: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnByteArray(boolExpression bool, trueReturnValue, falseReturnValue []byte) []byte {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

func ReturnBytePointer(boolExpression bool, trueReturnValue, falseReturnValue *byte) *byte {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

func ReturnBytePointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*byte) []*byte {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
