package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggestionBestNResults(t *testing.T) {
	s1 := &Suggestion{SourceSimilarity: .5}
	s2 := &Suggestion{SourceSimilarity: .2}
	s3 := &Suggestion{SourceSimilarity: .7}
	res := BestNResults([]*Suggestion{s1, s2, s3}, 2)
	assert.Len(t, res, 2)
	assert.Contains(t, res, s1)
	assert.Contains(t, res, s3)
}
