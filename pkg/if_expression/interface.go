package if_expression

// ReturnInterface
//  @Description:
//  @param boolExpression
//  @param trueReturnValue
//  @param falseReturnValue
//  @return interface{}
func ReturnInterface(boolExpression bool, trueReturnValue, falseReturnValue interface{}) interface{} {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnInterfaceSlice
//  @Description:
//  @param boolExpression
//  @param trueReturnValue
//  @param falseReturnValue
//  @return []interface{}
func ReturnInterfaceSlice(boolExpression bool, trueReturnValue, falseReturnValue []interface{}) []interface{} {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
