package metadata

import (
	"encoding/json"
	"reflect"
)

// RecordingID тип для перечисления идентификаторов записи трека во внешних БД.
type RecordingID uint8

// Допустимые значения идентификаторов записи во внешних БД.
const (
	MusicbrainzRecordingID RecordingID = iota + 1
	ISRC
)

// RecordingIDs представляет словарь идентификаторов записи во внешних БД.
type RecordingIDs map[RecordingID]string

// MarshalJSON преобразует словарь идентификаторов записи к JSON формату.
func (rids RecordingIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(rids))
	for k, v := range rids {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов записи из значения JSON.
func (rids *RecordingIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*rids = make(RecordingIDs, len(x))
	for k, v := range x {
		(*rids)[StrToRecordingID[k]] = v
	}
	return nil
}

// StrToRecordingID ..
var StrToRecordingID = map[string]RecordingID{
	"musicbrainz_recording_id": MusicbrainzRecordingID,
	"isrc":                     ISRC,
}

func (rid RecordingID) String() string {
	switch rid {
	case MusicbrainzRecordingID:
		return "musicbrainz_recording_id"
	case ISRC:
		return "isrc"
	}
	return ""
}

// RecordSession описывает общие свойства сессии записи.
type RecordSession struct {
	Place string `json:"place,omitempty"`
	Time  string `json:"time,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// Record содержит сведения о записи композиции.
type Record struct {
	Duration   int32                  `json:"duration,omitempty"`
	Actors     ActorsIDs              `json:"actors,omitempty"`
	ActorRoles ActorRoles             `json:"actor_roles,omitempty"`
	Moods      Moods                  `json:"moods,omitempty"`
	Genres     []string               `json:"genres,omitempty"`
	IDs        map[RecordingID]string `json:"ids,omitempty"`
	Notes      string                 `json:"notes,omitempty"`
}

// NewRecord создает новый объект Record.
func NewRecord() *Record {
	return &Record{
		Actors:     ActorsIDs{},
		ActorRoles: ActorRoles{},
		IDs:        map[RecordingID]string{},
	}
}

// IsEmpty проверяет коллекцию как не инициализированную.
func (p *Record) IsEmpty() bool {
	return reflect.DeepEqual(Record{}, *p)
}

// Clean сбрасывает всю коллекцию в nil, если поля структуры не отличаются от нулевых значений.
func (p *Record) Clean() {
	p.Actors.Clean()
	p.ActorRoles.Clean()
	if p.Actors.IsEmpty() {
		p.Actors = nil
	}
	if p.ActorRoles.IsEmpty() {
		p.ActorRoles = nil
	}
	if len(p.IDs) == 0 {
		p.IDs = nil
	}
}

// AddRole добавляет роль для актора записи.
func (p *Record) AddRole(name, role string) {
	p.ActorRoles[name] = append(p.ActorRoles[name], role)
}

// Performers return all album performers.
func (p *Record) Performers() ActorRoles {
	return p.ActorRoles.Filter(IsPerformer)
}
