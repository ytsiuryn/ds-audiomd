package metadata

import (
	"testing"
)

func TestAlbumPerformers(t *testing.T) {
	r := NewRecord()
	r.ActorRoles["Miles Davis"] = []string{"performer"}
	r.ActorRoles["Marcus Miller"] = []string{"guitar"}
	r.ActorRoles["Milt Jackson"] = []string{"performer"}
	if len(r.Performers()) != 2 {
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
