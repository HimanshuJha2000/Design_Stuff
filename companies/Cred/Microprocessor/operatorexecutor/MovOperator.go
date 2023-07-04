package service

import "Microprocessor/database"

func init() {
	MovOperatorInstance = &movOperator{}
}

type movOperator struct {
}

var MovOperatorInstance *movOperator

func (d movOperator) Execute(instruction []string) error {
	registerA, err := database.GetRegister(instruction[1])
	if err != nil {
		database.AddRegister(instruction[1])
		registerA, _ = database.GetRegister(instruction[1])
	}
	registerB, err := database.GetRegister(instruction[2])
	if err != nil {
		database.AddRegister(instruction[1])
		registerB, _ = database.GetRegister(instruction[2])
	}

	registerA.Value = registerB.Value
	database.UpdateRegister(registerA)

	return nil
}
