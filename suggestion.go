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

// Suggestion используется внешними сервисами для хранения единичного результата.
type Suggestion struct {
	Release          *Release `json:"release"`
	ServiceName      string   `json:"service"`
	SourceSimilarity float64  `json:"score"`
}

// SuggestionSet объединяет несколько результатов и оптимизирует размер за счет
// выноса сведений об акторах в отдельную сущность.
type SuggestionSet struct {
	Suggestions []*Suggestion `json:"suggestions"`
	Actors      ActorsIDs     `json:"actors,omitempty"`
}

// NewSuggestion ..
func NewSuggestion() *Suggestion {
	return &Suggestion{Release: NewRelease()}
}

// NewSuggestionSet создает объект коллекции результатов SuggestionSet.
func NewSuggestionSet() *SuggestionSet {
	return &SuggestionSet{
		Suggestions: []*Suggestion{},
		Actors:      ActorsIDs{}}
}

// Optimize оптимизирует релиз-данные для каждого результата и аггрегирует коды
// акторов во внешних БД в поле Actors.
func (ss *SuggestionSet) Optimize() {
	for _, s := range ss.Suggestions {
		s.Release.Optimize()
		if s.Release == nil {
			return
		}
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
