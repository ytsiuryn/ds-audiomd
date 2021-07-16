package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestDiscMediaMarshalAndUnmarshal(t *testing.T) {
	m := MediaLP
	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, m, MediaLP)

}

func TestDiscFormatCompare(t *testing.T) {
	df1 := &DiscFormat{Media: MediaLP}
	df2 := &DiscFormat{Media: MediaLP}
	assert.Equal(t, df1.Compare(df2), 1.)
}
