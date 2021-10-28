package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActorIDsAdd(t *testing.T) {
	var a ActorsIDs = map[ActorName]ActorIDs{}
	a.Add("John Doe", MusicbrainzAlbumArtistID, "12345")
	assert.Len(t, a, 1)
	assert.Len(t, a["John Doe"], 1)
	a.Add("John Doe", MusicbrainzAlbumArtistID, "12345") // не уникальные
	assert.Len(t, a, 1)
	assert.Len(t, a["John Doe"], 1)
	a.Add("Nemo", MusicbrainzAlbumArtistID, "234567")
	assert.Len(t, a, 2)
	a.Add("Nemo", MusicbrainzOriginalArtistID, "abcdefg")
	assert.Len(t, a, 2)
	assert.Len(t, a["Nemo"], 2)
}
func TestActorIDsMerge(t *testing.T) {
	var a1, a2 ActorsIDs
	a1 = map[ActorName]ActorIDs{
		"John Doe": {MusicbrainzAlbumArtistID: "12345"},
		"Nemo":     {MusicbrainzAlbumArtistID: "abcd"},
	}
	a2 = map[ActorName]ActorIDs{
		"John Doe": {MusicbrainzOriginalArtistID: "zyxwv"},
		"Nemo":     {MusicbrainzAlbumArtistID: "abcd"},
	}
	a1.Merge(a2)
	assert.Len(t, a1["John Doe"], 2)
	assert.Len(t, a1["Nemo"], 1)
}

func TestActorRolesAdd(t *testing.T) {
	ar := ActorRoles{}
	ar.Add("John Doe", "performer")
	assert.Len(t, ar, 1)
	ar.Add("John Doe", "performer")
	assert.Len(t, ar["John Doe"], 1)
	ar.Add("John Doe", "conductor")
	assert.Len(t, ar["John Doe"], 2)
}
func TestActorRolesCompare(t *testing.T) {

}

func TestActorRolesFilter(t *testing.T) {
	actorRoles := ActorRoles{}
	actorRoles["John Doe"] = []string{"performer"}
	assert.Len(t, actorRoles.Filter(IsPerformer), 1)
}

func TestActorRolesFirst(t *testing.T) {
	actorRoles := ActorRoles{}
	assert.Equal(t, "", actorRoles.First())
	actorRoles.Add("John Doe", "performer")
	actorRoles.Add("Nemo", "soloist")
	assert.Equal(t, "John Doe", actorRoles.First())
}

func TestActorIDsMarshal(t *testing.T) {
	m := ActorIDs{MusicbrainzArtistID: "12345"}
	data, err := json.Marshal(m)
	assert.Equal(t, `{"musicbrainz_artist_id":"12345"}`, string(data))
	assert.NoError(t, err)
}

func TestActorIDsUnmarshal(t *testing.T) {
	m := ActorIDs{}
	jsonData := []byte(`{"musicbrainz_artist_id": "12345"}`)
	err := json.Unmarshal(jsonData, &m)
	assert.NoError(t, err)
	assert.Contains(t, m, MusicbrainzArtistID)
}
