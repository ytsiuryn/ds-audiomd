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
)

// StrToPublishingID ..
var StrToPublishingID = map[string]PublishingID{
	"Barcode": Barcode,
}

func (pid PublishingID) String() string {
	switch pid {
	case Barcode:
		return "Barcode"
	}
	return ""
}

// MarshalJSON ..
func (pid PublishingID) MarshalJSON() ([]byte, error) {
	return json.Marshal(pid.String())
}

// UnmarshalJSON получает тип PublishingID из значения JSON.
func (pid *PublishingID) UnmarshalJSON(b []byte) error {
	k := string(b)
	*pid = StrToPublishingID[k[1:len(k)-1]]
	return nil
}

// LabelID тип для перечисления идентификаторов торговых лейблов издателей во внешних БД.
type LabelID uint8

// Допустимые значения идентификаторов торговых лейблов издателей во внешних БД.
const (
	DiscogsLabelID LabelID = iota + 1
	MusicbrainzLabelID
)

// StrToLabelID ..
var StrToLabelID = map[string]LabelID{
	"DiscogsLabelID":     DiscogsLabelID,
	"MusicbrainzLabelID": MusicbrainzLabelID,
}

func (lid LabelID) String() string {
	switch lid {
	case DiscogsLabelID:
		return "DiscogsLabelID"
	case MusicbrainzLabelID:
		return "MusicbrainzLabelID"
	}
	return ""
}

// MarshalJSON преобразует значение типа идентификатора лейбла к JSON формату.
func (lid LabelID) MarshalJSON() ([]byte, error) {
	return json.Marshal(lid.String())
}

// UnmarshalJSON получает тип идентификатора лейбла из значения JSON.
func (lid *LabelID) UnmarshalJSON(b []byte) error {
	k := string(b)
	*lid = StrToLabelID[k[1:len(k)-1]]
	return nil
}

// Label содержит информацию о лейбле и номере издания в каталоле
type Label struct {
	Label string             `json:"label,omitempty"`
	Catno string             `json:"catno,omitempty"`
	IDs   map[LabelID]string `json:"ids,omitempty"`
}

// NewLabel создает объект Label.
func NewLabel(lbl, catno string) *Label {
	return &Label{Label: lbl, Catno: catno, IDs: map[LabelID]string{}}
}

// Compare сравнивает 2 лейбла по номеру в каталоге.
func (lbl *Label) Compare(other *Label) float64 {
	return stringutils.JaroWinklerDistance(lbl.Catno, other.Catno)
}

// Publishing describes trade label of the release.
type Publishing struct {
	Labels []*Label                `json:"labels,omitempty"`
	IDs    map[PublishingID]string `json:"ids,omitempty"`
}

// NewPublishing creates a new copy of ReleaseLabel object.
func NewPublishing() *Publishing {
	return &Publishing{IDs: map[PublishingID]string{}}
}

// Compare a ReleaseLabel object with other one.
func (pub *Publishing) Compare(other *Publishing) float64 {
	var res, max float64
	for _, lbl := range pub.Labels {
		for _, otherLbl := range other.Labels {
			res = lbl.Compare(otherLbl)
			if res > max {
				max = res
			}
		}
	}
	if pub.IDs[Barcode] != "" && other.IDs[Barcode] != "" {
		max *= stringutils.JaroWinklerDistance(pub.IDs[Barcode], other.IDs[Barcode])
	}
	return max
}

// IsEmpty проверяет наличие в объекте значимой информации.
func (pub *Publishing) IsEmpty() bool {
	return len(pub.IDs) == 0 && len(pub.Labels) == 0
}
