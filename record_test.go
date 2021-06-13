package metadata

import (
	"testing"
)

func TestAlbumPerformers(t *testing.T) {
	r := NewRecord()
	r.Actors.AddRole("Miles Davis", "performer")
	r.Actors.AddRole("Marcus Miller", "guitar")
	r.Actors.AddRole("Milt Jackson", "performer")
	if len(*r.Performers()) != 2 {
		t.Fail()
	}
}

func TestRecordingIsEmptyAndClean(t *testing.T) {
	r := NewRecord()
	r.Clean()
	if !r.IsEmpty() {
		t.Fail()
	}
	if r.Actors != nil || r.IDs != nil {
		t.FailNow()
	}
}
