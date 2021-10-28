package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPictTypeMarshalAndUnmarshal(t *testing.T) {
	data, err := json.Marshal(PictType(0))
	require.NoError(t, err)
	assert.Equal(t, []byte(`""`), data)
	pt := PictTypeCoverFront
	data, err = json.Marshal(pt)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &pt))
	assert.Equal(t, pt, PictTypeCoverFront)
}
