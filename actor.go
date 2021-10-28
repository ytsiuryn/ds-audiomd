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
	DiscogsArtistID ActorID = iota + 1
	MusicbrainzAlbumArtistID
	MusicbrainzArtistID
	MusicbrainzOriginalArtistID
)

// StrToActorID ..
var StrToActorID = map[string]ActorID{
	"discogs_artist_id":              DiscogsArtistID,
	"musicbrainz_album_artist_id":    MusicbrainzAlbumArtistID,
	"musicbrainz_artist_id":          MusicbrainzArtistID,
	"musicbrainz_original_artist_id": MusicbrainzOriginalArtistID,
}

func (aid ActorID) String() string {
	switch aid {
	case DiscogsArtistID:
		return "discogs_artist_id"
	case MusicbrainzAlbumArtistID:
		return "musicbrainz_album_artist_id"
	case MusicbrainzArtistID:
		return "musicbrainz_artist_id"
	case MusicbrainzOriginalArtistID:
		return "musicbrainz_original_artist_id"
	}
	return ""
}

// ActorIDs представляет словарь идентификаторов акторов во внешних БД.
type ActorIDs map[ActorID]string

// MarshalJSON преобразует словарь идентификаторов идентификатора актора к JSON формату.
func (aid ActorIDs) MarshalJSON() ([]byte, error) {
	x := make(map[string]string, len(aid))
	for k, v := range aid {
		x[k.String()] = v
	}
	return json.Marshal(x)
}

// UnmarshalJSON получает словарь идентификаторов актора из значения JSON.
func (aid *ActorIDs) UnmarshalJSON(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*aid = make(ActorIDs, len(x))
	for k, v := range x {
		(*aid)[StrToActorID[k]] = v
	}
	return nil
}

// ActorsIDs хранит ссылки на их коды во внешних БД.
type ActorsIDs map[ActorName]ActorIDs

// ActorRoles хранит перечень ролей акторов для определенного контекста (релиза, трека,
// композиции, записи и т.д.).
type ActorRoles map[ActorName][]ActorRole

// IsPerformer предикатная функция фильтрации исполнителей альбома.
func IsPerformer(name ActorName, roles []ActorRole) bool {
	return collection.ContainsStr("performer", roles)
}

// Add добавляет сведеления об акторе и его коде во некоторой внешней БД, если необходимо.
func (ai ActorsIDs) Add(name ActorName, key ActorID, val string) {
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
func (ai ActorsIDs) Merge(other ActorsIDs) {
	for actor, ids := range other {
		for k, v := range ids {
			ai[actor][k] = v
		}
	}
}

// IsEmpty проверяет коллекцию на пустоту.
func (ai ActorsIDs) IsEmpty() bool {
	return len(ai) == 0
}

// First возвращает первое попавшееся имя актора.
func (ai ActorsIDs) First() string {
	for actorName := range ai {
		return string(actorName)
	}
	return ""
}

// Clean сбрасывает всю коллекцию в nil, если поля структуры не отличаются от нулевых значений.
func (ai ActorsIDs) Clean() {
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
