package search

import (
	"testing"
)

func TestSearchBreadthBasic(t *testing.T) {
	tree := makeAcyclicTree()
	var any_tree any = tree

	is_goal := func(item any) bool { return item == 5 }
	bfs := SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if len(bfs) != 2 {
		t.Errorf("Unexpected search results, len (%d): %v", len(bfs), bfs)
	}

	is_goal = func(item any) bool { return item == 10 }
	bfs = SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if len(bfs) != 3 {
		t.Errorf("Unexpected search results, len (%d): %v", len(bfs), bfs)
	}
	// Test that failure returns a null result
	is_goal = func(item any) bool { return item == -1000 }
	bfs = SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if bfs != nil {
		t.Errorf("Unexpected search results: %v", bfs)
	}
}

func TestDFSBasic(t *testing.T) {
	tree := makeAcyclicTree()
	var any_tree any = tree

	is_goal := func(item any) bool { return item == 12 }
	dfs := SearchDepthFirst(&any_tree, is_goal, get_next, 1000)
	if len(dfs) != 3 {
		t.Errorf("Unexpected search results, len (%d): %v", len(dfs), dfs)
	}

	is_goal = func(item any) bool { return item == 2 }
	dfs = SearchDepthFirst(&any_tree, is_goal, get_next, 1000)
	if len(dfs) != 2 {
		t.Errorf("Unexpected search results, len (%d): %v", len(dfs), dfs)
	}

	// Test that failure returns a null result
	is_goal = func(item any) bool { return item == -1000 }
	dfs = SearchDepthFirst(&any_tree, is_goal, get_next, 1000)
	if dfs != nil {
		t.Errorf("Unexpected search results: %v", dfs)
	}
}

func TestAcyclicBasic(t *testing.T) {
	var tree any = makeAcyclicTree()
	equal_weight := func(any, any) float64 { return 1.0 }

	is_goal := func(item any) bool { return item == 12 }

	step, _, wait := SearchAsync(&tree, is_goal, get_next, equal_weight, 1000.0)
	if !wait() {
		t.Errorf("unexpected failure to find goal.")
	} else if *step.Node != 12 {
		t.Errorf("unexpected endpoint: %v", step)
	}

	is_goal = func(item any) bool { return item == -1000 }
	step, _, wait = SearchAsync(&tree, is_goal, get_next, equal_weight, 1000)
	if wait() {
		t.Errorf("unexpected success in finding goal")
	} else if step.Depth >= 0 {
		t.Errorf("result step is malformed for what should be a failure: %v", *step)
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

func get_next(item *any) []*any {
	if node, ok := (*item).([]any); ok {
		result := make([]*any, len(node))
		for i, item := range node {
			result[i] = &item
		}
		return result
	}
	return nil
}
