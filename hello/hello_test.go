package hello

import (
	"testing"
)

func TestSum(t *testing.T) {
	hi := hello()
	if hi != "hello" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", hi, "hello")
	}
}
