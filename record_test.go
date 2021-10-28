package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbumPerformers(t *testing.T) {
	r := NewRecord()
	r.AddRole("Miles Davis", "performer")
	r.AddRole("Marcus Miller", "guitar")
	r.AddRole("Milt Jackson", "performer")
	assert.Len(t, r.Performers(), 2)
}

func TestRecordingIsEmptyAndClean(t *testing.T) {
	r := NewRecord()
	r.Clean()
	assert.True(t, r.IsEmpty())
	assert.Empty(t, r.Actors)
	assert.Empty(t, r.IDs)
}

func TestRecordingIDsMarshal(t *testing.T) {
	m := RecordingIDs{ISRC: "12345"}
	data, err := json.Marshal(m)
	assert.Equal(t, `{"isrc":"12345"}`, string(data))
	assert.NoError(t, err)
}

func TestRecordingIDsUnmarshal(t *testing.T) {
	m := RecordingIDs{}
	jsonData := []byte(`{"isrc": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	assert.NoError(t, err)
	assert.Contains(t, m, ISRC)
}
