package pathcompression_test

import (
	"testing"

	"github.com/disiqueira/algorithms/quickunion/weighted/pathcompression"
)

func TestQuickFind(t *testing.T) {
	q := pathcompression.New(10)
	q.Union(1, 2)
	if !q.Connected(1, 2) {
		t.Error("Should be true")
	}
	if !q.Connected(2, 1){
		t.Error("Should be true")
	}
	if q.Connected(1, 9){
		t.Error("Should be false")
	}
	q.Union(4, 5)
	q.Union(5, 6)
	q.Union(2, 6)
	if !q.Connected(4, 6){
		t.Error("Should be true")
	}
	if q.Connected(8, 9){
		t.Error("Should be false")
	}
}
