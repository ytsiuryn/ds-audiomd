package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAsumptionOptimize(t *testing.T) {
	assumption := NewAssumption(nil)
	assumption.Optimize()
	// assumption.Optimize() FIX: double optimizing test

	release := NewRelease()
	release.Actors["John Doe"] = map[ActorID]string{MusicbrainzAlbumArtistID: "12345"}
	assumption = NewAssumption(release)
	assumption.Optimize()
	assert.Empty(t, assumption.Release.Actors)
	assert.Equal(t, assumption.Actors.First(), "John Doe")

	release.Pictures = append(
		release.Pictures,
		&PictureInAudio{
			PictType: PictTypeCoverFront,
			Data:     []byte("JPEG"),
		})
	assumption.Optimize()
	assert.Empty(t, assumption.Release.Pictures)
	require.NotEmpty(t, assumption.Pictures)
	assert.Equal(t, assumption.Pictures[0].PictType, PictTypeCoverFront)
}
