package enum

type OperatorType int32

const (
	OperatorType_NO_TYPE OperatorType = 0
	OperatorType_ADD     OperatorType = 1
	OperatorType_ADR     OperatorType = 2
	OperatorType_MOV     OperatorType = 3
	OperatorType_SET     OperatorType = 4
	OperatorType_INR     OperatorType = 5
	OperatorType_DCR     OperatorType = 6
	OperatorType_RST     OperatorType = 7
)

// Enum value maps for OperatorType.
var (
	OperatorType_value = map[string]int32{
		"NO_TYPE": 0,
		"ADD":     1,
		"ADR":     2,
		"MOV":     3,
		"SET":     4,
		"INR":     5,
		"DCR":     6,
		"RST":     7,
	}
)
