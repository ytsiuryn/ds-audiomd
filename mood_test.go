package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoodFromName(t *testing.T) {
	assert.Equal(t, Mood(0), MoodFromName("unknown"))
	assert.Equal(t, HappyMood, MoodFromName("happy"))
}
