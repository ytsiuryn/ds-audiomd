package metadata

// Список сервисов:
// https://www.programmableweb.com/category/lyrics/api

// Lyrics describes the text of a song track.
type Lyrics struct {
	Text string `json:"text,omitempty"`
	// 3-символьное обозначение
	Language       string `json:"language,omitempty"`
	IsSynchronized bool   `json:"is_synchronized,omitempty"`
}

// NewLyrics создает новый объект Lyrics.
func NewLyrics() *Lyrics {
	return new(Lyrics)
}

// IsEmpty проверяет структуру на пустоту.
func (l *Lyrics) IsEmpty() bool {
	return Lyrics{} == *l
}

// Clean очищает объект до неинициализированного состояния, если это возможно.
func (l *Lyrics) Clean() {
	if l.IsEmpty() {
		l = nil
	}
}
