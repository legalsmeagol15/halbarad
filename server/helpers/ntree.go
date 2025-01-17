package helpers

import "errors"

type NTree[T comparable] interface {
	Add(item T, region Region) error
	Remove(item T) error
	GetIntersections(region Region) chan T
	GetRegion(item T) Region
	GetBoundary() Region
	Cardinality() int
}
type nTree[T comparable] struct {
	bounds      region
	root        nTreeNode[T]
	cardinality int
	contents    map[T]Region
}
type nTreeNode[T comparable] struct {
	quad1, quad2, quad3, quad4 *nTreeNode[T]
}

func (t *nTree[T]) Add(item T, region Region) error {
	if _, ok := t.contents[item]; ok {
		return errors.New("duplicate item")
	}
}
func (t *nTree[T]) Cardinality() int { return t.cardinality }

func (t *nTree[T]) Remove(item T) error {
	if _, ok := t.contents[item]; !ok {
		return errors.New("item does not exist")
	}
	delete(t.contents, item)
	return nil
}

func NewnTree[T comparable](cardinality int) nTree[T] {
	return nTree[T]{cardinality: cardinality}
}
