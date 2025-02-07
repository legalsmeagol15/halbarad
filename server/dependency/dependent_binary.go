package dependency

import "fmt"

type depBinary struct {
	nexts            []Dependent
	priors0, priors1 depValTuple
	formatter        func() string
	oper             func(a, b any) any
	value            any
}

func (bd depBinary) GetDependees() []Dependent  { return []Dependent{bd.priors0.dep, bd.priors1.dep} }
func (bd depBinary) GetDependents() []Dependent { return bd.nexts }
func (bd depBinary) GetFormatted() string       { return bd.formatter() }
func (bd depBinary) GetValue() any              { return bd.value }
func (bd depBinary) update(sender Dependent) bool {
	if sender == bd.priors0.dep {
		bd.priors0.value = sender.GetValue()
	} else if sender == bd.priors1.dep {
		bd.priors1.value = sender.GetValue()
	} else {
		panic("crap")
	}
	return bd.updateValue()
}
func (bd depBinary) updateValue() bool {
	oldValue := bd.GetValue()
	bd.value = bd.oper(bd.priors0.value, bd.priors1.value)
	return oldValue != bd.value
}

func NewDependentBinary(input0, input1 any, symbol string) *depBinary {
	priors0, priors1 := newDepValTuple(input0), newDepValTuple(input1)
	result := depBinary{
		priors0:   priors0,
		priors1:   priors1,
		formatter: func() string { return fmt.Sprintf("%s%s%s", priors0.value, symbol, priors1.value) },
	}

	switch symbol {
	case "+":
		result.oper = addition
	case "-":
		result.oper = subtraction
	case "*":
		result.oper = multiplication
	case "/":
		result.oper = division
	}
	result.updateValue()
	return &result
}

func addition(a, b any) any {
	switch _a := a.(type) {
	case float64:
		switch _b := b.(type) {
		case float64:
			return _a + _b
		}
	}
	return depError{
		message: "invalid addition types",
	}
}
func division(a, b any) any {
	switch _a := a.(type) {
	case float64:
		switch _b := b.(type) {
		case float64:
			if b == 0.0 {
				return depError{message: "division by zero error"}
			}
			return _a / _b
		}
	}
	return depError{
		message: "invalid division types",
	}
}
func multiplication(a, b any) any {
	switch _a := a.(type) {
	case float64:
		switch _b := b.(type) {
		case float64:
			return _a * _b
		}
	}
	return depError{
		message: "invalid multiplication types",
	}
}
func subtraction(a, b any) any {
	switch _a := a.(type) {
	case float64:
		switch _b := b.(type) {
		case float64:
			return _a - _b
		}
	}
	return depError{
		message: "invalid subtraction types",
	}
}
