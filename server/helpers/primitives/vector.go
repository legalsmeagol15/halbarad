package primitives

import "math"

type Vector []float64

func (v Vector) getLengthSquared() float64 {
	sum := 0.0
	for _, s := range v {
		sum += (s * s)
	}
	return sum
}
func (v Vector) GetUnit() Vector {
	length := v.GetLength()
	for i, s := range v {
		v[i] = s / length
	}
	return v
}
func (v Vector) GetLength() float64 { return math.Sqrt(v.getLengthSquared()) }
