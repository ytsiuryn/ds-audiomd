package metadata

import "encoding/json"

// PictType ..
type PictType int32

// Every picture into audio track has own type.
const (
	PictTypePNGIcon PictType = iota + 1
	PictTypeOtherIcon
	PictTypeCoverFront
	PictTypeCoverBack
	PictTypeLeaflet
	PictTypeMedia
	PictTypeLadArtist
	PictTypeArtist
	PictTypeConductor
	PictTypeOrchestra
	PictTypeComposer
	PictTypeLyricist
	PictTypeRecordingLocation
	PictTypeDuringRecording
	PictTypeDuringPerformance
	PictTypeMovieScreen
	PictTypeBrightColorFish
	PictTypeIllustration
	PictTypeArtistLogotype
	PictTypePublisherLogotype
)

// PictTypeToStr - строковое значение типа изображения.
var PictTypeToStr = map[PictType]string{
	PictTypePNGIcon:           "png_icon",
	PictTypeOtherIcon:         "other_icon",
	PictTypeCoverFront:        "cover_front",
	PictTypeCoverBack:         "cover_back",
	PictTypeLeaflet:           "leaflet",
	PictTypeMedia:             "media",
	PictTypeLadArtist:         "lad_artist",
	PictTypeArtist:            "artist",
	PictTypeConductor:         "conductor",
	PictTypeOrchestra:         "orchestra",
	PictTypeComposer:          "composer",
	PictTypeLyricist:          "lyricist",
	PictTypeRecordingLocation: "recording_location",
	PictTypeDuringRecording:   "during_recording",
	PictTypeDuringPerformance: "during_performance",
	PictTypeMovieScreen:       "movie_screen",
	PictTypeBrightColorFish:   "bright_color_fish",
	PictTypeIllustration:      "illustration",
	PictTypeArtistLogotype:    "artist_logotype",
	PictTypePublisherLogotype: "publisher_logotype",
}

// StrToPictType ..
var StrToPictType = map[string]PictType{
	"png_icon":           PictTypePNGIcon,
	"other_icon":         PictTypeOtherIcon,
	"cover_front":        PictTypeCoverFront,
	"cover_back":         PictTypeCoverBack,
	"leaflet":            PictTypeLeaflet,
	"media":              PictTypeMedia,
	"lad_artist":         PictTypeLadArtist,
	"artist":             PictTypeArtist,
	"conductor":          PictTypeConductor,
	"orchestra":          PictTypeOrchestra,
	"composer":           PictTypeComposer,
	"lyricist":           PictTypeLyricist,
	"recording_location": PictTypeRecordingLocation,
	"during_recording":   PictTypeDuringRecording,
	"during_performance": PictTypeDuringPerformance,
	"movie_screen":       PictTypeMovieScreen,
	"bright_color_fish":  PictTypeBrightColorFish,
	"illustration":       PictTypeIllustration,
	"artist_logotype":    PictTypeArtistLogotype,
	"publisher_logotype": PictTypePublisherLogotype,
}

// PictureMetadata describes the common picture metadata.
type PictureMetadata struct {
	MimeType   string `json:"mime_type,omitempty"`
	Width      uint32 `json:"width,omitempty"`
	Height     uint32 `json:"height,omitempty"`
	ColorDepth uint32 `json:"color_depth,omitempty"`
	Colors     uint32 `json:"colors,omitempty"`
	Size       uint32 `json:"size,omitempty"`
}

// PictureInAudio describes some picture of the type PictType.
// TODO: реализовать методы Clean() и IsEmpty().
type PictureInAudio struct {
	*PictureMetadata `json:"pict_meta,omitempty"`
	PictType         PictType `json:"pict_type"`
	Notes            string   `json:"description,omitempty"`
	CoverURL         string   `json:"cover_url,omitempty"`
	Data             []byte   `json:"data,omitempty"`
}

// TODO: посмотреть как извлекать фото артистов из online БД.
// ActorPicture describes the special picture for an album actor.
// type ActorPicture struct {
// 	URL   string `json:"url,omitempty"`
// 	Data  []byte `json:"data,omitempty"`
// 	Notes string `json:"notes,omitempty"`
// }

// MarshalJSON ..
func (pt PictType) MarshalJSON() ([]byte, error) {
	return json.Marshal(PictTypeToStr[pt])
}

// UnmarshalJSON ..
func (pt *PictType) UnmarshalJSON(b []byte) error {
	k := string(b)
	*pt = StrToPictType[k[1:len(k)-1]]
	return nil
}
