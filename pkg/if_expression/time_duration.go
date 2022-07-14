package if_expression

import "time"

// ReturnDuration
//  @Description: if实现的三元表达式，返回结果是time.Duration
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的time.Duration
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的time.Duration
//  @return time.Duration: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnDuration(boolExpression bool, trueReturnValue, falseReturnValue time.Duration) time.Duration {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnDurationSlice
//  @Description: if实现的三元表达式，返回结果是[]time.Duration
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]time.Duration
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]time.Duration
//  @return []time.Duration: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnDurationSlice(boolExpression bool, trueReturnValue, falseReturnValue []time.Duration) []time.Duration {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnDurationPointer
//  @Description: if实现的三元表达式，返回结果是*time.Duration
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*time.Duration
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*time.Duration
//  @return *time.Duration: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnDurationPointer(boolExpression bool, trueReturnValue, falseReturnValue *time.Duration) *time.Duration {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnDurationPointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*time.Duration
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*time.Duration
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*time.Duration
//  @return []*time.Duration: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnDurationPointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*time.Duration) []*time.Duration {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}