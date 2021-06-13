package metadata

import (
	"testing"
)

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
	if rs != ReleaseStatusOfficial {
		t.Fail()
	}
}

func TestReleaseTypeDecodeSlice(t *testing.T) {
	var rt ReleaseType
	rt.DecodeSlice(&[]string{"Text", "Album", "Another Text"})
	if rt != ReleaseTypeAlbum {
		t.Fail()
	}
}

func TestReleaseRepeatDecodeSlice(t *testing.T) {
	var rr ReleaseRepeat
	rr.DecodeSlice(&[]string{"Text", "Repress", "Another Text"})
	if rr != ReleaseRepeatRepress {
		t.Fail()
	}
}

func TestReleaseRemakeDecodeSlice(t *testing.T) {
	var rr ReleaseRemake
	rr.DecodeSlice(&[]string{"Text", "Remastered", "Another Text"})
	if rr != ReleaseRemakeRemastered {
		t.Fail()
	}
}

func TestReleaseOriginDecodeSlice(t *testing.T) {
	var ro ReleaseOrigin
	ro.DecodeSlice(&[]string{"Text", "Studio", "Another Text"})
	if ro != ReleaseOriginStudio {
		t.Fail()
	}
}
