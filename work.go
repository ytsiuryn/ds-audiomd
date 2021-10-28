package metadata

import (
	"encoding/json"
	"reflect"

	collection "github.com/ytsiuryn/go-collection"
)

// WorkID тип для перечисления идентификаторов композиции/произведения во внешних БД.
type WorkID uint8

// Допустимые значения идентификаторов композиции/произведения во внешних БД.
const (
	MusicbrainzWorkID WorkID = iota + 1
)

// StrToWorkID ..
var StrToWorkID = map[string]WorkID{
	"musicbrainz_work_id": MusicbrainzWorkID,
}

func (wid WorkID) String() string {
	switch wid {
	case MusicbrainzWorkID:
		return "musicbrainz_work_id"
	}
	return ""
}

// WorkIDs представляет словарь идентификаторов композиции/произведения во внешних БД.
type WorkIDs map[WorkID]string

// MarshalJSON преобразует словарь идентификаторов композиции/произведения к JSON формату.
func (wids WorkIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(wids))
	for k, v := range wids {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов композиции/произведения из значения JSON.
func (wids *WorkIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*wids = make(WorkIDs, len(x))
	for k, v := range x {
		(*wids)[StrToWorkID[k]] = v
	}
	return nil
}

// Work это часть произведения (композиция) или произведение целиком.
// Для высокоуровневых данных Position может применяться в дилогиях, трилогиях и т.д.
type Work struct {
	Parent     *Work             `json:"-"`
	Title      string            `json:"title,omitempty"`
	Position   int               `json:"index,omitempty"`
	Actors     ActorsIDs         `json:"actors,omitempty"`
	ActorRoles ActorRoles        `json:"actor_roles,omitempty"`
	Notes      string            `json:"notes,omitempty"`
	Lyrics     *Lyrics           `json:"lyrics,omitempty"`
	IDs        collection.StrMap `json:"ids,omitempty"` // ISWC
}

// NewWork создает новый объект Composition.
func NewWork() *Work {
	return &Work{
		Actors:     ActorsIDs{},
		ActorRoles: ActorRoles{},
		Lyrics:     NewLyrics(),
		IDs:        map[string]string{},
	}
}

// IsEmpty проверяет структуру на пустоту.
func (w *Work) IsEmpty() bool {
	return reflect.DeepEqual(Work{Parent: w.Parent}, *w)
}

// Clean сбрасывает ссылки полей в nil, если они не отличаются от нулевых значений.
func (w *Work) Clean() {
	w.Actors.Clean()
	w.ActorRoles.Clean()
	w.Lyrics.Clean()
	w.IDs.Clean()
	if w.Actors.IsEmpty() {
		w.Actors = nil
	}
	if w.ActorRoles.IsEmpty() {
		w.ActorRoles = nil
	}
	if w.IDs.IsEmpty() {
		w.IDs = nil
	}
	if w.Lyrics.IsEmpty() {
		w.Lyrics = nil
	}
}
