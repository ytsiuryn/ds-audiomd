package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseStatusMarshalAndUnmarshal(t *testing.T) {
	rs := ReleaseStatusOfficial
	data, err := json.Marshal(rs)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &rs); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rs, ReleaseStatusOfficial)
}

func TestReleaseStatusDecode(t *testing.T) {
	var rs ReleaseStatus
	for k, v := range StrToReleaseStatus {
		if rs.Decode(k); rs != v {
			t.Fail()
		}
	}
}

func TestReleaseStatusDecodeSlice(t *testing.T) {
	var rs ReleaseStatus
	rs.DecodeSlice(&[]string{"Text", "Oficial", "Another Text"})
	assert.Equal(t, rs, ReleaseStatusOfficial)
}

func TestReleaseTypeMarshalAndUnmarshal(t *testing.T) {
	rt := ReleaseTypeAlbum
	data, err := json.Marshal(rt)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &rt); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rt, ReleaseTypeAlbum)
}

func TestReleaseTypeDecodeSlice(t *testing.T) {
	var rt ReleaseType
	rt.DecodeSlice(&[]string{"Text", "Album", "Another Text"})
	assert.Equal(t, rt, ReleaseTypeAlbum)
}

func TestReleaseRepeatMarshalAndUnmarshal(t *testing.T) {
	rr := ReleaseRepeatRemake
	data, err := json.Marshal(rr)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &rr); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rr, ReleaseRepeatRemake)
}

func TestReleaseRepeatDecodeSlice(t *testing.T) {
	var rr ReleaseRepeat
	rr.DecodeSlice(&[]string{"Text", "Repress", "Another Text"})
	assert.Equal(t, rr, ReleaseRepeatRepress)
}

func TestReleaseRemakeMarshalAndUnmarshal(t *testing.T) {
	rr := ReleaseRemakeRemastered
	data, err := json.Marshal(rr)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &rr); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rr, ReleaseRemakeRemastered)
}

func TestReleaseRemakeDecodeSlice(t *testing.T) {
	var rr ReleaseRemake
	rr.DecodeSlice(&[]string{"Text", "Remastered", "Another Text"})
	assert.Equal(t, rr, ReleaseRemakeRemastered)
}

func TestReleaseOriginMarshalAndUnmarshal(t *testing.T) {
	ro := ReleaseOriginStudio
	data, err := json.Marshal(ro)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &ro); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ro, ReleaseOriginStudio)
}

func TestReleaseOriginDecodeSlice(t *testing.T) {
	var ro ReleaseOrigin
	ro.DecodeSlice(&[]string{"Text", "Studio", "Another Text"})
	assert.Equal(t, ro, ReleaseOriginStudio)
}
