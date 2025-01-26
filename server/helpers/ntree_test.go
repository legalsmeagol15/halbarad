package helpers

import (
	"testing"
)

func Test_ntree_init(t *testing.T) {
	tree := NewNTree[float64](2)
	if tree.GetCardinality() != 2 {
		t.Errorf("incorrect GetCardinality: %d", tree.GetCardinality())
	}
	if tree.GetCount() != 0 {
		t.Errorf("incorrect contents count: %d", tree.GetCount())
	}

}
