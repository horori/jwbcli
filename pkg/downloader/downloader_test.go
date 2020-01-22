package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPDownload(t *testing.T) {
	urlLink := "https://gist.githubusercontent.com/samdutton/ca37f3adaf4e23679957b8083e061177/raw/e19399fbccbc069a2af4266e5120ae6bad62699a/sample.vtt"
	err := HTTPDownload(urlLink, "test_file.vtt")
	assert.NoError(t, err)
}
