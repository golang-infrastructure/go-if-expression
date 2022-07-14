package if_expression

// ReturnRune
//  @Description: if实现的三元表达式，返回结果是rune
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的rune
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的rune
//  @return rune: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnRune(boolExpression bool, trueReturnValue, falseReturnValue rune) rune {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}


// ReturnRuneSlice
//  @Description: if实现的三元表达式，返回结果是[]rune
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]rune
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]rune
//  @return []rune: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnRuneSlice(boolExpression bool, trueReturnValue, falseReturnValue []rune) []rune {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnRunePointer
//  @Description: if实现的三元表达式，返回结果是*rune
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*rune
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*rune
//  @return *rune: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnRunePointer(boolExpression bool, trueReturnValue, falseReturnValue *rune) *rune {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnRunePointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*rune
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*rune
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*rune
//  @return []*rune: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnRunePointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*rune) []*rune {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
