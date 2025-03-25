package search

import (
	"sync"
)

type Step[TNode any] struct {
	Node   *TNode
	Weight float64
	Prior  *Step[TNode]
	Depth  int
}

func SearchAsync[TNode comparable](
	start *TNode,
	is_goal func(TNode) bool,
	get_next func(*TNode) []*TNode,
	get_weight func(TNode, TNode) float64,
	weight_limit float64) (*Step[TNode], func(), func() bool) {

	var (
		driver func(*Step[TNode], *TNode, float64, int)
		result = &Step[TNode]{Prior: nil, Depth: -1}

		ch_cancel = make(chan any, 1)
		reached   = map[*TNode]*Step[TNode]{start: {Node: start, Prior: nil, Weight: 0.0, Depth: 0}}
		wg        = sync.WaitGroup{}
		cancelled = false
		lock      = sync.RWMutex{}

		locked_read = func(node *TNode) *Step[TNode] {
			lock.RLocker().Lock()
			defer lock.RLocker().Unlock()
			if s, ok := reached[node]; ok {
				return s
			} else {
				return nil
			}
		}

		cancel = func() {
			defer func() { recover() }()
			close(ch_cancel)
		}
		wait = func() bool {
			wg.Wait()
			cancel()
			return result.Depth >= 0
		}
	)

	go func() {
		// This goroutine will run until the search is cancelled, and then sets the cancelled flag.
		<-ch_cancel
		cancelled = true
	}()

	driver = func(sender *Step[TNode], focus_node *TNode, weight float64, depth int) {
		defer wg.Done()
		focus_step := Step[TNode]{
			Node:   focus_node,
			Prior:  sender,
			Weight: weight,
			Depth:  depth,
		}

		if weight > weight_limit {
			return
		} else if is_goal(*focus_node) {
			*result = focus_step
			cancel()
			return
		} else if s := locked_read(focus_node); s != nil && weight > s.Weight {
			return
		} else if cancelled {
			return
		} else if children := get_next(focus_node); children != nil {
			func() {
				lock.Lock()
				defer lock.Unlock()
				if r, ok := reached[focus_node]; ok && weight > r.Weight {
					// We're checking weight once more because the reached[focus_node] step may
					// have been overwritten since the last locked_read.
					return
				}
				reached[focus_node] = &focus_step
			}()
			wg.Add(len(children))
			for _, child := range children {
				go driver(&focus_step, child, focus_step.Weight+get_weight(*focus_node, *child), focus_step.Depth+1)
			}
		}
	}

	// Kick it off
	wg.Add(1)
	driver(nil, start, 0.0, 0)

	return result, cancel, wait
}

// This does all the searchy things.
func search_sync[TNode comparable](start *TNode,
	pop func([]*Step[TNode]) (*Step[TNode], []*Step[TNode]),
	is_goal func(TNode) bool,
	get_next func(*TNode) []*TNode,
	get_weight func(TNode, TNode) float64,
	weight_limit float64) Step[TNode] {

	var (
		start_step = &Step[TNode]{Node: start, Weight: 0.0, Prior: nil, Depth: 0}
		reached    = map[*TNode]*Step[TNode]{start: start_step}
		work_queue = []*Step[TNode]{start_step}
	)

	for len(work_queue) > 0 {
		var focus_step *Step[TNode]
		focus_step, work_queue = pop(work_queue)
		focus_node := focus_step.Node
		if is_goal(*focus_step.Node) {
			return *focus_step
		} else if children := get_next(focus_node); children == nil {
			continue
		} else {
			for _, child := range children {
				child_wt := focus_step.Weight + get_weight(*focus_node, *child)
				if child_wt > weight_limit {
					continue
				} else if child_node, exists := reached[child]; exists && child_wt > child_node.Weight {
					continue
				} else {
					child_step := &Step[TNode]{
						Node:   child,
						Weight: child_wt,
						Prior:  focus_step,
						Depth:  focus_step.Depth + 1,
					}
					reached[child] = child_step
					work_queue = append(work_queue, child_step)
				}
			}
		}

	}
	return Step[TNode]{
		Node:   nil,
		Prior:  nil,
		Weight: 0.0,
		Depth:  -1,
	}
}

// Returns nil if there is no route between start and the matched goal.
func SearchBreadthFirst[TNode comparable](start *TNode, is_goal func(TNode) bool, get_next func(*TNode) []*TNode, depth_limit int) []*TNode {
	work_func := func(queue []*Step[TNode]) (*Step[TNode], []*Step[TNode]) { return queue[0], queue[1:] }
	wt_func := func(TNode, TNode) float64 { return 1.0 }
	result := search_sync(start, work_func, is_goal, get_next, wt_func, float64(depth_limit))
	return toTNodeArray(&result)
}

// Returns nil if there is no route between start and the matched goal.
func SearchDepthFirst[TNode comparable](start *TNode, is_goal func(TNode) bool, get_next func(*TNode) []*TNode, depth_limit int) []*TNode {
	work_func := func(stack []*Step[TNode]) (*Step[TNode], []*Step[TNode]) {
		return stack[len(stack)-1], stack[:len(stack)-1]
	}
	wt_func := func(TNode, TNode) float64 { return 1.0 }
	result := search_sync(start, work_func, is_goal, get_next, wt_func, float64(depth_limit))
	return toTNodeArray(&result)
}

func toTNodeArray[TNode comparable](step *Step[TNode]) []*TNode {
	if step.Depth < 0 {
		return nil
	}
	result := make([]*TNode, step.Depth)
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = step.Node
		step = step.Prior
	}
	return result
}
