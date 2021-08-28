package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		input  []interface{}
		expect []interface{}
	}{
		{
			[]interface{}{1, 2, 3},
			[]interface{}{1, 2, 3},
		},
		{
			[]interface{}{
				1, 2, []string{"a", "b", "c"},
			},
			[]interface{}{
				1, 2, "a", "b", "c",
			},
		},
		{
			[]interface{}{
				1, 2, []interface{}{"a", "b", "c", []int{10, 20, 30}, []string{"d", "e", "f"}},
			},
			[]interface{}{
				1, 2, "a", "b", "c", 10, 20, 30, "d", "e", "f",
			},
		},
	}

	for _, tt := range tests {
		var acc []interface{}
		Flatten(&acc, tt.input)

		if !reflect.DeepEqual(acc, tt.expect) {
			t.Errorf("expected: %v, got: %v", tt.expect, acc)
		}
	}
}
