package dependency

import (
	"sync"
)

var (
	Updates = make(chan Dependent, 512)
)

type Dependent interface {
	GetDependents() []Dependent
	GetValue() any
	GetFormatter() string
}
type dependent interface {
	getInputs() []any
	getOper() func(a, b any) any
	update(sender Dependent) bool
}

func update_func(wg *sync.WaitGroup, sender, focus Dependent) {
	defer wg.Done()
	if f, can_update := focus.(dependent); can_update {
		if f.update(sender) {
			Updates <- focus
		}
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
