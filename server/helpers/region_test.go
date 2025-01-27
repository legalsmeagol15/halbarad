package helpers

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

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
		{
			name: "Instantiation Test",
			args: args{[]float64{1, 1}, []float64{2, 2}},
			want: region{points: mat.NewDense(2, 2, []float64{1, 1, 2, 2})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegion(tt.args.pt0, tt.args.pt1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetCardinality(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Cardinality Test 1",
			fields: fields{points: mat.NewDense(2, 1, []float64{1, 2})},
			want:   1,
		},
		{
			name:   "Cardinality Test 2",
			fields: fields{points: mat.NewDense(2, 2, []float64{1, 1, 2, 2})},
			want:   2,
		},
		{
			name:   "Cardinality Test 3",
			fields: fields{points: mat.NewDense(2, 3, []float64{1, 1, 1, 2, 2, 2})},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetCardinality(); got != tt.want {
				t.Errorf("region.GetCardinality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetMin(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   []float64
	}{
		{
			name:   "GetMin2",
			fields: fields{points: mat.NewDense(2, 2, []float64{1, 1, 2, 2})},
			want:   []float64{1, 1},
		},
		{
			name:   "GetMin3",
			fields: fields{points: mat.NewDense(2, 3, []float64{1, 1, 1, 2, 2, 2})},
			want:   []float64{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetMin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("region.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetMax(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   []float64
	}{
		{
			name:   "GetMax2",
			fields: fields{points: mat.NewDense(2, 2, []float64{1, 1, 2, 2})},
			want:   []float64{2, 2},
		},
		{
			name:   "GetMax3",
			fields: fields{points: mat.NewDense(2, 3, []float64{1, 1, 1, 2, 2, 2})},
			want:   []float64{2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetMax(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("region.GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_region_GetPoints(t *testing.T) {
	// This tests a trivially simple function.  Yay for coverage.
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   mat.Matrix
	}{
		{
			name:   "GetPoints2",
			fields: fields{mat.NewDense(2, 2, []float64{1, 1, 2, 2})},
			want:   mat.NewDense(2, 2, []float64{1, 1, 2, 2}),
		},
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

func Test_region_GetArea(t *testing.T) {
	type fields struct {
		points mat.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "GetArea2a",
			fields: fields{points: mat.NewDense(2, 2, []float64{1, 1, 2, 2})},
			want:   1.0,
		},
		{
			name:   "GetArea2b",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 2, 2})},
			want:   4.0,
		},
		{
			name:   "GetArea3a",
			fields: fields{points: mat.NewDense(2, 3, []float64{1, 1, 1, 2, 2, 2})},
			want:   1.0,
		},
		{
			name:   "GetArea3b",
			fields: fields{points: mat.NewDense(2, 3, []float64{1, 1, 1, 3, 3, 3})},
			want:   8.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.GetArea(); got != tt.want {
				t.Errorf("region.GetArea() = %v, want %v", got, tt.want)
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
		{
			name:   "GetContains2a - contained",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{9, 9})},
			want:   true,
		},
		{
			name:   "GetContains2b - contained and border shared",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{10, 10})},
			want:   true,
		},
		{
			name:   "GetContains2c - not contained. overlapped",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{11, 11})},
			want:   false,
		},
		{
			name:   "GetContains2d - not contained and border shared",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{10, 10}, []float64{12, 12})},
			want:   false,
		},
		{
			name:   "GetContains2e - not contained and completely distinct",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 5, 5})},
			args:   args{other: NewRegion([]float64{6, 6}, []float64{10, 10})},
			want:   false,
		},
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
		{
			name:   "GetIntersection2a - contained",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{9, 9})},
			want:   NewRegion([]float64{1, 1}, []float64{9, 9}),
		},
		{
			name:   "GetIntersection2b - contained and border shared",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{10, 10})},
			want:   NewRegion([]float64{1, 1}, []float64{10, 10}),
		},
		{
			name:   "GetIntersection2c - not contained. overlapped",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{11, 11})},
			want:   NewRegion([]float64{1, 1}, []float64{10, 10}),
		},
		{
			name:   "GetIntersection2d - not contained and border shared",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{10, 10}, []float64{12, 12})},
			want:   NewRegion([]float64{10, 10}, []float64{10, 10}),
		},
		{
			name:   "GetIntersection2e - not contained and completely distinct",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 5, 5})},
			args:   args{other: NewRegion([]float64{6, 6}, []float64{10, 10})},
			want:   nil,
		},
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

func Test_region_GetPerimeter(t *testing.T) {
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
				t.Errorf("region.GetPerimeter() = %v, want %v", got, tt.want)
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
		{
			name:   "GetUnion2a - contained",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{9, 9})},
			want:   NewRegion([]float64{0, 0}, []float64{10, 10}),
		},
		{
			name:   "GetUnion2b - contained and border shared",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{10, 10})},
			want:   NewRegion([]float64{0, 0}, []float64{10, 10}),
		},
		{
			name:   "GetUnion2c - not contained. overlapped",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{1, 1}, []float64{11, 11})},
			want:   NewRegion([]float64{0, 0}, []float64{11, 11}),
		},
		{
			name:   "GetUnion2d - not contained and border shared",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 10, 10})},
			args:   args{other: NewRegion([]float64{10, 10}, []float64{12, 12})},
			want:   NewRegion([]float64{0, 0}, []float64{12, 12}),
		},
		{
			name:   "GetUnion2e - not contained and completely distinct",
			fields: fields{points: mat.NewDense(2, 2, []float64{0, 0, 5, 5})},
			args:   args{other: NewRegion([]float64{6, 6}, []float64{10, 10})},
			want:   NewRegion([]float64{0, 0}, []float64{10, 10}),
		},
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

func Test_region_checkCardinality(t *testing.T) {

}
