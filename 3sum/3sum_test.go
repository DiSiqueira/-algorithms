package sum_test

import (
	"testing"

	"github.com/disiqueira/algorithms/3sum"
)

func TestThreeSum(t *testing.T) {
	d := []int{30,-40, -20, -10, 40, 0, 10, 5}
	result := sum.ThreeSum(d)
	expected := 4

	if result != expected {
		t.Errorf("Expected: %d, result: %d", expected, result)
	}
}