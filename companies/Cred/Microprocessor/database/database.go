package database

import (
	"Microprocessor/model"
	"errors"
)

var registerDb = make(map[string]model.Register)

func GetAllRegister() []model.Register {
	var registers []model.Register
	for _, register := range registerDb {
		registers = append(registers, register)
	}
	return registers
}

func UpdateRegister(register model.Register) error {
	if _, ok := registerDb[register.Name]; ok {
		registerDb[register.Name] = register
		return nil
	}
	err := errors.New("register not exist")
	return err
}

func Reset() {
	for registerName, _ := range registerDb {
		register := registerDb[registerName]
		register.Value = 0
		registerDb[registerName] = register
	}
}

func AddRegister(registerName string) {
	register := model.Register{
		Name:  registerName,
		Value: 0,
	}
	registerDb[registerName] = register
}

func GetRegister(registerName string) (model.Register, error) {
	if register, ok := registerDb[registerName]; ok {
		return register, nil
	}
	register := model.Register{}
	err := errors.New("register not exist")
	return register, err
}
