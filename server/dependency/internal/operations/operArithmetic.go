package operations

import (
	"math"
	"strings"
)

var (
	addition = Oper{
		Priority: 9,
		CanFit:   FitChecks{inputs_2_plus},
		Call: func(inputs ...any) (any, error) {
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
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "+")
		},
	}
	division = Oper{
		Priority: 2,
		CanFit:   FitChecks{inputs_2},
		Call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				return f[0] / f[1], nil
			} else {
				return nil, err
			}
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "/")
		},
	}
	exponentiation = Oper{
		Priority: 2,
		CanFit:   FitChecks{inputs_2},
		Call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				return math.Pow(float64(f[0]), float64(f[1])), nil
			} else {
				return nil, err
			}
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "^")
		},
	}
	multiplication = Oper{
		Priority: 2,
		CanFit:   FitChecks{inputs_2_plus},
		Call: func(inputs ...any) (any, error) {
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
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "*")
		},
	}
	negation = Oper{
		Priority: 1,
		CanFit:   FitChecks{inputs_1},
		Call: func(inputs ...any) (any, error) {
			if f, err := as_numbers(inputs); err == nil {
				return -f[0], nil
			} else {
				return nil, err
			}
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "-")
		},
	}

	subtraction = Oper{
		Priority: 10,
		CanFit:   FitChecks{inputs_2_plus},
		Call: func(inputs ...any) (any, error) {
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
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "-")
		},
	}
)
