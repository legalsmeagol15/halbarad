package dependency

import "fmt"

type depUnary struct {
	nexts     []Dependent
	prior     depValTuple
	formatter func() string
	oper      func(a any) any
	value     any
}

func (du depUnary) GetDependees() []Dependent  { return []Dependent{du.prior.dep} }
func (du depUnary) GetDependents() []Dependent { return du.nexts }
func (du depUnary) GetFormatted() string       { return du.formatter() }
func (du depUnary) GetValue() any              { return du.value }
func (du depUnary) update(sender Dependent) bool {
	if sender == du.prior.dep {
		du.prior.value = sender.GetValue()
	} else {
		panic("crap")
	}
	return du.updateValue()
}
func (du depUnary) updateValue() bool {
	oldValue := du.value
	du.value = du.oper(du.prior.value)
	return oldValue != du.value
}

func NewDependentUnary(input Dependent, symbol string) *depUnary {
	priors := newDepValTuple(input)
	result := depUnary{
		prior:     priors,
		formatter: func() string { return fmt.Sprintf("%s%s", symbol, priors.value) },
	}

	switch symbol {
	case "-":
		result.oper = negation
	}
	result.updateValue()
	return &result
}

func negation(a any) any {
	switch _a := a.(type) {
	case float64:
		return 0.0 - _a
	}
	return depError{
		message: "invalid negation types",
	}
}
