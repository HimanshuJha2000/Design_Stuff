package main

import (
	"Microprocessor/instructionexecutor"
	"fmt"
)

func main() {
	var instructions = []string{
		"RST",
		"SET A 2",
		"ADD B 2",
		"ADD A 10",
		"INR A",
		"DCR A",
		"ADR B A",
		"MOV A B",
	}
	result, err := instructionexecutor.Execute(instructions)
	fmt.Println(result, err)

}
