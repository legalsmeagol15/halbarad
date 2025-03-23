package helpers

// This does all the searchy things.
func search_sync[TNode comparable](start *TNode,
	work_func func([]*TNode) (*TNode, []*TNode),
	is_goal func(TNode) bool,
	get_next func(*TNode) []*TNode,
	get_weight func(TNode, TNode) float64,
	weight_limit float64) ([]TNode, float64) {

	type datum struct {
		path   []TNode
		weight float64
	}
	var (
		result  = datum{path: nil, weight: weight_limit}
		reached = map[*TNode]datum{start: {path: []TNode{}, weight: 0.0}}
		queue   = []*TNode{start}
	)

	for len(queue) > 0 {
		var focus_node *TNode
		focus_node, queue = work_func(queue)
		focus_d := reached[focus_node]

		if children := get_next(focus_node); children != nil {
			for _, child := range children {
				child_wt := focus_d.weight + get_weight(*focus_node, *child)
				if child_wt > result.weight {
					continue
				} else if is_goal(*child) {
					result = datum{weight: child_wt, path: append(focus_d.path, *focus_node)}

					// Keep going in case a lower-weight path to goal can be found
					continue
				} else if child_d, encountered := reached[child]; encountered && child_d.weight < child_wt {
					continue
				} else {
					reached[child] = datum{
						path:   append(focus_d.path, *focus_node), // The non-assigning use of 'append' is legit. I checked on PG
						weight: child_wt,
					}
					queue = append(queue, child)
				}
			}
		}
	}
	return result.path, result.weight
}

// Returns nil if there is no route between start and the matched goal.
func SearchBreadthFirst[TNode comparable](start *TNode, is_goal func(TNode) bool, get_next func(*TNode) []*TNode, depth_limit int) []TNode {
	work_func := func(queue []*TNode) (*TNode, []*TNode) { return queue[0], queue[1:] }
	wt_func := func(TNode, TNode) float64 { return 1.0 }
	result, _ := search_sync(start, work_func, is_goal, get_next, wt_func, float64(depth_limit))
	return result
}

// Returns nil if there is no route between start and the matched goal.
func SearchDepthFirst[TNode comparable](start *TNode, is_goal func(TNode) bool, get_next func(*TNode) []*TNode, depth_limit int) []TNode {
	work_func := func(stack []*TNode) (*TNode, []*TNode) { return stack[len(stack)-1], stack[:len(stack)-1] }
	wt_func := func(TNode, TNode) float64 { return 1.0 }
	result, _ := search_sync(start, work_func, is_goal, get_next, wt_func, float64(depth_limit))
	return result
}
