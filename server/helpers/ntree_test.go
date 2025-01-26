package helpers

import (
	"testing"
)

func Test_ntree_init(t *testing.T) {
	tree := NewNTree[float64](2)
	if tree.cardinality != 2 {
		t.Errorf("incorrect cardinality: %d", tree.cardinality)
	} else if tree.GetCardinality() != 2 {
		t.Errorf("incorrect GetCardinality: %d", tree.cardinality)
	}

}
