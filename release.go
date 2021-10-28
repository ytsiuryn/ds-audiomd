package metadata

import (
	"encoding/json"
	"reflect"
	"sync"

	collection "github.com/ytsiuryn/go-collection"
	stringutils "github.com/ytsiuryn/go-stringutils"
)

// ReleaseID тип для перечисления идентификаторов релиза во внешних БД.
type ReleaseID uint8

// Допустимые значения идентификаторов релиза во внешних БД.
const (
	DiscogsReleaseID ReleaseID = iota + 1
	DiscogsMasterID
	MusicbrainzAlbumID
	MusicbrainzOriginalAlbumID
	MusicbrainzReleaseGroupID
	Rutracker
	AccurateRip
	Asin
)

func (rid ReleaseID) String() string {
	switch rid {
	case DiscogsReleaseID:
		return "discogs_release_id"
	case DiscogsMasterID:
		return "discogs_master_id"
	case MusicbrainzAlbumID:
		return "musicbrainz_album_id"
	case MusicbrainzOriginalAlbumID:
		return "musicbrainz_original_album_id"
	case MusicbrainzReleaseGroupID:
		return "musicbrainz_release_group_id"
	case Rutracker:
		return "rutracker"
	case AccurateRip:
		return "accurate_rip"
	case Asin:
		return "asin"
	}
	return ""
}

// StrToReleaseID ..
var StrToReleaseID = map[string]ReleaseID{
	"discogs_release_id":            DiscogsReleaseID,
	"discogs_master_id":             DiscogsMasterID,
	"musicbrainz_album_id":          MusicbrainzAlbumID,
	"musicbrainz_original_album_id": MusicbrainzOriginalAlbumID,
	"musicbrainz_release_group_id":  MusicbrainzReleaseGroupID,
	"rutracker":                     Rutracker,
	"accurate_rip":                  AccurateRip,
	"asin":                          Asin,
}

// ReleaseIDs представляет словарь идентификаторов релиза во внешних БД.
type ReleaseIDs map[ReleaseID]string

// MarshalJSON преобразует словарь идентификаторов релиза к JSON формату.
func (rids ReleaseIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(rids))
	for k, v := range rids {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов релиза из значения JSON.
func (rids *ReleaseIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*rids = make(ReleaseIDs, len(x))
	for k, v := range x {
		(*rids)[StrToReleaseID[k]] = v
	}
	return nil
}

// Release описывает коммерческую суть альбома из репозитория.
type Release struct {
	*ReleaseStub
	Original *ReleaseStub `json:"original,omitempty"`
}

// ReleaseStub отражает коммерческую суть продажи альбома.
type ReleaseStub struct {
	Title         string      `json:"title"`
	TotalDiscs    int         `json:"total_discs,omitempty"`
	Discs         []*Disc     `json:"discs,omitempty"`
	TotalTracks   int         `json:"total_tracks,omitempty"`
	Tracks        []*Track    `json:"tracks,omitempty"`
	Publishing    *Publishing `json:"publishing,omitempty"`
	Country       string      `json:"country,omitempty"`
	Year          int         `json:"year,omitempty"`
	Notes         string      `json:"notes,omitempty"`
	ReleaseStatus `json:"release_status,omitempty"`
	ReleaseType   `json:"release_type,omitempty"`
	ReleaseRepeat `json:"release_repeat,omitempty"`
	ReleaseRemake `json:"release_remake,omitempty"`
	ReleaseOrigin `json:"release_origin,omitempty"`
	Actors        ActorsIDs            `json:"actors,omitempty"`
	ActorRoles    ActorRoles           `json:"actors_roles,omitempty"`
	IDs           map[ReleaseID]string `json:"ids,omitempty"`
	Pictures      []*PictureInAudio    `json:"pictures,omitempty"`
	Unprocessed   collection.StrMap    `json:"unprocessed,omitempty"` // for ext view mode
	wg            sync.WaitGroup
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
		Actors:      ActorsIDs{},
		ActorRoles:  ActorRoles{},
		Publishing:  NewPublishing(),
		IDs:         map[ReleaseID]string{},
		Unprocessed: map[string]string{},
	}
}

// IsEmpty проверяет возможность сбросить ссылку на объект в nil, если все его поля
// установлены в нулевое значение.
func (stub *ReleaseStub) IsEmpty() bool {
	return reflect.DeepEqual(*stub, ReleaseStub{}) ||
		reflect.DeepEqual(*stub, *NewReleaseStub())
}

// Clean ..
func (stub *ReleaseStub) Clean() {
	if stub.Actors != nil {
		stub.Actors.Clean()
		if stub.Actors.IsEmpty() {
			stub.Actors = nil
		}
	}
	stub.ActorRoles.Clean()
	if stub.ActorRoles.IsEmpty() {
		stub.ActorRoles = nil
	}
	for _, d := range stub.Discs {
		d.Clean()
	}
	if len(stub.IDs) == 0 {
		stub.IDs = nil
	}
	// stub.Pictures.Clean()
	// if stub.Pictures.IsEmpty() {
	// 	stub.Pictures = nil
	// }
	stub.Unprocessed.Clean()
	if stub.Unprocessed.IsEmpty() {
		stub.Unprocessed = nil
	}
	for _, tr := range stub.Tracks {
		tr.Clean()
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
	performers := r.ActorRoles.Filter(IsPerformer)
	otherPerformers := other.ActorRoles.Filter(IsPerformer)
	res := performers.Compare(otherPerformers)
	if res == 0 {
		return 0., 0.
	}
	return res, 5.
}

func (r *Release) pubCompare(other *Release) (float64, float64) {
	if r.Publishing.IsEmpty() || other.Publishing.IsEmpty() {
		return 0., 0.
	}
	return r.Publishing.Compare(other.Publishing), 1.
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
	go r.aggregateActors()
	go r.aggregateUnprocessed()
	r.wg.Wait()
	r.Clean()
}

// Clean оптимизирует структуры по занимаемой памяти.
func (r *Release) Clean() {
	if r.Original != nil {
		r.Original.Clean()
		if r.Original.IsEmpty() {
			r.Original = nil
		}
	}
	if r.ReleaseStub != nil {
		r.ReleaseStub.Clean()
		if r.ReleaseStub.IsEmpty() {
			r.ReleaseStub = nil
		}
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

func (r *Release) aggregateActors() {
	defer r.wg.Done()
	for _, t := range r.Tracks {
		for name, ids := range t.Actors {
			if _, ok := r.Actors[name]; !ok {
				r.Actors[name] = ids
			}
			for k, v := range ids {
				if _, ok := r.Actors[name][k]; !ok {
					r.Actors[name][k] = v
				}
				delete(ids, k)
			}
			if len(ids) == 0 {
				delete(t.Actors, name)
			}
		}
	}
}
