package metadata

import (
	"encoding/json"

	collection "github.com/ytsiuryn/go-collection"
	stringutils "github.com/ytsiuryn/go-stringutils"
)

// ActorName ..
type ActorName = string

// ActorRole ..
type ActorRole = string

// ActorID тип для перечисления идентификаторов актора во внешних БД.
type ActorID uint8

// Допустимые значения идентификаторов релиза во внешних БД.
const (
	MusicbrainzAlbumArtistID ActorID = iota + 1
	MusicbrainzArtistID
	MusicbrainzOriginalArtistID
)

func (aid ActorID) String() string {
	switch aid {
	case MusicbrainzAlbumArtistID:
		return "MusicbrainzAlbumArtistID"
	case MusicbrainzArtistID:
		return "MusicbrainzArtistID"
	case MusicbrainzOriginalArtistID:
		return "MusicbrainzOriginalArtistID"
	}
	return ""
}

// MarshalJSON ..
func (aid ActorID) MarshalJSON() ([]byte, error) {
	return json.Marshal(aid.String())
}

// ActorIDs хранит ссылки на их коды во внешних БД.
type ActorIDs map[ActorName]map[ActorID]string

// ActorRoles хранит перечень ролей акторов для определенного контекста (релиза, трека,
// композиции, записи и т.д.).
type ActorRoles map[ActorName][]ActorRole

// IsPerformer предикатная функция фильтрации исполнителей альбома.
func IsPerformer(name ActorName, roles []ActorRole) bool {
	return collection.ContainsStr("performer", roles)
}

// Add добавляет сведеления об акторе и его коде во некоторой внешней БД, если необходимо.
func (ai ActorIDs) Add(name ActorName, key ActorID, val string) {
	ids, ok := ai[name]
	if !ok {
		ai[name] = map[ActorID]string{key: val}
	} else {
		if _, ok := ids[key]; !ok {
			ai[name][key] = val
		}
	}
}

// Merge объединяет данные в целевой исходный объект.
func (ai ActorIDs) Merge(other ActorIDs) {
	for actor, ids := range other {
		for k, v := range ids {
			ai[actor][k] = v
		}
	}
}

// IsEmpty проверяет коллекцию на пустоту.
func (ai ActorIDs) IsEmpty() bool {
	return len(ai) == 0
}

// First возвращает первое попавшееся имя актора.
func (ai ActorIDs) First() string {
	for actorName := range ai {
		return string(actorName)
	}
	return ""
}

// Clean сбрасывает всю коллекцию в nil, если поля структуры не отличаются от нулевых значений.
func (ai ActorIDs) Clean() {
	for name, ids := range ai {
		if len(ids) == 0 {
			delete(ai, name)
		}
	}
}

// Compare сравнивает объект с аналогичным и определяет его степень схожести в числовом
// выражении.
func (ar ActorRoles) Compare(other ActorRoles) float64 {
	if len(ar) == 0 || len(other) == 0 {
		return 0.
	}
	var max, res float64
	for name := range ar {
		for otherName := range other {
			res = stringutils.JaroWinklerDistance(string(name), string(otherName))
			if max < res {
				max = res
			}
		}
	}
	return max
}

// Add добавляет актора и роль, если необходимо.
func (ar ActorRoles) Add(name ActorName, role ActorRole) {
	roles := ar[name]
	if !collection.ContainsStr(role, roles) {
		roles = append(roles, role)
	}
	ar[name] = roles
}

// Filter фильтрует коллекцию акторов с определенной функцией-предикатом.
func (ar ActorRoles) Filter(predicat func(name ActorName, roles []ActorRole) bool) ActorRoles {
	ret := ActorRoles{}
	for name, roles := range ar {
		if predicat(name, roles) {
			ret[name] = roles
		}
	}
	return ret
}

// First возвращает первый попавшийся ключ-имя или пустую строку.
func (ar ActorRoles) First() string {
	for actorName := range ar {
		return actorName
	}
	return ""
}

// IsEmpty проверяет коллекцию как не инициализированную.
func (ar ActorRoles) IsEmpty() bool {
	return len(ar) == 0
}

// Clean удаляет пары, где список ролей пуст.
func (ar ActorRoles) Clean() {
	for name, roles := range ar {
		if len(roles) == 0 {
			delete(ar, name)
		}
	}
}
