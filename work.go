package metadata

import (
	"reflect"

	collection "github.com/ytsiuryn/go-collection"
)

// Work это часть произведения (композиция) или произведение целиком.
// Для высокоуровневых данных Position может применяться в дилогиях, трилогиях и т.д.
type Work struct {
	Parent   *Work             `json:"-"`
	Title    string            `json:"title,omitempty"`
	Position int               `json:"index,omitempty"`
	Actors   *Actors           `json:"actors,omitempty"`
	Notes    string            `json:"notes,omitempty"`
	Lyrics   *Lyrics           `json:"lyrics,omitempty"`
	IDs      collection.StrMap `json:"ids,omitempty"` // ISWC
}

// NewWork создает новый объект Composition.
func NewWork() *Work {
	return &Work{
		Actors: NewActorCollection(),
		Lyrics: NewLyrics(),
		IDs:    map[string]string{},
	}
}

// IsEmpty проверяет структуру на пустоту.
func (w *Work) IsEmpty() bool {
	return reflect.DeepEqual(Work{Parent: w.Parent}, *w)
}

// Clean сбрасывает ссылки полей в nil, если они не отличаются от нулевых значений.
func (w *Work) Clean() {
	w.Actors.Clean()
	w.Lyrics.Clean()
	w.IDs.Clean()
	if w.Actors.IsEmpty() {
		w.Actors = nil
	}
	if w.IDs.IsEmpty() {
		w.IDs = nil
	}
	if w.Lyrics.IsEmpty() {
		w.Lyrics = nil
	}
}
