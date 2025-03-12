package operations

import (
	"fmt"
	"halbarad/server/dependency/internal/literals"
	"strconv"
)

var (
	Operations = map[string][]Oper{
		"+": {addition},
		"^": {exponentiation},
		"/": {division},
		"-": {subtraction, negation},
		"*": {multiplication},
	}
)

type Oper struct {
	priority int

	can_fit   func(inputs ...any) bool
	call      func(inputs ...any) (any, error)
	to_string func(...any) string
}

func As_strings(items ...any) []string {
	result := make([]string, len(items))
	for i := 0; i < len(items); i++ {
		switch item := items[i].(type) {
		case literals.Number:
			result[i] = strconv.FormatFloat(float64(item), 'f', -1, 64)
		case string:
			result[i] = item
		default:
			panic(fmt.Errorf("unimplemented type: %t", item))
		}
	}
	return result
}
