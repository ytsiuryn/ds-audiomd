package metadata

import (
	"testing"
)

func TestReleaseCover(t *testing.T) {
	r := NewRelease()
	r.Pictures = append(r.Pictures, &PictureInAudio{PictType: PictTypeArtist})
	if r.Cover() != nil {
		t.Fail()
	}
	r.Pictures = append(r.Pictures, &PictureInAudio{PictType: PictTypeCoverFront})
	if r.Cover() == nil {
		t.Fail()
	}
}

func TestReleaseTrackByPosition(t *testing.T) {
	track := NewTrack()
	track.Position = "1"
	r := NewRelease()
	r.Tracks = []*Track{track}
	if r.TrackByPosition("3") != nil {
		t.Fail()
	}
	if r.TrackByPosition("1") == nil {
		t.Fail()
	}
}

func TestReleaseDisc(t *testing.T) {
	r := NewRelease()
	d := r.Disc(3)
	if len(r.Discs) != 3 || d == nil {
		t.Fail()
	}
}

func TestReleasePerformersCompare(t *testing.T) {
	r := NewRelease()
	r2 := NewRelease()
	if res, weight := r.performersCompare(r2); res != 0. || weight != 0. {
		t.Fail()
	}
	r.Actors.AddRole("Miles Davis", "performer")
	r2.Actors.AddRole("Miles Davis", "performer")
	if res, weight := r.performersCompare(r2); res != 1. || weight != 5. {
		t.Fail()
	}
}

func TestReleasePubCompare(t *testing.T) {
	r := NewRelease()
	r.Publishing = append(r.Publishing, NewReleaseLabel("Analog Audio"))
	r2 := NewRelease()
	if res, weight := r.pubCompare(r2); res != 0. || weight != 0. {
		t.Fail()
	}
	r2.Publishing = append(r2.Publishing, NewReleaseLabel("RCA"))
	r2.Publishing = append(r2.Publishing, NewReleaseLabel("Analog Audio"))
	if res, weight := r.pubCompare(r2); res != .99 || weight != 1. {
		t.Fail()
	}
}

func TestReleaseTracksCompare(t *testing.T) {
	r := NewRelease()
	track := NewTrack()
	track.Title = "Some Prince will come"
	r.Tracks = append(r.Tracks, track)
	r2 := NewRelease()
	if res, weight := r.tracksCompare(r2); res != 0. || weight != 0. {
		t.Fail()
	}
	track2 := NewTrack()
	track2.Title = "Some Prince will come"
	r2.Tracks = append(r2.Tracks, track2)
	if res, weight := r.tracksCompare(r2); res != 1. || weight != 1. {
		t.Fail()
	}
}

func TestReleaseDiscFormatsCompare(t *testing.T) {
	r := NewRelease()
	r2 := NewRelease()
	if res, weight := r.discFormatsCompare(r2); res != 0. || weight != 0. {
		t.Fail()
	}
	d := NewDisc(1)
	d.Format.Media = MediaLP
	r.Discs = append(r.Discs, d)
	d2 := NewDisc(1)
	d2.Format.Media = MediaLP
	r2.Discs = append(r2.Discs, d2)
	if res, weight := r.discFormatsCompare(r2); res != 1. || weight != 1. {
		t.Fail()
	}
}

func TestReleaseOptimizeNotes(t *testing.T) {
	r := NewRelease()
	t1 := NewTrack()
	t1.Notes = "Notes"
	t2 := NewTrack()
	t2.Notes = "Notes"
	r.Tracks = append(r.Tracks, t1, t2)
	r.wg.Add(1)
	r.aggregateNotes()
	if len(r.Notes) == 0 || len(t1.Notes) != 0 || len(t2.Notes) != 0 {
		t.Fail()
	}
}

func TestReleaseAggregateUnprocessed(t *testing.T) {
	r := NewRelease()
	t1 := NewTrack()
	t1.Unprocessed = map[string]string{"A": "AA", "B": "BB"}
	t2 := NewTrack()
	t2.Unprocessed = map[string]string{"A": "AA", "C": "CC"}
	r.Tracks = append(r.Tracks, t1, t2)
	r.wg.Add(1)
	r.aggregateUnprocessed()
	if len(r.Unprocessed) != 0 || len(t1.Unprocessed) != 2 || len(t2.Unprocessed) != 2 {
		t.Fail()
	}
	t1.Position = "1"
	t2.Position = "2"
	r.wg.Add(1)
	r.aggregateUnprocessed()
	if len(r.Unprocessed) != 1 || len(t1.Unprocessed) != 1 || len(t2.Unprocessed) != 1 {
		t.Fail()
	}
}

func TestReleaseAggregateReleaseActorRoles(t *testing.T) {
	r := NewRelease()
	t1 := NewTrack()
	t1.Actors.AddRole("Nemo", "engineer")
	t2 := NewTrack()
	t2.Actors.AddRole("Nemo", "engineer")
	r.Tracks = append(r.Tracks, t1, t2)
	r.wg.Add(1)
	r.aggregateReleaseActorRoles()
	if len(*r.Actors) != 1 || len(*t1.Actors) != 0 || len(*t2.Actors) != 0 {
		t.Fail()
	}
}

func TestReleaseClean(t *testing.T) {
	r := NewRelease()
	r.Clean()
	if r.Original.IDs != nil || r.Original.Tracks != nil || r.Original.Unprocessed != nil ||
		r.IDs != nil || r.Tracks != nil || r.Unprocessed != nil {
		t.FailNow()
	}
}
