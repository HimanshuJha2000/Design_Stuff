package validator

import (
	"Microprocessor/enum"
)

func ValidateRegister(registerName string) bool {
	registerType := enum.RegisterType(enum.RegisterTypeType_value[registerName])
	if registerType == enum.RegisterTypeType_NO_TYPE {
		return false
	}
	return true
}

func ValidateOperator(operatorName string) bool {
	operatorType := enum.OperatorType(enum.OperatorType_value[operatorName])
	if operatorType == enum.OperatorType_NO_TYPE {
		return false
	}
	return true
}
