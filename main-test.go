package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println(Sum(5, 5))
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
