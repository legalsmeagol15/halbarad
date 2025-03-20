package helpers

import (
	"testing"
)

func TestSearchBreadthFirst(t *testing.T) {
	tree := makeAcyclicTree()
	var any_tree any = tree

	is_goal := func(item any) bool { return item == 5 }
	get_next := func(item *any) []*any {
		if node, ok := (*item).([]any); ok {
			result := make([]*any, len(node))
			for i, item := range node {
				result[i] = &item
			}
			return result
		}
		return nil
	}

	bfs := SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if len(bfs) != 2 {
		t.Errorf("Unexpected search results: %s", bfs)
	}
}

func makeAcyclicTree() []any {
	tree := []any{
		[]any{1, 2, 3},
		[]any{4, 5, 6},
		[]any{
			[]any{7, 8, 9},
			[]any{10, 11, 12},
		},
	}
	return tree
}
