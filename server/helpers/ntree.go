package helpers

import "errors"

var (
	nonContainError error = errors.New("node cannot contain error")
	maxNodeItems    int   = 32
	maxNodeCapacity int   = 32
)

type NTree[T comparable] interface {
	Add(item T, region Region) error
	Remove(item T) error
	GetIntersections(region Region) chan T
	GetRegion(item T) Region
	GetBoundary() Region
	GetCardinality() int
}
type nTree[T comparable] struct {
	root        *nTreeNode[T]
	cardinality int
	contents    map[T]Region
}

func (t *nTree[T]) Add(item T, itemRegion Region) error {
	if !itemRegion.IsDefined() {
		return errors.New("region not completely defined.")
	}
	if region.PointA().Cardinality() > t.cardinality {
		return errors.New("inconsistent cardinality")
	} else if _, ok := t.contents[item]; ok {
		return errors.New("duplicate item")
	}

	var n *nTreeNode[T] = t.root
	if !n.bounds.Contains(itemRegion) {

	}

}
func (t *nTree[T]) GetBoundary() Region { return t.root.bounds }
func (t *nTree[T]) GetCardinality() int { return t.cardinality }

func (t *nTree[T]) Remove(item T) error {
	if _, ok := t.contents[item]; !ok {
		return errors.New("item does not exist")
	}
	delete(t.contents, item)
	return nil
}

func NewNTree[T comparable](cardinality int) nTree[T] {
	return nTree[T]{cardinality: cardinality}
}

type nTreeNode[T comparable] struct {
	quad1, quad2, quad3, quad4 *nTreeNode[T]
	bounds                     region
	items                      []T
	depth                      int
}

func (n *nTreeNode[T]) add(item T, itemBounds region) error {
	if !n.bounds.GetContains(itemBounds) {
		return nonContainError
	} else if len(n.items) < maxNodeItems || n.depth > maxNodeCapacity {
		n.items = append(n.items, item)
		return nil
	}
}
