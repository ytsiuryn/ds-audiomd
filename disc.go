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

func (did MediaID) String() string {
	switch did {
	case DiscID:
		return "DiscID"
	}
	return ""
}

// MarshalJSON ..
func (did MediaID) MarshalJSON() ([]byte, error) {
	return json.Marshal(did.String())
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
	if df.Media == other.Media {
		return 1.
	}
	return 0.
}

// IsEmpty проверяет объект на пустоту.
func (df *DiscFormat) IsEmpty() bool {
	return reflect.DeepEqual(DiscFormat{}, *df)
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
