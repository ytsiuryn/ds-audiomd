package metadata

import (
	"encoding/json"
	"strings"
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

func (rs ReleaseStatus) String() string {
	switch rs {
	case ReleaseStatusOfficial:
		return "oficial"
	case ReleaseStatusContrafact:
		return "contrafact"
	case ReleaseStatusBootleg:
		return "bootleg"
	case ReleaseStatusDemonstration:
		return "demonstration"
	case ReleaseStatusPromotion:
		return "promotion"
	case ReleaseStatusSampler:
		return "sampler"
	case ReleaseStatusUpcoming:
		return "upcoming"
	case ReleaseStatusOuttake:
		return "outtake"
	}
	return ""
}

// MarshalJSON ..
func (rs ReleaseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(rs.String())
}

// UnmarshalJSON ..
func (rs *ReleaseStatus) UnmarshalJSON(b []byte) error {
	k := string(b)
	*rs = StrToReleaseStatus[k[1:len(k)-1]]
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

func (rt ReleaseType) String() string {
	switch rt {
	case ReleaseTypeSingle:
		return "single"
	case ReleaseTypeMaxiSingle:
		return "maxisingle"
	case ReleaseTypeMiniAlbum:
		return "minialbum"
	case ReleaseTypeAlbum:
		return "album"
	}
	return ""
}

// MarshalJSON ..
func (rt ReleaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(rt.String())
}

// UnmarshalJSON ..
func (rt *ReleaseType) UnmarshalJSON(b []byte) error {
	k := string(b)
	*rt = StrToReleaseType[k[1:len(k)-1]]
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

func (rr ReleaseRepeat) String() string {
	switch rr {
	case ReleaseRepeatRepress:
		return "repress"
	case ReleaseRepeatReissue:
		return "reissue"
	case ReleaseRepeatCompilation:
		return "compilation"
	case ReleaseRepeatDiscography:
		return "discography"
	case ReleaseRepeatRemake:
		return "remake"
	}
	return ""
}

// MarshalJSON ..
func (rr ReleaseRepeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(rr.String())
}

// UnmarshalJSON ..
func (rr *ReleaseRepeat) UnmarshalJSON(b []byte) error {
	k := string(b)
	*rr = StrToReleaseRepeat[k[1:len(k)-1]]
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

func (rr ReleaseRemake) String() string {
	switch rr {
	case ReleaseRemakeRemastered:
		return "remastered"
	case ReleaseRemakeTribute:
		return "tribute"
	case ReleaseRemakeCover:
		return "cover"
	case ReleaseRemakeRemix:
		return "remix"
	}
	return ""
}

// MarshalJSON ..
func (rr ReleaseRemake) MarshalJSON() ([]byte, error) {
	return json.Marshal(rr.String())
}

// UnmarshalJSON ..
func (rr *ReleaseRemake) UnmarshalJSON(b []byte) error {
	k := string(b)
	*rr = StrToReleaseRemake[k[1:len(k)-1]]
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

func (ro ReleaseOrigin) String() string {
	switch ro {
	case ReleaseOriginStudio:
		return "studio"
	case ReleaseOriginLive:
		return "live"
	case ReleaseOriginRehearsal:
		return "rehearsal"
	case ReleaseOriginHome:
		return "home"
	case ReleaseOriginFieldRecording:
		return "fieldrecording"
	case ReleaseOriginRadio:
		return "radio"
	case ReleaseOriginTV:
		return "tv"
	}
	return ""
}

// MarshalJSON ..
func (ro ReleaseOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(ro.String())
}

// UnmarshalJSON ..
func (ro *ReleaseOrigin) UnmarshalJSON(b []byte) error {
	k := string(b)
	*ro = StrToReleaseOrigin[k[1:len(k)-1]]
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
