// I am specifically avoiding github.com/go-spatial/geom because it is designed to focus on 2D
// geometry, and I can imagine someday needing 3D or even N-dimensional geometry.
package helpers

import "errors"

type Geometry interface {
	Area() float64
	Perimeter() float64
	Cardinality() int
}

type Point []float64

func (p Point) Area() float64      { return 0 }
func (p Point) Perimeter() float64 { return 0 }
func (p Point) Cardinality() int   { return len(p) }

type Region interface {
	PointA() Point
	PointB() Point

	GetIntersection(other Region) (Region, error)
}

type region struct {
	a, b Point
}

func (r region) PointA() Point    { return r.a }
func (r region) PointB() Point    { return r.b }
func (r region) Cardinality() int { return len(r.a) }

func NewRect(a, b Point) (Region, error) {
	if len(a) != len(b) {
		empty := Point{}
		return region{empty, empty}, errors.New("inconsistent cardinality")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return region{a, b}, nil
		}
	}
	empty := Point{}
	return region{empty, empty}, errors.New("point given")
}

func (r region) GetIntersection(other Region) (Region, error) {

}
