package dependency

type NumberType float64

var (
	//TODO: there's gotta be a better way of populating this in init()
	opers = make(map[string]oper)
)

func float_or_error(item any) (float64, any) {
	if f, can_be_float := item.(float64); can_be_float {
		return f, nil
	}
	if _, is_error := item.(depError); is_error {
		return 0.0, item
	}

	return 0.0, depError{message: "invalid operation"}
}
func operMinus(inputs ...any) oper {
	if len(inputs) == 1 {
		return _negation
	} else if len(inputs) >= 2 {
		_subtraction
	}
}
func _negation(inputs ...any) any {
	if f, is_num := inputs[0].(numberType); is_num {
		return -f
	} else {
		return inputs[0]
	}
}
func _subtraction(inputs ...any) any {
	var result numberType
	if result, is_num := inputs[0]; !is_num {
		return depError{message: "Type mismatch"}
	}
	for i := 1; i < len(inputs); i++ {
		if f, is_num := inputs[i].(numberType); !is_num {
			return depError{message: "Type mismatch"}
		} else {
			result -= f
		}
	}
	return result
}
