package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbumPerformers(t *testing.T) {
	r := NewRecord()
	r.ActorRoles["Miles Davis"] = []string{"performer"}
	r.ActorRoles["Marcus Miller"] = []string{"guitar"}
	r.ActorRoles["Milt Jackson"] = []string{"performer"}
	assert.Len(t, r.Performers(), 2)
}

func TestRecordingIsEmptyAndClean(t *testing.T) {
	r := NewRecord()
	r.Clean()
	assert.True(t, r.IsEmpty())
	assert.Empty(t, r.Actors)
	assert.Empty(t, r.IDs)
}
