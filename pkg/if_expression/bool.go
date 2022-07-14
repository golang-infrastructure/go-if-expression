package if_expression

// ReturnBool
//  @Description: if实现的三元表达式，返回结果是布尔bool
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的bool值
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的bool值
//  @return bool: 三元表达式的结果
func ReturnBool(boolExpression bool, trueReturnValue, falseReturnValue bool) bool {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

func ReturnBoolSlice(boolExpression bool, trueReturnValue, falseReturnValue []bool) []bool {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnBoolPointer
//  @Description: if实现的三元表达式，返回结果是布尔指针*bool
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*bool值
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*bool值
//  @return *bool: 三元表达式的结果
func ReturnBoolPointer(boolExpression bool, trueReturnValue, falseReturnValue *bool) *bool {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

func ReturnBoolPointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*bool) []*bool {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
