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
}

func TestAcyclicBasic(t *testing.T) {
	var tree any = makeAcyclicTree()

	is_goal := func(item any) bool { return item == 12 }

	step, _, wait := SearchAsync(&tree, is_goal, get_next, func(any, any) float64 { return 1.0 }, 1000.0)

	wait()
	if *step.Node != 12 {
		t.Errorf("unexpected endpoint: %v", step)
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
