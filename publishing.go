package metadata

import (
	"encoding/json"

	stringutils "github.com/ytsiuryn/go-stringutils"
)

// PublishingID тип для перечисления идентификаторов публикации во внешних БД.
type PublishingID uint8

// Допустимые значения идентификаторов публикации во внешних БД.
const (
	// он же UPC?
	Barcode PublishingID = iota + 1
	Catno
)

func (pid PublishingID) String() string {
	switch pid {
	case Barcode:
		return "Barcode"
	case Catno:
		return "Catno"
	}
	return ""
}

// MarshalJSON ..
func (pid PublishingID) MarshalJSON() ([]byte, error) {
	return json.Marshal(pid.String())
}

// Publishing describes trade label of the release.
type Publishing struct {
	Name  string                  `json:"name,omitempty"`
	Catno string                  `json:"catno,omitempty"`
	IDs   map[PublishingID]string `json:"ids,omitempty"`
}

// NewReleaseLabel creates a new copy of ReleaseLabel object.
func NewReleaseLabel(name string) *Publishing {
	return &Publishing{Name: name, IDs: map[PublishingID]string{}}
}

// Compare a ReleaseLabel object with other one.
func (rl *Publishing) Compare(other *Publishing) float64 {
	if rl.Catno != "" && rl.Catno != other.Catno {
		return 1.
	}
	return stringutils.JaroWinklerDistance(rl.Name, other.Name) * .99
}
