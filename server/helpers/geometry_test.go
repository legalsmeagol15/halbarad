// I am specifically avoiding github.com/go-spatial/geom because it is designed to focus on 2D
// geometry, and I can imagine someday needing 3D or even N-dimensional geometry.
// Author Wesley Oates
package helpers

import (
	"reflect"
	"testing"
)

func TestVecSubtract(t *testing.T) {
	type args struct {
		b []float64
		a []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VecSubtract(tt.args.b, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VecSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVecLengthSquared(t *testing.T) {
	type args struct {
		v []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VecLengthSquared(tt.args.v); got != tt.want {
				t.Errorf("VecLengthSquared() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVecLength(t *testing.T) {
	type args struct {
		v []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VecLength(tt.args.v); got != tt.want {
				t.Errorf("VecLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
