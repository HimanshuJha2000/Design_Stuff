package service

import (
	"Microprocessor/database"
	"github.com/spf13/cast"
)

func init() {
	SetOperatorInstance = &setOperator{}
}

type setOperator struct {
}

var SetOperatorInstance *setOperator

func (d setOperator) Execute(instruction []string) error {
	register, err := database.GetRegister(instruction[1])
	if err != nil {
		database.AddRegister(instruction[1])
		register, _ = database.GetRegister(instruction[1])
	}
	register.Value = cast.ToInt32(instruction[2])
	database.UpdateRegister(register)
	return nil
}
