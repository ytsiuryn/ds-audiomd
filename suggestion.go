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
	Pictures         []*PictureInAudio `json:"pictures,omitempty"`
	ServiceName      string            `json:"service"`
	SourceSimilarity float64           `json:"score"`
}

type SuggestionSet struct {
	Suggestions []*Suggestion     `json:"suggestions"`
	Actors      ActorIDs          `json:"actors,omitempty"`
	Pictures    []*PictureInAudio `json:"pictures,omitempty"`
}

func NewSuggestionSet() *SuggestionSet {
	return &SuggestionSet{
		Suggestions: []*Suggestion{},
		Actors:      ActorIDs{},
		Pictures:    []*PictureInAudio{}}
}

func (ss *SuggestionSet) Optimize() {
	for _, s := range ss.Suggestions {
		s.Release.Optimize()
		if s.Release == nil {
			continue
		}
		s.Pictures = s.Release.Pictures
		s.Release.Pictures = nil
		for actor, ids := range s.Release.Actors {
			if oldIDs, ok := ss.Actors[actor]; !ok {
				ss.Actors[actor] = ids
			} else {
				for k, v := range ids {
					oldIDs[k] = v
				}
			}
		}
		s.Release.Actors = nil
	}
}
