package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/iamsayantan/chhobi/manipulation"

	"github.com/disintegration/imaging"
)

var imageName = "./images/manhattan.jpg"

func main() {
	imageFile, err := imaging.Open(imageName)
	if err != nil {
		fmt.Println("Error opening image")
		log.Fatal(err)
	}

	cropptingStart := time.Now()
	croppedImages := manipulation.ResizeMultipleCrop(imageFile, manipulation.LargeCrop, manipulation.MediumCrop, manipulation.ThumbnailCrop)
	croppingEnd := time.Now()

	diff := croppingEnd.Sub(cropptingStart)

	fmt.Println("Cropping time taken", diff.Seconds())
	for i, img := range croppedImages {
		imaging.Save(img, "./thumbs/thumb_"+strconv.Itoa(i)+".jpg")
	}
}
