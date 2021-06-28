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

// IsEmpty проверяет коллекцию на пустоту.
func (ai ActorIDs) IsEmpty() bool {
	return len(ai) == 0
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

// ActorInRoles describes album/track actor with all his roles.
// type ActorInRoles struct {
// 	Name  string            `json:"name,omitempty"`
// 	Notes string            `json:"notes,omitempty"`
// 	IDs   collection.StrMap `json:"ids,omitempty"`
// 	Roles []string          `json:"roles,omitempty"`
// }

// Actors обрабатывает коллекцию людей и их роли.
// type Actors []*ActorInRoles

// NewActorInRoles initializes a new copy of `Actor` object.
// func NewActorInRoles(name string) *ActorInRoles {
// 	return &ActorInRoles{
// 		Name:  name,
// 		IDs:   map[string]string{},
// 		Roles: []string{},
// 	}
// }

// NewActorCollection инициалирует коллекцию акторов.
// func NewActorCollection() *Actors {
// 	return (*Actors)(&[]*ActorInRoles{})
// }

// --- Используемые функции-предикаты ---

// IsPerformer определение является ли актор главным исполнителем альбома.
// func IsPerformer(air *ActorInRoles) bool {
// 	return collection.ContainsStr("performer", air.Roles)
// }

// Actor возвращает ссылку на актора по имени или nil, если отсутствует в коллекции.
// func (as *Actors) Actor(name string) *ActorInRoles {
// 	for _, air := range *as {
// 		if air.Name == name {
// 			return air
// 		}
// 	}
// 	return nil
// }

// ActorIndexByName возвращает индекс входа по имени актора.
// func (as *Actors) ActorIndexByName(name string) int {
// 	for i := range *as {
// 		if (*as)[i].Name == name {
// 			return i
// 		}
// 	}
// 	return -1
// }

// ActorByIndex возвращает данные актора в соответствии с хронологии его занесения в коллекцию.
// func (as *Actors) ActorByIndex(ind int) *ActorInRoles {
// 	return (*as)[ind]
// }

// ActorByID возвращает ссылку на объект Actor по его идектификатору во внешекй БД.
// func (as *Actors) ActorByID(dbkey, id string) *ActorInRoles {
// 	for _, air := range *as {
// 		for k, v := range air.IDs {
// 			if k == dbkey && v == id {
// 				return air
// 			}
// 		}
// 	}
// 	return nil
// }

// Filter выполняет отбор по условию для кажого элемента, указанного в функции-предикате
// и возвращает новую коллекцию.
// func (as *Actors) Filter(predicat func(elem *ActorInRoles) bool) *Actors {
// 	ret := []*ActorInRoles{}
// 	for _, air := range *as {
// 		if predicat(air) {
// 			ret = append(ret, air)
// 		}
// 	}
// 	return (*Actors)(&ret)
// }

// Roles возвращает список ролей актора. Если актор не существует, возвращает nil.
// func (as *Actors) Roles(name string) []string {
// 	if actor := as.Actor(name); actor != nil {
// 		return actor.Roles
// 	}
// 	return nil
// }

// Compare расчитывает схожесть двух коллекций акторов.
// func (as *Actors) Compare(other *Actors) float64 {
// 	if len(*as) == 0 || len(*other) == 0 {
// 		return 0.
// 	}
// 	var max, res float64
// 	for _, air := range *as {
// 		for _, otherAir := range *other {
// 			res = stringutils.JaroWinklerDistance(air.Name, otherAir.Name)
// 			if max < res {
// 				max = res
// 			}
// 		}
// 	}
// 	return max
// }

// AddActorEntry добавляет новый вход в коллекцию, если актор в ней не найден.
// func (as *Actors) AddActorEntry(name string) *ActorInRoles {
// 	actor := as.Actor(name)
// 	if actor == nil {
// 		actor = NewActorInRoles(name)
// 		*as = append(*as, actor)
// 		return actor
// 	}
// 	return actor
// }

// AddRole добавляет уникальную роль для входа с указанным именем и возвращает ссылка на актора.
// func (as *Actors) AddRole(name, role string) *ActorInRoles {
// 	actor := as.AddActorEntry(name)
// 	if !collection.Contains(role, actor.Roles) {
// 		actor.Roles = append(actor.Roles, role)
// 	}
// 	return actor
// }

// Merge создает новый вход или дополняет существующий недостающими атрибутами.
// func (as *Actors) Merge(air *ActorInRoles) {
// 	if actor := as.Actor(air.Name); actor == nil {
// 		airCopy := ActorInRoles{}
// 		airCopy = *air
// 		*as = append(*as, &airCopy)
// 	} else {
// 		if len(actor.Notes) == 0 && len(air.Notes) > 0 {
// 			actor.Notes = air.Notes
// 		}
// 		for db, id := range air.IDs {
// 			actor.IDs.Add(db, id)
// 		}
// 		for _, role := range air.Roles {
// 			if !collection.ContainsStr(role, actor.Roles) {
// 				actor.Roles = append(actor.Roles, role)
// 			}
// 		}
// 	}
// }

// DeleteRole удаляет роль определенного актора.
// func (as *Actors) DeleteRole(name, role string) error {
// 	var actor *ActorInRoles
// 	if actor = as.Actor(name); actor == nil {
// 		return errors.New("deleting a role for non existing actor")
// 	}
// 	if i := collection.Index(role, actor.Roles); i != -1 {
// 		actor.Roles = append(actor.Roles[:i], actor.Roles[i+1:]...)
// 	}
// 	return nil
// }

// DeleteActor удаляет актора из коллекции целиком.
// func (as *Actors) DeleteActor(name string) {
// 	if ind := as.ActorIndexByName(name); ind != -1 {
// 		(*as) = append((*as)[:ind], (*as)[ind+1:]...)
// 	}
// }

// IsEmpty проверяет коллекцию как не инициализированную.
// func (as *Actors) IsEmpty() bool {
// 	return len(*as) == 0
// }

// Clean сбрасывает всю коллекцию в nil, если поля структуры не отличаются от нулевых значений.
// func (as *Actors) Clean() {
// 	for i := len(*as) - 1; i >= 0; i-- {
// 		(*as)[i].IDs.Clean()
// 		if (*as)[i].IDs.IsEmpty() {
// 			(*as)[i].IDs = nil
// 		}
// 		if reflect.DeepEqual(*(*as)[i], *NewActorInRoles("")) {
// 			(*as) = append((*as)[:i], (*as)[i+1:]...)
// 		}
// 	}
// }
