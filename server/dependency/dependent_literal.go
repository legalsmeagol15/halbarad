package dependency

import "fmt"

type depLiteral struct {
	value any
}

func (dl depLiteral) GetDependents() []Dependent { panic("No dependents for a literal.") }
func (dl depLiteral) GetDependees() []Dependent  { panic("No dependees for a literal.") }
func (dl depLiteral) GetFormatter() string       { return fmt.Sprintf("%s", dl.value) }
func (dl depLiteral) GetValue() any              { return dl.value }

func link_or_literal(d Dependent) dependentLink {
	if dl, ok := d.(depLiteral); ok {
		return dependentLink{dep: nil, value: dl.value}
	} else {
		return dependentLink{dep: d, value: d.GetValue()}
	}
}
