package metadata

import (
	"github.com/ytsiuryn/go-collection"
)

type IDs collection.StrMap

// Merge объединяет данные двух объектов в исходном.
func (ids IDs) Merge(other IDs) {
	for k, v := range other {
		ids[k] = v
	}
}
