package service

import (
	"Microprocessor/database"
	"Microprocessor/validator"
	"errors"
	"github.com/spf13/cast"
)

func init() {
	AddOperatorInstance = &addOperator{}
}

type addOperator struct {
}

var AddOperatorInstance *addOperator

func (d addOperator) Execute(instruction []string) error {
	valid := instructionValidation(instruction)
	if !valid {
		return errors.New("Not a valid register or operator")
	}
	register, err := database.GetRegister(instruction[1])
	if err != nil {
		database.AddRegister(instruction[1])
		register, _ = database.GetRegister(instruction[1])
	}
	register.Value = register.Value + cast.ToInt32(instruction[2])
	database.UpdateRegister(register)
	return nil
}

func instructionValidation(instruction []string) bool {
	valid := validator.ValidateOperator(instruction[0])
	valid = validator.ValidateRegister(instruction[1])
	return valid
}
