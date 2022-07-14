package if_expression

import "time"

// ReturnTime
//  @Description: if实现的三元表达式，返回结果是time.Time
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的time.Time
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的time.Time
//  @return time.Time: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnTime(boolExpression bool, trueReturnValue, falseReturnValue time.Time) time.Time {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnTimeSlice
//  @Description: if实现的三元表达式，返回结果是[]time.Time
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]time.Time
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]time.Time
//  @return []time.Time: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnTimeSlice(boolExpression bool, trueReturnValue, falseReturnValue []time.Time) []time.Time {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnTimePointer
//  @Description: if实现的三元表达式，返回结果是*time.Time
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的*time.Time
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的*time.Time
//  @return *time.Time: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnTimePointer(boolExpression bool, trueReturnValue, falseReturnValue *time.Time) *time.Time {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnTimePointerSlice
//  @Description: if实现的三元表达式，返回结果是[]*time.Time
//  @param boolExpression: 表达式，最终返回一个布尔值
//  @param trueReturnValue: 当boolExpression返回值为true的时候返回的[]*time.Time
//  @param falseReturnValue: 当boolExpression返回值为false的时候返回的[]*time.Time
//  @return []*time.Time: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnTimePointerSlice(boolExpression bool, trueReturnValue, falseReturnValue []*time.Time) []*time.Time {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

