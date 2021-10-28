package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	m := RecordingIDs{0: ""}
	data, err := json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"":""}`, string(data))
	m = RecordingIDs{ISRC: "12345"}
	data, err = json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"isrc":"12345"}`, string(data))
}

func TestRecordingIDsUnmarshal(t *testing.T) {
	m := RecordingIDs{}
	jsonData := []byte(`{"unknown": 0}`)
	assert.Error(t, json.Unmarshal(jsonData, &m))
	jsonData = []byte(`{"isrc": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	require.NoError(t, err)
	assert.Contains(t, m, ISRC)
}
