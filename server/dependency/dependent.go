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

type dep graph.Dep

func (d dep) GetDependees() []Dependent {
	panic("do I need this?")
	// result := make([]Dependent, len(d.Priors))
	// for i := 0; i < len(d.Priors); i++ {
	// 	result[i] = d.Priors[i].Dep
	// }
	// return result
}
func (d dep) GetDependents() []Dependent {
	panic("to be implemented")
	// result := make([]Dependent, len(d.Priors))
	// for i := 0; i < len(d.Priors); i++ {

	// }
	// return result
}
func (d dep) GetValue() any { return d.Value }

func Update(dependent Dependent, new_value any, wait bool) bool {

	var (
		wg          sync.WaitGroup
		update_func func(graph.Dep, graph.DepValTuple)
		changed     = false
	)
	update_func = func(d graph.Dep, sender graph.DepValTuple) {
		defer wg.Done()
		if d.UpdateValue(sender) {
			changed = true
		}
		wg.Add(len(d.Nexts))
		for _, child := range d.Nexts {
			sender := graph.DepValTuple{Dep: &d, Value: d.Value}
			go update_func(child, sender)
		}
	}

	if d, can_update := dependent.(dep); !can_update {
		return false
	} else if d.Value == new_value {
		return false
	} else {
		d.Value = new_value
		wg.Add(1)
		updates_in <- d
		self_sender := graph.DepValTuple{Dep: &d, Value: new_value}
		go update_func(d, self_sender)
		if wait {
			wg.Wait()
		}
		return changed
	}

}
