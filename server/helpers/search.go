package helpers

// This does all the searchy things.
func search_sync[T comparable](start T,
	work_func func([]T) (T, []T),
	is_match func(T) bool,
	get_next func(T) []T,
	get_weight func(T, T) float64,
	allow_cycles bool,
	depth_limit int) ([]T, float64) {

	type datum struct {
		item   T
		path   []T
		weight float64
	}
	var (
		reached = map[T]datum{start: datum{item: start, path: []T{}, weight: 0.0}}
		queue   = []T{start}
	)

	for len(queue) > 0 {
		var item T
		item, queue = work_func(queue)
		d := reached[item]
		if is_match(item) {
			return d.path, d.weight
		} else {
			for _, child := range get_next(item) {
				child_wt := get_weight(item, child)
				if child_d, encountered := reached[child]; encountered && child_d.weight < d.weight+child_wt {
					continue
				} else {
					reached[child] = datum{item: child, path: append(d.path, item), weight: child_wt} // The non-assigning use of 'append' is legit. I checked on PG
					queue = append(queue, child)
				}
			}
		}
	}
	return nil, 0.0
}

// Returns nil if there is no route between start and the matched goal.
func SearchBreadthFirst[T comparable](start T, is_match func(T) bool, get_next func(T) []T, allow_cycles bool, depth_limit int) []T {
	work_func := func(stack []T) (T, []T) {
		return stack[0], stack[:1]
	}
	wt_func := func(T, T) float64 { return 1.0 }
	result, _ := search_sync(start, work_func, is_match, get_next, wt_func, allow_cycles, depth_limit)
	return result
}

// Returns nil if there is no route between start and the matched goal.
func SearchDepthFirst[T comparable](start T, is_match func(T) bool, get_next func(T) []T, allow_cycles bool, depth_limit int) []T {
	work_func := func(stack []T) (T, []T) {
		return stack[len(stack)-1], stack[:len(stack)-1]
	}
	wt_func := func(T, T) float64 { return 1.0 }
	result, _ := search_sync(start, work_func, is_match, get_next, wt_func, allow_cycles, depth_limit)
	return result
}

func SearchAsync[T comparable](start T, is_match func(T) bool, get_next func(T) []T, allow_cycles bool) []T {
	panic("To be implemented")
}
