package helpers

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Region interface {
	GetPoints() mat.Matrix
	Area() float64
	Perimeter() float64
	GetIntersection(other Region) Region
	GetUnion(other Region) Region
	GetContains(other Region) bool
}

type region struct {
	points mat.Matrix
}

func (r region) GetPoints() mat.Matrix { return r.points }
func (r region) Cardinality() int      { _, c := r.points.Dims(); return c }
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
func (r region) GetContains(other Region) bool {
	_, card := other.GetPoints().Dims()
	if r.Cardinality() != card {
		panic("inconsistent cardinality")
	}
	for i := 0; i < card; i++ {
		a := mat.Col(nil, i, r.points)
		b := mat.Col(nil, i, other.GetPoints())
		if b[0] < a[0] || a[1] < b[1] {
			return false
		}
	}
	return true
}
func (r region) GetIntersection(other Region) Region {
	_, card := other.GetPoints().Dims()
	if r.Cardinality() != card {
		panic("inconsistent cardinality")
	}
	res_a, res_b := make([]float64, card), make([]float64, card)

	for i := 0; i < card; i++ {
		a := mat.Col(nil, i, r.points)
		b := mat.Col(nil, i, other.GetPoints())

		if a[0] <= b[0] {
			if a[1] < b[0] {
				return nil
			}
			res_a[i], res_b[i] = b[0], min(a[1], b[1])
		} else if b[1] < a[0] {
			return nil
		} else {
			res_a[i], res_b[i] = max(b[1], a[0]), min(b[1], a[1])
		}
	}
	return NewRegion(res_a, res_b)
}
func (r region) GetUnion(other Region) Region {
	_, card := other.GetPoints().Dims()
	if r.Cardinality() != card {
		panic("inconsistent cardinality")
	}
	res_a, res_b := make([]float64, card), make([]float64, card)
	for i := 0; i < card; i++ {
		a := mat.Col(nil, i, r.points)
		b := mat.Col(nil, i, other.GetPoints())
		res_a[i], res_b[i] = min(a[0], b[0]), max(a[1], b[1])
	}
	return NewRegion(res_a, res_b)
}
func (r region) equals(other region) bool { return mat.Equal(r.points, other.points) }
func (r region) newQuadrants() []region {
	center, a1, b3 := make([]float64, r.Cardinality())
	a, b := mat.Row(nil, 0, r.points), mat.Row(nil, 1, r.points)
	for i := 1; i < r.Cardinality(); i++ {
		center[i] = (a[i] + b[i]) / 2
	}
	return []region {
		NewRegion(center, 
	}
}
func NewRegion(pt0, pt1 []float64) Region {
	if len(pt0) != len(pt1) {
		panic("inconsistent cardinality")
	}
	min_pt, max_pt := make([]float64, len(pt0)), make([]float64, len(pt0))
	for i, m0 := range pt0 {
		m1 := pt1[i]
		if m0 <= m1 {
			min_pt[i], max_pt[i] = m0, m1
		} else {
			min_pt[i], max_pt[i] = m1, m0
		}
	}
	data := append(min_pt, max_pt...)
	return region{
		points: mat.NewDense(2, len(pt0), data),
	}
}
