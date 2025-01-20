// I am specifically avoiding github.com/go-spatial/geom because it is designed to focus on 2D
// geometry, and I can imagine someday needing 3D or even N-dimensional geometry.
// Author Wesley Oates
package helpers

import (
	"errors"
	"math"
)

type GeomProperty uint16
type GeomIntersection uint16

const (
	VECTOR GeomProperty = 1 << iota
	REGION
	COTERMINAL
	COLINEAR
	COPLANAR
)

const (
	NO_OVERLAP GeomIntersection = 1 << iota
	OVERLAP_EAST
	OVERLAP_NORTH
	OVERLAP_WEST
	OVERLAP_SOUTH
	CONTAINER = OVERLAP_EAST | OVERLAP_NORTH | OVERLAP_SOUTH | OVERLAP_WEST
)

type Geometry interface {
	Cardinality() int
	Simplify() Geometry
	IsDefined() bool
	GetControlPoints() []Point
}

type Point []float64

func (p Point) Cardinality() int          { return len(p) }
func (p Point) Simplify() Geometry        { return p }
func (p Point) GetControlPoints() []Point { return []Point{p} }
func (p Point) Diff(other Point) Vector {
	result := make([]float64, len(p), len(p))
	for d := 0; d < len(p); d++ {
		result[d] = p[d] - other[d]
	}
	return result
}
func (p Point) Add(other Vector) Point {
	result := make([]float64, len(p), len(p))
	for d := 0; d < len(p); d++ {
		result[d] = p[d] + other[d]
	}
	return result
}
func (p Point) IsDefined() bool {
	for _, d := range p {
		if math.IsNaN(d) || math.IsInf(d, -1) || math.IsInf(d, 1) {
			return false
		}
	}
	return true
}
func (p Point) Scale(factor float64) Point {
	result := Point{}
	for d := 0; d < p.Cardinality(); d++ {
		result[d] = p[d] * factor
	}
	return result
}
func (p Point) Negate() Point { return p.Scale(-1.0) }

type Vector Point

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
	IsDefined() bool
}

type region struct {
	a, b Point
}

func (r region) PointA() Point             { return r.a }
func (r region) PointB() Point             { return r.b }
func (r region) Cardinality() int          { return len(r.a) }
func (r region) GetControlPoints() []Point { return []Point{r.a, r.b} }
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
func (r region) IsDefined() bool { return r.a.IsDefined() && r.b.IsDefined() }
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
func getOr(start_a, end_a, start_b, end_b float64) (float64, float64, bool) {
	if start_b < start_a {
		return getOr(start_b, end_b, start_a, end_a)
	} else if start_b > end_a {
		return 0, 0, false
	} else {
		return start_a, max(end_a, end_b), true
	}
}
func getAnd(start_a, end_a, start_b, end_b float64) (float64, float64, bool) {
	if start_b < start_a {
		return getAnd(start_b, end_b, start_a, end_a)
	} else if start_b > end_a {
		return 0, 0, false
	} else {
		return start_b, min(end_b, end_a), true
	}
}
func (r region) GetIntersection(other Region) (Region, bool) {
	result := region{}
	for d := 0; d < r.Cardinality(); d++ {
		var overlap bool
		if result.a[d], result.b[d], overlap = getAnd(r.a[d], r.b[d], other.PointA()[d], other.PointB()[d]); !overlap {
			return nil, false
		}
	}
	return result, true
}
func (r region) GetUnion(other region) (Region, bool) {
	result := region{}
	for d := 0; d < r.Cardinality(); d++ {
		var overlap bool
		if result.a[d], result.b[d], overlap = getOr(r.a[d], r.b[d], other.PointA()[d], other.PointB()[d]); !overlap {
			return nil, false
		}
	}
	return result, true
}
func (r region) GetScaled(factor float64, factors []float64) Region {
	// TODO:  this ought to be a standard linear algebra rotation-then-translate, not this custom thing.
	for d := 0; d < r.Cardinality(); d++ {
		scale := r.b[d] - r.a[d]
		f := factors[d]
		if f > 0.0 {
			r.b[d] = r.a[d] + (scale * f)
		} else if f < 0.0 {
			r.a[d] = r.b[d] - (scale * f)
		} else {
			r.a[d] = (r.a[d] + r.b[d]) / 2.0
			r.b[d] = r.a[d]
		}
	}
	return r
}
