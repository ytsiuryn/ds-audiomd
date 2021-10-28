package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkIsEmptyAndClean(t *testing.T) {
	w := NewWork()
	w.Clean()
	assert.True(t, w.IsEmpty())
	assert.Empty(t, w.Actors)
	assert.Empty(t, w.IDs)
}

func TestWorkIDsMarshal(t *testing.T) {
	m := WorkIDs{0: ""}
	data, err := json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"":""}`, string(data))
	m = WorkIDs{MusicbrainzWorkID: "12345"}
	data, err = json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"musicbrainz_work_id":"12345"}`, string(data))
}

func TestWorkIDsUnmarshal(t *testing.T) {
	m := WorkIDs{}
	jsonData := []byte(`{"unknown": 0}`)
	assert.Error(t, json.Unmarshal(jsonData, &m))
	jsonData = []byte(`{"musicbrainz_work_id": "12345"}`)
	require.NoError(t, json.Unmarshal(jsonData, &m))
	assert.Contains(t, m, MusicbrainzWorkID)
}
