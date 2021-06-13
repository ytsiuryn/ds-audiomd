package metadata

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"

	collection "github.com/gtyrin/go-collection"
	intutils "github.com/gtyrin/go-intutils"
	stringutils "github.com/gtyrin/go-stringutils"
	world "github.com/gtyrin/go-world"
)

// FileInfo describes the common file track properties.
type FileInfo struct {
	FileName string `json:"file_name,omitempty"`
	ModTime  int64  `json:"mod_time,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}

// IsEmpty проверяет коллекцию как не инициализированную.
func (fi *FileInfo) IsEmpty() bool {
	return *fi == FileInfo{}
}

// Clean сбрасывает всю коллекцию в nil, если поля структуры не отличаются от нулевых значений.
func (fi *FileInfo) Clean() {}

// Track describes the common track metadata, audio and file properties.
type Track struct {
	disc        *Disc
	Composition *Work             `json:"composition,omitempty"`
	Record      *Record           `json:"record,omitempty"`
	Position    string            `json:"position,omitempty"`
	Title       string            `json:"title,omitempty"`
	Notes       string            `json:"notes,omitempty"`
	Duration    intutils.Duration `json:"duration,omitempty"` // TODO: aggregate release track Actors?
	Actors      *Actors           `json:"actors,omitempty"`
	IDs         collection.StrMap `json:"ids,omitempty"`
	Unprocessed collection.StrMap `json:"unprocessed,omitempty"`
	*FileInfo   `json:"file_info,omitempty"`
	*AudioInfo  `json:"audio_info,omitempty"`
}

// DiscNumberByTrackPos calculate disc number from track position value.
// CD ("disk-track"): https://api.discogs.com/releases/2528044
// LP ("A,B,C,D,.."): https://api.discogs.com/releases/2373051  https://api.discogs.com/releases/10131989
// SubIndex ex.: https://api.discogs.com/releases/13452282
func DiscNumberByTrackPos(tpos string) int {
	if len(tpos) == 0 {
		return 1
	}
	flds := stringutils.SplitIntoRegularFieldsWithDelimiters(tpos, []rune{'-', '.'})
	if len(flds) == 2 {
		ret, err := strconv.Atoi(flds[0])
		if err != nil {
			return 1
		}
		return ret
	}
	diff := int(tpos[0] - 'A')
	if diff < 128 {
		return diff/2 + 1
	}
	return 1
}

// ComplexPosition ..
func ComplexPosition(position, subPosition string) string {
	return fmt.Sprintf("%s.%s", position, subPosition)
}

// ComplexTitle ..
func ComplexTitle(title, subTitle string) string {
	return fmt.Sprintf("%s. %s", title, subTitle)
}

// NewTrack creates empty object and initializes all compound fields.
func NewTrack() *Track {
	return &Track{
		Record:      NewRecord(),
		Composition: NewWork(),
		Actors:      NewActorCollection(),
		IDs:         make(map[string]string),
		Unprocessed: make(map[string]string),
		FileInfo:    &FileInfo{},
		AudioInfo:   &AudioInfo{},
	}
}

// NewFileTrack creates a new Track object for some file.
func NewFileTrack(pathName string, modTime int64) *Track {
	track := NewTrack()
	track.FileInfo = &FileInfo{FileName: pathName, ModTime: modTime}
	return track
}

// AddComment extends comment multiline string with the new one.
// Posible format is "lng_comment".
func (track *Track) AddComment(comment string) {
	var delimiter string
	if len(track.Notes) > 0 {
		delimiter = "\n"
	}
	track.Notes += fmt.Sprintf("%s%s", delimiter, comment)
}

// AddUnprocessed gatheres unsupported metadata tags info.
func (track *Track) AddUnprocessed(key, value string) {
	if utf8.ValidString(value) {
		track.Unprocessed[key] = value
	}
}

// SetLyrics save tracks's lyrics info.
func (track *Track) SetLyrics(text string, isSynchronized bool) {
	if track.Composition.Lyrics == nil {
		track.Composition.Lyrics = new(Lyrics)
	}
	track.Composition.Lyrics.Text = text
	track.Composition.Lyrics.IsSynchronized = isSynchronized
}

// SetLyricsLanguage устанавливает язык текста трека.
func (track *Track) SetLyricsLanguage(lang string) {
	if track.Composition.Lyrics == nil {
		track.Composition.Lyrics = new(Lyrics)
	} else {
		track.Composition.Lyrics.Language = world.LanguageFromString(lang)
	}
}

// SetMood is the setter for 'mood' struture field.
// func (track *Track) SetMood(moods string, album *Album) {
// 	for _, mood := range tp.SplitIntoRegularFields(moods) {
// 		track.Moods = append(track.Moods, mood)
// 		if album != nil {
// 			album.Moods[mood] = true
// 		}
// 	}
// }

// LinkWithDisc связывет трек с определенным диском.
func (track *Track) LinkWithDisc(d *Disc) {
	track.disc = d
}

// Disc возвращает ссылку на диск, связанный с треком.
func (track *Track) Disc() *Disc {
	return track.disc
}

// NormalizePosition returns a correct form of track position representation.
func NormalizePosition(position string) string {
	if len(position) == 1 && unicode.IsDigit(rune(position[0])) {
		return "0" + position
	}
	return position
}

// SetPosition ..
func (track *Track) SetPosition(position string) {
	track.Position = NormalizePosition(position)
}

// SetTitle is a setter for title track value.
func (track *Track) SetTitle(title string) {
	track.Title = title
}

// SetISRC is a setter for ISRC code of record.
func (track *Track) SetISRC(isrc string) {
	track.IDs.Add("isrc", isrc)
}

// --- Helper functions ---

// Compare a Track object with other one.
func (track *Track) Compare(other *Track) float64 {
	return stringutils.JaroWinklerDistance(track.Title, other.Title)
}

// Clean оптимизирует структуры по занимаемой памяти.
func (track *Track) Clean() {
	track.Composition.Clean()
	if track.Composition.IsEmpty() {
		track.Composition = nil
	}
	track.Record.Clean()
	if track.Record.IsEmpty() {
		track.Record = nil
	}
	track.Actors.Clean()
	if track.Actors.IsEmpty() {
		track.Actors = nil
	}
	track.IDs.Clean()
	if track.IDs.IsEmpty() {
		track.IDs = nil
	}
	track.Unprocessed.Clean()
	if track.Unprocessed.IsEmpty() {
		track.Unprocessed = nil
	}
	track.AudioInfo.Clean()
	if track.AudioInfo.IsEmpty() {
		track.AudioInfo = nil
	}
	track.FileInfo.Clean()
	if track.FileInfo.IsEmpty() {
		track.FileInfo = nil
	}
}
