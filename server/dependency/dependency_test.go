package dependency

import (
	"testing"
)

func TestInit(t *testing.T) {
	simple := NewDependentBinary(4, 2, "+")
	if simple.value != 6 {
		t.Errorf("Incorrect value: %s", simple.value)
	}

}
