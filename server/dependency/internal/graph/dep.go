package graph

import "halbarad/server/dependency/internal/operations"

type Dep struct {
	Nexts  []Dep
	Priors []depValTuple
	Oper   operations.Oper
	Value  any
}

func (d *Dep) UpdateValue(sender DepValTuple) bool {
	inputs := make([]any, len(d.Priors))
	changed := false

	for i := 0; i < len(d.Priors); i++ {
		if d.Priors[i].dep == sender.Dep && d.Priors[i].Value != sender.Value {
			d.Priors[i] = sender
			changed = true
		}
		inputs[i] = d.Priors[i].Value
	}
	if changed {
		old_value := d.Value

		d.Value = d.oper(inputs...)
		return old_value != d.Value
	}
	return false
}
