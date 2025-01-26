package helpers

import (
	"errors"
	"math"
)

var (
	errCardMismatch     error = errors.New("cardinality mismatch")
	errItemContained    error = errors.New("item already contained")
	errItemNotContained error = errors.New("item not contained")
	maxNodeItems        int   = 32
)

type NTree[T comparable] struct {
	root            *nTreeNode[T]
	contents_node   map[any]*nTreeNode[T]
	contents_region map[any]Region
	cardinality     int
	count           int
}

func (t *NTree[T]) Add(item any, itemRegion Region) error {
	if _, c := itemRegion.GetPoints().Dims(); c != t.cardinality {
		return errCardMismatch
	} else if t.root == nil {
		t.root = &nTreeNode[T]{
			parent: nil,
			bounds: itemRegion.(region),
			depth:  0,
		}
	} else if _, contained := t.contents_node[item]; contained {
		return errItemContained
	}

	containingNode := t.root.add(item, itemRegion)
	t.contents_region[item] = itemRegion
	t.count += 1

	if containingNode != nil {
		// If we found a containing node, it means the item fit on the root or one of its sub-
		// nodes
		t.contents_node[item] = containingNode
	} else {
		// If the item couldn't be added to the the root node, it's because it wouldn't fit.  We'll need
		// to build bigger nodes until we get one that will fit the item.
		compareNode := t.root
		card := t.root.bounds.Cardinality()
		for {
			// Find which way to expand for each dimension.  If the item's region is within bounds for
			// any particular dimension, we arbitrarily choose to expand closer to 0.0 for that
			// dimension.
			pt0, pt1 := make([]float64, card), make([]float64, card)
			for d := 0; d < card; d++ {
				i_min, i_max := itemRegion.GetMin(d), itemRegion.GetMax(d)
				c_min, c_max := compareNode.bounds.GetMin(d), compareNode.bounds.GetMax(d)

				if i_min < c_min {
					pt0[d] = c_min - (c_max - c_min)
					pt1[d] = c_max
				} else if c_max < i_max {
					pt0[d] = c_min
					pt1[d] = c_max + (c_max - c_min)
				} else {
					if math.Abs(c_min) < math.Abs(c_max) {
						pt0[d] = c_min - (c_max - c_min)
						pt1[d] = c_max
					} else {
						pt0[d] = c_min
						pt1[d] = c_max + (c_max - c_min)
					}
				}
			}

			// We now know the dimensions of the containing node.  Create it and assign the compareNode
			// as a subnode.
			newRegion := NewRegion(pt0, pt1)
			newNode := nTreeNode[T]{
				bounds: newRegion.(region),
				depth:  compareNode.depth - 1,
				parent: compareNode,
			}
			subIdx, _ := newNode.bounds.getContainingSubRegion(compareNode.bounds)
			if subIdx >= 0 {
				newNode.subs[subIdx] = compareNode
				compareNode = &newNode
			}

			// Check if the newly-created node contains the given item.  If so, we're done.  If not,
			// we'll create an even bigger node containing this new one.
			if newRegion.GetContains(itemRegion) {
				newNode.items = append(newNode.items, item)
				t.root = &newNode
				return nil
			}
		}
	}
	return nil
}
func (t *NTree[T]) Contains(item any) bool { _, contained := t.contents_node[item]; return contained }
func (t *NTree[T]) GetBounds() Region      { return t.root.bounds }
func (t *NTree[T]) GetCardinality() int    { return t.cardinality }
func (t *NTree[T]) GetCount() int          { return t.count }
func (t *NTree[T]) GetIntersections(region Region) []any {
	nodes := []*nTreeNode[T]{t.root}
	result := make([]any, 0)
	for len(nodes) > 0 {
		// pop
		n := nodes[0]
		nodes = nodes[1:]

		i := n.bounds.GetIntersection(region)
		if i == nil {
			continue
		} else {
			// Get intersectors at this node.
			for _, item := range n.items {
				if itemBounds, ok := t.contents_region[item]; ok && itemBounds.GetIntersection(region) != nil {
					result = append(result, item)
				}
			}
			// Append the sub-nodes to the nodes stack
			for _, sub := range n.subs {
				if sub != nil {
					nodes = append(nodes, sub)
				}
			}
		}
	}
	return result
}
func (t *NTree[T]) Remove(item T) error {
	if node, ok := t.contents_node[item]; !ok {
		return errItemNotContained
	} else {
		// Delete from the tree' registry, and from the node's list.
		delete(t.contents_node, item)
		delete(t.contents_region, item)
		idx := -1
		for i, compare := range node.items {
			if compare == item {
				idx = i
				break
			}
		}
		node.items = append(node.items[:idx], node.items[idx+1:]...)

		// Now check if the node is empty, and remove it from the parent if so (and so on up the
		// chain).
		for node.isEmpty() && node.parent != nil {
			parent := node.parent
			i, _ := parent.bounds.getContainingSubRegion(node.bounds)
			parent.subs[i] = nil
			node.parent = nil
			node = parent
		}
		t.count -= 1
		return nil
	}
}

func NewNTree[T comparable](cardinality int) *NTree[T] {
	return &NTree[T]{cardinality: cardinality}
}

type nTreeNode[T comparable] struct {
	subs   []*nTreeNode[T]
	parent *nTreeNode[T]
	bounds region
	items  []any
	depth  int
}

// Returns the node that contains the given object after the 'add' operation is done.
func (n *nTreeNode[T]) add(item any, itemBounds Region) *nTreeNode[T] {
	if !n.bounds.GetContains(itemBounds) {
		// The item doesn't fit.  The caller will have to create a larger node
		return nil
	} else if len(n.items) < maxNodeItems {
		// The item fits and can go into the current collection
		n.items = append(n.items, item)
		return n
	} else {
		// The item might be able to go into a subregion of this node
		i, subRegion := n.bounds.getContainingSubRegion(itemBounds)
		if subRegion != nil {
			newNode := nTreeNode[T]{
				bounds: subRegion.(region),
				depth:  n.depth + 1,
			}
			n.subs[i] = &newNode
			return newNode.add(item, itemBounds)
		}

		// If none of the subregions could take it, item MUST be added at this node.
		n.items = append(n.items, item)
		return n
	}
}
func (n *nTreeNode[T]) isEmpty() bool {
	if len(n.items) > 0 {
		return false
	} else {
		for _, sub := range n.subs {
			if sub != nil && !sub.isEmpty() {
				return false
			}
		}
	}
	return true
}
