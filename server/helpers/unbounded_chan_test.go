package helpers

import (
	"reflect"
	"testing"
)

func TestNewUnboundedChan(t *testing.T) {
	type args struct {
		bfr int
	}
	tests := []struct {
		name       string
		args       args
		wantCh_in  chan<- int
		wantCh_out <-chan int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCh_in, gotCh_out := NewUnboundedChan[int](tt.args.bfr)
			if !reflect.DeepEqual(gotCh_in, tt.wantCh_in) {
				t.Errorf("NewUnboundedChan() gotCh_in = %v, want %v", gotCh_in, tt.wantCh_in)
			}
			if !reflect.DeepEqual(gotCh_out, tt.wantCh_out) {
				t.Errorf("NewUnboundedChan() gotCh_out = %v, want %v", gotCh_out, tt.wantCh_out)
			}
		})
	}
}
