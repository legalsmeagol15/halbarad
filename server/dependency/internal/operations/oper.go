package operations

import (
	"fmt"
	"halbarad/server/dependency/internal/literals"
	"strconv"
)

const (
	MAX_INPUTS = 64
)

var (
	Operations = map[string][]Oper{
		"+": {addition},
		"^": {exponentiation},
		"/": {division},
		"-": {negation, subtraction},
		"*": {multiplication},

		"|":  {disjunction},
		"&":  {conjunction},
		"<<": {left_shift},
		">>": {right_shift},
	}
)

type Oper struct {
	Priority int

	CanFit   FitChecks
	Call     func(inputs ...any) (any, error)
	ToString func(...any) string
}

type FitChecks []func(...any) bool

func inputs_1(inputs ...any) bool      { return len(inputs) == 1 }
func inputs_2(inputs ...any) bool      { return len(inputs) == 2 }
func inputs_2_plus(inputs ...any) bool { return len(inputs) >= 2 && len(inputs) <= MAX_INPUTS }

func as_bools(inputs ...any) ([]bool, error) {
	result := make([]bool, len(inputs))
	for i, item := range inputs {
		switch typed := item.(type) {
		case bool:
			result[i] = typed
		case literals.Number:
			result[i] = typed != 0.0
		case int:
			result[i] = typed != 0
		}
	}
	return result, nil
}
func as_numbers(inputs ...any) ([]literals.Number, error) {
	result := make([]literals.Number, len(inputs))
	for i, item := range inputs {
		switch typed := item.(type) {
		case literals.Number:
			result[i] = typed
		case string:
			if f, err := strconv.ParseFloat(typed, 64); err != nil {
				result[i] = literals.Number(f)
			} else {
				return nil, fmt.Errorf("type mismatch: unparseable string")
			}
		default:
			return nil, fmt.Errorf("type mismatch: %T", inputs[i])
		}
	}
	return result, nil
}
func as_strings(items ...any) []string {
	result := make([]string, len(items))
	for i, _any := range items {
		switch item := _any.(type) {
		case literals.Number:
			result[i] = strconv.FormatFloat(float64(item), 'f', -1, 64)
		case string:
			result[i] = item
		default:
			panic(fmt.Errorf("unimplemented type: %t", item))
		}
	}
	return result
}
