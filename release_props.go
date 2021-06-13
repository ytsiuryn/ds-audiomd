package metadata

import (
	"encoding/json"
	"strings"

	collection "github.com/ytsiuryn/go-collection"
	stringutils "github.com/ytsiuryn/go-stringutils"
)

// ReleaseStatus ..
type ReleaseStatus int8

// Release status
const (
	ReleaseStatusOfficial ReleaseStatus = iota + 1
	ReleaseStatusContrafact
	ReleaseStatusBootleg
	ReleaseStatusDemonstration
	ReleaseStatusPromotion
	ReleaseStatusSampler
	ReleaseStatusUpcoming
	ReleaseStatusOuttake
)

// StrToReleaseStatus ..
var StrToReleaseStatus = map[string]ReleaseStatus{
	"oficial":       ReleaseStatusOfficial,
	"contrafact":    ReleaseStatusContrafact,
	"bootleg":       ReleaseStatusBootleg,
	"demonstration": ReleaseStatusDemonstration,
	"promotion":     ReleaseStatusPromotion,
	"sampler":       ReleaseStatusSampler,
	"upcoming":      ReleaseStatusUpcoming,
	"outtake":       ReleaseStatusOuttake,
}

// ReleaseStatusToStr ..
var ReleaseStatusToStr = map[ReleaseStatus]string{
	ReleaseStatusOfficial:      "oficial",
	ReleaseStatusContrafact:    "contrafact",
	ReleaseStatusBootleg:       "bootleg",
	ReleaseStatusDemonstration: "demonstration",
	ReleaseStatusPromotion:     "promotion",
	ReleaseStatusSampler:       "sampler",
	ReleaseStatusUpcoming:      "upcoming",
	ReleaseStatusOuttake:       "outtake",
}

// ReleaseType ..
type ReleaseType int8

// Release Type
const (
	ReleaseTypeSingle ReleaseType = iota + 1
	ReleaseTypeMaxiSingle
	ReleaseTypeMiniAlbum
	ReleaseTypeAlbum
)

// StrToReleaseType ..
var StrToReleaseType = map[string]ReleaseType{
	"single":     ReleaseTypeSingle,
	"maxisingle": ReleaseTypeMaxiSingle,
	"minialbum":  ReleaseTypeMiniAlbum,
	"album":      ReleaseTypeAlbum,
}

// ReleaseTypeToStr ..
var ReleaseTypeToStr = map[ReleaseType]string{
	ReleaseTypeSingle:     "single",
	ReleaseTypeMaxiSingle: "maxisingle",
	ReleaseTypeMiniAlbum:  "minialbum",
	ReleaseTypeAlbum:      "album",
}

// ReleaseRepeat ..
type ReleaseRepeat int8

// Release repeat
const (
	ReleaseRepeatRepress ReleaseRepeat = iota + 1
	ReleaseRepeatReissue
	ReleaseRepeatCompilation
	ReleaseRepeatDiscography
	ReleaseRepeatRemake
)

// StrToReleaseRepeat ..
var StrToReleaseRepeat = map[string]ReleaseRepeat{
	"repress":     ReleaseRepeatRepress,
	"reissue":     ReleaseRepeatReissue,
	"compilation": ReleaseRepeatCompilation,
	"discography": ReleaseRepeatDiscography,
	"remake":      ReleaseRepeatRemake,
}

// ReleaseRepeatToStr ..
var ReleaseRepeatToStr = map[ReleaseRepeat]string{
	ReleaseRepeatRepress:     "repress",
	ReleaseRepeatReissue:     "reissue",
	ReleaseRepeatCompilation: "compilation",
	ReleaseRepeatDiscography: "discography",
	ReleaseRepeatRemake:      "remake",
}

// ReleaseRemake ..
type ReleaseRemake int8

// Release remake
const (
	ReleaseRemakeRemastered ReleaseRemake = iota + 1
	ReleaseRemakeTribute
	ReleaseRemakeCover
	ReleaseRemakeRemix
)

// StrToReleaseRemake ..
var StrToReleaseRemake = map[string]ReleaseRemake{
	"remastered": ReleaseRemakeRemastered,
	"tribute":    ReleaseRemakeTribute,
	"cover":      ReleaseRemakeCover,
	"remix":      ReleaseRemakeRemix,
}

// ReleaseRemakeToStr ..
var ReleaseRemakeToStr = map[ReleaseRemake]string{
	ReleaseRemakeRemastered: "remastered",
	ReleaseRemakeTribute:    "tribute",
	ReleaseRemakeCover:      "cover",
	ReleaseRemakeRemix:      "remix",
}

// ReleaseOrigin ..
type ReleaseOrigin int8

// Release origin
const (
	ReleaseOriginStudio ReleaseOrigin = iota + 1
	ReleaseOriginLive
	ReleaseOriginRehearsal
	ReleaseOriginHome
	ReleaseOriginFieldRecording
	ReleaseOriginRadio
	ReleaseOriginTV
)

// StrToReleaseOrigin ..
var StrToReleaseOrigin = map[string]ReleaseOrigin{
	"studio":         ReleaseOriginStudio,
	"live":           ReleaseOriginLive,
	"rehearsal":      ReleaseOriginRehearsal,
	"home":           ReleaseOriginHome,
	"fieldrecording": ReleaseOriginFieldRecording,
	"radio":          ReleaseOriginRadio,
	"tv":             ReleaseOriginTV,
}

// ReleaseOriginToStr ..
var ReleaseOriginToStr = map[ReleaseOrigin]string{
	ReleaseOriginStudio:         "studio",
	ReleaseOriginLive:           "live",
	ReleaseOriginRehearsal:      "rehearsal",
	ReleaseOriginHome:           "home",
	ReleaseOriginFieldRecording: "fieldrecording",
	ReleaseOriginRadio:          "radio",
	ReleaseOriginTV:             "tv",
}

// Publishing describes trade label of the release.
type Publishing struct {
	Name  string            `json:"name,omitempty"`
	Catno string            `json:"catno,omitempty"`
	IDs   collection.StrMap `json:"ids,omitempty"`
}

// MarshalJSON ..
func (rs ReleaseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ReleaseStatusToStr[rs])
}

// UnmarshalJSON ..
func (rs *ReleaseStatus) UnmarshalJSON(b []byte) error {
	*rs = StrToReleaseStatus[string(b)]
	return nil
}

// Decode parses a string value into some enumeration value.
func (rs *ReleaseStatus) Decode(s string) {
	if val, ok := StrToReleaseStatus[strings.ToLower(s)]; ok {
		*rs = val
	}
}

// DecodeSlice ..
func (rs *ReleaseStatus) DecodeSlice(props *[]string) {
	for i := len(*props) - 1; i >= 0; i-- {
		if v, ok := StrToReleaseStatus[strings.ToLower((*props)[i])]; ok {
			*rs = v
			*props = append((*props)[:i], (*props)[i+1:]...)
			return
		}
	}
}

// MarshalJSON ..
func (rt ReleaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ReleaseTypeToStr[rt])
}

// UnmarshalJSON ..
func (rt *ReleaseType) UnmarshalJSON(b []byte) error {
	*rt = StrToReleaseType[string(b)]
	return nil
}

// DecodeSlice ..
func (rt *ReleaseType) DecodeSlice(props *[]string) {
	for i := len(*props) - 1; i >= 0; i-- {
		if v, ok := StrToReleaseType[strings.ToLower((*props)[i])]; ok {
			*rt = v
			*props = append((*props)[:i], (*props)[i+1:]...)
			return
		}
	}
}

// MarshalJSON ..
func (rr ReleaseRepeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(ReleaseRepeatToStr[rr])
}

// UnmarshalJSON ..
func (rr *ReleaseRepeat) UnmarshalJSON(b []byte) error {
	*rr = StrToReleaseRepeat[string(b)]
	return nil
}

// DecodeSlice ..
func (rr *ReleaseRepeat) DecodeSlice(props *[]string) {
	for i := len(*props) - 1; i >= 0; i-- {
		if v, ok := StrToReleaseRepeat[strings.ToLower((*props)[i])]; ok {
			*rr = v
			*props = append((*props)[:i], (*props)[i+1:]...)
			return
		}
	}
}

// MarshalJSON ..
func (rr ReleaseRemake) MarshalJSON() ([]byte, error) {
	return json.Marshal(ReleaseRemakeToStr[rr])
}

// UnmarshalJSON ..
func (rr *ReleaseRemake) UnmarshalJSON(b []byte) error {
	*rr = StrToReleaseRemake[string(b)]
	return nil
}

// DecodeSlice ..
func (rr *ReleaseRemake) DecodeSlice(props *[]string) {
	for i := len(*props) - 1; i >= 0; i-- {
		if v, ok := StrToReleaseRemake[strings.ToLower((*props)[i])]; ok {
			*rr = v
			*props = append((*props)[:i], (*props)[i+1:]...)
			return
		}
	}
}

// MarshalJSON ..
func (ro ReleaseOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(ReleaseOriginToStr[ro])
}

// UnmarshalJSON ..
func (ro *ReleaseOrigin) UnmarshalJSON(b []byte) error {
	*ro = StrToReleaseOrigin[string(b)]
	return nil
}

// DecodeSlice ..
func (ro *ReleaseOrigin) DecodeSlice(props *[]string) {
	for i := len(*props) - 1; i >= 0; i-- {
		if v, ok := StrToReleaseOrigin[strings.ToLower((*props)[i])]; ok {
			*ro = v
			*props = append((*props)[:i], (*props)[i+1:]...)
			return
		}
	}
}

// NewReleaseLabel creates a new copy of ReleaseLabel object.
func NewReleaseLabel(name string) *Publishing {
	return &Publishing{Name: name, IDs: *collection.NewStrMap()}
}

// Compare a ReleaseLabel object with other one.
func (rl *Publishing) Compare(other *Publishing) float64 {
	if rl.Catno != "" && rl.Catno != other.Catno {
		return 1.
	}
	return stringutils.JaroWinklerDistance(rl.Name, other.Name) * .99
}
