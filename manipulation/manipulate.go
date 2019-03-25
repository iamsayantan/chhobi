// Package manipulation provides helper functions for resizing
// or adding watermark. This package uses https://github.com/disintegration/imaging/
// package for the actual manipulation
package manipulation

import (
	"image"
	"sync"

	"github.com/disintegration/imaging"
)

// CropSize struct represents the cropping dimensions
type CropSize struct {
	Height int
	Width  int
}

// cropJob stuctures individual crop operation with the original image
// and the cropping details
type cropJob struct {
	image image.Image
	crop  CropSize
}

var cropJobCh = make(chan cropJob, 20)
var cropResCh = make(chan *image.NRGBA, 20)

var noOfWorkers = 10

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
func ResizeImage(img image.Image, crop CropSize) *image.NRGBA {
	return resize(img, crop.Width, crop.Height)
}

// ResizeImageMultiple resizes multiple images and returns the references of
// the resized images
func ResizeImageMultiple(imgs []image.Image, crop CropSize) []*image.NRGBA {
	var cropped []*image.NRGBA

	// add all the images to the cropJobCh channel as cropJob for processing.
	go func(imgsToCrop []image.Image, cropSize CropSize) {
		for _, img := range imgs {
			cropjob := cropJob{image: img, crop: cropSize}
			cropJobCh <- cropjob
		}
		close(cropJobCh)
	}(imgs, crop)
	go createCropWorker(noOfWorkers)

	for croppedImage := range cropResCh {
		cropped = append(cropped, croppedImage)
	}

	return cropped
}

// createCropWorker creates a pool of cropWorker for concurrently processing tasks
func createCropWorker(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go cropWorker(&wg)
	}
	wg.Wait()
	close(cropResCh)
}

// cropWorker receives cropJobs from cropJobCh and processes them and then
// sends out the outputs to cropResCh channel.
func cropWorker(wg *sync.WaitGroup) {
	for cropjob := range cropJobCh {
		cropped := resize(cropjob.image, cropjob.crop.Width, cropjob.crop.Height)
		cropResCh <- cropped
	}
	wg.Done()
}

func resize(img image.Image, height, width int) *image.NRGBA {
	return imaging.Resize(img, width, height, imaging.Lanczos)
}
