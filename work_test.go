package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkIsEmptyAndClean(t *testing.T) {
	w := NewWork()
	w.Clean()
	assert.True(t, w.IsEmpty())
	assert.Empty(t, w.Actors)
	assert.Empty(t, w.IDs)
}

func TestWorkIDsMarshal(t *testing.T) {
	m := WorkIDs{MusicbrainzWorkID: "12345"}
	data, err := json.Marshal(m)
	assert.Equal(t, `{"musicbrainz_work_id":"12345"}`, string(data))
	assert.NoError(t, err)
}

func TestWorkIDsUnmarshal(t *testing.T) {
	m := WorkIDs{}
	jsonData := []byte(`{"musicbrainz_work_id": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	assert.NoError(t, err)
	assert.Contains(t, m, MusicbrainzWorkID)
}
