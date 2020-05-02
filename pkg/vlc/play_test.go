package vlc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayNow(t *testing.T) {
	err := PlayNow("test.mp4", "test.vtt")
	assert.NoError(t, err)
}
