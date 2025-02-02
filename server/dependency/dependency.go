package dependency

type Dependent interface {
	GetDependents() []Dependent
	GetValue() any
	GetFormatter() string
}
type dependent interface {
	getInputs() []any
	getOper() func(a, b any) any
	update(sender dependent) bool
}

type binaryDependent struct {
	nexts            []Dependent
	priors0, priors1 struct {
		dep   Dependent
		input any
	}
	oper  func(a, b any) any
	value any
}

func (bd binaryDependent) GetDependents() []Dependent {
	return []Dependent{bd.priors0.dep, bd.priors1.dep}
}
func (bd binaryDependent) GetValue() any    { return bd.value }
func (bd binaryDependent) getInputs() []any { return []any{bd.priors0.input, bd.priors1.input} }
func (bd binaryDependent) update(sender Dependent) bool {
	if sender == bd.priors0.dep {
		bd.priors0.input = sender.GetValue()
	} else if sender == bd.priors1.dep {
		bd.priors1.input = sender.GetValue()
	} else {
		panic("crap")
	}
	oldValue := bd.value
	bd.value = bd.oper(bd.priors0.input, bd.priors1.input)
	return oldValue != bd.value
}

func Update(d Dependent, wait bool) bool {
	if wait {

	}
}
