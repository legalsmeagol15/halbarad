package primitives

import (
	"errors"
	"math"
)

var (
	error_dimension_mismatch = errors.New("dimension mismatch")
)

type Matrix interface {
	Get(coords ...int) (float64, error)
	Set(value float64, coords ...int) error
}

type matrix struct {
	dimensions []int
	items      []any
}

func (m matrix) Get(coords ...int) (float64, error) {
	items := m.items
	for _, c := range coords {
		item := items[c]
		if next, ok := item.([]any); ok {
			items = next
		} else {
			return item.(float64), nil
		}
	}
	return 0.0, error_dimension_mismatch
}
func (m matrix) Set(value float64, coords ...int) error {
	items := m.items
	for _, c := range coords {
		item := items[c]
		if next, ok := item.([]any); ok {
			items = next
		} else {
			items[c] = value
		}
	}
	return nil
}
func (m matrix) GetRotator2D(angle float64) (Matrix, error) {
	// Source: https://en.wikipedia.org/wiki/Rotation_matrix#In_two_dimensions
	if len(m.dimensions) != 2 {
		return nil, error_dimension_mismatch
	}
	c := math.Cos(angle)
	s := math.Sin(angle)
	result := matrix{
		dimensions: []int{2, 2},
		items: []any{[]float64{c, -s},
			[]float64{s, c},
		},
	}
	return result, nil
}
func (m matrix) GetRotator3D(axis Vector, angle float64) (Matrix, error) {
	// Source: https://en.wikipedia.org/wiki/Rotation_matrix#In_three_dimensions
	if len(m.dimensions) != 3 {
		return nil, error_dimension_mismatch
	}
	unit_axis := axis.GetUnit()
	ux, uy, uz := unit_axis[0], unit_axis[1], unit_axis[2]
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	less_cos := 1 - cos
	result := matrix{
		dimensions: []int{3, 3},
		items: []any{[]float64{(ux * ux * less_cos) + cos, (ux * uy * less_cos) - (uz * sin), (ux * uz * less_cos) + (uy * sin)},
			[]float64{(ux * uy * less_cos) + (uz * sin), (uy * uy * less_cos) + cos, (uy * uz * less_cos) - (ux * sin)},
			[]float64{(ux * uz * less_cos) - (uy * sin), (uy * uz * less_cos) - (ux * sin), (uz * uz * less_cos) + cos},
		},
	}
	return result, nil
}
func (m matrix) GetScalor(scalars ...float64) (Matrix, error) {
	dim := len(scalars)
	if len(m.dimensions) != dim {
		return nil, error_dimension_mismatch
	}
	result := matrix{
		dimensions: []int{dim, dim},
		items:      make([]any, dim, dim),
	}
	for i, s := range scalars {
		row := make([]any, dim, dim)
		row[i] = s
		result.items[i] = row
	}
	return result, nil
}
func (m matrix) GetTranslated(v Vector) (Matrix, error) {
	if len(m.dimensions) != len(v) {
		return nil, error_dimension_mismatch
	}
	result := matrix{
		dimensions: m.dimensions,
		items:      make([]any, len(v), len(v)),
	}
	for i, row := range m.items {

		new_row := make([]any, len(v), len(v))

		result.items[i] = new_row

	}
}
