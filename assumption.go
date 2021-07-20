package metadata

// Assumption хранит результат считывания метаданных из файловых треков.
type Assumption struct {
	*Release `json:"release"`
	Pictures []*PictureInAudio `json:"pictures,omitempty"`
	Actors   ActorIDs          `json:"actors,omitempty"`
}

// NewAssumption создает объект типа Assumption и возвращает ссылку на него.
func NewAssumption(release *Release) *Assumption {
	return &Assumption{
		Release:  release,
		Actors:   ActorIDs{},
		Pictures: []*PictureInAudio{}}
}

// Optimize оптимизирует исходный релиз и выносит графический материал из Release на уровень
// выше, если этот материал содержит образ картинки.
func (as *Assumption) Optimize() {
	as.Optimize()
	if as.Release == nil {
		return
	}
	as.Actors = as.Release.Actors
	as.Release.Actors = nil
	for i := len(as.Release.Pictures) - 1; i >= 0; i-- {
		if len(as.Release.Pictures[i].Data) > 0 {
			as.Pictures = append(as.Pictures, as.Release.Pictures[i])
			as.Release.Pictures = as.Release.Pictures[:i]
		}
	}
	if len(as.Release.Pictures) == 0 {
		as.Release.Pictures = nil
	}
}