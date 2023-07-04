package enum

type RegisterType int32

const (
	RegisterTypeType_NO_TYPE RegisterType = 0
	RegisterTypeType_A       RegisterType = 1
	RegisterTypeTypee_B      RegisterType = 2
	RegisterTypeType_C       RegisterType = 3
	RegisterTypeType_D       RegisterType = 4
)

// Enum value maps for RegisterType.
var (
	RegisterTypeType_name = map[int32]string{
		0: "NO_TYPE",
		1: "A",
		2: "B",
		3: "C",
		4: "D",
	}
	RegisterTypeType_value = map[string]int32{
		"NO_TYPE": 0,
		"A":       1,
		"B":       2,
		"C":       3,
		"D":       4,
	}
)
