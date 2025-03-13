package graph

import (
	"halbarad/server/dependency/internal/operations"
	"halbarad/server/helpers"
	"sync"
)

var (
	ch_in, Updates = helpers.NewUnboundedChan[*Dep](512)
)

type Dep struct {
	Nexts  []*Dep
	Priors []DepValTuple
	Oper   operations.Oper
	Value  any
}

func (d *Dep) update_value(sender DepValTuple) bool {

	var (
		inputs        = make([]any, len(d.Priors))
		input_changed = false
	)

	if sender.Dep == d {
		for i, p := range d.Priors {
			inputs[i] = p.Value
		}
		input_changed = true
	} else {
		for i, p := range d.Priors {
			if p.Dep == sender.Dep && p.Value != sender.Value {
				d.Priors[i] = sender
				input_changed = true
			}
			inputs[i] = p.Value
		}
	}

	if !input_changed {
		return false
	}

	if new_value, err := d.Oper.Call(inputs...); err == nil {
		old_value := d.Value
		d.Value = new_value
		return d.Value != old_value
	} else {
		old_error := d.Value.(depError)
		new_error := NewError(err.Error(), d)
		if old_error.Equals(&new_error) {
			return false
		}
		d.Value = new_error
		return true
	}

}

func (d *Dep) update(wait bool) {
	var (
		wg           sync.WaitGroup
		_update_func func(DepValTuple, *Dep)
	)

	_update_func = func(input_changed DepValTuple, focus *Dep) {
		defer wg.Done()
		if focus.update_value(input_changed) {
			ch_in <- focus
			wg.Add(len(focus.Nexts))
			dvt := DepValTuple{Dep: focus, Value: focus.Value}
			for _, n := range focus.Nexts {
				go _update_func(dvt, n)
			}
		}
	}

	wg.Add(1)
	self_sender := DepValTuple{Dep: d, Value: nil}
	go _update_func(self_sender, d)

	if wait {
		wg.Wait()
	}

}

func (d *Dep) Update() {
	d.update(true)
}
func (d *Dep) UpdateAsync() {
	d.update(false)
}
