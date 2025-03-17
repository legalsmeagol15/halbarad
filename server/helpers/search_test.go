package helpers

import (
	"reflect"
	"testing"
)

func TestSearchBreadthFirst(t *testing.T) {
	tree := []any{
		[]any{1, 2, 3},
		[]any{4, 5, 6},
		[]any{
			[]any{7, 8, 9},
			[]any{10, 11, 12},
		},
	}
	is_match := func(item any) bool {
		if i, ok := item.(int); ok && i == 5 {
			return true
		}
		return false
	}
	get_next := func(item any) []any {
		if node, ok := item.([]any); ok {
			return node
		}
		return nil

	}

}

func TestSearchDepthFirst(t *testing.T) {
	type args struct {
		start        T
		is_match     func(T) bool
		get_next     func(T) []T
		allow_cycles bool
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchDepthFirst(tt.args.start, tt.args.is_match, tt.args.get_next, tt.args.allow_cycles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchDepthFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchAsync(t *testing.T) {
	type args struct {
		start        T
		is_match     func(T) bool
		get_next     func(T) []T
		allow_cycles bool
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchAsync(tt.args.start, tt.args.is_match, tt.args.get_next, tt.args.allow_cycles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchAsync() = %v, want %v", got, tt.want)
			}
		})
	}
}
