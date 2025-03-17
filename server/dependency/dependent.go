package dependency

import (
	"halbarad/server/dependency/internal/graph"
	"halbarad/server/dependency/internal/operations"
	"halbarad/server/helpers"
	"sync"
)

var (
	updates_in, Updates = helpers.NewUnboundedChan[Dependent](512)
)

// An interface which may have dependents (listeners) and dependees (sources).
type Dependent interface {
	GetValue() any
	SetInputs(inputs ...any) graph.DepError
	SetOper(oper operations.Oper) graph.DepError
}

// Creates a new dependency object that will perform the given operation on the given inputs.
func NewDependent(oper operations.Oper, inputs ...any) Dependent {
	c := make(chan graph.DepValTuple)
	d := &graph.Dep{
		Oper:    oper,
		Updates: c,
	}
	for _, input := range inputs {
		if child, ok := input.(*graph.Dep); ok {
			// Since this is a new dependent, we can dispense with search for dependency cycles.
			child.AddListener(d)
			d.AddSource(child)
		}
	}
	go func() {
		for dvt := range c {
			if _, updated := d.UpdateValue(dvt); updated {
				updates_in <- d
				for n := range d.GetNexts() {
					n.Updates <- d.ToDVT()

				}
			}
		}

	}()
	return d
}

func Update(dependent Dependent, new_value any, wait bool) bool {

	var (
		wg  sync.WaitGroup
		dep = dependent.(graph.Dep)
	)

	if dep.Value == new_value {
		return false
	} else {
		dep.Value = new_value
		wg.Add(1)
		updates_in <- dep
		self_sender := graph.DepValTuple{Dep: &dep, Value: new_value}
		dep.Updates <- self_sender
		if wait {
			wg.Wait()
		}
		return true
	}
}
