package binarysearch_test

import (
	"github.com/disiqueira/algorithms/binarysearch"
	"testing"
)

func TestSearch(t *testing.T) {
	data := []int{6, 13, 14, 25, 33, 43, 51, 53, 64, 72, 84, 93, 95, 96, 97}
	if search := 33; !binarysearch.Search(search, data) {
		t.Errorf("The Search should have found %d in the list", search)
	}
	if search := 25; !binarysearch.Search(search, data) {
		t.Errorf("The Search should have found %d in the list", search)
	}
	if search := 27; binarysearch.Search(search, data) {
		t.Errorf("The Search shouldn't have found %d in the list", search)
	}
}
