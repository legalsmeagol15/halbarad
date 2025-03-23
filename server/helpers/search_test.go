package helpers

import (
	"testing"
)

func TestSearchBreadthBasic(t *testing.T) {
	tree := makeAcyclicTree()
	var any_tree any = tree

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

	is_goal := func(item any) bool { return item == 5 }
	bfs := SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if len(bfs) != 2 {
		t.Errorf("Unexpected search results, len (%d): %s", len(bfs), bfs)
	}

	is_goal = func(item any) bool { return item == 10 }
	bfs = SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if len(bfs) != 3 {
		t.Errorf("Unexpected search results, len (%d): %s", len(bfs), bfs)
	}
}

func TestDFSBasic(t *testing.T) {
	tree := makeAcyclicTree()
	var any_tree any = tree

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

	is_goal := func(item any) bool { return item == 12 }
	dfs := SearchDepthFirst(&any_tree, is_goal, get_next, 1000)
	if len(dfs) != 3 {
		t.Errorf("Unexpected search results, len (%d): %s", len(dfs), dfs)
	}

	is_goal = func(item any) bool { return item == 2 }
	dfs = SearchBreadthFirst(&any_tree, is_goal, get_next, 1000)
	if len(dfs) != 2 {
		t.Errorf("Unexpected search results, len (%d): %s", len(dfs), dfs)
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
