package metadata

import (
	"sort"

	intutils "github.com/ytsiuryn/go-intutils"
)

// Suggestion has a search result of an online service.
type Suggestion struct {
	Entity           interface{} `json:"entity"`
	ServiceName      string      `json:"service"`
	SourceSimilarity float64     `json:"score"`
	// Confirmed   bool        `json:"confirmed,omitempty"`
}

// BestNResults returns the N suggestion results.
func BestNResults(in []*Suggestion, n int) []*Suggestion {
	sort.SliceStable(in, func(i, j int) bool {
		return in[i].SourceSimilarity > in[j].SourceSimilarity
	})
	return in[:intutils.MinOf(n, len(in))]
}
