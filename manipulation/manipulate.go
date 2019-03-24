// Package manipulation provides helper functions for resizing
// or adding watermark. This package uses https://github.com/disintegration/imaging/
// package for the actual manipulation
package manipulation

import (
	"image"

	"github.com/disintegration/imaging"
)

// CropSize struct represents the cropping dimensions
type CropSize struct {
	Height int
	Width  int
}

// These are default crops available. Custom crop can be acheive by
// providing a CropSize type.

// ThumbnailCrop default
var ThumbnailCrop CropSize

// MediumCrop default
var MediumCrop CropSize

// LargeCrop default
var LargeCrop CropSize

func init() {
	ThumbnailCrop = CropSize{
		Height: 100,
		Width:  0,
	}

	MediumCrop = CropSize{
		Height: 250,
		Width:  0,
	}

	LargeCrop = CropSize{
		Height: 500,
		Width:  0,
	}
}

// ResizeImage resizes the given image
func ResizeImage(img image.Image, crop CropSize) (*image.NRGBA, error) {
	return resize(img, crop.Width, crop.Height), nil
}

// ResizeImageMultiple resizes multiple images and returns the references of
// the resized images
func ResizeImageMultiple(imgs []image.Image, crop CropSize)([]*image.NRGBA) {
	var cropped []*image.NRGBA
	for _, image := range imgs {
		cropped = append(cropped, resize(image, crop.Width, crop.Height))
	}
	return cropped
}

func resize(img image.Image, height, width int) *image.NRGBA {
	return imaging.Resize(img, width, height, imaging.Lanczos)
}