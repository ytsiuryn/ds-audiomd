package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelIDsMarshal(t *testing.T) {
	m := LabelIDs{MusicbrainzLabelID: "12345"}
	data, err := json.Marshal(m)
	assert.Equal(t, `{"musicbrainz_label_id":"12345"}`, string(data))
	assert.NoError(t, err)
}

func TestLabelIDsUnmarshal(t *testing.T) {
	m := LabelIDs{}
	jsonData := []byte(`{"musicbrainz_label_id": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	assert.NoError(t, err)
	assert.Contains(t, m, MusicbrainzLabelID)
}

func TestPublishingIDMarshal(t *testing.T) {
	data, err := json.Marshal(PublishingBarcode)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"barcode"`), data)
}

func TestPubIDsUnmarshal(t *testing.T) {
	m := PubIDs{}
	jsonData := []byte(`{"barcode": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	assert.NoError(t, err)
	assert.Contains(t, m, PublishingBarcode)
}
