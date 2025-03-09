package dependency

type depBinary struct {
	nexts     []Dependent
	priors    [2]depValTuple
	formatter func() string
	oper      func(...any) any
	value     any
}

func (db depBinary) GetDependees() []Dependent {
	return []Dependent{db.priors[0].dep, db.priors[1].dep}
}
func (db depBinary) GetDependents() []Dependent { return db.nexts }
func (db depBinary) GetValue() any              { return db.value }
func (db depBinary) String() string             { return db.formatter() }

func (db depBinary) getOper() func(...any) any { return db.oper }
func (db depBinary) getValue() any             { return db.value }
func (db depBinary) setValue(value any)        { db.value = value }
func (db depBinary) getPriorAddr(idx int) (*depValTuple, int) {
	return &db.priors[idx], len(db.priors)
}

func newDepBinary(inputs [2]any) *depBinary {
	db := depBinary{}
	for idx, item := range inputs {
		switch i := item.(type) {
		case depValTuple:
			db.priors[idx] = i
		case Dependent:
			db.priors[idx] = depValTuple{dep: i, value: i.GetValue()}
		default:
			db.priors[idx] = depValTuple{dep: nil, value: i}
		}
	}
}
