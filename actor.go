package metadata

import (
	collection "github.com/ytsiuryn/go-collection"
	stringutils "github.com/ytsiuryn/go-stringutils"
)

// ActorName ..
type ActorName string

// Actors хранит ссылки на их коды во внешних БД.
type ActorIDs map[ActorName]IDs

// ActorRoles хранит перечень ролей акторов для определенного контекста (релиза, трека,
// композиции, записи и т.д.).
type ActorRoles map[ActorName][]string

func IsPerformer(name ActorName, roles []string) bool {
	return collection.ContainsStr("performer", roles)
}

// Add добавиляет сведеления об акторе и его коде во некоторой внешней БД, если необходимо.
func (ai ActorIDs) Add(name, key, val string) {
	ids, ok := ai[ActorName(name)]
	if !ok {
		ai[ActorName(name)] = IDs{key: val}
	} else {
		if _, ok := ids[key]; !ok {
			ai[ActorName(name)][key] = val
		}
	}
}

// Merge объединяет данные в целевой исходный объект.
func (ai ActorIDs) Merge(other ActorIDs) {
	for k, v := range other {
		ai[k].Merge(v)
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

// Добавить актора и роль, если необходимо.
func (ar ActorRoles) Add(name, role string) {
	roles := ar[ActorName(name)]
	if !collection.ContainsStr(role, roles) {
		roles = append(roles, role)
	}
	ar[ActorName(name)] = roles
}

func (ar ActorRoles) Filter(predicat func(name ActorName, roles []string) bool) ActorRoles {
	ret := ActorRoles{}
	for name, roles := range ar {
		if predicat(name, roles) {
			ret[name] = roles
		}
	}
	return ret
}

// First возвращает первый попавшийся ключ-имя или пустую строку.
func (ac ActorRoles) First() ActorName {
	for actorName := range ac {
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
