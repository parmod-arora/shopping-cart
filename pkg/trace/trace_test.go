package trace

import (
	"testing"
)

func TestTrace(t *testing.T) {
	tests := []struct {
		name string
		want Info
	}{
		{
			name: "TraceInfo Test",
			want: Info{
				Line:         24,
				FunctionName: "test_functionName",
			},
		},
	}
	for _, tt := range tests {
		testFunc:= func(t *testing.T) {
			if got := Trace(); got.Line == tt.want.Line && got.FunctionName == tt.want.FunctionName {
				t.Errorf("Trace() = %v, want %v", got, tt.want)
			}
		}
		t.Run(tt.name, testFunc)
	}
}
