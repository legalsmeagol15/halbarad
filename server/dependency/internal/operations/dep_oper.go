package operations

type oper struct {
	priority, min_inputs, max_inputs int

	call      func(caller *dep, inputs ...any) any
	to_string func(...any) string
}
