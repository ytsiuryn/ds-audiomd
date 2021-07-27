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

func (pt PictType) String() string {
	switch pt {
	case PictTypePNGIcon:
		return "png_icon"
	case PictTypeOtherIcon:
		return "other_icon"
	case PictTypeCoverFront:
		return "cover_front"
	case PictTypeCoverBack:
		return "cover_back"
	case PictTypeLeaflet:
		return "leaflet"
	case PictTypeMedia:
		return "media"
	case PictTypeLadArtist:
		return "lad_artist"
	case PictTypeArtist:
		return "artist"
	case PictTypeConductor:
		return "conductor"
	case PictTypeOrchestra:
		return "orchestra"
	case PictTypeComposer:
		return "composer"
	case PictTypeLyricist:
		return "lyricist"
	case PictTypeRecordingLocation:
		return "recording_location"
	case PictTypeDuringRecording:
		return "during_recording"
	case PictTypeDuringPerformance:
		return "during_performance"
	case PictTypeMovieScreen:
		return "movie_screen"
	case PictTypeBrightColorFish:
		return "bright_color_fish"
	case PictTypeIllustration:
		return "illustration"
	case PictTypeArtistLogotype:
		return "artist_logotype"
	case PictTypePublisherLogotype:
		return "publisher_logotype"
	}
	return ""
}

// MarshalJSON ..
func (pt PictType) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.String())
}

// UnmarshalJSON ..
func (pt *PictType) UnmarshalJSON(b []byte) error {
	k := string(b)
	*pt = StrToPictType[k[1:len(k)-1]]
	return nil
}
