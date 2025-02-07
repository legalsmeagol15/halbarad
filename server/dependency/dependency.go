package dependency

import (
	"fmt"
	"halbarad/server/helpers"
	"sync"
)

var (
	updates_in, Updates = helpers.NewUnboundedChan[Dependent](512)
)

type Dependent interface {
	GetDependents() []Dependent
	GetDependees() []Dependent
	GetFormatted() string
	GetValue() any
}
type dependent interface {
	update(sender Dependent) bool
}
type depValTuple struct {
	dep   Dependent
	value any
}

func (dl depValTuple) String() string {
	return fmt.Sprintf("->%s", dl.value)
}
func newDepValTuple(input any) depValTuple {
	if _i, ok := input.(Dependent); ok {
		return depValTuple{dep: _i, value: _i.GetValue()}
	} else {
		return depValTuple{dep: nil, value: input}
	}
}

func update_func(wg *sync.WaitGroup, sender, focus Dependent) {
	defer wg.Done()
	any_updated := true
	if f, can_update := focus.(dependent); can_update {
		if f.update(sender) {
			updates_in <- focus
			any_updated = true
		}
	}
	if !any_updated {
		return
	}
	focus_deps := focus.GetDependents()
	wg.Add(len(focus_deps))
	for _, d := range focus_deps {
		go update_func(wg, focus, d)
	}
}

func Update(d Dependent, wait bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	go update_func(&wg, nil, d)
	if wait {
		wg.Wait()
	}
}
