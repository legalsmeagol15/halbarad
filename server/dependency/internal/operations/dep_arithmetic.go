package operations

const (
	MAX_INPUTS = 64
)

var (
	minus = oper{
		priority:   10,
		min_inputs: 2,
		max_inputs: MAX_INPUTS,
		call: func(inputs ...any) any {
			var result NumberType
			if result, is_num := inputs[0].(NumberType); !is_num {
				return NewError("invalid input: %s", inputs[0])
			}
			for i := 1; i < len(inputs); i++ {
				if f, is_num := inputs[i].(numberType); !is_num {
					return depError{message: "Type mismatch"}
				} else {
					result -= f
				}
			}
			return result
		},
	}
)
