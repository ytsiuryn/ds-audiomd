package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLabelIDsMarshal(t *testing.T) {
	m := LabelIDs{0: ""}
	data, err := json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"":""}`, string(data))
	m = LabelIDs{MusicbrainzLabelID: "12345"}
	data, err = json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"musicbrainz_label_id":"12345"}`, string(data))
}

func TestLabelIDsUnmarshal(t *testing.T) {
	m := LabelIDs{}
	jsonData := []byte(`{"musicbrainz_label_id": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	require.NoError(t, err)
	assert.Contains(t, m, MusicbrainzLabelID)
}

func TestPublishingIDsMarshal(t *testing.T) {
	m := PubIDs{0: ""}
	data, err := json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"":""}`, string(data))
	m = PubIDs{PublishingBarcode: "12345"}
	data, err = json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"barcode":"12345"}`, string(data))
}

func TestPubIDsUnmarshal(t *testing.T) {
	m := PubIDs{}
	jsonData := []byte(`{"unknown": 0}`)
	assert.Error(t, json.Unmarshal(jsonData, &m))
	jsonData = []byte(`{"barcode": "12345"}`)
	require.NoError(t, json.Unmarshal(jsonData, &m))
	assert.Contains(t, m, PublishingBarcode)
}
