package service

import (
	"Microprocessor/database"
)

func init() {
	DcrOperatorInstance = &dcrOperator{}
}

type dcrOperator struct {
}

var DcrOperatorInstance *dcrOperator

func (d dcrOperator) Execute(instruction []string) error {
	register, err := database.GetRegister(instruction[1])
	if err != nil {
		database.AddRegister(instruction[1])
		register, _ = database.GetRegister(instruction[1])
	}
	register.Value = register.Value - 1
	database.UpdateRegister(register)
	return nil
}
