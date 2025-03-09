package dependency

type depTrinary struct {
	nexts     []Dependent
	priors    [3]depValTuple
	formatter func() string
	oper      func(...any) any
	value     any
}

func (dt depTrinary) GetDependees() []Dependent {
	return []Dependent{dt.priors[0].dep, dt.priors[1].dep, dt.priors[2].dep}
}
func (dt depTrinary) GetDependents() []Dependent { return dt.nexts }
func (dt depTrinary) GetValue() any              { return dt.value }
func (dt depTrinary) String() string             { return dt.formatter() }

func (dt depTrinary) getOper() func(...any) any { return dt.oper }
func (dt depTrinary) getValue() any             { return dt.value }
func (dt depTrinary) setValue(value any)        { dt.value = value }
func (dt depTrinary) getPriorAddr(idx int) (*depValTuple, int) {
	return &dt.priors[idx], len(dt.priors)
}
