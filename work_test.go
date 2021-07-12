package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkIsEmptyAndClean(t *testing.T) {
	w := NewWork()
	w.Clean()
	assert.True(t, w.IsEmpty())
	assert.Empty(t, w.Actors)
	assert.Empty(t, w.IDs)
}
