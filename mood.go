package metadata

import "encoding/json"

// Описание музыкальных настроений в:
// https://sites.tufts.edu/eeseniordesignhandbook/2015/music-mood-classification/

// Mood описывает настроение при прослушивании музыки.
type Mood int8

// Перечень допустимых значений настроений.
const (
	HappyMood Mood = iota + 1
	ExuberantMood
	EnergeticMood
	FranticMood
	AnxiousSadMood
	DepressionMood
	CalmMood
	ContentmentMood
)

// StrToMood ..
// TODO: преобразовывать разброс строковых значений к ключам StrToMood
// за счет добавления альтернативных ключей
var StrToMood = map[string]Mood{
	"happy":       HappyMood,
	"exuberant":   ExuberantMood,
	"energetic":   EnergeticMood,
	"frantic":     FranticMood,
	"anxious_sad": AnxiousSadMood,
	"depression":  DepressionMood,
	"calm":        CalmMood,
	"contentment": ContentmentMood,
}

// Moods хранит перечень настроений, характерных при прослушивании трека/альбома.
type Moods []Mood

// MoodFromName преобразовывает произвольное строковое значение к одному из допустимых констант
// типа Mood.
func MoodFromName(moodName string) Mood {
	var ret Mood
	if _, ok := StrToMood[moodName]; ok {
		ret = StrToMood[moodName]
	}
	return ret
}

func (m Mood) String() string {
	switch m {
	case HappyMood:
		return "happy"
	case ExuberantMood:
		return "exuberant"
	case EnergeticMood:
		return "energetic"
	case FranticMood:
		return "frantic"
	case AnxiousSadMood:
		return "anxious_sad"
	case DepressionMood:
		return "depression"
	case CalmMood:
		return "calm"
	case ContentmentMood:
		return "contentment"
	}
	return ""
}

// MarshalJSON ..
func (m Mood) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// UnmarshalJSON ..
func (m *Mood) UnmarshalJSON(b []byte) error {
	*m = StrToMood[string(b)]
	return nil
}
