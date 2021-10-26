package metadata

import (
	"github.com/ytsiuryn/go-collection"
)

// IDs - тип для хранения идентификаторов объектов внешних (online) БД.
type IDs collection.StrMap

// Merge объединяет данные двух объектов в исходном.
func (ids IDs) Merge(other IDs) {
	for k, v := range other {
		ids[k] = v
	}
}
