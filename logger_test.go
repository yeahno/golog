package golog

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	log := NewLogger(os.Stdout)
	log.Debug("Debug", "this is a string", []string{"string[0]", "string[1]"})
	log.Trace("Trace", "log a int[]", []int{0, 1})
	log.Info("Info", "test map", map[string]string{"string": "string_value1"})
	log.Warning("Warning", true)
	log.Error("Error", 1, 2, 3, 4, 5, "string")

	log.Disable(DEBUG)
	log.Debug("debug not show", "test", "debug", []string{"string1", "string2"}, log)
	log.Enable(DEBUG)
	log.Debug("debug ought show", "test", "debug", []string{"string1", "string2"}, log)
}
