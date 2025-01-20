package helpers

import "golang.org/x/exp/constraints"

type Matrix[T constraints.Float] interface {
	Translate(v Vector) Matrix[T]
	Scale(v Vector) Matrix[T]
	Rotate(v Vector) Matrix[T]
	Get(coords ...uint16) T
	Set(value T, coords ...uint16)
}

type matrix[T constraints.Float] struct {
	contents []T
	offsets  []uint16
}

func NewMatrix[T constraints.Float](dimensions ...uint16) Matrix[T] {
	if len(dimensions) < 1 {
		return matrix[T]{}
	}
	reserve := uint16(1)
	offsets := make([]uint16, len(dimensions), len(dimensions))
	for i, d := range dimensions {
		offsets[len(dimensions)-(1+i)] = reserve
		reserve *= d
	}
	result := matrix[T]{
		offsets:  offsets,
		contents: make([]T, reserve, reserve),
	}
	return result
}

func (m *matrix[T]) getOffset(coords []uint16) uint16 {
	sum_c := uint16(0)
	for i, c := range coords {
		sum_c += c * m.offsets[i]
	}
	return sum_c
}
func (m matrix[T]) Get(coords ...uint16) T        { return m.contents[m.getOffset(coords)] }
func (m matrix[T]) Set(value T, coords ...uint16) { m.contents[m.getOffset(coords)] = value }

func GetRotator[T constraints.Float](rotations ...T) Matrix[T] {

}
func GetTranslator[T constraints.Float](translations ...T) Matrix[T] {

}
func GetScalor[T constraints.Float](scalors ...T) Matrix[T] {

}
