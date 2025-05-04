package graph

import (
	"halbarad/server/dependency/internal/operations"
	"halbarad/server/helpers"
	"halbarad/server/helpers/search"
	"iter"
)

// A node in an auto-propogating dependency structure
type Dep struct {

	// The dependent (ie, child) nodes.
	Nexts []*Dep

	// A list of the dependee (ie, parent) nodes, along with their cached values.
	Priors []DepValTuple
	Oper   operations.Oper
	Value  any

	// Close this channel to clean up the Dep
	Updates chan<- DepValTuple
}

func (d Dep) GetValue() any { return d.Value }
func (d *Dep) SetInputs(new_inputs ...any) DepError {
	is_goal := func(a any) bool { return a == d }
	get_next := func(node *any) {

	}
	for _, ni := range new_inputs {
		step := search.SearchAsync(ni, func(a any) { return a == d }, func(*any))

		path := helpers.SearchDepthFirst(ni, func(item Dep) bool { return item == d })
	}

}
func (d Dep) SetOper(oper operations.Oper) DepError {

}

func (d *Dep) UpdateValue(sender_dvt DepValTuple) (any, bool) {

	// The priors are a list of the dependee (ie, parent) Deps stored along with their cached
	// values. We will presume a prior has changed.
	var (
		inputs        = make([]any, len(d.Priors))
		input_changed = false
	)

	if sender_dvt.Dep == d {
		for i, p := range d.Priors {
			inputs[i] = p.Value
		}
		input_changed = true
	} else {
		for i, p := range d.Priors {
			if p.Dep == sender_dvt.Dep {
				d.Priors[i] = sender_dvt
				input_changed = true
			}
			inputs[i] = p.Value
		}
	}

	if !input_changed {
		return nil, false
	}

	if new_value, err := d.Oper.Call(inputs...); err == nil {
		old_value := d.Value
		d.Value = new_value
		return d.Value, d.Value != old_value
	} else {
		old_error := d.Value.(DepError)
		new_error := NewError(err.Error(), d)
		if old_error.Equals(&new_error) {
			return nil, false
		}
		d.Value = new_error
		return nil, true
	}

}

func (d *Dep) ToDVT() DepValTuple     { return DepValTuple{Dep: d, Value: d.Value} }
func (d *Dep) AddListener(other *Dep) { d.Nexts = append(d.Nexts, other) }
func (d *Dep) AddSource(other *Dep)   { d.Priors = append(d.Priors, other.ToDVT()) }
func (d *Dep) GetNextsCount() int     { return len(d.Nexts) }

func (d *Dep) GetNexts() iter.Seq[*Dep] {
	return func(yield func(*Dep) bool) {
		for _, v := range d.Nexts {
			if !yield(v) {
				return
			}
		}
	}
}
func (d *Dep) GetPriors() iter.Seq[*Dep] {
	return func(yield func(*Dep) bool) {
		for _, v := range d.Priors {
			if !yield(v.Dep) {
				return
			}
		}
	}
}
