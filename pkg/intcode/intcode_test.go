package intcode

import (
	"fmt"
	"reflect"
	"testing"
)

var runData = []struct {
	initial Program
	final   Memory
}{
	{
		Program{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		Memory{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
	},
	{
		Program{1, 0, 0, 0, 99},
		Memory{2, 0, 0, 0, 99},
	},
	{
		Program{2, 3, 0, 3, 99},
		Memory{2, 3, 0, 6, 99},
	},
	{
		Program{2, 4, 4, 5, 99, 0},
		Memory{2, 4, 4, 5, 99, 9801},
	},
	{
		Program{1, 1, 1, 4, 99, 5, 6, 0, 99},
		Memory{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
}

func TestRun(t *testing.T) {
	for _, tt := range runData {
		t.Run(fmt.Sprintf("test initial %v", tt.initial), func(t *testing.T) {
			m := New()
			m.Load(tt.initial)
			m.Run()
			final := m.Memory()
			if !reflect.DeepEqual(final, tt.final) {
				t.Errorf("incorrect final, got %v want %v", final, tt.final)
			}
		})
	}
}
