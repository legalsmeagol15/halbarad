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
	GetFormatter() string
	GetValue() any
}
type dependent interface {
	getInputs() []any
	getOper() func(a, b any) any
	update(sender Dependent) bool
}
type dependentLink struct {
	dep   Dependent
	value any
}

func (dl dependentLink) String() string {
	return fmt.Sprintf("%s", dl.value)
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
