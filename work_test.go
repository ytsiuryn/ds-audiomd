package metadata

import "testing"

func TestWorkIsEmptyAndClean(t *testing.T) {
	w := NewWork()
	w.Clean()
	if !w.IsEmpty() {
		t.Fail()
	}
	if w.Actors != nil || w.IDs != nil {
		t.Fail()
	}
}
