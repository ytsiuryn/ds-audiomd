package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoodFromName(t *testing.T) {
	assert.Equal(t, Mood(0), MoodFromName("unknown"))
	assert.Equal(t, HappyMood, MoodFromName("happy"))
}

func TestMoodMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(Mood(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	mood := CalmMood
	data, err = json.Marshal(mood)
	require.NoError(t, err)
	assert.Equal(t, []byte(`"calm"`), data)
	require.NoError(t, json.Unmarshal(data, &mood))
	assert.Equal(t, CalmMood, mood)
}
