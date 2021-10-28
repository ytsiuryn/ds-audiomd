package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeMedia(t *testing.T) {
	results := map[string]Media{
		"LP":             MediaLP,
		"VINYL":          MediaLP,
		"SACD":           MediaSACD,
		"CD":             MediaCD,
		"[TR24][OF]":     MediaDigital,
		"[TR24][SM][OF]": MediaDigital,
		"[DSD][OF]":      MediaDigital,
		"[DXD][OF]":      MediaDigital,
		"[DVDA][OF]":     MediaDigital,
		"REEL":           MediaReeL,
	}
	for k, v := range results {
		if DecodeMedia(k) != v {
			t.Fail()
		}
	}
}

func TestDiscMediaMarshalAndUnmarshal(t *testing.T) {
	m := MediaLP
	data, err := json.Marshal(m)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &m))
	assert.Equal(t, m, MediaLP)
}

func TestDiscFormatCompare(t *testing.T) {
	df1 := &DiscFormat{Media: MediaLP}
	df2 := &DiscFormat{Media: MediaLP}
	assert.Equal(t, df1.Compare(df2), 1.)
}

func TestMediaIDsMarshal(t *testing.T) {
	m := MediaIDs{0: ""}
	data, err := json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"":""}`, string(data))
	m = MediaIDs{DiscID: "12345"}
	data, err = json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"disc_id":"12345"}`, string(data))
}

func TestMediaIDsUnmarshal(t *testing.T) {
	m := MediaIDs{}
	jsonData := []byte(`{"unknown": 0}`)
	assert.Error(t, json.Unmarshal(jsonData, &m))
	jsonData = []byte(`{"disc_id": "12345"}`)
	require.NoError(t, json.Unmarshal(jsonData, &m))
	assert.Contains(t, m, DiscID)
}

func TestDiscClean(t *testing.T) {
	d := NewDisc(1)
	d.Clean()
	assert.True(t, d.Format.IsEmpty())
}
