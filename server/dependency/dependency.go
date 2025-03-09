package dependency

import (
	"fmt"
	"halbarad/server/helpers"
	"sync"
)

var (
	updates_in, Updates = helpers.NewUnboundedChan[Dependent](512)
)

// A struct which may have dependents and dependees
type Dependent interface {
	GetDependents() []Dependent
	GetDependees() []Dependent
	GetValue() any
}

// A struct which may be updated in response to changes in the dependees, and then signal
// dependents to update in turn.
type dependent interface {
	getOper() func(...any) any
	getPriorAddr(idx int) (*depValTuple, int)
	setValue(any)
	getValue() any
}

// A struct representing a Dependent object and its value
type depValTuple struct {
	dep   Dependent
	value any
}

func connectDependency(d dependent, inputs ...any) (err error) {
	// It's possible for a panic due to bad matrix indexing to happen here.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("some error - tbd")
		}
	}()

	inputVals := make([]any, len(inputs))
	for i := 0; i < len(inputs); i++ {
		prior, _ := d.getPriorAddr(i)
		switch item := inputs[i].(type) {
		case depValTuple:
			*prior = item
		case Dependent:
			*prior = depValTuple{dep: item, value: item.GetValue()}
		default:
			*prior = depValTuple{dep: nil, value: i}
		}
		inputVals[i] = prior.value

		if prior_D, prior_can_update := (prior.dep).(Dependent); prior_can_update {

			prior_deps := prior_D.GetDependents()
			prior_deps = append(prior_deps, d)
		}
	}
	d.setValue(d.getOper()(inputVals...))

}

func (dl depValTuple) String() string {
	return fmt.Sprintf("->%s", dl.value)
}

func update(d dependent, sender depValTuple) bool {
	var inputs []any
	changed := false
	write_from_sender := func(dvp *depValTuple) {
		if dvp.dep == sender.dep {
			if dvp.value == sender.value {
				return
			} else {
				oldVal := dvp.value
				*dvp = sender
				changed = changed || (oldVal != sender.value)
			}
		}
	}

	// Write the new sender value to this dependent
	prior, length := d.getPriorAddr(0)
	write_from_sender(prior)
	if length > 0 && prior != nil {
		inputs = make([]any, length)
		inputs[0] = prior.value
		for i := 1; i < length; i++ {
			prior, _ = d.getPriorAddr(i)
			write_from_sender(prior)
			inputs[i] = prior.value
		}
	}

	// Perform this dependent's operation, and whether a change happened.
	if changed {
		oldVal := d.getValue()
		newVal := d.getOper()(inputs...)
		if newVal != oldVal {
			d.setValue(newVal)
			return true
		}
	}
	return false
}

func update_func(wg *sync.WaitGroup, focus Dependent, sender depValTuple) {
	defer wg.Done()

	// Not all Dependents can update.  Check.
	if d, can_update := focus.(dependent); can_update {
		if update(d, sender) {
			updates_in <- focus
		} else {
			return
		}
	}

	// Update and children of the focus.
	focus_deps := focus.GetDependents()
	wg.Add(len(focus_deps))
	for _, d := range focus_deps {
		go update_func(wg, d, depValTuple{dep: focus, value: focus.GetValue()})
	}
}

func Update(d Dependent, newVal any, wait bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	go update_func(&wg, d, depValTuple{dep: nil, value: newVal})
	if wait {
		wg.Wait()
	}
}
