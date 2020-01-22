package jwapi

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetVttURLByNaturalKey(t *testing.T) {
	t.Run("Get J VTT URL", func(t *testing.T) {
		// Get English list
		data, err := GetLatestVideo("E")
		m := GetLatestVideoTitles(data, 20)
		for i := 0; i < len(m); i++ {
			// Find a Japanese available video
			arr := GetAvailableLanguage(data, i)
			for _, a := range arr {
				if a == "J" {
					// Get Japanese VTT
					naturalKey := GetNaturalKey(data, i)
					t.Logf("naturalKey [ %s ]", naturalKey)
					vttURL, err := GetVttURLByNaturalKey( naturalKey, "J" )
					t.Logf("vttURL J [ %s ]", vttURL)
					assert.NoError(t, err)
					vttURL, err = GetVttURLByNaturalKey( naturalKey, "E" )
					t.Logf("vttURL E [ %s ]", vttURL)
					assert.NoError(t, err)
					return
				}
			}
		}
		assert.NoError(t, err)
	})
}
