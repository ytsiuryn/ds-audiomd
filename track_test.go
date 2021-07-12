package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	}
	for k, v := range results {
		if DiscNumberByTrackPos(k) != v {
			t.Fail()
		}
	}
}

func TestTrackComplexPosition(t *testing.T) {
	assert.Equal(t, ComplexPosition("1", "1"), "1.1")
}

func TestTrackComplexTitle(t *testing.T) {
	res := ComplexTitle("Sym.5 in C minor, op.67", "1. Allegro con brio")
	assert.Equal(t, res, "Sym.5 in C minor, op.67. 1. Allegro con brio")
}

func TestTrackAddComment(t *testing.T) {
	track := NewTrack()
	track.AddComment("1st comment")
	assert.Equal(t, track.Notes, "1st comment")
	track.AddComment("2nd comment")
	assert.Equal(t, track.Notes, "1st comment\n2nd comment")
}

func TestTrackAddUnprocessed(t *testing.T) {
	track := NewTrack()
	track.AddUnprocessed("discogs", "12345")
	assert.True(t, track.Unprocessed.Exists("discogs"))
}

func TestTrackSetLyrics(t *testing.T) {
	track := NewTrack()
	track.SetLyrics("Bla-bla", false)
	assert.Equal(t, track.Composition.Lyrics.Text, "Bla-bla")
}

func TestTrackNormalizePosition(t *testing.T) {
	assert.Equal(t, NormalizePosition("1"), "01")
	assert.Equal(t, NormalizePosition("01"), "01")
}

func TestTrackSetPosition(t *testing.T) {
	track := NewTrack()
	track.SetPosition("")
	assert.Empty(t, track.Position)
	track.SetPosition("1")
	assert.Equal(t, track.Position, "01")
}

func TestTrackClean(t *testing.T) {
	tr := NewTrack()
	tr.Clean()
	if tr.Composition != nil || tr.Record != nil || tr.Actors != nil || tr.ActorRoles != nil ||
		tr.IDs != nil || tr.Unprocessed != nil || tr.AudioInfo != nil || tr.FileInfo != nil {
		t.Fail()
	}
}
