package graph

type Dep struct {
	Nexts  []Dep
	Priors []depValTuple
	Oper   func(...any) any
	Value  any
}
