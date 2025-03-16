package helpers

// Returns nil if there is no route between start and the matched goal.
func SearchBreadthFirst[T comparable](start T, is_match func(T) bool, get_next func(T) []T, allow_cycles bool) []T {
	var (
		queue       = []T{start}
		encountered = make(map[T]bool)
		path        = []T{}
	)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		path = append(path, item)
		if _, contained := encountered[item]; contained && !allow_cycles {
			return path
		} else if is_match(item) {
			return path
		} else {
			encountered[item] = true
			queue = append(queue, get_next(item)...)
		}
	}
	return nil
}

// Returns nil if there is no route between start and the matched goal.
func SearchDepthFirst[T comparable](start T, is_match func(T) bool, get_next func(T) []T, allow_cycles bool) []T {
	var (
		queue       = []T{start}
		encountered = make(map[T]bool)
		path        = []T{}
	)
	for len(queue) > 0 {
		item := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		path = append(path, item)
		if _, contained := encountered[item]; contained && !allow_cycles {
			return path
		} else if is_match(item) {
			return path
		} else {
			encountered[item] = true
			queue = append(queue, get_next(item)...)
		}
	}
	return nil
}

func SearchAsync[T comparable](start T, is_match func(T) bool, get_next func(T) []T, allow_cycles bool) []T {
	panic("To be implemented")
}
