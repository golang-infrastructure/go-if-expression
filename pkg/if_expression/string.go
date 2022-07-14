package if_expression

// ReturnString
//  @Description: if实现的三元表达式，返回结果是string
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的string
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的string
//  @return string: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnString(boolExpression bool, trueReturnValue, falseReturnValue string) string {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnStringSlice
//  @Description: if实现的三元表达式，返回结果是[]string
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]string
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]string
//  @return []string: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnStringSlice(boolExpression bool, trueReturnValue, falseReturnValue []string) []string {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnStringPointer
//  @Description: if实现的三元表达式，返回结果是*string
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*string
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*string
//  @return *string: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnStringPointer(boolExpression bool, trueReturnValue, falseReturnValue *string) *string {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnStringPointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*string
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*string
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*string
//  @return []*string: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnStringPointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*string) []*string {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
