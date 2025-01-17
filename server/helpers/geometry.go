// I am specifically avoiding github.com/go-spatial/geom because it is designed to focus on 2D
// geometry, and I can imagine someday needing 3D or even N-dimensional geometry.
// Author Wesley Oates
package helpers

import (
	"errors"
	"math"
)

type GeomProperty uint16

const (
	VECTOR GeomProperty = 1 << iota
	REGION
	COTERMINAL
	COLINEAR
	COPLANAR
)

type Geometry interface {
	Cardinality() int
	Simplify() Geometry
}

type Point []float64

func NewPoint(p ...float64) Point {
	// This is a dumb method, but it helps to make the code more readable
	return p
}

func (p Point) Cardinality() int   { return len(p) }
func (p Point) Simplify() Geometry { return p }
func (p Point) Diff(other Point) Vector {
	result := make([]float64, len(p), len(p))
	for d := 0; d < len(p); d++ {
		result[d] = p[d] - other[d]
	}
	return result
}

type Vector Point

func NewVector(v ...float64) Point {
	// This is a dumb method, but it helps to make the code more readable
	return v
}

func (v Vector) LengthSquared() float64 {
	var sum float64 = 0
	for _, c := range v {
		sum += (c * c)
	}
	return sum
}
func (v Vector) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

type Region interface {
	PointA() Point
	PointB() Point

	Area() float64
	Perimeter() float64
	GetIntersection(other Region) (Region, bool)
}

type region struct {
	a, b Point
}

func (r region) PointA() Point    { return r.a }
func (r region) PointB() Point    { return r.b }
func (r region) Cardinality() int { return len(r.a) }
func (r region) Area() float64 {
	if len(r.a) <= 1 {
		return 0
	}
	var area_volume float64 = 1.0
	for d := 0; d < len(r.a); d++ {
		area_volume *= math.Abs(r.a[d] - r.b[d])
	}
	return area_volume
}
func (r region) Perimeter() float64 {
	switch len(r.a) {
	case 0:
		return 0
	case 1:
		return math.Abs(r.a[0] - r.b[0])
	default:
		// TODO:  implement at least for cardinality = 2
		panic("not implemented")
	}

}
func (r region) Simplify() Geometry {
	if r.Area() == 0 {
		return r.a
	}
	return r
}

func NewRect(a, b Vector) (Region, GeomProperty, error) {
	if len(a) != len(b) {
		empty := Point{}
		return region{empty, empty}, 0, errors.New("inconsistent cardinality")
	}

	is_Vector := true

	// The defining Vectors should be in smaller, larger order for each dimension.
	new_a := make([]float64, len(a), len(a))
	new_b := make([]float64, len(a), len(a))
	for d := 0; d < len(a); d++ {
		this_a, this_b := a[d], b[d]
		if this_a <= this_b {
			new_a[d] = this_a
			new_b[d] = this_b
		} else {
			new_a[d] = this_b
			new_b[d] = this_a
		}
		is_Vector = is_Vector && (this_a == this_b)
	}
	var p GeomProperty = REGION
	if is_Vector {
		p = VECTOR
	}

	return region{new_a, new_b}, p, nil
}

func getOverlap(start_a, end_a, start_b, end_b float64) (float64, float64, GeomProperty) {
	if start_b < start_a {
		return getOverlap(start_b, end_b, start_a, end_a)
	} else if start_b > end_a {
		return -1, -1, 0
	} else {
		relation := COPLANAR
		if end_a == start_b {
			relation = COTERMINAL
		}
		return start_b, min(end_a, start_b), relation
	}
}
func (r region) GetIntersection(other Region) (Region, bool) {
	result := region{}
	for d := 0; d < r.Cardinality(); d++ {
		var overlap GeomProperty
		result.a[d], result.b[d], overlap = getOverlap(r.a[d], r.b[d], other.PointA()[d], other.PointB()[d])
		if overlap == 0 {
			return nil, false
		}
	}
	return result, true
}
func (r region) GetContains(other Region) bool {
	otherCard := other.PointA().Cardinality()
	if otherCard > r.Cardinality() {
		return false
	}
	for d := 0; d < otherCard; d++ {
		if r.a[d] > other.PointA()[d] || r.b[d] < other.PointB()[d] {
			return false
		}
	}
	return true
}
