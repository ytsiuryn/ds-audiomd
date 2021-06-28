package metadata

import (
	"reflect"

	collection "github.com/ytsiuryn/go-collection"
)

// RecordSession описывает общие свойства сессии записи.
type RecordSession struct {
	Place string `json:"place,omitempty"`
	Time  string `json:"time,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// Record содержит сведения о записи композиции.
type Record struct {
	Duration   int32             `json:"duration,omitempty"`
	Actors     ActorIDs          `json:"actors,omitempty"`
	ActorRoles ActorRoles        `json:"actor_roles,omitempty"`
	Moods      Moods             `json:"moods,omitempty"`
	Genres     []string          `json:"genres,omitempty"`
	IDs        collection.StrMap `json:"ids,omitempty"` // ISRC, MusicbrainzRecordID
	Notes      string            `json:"notes,omitempty"`
}

// NewRecord создает новый объект Record.
func NewRecord() *Record {
	return &Record{
		Actors:     ActorIDs{},
		ActorRoles: ActorRoles{},
		IDs:        map[string]string{},
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
	p.IDs.Clean()
	if p.Actors.IsEmpty() {
		p.Actors = nil
	}
	if p.ActorRoles.IsEmpty() {
		p.ActorRoles = nil
	}
	if p.IDs.IsEmpty() {
		p.IDs = nil
	}
}

// AddRole добавляет роль для актора записи.
func (p *Record) AddRole(name, role string) {
	p.ActorRoles[ActorName(name)] = append(p.ActorRoles[ActorName(name)], role)
}

// Performers return all album performers.
func (p *Record) Performers() ActorRoles {
	return p.ActorRoles.Filter(IsPerformer)
}
