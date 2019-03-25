package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/iamsayantan/chhobi/manipulation"

	"github.com/disintegration/imaging"
)

var imageName = "DSC_6119.jpg"

func main() {

	// imageFile, err := imaging.Open(imageName)

	// if err != nil {
	// 	fmt.Println("Error opening image")
	// 	log.Fatal(err)
	// }
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	var images []image.Image
	for _, f := range files {
		imageFile, err := imaging.Open("./images/" + f.Name())

		if err != nil {
			fmt.Println("Error opening image")
			log.Fatal(err)
		}
		images = append(images, imageFile)
	}

	croppedImages := manipulation.ResizeImageMultiple(images, manipulation.ThumbnailCrop)

	for i, img := range croppedImages {
		imaging.Save(img, "./thumbs/thumb_"+strconv.Itoa(i)+".jpg")
	}
	// cropped, err := manipulation.ResizeImage(imageFile, manipulation.MediumCrop)
	// if err != nil {
	// 	fmt.Println("Error Resizing image")
	// 	log.Fatal(err)
	// }

	// err = imaging.Save(cropped, "./thumbs/thumb_"+imageName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
