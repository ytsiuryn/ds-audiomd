package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReleaseStatusMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseStatus(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rs := ReleaseStatusOfficial
	data, err = json.Marshal(rs)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rs))
	assert.Equal(t, ReleaseStatusOfficial, rs)
}

func TestReleaseStatusDecode(t *testing.T) {
	var rs ReleaseStatus
	for k, v := range StrToReleaseStatus {
		rs.Decode(k)
		assert.Equal(t, v, rs)
	}
}

func TestReleaseStatusDecodeSlice(t *testing.T) {
	var rs ReleaseStatus
	rs.DecodeSlice(&[]string{"Text", "Oficial", "Another Text"})
	assert.Equal(t, ReleaseStatusOfficial, rs)
}

func TestReleaseTypeMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseType(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rt := ReleaseTypeAlbum
	data, err = json.Marshal(rt)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rt))
	assert.Equal(t, ReleaseTypeAlbum, rt)
}

func TestReleaseTypeDecodeSlice(t *testing.T) {
	var rt ReleaseType
	rt.DecodeSlice(&[]string{"Text", "Album", "Another Text"})
	assert.Equal(t, ReleaseTypeAlbum, rt)
}

func TestReleaseRepeatMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseRepeat(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rr := ReleaseRepeatRemake
	data, err = json.Marshal(rr)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rr))
	assert.Equal(t, ReleaseRepeatRemake, rr)
}

func TestReleaseRepeatDecodeSlice(t *testing.T) {
	var rr ReleaseRepeat
	rr.DecodeSlice(&[]string{"Text", "Repress", "Another Text"})
	assert.Equal(t, ReleaseRepeatRepress, rr)
}

func TestReleaseRemakeMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseRemake(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rr := ReleaseRemakeRemastered
	data, err = json.Marshal(rr)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rr))
	assert.Equal(t, ReleaseRemakeRemastered, rr)
}

func TestReleaseRemakeDecodeSlice(t *testing.T) {
	var rr ReleaseRemake
	rr.DecodeSlice(&[]string{"Text", "Remastered", "Another Text"})
	assert.Equal(t, ReleaseRemakeRemastered, rr)
}

func TestReleaseOriginMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseOrigin(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	ro := ReleaseOriginStudio
	data, err = json.Marshal(ro)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &ro))
	assert.Equal(t, ReleaseOriginStudio, ro)
}

func TestReleaseOriginDecodeSlice(t *testing.T) {
	var ro ReleaseOrigin
	ro.DecodeSlice(&[]string{"Text", "Studio", "Another Text"})
	assert.Equal(t, ReleaseOriginStudio, ro)
}
