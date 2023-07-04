package service

type OperatorHandler interface {
	Execute(instruction []string) error
}
