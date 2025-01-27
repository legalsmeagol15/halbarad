package helpers

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func Test_region_GetMin(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		dimension int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		struct {
			name   string
			fields fields
			args   args
			want   float64
		}{
			name: "firstTest",
			fields: fields{
				points: mat.NewDense(2, 2, []float64{1, 1, 2, 2}),
			},
		},
		struct {
			name   string
			fields fields
			args   args
			want   float64
		}{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetMin(tt.args.dimension); got != tt.want {
				t.Errorf("region.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetMax(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		dimension int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetMax(tt.args.dimension); got != tt.want {
				t.Errorf("region.GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetPoints(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   mat.Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("region.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_Cardinality(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetCardinality(); got != tt.want {
				t.Errorf("region.Cardinality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_Area(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetArea(); got != tt.want {
				t.Errorf("region.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_Perimeter(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetPerimeter(); got != tt.want {
				t.Errorf("region.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetContains(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		other Region
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetContains(tt.args.other); got != tt.want {
				t.Errorf("region.GetContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetIntersection(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		other Region
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Region
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetIntersection(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("region.GetIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetUnion(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		other Region
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Region
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetUnion(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("region.GetUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_getSubRegion(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Region
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.getSubRegion(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("region.getSubRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_getContainingSubRegion(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	type args struct {
		bounds Region
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  Region
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			got, got1 := r.getContainingSubRegion(tt.args.bounds)
			if got != tt.want {
				t.Errorf("region.getContainingSubRegion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("region.getContainingSubRegion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewRegion(t *testing.T) {
	type args struct {
		pt0 []float64
		pt1 []float64
	}
	tests := []struct {
		name string
		args args
		want Region
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegion(tt.args.pt0, tt.args.pt1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
