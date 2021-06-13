package metadata

import (
	"testing"
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
	if ComplexPosition("1", "1") != "1.1" {
		t.Fail()
	}
}

func TestTrackComplexTitle(t *testing.T) {
	res := ComplexTitle("Sym.5 in C minor, op.67", "1. Allegro con brio")
	if res != "Sym.5 in C minor, op.67. 1. Allegro con brio" {
		t.Fail()
	}
}

func TestTrackAddComment(t *testing.T) {
	track := NewTrack()
	track.AddComment("1st comment")
	if track.Notes != "1st comment" {
		t.Fail()
	}
	track.AddComment("2nd comment")
	if track.Notes != "1st comment\n2nd comment" {
		t.Fail()
	}
}

func TestTrackAddUnprocessed(t *testing.T) {
	track := NewTrack()
	track.AddUnprocessed("discogs", "12345")
	if !track.Unprocessed.Exists("discogs") {
		t.Fail()
	}
}

func TestTrackSetLyrics(t *testing.T) {
	track := NewTrack()
	track.SetLyrics("Bla-bla", false)
	if track.Composition.Lyrics.Text != "Bla-bla" {
		t.Fail()
	}
}

func TestTrackNormalizePosition(t *testing.T) {
	if NormalizePosition("1") != "01" {
		t.Fail()
	}
	if NormalizePosition("01") != "01" {
		t.Fail()
	}
}

func TestTrackSetPosition(t *testing.T) {
	track := NewTrack()
	track.SetPosition("")
	if track.Position != "" {
		t.Fail()
	}
	track.SetPosition("1")
	if track.Position != "01" {
		t.Fail()
	}
}

func TestTrackClean(t *testing.T) {
	tr := NewTrack()
	tr.Clean()
	if tr.Composition != nil || tr.Record != nil || tr.Actors != nil || tr.IDs != nil ||
		tr.Unprocessed != nil || tr.AudioInfo != nil || tr.FileInfo != nil {
		t.Fail()
	}
}
