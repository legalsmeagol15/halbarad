package dependency

import (
	"halbarad/server/dependency/internal/graph"
	"halbarad/server/helpers"
	"sync"
)

var (
	updates_in, Updates = helpers.NewUnboundedChan[Dependent](512)
)

// An interface which may have dependents and dependees
type Dependent interface {
	GetDependents() []Dependent
	GetDependees() []Dependent
	GetValue() any
}

func Update(dependent Dependent, new_value any, wait bool) bool {

	var (
		wg          sync.WaitGroup
		update_func func(graph.Dep, graph.DepValTuple)
		changed     = false
	)
	update_func = func(d graph.Dep, sender graph.DepValTuple) {
		defer wg.Done()
		if d.update_value(sender) {
			changed = true
		}
		wg.Add(len(d.Nexts))
		for _, child := range d.Nexts {
			sender := graph.DepValTuple{dep: &d, value: d.Value}
			go update_func(child, sender)
		}
	}

	if d, can_update := dependent.(graph.Dep); !can_update {
		return false
	} else if d.Value == new_value {
		return false
	} else {
		d.Value = new_value
		wg.Add(1)
		updates_in <- d
		self_sender := graph.DepValTuple{dep: &d, value: new_value}
		go update_func(d, self_sender)
		if wait {
			wg.Wait()
		}
		return changed
	}

}

func (d graph.Dep) GetDependees() []Dependent {
	result := make([]Dependent, len(d.Priors))
	for i := 0; i < len(d.Priors); i++ {
		result[i] = d.Priors[i].dep
	}
	return result
}
func (d graph.Dep) GetDependents() []Dependent {
	result := make([]Dependent, len(d.Priors))
	for i := 0; i < len(d.Priors); i++ {
		result[i] = d.Nexts[i]
	}
	return result
}
func (d graph.Dep) GetValue() any { return d.Value }

func (d *graph.Dep) update_value(sender graph.DepValTuple) bool {
	inputs := make([]any, len(d.Priors))
	changed := false

	for i := 0; i < len(d.Priors); i++ {
		if d.Priors[i].dep == sender.dep && d.Priors[i].Value != sender.Value {
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
