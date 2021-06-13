package metadata

import (
	"testing"
)

func TestDecodeMedia(t *testing.T) {
	results := map[string]Media{
		"LP":             MediaLP,
		"VINYL":          MediaLP,
		"SACD":           MediaSACD,
		"CD":             MediaCD,
		"[TR24][OF]":     MediaDigital,
		"[TR24][SM][OF]": MediaDigital,
		"[DSD][OF]":      MediaDigital,
		"[DXD][OF]":      MediaDigital,
		"[DVDA][OF]":     MediaDigital,
		"REEL":           MediaReeL,
	}
	for k, v := range results {
		if DecodeMedia(k) != v {
			t.Fail()
		}
	}
}

func TestDiscFormatCompare(t *testing.T) {
	df1 := &DiscFormat{Media: MediaLP}
	df2 := &DiscFormat{Media: MediaLP}
	if df1.Compare(df2) != 1. {
		t.Fail()
	}
}
