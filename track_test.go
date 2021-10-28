package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiscNumberByTrackPos(t *testing.T) {
	results := map[string]int{
		"A1":    1,
		"B2":    1,
		"C3":    2,
		"D4":    2,
		"E5":    3,
		"":      1,
		"1":     1,
		"2.10":  2,
		"3 - 1": 3,
		"A.2":   1,
	}
	for k, v := range results {
		if DiscNumberByTrackPos(k) != v {
			t.Fail()
		}
	}
}

func TestTrackComplexPosition(t *testing.T) {
	assert.Equal(t, "1.1", ComplexPosition("1", "1"))
}

func TestTrackComplexTitle(t *testing.T) {
	res := ComplexTitle("Sym.5 in C minor, op.67", "1. Allegro con brio")
	assert.Equal(t, "Sym.5 in C minor, op.67. 1. Allegro con brio", res)
}

func TestTrackAddComment(t *testing.T) {
	track := NewTrack()
	track.AddComment("1st comment")
	assert.Equal(t, "1st comment", track.Notes)
	track.AddComment("2nd comment")
	assert.Equal(t, "1st comment\n2nd comment", track.Notes)
}

func TestTrackAddUnprocessed(t *testing.T) {
	track := NewTrack()
	track.AddUnprocessed("discogs", "12345")
	assert.True(t, track.Unprocessed.Exists("discogs"))
}

func TestTrackSetLyrics(t *testing.T) {
	track := NewTrack()
	track.SetLyrics("Bla-bla", false)
	assert.Equal(t, "Bla-bla", track.Composition.Lyrics.Text)
}

func TestTrackNormalizePosition(t *testing.T) {
	assert.Equal(t, "01", NormalizePosition("1"))
	assert.Equal(t, "01", NormalizePosition("01"))
}

func TestTrackSetPosition(t *testing.T) {
	track := NewTrack()
	track.SetPosition("")
	assert.Empty(t, track.Position)
	track.SetPosition("1")
	assert.Equal(t, "01", track.Position)
}

func TestTrackClean(t *testing.T) {
	tr := NewTrack()
	tr.Clean()
	if tr.Composition != nil || tr.Record != nil || tr.Actors != nil || tr.ActorRoles != nil ||
		tr.IDs != nil || tr.Unprocessed != nil || tr.AudioInfo != nil || tr.FileInfo != nil {
		t.Fail()
	}
}

func TestTrackIDsMarshal(t *testing.T) {
	m := TrackIDs{0: ""}
	data, err := json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"":""}`, string(data))
	m = TrackIDs{MusicbrainzTrackID: "12345"}
	data, err = json.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, `{"musicbrainz_track_id":"12345"}`, string(data))
}

func TestTrackIDsUnmarshal(t *testing.T) {
	m := TrackIDs{}
	jsonData := []byte(`{"unknown": 0}`)
	assert.Error(t, json.Unmarshal(jsonData, &m))
	jsonData = []byte(`{"musicbrainz_track_id": "12345"}`)
	require.NoError(t, json.Unmarshal(jsonData, &m))
	assert.Contains(t, m, MusicbrainzTrackID)
}

func TestLinkWithDisc(t *testing.T) {
	d := NewDisc(1)
	tr := NewTrack()
	tr.LinkWithDisc(d)
	assert.Equal(t, d.Number, tr.Disc().Number)
}
