package helpers

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Region interface {
	GetPoints() *mat.Matrix
	Area() float64
	Perimeter() float64
	GetIntersection(other Region) Region
	GetUnion(other Region) Region
}

type region struct {
	points mat.Matrix
}

func (r region) GetPoints() *mat.Matrix { return &r.points }
func (r region) Cardinality() int       { _, c := r.points.Dims(); return c }
func (r region) Area() float64 {
	if r.Cardinality() <= 1 {
		return 0
	} else {
		var area_volume float64 = 1.0
		for d := 0; d < r.Cardinality(); d++ {
			area_volume *= math.Abs(r.points.At(0, d) - r.points.At(1, d))
		}
		return area_volume
	}
}
func (r region) Perimeter() float64 {

	switch r.Cardinality() {
	case 0:
		return 0
	case 1:
		return 2 * VecLength(VecSubtract(mat.Row(nil, 0, r.points), mat.Row(nil, 1, r.points)))
	default:
		// TODO:  implement at least for cardinality = 2
		panic("not implemented")
	}
}

func (r region) GetIntersection(other Region) Region {
	panic("not implemented")
}
func (r region) GetUnion(other Region) Region {
	panic("not implemented")
}
