package operations

import (
	"fmt"
	"halbarad/server/dependency/internal/literals"
	"math"
	"strings"
)

const (
	MAX_INPUTS = 64
)

var (
	addition = Oper{
		priority: 9,
		can_fit:  inputs_2_plus,
		call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				result := f[0]
				for i := 1; i < len(f); i++ {
					result += f[i]
				}
				return result, nil
			} else {
				return nil, err
			}
		},
		to_string: func(inputs ...any) string {
			return strings.Join(As_strings(inputs), "+")
		},
	}
	division = Oper{
		priority: 2,
		can_fit:  inputs_2,
		call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				return f[0] / f[1], nil
			} else {
				return nil, err
			}
		},
		to_string: func(inputs ...any) string {
			return strings.Join(As_strings(inputs), "/")
		},
	}
	exponentiation = Oper{
		priority: 2,
		can_fit:  inputs_2,
		call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				return math.Pow(float64(f[0]), float64(f[1])), nil
			} else {
				return nil, err
			}
		},
		to_string: func(inputs ...any) string {
			return strings.Join(As_strings(inputs), "^")
		},
	}
	multiplication = Oper{
		priority: 2,
		can_fit:  inputs_2_plus,
		call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				result := f[0]
				for i := 1; i < len(f); i++ {
					result *= f[i]
				}
				return result, nil
			} else {
				return nil, err
			}
		},
		to_string: func(inputs ...any) string {
			return strings.Join(As_strings(inputs), "*")
		},
	}
	negation = Oper{
		priority: 1,
		can_fit:  inputs_1,
		call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				return -f[0], nil
			} else {
				return nil, err
			}
		},
		to_string: func(inputs ...any) string {
			return strings.Join(As_strings(inputs), "-")
		},
	}

	subtraction = Oper{
		priority: 10,
		can_fit:  inputs_2_plus,
		call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				result := f[0]
				for i := 1; i < len(f); i++ {
					result -= f[i]
				}
				return result, nil
			} else {
				return nil, err
			}
		},
		to_string: func(inputs ...any) string {
			return strings.Join(As_strings(inputs), "-")
		},
	}
)

func inputs_1(inputs ...any) bool      { return len(inputs) == 1 }
func inputs_2(inputs ...any) bool      { return len(inputs) == 2 }
func inputs_2_plus(inputs ...any) bool { return len(inputs) >= 2 && len(inputs) <= MAX_INPUTS }
func as_numbers(inputs ...any) ([]literals.Number, error) {
	result := make([]literals.Number, len(inputs))
	for i := 0; i < len(inputs); i++ {
		if f, is_num := inputs[i].(literals.Number); is_num {
			result[i] = f
		} else {
			return nil, fmt.Errorf("type mismatch: %T", inputs[i])
		}
	}
	return result, nil
}
