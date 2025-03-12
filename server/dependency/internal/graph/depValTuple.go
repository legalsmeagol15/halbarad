package graph

import "fmt"

// A struct representing a Dependent object and its value. Used mostly to cache the values of
// inputs for a dep, but for a few other things too.
type DepValTuple struct {
	Dep   *Dep
	Value any
}

func (dl DepValTuple) String() string {
	return fmt.Sprintf("->%s", dl.Value)
}
