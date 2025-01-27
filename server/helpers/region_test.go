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

func Test_region_GetArea(t *testing.T) {
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

func Test_region_checkCardinality(t *testing.T) {
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
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := region{
				points: tt.fields.points,
			}
			if got := r.checkCardinality(tt.args.other); got != tt.want {
				t.Errorf("region.checkCardinality() = %v, want %v", got, tt.want)
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
