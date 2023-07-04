package service

import "Microprocessor/database"

func init() {
	InrOperatorInstance = &inrOperator{}
}

type inrOperator struct {
}

var InrOperatorInstance *inrOperator

func (d inrOperator) Execute(instruction []string) error {
	register, err := database.GetRegister(instruction[1])
	if err != nil {
		database.AddRegister(instruction[1])
		register, _ = database.GetRegister(instruction[1])
	}
	register.Value = register.Value + 1
	database.UpdateRegister(register)
	return nil
}
