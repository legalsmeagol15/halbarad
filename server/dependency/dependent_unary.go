package dependency

type depUnary struct {
	nexts     []Dependent
	prior     depValTuple
	formatter func() string
	oper      func(...any) any
	value     any
}

func (du depUnary) GetDependees() []Dependent {
	return []Dependent{du.prior.dep}
}
func (du depUnary) GetDependents() []Dependent { return du.nexts }
func (du depUnary) GetValue() any              { return du.value }
func (du depUnary) String() string             { return du.formatter() }

func (du depUnary) getOper() func(...any) any                { return du.oper }
func (du depUnary) getValue() any                            { return du.value }
func (du depUnary) setValue(value any)                       { du.value = value }
func (du depUnary) getPriorAddr(idx int) (*depValTuple, int) { return &du.prior, 1 }
