package operatorfactory

import (
	"Microprocessor/enum"
	"Microprocessor/operatorexecutor"
	"errors"
)

var operatorMap = make(map[enum.OperatorType]service.OperatorHandler)

func init() {
	operatorMap[enum.OperatorType_ADD] = service.AddOperatorInstance
	operatorMap[enum.OperatorType_ADR] = service.AdrOperatorInstance
	operatorMap[enum.OperatorType_MOV] = service.MovOperatorInstance
	operatorMap[enum.OperatorType_SET] = service.SetOperatorInstance
	operatorMap[enum.OperatorType_INR] = service.InrOperatorInstance
	operatorMap[enum.OperatorType_DCR] = service.DcrOperatorInstance
	operatorMap[enum.OperatorType_RST] = service.RstOperatorInstance
}

func GetOperatorInstance(operatorType enum.OperatorType) (service.OperatorHandler, error) {
	if operator, ok := operatorMap[operatorType]; ok {
		return operator, nil
	}
	return nil, errors.New("Invalid Operator")
}
