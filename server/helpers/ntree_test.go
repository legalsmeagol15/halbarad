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

func Test_ntree_Add(t *testing.T) {
	tree := NewNTree[float64](2)

	// This should error with a cardinality mismatch, and state should be unchanged
	e := tree.Add("smooo", NewRegion([]float64{1, 1, 1}, []float64{2, 2, 2}))
	if e != errCardMismatch {
		t.Errorf("failed to detect cardinality mismatch")
	}
	if tree.GetCount() != 0 {
		t.Errorf("state was changed when it shouldn't have been")
	}
	if tree.Contains("smooo") {
		t.Errorf("state was changed when it shouldn't have been")
	}
	if tree.Remove("smooo") != errItemNotContained {
		t.Errorf("Remove didn't fail when it should have")
	}

	// This should not error out
	e = tree.Add("smooo", NewRegion([]float64{1, 1}, []float64{2, 2}))
	if e != nil {
		t.Errorf("cardinality mismatch detected where there isn't one")
	}
	if tree.GetCount() != 1 {
		t.Errorf("state change not reflected in contents count")
	}
	if !tree.Contains("smooo") {
		t.Errorf("the item isn't on the tree")
	}
	if tree.Remove("smooo") != nil {
		t.Errorf("removal failed")
	}

}
