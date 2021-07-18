package metadata

import (
	"sort"

	intutils "github.com/ytsiuryn/go-intutils"
)

// BestNResults returns the N suggestion results.
func BestNResults(in []*Suggestion, n int) []*Suggestion {
	sort.SliceStable(in, func(i, j int) bool {
		return in[i].SourceSimilarity > in[j].SourceSimilarity
	})
	return in[:intutils.MinOf(n, len(in))]
}

// Suggestion has a search result of an online service.
type Suggestion struct {
	*Release         `json:"release"`
	Actors           ActorIDs          `json:"actors"`
	Pictures         []*PictureInAudio `json:"pictures,omitempty"`
	ServiceName      string            `json:"service"`
	SourceSimilarity float64           `json:"score"`
}

func (s *Suggestion) Optimize() {
	s.Release.Optimize()
	if s.Release == nil {
		return
	}
	s.Actors = s.Release.Actors
	s.Release.Actors = nil
	s.Pictures = s.Release.Pictures
	s.Release.Pictures = nil
}
