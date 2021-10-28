package metadata

import (
	"encoding/json"
	"reflect"
	"strings"

	collection "github.com/ytsiuryn/go-collection"
)

// Media ..
type Media int8

// Release media types.
const (
	MediaSACD Media = iota + 1
	MediaCD
	MediaDigital
	MediaReeL
	MediaLP
)

// MediaID тип для перечисления идентификаторов дисков во внешних БД.
type MediaID uint8

// Допустимые значения идентификаторов дисков во внешних БД.
const (
	DiscID MediaID = iota + 1
)

// StrToMediaID ..
var StrToMediaID = map[string]MediaID{
	"disc_id": DiscID,
}

func (mid MediaID) String() string {
	switch mid {
	case DiscID:
		return "disc_id"
	}
	return ""
}

// MediaIDs представляет словарь идентификаторов медиа-дисков релиза во внешних БД.
type MediaIDs map[MediaID]string

// MarshalJSON преобразует словарь идентификаторов диска к JSON формату.
func (mids MediaIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(mids))
	for k, v := range mids {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов диска из значения JSON.
func (mids *MediaIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*mids = make(MediaIDs, len(x))
	for k, v := range x {
		(*mids)[StrToMediaID[k]] = v
	}
	return nil
}

// StrToMedia ..
var StrToMedia = map[string]Media{
	"sacd":    MediaSACD,
	"cd":      MediaCD,
	"digital": MediaDigital,
	"reel":    MediaReeL,
	"lp":      MediaLP,
}

// DiscFormat ..
type DiscFormat struct {
	Media `json:"media,omitempty"`
	Attrs []string `json:"attrs,omitempty"`
}

// Disc описывает дополнительные свойства диска. Сам номер диска указывается в объекте трека.
type Disc struct {
	Number int                `json:"number"`
	Title  string             `json:"title,omitempty"`
	Format *DiscFormat        `json:"format,omitempty"`
	IDs    map[MediaID]string `json:"ids,omitempty"`
}

// NewDisc creates and initialize a new DiscExtra object.
func NewDisc(num int) *Disc {
	return &Disc{Number: num, Format: &DiscFormat{}, IDs: make(map[MediaID]string)}
}

// DecodeMedia converts a string representation of media to a const of Media type.
func DecodeMedia(v string) Media {
	var ret Media
	upperVal := strings.ToUpper(v)
	switch {
	case collection.ContainsStr(upperVal, []string{"LP", "VINYL"}):
		ret = MediaLP
	case strings.Index(upperVal, "SACD") != -1:
		ret = MediaSACD
	case strings.Index(upperVal, "CD") != -1:
		ret = MediaCD
	case collection.ContainsStr(upperVal, []string{
		"[TR24][OF]", "[TR24][SM][OF]", "[DSD][OF]", "[DXD][OF]", "[DVDA][OF]"}):
		ret = MediaDigital
	case strings.Index(upperVal, "REEL") != -1:
		ret = MediaReeL
	}
	return ret
}

func (m Media) String() string {
	switch m {
	case MediaSACD:
		return "sacd"
	case MediaCD:
		return "cd"
	case MediaDigital:
		return "digital"
	case MediaReeL:
		return "reel"
	case MediaLP:
		return "lp"
	}
	return ""
}

// MarshalJSON преобразует значение типа медиа к JSON формату.
func (m Media) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// UnmarshalJSON получает тип медиа из значения JSON.
func (m *Media) UnmarshalJSON(b []byte) error {
	k := string(b)
	*m = StrToMedia[k[1:len(k)-1]]
	return nil
}

// Compare a DiscFormat object with other one.
func (df *DiscFormat) Compare(other *DiscFormat) float64 {
	if df != nil && other != nil && df.Media == other.Media {
		return 1.
	}
	return 0.
}

// IsEmpty проверяет объект на пустоту.
func (df *DiscFormat) IsEmpty() bool {
	return df == nil || reflect.DeepEqual(DiscFormat{}, *df)
}

// Clean сбрасывает поля структуры в nil, если поля структуры не отличаются от нулевых значений.
func (df *DiscFormat) Clean() {}

// Clean сбрасывает поля структуры в nil, если поля структуры не отличаются от нулевых значений.
func (d *Disc) Clean() {
	d.Format.Clean()
	if d.Format.IsEmpty() {
		d.Format = nil
	}
}
