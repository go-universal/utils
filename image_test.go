package utils_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/go-universal/utils"
)

func createTestImage(format string) ([]byte, error) {
	var sampleImageURL string
	switch format {
	case "jpeg":
		sampleImageURL = "https://placehold.co/600x400/jpg"
	case "png":
		sampleImageURL = "https://placehold.co/200x400/png"
	default:
		return nil, nil
	}

	resp, err := http.Get(sampleImageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch sample image: %s", resp.Status)
	}

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return imgData, nil
}

func TestCreateThumbnail_JPEG(t *testing.T) {
	imgData, err := createTestImage("jpeg")
	require.NoError(t, err, "failed to create test JPEG image")

	thumbnail, format, err := utils.CreateThumbnail(imgData, 150)
	require.NoError(t, err, "CreateThumbnail failed")

	assert.Equal(t, "jpeg", format, "expected format 'jpeg'")
	assert.NotEmpty(t, thumbnail, "thumbnail is empty")
}

func TestCreateThumbnail_PNG(t *testing.T) {
	imgData, err := createTestImage("png")
	require.NoError(t, err, "failed to create test PNG image")

	thumbnail, format, err := utils.CreateThumbnail(imgData, 150)
	require.NoError(t, err, "CreateThumbnail failed")

	assert.Equal(t, "png", format, "expected format 'png'")
	assert.NotEmpty(t, thumbnail, "thumbnail is empty")
}
