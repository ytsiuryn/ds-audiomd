package metadata

import (
	"reflect"
	"sync"

	collection "github.com/ytsiuryn/go-collection"
	stringutils "github.com/ytsiuryn/go-stringutils"
)

// Release описывает коммерческую суть альбома из репозитория.
type Release struct {
	*ReleaseStub
	Original *ReleaseStub `json:"original,omitempty"`
}

// ReleaseStub отражает коммерческую суть продажи альбома.
type ReleaseStub struct {
	Title         string        `json:"title"`
	TotalDiscs    int           `json:"total_discs,omitempty"`
	Discs         []*Disc       `json:"discs,omitempty"`
	TotalTracks   int           `json:"total_tracks,omitempty"`
	Tracks        []*Track      `json:"tracks,omitempty"`
	Publishing    []*Publishing `json:"publishing,omitempty"`
	Country       string        `json:"country,omitempty"`
	Year          int           `json:"year,omitempty"`
	Notes         string        `json:"notes,omitempty"`
	ReleaseStatus `json:"release_status,omitempty"`
	ReleaseType   `json:"release_type,omitempty"`
	ReleaseRepeat `json:"release_repeat,omitempty"`
	ReleaseRemake `json:"release_remake,omitempty"`
	ReleaseOrigin `json:"release_origin,omitempty"`
	Actors        *Actors `json:"actors,omitempty"`
	// идентификаторы online БД ("discogs", "musicbrainz", "rutracker")
	IDs         collection.StrMap `json:"ids,omitempty"`
	Pictures    []*PictureInAudio `json:"pictures,omitempty"`
	Unprocessed collection.StrMap `json:"unprocessed,omitempty"` // for ext view mode
	wg          sync.WaitGroup
}

// NewRelease construct a new release object.
func NewRelease() *Release {
	return &Release{
		ReleaseStub: NewReleaseStub(),
		Original:    NewReleaseStub(),
	}
}

// NewReleaseStub construct a new release object.
func NewReleaseStub() *ReleaseStub {
	return &ReleaseStub{
		Actors:      NewActorCollection(),
		IDs:         map[string]string{},
		Unprocessed: map[string]string{},
	}
}

// Cover возвращает объект PictureInAudio, если он описывает обложку альбома, или nil.
func (r *Release) Cover() *PictureInAudio {
	for _, pia := range r.Pictures {
		if pia.PictType == PictTypeCoverFront {
			return pia
		}
	}
	return nil
}

// TrackByPosition возвращает объект трека по его позиции.
func (r *Release) TrackByPosition(pos string) *Track {
	for _, tr := range r.Tracks {
		if tr.Position == pos {
			return tr
		}
	}
	return nil
}

// Disc возвращает ссылку на объект диска.
// Если диск с указанным номером не существует, он добавляется в колекцию в позицию,
// соответствующую его номеру с заполнением "пробелов".
func (r *Release) Disc(num int) *Disc {
	lenDiff := num - len(r.Discs)
	for i := 0; i < lenDiff; i++ {
		r.Discs = append(r.Discs, NewDisc(len(r.Discs)+i+1))
	}
	return r.Discs[num-1]
}

// --- COMPARE METHODS ---

// Compare compare two albums by important metadata.
// Если номера каталогов изданий совпадают, объекты считаются идентичными досрочно.
func (r *Release) Compare(other *Release) float64 {
	labsR, labsW := r.pubCompare(other)
	if labsR == 1. {
		return 1.
	}
	titleR, titleW := stringutils.JaroWinklerDistance(r.Title, other.Title), 5.
	perfR, perfW := r.performersCompare(other)
	trcksR, trcksW := r.tracksCompare(other)
	frmsR, frmsW := r.discFormatsCompare(other)
	return (titleW*titleR + perfW*perfR + labsW*labsR + trcksW*trcksR + frmsW*frmsR) /
		(titleW + perfW + labsW + trcksW + frmsW)
}

func (r *Release) performersCompare(other *Release) (float64, float64) {
	performers := r.Actors.Filter(IsPerformer)
	otherPerformers := other.Actors.Filter(IsPerformer)
	res := performers.Compare(otherPerformers)
	if res == 0 {
		return 0., 0.
	}
	return res, 5.
}

func (r *Release) pubCompare(other *Release) (float64, float64) {
	if len(r.Publishing) == 0 || len(other.Publishing) == 0 {
		return 0., 0.
	}
	var max, res float64
	for _, pub := range r.Publishing {
		for _, pub2 := range other.Publishing {
			res = pub.Compare(pub2)
			if res == 1. {
				return 1., 0
			}
			if max < res {
				max = res
			}
		}
	}
	return max, 1.
}

func (r *Release) tracksCompare(other *Release) (float64, float64) {
	if len(r.Tracks) != len(other.Tracks) {
		return 0., 0.
	}
	sum := 0.
	for i, tr := range r.Tracks {
		sum += tr.Compare(other.Tracks[i])
	}
	return sum, float64(len(r.Tracks))
}

func (r *Release) discFormatsCompare(other *Release) (float64, float64) {
	var res, max float64
	for i, disc := range r.Discs {
		res = disc.Format.Compare(other.Discs[i].Format)
		if max < res {
			max = res
		}
	}
	if max == 0 {
		return 0., 0.
	}
	return max, 1.
}

// --- OPTIMIZATION METHODS ---

type void struct{}

// Optimize улучшает хранение данных за счет делигирования повторяющихся данных на
// уровень выше.
func (r *Release) Optimize() {
	r.wg.Add(3)
	go r.aggregateNotes()
	go r.aggregateReleaseActorRoles()
	go r.aggregateUnprocessed()
	r.wg.Wait()
	r.Clean()
}

// Clean оптимизирует структуры по занимаемой памяти.
func (r *Release) Clean() {
	r.Original.Actors.Clean()
	if r.Original.Actors.IsEmpty() {
		r.Original.Actors = nil
	}
	for _, d := range r.Discs {
		d.Clean()
	}
	r.Original.IDs.Clean()
	if r.Original.IDs.IsEmpty() {
		r.Original.IDs = nil
	}
	r.Original.Unprocessed.Clean()
	if r.Original.Unprocessed.IsEmpty() {
		r.Original.Unprocessed = nil
	}
	for _, tr := range r.Original.Tracks {
		tr.Clean()
	}
	r.Actors.Clean()
	if r.Actors.IsEmpty() {
		r.Actors = nil
	}
	r.IDs.Clean()
	if r.IDs.IsEmpty() {
		r.IDs = nil
	}
	r.Unprocessed.Clean()
	if r.Unprocessed.IsEmpty() {
		r.Unprocessed = nil
	}
	for _, tr := range r.Tracks {
		tr.Clean()
	}
}

func (r *Release) aggregateNotes() {
	defer r.wg.Done()
	var member void
	commentMap := make(map[string]void)
	for _, track := range r.Tracks {
		if _, ok := commentMap[track.Notes]; !ok {
			commentMap[track.Notes] = member
		}
	}
	if len(commentMap) == 1 {
		for k := range commentMap {
			r.Notes = k
			for _, track := range r.Tracks {
				track.Notes = ""
			}
		}
	}
}

func (r *Release) aggregateUnprocessed() {
	defer r.wg.Done()
	trackCount := len(r.Tracks)
	unprocessed := map[string]map[string]int{}
	for _, track := range r.Tracks {
		if track.Position == "" {
			return
		}
	}
	for _, track := range r.Tracks {
		for k, v := range track.Unprocessed {
			if _, ok := unprocessed[k]; !ok {
				unprocessed[k] = map[string]int{}
			}
			if _, ok := unprocessed[k][v]; !ok {
				unprocessed[k][v] = 0
			}
			unprocessed[k][v]++
		}
	}
	for k, m := range unprocessed {
		for _, v := range reflect.ValueOf(m).MapKeys() {
			val := v.String()
			if unprocessed[k][val] == trackCount {
				r.Unprocessed[k] = val
				for i := trackCount - 1; i >= 0; i-- {
					delete(r.Tracks[i].Unprocessed, k)
				}
			}
		}
	}
}

func (r *Release) aggregateReleaseActorRoles() {
	r.aggregateActors(
		func(t *Track) *Actors {
			return t.Actors
		},
		r.Actors)
}

func (r *Release) aggregateActors(fnc func(*Track) *Actors, acc *Actors) {
	defer r.wg.Done()
	tCount := len(r.Tracks)
	counters := map[string]map[string]int{} // name -> role -> track_count
	for _, t := range r.Tracks {
		for _, air := range *fnc(t) {
			if _, ok := counters[air.Name]; !ok {
				counters[air.Name] = map[string]int{}
			}
			for _, role := range air.Roles {
				if _, ok := counters[air.Name][role]; !ok {
					counters[air.Name][role] = 0
				}
				counters[air.Name][role]++
			}
		}
	}
	for _, t := range r.Tracks {
		actors := fnc(t)
		for actorInd := len(*actors) - 1; actorInd >= 0; actorInd-- {
			air := (*actors)[actorInd]
			for roleInd := len(air.Roles) - 1; roleInd >= 0; roleInd-- {
				role := air.Roles[roleInd]
				if counters[air.Name][role] == tCount {
					acc.Merge(air)
					actors.DeleteRole(air.Name, role)
					if len(actors.Roles(air.Name)) == 0 {
						actors.DeleteActor(air.Name)
					}
				}
			}
		}
	}
}
