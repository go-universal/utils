package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"
)

// CreateThumbnail generates a thumbnail image from the provided image data.
// It resizes the image to fit within the specified maximum size while preserving
// the aspect ratio. The function supports decoding images in JPEG and PNG formats.
//
// The thumbnail is encoded in PNG format if the original image is in PNG format;
// otherwise, it is encoded in JPEG format.
func CreateThumbnail(imgData []byte, maxSize int) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, "", err
	}

	// Calculate new dimensions while preserving aspect ratio
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	if width > height {
		height = (maxSize * height) / width
		width = maxSize
	} else {
		width = (maxSize * width) / height
		height = maxSize
	}

	// Create a resized image
	thumb := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(thumb, thumb.Bounds(), img, bounds, draw.Over, nil)

	// Convert to PNG
	if format == "jpeg" {
		var buf bytes.Buffer
		err = png.Encode(&buf, thumb)
		if err != nil {
			return nil, "", err
		}
		return buf.Bytes(), format, nil
	}

	// Convert to JPEG (default)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, thumb, &jpeg.Options{Quality: 80})
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), format, nil
}
