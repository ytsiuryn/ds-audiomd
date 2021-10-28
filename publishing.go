package metadata

import (
	"encoding/json"

	stringutils "github.com/ytsiuryn/go-stringutils"
)

// PublishingID тип для перечисления идентификаторов публикации во внешних БД.
type PublishingID uint8

// PubIDs представляет словарь идентификаторов издателя во внешних БД.
type PubIDs map[PublishingID]string

// Допустимые значения идентификаторов публикации во внешних БД.
const (
	// он же UPC?
	PublishingBarcode PublishingID = iota + 1
)

// StrToPublishingID ..
var StrToPublishingID = map[string]PublishingID{
	"barcode": PublishingBarcode,
}

func (pid PublishingID) String() string {
	if pid == PublishingBarcode {
		return "barcode"
	}
	return ""
}

// MarshalJSON ..
func (pid PublishingID) MarshalJSON() ([]byte, error) {
	return json.Marshal(pid.String())
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
	"discogs_label_id":     DiscogsLabelID,
	"musicbrainz_label_id": MusicbrainzLabelID,
}

func (lid LabelID) String() string {
	switch lid {
	case DiscogsLabelID:
		return "discogs_label_id"
	case MusicbrainzLabelID:
		return "musicbrainz_label_id"
	}
	return ""
}

// LabelIDs представляет словарь идентификаторов лейблов во внешних БД.
type LabelIDs map[LabelID]string

// MarshalJSON преобразует словарь идентификаторов лейбла к JSON формату.
func (lbl LabelIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(lbl))
	for k, v := range lbl {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов лейбла из значения JSON.
func (lbl *LabelIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*lbl = make(LabelIDs, len(x))
	for k, v := range x {
		(*lbl)[StrToLabelID[k]] = v
	}
	return nil
}

// MarshalJSON преобразует словарь идентификаторов лейбла к JSON формату.
func (pids PubIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(pids))
	for k, v := range pids {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов издателя из значения JSON.
func (pids *PubIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*pids = make(PubIDs, len(x))
	for k, v := range x {
		(*pids)[StrToPublishingID[k]] = v
	}
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

// Compare сравнивает 2 лейбла наименованию самого лейбла и по номеру в каталоге.
func (lbl *Label) Compare(other *Label) float64 {
	return stringutils.JaroWinklerDistance(lbl.Label, other.Label) *
		stringutils.JaroWinklerDistance(lbl.Catno, other.Catno)
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

// AddLabel добавляет лейбл в данные об издании.
func (pub *Publishing) AddLabel(lbl *Label) {
	pub.Labels = append(pub.Labels, lbl)
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
	if pub.IDs[PublishingBarcode] != "" && other.IDs[PublishingBarcode] != "" {
		max *= stringutils.JaroWinklerDistance(
			pub.IDs[PublishingBarcode], other.IDs[PublishingBarcode])
	}
	return max
}

// IsEmpty проверяет наличие в объекте значимой информации.
func (pub *Publishing) IsEmpty() bool {
	return len(pub.IDs) == 0 && len(pub.Labels) == 0
}
