package metadata

import (
	"testing"

	collection "github.com/gtyrin/go-collection"
)

func TestSuggestionBestNResults(t *testing.T) {
	s1 := &Suggestion{SourceSimilarity: .5}
	s2 := &Suggestion{SourceSimilarity: .2}
	s3 := &Suggestion{SourceSimilarity: .7}
	res := BestNResults([]*Suggestion{s1, s2, s3}, 2)
	if len(res) != 2 || !collection.Contains(s1, res) || !collection.Contains(s3, res) {
		t.Fail()
	}
}
