package service

import "Microprocessor/database"

func init() {
	AdrOperatorInstance = &adrOperator{}
}

type adrOperator struct {
}

var AdrOperatorInstance *adrOperator

func (d adrOperator) Execute(instruction []string) error {
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

	registerA.Value = registerA.Value + registerB.Value
	database.UpdateRegister(registerA)

	return nil
}
