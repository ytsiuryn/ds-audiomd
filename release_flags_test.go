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
	assert.Equal(t, rs, ReleaseStatusOfficial)
}

func TestReleaseStatusDecode(t *testing.T) {
	var rs ReleaseStatus
	for k, v := range StrToReleaseStatus {
		rs.Decode(k)
		assert.Equal(t, rs, v)
	}
}

func TestReleaseStatusDecodeSlice(t *testing.T) {
	var rs ReleaseStatus
	rs.DecodeSlice(&[]string{"Text", "Oficial", "Another Text"})
	assert.Equal(t, rs, ReleaseStatusOfficial)
}

func TestReleaseTypeMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseType(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rt := ReleaseTypeAlbum
	data, err = json.Marshal(rt)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rt))
	assert.Equal(t, rt, ReleaseTypeAlbum)
}

func TestReleaseTypeDecodeSlice(t *testing.T) {
	var rt ReleaseType
	rt.DecodeSlice(&[]string{"Text", "Album", "Another Text"})
	assert.Equal(t, rt, ReleaseTypeAlbum)
}

func TestReleaseRepeatMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseRepeat(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rr := ReleaseRepeatRemake
	data, err = json.Marshal(rr)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rr))
	assert.Equal(t, rr, ReleaseRepeatRemake)
}

func TestReleaseRepeatDecodeSlice(t *testing.T) {
	var rr ReleaseRepeat
	rr.DecodeSlice(&[]string{"Text", "Repress", "Another Text"})
	assert.Equal(t, rr, ReleaseRepeatRepress)
}

func TestReleaseRemakeMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseRemake(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	rr := ReleaseRemakeRemastered
	data, err = json.Marshal(rr)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &rr))
	assert.Equal(t, rr, ReleaseRemakeRemastered)
}

func TestReleaseRemakeDecodeSlice(t *testing.T) {
	var rr ReleaseRemake
	rr.DecodeSlice(&[]string{"Text", "Remastered", "Another Text"})
	assert.Equal(t, rr, ReleaseRemakeRemastered)
}

func TestReleaseOriginMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(ReleaseOrigin(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	ro := ReleaseOriginStudio
	data, err = json.Marshal(ro)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &ro))
	assert.Equal(t, ro, ReleaseOriginStudio)
}

func TestReleaseOriginDecodeSlice(t *testing.T) {
	var ro ReleaseOrigin
	ro.DecodeSlice(&[]string{"Text", "Studio", "Another Text"})
	assert.Equal(t, ro, ReleaseOriginStudio)
}
