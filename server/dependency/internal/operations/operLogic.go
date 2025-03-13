package operations

import "strings"

var (
	conjunction = Oper{
		Priority: 9,
		CanFit:   FitChecks{inputs_2_plus},
		Call: func(inputs ...any) (any, error) {
			if f, err := as_bools(inputs); err == nil {
				b := true
				for _, _b := range f {
					b = b && _b
				}
				return b, nil
			} else {
				return nil, err
			}
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "&")
		},
	}
	disjunction = Oper{
		Priority: 9,
		CanFit:   FitChecks{inputs_2_plus},
		Call: func(inputs ...any) (any, error) {
			if f, err := as_bools(inputs); err == nil {
				for _, _b := range f {
					if _b {
						return true, nil
					}
				}
				return false, nil
			} else {
				return nil, err
			}
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "&")
		},
	}

	left_shift = Oper{
		Priority: 9,
		CanFit:   FitChecks{inputs_2},
		Call: func(inputs ...any) (any, error) {
			panic("not implemented yet")
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), "<<")
		},
	}
	right_shift = Oper{
		Priority: 9,
		CanFit:   FitChecks{inputs_2},
		Call: func(inputs ...any) (any, error) {
			panic("not implemented yet")
		},
		ToString: func(inputs ...any) string {
			return strings.Join(as_strings(inputs), ">>")
		},
	}
)
